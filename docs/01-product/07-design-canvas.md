# 07 — Design canvas (R0.5)

- **Trạng thái:** Proposed — chờ chủ dự án duyệt
- **Ngày:** 18-11-2026
- **Chủ sở hữu:** chủ dự án
- **Liên quan:** `05-brand-ui-direction.md`, `07-screen-map.html`, `08-accounts-roles-partners.md`, `09-stories-explore-search.md`, `../00-foundation/02-scope.md` (R0.5)

## TL;DR

- Toàn bộ màn hình R0.5 nằm trên **một canvas claude.ai (loại Design)**: <https://claude.ai/artifact/Wt2LzogXbLm7msPKHBaUjY> (riêng tư, chỉ chủ dự án mở được).
- Ngày 18-11-2026, bản canvas xuất ra có **152 màn trong 25 cụm**, mỗi cụm một hàng. Hàng đầu là **Sơ đồ màn hình** (`project/ScreenMap.dc.html`). Bản export đầy đủ (source/static/assets/png) do chủ dự án giữ **ngoài repo**; repo chỉ giữ bản sao sơ đồ `07-screen-map.html` và các quy ước trong file này.
- Bản sao của sơ đồ cho agent đọc trong repo: [`07-screen-map.html`](./07-screen-map.html). Khi canvas thay đổi, bản sao phải cập nhật trong cùng PR.
- Canvas chỉ là **thiết kế**: số liệu là DEMO, nội dung địa điểm và câu chuyện là ví dụ, chưa kiểm chứng. Quyết định sản phẩm nằm trong docs, không nằm trên canvas.
- Design System artifact riêng (<https://claude.ai/artifact/94b7v6RALKbTwBu887XGHt>) **vẫn dùng token blue cũ**, cần cập nhật sang v4.

## Cách agent đọc canvas

- Mỗi màn là một file `project/<Tên>.dc.html` (HTML + CSS inline). Đọc bằng công cụ Artifact (`action: read`, `path: "project/<Tên>.dc.html"`), không dùng web fetch.
- Nếu môi trường không có công cụ Artifact hoặc canvas riêng tư, dùng **bản export do chủ dự án giữ ngoài repo** (nếu có): `source/` cho bản gốc `.dc.html`, `static/` cho HTML mở bằng trình duyệt, `png/` để xem bằng mắt. Bản export là snapshot, có thể cũ; ảnh PNG có ô xám ở component lồng nhau (Gô, ảnh minh hoạ) — khi cần đầy đủ thì mở canvas.
- `project/canvas.json` là chỉ mục: vị trí, cỡ, tiêu đề từng khung, tiêu đề cụm.
- Liên kết giữa các màn là `<a href="X.dc.html">`; đi theo link để biết luồng.
- Trước khi sửa một màn, **đọc lại bản mới nhất**: chủ dự án có thể sửa tay trên canvas.

## Cụm màn hình

| # | Cụm | Ví dụ màn |
|---|---|---|
| 0 | Sơ đồ màn hình | ScreenMap |
| 1 | Nền tảng thiết kế | Direction, Mascot, Chip (Gô), PlaceArt |
| 2 | App · Bắt đầu & giao diện | Splash, Onboarding, Xin quyền vị trí, ThemePicker |
| 3 | App · Bản đồ & tìm kiếm | Main, NavMenu (thanh icon xổ ra), Search, SearchFood, SearchHospital, PlaceDetail, sheet |
| 4 | App · Tìm tuyến & cảnh báo | RouteResults, RouteDetail, QuickAlert (loa 3 mức), Unsupported, Offline |
| 5 | App · Khám phá, câu chuyện, Việt Nam | Discover, StoryList, StoryReader, ExploreVietnam, bản tối |
| 6 | App · Lưu, kế hoạch, hồ sơ | Saved, Plan, Notifications, NotificationDetail, Settings, MyBookings, BookingCancel |
| 7 | App · Tài khoản | Login, FaceID, Register, VerifyEmail, ForgotChannel, Forgot (mã), ResetPassword |
| 8 | App · Đóng góp & trạng thái chung | SuggestPlace, MySubmissions, Loading, ComingSoon |
| 9 | App · Chế độ quán | ShopDashboard (thống kê theo tháng + món đặt nhiều), ShopLogin, ShopStaffLogin, ShopToday, ShopOrderDetail, ShopStatus, ShopMenuQuick, ShopNotifications (`MyShops`, `ShopInvite` còn trên canvas, gắn nhãn “Không dùng nữa”; `ShopStaffLogin` = đổi người đứng ca trên máy dùng chung) |
| 10 | App · iPad ngang | TabletMap, TabletSearch, TabletDiscover, TabletStory, TabletBooking |
| 11 | App · R2–R7 | Wander, JourneyMemory, Community, AICompanion, Booking, BookingEmails |
| 12 | Admin · Đăng nhập | Login, 2FA, Invite, Error, Locked |
| 13 | Admin · Giám sát & tài chính | Overview, Finance, Logs |
| 14 | Admin · Nội dung | PlaceEdit, Map, Places (duyệt), Stories, Promos, Landing, Feedback |
| 15 | Admin · Partner | Partners, PartnerCreate, PartnerDetail, PartnerRequests |
| 16 | Admin · Hệ thống | OTP, Backup, Users, Notifications |
| 17 | Partner · Đăng ký hợp tác | Apply, ApplyDocs, ApplyReview, ApplyPending, ApplyFix, Emails |
| 18 | Partner · Đăng nhập quán + quên mật khẩu | PartnerLogin, AuthShopSetPassword, AuthShopForgot, AuthShopEmails, ForgotChannel, ForgotCode, ForgotReset |
| 19 | Partner web | Dashboard, Orders, Menu, Booking, Profile, Staff, Branches, Devices, Payout |
| 20 | Showcase web | Home, Policy |
| 21–24 | Responsive | Admin/Partner ở iPad 834 và điện thoại 390 (khung thường) |

Danh sách đầy đủ và luồng chính: [`07-screen-map.html`](./07-screen-map.html).

## Quy ước vẽ

- Màu, chữ, component: `05-brand-ui-direction.md` (palette v4).
- Màn app 390 × 844; iPad 1194 × 834; web 1440 fluid.
- Mỗi bước của một luồng là **một khung riêng** (chủ dự án yêu cầu; ví dụ đăng nhập không vẽ gộp).
- Số liệu ghi “DEMO”; nội dung lịch sử ghi “Nội dung mẫu, cần biên tập kiểm chứng”.
- Ảnh thật có ghi công; nguồn trong ghi chú `photoCredit` của canvas.
- Tính năng chưa có: vẽ trạng thái “Sắp có”/làm mờ, không giấu.
- **Màn nhỏ (điện thoại, iPad):** không lướt ngang. Bảng thành card (nhãn lấy từ tiêu đề cột, ≤ 1100 px), nhiều hơn 4 card thì có phân trang; hàng chip thành 2–3 chip + “Thêm ▾” hoặc ô chọn; sửa/nhân bản/xoá là nút icon (vẫn có `aria-label`); danh sách món là card ảnh lớn.
- **Món ăn luôn có ảnh;** ảnh là **bắt buộc** khi thêm/sửa món (hiện dùng ảnh Wikimedia Commons có ghi công).
- **Chuông thông báo cho mọi loại người dùng** (chỉ icon + số nhỏ), mỗi loại một hộp thư: khách `Notifications`, chế độ quán `ShopNotifications`, Partner `PartnerNotifications`, Admin `AdminInbox` (nơi *nhận*; `AdminNotifications` là nơi *gửi*).
- **Thanh icon bấm xổ ra** cho mọi loại: khách điện thoại (`NavMenu`), chế độ quán (mục đang mở hiện tên), iPad (`TabletNavMenu`), Partner trên điện thoại (thanh dưới + “⋯”).

## Việc còn thiếu trên canvas (18-11-2026)

- Bản tối cho các màn chính khác ngoài Câu chuyện.
- Responsive cho Showcase; responsive điện thoại cho Admin (hiện chỉ có iPad).
- Bản đồ 2D/3D, chỉ dẫn từng bước dạng chia đôi, cảnh báo đi lệch tuyến, thời tiết theo chuyến (đã có bong bóng thời tiết `Weather`; chưa có thời tiết dọc tuyến).
- Bộ icon thống nhất; Design System artifact sang v4.
- Làm lại bố cục SuggestPlace, MySubmissions, WanderSetup, WanderItinerary (mới chỉ đổi màu).
- Kiểm tra hiển thị thực tế: đã quét tự động mọi khung ở đúng bề rộng (tràn, cắt, lướt ngang); mới xem bằng mắt một số màn. **Chưa duyệt từng màn bằng mắt.**
