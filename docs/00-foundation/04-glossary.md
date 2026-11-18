# 04 — Glossary

- **Trạng thái:** Draft
- **Ngày:** 18-11-2026
- **Chủ sở hữu:** chủ dự án
- **Liên quan:** `01-vision.md`, `02-scope.md`, `03-principles.md`

Mục đích: **một nghĩa cho mỗi từ**. Nếu một từ có nhiều nghĩa, ghi rõ nghĩa dùng trong GoWander. Khi viết docs/code/UI, dùng đúng từ ở đây.

## 1. Sản phẩm & UX

| Từ | Nghĩa trong GoWander |
|---|---|
| **GoWander Client** | App cho du khách/người dùng cuối: Android, iOS; Client Web sau |
| **GoWander Business** | App cho Partner + Admin, dùng phân quyền để hiện giao diện phù hợp |
| **Client Web** | Web cho người dùng cuối: **Showcase marketing** trong R1; app/lập kế hoạch sau |
| **Showcase Web** | Phần marketing công khai của Client Web: dự án, dịch vụ, giá, chính sách, tải app, liên hệ |
| **Partner Web** | Web cho đối tác quản lý dịch vụ/khuyến mãi/booking; R6 |
| **Admin Web** | Web vận hành nội dung, địa điểm, ảnh, feedback, người dùng; có trong R1 |
| **Guest** | Người dùng không đăng nhập; R1 dùng guest-first |
| **Plan** | Kế hoạch do người dùng tạo: nhiều điểm, thứ tự, thời gian, ghi chú |
| **Stop** (trong plan) | Một điểm dừng trong plan, có thể là địa điểm hoặc điểm xuất phát |
| **Itinerary** | Lịch trình có thời gian cụ thể, thường do Wander Mode tạo |
| **Wander Mode** | Hero feature: từ thời gian/sở thích/điểm xuất phát/ngân sách → hành trình có lý do, có Plan B |
| **Plan B** | Phương án thay thế khi mưa, trễ, đóng cửa, quá đông |
| **Discover** | Duyệt/tìm địa điểm, món ăn, câu chuyện, hidden gems |
| **Guide** | Nội dung biên tập theo vùng/chủ đề (ví dụ cẩm nang TP.HCM) |
| **Hidden Gem** | Địa điểm ít người biết, do biên tập chọn; không phải “review nhiều sao” |
| **Recommendation** | Gợi ý có giải thích; tách khỏi quảng cáo trả tiền |
| **Saved place** | Địa điểm người dùng lưu lại; local-first, không public |
| **Place submission** | Đề xuất địa điểm mới của user; pending đến khi admin duyệt |
| **Moderation** | Duyệt/sửa/gộp/từ chối nội dung user; có audit |
| **Verification** | Trạng thái tin cậy của place: unverified/verified/official |
| **UGC** | User-generated content: địa điểm, ảnh, bài viết, bình luận |
| **Photo submission** | Ảnh user gửi kèm đề xuất địa điểm; **R1.6** (sau cổng pháp lý); pending, phải duyệt, strip EXIF |
| **Comment** | Bình luận của user; R4 text-only, cần account/moderation; ảnh comment sau |
| **Journey Memory** | Bản ghi hành trình đã đi: plan, ảnh, ghi chú, thời gian |
| **Storybook** | Cách trình bày Journey Memory thành câu chuyện/album riêng tư |
| **Explore** | Giai đoạn đi theo tuyến, xem bản đồ, chỉ dẫn, thời tiết |
| **Community** | Post, follow, comment, post-to-journey (R4) |
| **AI Companion** | Trợ lý AI hỗ trợ lập kế hoạch/giải thích, có giới hạn nguồn (R5) |
| **Business** | Lớp đối tác/doanh thu: Partner Web, Ads, Affiliate, Booking |
| **Services** | Dịch vụ GoWander cung cấp: free, business, ads, affiliate, booking… |
| **Pricing** | Giá/gói/dùng thử/huỷ; phải minh bạch, có VAT |
| **Policy** | Chính sách privacy, terms, cookies, content, map, refund… |
| **Ads** | Quảng cáo có nhãn; không thay đổi ranking tự nhiên |
| **Affiliate** | Giới thiệu dịch vụ, theo dõi referral/hoa hồng hợp lệ |
| **Booking** | Đặt dịch vụ; R7, bắt đầu bằng external link |

## 2. Di chuyển & dữ liệu transit

| Từ | Nghĩa |
|---|---|
| **Mode** | Phương tiện: walk, bus, metro, bicycle, car, motorbike, ferry |
| **Bike** | Trong tài liệu này = **xe đạp** (`BICYCLE`), không phải xe máy |
| **Motorbike** | Xe máy; mode riêng, chưa làm trong R1 |
| **Route** | Một đường đi từ A đến B, gồm nhiều leg |
| **Leg** | Một chặng của route: đi bộ, một chuyến bus, một chuyến metro… |
| **Step** | Một chỉ dẫn rẽ/đi thẳng trong leg đi bộ/xe đạp/ô tô |
| **Transfer** | Đổi tuyến/phương tiện; có thể là walking transfer |
| **Trip** | Chuyến đi của người dùng, có thể gồm route + plan + memory |
| **Origin / Destination** | Điểm đi / điểm đến |
| **Stop** (transit) | Trạm bus/điểm dừng |
| **Station** | Ga metro/rail |
| **Place / POI** | Địa điểm/point of interest |
| **Timetable** | Lịch chạy theo giờ; có thể là giả định |
| **Headway / Frequency** | Khoảng cách/ tần suất giữa hai chuyến |
| **Realtime** | Dữ liệu trễ/giờ thực; R1 chưa có GTFS-RT |
| **GTFS** | Định dạng dữ liệu lịch trình transit |
| **GTFS-RT** | GTFS Realtime: cập nhật trễ/giờ thực |
| **OTP** | OpenTripPlanner — engine tìm tuyến dùng GTFS + OSM |
| **OSM** | OpenStreetMap — dữ liệu bản đồ mở |
| **PBF** | Định dạng file OSM nén |
| **Overture** | Nguồn dữ liệu địa điểm/bản đồ mở thay thế/bổ sung |
| **Geocoding** | Đổi địa chỉ/tên thành toạ độ |
| **Reverse geocoding** | Đổi toạ độ thành địa chỉ/tên |
| **Photon** | Geocoder dùng OSM; fallback |
| **Vietmap** | Nhà cung cấp bản đồ/geocoding Việt Nam; dùng có key |
| **`dataStatus`** | Trạng thái dữ liệu trong API: `OFFICIAL`, `ASSUMED`, `BEST_EFFORT`, `UNSUPPORTED` |
| **Pilot area** | TP.HCM — nơi transit được làm sâu (bus + metro L1 + walk/bike) |
| **Coverage** | Vùng/mode được hỗ trợ routing; ngoài coverage chỉ map/search/place |
| **ASSUMED** | Dữ liệu giả định/ước tính cho pilot, không phải lịch chính thức; phải ghi rõ phần có nguồn và phần ước tính |
| **BEST_EFFORT** | Route/geometry từ OSM, không có timetable đầy đủ; chỉ tham khảo. **Chưa dùng ở R1** (để dành cho vùng mở sau) |
| **UNSUPPORTED** | Chưa hỗ trợ routing tại khu vực/mode |
| **OFFICIAL** | Dữ liệu từ nguồn chính thức; chỉ dùng khi có bằng chứng |
| **DEMO** | Dữ liệu mẫu cho demo; không dùng như fact |
| **Mock** | Dữ liệu giả trong code/UI; phải đăng ký và thay trước pilot |

## 3. Kiến trúc & kỹ thuật

| Từ | Nghĩa |
|---|---|
| **Modular monolith** | Một backend deployable, chia module domain rõ ràng; không microservice |
| **Domain module** | Identity, Mobility, Discover, Journey, AI, Community, Ads, Affiliate, Booking |
| **Seam** | Interface/role/feature flag để thêm domain sau mà không phá kiến trúc |
| **API** | Giao diện HTTP giữa client/Admin và backend |
| **OpenAPI** | Spec hợp đồng API; nguồn sự thật |
| **Contract test** | Test đảm bảo code khớp OpenAPI |
| **DTO** | Data Transfer Object — kiểu dữ liệu qua API |
| **Error envelope** | Format lỗi thống nhất: `code`, `message`, `requestId` |
| **KMP** | Kotlin Multiplatform — chia sẻ code Android/iOS |
| **Compose Multiplatform** | UI Kotlin dùng chung Android/iOS |
| **MapLibre** | Thư viện bản đồ mã nguồn mở |
| **OpenFreeMap** | Nhà cung cấp tile vector miễn phí dựa OSM |
| **Tile** | Mảnh bản đồ theo z/x/y |
| **Raster / Vector** | Ảnh tile / dữ liệu vector (đường, nhà, label) |
| **PostgreSQL** | Database quan hệ |
| **PostGIS** | Extension địa lý cho PostgreSQL |
| **Migration** | File thay đổi schema có version |
| **Seed** | Dữ liệu nạp ban đầu (category, place, guide…) |
| **RBAC** | Role-based access control: client/partner/admin/editor |
| **Feature flag** | Bật/tắt tính năng theo môi trường/role |
| **ADR** | Architecture Decision Record — quyết định kiến trúc |
| **Change record** | Bản ghi thay đổi theo ngày, có bằng chứng |
| **Evidence** | Ảnh/log/biên bản chứng minh đã chạy/đã đo |
| **Gate** | Điều kiện bắt buộc để qua milestone/release |
| **R1, R2…** | Release/phase, xem `02-scope.md` |
| **MVP / PoC** | Sản phẩm tối thiểu / proof of concept; không đồng nghĩa production |
| **SLO** | Service Level Objective — mục tiêu chất lượng dịch vụ |
| **RPO / RTO** | Mất dữ liệu tối đa / thời gian khôi phục tối đa |

## 4. Bảo mật, riêng tư & pháp lý

| Từ | Nghĩa |
|---|---|
| **PII** | Thông tin định danh cá nhân |
| **Dữ liệu cá nhân** | Thông tin gắn với một người hoặc có thể dùng để xác định người đó |
| **Dữ liệu cá nhân nhạy cảm** | Ví dụ vị trí, theo dõi hành vi; cần bảo vệ cao hơn |
| **DPIA** | Đánh giá tác động xử lý dữ liệu cá nhân |
| **Consent** | Sự đồng ý rõ ràng, có thể rút lại |
| **Data localization** | Yêu cầu lưu trữ dữ liệu trong nước (tùy luật/ngành) |
| **Map sovereignty** | Yêu cầu thể hiện đúng chủ quyền/biên giới trên bản đồ Việt Nam |
| **ODbL** | Giấy phép dữ liệu OSM; cần attribution + share-alike cho derivative database |
| **CC BY-SA** | Giấy phép ảnh/nội dung; cần ghi công + share-alike |
| **Attribution** | Ghi công tác giả/nguồn/giấy phép |
| **Share-alike** | Tác phẩm phái sinh phải dùng cùng giấy phép |
| **License register** | Danh sách nguồn/giấy phép/điều kiện dùng |
| **Third-party** | Bên thứ ba: OTP, Mapillary, Open-Meteo, Vietmap, R2… |

## 5. Thời tiết & tiện ích

| Từ | Nghĩa |
|---|---|
| **Open-Meteo** | Nguồn dự báo thời tiết miễn phí phi thương mại; cần mua/self-host trước thương mại |
| **Apparent temperature** | Nhiệt độ cảm nhận (nóng bức) |
| **Precipitation probability** | Xác suất mưa |
| **Indoor / Outdoor** | Điểm trong nhà / ngoài trời; dùng cho gợi ý mưa/nắng |
| **Weather-aware** | Có xét thời tiết khi gợi ý tuyến/plan |
| **Plan B** | Phương án thay thế khi thời tiết/đóng cửa/trễ |

## 6. Từ dễ nhầm

| Cặp | Phân biệt |
|---|---|
| **Bike vs Motorbike** | Bike = xe đạp; motorbike = xe máy. Không dùng “bike” cho xe máy. |
| **Route vs Trip** | Route = đường A→B; Trip = chuyến đi của người dùng (có thể gồm route, plan, memory). |
| **Plan vs Itinerary** | Plan = người dùng tạo; Itinerary = lịch trình có giờ, thường do Wander Mode sinh. |
| **Stop vs Station** | Stop = trạm bus; Station = ga metro/rail. |
| **Place vs POI** | Place = địa điểm trong Discover; POI = điểm trên bản đồ. |
| **ASSUMED vs OFFICIAL** | ASSUMED = giả định/ước tính; OFFICIAL = có nguồn chính thức. Không được gọi ASSUMED là official. |
| **Recommendation vs Ads** | Recommendation = gợi ý tự nhiên có giải thích; Ads = trả tiền, phải gắn nhãn. |
| **Memory vs Storybook** | Memory = dữ liệu hành trình; Storybook = cách trình bày thành câu chuyện. |
| **Guest vs Account** | Guest = không đăng nhập, dữ liệu trên máy; Account = đề xuất địa điểm (R1.5), đồng bộ (R3), community (R4). |
| **Client vs Business** | Client = người dùng cuối; Business = Partner/Admin. |
| **MVP vs Pilot** | MVP = phạm vi tối thiểu; Pilot = chạy thử với người dùng thật, có gate. |

## 6b. Tài khoản, đối tác, nội dung (thêm 18-11-2026)

| Thuật ngữ | Nghĩa |
|---|---|
| Tài khoản quán | Tài khoản do Admin cấp sau khi duyệt hồ sơ; mỗi người trong quán đăng nhập bằng **mã nhân viên** + mật khẩu riêng (không có mật khẩu quán); mở giao diện kinh doanh trong app (ADR-0002). Email chỉ để khôi phục, không đăng nhập |
| Chủ quán / Quản lý / Nhân viên | Ba vai trò trong một quán, gắn với **mã nhân viên** (`CQ`/`QL`/`NV`); Chủ quán thấy đủ tab và tự CRUD nhân viên; bảng quyền ở `01-product/08-accounts-roles-partners.md` |
| Chi nhánh | Địa điểm của một doanh nghiệp, có địa chỉ, giờ mở, khung đặt trước và đơn riêng |
| Giai đoạn 1/2/3 (kinh doanh) | Đặt trước miễn phí + affiliate → gói tháng → hoa hồng thanh toán trong app (`01-product/10-business-model.md`) |
| Mã quán | Mã 2–6 chữ cái in hoa, duy nhất, do Admin đặt; tiền tố của mã nhân viên |
| Mã nhân viên | `<mã quán>-<vai trò><số thứ tự>` (ví dụ `LL-NV007`), dùng làm tên đăng nhập; **định danh, không bí mật**; bí mật là mật khẩu riêng; không dùng một mình để đăng nhập |
| Mật khẩu tạm | Mật khẩu dùng một lần do Quản lý/Chủ quán cấp lại; hiện một lần, hết hạn 24 giờ, bắt đổi khi đăng nhập |
| Đổi người đứng ca | Thao tác trên máy dùng chung: đăng xuất người hiện tại rồi mở lại ô mã nhân viên + mật khẩu; máy nhớ vài mã gần đây nhưng vẫn phải nhập mật khẩu |
| Hồ sơ hợp tác | Yêu cầu làm đối tác kèm giấy tờ, Admin xác minh trước khi cấp tài khoản (`partner_application`) |
| Đặt món trước | Khách chọn món (tên + giá) khi đặt trước; trả tại quán; tuỳ quán bật |
| Cọc | Tiền giữ bàn theo chính sách quán; được chụp lại lúc đặt; hoàn theo mốc thời gian huỷ |
| Câu chuyện | Bài về khu lịch sử/công trình, soạn bằng editor theo khối, có audio/video và ghi công media |
| Gô | Linh vật GoWander (robot hộp kiểu xe buýt mini); “Gô đọc” là loa đọc chỉ dẫn 3 mức |
| DEMO | Nhãn cho số liệu/nội dung giả trong bản vẽ |

## 7. Viết tắt

| Viết tắt | Nghĩa |
|---|---|
| ADR | Architecture Decision Record |
| API | Application Programming Interface |
| DTO | Data Transfer Object |
| GTFS | General Transit Feed Specification |
| GTFS-RT | GTFS Realtime |
| KMP | Kotlin Multiplatform |
| OSM | OpenStreetMap |
| OTP | OpenTripPlanner |
| PBF | Protocolbuffer Binary Format |
| PII | Personally Identifiable Information |
| POI | Point of Interest |
| RBAC | Role-Based Access Control |
| RPO/RTO | Recovery Point/Time Objective |
| SLO | Service Level Objective |
| UI/UX | User Interface / User Experience |

## Quy ước cập nhật

- Thêm từ mới khi nó xuất hiện trong docs/API/UI.
- Nếu một từ đổi nghĩa, sửa ở đây trước, rồi cập nhật docs liên quan.
- Không dùng từ viết tắt trong UI nếu người dùng không hiểu; ưu tiên tiếng Việt rõ nghĩa.
