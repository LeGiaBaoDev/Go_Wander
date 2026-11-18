# ADR-0005 — Thuật toán tìm đường: lọc đúng luật, đa tiêu chí Pareto, bốn nhánh

- **Trạng thái:** Proposed — chờ chủ dự án chốt
- **Ngày:** 18-11-2026
- **Người quyết định:** chủ dự án (chờ chốt)
- **Liên quan:** `ADR-0003-api-contract-auth-rbac.md`, `ADR-0004-data-model-postgis.md`, `../00-foundation/02-scope.md`, `../04-development/02-conventions.md`, `../01-product/07-design-canvas.md`, canvas `RouteLogic`/`RouteBranches`

## Bối cảnh

R1 chỉ tìm tuyến trong **TP.HCM**, cho bốn mode: đi bộ, bus, metro L1, xe đạp. Ngoài TP.HCM chỉ xem bản đồ/địa điểm; tab tuyến trả `UNSUPPORTED`, không bịa route (`02-scope.md`).

Sản phẩm yêu cầu:

- **Bốn nhánh** khi chọn điểm đến: Gợi ý / Nhanh / Gần / Rẻ (canvas `RouteBranches`).
- **Lọc “đúng luật” là ràng buộc cứng, chạy trước khi tính** (chủ dự án chốt; canvas `RouteLogic`; `02-conventions.md`).
- Giải thích vì sao chọn tuyến, có `dataStatus`/nhãn `ASSUMED`; không quảng cáo “đúng luật tuyệt đối” khi dữ liệu chưa đủ.
- Không để quảng cáo/khuyến mãi đổi thứ hạng (nguyên tắc 3).
- Kết quả tái lập được với cùng đầu vào; có test case cố định.

Bản export canvas đọc ngày 18-11-2026 (`README.md` ghi version `1794999156-c9a6`, 144 màn) xác nhận:

- `RouteLogic` dùng **3 tiêu chí Pareto**: thời gian, quãng đường, tiền; **Gợi ý** = 0,5·thời gian + 0,3·tiền + 0,2·số lần đổi xe (chuẩn hoá 0–1, `ASSUMED`).
- **Gần nhất** = quãng đường nhỏ nhất, thời gian **không quá 1,5 lần** nhanh nhất; **Rẻ nhất** = tiền nhỏ nhất, thời gian **không quá 2 lần** nhanh nhất.
- Đi bộ tới ga/trạm trong bán kính **800 m**, bắt đầu từ **3 điểm gần nhất**; mục tiêu phản hồi **< 2 giây**, có cache theo ô lưới + khung giờ.
- Canvas ghi “Mưa: cộng phạt đoạn đi bộ ngoài trời”; đây là hành vi **R2** (weather-aware), không phải R1.
- Badge trên `RouteBranches` là “✓ Đúng luật” kèm ghi chú “Giờ, giá: ước tính”; ADR này giữ badge đó nhưng bắt buộc kèm điều kiện dữ liệu (xem 3.3).

Dữ liệu hiện có:

- **OSM** cho đường đi bộ/xe đạp, ODbL; graph cắt theo khung TP.HCM.
- **GTFS Metro L1 giả định** và **GTFS bus dựng từ OSM giả định**, nhãn `ASSUMED`; chưa có GTFS/GTFS-RT chính thức, chưa có giá vé trong feed cũ.
- **OTP** là engine đã chọn (`02-containers.md`); repo cũ đã chạy OTP 2.10.0 với graph TP.HCM và 10 route case (`ADR-020`, `ADR-023`, `ADR-032`).
- Chưa có dữ liệu giao thông thời gian thực, chưa có turn-by-turn traffic; ô tô/xe máy ngoài R1.

Các câu hỏi phải chốt:

1. Tự viết RAPTOR/A* trong Go hay dùng OTP?
2. Lọc luật ở đâu, xử lý dữ liệu thiếu thế nào?
3. Đa tiêu chí tính ra sao, bốn nhánh khác nhau ở điểm nào?
4. Xử lý giá vé/giờ giả định và coverage như thế nào?
5. Làm sao để “Gợi ý” không bịa và “Rẻ” không sai khi thiếu giá?

## Quyết định

### 1. Dùng OTP làm engine, Go API là lớp điều phối

- **Không tự viết engine tìm đường trong Go ở R1.** Dùng **OpenTripPlanner 2.x** (phiên bản khoá ở scaffold; repo cũ dùng 2.10.0):
  - **Transit:** RAPTOR (multi-criteria) cho bus/metro, theo OTP.
  - **Đường bộ:** A* trên graph OSM cho đi bộ/xe đạp, theo OTP.
- Go API chịu trách nhiệm: validate request, kiểm coverage, chuẩn hoá toạ độ, gọi OTP, **lọc luật trước khi tính**, **hậu kiểm** kết quả, tính Pareto/chọn nhánh, gắn nhãn/giải thích, trả response.
- OTP không lộ ra client; client chỉ thấy schema của Go API. Đổi engine sau này là thay adapter `RoutingProvider`, không đổi hợp đồng.
- Lý do: OTP đã được chứng minh với graph TP.HCM ở repo cũ; tự viết RAPTOR + A* + turn restrictions + GTFS là quá lớn cho R1 và dễ sai luật hơn.

### 2. Pipeline cố định: lọc trước, tính sau, hậu kiểm

```text
Request
  → 1. Validate + chuẩn hoá (toạ độ, thời gian, mode, locale)
  → 2. Kiểm coverage (polygon + mode)         → UNSUPPORTED nếu ngoài
  → 3. Snap origin/destination vào graph      → lỗi rõ nếu quá xa/không snap
  → 4. Lọc luật cứng trên graph/request       → loại cạnh/chuyến sai luật
  → 5. Gọi OTP (RAPTOR + A*)                  → candidates
  → 6. Hậu kiểm luật trên từng leg             → loại route vi phạm
  → 7. Chuẩn hoá + Pareto + chọn 4 nhánh
  → 8. Gắn dataStatus/legality/reasons        → response
```

Quy tắc:

- **Không đảo thứ tự.** Lọc luật không phải “phạt” rồi vẫn trả route; cạnh sai luật bị loại khỏi graph/không được đi.
- **Transit access:** đi bộ tới ga/trạm tối đa **800 m**, bắt đầu từ **3 điểm gần nhất** (canvas `RouteLogic`; `ASSUMED`); nếu không có trạm trong bán kính, route transit không khả dụng nhưng walk/bike vẫn có thể có.
- **Mode R1:** walk, bus, metro, bicycle. Xe máy/ô tô dùng A* về mặt kỹ thuật nhưng **không mở ở R1**; không nhận mode ngoài danh sách.
- **Weather:** R1 không có weather; “cộng phạt đi bộ ngoài trời” là R2, không đưa vào R1.
- Nếu một ràng buộc không thể lọc (dữ liệu thiếu), route có thể vẫn trả nhưng `legality.status=PARTIAL/UNKNOWN` và UI phải nói rõ; **không** gắn nhãn “đúng luật”.
- Hậu kiểm là lưới an toàn: nếu OTP trả leg vi phạm, **bỏ route**, ghi log mã `ROUTING_LEGALITY_VIOLATION`, không trả cho client. Nếu tất cả bị bỏ → `NO_ROUTE` kèm lý do.
- Mọi bước có `deadline` chung; hết giờ trả `ROUTING_TIMEOUT` thay vì treo.

### 3. Lọc luật cứng

#### 3.1 Hai tầng lọc

| Tầng | Khi nào | Nội dung |
|---|---|---|
| **Graph build** | Khi dựng graph OTP | Loại cạnh/chuyến sai luật khỏi graph; cấu hình profile theo mode; kiểm OSM/GTFS |
| **Request** | Mỗi request | Ràng buộc người dùng (tránh bậc thang, xe lăn, max walk) và mode; nếu OTP không hỗ trợ thì trả `UNSUPPORTED` thay vì bỏ qua |

#### 3.2 Bảng luật tối thiểu

| Dữ liệu OSM/GTFS | Hành động | Ghi chú |
|---|---|---|
| `access=no/private` | Loại | Trừ `access=destination` khi điểm đi/đến nằm trong cùng way/area; R1 có thể trả `NO_ROUTE` nếu không xác định được |
| `access=customers` | Loại | Không coi là đường công cộng |
| `foot=no` | Loại với đi bộ | Nếu lối đi duy nhất, trả `NO_ROUTE`, không đi xuyên |
| `bicycle=no` | Loại với xe đạp | Kể cả `highway=steps` nếu không có `bicycle=yes` |
| `highway=construction/proposed` | Loại | Không đi vào đường chưa tồn tại |
| `highway=steps` | Đi bộ: cho; xe đạp: loại trừ khi `bicycle=yes` | Không mô hình hoá dắt bộ ở R1 |
| `oneway=yes` | Xe đạp: tôn trọng; đi bộ: không áp dụng | `oneway:bicycle=no` được ưu tiên |
| `barrier=gate/bollard` + `access=no` | Loại | `bicycle=yes`/`foot=yes` được xét theo tag cụ thể |
| `route=ferry` | Loại | Phà ngoài R1 |
| Turn restriction relations | OTP xử lý | Nếu relation lỗi/không đọc được → `PARTIAL` |
| `access:conditional`, `opening_hours` phức tạp | Không đánh giá tự động ở R1 | Gắn `PARTIAL` nếu route đi qua; UI nhắc “chưa xác minh giờ/điều kiện” |
| Đóng đường tạm thời | Không có dữ liệu realtime | Không khẳng định; feedback người dùng là kênh bổ sung |
| GTFS `calendar`/`frequencies` | Chỉ dùng chuyến có trong ngày/khung giờ | Ngoài giờ → `DATA_UNAVAILABLE` hoặc route không transit |
| GTFS `wheelchair_accessible` | Nếu user yêu cầu xe lăn | Thiếu dữ liệu → `PARTIAL`/`UNSUPPORTED`, không đoán |

- Luật chi tiết theo profile (walk/bike/transit) nằm trong config có tên, không rải trong code.
- Không dùng `access=private` cho lối tắt; nếu điểm đến nằm trong khu riêng, snap về điểm công cộng gần nhất và ghi `DESTINATION_ACCESS_RESTRICTED` (đề xuất; cần test với dữ liệu thật).
- Không tự “sửa” OSM bằng tay trong graph; mọi sửa phải là script/import có ghi vết và rebuild graph.

#### 3.3 Trạng thái luật

`legality.status`:

| Giá trị | Nghĩa | UI |
|---|---|---|
| `FILTERED` | Đã áp dụng mọi luật cứng đã biết; không có tag thiếu/lỗi liên quan trên route | Badge “Đúng luật” + helper “Theo dữ liệu bản đồ hiện có” |
| `PARTIAL` | Có cạnh/chuyến thiếu dữ liệu hoặc tag điều kiện không đánh giá được | “Một số quy định chưa xác minh” (không dùng badge xanh) |
| `UNKNOWN` | Không đủ dữ liệu để kết luận | “Chưa xác minh được luật” |

- Badge “Đúng luật” của canvas chỉ dùng khi `FILTERED`; helper/tooltip phải nói rõ “theo dữ liệu bản đồ hiện có”. `FILTERED` **không** có nghĩa “đúng luật tuyệt đối”; OSM là dữ liệu cộng đồng, có thể thiếu/sai.
- `PARTIAL`/`UNKNOWN` không được hiện badge “Đúng luật”; nếu thiết kế muốn giữ badge, phải đổi trạng thái/thông điệp trước khi code.
- `legality.unverifiedRules[]` liệt kê lý do (ví dụ `CONDITIONAL_ACCESS`, `TURN_RESTRICTION_DATA`, `FERRY_AVOIDED`).
- Nếu route vi phạm sau hậu kiểm: không trả; log mã riêng; không “hạ nhãn” để vẫn trả.

### 4. Đa tiêu chí và Pareto

#### 4.1 Vector tiêu chí

Mỗi route candidate có:

- **3 tiêu chí Pareto:** `duration_s`, `distance_m`, `fare_vnd`.
- **Metric phụ:** `walk_m`, `transfers`, `waiting_s` — dùng cho hiển thị, tie-break và điểm “Gợi ý”, không tham gia Pareto.

Trong đó:

- `duration_s`: tổng thời gian, gồm chờ/chuyển tuyến.
- `distance_m`: tổng quãng đường của route (không phải đường chim bay).
- `fare_vnd`: tổng giá vé; có thể `unknown`.
- `walk_m`: tổng quãng đường đi bộ (kể cả chuyển tuyến).
- `transfers`: số lần chuyển tuyến (không tính đi bộ nối).

#### 4.2 Pareto

- Route A **chi phối** B nếu A không tệ hơn B ở cả 3 tiêu chí và tốt hơn ở ít nhất một.
- **Pareto trên 3 tiêu chí** (`duration_s`, `distance_m`, `fare_vnd`); route bị chi phối bị loại. Nếu `fare` unknown, bỏ `fare` khỏi so sánh của route đó và gắn cờ; không coi unknown là 0.
- Loại trùng: cùng chuỗi leg + giờ lên/xuống gần nhau (≤ 60 giây) → giữ route tốt hơn.
- Giới hạn **tối đa 12 candidate** sau Pareto để chọn nhánh; không trả quá 4 nhánh chính.

#### 4.3 Chọn bốn nhánh

| Nhánh (`kind`) | Mục tiêu | Điều kiện | Tie-break | Nhãn UI | Khi không có |
|---|---|---|---|---|---|
| `recommended` | Điểm tổng hợp thấp nhất | Không | Ít chuyển tuyến, ít đi bộ | Gợi ý | Không có route → `NO_ROUTE` |
| `fastest` | `duration_s` nhỏ nhất | Không | Ít chuyển tuyến, ít đi bộ | Nhanh | Không có candidate → `available:false` |
| `shortest` | `distance_m` nhỏ nhất | `duration_s ≤ 1,5 × fastest.duration_s` | Ít đi bộ, thời gian | Gần | Không có candidate trong cap → `available:false`, lý do `TIME_CAP_1_5` |
| `cheapest` | `fare_vnd` nhỏ nhất | `duration_s ≤ 2 × fastest.duration_s` | Thời gian, ít chuyển tuyến | Rẻ | Không có giá/candidate trong cap → `available:false`, lý do `FARE_UNKNOWN`/`TIME_CAP_2_0` |

- **“Gần” = tổng quãng đường ngắn nhất** trong cap 1,5× nhanh nhất, không phải “ít đi bộ”; walking chỉ là tie-break/hiển thị. Ghi rõ trong UI/response để tránh hiểu nhầm.
- **Bỏ trùng theo route** (đúng canvas `RouteLogic`): nếu cùng một route thắng nhiều nhánh, `routes[]` chỉ có một route và `branches[]` có thể tham chiếu cùng `routeId` với nhiều nhãn; client gộp badge trên một thẻ. Nếu muốn 4 thẻ riêng, lấy candidate kế tiếp trong cap; nếu không còn, nhánh đó `available:false`.
- Không bịa nhánh: nếu không có candidate cho “Rẻ”, trả `available:false` + lý do; không lấy route đắt gán nhãn “Rẻ”.
- Không để ads/affiliate/khuyến mãi tham gia chọn nhánh hoặc đổi thứ tự.

#### 4.4 Điểm “Gợi ý”

- Chuẩn hoá min-max từng tiêu chí trong tập candidate:
  - `timeScore = (duration - min) / (max - min)`
  - `fareScore = (fare - min) / (max - min)`
  - `transferScore = (transfers - min) / (max - min)`
- Công thức đúng canvas `RouteLogic`: **0,5 · timeScore + 0,3 · fareScore + 0,2 · transferScore** (trọng số `ASSUMED`, chờ chốt).
- Nếu `fare` unknown: bỏ term giá, chuẩn hoá lại trọng số còn lại (`time`/`transfers` theo tỉ lệ 0,5/0,2), gắn cờ `FARE_UNKNOWN`; route vẫn có thể được “Gợi ý” nhưng không được gán “Rẻ”.
- Nếu mọi candidate bằng nhau ở một tiêu chí, term = 0; không chia cho 0.
- Không cộng thêm “phạt chuyển tuyến 180 giây” vào thời gian khi tính điểm; số lần chuyển tuyến đã là term riêng. `walk_m` dùng tie-break/hiển thị, không đưa vào công thức Gợi ý.
- Điểm chỉ dùng để chọn “Gợi ý”; không hiển thị như “điểm chất lượng” tuyệt đối.

### 5. Dữ liệu, graph và coverage

#### 5.1 Coverage

- `mobility.coverage_area` (ADR-0004) là nguồn sự thật: polygon + modes + `data_status`.
- Kiểm **trước** khi gọi OTP:
  - Điểm đi/đến ngoài polygon → `UNSUPPORTED` (`details.reason=OUTSIDE_COVERAGE`).
  - Mode không có trong `modes` → `UNSUPPORTED` (`MODE_NOT_SUPPORTED`).
  - R1: polygon TP.HCM; khung OSM cắt thô có thể rộng hơn polygon, nhưng API chỉ nhận trong polygon.
- Không gộp graph toàn quốc vào graph TP.HCM; vùng mới mở sau, từng vùng một.

#### 5.2 OSM

- Nguồn: extract OSM (Geofabrik hoặc nguồn tương đương), cắt theo khung TP.HCM; ghi ngày extract, checksum, ODbL/“© OpenStreetMap contributors”.
- Graph build loại cạnh sai luật theo bảng mục 3; cấu hình profile walk/bike; không dùng dữ liệu traffic.
- Sửa OSM bằng script/import có ghi vết; không sửa file graph bằng tay.

#### 5.3 GTFS transit

- Metro L1: feed giả định, `ASSUMED`; phần có nguồn (giờ đầu/cuối, số lượt, giãn cách) và phần ước tính (giờ từng chặng) ghi trong `transit_feed.notes`.
- Bus: feed dựng từ OSM, `ASSUMED`; tốc độ 15 km/h là giả định của repo cũ, **chưa có nguồn**; nếu dùng, ghi rõ.
- Giá vé: feed cũ **không có** fare. R1 dùng **bảng giá cấu hình `ASSUMED`** cho metro/bus (có version, nguồn/ghi chú), hoặc `UNKNOWN` nếu thiếu; UI ghi “Giá ước tính”. Không gọi là giá chính thức.
- Không có GTFS-RT ở R1; `realtime=false`; không hứa giờ thực.
- Mỗi graph build ghi `graphVersion`, `otpVersion`, `osmDate`, feed IDs/checksums, `configHash`; response trả `graphVersion` để tái lập.

#### 5.4 Ngưỡng và trọng số (tất cả là đề xuất)

| Tham số | Giá trị khởi điểm | Nhãn | Ghi chú |
|---|---|---|---|
| Tốc độ đi bộ | 4,5 km/h | `ASSUMED` | Cấu hình; OTP có thể dùng tốc độ riêng |
| Tốc độ xe đạp | 12 km/h | `ASSUMED` | Cấu hình |
| Tốc độ bus | 15 km/h | `ASSUMED` | Kế thừa ADR-032; không có nguồn |
| Phạt chuyển tuyến | 180 s | `ASSUMED` | Chỉ dùng tie-break/giải thích; không cộng vào công thức Gợi ý |
| Snap tối đa | 300 m | `ASSUMED` | Quá xa → `ORIGIN_NOT_SNAPPABLE`/`DESTINATION_NOT_SNAPPABLE` |
| Bán kính đi bộ tới trạm/ga | 800 m | `ASSUMED` | Canvas `RouteLogic`; bắt đầu từ 3 điểm gần nhất |
| Cap “Gần” / “Rẻ” | 1,5× / 2× nhanh nhất | `ASSUMED` | Canvas `RouteLogic`; ngoài cap → nhánh không khả dụng |
| Bán kính nearby | 2 km | `ASSUMED` | Không liên quan routing |
| Timeout OTP | 8 s | Đề xuất | API cắt trước OTP |
| Mục tiêu p95 | < 2 s; gate R1 ≤ 3 s | Canvas / vision | Đo trước pilot; 2–3 s ghi nhận và tối ưu |
| Số candidate tối đa | 12 | Đề xuất | Sau Pareto |
| Số nhánh trả | 4 | Sản phẩm | Có thể ít hơn nếu thiếu/trùng |
| Trọng số Gợi ý | 0,5 thời gian / 0,3 tiền / 0,2 chuyển tuyến | `ASSUMED` | Canvas RouteLogic |

- Tất cả tham số nằm trong config có tên/version; đổi phải có change record và chạy lại test.
- Không hard-code trong handler; không để client gửi trọng số tùy ý (chỉ nhận constraint an toàn).
- **Cache:** R1 không cache route theo người dùng. Nếu cần đạt < 2 s, chỉ cache/precompute theo **ô lưới + khung giờ** cho cặp OD phổ biến, không định danh, TTL ngắn; làm khi có ADR/đo đạc, không mặc định.
- **Weather:** “cộng phạt đoạn đi bộ ngoài trời” là R2; R1 không nhận weather làm input và không đổi thứ hạng vì thời tiết.

### 6. Hợp đồng response

```json
{
  "requestId": "req_01J...",
  "graphVersion": "hcmc-2026-11-18-a1b2c3",
  "dataStatus": "ASSUMED",
  "coverage": { "status": "SUPPORTED", "areaId": "hcmc" },
  "routes": [
    {
      "id": "rt_...",
      "metrics": {
        "durationS": 1840,
        "distanceM": 8200,
        "walkM": 950,
        "fareVnd": 15000,
        "fareStatus": "ASSUMED",
        "transfers": 1,
        "waitingS": 240
      },
      "legs": [
        {
          "mode": "WALK|BUS|METRO|BICYCLE",
          "from": { "name": "...", "lat": 10.77, "lon": 106.69 },
          "to": { "name": "...", "lat": 10.78, "lon": 106.70 },
          "startAt": "2026-11-18T09:00:00Z",
          "endAt": "2026-11-18T09:08:00Z",
          "distanceM": 650,
          "geometry": "encoded-polyline",
          "routeName": "L1",
          "realtime": false
        }
      ],
      "legality": {
        "status": "FILTERED|PARTIAL|UNKNOWN",
        "unverifiedRules": []
      }
    }
  ],
  "branches": [
    {
      "kind": "recommended|fastest|shortest|cheapest",
      "routeId": "rt_...",
      "available": true,
      "score": 0.42,
      "reasons": [
        { "code": "BEST_OVERALL", "params": {} },
        { "code": "LEGAL_FILTER_APPLIED", "params": {} },
        { "code": "DATA_ASSUMED", "params": {} }
      ]
    }
  ],
  "assumptions": ["FARE_ASSUMED", "SCHEDULE_ASSUMED"]
}
```

- `routes[]` chứa route duy nhất; `branches[]` tham chiếu `routeId` (có thể cùng route nhiều nhãn).
- `branches[].available=false` có `reason` (ví dụ `FARE_UNKNOWN`, `NO_FASTEST_CANDIDATE`); không bịa route.
- `reasons[].code` là enum ổn định, client dịch; không trả câu giải thích tự do từ LLM.
- `legality` và `dataStatus` bắt buộc; `assumptions` liệt kê yếu tố giả định.
- Không trả toạ độ chính xác của người dùng trong log; response chỉ trả điểm đi/đến người dùng đã nhập.

### 7. Lỗi

| Mã | HTTP | Khi nào |
|---|---|---|
| `VALIDATION_FAILED` | 400 | Toạ độ/thời gian/mode sai |
| `UNSUPPORTED` | 422 | Ngoài coverage/mode/constraint; `details.reason` |
| `NO_ROUTE` | 422 | Trong coverage nhưng không có route thỏa luật/ràng buộc |
| `DATA_UNAVAILABLE` | 422 | Ngoài giờ chạy/không có lịch |
| `ORIGIN_NOT_SNAPPABLE` / `DESTINATION_NOT_SNAPPABLE` | 422 | Điểm quá xa đường/không snap được; `details.maxSnapM` |
| `ROUTING_TIMEOUT` | 504 | Quá deadline |
| `ROUTING_UNAVAILABLE` | 503 | OTP lỗi/không kết nối |
| `ROUTING_LEGALITY_VIOLATION` | 500 | OTP trả route vi phạm sau hậu kiểm; log nội bộ, không trả chi tiết |
| `RATE_LIMITED` | 429 | Vượt giới hạn |

- `UNSUPPORTED` khác `NO_ROUTE`: ngoài vùng/mode không hỗ trợ vs trong vùng nhưng không có tuyến.
- Không trả `200` với mảng rỗng để che lỗi.
- Nếu OTP trả route kèm cảnh báo (ví dụ đi bộ nhanh hơn transit), vẫn trả route + reason; không tự loại.

### 8. Hiệu năng, cache, tái lập

- Mục tiêu p95 **< 2 giây** (canvas `RouteLogic`); gate R1 **≤ 3 giây** (vision). Nếu p95 2–3 giây, ghi nhận và tối ưu trước pilot; quá 3 giây không qua gate. Timeout API **8 giây** (OTP timeout nhỏ hơn).
- Một request thường **một lần gọi OTP** với `numItineraries` đủ lớn (ví dụ 12) + cửa sổ tìm kiếm hợp lý. Chỉ gọi lần hai có mục tiêu (ví dụ tối ưu ít chuyển) nếu lần đầu thiếu nhánh và còn ngân sách; tối đa 2 lần gọi/request. *(Đề xuất; đo trước khi chốt.)*
- **Không cache route cá nhân ở R1.** Nếu cần đạt < 2 giây, chỉ cache/precompute theo ô lưới + khung giờ cho cặp OD phổ biến, không định danh, TTL ngắn; không lưu toạ độ thô. Làm khi có đo đạc/ADR, không mặc định.
- Tái lập: cùng `graphVersion` + cùng request + cùng config → cùng tập candidate/thứ tự (tie-break xác định). Ghi `graphVersion`/`configHash` trong response và test.
- OTP chạy cùng host pilot; graph TP.HCM; không K8s.

### 9. Kiểm thử và bằng chứng

- **Unit:** bảng luật (mỗi tag → loại/cho), Pareto (chi phối/không), chuẩn hoá điểm, chọn nhánh, fare unknown, tie-break.
- **Integration:** graph nhỏ cố định (OSM + GTFS fixture) có đáp án; chạy OTP thật trong CI hoặc nightly.
- **Route case:** kế thừa 10 case repo cũ (RC-01…RC-10) làm nền; thêm case:
  - đi bộ qua đường cấm/không có lối đi;
  - xe đạp vào đường cấm/ bậc thang;
  - một chiều;
  - ngoài coverage (Hà Nội) → `UNSUPPORTED`;
  - origin/destination quá xa đường → snap error;
  - không có tuyến trong giờ → `DATA_UNAVAILABLE`;
  - 4 nhánh khác nhau đúng chỗ; nhánh không có → `available:false`.
- **Hậu kiểm:** cố tình chèn route vi phạm trong fixture → API phải bỏ, không trả.
- **Hiệu năng:** đo p50/p95 trên máy pilot, ghi môi trường; chưa đạt thì không qua gate R1.
- **Bằng chứng:** log/test output/lệnh chạy trong change record; không ghi “đã test” nếu chưa chạy.

### 10. Riêng tư

- Không lưu route/legs/origin/destination; không log toạ độ chính xác.
- Nếu cần debug, log `requestId`, `graphVersion`, mode, khoảng cách làm tròn; không log địa chỉ/tên điểm.
- Feedback “báo sai tuyến” chỉ gửi leg/route ID + mô tả, không gửi vị trí hiện tại mặc định.
- Không dùng dữ liệu route để huấn luyện/ads khi chưa có consent và ADR riêng.

## Lý do

1. **OTP đã chứng minh** với OSM/GTFS TP.HCM ở repo cũ; tự viết RAPTOR/A* trong R1 tăng rủi ro và chi phí, không cần để có 4 nhánh.
2. **Lọc trước + hậu kiểm** là cách duy nhất để vừa giữ “đúng luật là ràng buộc cứng” vừa không trả route sai khi OTP/dữ liệu có lỗi.
3. **Pareto** cho phép chọn 4 nhánh từ cùng tập candidate, không phải chạy 4 engine; “Gợi ý” chỉ là một cách chọn có trọng số, không thay đổi dữ liệu.
4. **Không bịa nhánh** quan trọng hơn việc luôn có 4 thẻ; người dùng cần biết “không có giá” hơn là bị gán nhãn “Rẻ” sai.
5. **Nhãn `ASSUMED`/`PARTIAL`** giữ cam kết trung thực; không quảng cáo “đúng luật tuyệt đối” khi OSM/GTFS chưa đủ.
6. **Không lưu route/vị trí** giữ privacy by default; cache chỉ thêm khi có nhu cầu và biện pháp.

## Phương án đã cân nhắc

| Phương án | Vì sao không chọn |
|---|---|
| **Tự viết RAPTOR + A* trong Go** | Kiểm soát tốt nhưng phải làm GTFS parser, turn restrictions, OSM access, multi-criteria, profiling; quá lớn cho R1, dễ sai luật |
| **Google Directions / Mapbox** | Không kiểm soát luật/dữ liệu, chi phí, điều khoản hiển thị, không dùng được với OSM/GTFS giả định; ranking hộp đen |
| **Valhalla** | Đường bộ tốt, transit yếu/không có RAPTOR đầy đủ; thêm engine, graph riêng |
| **GraphHopper** | Đường bộ tốt, transit hạn chế/trả phí; không giải quyết bài toán bus/metro |
| **OTP1** | Cũ hơn, không có lý do chọn; OTP2 đã chạy với graph cũ |
| **Nhiều engine (OTP + Valhalla)** | Tăng vận hành, hai bộ dữ liệu/không nhất quán; chưa cần |
| **Chạy 4 request OTP cho 4 nhánh** | Tốn latency, khó nhất quán; Pareto từ một tập đủ tốt hơn |
| **Luôn trả 4 nhánh, gán nhãn gần đúng** | Vi phạm trung thực dữ liệu; người dùng bị dẫn sai |
| **Không lọc luật, chỉ phạt** | Trái yêu cầu “ràng buộc cứng”; vẫn trả route sai |
| **Lưu route để cache/analytics** | Rủi ro PII; R1 không cần; nếu làm phải có ADR riêng |

## Hệ quả

### Tích cực

- Tận dụng OTP/RAPTOR/A* đã có; R1 tập trung vào lọc luật, Pareto, giải thích và UX.
- 4 nhánh nhất quán từ một tập candidate; dễ test/tái lập.
- Nhãn luật/dữ liệu rõ; không hứa quá.
- Hợp đồng route ổn định, client không phụ thuộc OTP.

### Tiêu cực / chi phí

- Phụ thuộc OTP/graph build; OTP lỗi/thiếu RAM ảnh hưởng toàn bộ route.
- Lọc luật ở graph build đòi hỏi script/kiểm OSM kỹ; tag Việt Nam không đầy đủ.
- Giá vé/giờ bus là giả định; “Rẻ” có thể không có hoặc sai nếu bảng giá cấu hình sai.
- Một lần gọi OTP + Pareto có thể không đủ đa dạng 4 nhánh; cần đo và có thể gọi lần hai.
- Không có realtime; route có thể lệch giờ tàu/xe.
- Hậu kiểm luật cần hiểu OTP response; nếu OTP đổi schema phải sửa adapter.

### Rủi ro

- OSM thiếu `foot`/`bicycle`/`access` hoặc turn restriction lỗi → route “FILTERED” vẫn có thể sai thực tế; UI phải trung thực.
- GTFS giả định (bus 15 km/h, metro) có thể sai giờ; người dùng trễ nếu tin tuyệt đối.
- Bảng giá `ASSUMED` có thể lệch giá thật; phải ghi “ước tính” và cập nhật.
- OTP 2.x có thể đổi API/schema; pin version và contract test.
- Graph TP.HCM cần RAM/CPU; pilot cùng host có thể hết bộ nhớ khi route nhiều.
- “Gần” theo tổng quãng đường có thể không phải điều người dùng muốn; cần usability test, có thể đổi nhãn/tiêu chí.

## Việc tiếp theo

- [ ] Chủ dự án chốt: dùng OTP (không tự viết engine), định nghĩa “Gần”, trọng số 0,5 thời gian / 0,3 tiền / 0,2 chuyển tuyến, cap 1,5×/2×, bán kính 800 m, p95 < 2 s, bảng giá `ASSUMED`.
- [ ] Viết `../02-architecture/` tài liệu tìm đường chi tiết (pipeline, config, graph build) — task 2.
- [ ] Chốt OTP version/graph build ở `../04-development/01-toolchain.md`; pin digest.
- [ ] Dựng graph nhỏ fixture + 10 route case + case luật; chạy OTP thật, ghi bằng chứng.
- [ ] Đo p95 trên máy pilot; chỉnh timeout/số candidate.
- [ ] Viết test hậu kiểm luật và test “không bịa nhánh”.
- [ ] Rà OSM/GTFS license/attribution với luật sư trước pilot.
- [ ] Khi có dữ liệu chính thức/GTFS-RT: ADR mới, không sửa lịch sử ADR này.

## Tham khảo

- `../00-foundation/02-scope.md` — coverage TP.HCM, `ASSUMED`/`UNSUPPORTED`, gate route test.
- `../04-development/02-conventions.md` — thứ tự lọc → đa tiêu chí → chọn → giải thích; trọng số có tên.
- `../01-product/07-design-canvas.md` — canvas `RouteBranches`/`RouteLogic`; 4 nhánh, lọc đúng luật trước.
- Bản export canvas do chủ dự án giữ ngoài repo (đọc 18-11-2026): `source/RouteLogic.dc.html`, `source/RouteBranches.dc.html`, `INDEX.md`.
- `../07-changes/change-18-11-2026.md` mục 9 — trọng số 0,5 thời gian / 0,3 tiền / 0,2 chuyển tuyến và ngưỡng 1,5×/2× là giả định.
- Repo cũ `../GoWander`: `ADR-020` (Metro giả định), `ADR-023` (OTP 2.10.0), `ADR-027` (API Go/mã lỗi), `ADR-032` (bus OSM, graph TP.HCM, tốc độ 15 km/h), `data/test-cases/route-cases.csv` (RC-01…RC-10).
- OTP 2.x: RAPTOR cho transit, A* cho street (theo tài liệu OTP; **chưa kiểm chứng lại trên graph v2**).
