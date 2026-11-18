# 05 — Kế thừa từ repo cũ

- **Trạng thái:** Draft
- **Ngày:** 18-11-2026
- **Chủ sở hữu:** chủ dự án
- **Liên quan:** `02-scope.md`, `../03-decisions/README.md`, `../06-research/README.md`

## TL;DR

Repo cũ (`../GoWander`, nhánh `develop`) không chỉ có docs: nó có 22 ADR (ADR-020 đến ADR-041), dữ liệu giả định, test case, bằng chứng chạy trên Android/iOS, backend Go và app KMP. V2 là “docs-first rebuild”, nên mỗi quyết định cũ phải được ghi là **giữ ý / thay / hoãn / chưa xem xét**, theo quy tắc ở `../README.md` (“Quyết định của repo cũ không mặc nhiên là quyết định của v2”).

> **Giới hạn của tài liệu này:** chỉ dựa trên ADR, research và dữ liệu đã đọc ngày 18-11-2026. **Chưa audit code** trong `backend/` và `mobile/`; trạng thái triển khai là theo lời các ADR/change record cũ, chưa kiểm chứng lại.

## Bảng kế thừa theo ADR

| ADR cũ | Chủ đề | Ý chính (đã đọc) | Ở v2 |
|---|---|---|---|
| 020 | Transit giả định cho Metro L1 | Feed tự biên soạn có nhãn; phân biệt **có nguồn** (báo chí: giờ đầu/cuối, số lượt, giãn cách) và **giả định** (giờ cao điểm, thời gian từng chặng). Điều kiện: trước pilot người thật phải có dữ liệu được phép dùng hoặc xác nhận của nhà vận hành | **Giữ ý.** Điều kiện pilot đã thêm vào gate R1 (`02-scope.md`). Cần ADR-0004 |
| 021 | Thời tiết | Có hiển thị thời tiết nhẹ, gắn với di chuyển; nhà cung cấp chưa chọn | **Lệch:** v2 xếp weather vào R2; app cũ đã có trip weather. Cần quyết định có dời về R1 không |
| 022 | Đăng nhập, ngôn ngữ, quy mô Admin | Mật khẩu + Google song song; OTP qua email | Giữ hướng ở R1.5 |
| 023 | OTP 2.10.0 khoá theo digest | Khoá phiên bản | Giữ ý; cần ADR-0004 |
| 025 | Tile OpenFreeMap | Không SLA; phải quyết định lại trước pilot | **Giữ**; thêm gate “tile có SLA/tự host” |
| 026 | Toolchain mobile | Bộ phiên bản khớp maplibre-compose 0.20.0 | Giữ ý; chưa ADR v2 |
| 027 | API tìm tuyến bằng Go | Lớp Go giữa app và OTP | Giữ; chưa ADR v2 |
| 028 | PostgreSQL/PostGIS cho Discover | Image, quyền DB, migration, nguồn dữ liệu | Giữ; ADR-0003 (planned) |
| 031 | Geocoding | Photon tự chạy, Vietmap là hướng nâng cấp | Giữ; chưa ADR v2 |
| 032 | Bus từ OSM, graph TP.HCM, CAR + BICYCLE | Bus giả định tốc độ 15 km/h (không nguồn), bỏ 119/349 relation; **một graph OTP cho khung TP.HCM**; API có CAR/BICYCLE, xe máy bị từ chối | **Giữ graph TP.HCM + bus OSM.** **Thay:** CAR hoãn (cần dữ liệu giao thông). Không hứa tìm tuyến ngoài TP.HCM |
| 033 | Nav/ảnh đường phố/loại bản đồ | **Vệ tinh chưa bật** vì cần giấy phép (Esri/ArcGIS); chọn 2D/3D | **Giữ:** vệ tinh ngoài R1 |
| 035 | Lưu ảnh Cloudflare R2 | R2 bucket riêng, CDN; adapter R2 chưa làm | Giữ ý; cần rà data-localization (`../06-research/README.md`) |
| 036 | Lưu plan/saved trên máy | Một chuỗi JSON trong Prefs; nâng cấp SQLDelight khi cần | **Giữ** (đã link ở `06-saved-and-contributed-places.md`) |
| 037 | Admin web | Go server-rendered + htmx, loại React/Next vì “một người, chưa có Node”; §2–3 bị ADR-039 thay | **Thay:** v2 chọn React + Vite (ADR-0001, chủ dự án chốt 18-11-2026). Đánh đổi đã ghi ở ADR-0001 |
| 038 | Đo lường sử dụng | Chỉ bộ đếm gộp theo ngày, không lưu sự kiện từng người | **Giữ nguyên tắc** (đã ghi ở seam analytics) |
| 039 | Tài khoản Admin | Tài khoản riêng + mật khẩu; email qua SMTP | Giữ hướng; chi tiết ở ADR-0002 (planned) |
| 040 | Tài khoản app | Guest được dùng; email/SĐT/tên đăng nhập + Google; phiên 60 ngày; **có xoá tài khoản trong app**; Apple Sign-In để sau nhưng ghi nhận Guideline 4.8 | **Giữ.** Đã thêm vào R1.5: xoá tài khoản + cân nhắc Sign in with Apple |
| 041 | Nội dung Editor.js | Block JSON, app render native | **Chưa chốt:** ADR-0001 còn “TipTap hoặc Editor.js” |
| 024, 029, 030, 034 | Giám sát Admin, API client mobile, UI kit, IA UX-2 | Chưa đọc kỹ | **Chưa xem xét** |

## Dữ liệu và bằng chứng đáng mang sang

| Thứ | Vị trí ở repo cũ | Ghi chú |
|---|---|---|
| Feed Metro L1 giả định | `data/assumed/hcmc-metro-l1-gtfs/` | Sinh bằng script; không sửa tay (`ADR-020`) |
| Feed bus từ OSM | `data/assumed/hcmc-bus-osm-gtfs/` | Chỉ manifest được commit; dựng lại bằng script (`ADR-032`) |
| Route test case | `data/test-cases/route-cases.csv` | 10 case (metro, đi bộ, ngoài giờ, ngoài vùng phủ RC-09, input sai RC-10). Là nền cho gate “route test ≥ 90%” |
| Danh mục nguồn dữ liệu | `data/sources/`, `docs/research/data-sources-inventory.md` | Xem `../06-research/README.md` |
| Bằng chứng chạy app | `docs/evidence/m03`, `m06`, `m06b`, `nav1`, `rd0`, `ux`, `ux2` | Ảnh Android/iOS có ngày |

## Việc cần chốt

- [x] **Code cũ: viết lại** (chủ dự án chốt 18-11-2026). Code repo cũ chỉ làm tham chiếu đọc; không copy máy móc. Dữ liệu, test case, script dựng feed và bằng chứng ở bảng trên vẫn được mang sang có ghi nguồn.
- [ ] Vì viết lại: mỗi công nghệ giữ nguyên (Go, KMP, OTP, PostGIS, MapLibre, R2) cần một ADR v2 ngắn, ghi “kế thừa” kèm lý do, trước khi dựng repo code.
- [ ] Weather: R1 hay R2 (xem ADR-021 ở trên).
- [ ] Soạn ADR v2 cho: Go/Compose-KMP/OTP/PostGIS/MapLibre/R2 (hoặc ghi “kế thừa nguyên văn” kèm lý do).
- [ ] Đọc nốt ADR-024, 029, 030, 034.
