# 08 — Tài khoản, phân quyền, đối tác, đặt bàn

- **Trạng thái:** Proposed — chờ chủ dự án duyệt; phần pháp lý **chưa được luật sư kiểm chứng**
- **Ngày:** 18-11-2026
- **Chủ sở hữu:** chủ dự án
- **Liên quan:** `../03-decisions/ADR-0002-account-roles.md`, `07-design-canvas.md`, `06-saved-and-contributed-places.md`, `../00-foundation/02-scope.md`, `../06-research/README.md`

## TL;DR

- **Một kiểu tài khoản** (email hoặc số điện thoại + mật khẩu) dùng cho app và web. Quyền gắn theo **vai trò ở từng quán**.
- App chỉ có một: tài khoản có vai trò ở quán thì thấy thêm **“Quán của tôi”** (chế độ quán). Việc cấu hình làm trên **Partner web**.
- **Không đăng ký Partner tự do:** quán gửi **hồ sơ hợp tác kèm giấy tờ**, Admin xác minh (1–3 ngày làm việc, báo qua email), rồi **Admin cấp tài khoản Chủ quán**. Chủ quán tự mời Quản lý / Nhân viên.
- Tài khoản **Admin** riêng, chỉ dùng web, do Owner cấp, **bắt buộc mã 2 lớp**.
- Đặt bàn (R7): khung giờ, sức chứa, **cọc theo chính sách quán**, **đặt món trước tuỳ chọn** (chỉ tên + giá).

## 1. Mô hình tài khoản

| Loại | Đăng nhập ở | Cách có tài khoản | Xác thực |
|---|---|---|---|
| Khách | App (điện thoại, iPad) | Tự đăng ký (R1.5); dùng app không cần tài khoản | Mật khẩu; Face ID đăng nhập nhanh |
| Thành viên quán | App + Partner web | Chủ quán: Admin cấp sau khi duyệt hồ sơ. Quản lý/Nhân viên: chủ quán mời | Mật khẩu; nhận lời mời qua email/SMS |
| Admin GoWander | Chỉ Admin web | Owner cấp; không đăng ký tự do | Mật khẩu + **mã 2 lớp bắt buộc** (ứng dụng xác thực hoặc passkey; mã dự phòng) |

- Một người có thể thuộc nhiều quán, mỗi quán một vai trò. Nhân viên đã có tài khoản app thì chỉ gắn thêm vai trò.
- **Quyền kiểm ở server.** App/web chỉ ẩn màn không đủ quyền; gọi API khi thiếu quyền phải bị chặn.
- Mọi thao tác của thành viên quán và admin ghi nhật ký (ai, lúc nào, làm gì).

### Quên mật khẩu (mọi loại tài khoản)

1. Nhập email hoặc số điện thoại → chọn kênh **Email (SMTP)** hoặc **SMS**.
2. Nhập mã 6 số; hết hạn sau 10 phút; tối đa 5 lần gửi mã/giờ/tài khoản; mã lưu dạng băm, dùng 1 lần.
3. Đặt mật khẩu mới (≥ 10 ký tự, có chữ và số, không trùng 3 mật khẩu gần nhất). Mọi phiên khác bị đăng xuất.

Admin: sai mật khẩu nhiều lần → khoá 15 phút; sai mã 2 lớp 5 lần → khoá, chỉ Owner mở. Phiên admin hết hạn sau 30 phút không thao tác. *(Các con số này là đề xuất, chưa chốt.)*

## 2. Vai trò trong quán

| Việc | Chủ quán | Quản lý | Nhân viên |
|---|:-:|:-:|:-:|
| Xem, xác nhận đơn hôm nay; đánh dấu đã đến | ✓ | ✓ | ✓ |
| Bật/tắt món Còn bán / Tạm hết | ✓ | ✓ | ✓ |
| Thực đơn/dịch vụ: thêm, sửa, xoá, đổi giá | ✓ | ✓ | — |
| Từ chối/huỷ đơn, chặn giờ, bật/tắt nhận khách | ✓ | ✓ | — |
| Khung giờ, sức chứa, hồ sơ quán, khuyến mãi, trả lời đánh giá | ✓ | ✓ | — |
| Chính sách cọc và hoàn tiền | ✓ | — | — |
| Đối soát tiền, tài khoản nhận tiền | ✓ | — | — |
| Mời, đổi vai trò, gỡ thành viên | ✓ | — | — |

- Mỗi quán đúng **một Chủ quán**; đổi chủ phải qua GoWander xác minh lại.
- Đổi tên hoặc địa chỉ quán cần GoWander duyệt lại.

## 3. Vai trò Admin GoWander

| Vai trò | Làm được |
|---|---|
| Owner | Mọi thứ; cấp/khoá tài khoản admin; khôi phục dữ liệu |
| Vận hành hệ thống | OTP, sao lưu, nhật ký |
| Biên tập nội dung | Địa điểm, câu chuyện, khuyến mãi, landing; đăng bài cần người thứ hai duyệt |
| Duyệt đề xuất | Hàng chờ đề xuất địa điểm, phản hồi từ app |
| Kế toán | Tài chính, đối soát Partner; chỉ xem |
| Hỗ trợ Partner | Hồ sơ hợp tác, tạo quán, cấp tài khoản chủ quán, xem đơn; không xem số tài khoản ngân hàng |

Admin không tự sửa tiền hay chính sách cọc của quán; chỉ gửi yêu cầu để chủ quán đồng ý.

## 4. Đăng ký hợp tác và xác minh

Luồng (web, có bản iPad/điện thoại):

1. **Thông tin quán:** loại hình (ăn uống, cà phê, lưu trú, tham quan/tour, dịch vụ khác), hình thức (hộ kinh doanh / doanh nghiệp), tên, **mã số thuế**, địa chỉ (chọn trên bản đồ; nếu quán đã có trên GoWander thì gắn vào đó), người đại diện, SĐT (xác thực OTP), email nhận kết quả.
2. **Giấy tờ:** danh sách đổi theo loại hình (bảng dưới). PDF/JPG/PNG ≤ 10 MB/tệp.
3. **Xem lại & gửi:** 3 cam kết bắt buộc: thông tin thật và còn hiệu lực; đồng ý điều khoản đối tác (hợp đồng điện tử); đồng ý xử lý dữ liệu cá nhân trong hồ sơ.
4. **Chờ duyệt:** thường **1–3 ngày làm việc**; chưa đăng nhập Partner được.
5. Kết quả qua email: **duyệt** (link kích hoạt dùng 1 lần, hết hạn 48 giờ) · **cần bổ sung** (lý do, hạn 14 ngày; quá hạn tự đóng và xoá giấy tờ) · **từ chối**.

Admin xử lý ở **Yêu cầu hợp tác**: đối chiếu mã số thuế, CCCD, giấy tờ theo loại hình; mỗi quyết định gửi email tự động và ghi nhật ký. Hạn xử lý đề xuất: 3 ngày làm việc.

### Giấy tờ theo loại hình (chưa kiểm chứng pháp lý)

| Giấy tờ | Áp dụng | Căn cứ (cần luật sư xác nhận) |
|---|---|---|
| Giấy chứng nhận đăng ký hộ kinh doanh / doanh nghiệp (có MST) | Mọi quán | NĐ 01/2021/NĐ-CP |
| CCCD người đại diện (2 mặt); giấy uỷ quyền nếu không phải chủ | Mọi quán | — |
| Giấy chứng nhận cơ sở đủ điều kiện ATTP, hoặc bản cam kết bảo đảm ATTP (diện miễn cấp) | Ăn uống, cà phê | Luật ATTP 2010; NĐ 15/2018/NĐ-CP |
| Giấy chứng nhận đủ điều kiện về an ninh, trật tự | Lưu trú | NĐ 96/2016/NĐ-CP |
| Giấy phép kinh doanh dịch vụ lữ hành | Tour | Luật Du lịch 2017 |
| Xác nhận tài khoản ngân hàng **cùng tên** chủ hộ/doanh nghiệp | Khi bật thu cọc | — |
| Đồng ý xử lý dữ liệu cá nhân | Mọi hồ sơ | Luật 91/2025/QH15 và NĐ 356/2025/NĐ-CP (thay NĐ 13/2023 từ 01-01-2026, theo `06-research`) |

- Giấy tờ mã hoá khi lưu; chỉ người xác minh xem; **xoá bản sao sau 30 ngày** nếu hồ sơ bị từ chối hoặc đóng *(đề xuất, cần chốt cùng chính sách lưu trữ)*.
- Việc của GoWander (không phải màn hình): nếu GoWander **thu cọc hộ quán**, có thể phải đăng ký **sàn thương mại điện tử** với Bộ Công Thương (NĐ 52/2013, sửa bởi NĐ 85/2021). **Phải hỏi luật sư trước khi bật thu cọc.**

## 5. Admin cấp tài khoản quán

- **Tạo quán & cấp tài khoản:** gắn địa điểm có sẵn, loại hình, gói (Cơ bản/Pro), thông tin chủ quán, kênh gửi lời mời (Email/SMS/cả hai), bật/tắt tính năng (đặt bàn, đặt món trước, thu cọc).
- **Chi tiết đối tác:** thành viên, lời mời đang chờ, tính năng, nhật ký; hành động: đổi chủ quán, đổi gói, gửi lại lời mời, khoá, gỡ.
- Trạng thái quán: Chờ kích hoạt → Hoạt động → (Đã khoá).

## 6. App: chế độ quán

Chỉ hiện khi tài khoản có vai trò ở ít nhất một quán; vào từ **Hồ sơ → Quán của tôi**. Thanh điều hướng riêng (tối) để không nhầm với chế độ khách.

| App (việc hằng ngày) | Partner web (cấu hình) |
|---|---|
| Đơn hôm nay: xác nhận, từ chối (có lý do), đổi giờ, đã đến | Khung giờ, sức chứa, chính sách cọc |
| Bật/tắt nhận khách, tạm ngưng 30 phút/1 giờ/hết ngày, chặn giờ | Thực đơn đầy đủ, nhập Excel/CSV |
| Món Còn bán / Tạm hết; sửa giá nhanh | Hồ sơ quán, ảnh |
| Nhận lời mời vào quán | Nhân viên & quyền; đối soát; thống kê |

Đơn chưa xác nhận: nhắc sau 30 phút; tự huỷ và hoàn cọc sau 2 giờ *(đề xuất)*.

## 7. Thực đơn / Dịch vụ

- Mỗi mục: **tên, giá** (bắt buộc); ảnh, mô tả, nhóm (tuỳ chọn); **Còn bán / Tạm hết**.
- Thêm, sửa, xoá (có hỏi lại), nhân bản, kéo sắp xếp, nhập file Excel/CSV.
- Tên mục đổi theo loại quán: “Thực đơn” hoặc “Dịch vụ” (vé, gói, phòng).
- Xoá món không ảnh hưởng đơn cũ: đơn giữ tên và giá lúc đặt.
- **Không làm ở bước này:** tuỳ chọn món (size/topping), giảm giá theo món, tồn kho.

## 8. Đặt bàn, đặt món trước, cọc (R7)

- Khách chọn ngày, giờ (theo khung quán mở), số người; **đặt món trước tuỳ chọn** nếu quán bật. Món **trả tại quán**; GoWander chỉ gửi danh sách món kèm đơn.
- Quán cấu hình: khung giờ theo tuần + sức chứa từng ô, thời lượng mỗi lượt, đặt trước tối thiểu/tối đa, số khách mỗi bàn, ngày nghỉ, loại bàn.
- **Cọc:** không cọc (giữ bàn 15 phút) · cọc theo người · cọc % hoá đơn dự kiến; ngưỡng số khách áp dụng; cọc trừ vào hoá đơn.
- **Huỷ và hoàn cọc** (ví dụ mặc định): ≥ 24 giờ hoàn 100%; 2–24 giờ hoàn 50%; < 2 giờ hoặc không đến: không hoàn; **quán huỷ: hoàn 100% cho khách**.
- Chính sách được **chụp lại lúc khách đặt**; quán đổi chính sách sau đó không ảnh hưởng đơn cũ.
- Khách xem đơn ở **Hồ sơ → Đặt bàn của tôi**; màn huỷ hiện rõ số tiền được hoàn trước khi xác nhận.
- Thu cọc qua cổng thanh toán; đối soát hằng tuần; phí GoWander (ví dụ 5%) *(chưa chốt; phụ thuộc pháp lý mục 4)*.

## Câu hỏi mở

1. Mời nhân viên: chủ quán tự mời (đang vẽ) hay chỉ Admin cấp? Chủ dự án **chưa chốt rõ** giữa hai cách.
2. Các con số đề xuất (khoá tài khoản, hết hạn mã, tự huỷ đơn, phí %, xoá giấy tờ 30 ngày) cần chốt.
3. Có thu cọc qua GoWander không, hay chỉ ghi nhận cọc quán tự thu? (phụ thuộc pháp lý TMĐT)
4. Cổng thanh toán, nhà cung cấp SMS/SMTP.
5. **Phí đối tác — chưa chốt, cần nghiên cứu** (xem `../06-research/README.md`, mục “Phí đối tác”). Bản vẽ chỉ có giữ chỗ DEMO: gói Cơ bản 0 ₫, gói Pro chưa có giá, phí 5% trên cọc. Hướng đang cân nhắc: không thu phí đăng ký; Cơ bản 0 ₫; Pro trả theo tháng/năm; phí % chỉ trên cọc thu qua GoWander, không trên tiền món.
