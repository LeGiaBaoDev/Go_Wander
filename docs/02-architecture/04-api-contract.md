# 04 — API contract (v0, R1)

- **Trạng thái:** Proposed — chờ chủ dự án chốt
- **Ngày:** 18-11-2026
- **Chủ sở hữu:** chủ dự án
- **Liên quan:** `01-system-context.md`, `02-containers.md`, `03-data-model.md`, `../03-decisions/ADR-0003-api-contract-auth-rbac.md`, `../03-decisions/ADR-0005-routing-algorithm.md`, `../04-development/02-conventions.md`
- **Spec máy đọc:** [`api/openapi.yaml`](../../api/openapi.yaml) — nguồn sự thật của hợp đồng; tài liệu này giải thích và ghi quyết định.

## TL;DR

- REST/JSON trên HTTPS, mọi endpoint sản phẩm dưới **`/api/v1`**; OpenAPI 3.1 là nguồn sự thật.
- **v0 chỉ phủ API công khai cho Client Mobile R1**: cấu hình app, category, place, story, geocode, route, feedback. **Chưa có** endpoint Admin, auth client (R1.5), Partner (R6) — các slice đó mở sau, không tạo endpoint rỗng trước.
- R1 **guest-first**: không endpoint nào yêu cầu đăng nhập. Feedback gửi được không cần tài khoản.
- `POST /routes/plan` nhận **mảng `modes`** (một hoặc nhiều mode) và trả nhiều lựa chọn tuyến; hình học mặc định là encoded polyline **precision 5**, có thể xin GeoJSON.
- Mọi response lỗi cùng một envelope `{error:{code,message,requestId,details?}}`; ngoài coverage trả `UNSUPPORTED` (422) với `details.reason`, **không** trả tuyến rỗng.
- Thời gian ISO 8601 UTC (request route nhận offset), toạ độ WGS84 6 chữ số thập phân, ID UUID dạng chuỗi, tiền `bigint` VND.
- Phân trang public dùng **cursor**; tham số lọc/sắp xếp có whitelist.
- Các con số cache/rate limit dưới đây là **đề xuất** (`ASSUMED`), chỉnh sau khi đo.

## 1. Phạm vi v0

| Nhóm | Endpoint | Trạng thái |
|---|---|---|
| Health | `GET /health/live`, `GET /health/ready` | ✅ v0, ngoài `/api/v1`, không dữ liệu nghiệp vụ |
| Cấu hình | `GET /api/v1/app/config` | ✅ v0 |
| Category | `GET /api/v1/categories` | ✅ v0 |
| Place | `GET /api/v1/places`, `GET /api/v1/places/nearby`, `GET /api/v1/places/{id}` | ✅ v0 |
| Story/Guide | `GET /api/v1/stories`, `GET /api/v1/stories/{slug}` | ✅ v0 |
| Geocode | `GET /api/v1/geocode`, `GET /api/v1/geocode/reverse` | ✅ v0 |
| Route | `POST /api/v1/routes/plan` | ✅ v0 |
| Feedback | `POST /api/v1/feedback` | ✅ v0 |
| Admin | `/api/v1/admin/auth/*`, `/places`, `/stories`, `/media`, `/feedback`, `/audit` | ⏳ slice riêng khi làm Admin Web |
| Auth client | `/api/v1/auth/*`, `/api/v1/me/*` | ❌ R1.5 |
| Partner | `/api/v1/partner/*` | ❌ R6 |

Quy tắc: endpoint chỉ được thêm khi phase tương ứng mở; không tạo endpoint rỗng “chuẩn bị trước” (ADR-0003 mục 12).

## 2. Quy ước hợp đồng

| Chủ đề | Quy ước |
|---|---|
| Base path | `/api/v1`; health ở `/health/*` |
| Định dạng | `application/json; charset=utf-8`; upload sau này `multipart/form-data` |
| Tên field | `camelCase` trong JSON; `snake_case` trong DB, map ở repository |
| ID | UUIDv7 dạng chuỗi (`0192f3a4-5b6c-7d8e-9f01-234567890abc`); không lộ số thứ tự |
| Thời gian | ISO 8601; DB/API dùng UTC (`Z`); request route nhận offset (`+07:00`) và đổi sang UTC |
| Toạ độ | WGS84, `lat`/`lon` số thực, trả 6 chữ số thập phân; validate `-90..90`, `-180..180` |
| Tiền | Số nguyên VND (`bigint`), field mang hậu tố `Vnd` (`fareVnd`); không float, không dùng `currency` rời vì toàn hệ thống là VND |
| Ngôn ngữ | `Accept-Language` hoặc `?locale=vi|en`; fallback `vi`; response trả `locale` thực tế khi có bản dịch |
| Phân trang | Cursor: request `limit`, `cursor`; response `items`, `nextCursor` (null khi hết) |
| Lọc/sắp xếp | Tham số có tên, whitelist; giá trị ngoài whitelist → `VALIDATION_FAILED` |
| Idempotency | `Idempotency-Key` bắt buộc cho `POST /feedback`; lưu 24 giờ; cùng key khác payload → `409` |
| Request id | Mọi response có `X-Request-Id`; lỗi có `requestId` trong body |
| Phiên bản | Chỉ thêm trong `/api/v1`; breaking change mở `/api/v2`, giữ `/api/v1` thời gian chuyển |
| Log | Không log body mặc định, không log toạ độ/token/PII |

## 3. Xác thực và giới hạn tần suất

### 3.1 Xác thực

- v0 **public hoàn toàn**: không `Authorization`, không cookie. Không có dữ liệu riêng của người dùng.
- Khi mở R1.5, endpoint client dùng `Authorization: Bearer <token>` (opaque 256-bit, lưu hash). Admin dùng cookie `__Host-gw_admin`; Partner dùng `__Host-gw_partner` (ADR-0003 mục 2–3).
- Không endpoint nào nhận `accountId`, `role` hay `shopId` do client gửi để quyết định quyền.

### 3.2 Giới hạn tần suất (đề xuất, `ASSUMED`)

| Nhóm | Giới hạn | Ghi chú |
|---|---|---|
| Route | 30/phút/IP; 120/giờ/token khi có account | Tránh lạm dụng OTP |
| Geocode | 60/phút/IP | Không log query |
| Feedback | 5/ngày/IP | Có thể thêm captcha/honeypot nếu bị spam |
| Place/Story/Category/Config | 120/phút/IP | Đề xuất; chưa đo tải |
| Health | Không giới hạn | Chỉ nội bộ/giám sát |

Vượt giới hạn → `429 RATE_LIMITED`, header `Retry-After` (giây).

## 4. Mô hình lỗi

```json
{
  "error": {
    "code": "VALIDATION_FAILED",
    "message": "Toạ độ không hợp lệ.",
    "requestId": "req_01J8Z3K4M5N6P7Q8R9S0T1",
    "details": { "field": "origin.lat" }
  }
}
```

| Mã | HTTP | Khi nào |
|---|---|---|
| `VALIDATION_FAILED` | 400 | JSON sai, thiếu field, toạ độ/thời gian/mode sai, tham số ngoài whitelist |
| `UNAUTHENTICATED` | 401 | Chưa dùng ở v0; dành cho R1.5 |
| `FORBIDDEN` | 403 | Chưa dùng ở v0 |
| `NOT_FOUND` | 404 | Không có place/story, hoặc không được phép biết là có |
| `CONFLICT` | 409 | `Idempotency-Key` khác payload |
| `METHOD_NOT_ALLOWED` | 405 | Phương thức không được hỗ trợ cho endpoint |
| `PAYLOAD_TOO_LARGE` | 413 | Body quá giới hạn |
| `UNSUPPORTED_MEDIA_TYPE` | 415 | Loại nội dung không hỗ trợ |
| `UNSUPPORTED` | 422 | Ngoài coverage hoặc mode/constraint không hỗ trợ; `details.reason` = `OUTSIDE_COVERAGE`, `MODE_NOT_SUPPORTED`, `CONSTRAINT_NOT_SUPPORTED` |
| `NO_ROUTE` | 422 | Trong coverage nhưng không có tuyến thoả ràng buộc |
| `DATA_UNAVAILABLE` | 422 | Ngoài giờ chạy/không có lịch cho ngày đó |
| `ORIGIN_NOT_SNAPPABLE` | 422 | Điểm đi quá xa đường; `details.maxSnapM` |
| `DESTINATION_NOT_SNAPPABLE` | 422 | Điểm đến quá xa đường; `details.maxSnapM` |
| `RATE_LIMITED` | 429 | Vượt giới hạn |
| `INTERNAL` | 500 | Lỗi không lường trước; không trả chi tiết nội bộ |
| `ROUTING_LEGALITY_VIOLATION` | 500 | OTP trả route vi phạm sau hậu kiểm; log nội bộ, không trả chi tiết |
| `DEPENDENCY_UNAVAILABLE` | 503 | DB/geocoder/media không sẵn sàng |
| `ROUTING_UNAVAILABLE` | 503 | OTP lỗi/không kết nối |
| `ROUTING_TIMEOUT` | 504 | Quá deadline (API 8 giây) |

- `message` theo locale, không lộ stack/SQL/tên bảng; `details` chỉ chứa dữ liệu an toàn.
- **Không** dùng `200` + mảng rỗng để che lỗi coverage/route.

## 5. Nhãn dữ liệu

| Field | Giá trị | Nơi dùng |
|---|---|---|
| `dataStatus` (place/story/coverage) | `DEMO`, `OFFICIAL`, `ASSUMED`, `BEST_EFFORT` | Place/story công khai, coverage area |
| `dataStatus` (route) | `OFFICIAL`, `ASSUMED`, `BEST_EFFORT`, `UNSUPPORTED` | Route response |
| `legality.status` | `FILTERED`, `PARTIAL`, `UNKNOWN` | Mỗi route |
| `fareStatus` | `OFFICIAL`, `ASSUMED`, `UNKNOWN` | `metrics.fareVnd` |

- Client **phải** hiển thị nhãn `ASSUMED`/`DEMO` và không gọi là dữ liệu chính thức.
- Badge “Đúng luật” chỉ khi `legality.status='FILTERED'`, kèm helper “Theo dữ liệu bản đồ hiện có”.

## 6. Tài nguyên và schema

### 6.1 `PlaceSummary`

| Field | Kiểu | Ghi chú |
|---|---|---|
| `id` | uuid | |
| `slug` | string | |
| `name` | string | Theo `locale`, fallback `vi` |
| `category` | `CategoryRef` | `id`, `code`, `name` |
| `location` | `{lat, lon}` | 6 chữ số thập phân |
| `address` | string? | |
| `distanceM` | integer? | Chỉ khi request có điểm tham chiếu |
| `verification` | `unverified|verified|official` | |
| `dataStatus` | `DEMO|OFFICIAL|ASSUMED|BEST_EFFORT` | |
| `source` | `SourceRef` | `code`, `name`, `licenseCode`, `attribution` |
| `cover` | `MediaRef?` | Ảnh bìa nếu có |

### 6.2 `PlaceDetail` = `PlaceSummary` +

`summary`, `description`, `openingHours`, `phone`, `website`, `priceLevel` (1–4, ước tính), `indoor`, `wheelchair`, `categories[]`, `media[]`, `entrance?`, `createdAt`, `updatedAt`.

### 6.3 `StorySummary` / `StoryDetail`

- Summary: `id`, `slug`, `kind` (`historical|guide|food`), `title`, `summary`, `cover?`, `readTimeMin?`, `audioDurationS?`, `placeCount`, `dataStatus`, `publishedAt`.
- Detail thêm: `blocks[]` (allowlist paragraph/header/list/quote/image/image-compare/timeline/place-card/audio/video/divider), `media[]` (credit/license), `places[]` (PlaceSummary rút gọn), `attribution`.
- Ảnh/video trong block tham chiếu `mediaId`; client tra trong `media[]`.

### 6.4 `GeocodeResult`

`name`, `address?`, `location {lat,lon}`, `category?`, `placeId?`, `source {code,name,licenseCode,attribution}`. Kết quả ngoài DB GoWander **không** có `placeId`.

### 6.5 `RoutePlanRequest`

| Field | Kiểu | Bắt buộc | Ghi chú |
|---|---|---|---|
| `origin` | `Waypoint` | ✅ | `{lat, lon, name?, placeId?}` |
| `destination` | `Waypoint` | ✅ | |
| `modes` | mảng `WALK|BUS|METRO|BICYCLE` (1–4, không trùng) | ✅ | `bike` trong yêu cầu sản phẩm = `BICYCLE`; `BUS`/`METRO` được có leg đi bộ nối |
| `departAt` | string (ISO có offset) | ❌ | Mặc định hiện tại; quá khứ > 5 phút → `VALIDATION_FAILED` |
| `geometry` | `polyline|geojson` | ❌ | Mặc định `polyline` (precision 5); `geojson` trả LineString |
| `constraints` | `{avoidStairs?, wheelchair?}` | ❌ | Không hỗ trợ được → `UNSUPPORTED`/`CONSTRAINT_NOT_SUPPORTED` |

- Không nhận `car`, `motorbike`, `ferry`, `TRANSIT` ở R1.
- Mọi phần tử `modes` phải thuộc enum viết hoa `WALK|BUS|METRO|BICYCLE` (khớp `legs[].mode`); giá trị ngoài enum → `VALIDATION_FAILED`. Mode hợp lệ nhưng không có trong coverage (ví dụ `METRO` khi vùng chưa có metro) → `UNSUPPORTED` với `details.reason=MODE_NOT_SUPPORTED`, trả trước khi gọi OTP.
- Response vẫn là **nhiều lựa chọn tuyến** (các nhánh Gợi ý/Nhanh/Gần/Rẻ), không phải một route cho mỗi mode.
- Không nhận trọng số/tham số xếp hạng từ client.

### 6.6 `RoutePlanResponse`

Theo ADR-0005 mục 6:

- `requestId`, `graphVersion`, `dataStatus`, `coverage {status, areaId?, modes?}`, `routes[]`, `branches[]`, `assumptions[]`.
- `routes[]` là route duy nhất; `branches[]` tham chiếu `routeId`, có thể cùng route nhiều nhãn.
- Mỗi route: `id`, `metrics {durationS, distanceM, walkM, fareVnd?, fareStatus, transfers, waitingS}`, `legs[]`, `legality {status, unverifiedRules[]}`.
- Mỗi leg: `mode`, `from`, `to`, `startAt`, `endAt`, `distanceM`, `geometry` (encoded polyline precision 5, hoặc GeoJSON LineString theo field `geometry` của request), `routeName?`, `headsign?`, `realtime: false`, `steps[]?` cho leg đi bộ/xe đạp.
- Mỗi step: `instruction` (`CONTINUE|TURN_LEFT|TURN_RIGHT|UTURN|ARRIVE|DEPART`), `streetName?`, `distanceM`, `geometry?` cùng định dạng với leg.
- GeoJSON trả `{"type":"LineString","coordinates":[[lon,lat],...]}` — thứ tự **lon trước lat**; polyline dùng precision 5 (1e5), không phải 1e6.
- `branches[]`: `{kind: recommended|fastest|shortest|cheapest, routeId?, available, score?, reasons[], reason?}`; `available=false` có `reason` (`FARE_UNKNOWN`, `TIME_CAP_1_5`, `TIME_CAP_2_0`, `NO_FASTEST_CANDIDATE`…).

## 7. Endpoint

### 7.1 `GET /health/live`, `GET /health/ready`

- 200 `{"status":"ok"}`; ready 503 `{"status":"unavailable","checks":{...}}` khi DB/OTP chưa sẵn sàng.
- Không trả version chi tiết, connection string hay thông tin nội bộ.

### 7.2 `GET /api/v1/app/config`

Trả cấu hình công khai, không bí mật:

```json
{
  "minAppVersion": { "android": "0.1.0", "ios": "0.1.0" },
  "locales": ["vi", "en"],
  "defaultLocale": "vi",
  "map": {
    "styleUrl": "https://…/style.json",
    "attribution": "© OpenStreetMap contributors"
  },
  "coverage": {
    "areas": [{
      "id": "0192…",
      "code": "hcmc",
      "name": "TP.HCM",
      "modes": ["WALK", "BUS", "METRO", "BICYCLE"],
      "dataStatus": "ASSUMED"
    }]
  },
  "features": { "stories": true, "bicycle": true, "feedback": true },
  "maintenance": null
}
```

- `Cache-Control: public, max-age=60` (đề xuất).
- Không trả khoá Vietmap/OTP hay URL nội bộ.

### 7.3 `GET /api/v1/categories`

`items[]`: `id`, `code`, `name`, `icon?`, `sortOrder`, `parentId?`. Chỉ `status='active'`, có bản dịch theo locale (fallback VI).

### 7.4 `GET /api/v1/places`

| Tham số | Kiểu | Ghi chú |
|---|---|---|
| `q` | string 1–120 | Tìm theo tên; chuẩn hoá bỏ dấu |
| `category` | string (code) | Có thể lặp để lọc nhiều |
| `bbox` | `minLon,minLat,maxLon,maxLat` | Lọc khung bản đồ |
| `lat`, `lon` | double | Điểm tham chiếu để tính `distanceM`/sắp xếp; phải truyền cùng nhau |
| `sort` | `relevance|name` | Mặc định `relevance` khi có `q`, ngược lại theo thứ tự nổi bật |
| `limit` | int 1–50 | Mặc định 20 |
| `cursor` | string | Opaque |
| `locale` | `vi|en` | |

- `q` và `bbox` có thể kết hợp. Khi không có `q`/`bbox`/`category`, danh sách “nổi bật” sắp theo `featured_rank` tăng dần (null xuống cuối), rồi `distanceM` nếu có `lat`/`lon`, cuối cùng theo tên.
- `sort=relevance` chỉ hợp lệ khi có `q`; thiếu `q` → `VALIDATION_FAILED`.
- Chỉ trả `status='published'`, `deleted_at is null`.
- Response: `{items: PlaceSummary[], nextCursor}`.

### 7.5 `GET /api/v1/places/nearby`

| Tham số | Kiểu | Ghi chú |
|---|---|---|
| `lat`, `lon` | double | Bắt buộc |
| `radiusM` | int 100–10000 | Mặc định 2000 (`ASSUMED`) |
| `category` | string | Tuỳ chọn |
| `limit`, `cursor`, `locale` | | |

- Sắp xếp theo khoảng cách tăng dần; `distanceM` làm tròn mét.
- Ngoài TP.HCM vẫn trả place (map/search toàn quốc); **không** liên quan coverage route.

### 7.6 `GET /api/v1/places/{id}`

- 200 `PlaceDetail`; 404 `NOT_FOUND` nếu không publish/không tồn tại.
- `Cache-Control: public, max-age=60`.

### 7.7 `GET /api/v1/stories`, `GET /api/v1/stories/{slug}`

- List: `kind?`, `lat?`, `lon?`, `radiusM?`, `hasAudio?`, `limit`, `cursor`, `locale`. Lọc “Gần bạn” dựa `story_place` + `ST_DWithin`.
- Detail theo `slug`: 200 `StoryDetail`, 404 nếu chưa publish.
- Story chưa có bản dịch EN: trả VI + `locale: "vi"`; không tự dịch máy.

### 7.8 `GET /api/v1/geocode`

| Tham số | Kiểu | Ghi chú |
|---|---|---|
| `q` | string 1–200 | Địa chỉ/tên |
| `limit` | int 1–10 | Mặc định 5 |
| `lat`, `lon` | double | Thiên vị kết quả gần (không bắt buộc) |
| `locale` | `vi|en` | |

- Response `{items: GeocodeResult[]}`; có thể rỗng.
- Nguồn: Vietmap nếu có key, fallback Photon; response ghi `source` + `attribution`.
- Không log `q`; `Cache-Control: no-store`.

### 7.9 `GET /api/v1/geocode/reverse`

`lat`, `lon` bắt buộc; `locale`. Trả **200** với `{items: GeocodeResult[]}` (thường 1); khi không tìm thấy địa chỉ trả **200 với `items: []`**, **không** trả 404. Không bịa địa chỉ khi geocoder không có dữ liệu.

- Test: toạ độ giữa biển/không có địa chỉ → 200 `{items: []}`.

### 7.10 `POST /api/v1/routes/plan`

Request theo mục 6.5 (`modes` là mảng, `geometry` tuỳ chọn). Pipeline và mã lỗi theo ADR-0005:

1. Validate + chuẩn hoá.
2. Kiểm coverage: điểm ngoài polygon → `UNSUPPORTED`/`OUTSIDE_COVERAGE`; mode hợp lệ nhưng không có trong `coverage_area.modes` → `UNSUPPORTED`/`MODE_NOT_SUPPORTED`. Cả hai trả **trước khi** gọi OTP.
3. Snap origin/destination (tối đa 300 m, `ASSUMED`) → lỗi snap riêng.
4. Lọc luật cứng, gọi OTP, hậu kiểm, Pareto, chọn 4 nhánh.
5. Trả route + `dataStatus`/`legality`/`assumptions`.

- `200` với ít nhất một route; nhánh không có → `available:false` + `reason`, **không** bịa.
- Không cache response; `Cache-Control: no-store`.
- Timeout API 8 giây → `504 ROUTING_TIMEOUT`.

### 7.11 `POST /api/v1/feedback`

Header `Idempotency-Key` bắt buộc. Body:

```json
{
  "category": "route|place|app|other",
  "message": "Nội dung 10–2000 ký tự",
  "contactEmail": "optional@example.com",
  "contactPhone": "+84…",
  "appVersion": "0.1.0",
  "locale": "vi"
}
```

- 201 `{id, status:"new", createdAt}`; cùng key + cùng payload trả lại kết quả cũ; khác payload → `409`.
- Guest không cần tài khoản; nếu có account (R1.5) lấy email từ account, không tin client.
- `contactEmail`/`contactPhone` không bắt buộc; không public, không log.
- Rate limit 5/ngày/IP.

## 8. Ví dụ

### 8.1 Danh sách place

```json
{
  "items": [{
    "id": "0192f3a4-5b6c-7d8e-9f01-234567890abc",
    "slug": "cho-ben-thanh",
    "name": "Chợ Bến Thành",
    "category": { "id": "0192…", "code": "market", "name": "Chợ" },
    "location": { "lat": 10.772500, "lon": 106.698000 },
    "address": "Lê Lợi, Quận 1",
    "distanceM": 420,
    "verification": "unverified",
    "dataStatus": "DEMO",
    "source": {
      "code": "demo",
      "name": "Dữ liệu mẫu",
      "licenseCode": "internal",
      "attribution": "DEMO"
    }
  }],
  "nextCursor": null
}
```

### 8.2 Route ngoài coverage

```json
{
  "error": {
    "code": "UNSUPPORTED",
    "message": "Chưa hỗ trợ tìm tuyến tại khu vực này.",
    "requestId": "req_01J8Z3K4M5N6P7Q8R9S0T1",
    "details": { "reason": "OUTSIDE_COVERAGE" }
  }
}
```

### 8.3 Route request

```json
{
  "origin": { "lat": 10.770000, "lon": 106.690000, "name": "Điểm đi" },
  "destination": { "lat": 10.771000, "lon": 106.692000, "name": "Trạm Bến Thành" },
  "modes": ["WALK", "BUS"],
  "departAt": "2026-11-19T09:00:00+07:00",
  "geometry": "polyline"
}
```

`geometry` bỏ trống hoặc `polyline` → encoded polyline precision 5; `geojson` → GeoJSON LineString.

### 8.4 Route response (rút gọn)

```json
{
  "requestId": "req_01J8Z3K4M5N6P7Q8R9S0T1",
  "graphVersion": "hcmc-2026-11-18-a1b2c3",
  "dataStatus": "ASSUMED",
  "coverage": { "status": "SUPPORTED", "areaId": "0192…" },
  "routes": [{
    "id": "rt_01J8Z4…",
    "metrics": {
      "durationS": 1840,
      "distanceM": 8200,
      "walkM": 950,
      "fareVnd": 15000,
      "fareStatus": "ASSUMED",
      "transfers": 1,
      "waitingS": 240
    },
    "legs": [{
      "mode": "WALK",
      "from": { "name": "Điểm đi", "lat": 10.770000, "lon": 106.690000 },
      "to": { "name": "Trạm Bến Thành", "lat": 10.771000, "lon": 106.692000 },
      "startAt": "2026-11-18T02:00:00Z",
      "endAt": "2026-11-18T02:08:00Z",
      "distanceM": 650,
      "geometry": "oov`Ao{djSgEoK",
      "realtime": false,
      "steps": [
        { "instruction": "TURN_LEFT", "streetName": "Lê Lợi", "distanceM": 120 }
      ]
    }],
    "legality": { "status": "FILTERED", "unverifiedRules": [] }
  }],
  "branches": [
    {
      "kind": "recommended",
      "routeId": "rt_01J8Z4…",
      "available": true,
      "score": 0.42,
      "reasons": [{ "code": "BEST_OVERALL" }, { "code": "LEGAL_FILTER_APPLIED" }]
    },
    { "kind": "cheapest", "available": false, "reason": "FARE_UNKNOWN", "reasons": [] }
  ],
  "assumptions": ["FARE_ASSUMED", "SCHEDULE_ASSUMED"]
}
```

### 8.5 Giải mã polyline (precision 5)

Chuỗi trong ví dụ trên:

```text
oov`Ao{djSgEoK
```

- Cắt chuỗi thành từng cặp (lat, lon); mỗi giá trị là **delta** so với điểm trước.
- Mỗi giá trị giải mã 5 bit/lần, trừ offset 63, rồi nhân với **1e-5** (precision 5), không phải 1e-6.
- Kết quả: `(10.77000, 106.69000)` rồi `(10.77100, 106.69200)`.
- Nếu client chia 1e-6 sẽ ra `(1.077, 10.669)` — sai. Đây là lỗi thường gặp khi dùng decoder precision 6.

Với `geometry=geojson`, cùng route trả:

```json
{"type": "LineString", "coordinates": [[106.690000, 10.770000], [106.692000, 10.771000]]}
```

Lưu ý thứ tự **lon, lat** của GeoJSON.

## 9. Cache và hiệu năng (đề xuất)

| Endpoint | Cache | Ghi chú |
|---|---|---|
| `/app/config` | `public, max-age=60` | Cấu hình đổi ít |
| `/categories` | `public, max-age=300` | |
| `/places`, `/places/nearby` | `public, max-age=30` | Danh sách đổi khi Admin sửa |
| `/places/{id}` | `public, max-age=60` | Không `ETag` ở v0 |
| `/stories`, `/stories/{slug}` | `public, max-age=60` | |
| `/geocode*` | `no-store` | Không cache query/PII |
| `/routes/plan` | `no-store` | Không lưu vị trí |
| `/feedback` | `no-store` | |

- Mục tiêu route p95 < 2 giây, gate R1 ≤ 3 giây (ADR-0005); chưa đo.
- Không cache route cá nhân ở R1; nếu cần, chỉ cache theo ô lưới + khung giờ khi có ADR riêng.
- Nén `gzip`/`br`; giới hạn body mặc định 64 KB (route), 8 KB (feedback).

## 10. Tương thích và quy trình đổi hợp đồng

1. Sửa `api/openapi.yaml` trước, review diff.
2. Lint spec (Spectral hoặc tương đương), kiểm breaking change trong CI.
3. Sinh client TypeScript cho web; mobile dùng client Ktor mỏng + contract test với mock sinh từ spec.
4. Contract test ở Go: response thật validate theo schema, có test error envelope và các mã chính.
5. Chỉ thêm field/endpoint trong `/api/v1`; bỏ/đổi nghĩa/đổi kiểu là breaking → `/api/v2` + `Sunset`.
6. `operationId` ổn định; mỗi operation ghi realm/quyền (v0: public).

## 11. Đã chốt 19-11-2026 và câu hỏi mở

### Đã chốt

1. **`modes` là mảng** (một hoặc nhiều mode), không phải một `mode`. `bike` trong yêu cầu sản phẩm = `BICYCLE`. Phần tử nào không được coverage hỗ trợ → `UNSUPPORTED` với `details.reason=MODE_NOT_SUPPORTED`. Response vẫn trả nhiều lựa chọn tuyến (khớp màn RouteResults).
2. **Hình học tuyến**: mặc định encoded polyline **precision 5 (1e5)**; thêm field `geometry=polyline|geojson` (mặc định `polyline`). `geojson` trả GeoJSON `LineString`, toạ độ `[lon, lat]`.
3. **Reverse geocode không có địa chỉ**: trả **200** với `items: []`, không trả 404.
4. **ETag**: không làm ở v0; thêm khi thấy bản đồ cần cache (xem “Chưa làm”).
5. **`featured_rank`**: thêm cột `int` nullable vào `catalog.place`; danh sách “nổi bật” sắp `featured_rank` tăng dần (null xuống cuối), rồi khoảng cách nếu có `lat`/`lon`, cuối cùng theo tên.

### Còn mở

1. Có mở `GET /api/v1/coverage` công khai không, hay chỉ trả trong `/app/config` và route response? (v0: không mở)

## 12. Việc tiếp theo

### Chưa làm ở v0

- **`ETag`/`If-None-Match`/`304`: không làm.** Lý do: chưa thấy bản đồ cần cache; thêm khi đo thấy cần cache tile/danh sách place.
- Endpoint Admin, auth client, Partner: các slice sau (xem mục 1), không tạo trước.

### Việc cần làm

- [ ] Chủ dự án duyệt v0 và câu hỏi còn mở ở mục 11.
- [ ] Lint `api/openapi.yaml`, sinh mock + client thử ở `chore/repo-scaffold`.
- [ ] Viết slice Admin API khi làm Admin Web (đăng nhập 2FA, CRUD place/story/media, feedback, audit).
- [ ] Viết contract test Go cho place/geocode/route/feedback và các mã lỗi; có case reverse geocode rỗng (200, không 404) và giải mã polyline precision 5.
- [ ] Đo p95 route trên máy pilot; chỉnh timeout/rate limit.
- [ ] Bổ sung auth client + `/me/*` ở R1.5, Partner ở R6.
