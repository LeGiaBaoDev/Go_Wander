# 08 — Tài khoản, phân quyền, đối tác, đặt trước

- **Trạng thái:** Proposed — chờ chủ dự án duyệt; phần pháp lý **chưa được luật sư kiểm chứng**
- **Ngày:** 18-11-2026
- **Chủ sở hữu:** chủ dự án
- **Liên quan:** `../03-decisions/ADR-0002-account-roles.md`, `10-business-model.md`, `07-design-canvas.md`, `06-saved-and-contributed-places.md`, `../00-foundation/02-scope.md`, `../06-research/README.md`

## TL;DR

- **Ba loại đăng nhập, một app:** tài khoản **cá nhân** (khách), tài khoản **quán** do Admin cấp, tài khoản **Admin** (web nội bộ, 2 lớp). Chi tiết: `../03-decisions/ADR-0002-account-roles.md`.
- **Không đăng ký Partner tự do:** khách gửi **hồ sơ hợp tác kèm giấy tờ**, Admin xác minh (1–3 ngày làm việc), rồi hệ thống **cấp tài khoản quán**: mã Chủ quán (`LL-CQ001`) + **liên kết đặt mật khẩu một lần** gửi tới email đăng ký. Quán đăng nhập **cùng app**, thấy giao diện kinh doanh theo quyền.
- **Đăng nhập quán bằng một mật khẩu:** **mã nhân viên** (đã chứa mã quán, ví dụ `LL-NV007`) + **mật khẩu riêng** của người đó; nhiều thiết bị cùng lúc; thiết bị mới gửi email báo Chủ quán. **Chủ quán** thấy đủ tab và **tự CRUD nhân viên**; Quản lý/Nhân viên thấy ít hơn theo chi nhánh. **Admin có toàn quyền.**
- **Email chỉ để khôi phục mật khẩu và xác định tài khoản thuộc về ai**, không dùng đăng nhập.
- **Chi nhánh:** mỗi chi nhánh có địa chỉ, giờ mở, khung đặt trước và đơn riêng.
- **Đặt trước bắt buộc có tài khoản GoWander, không đặt vãng lai.** Quán chỉ nhận đơn từ khách đã đăng nhập (xem mục 8).
- **Món đặt nhiều:** khách thấy nhãn **“Hot”**; quán có dashboard (Partner web và app) với **lọc theo tháng**, biểu đồ đơn 12 tháng, danh sách **món đặt nhiều** của tháng (so với tháng trước), thống kê theo ca (xem mục 9).
- **Đặt trước** giai đoạn 1: **không thu tiền**, xác nhận qua email, tuỳ chọn món (tên, giá, **ảnh**). Gói tháng (giai đoạn 2) và hoa hồng thanh toán trong app (giai đoạn 3): `10-business-model.md`.

## 1. Mô hình tài khoản

| Loại | Dùng cho | Đăng nhập bằng | Cấp bởi |
|---|---|---|---|
| **Cá nhân** (khách) | App khách | Email hoặc SĐT + mật khẩu | Tự đăng ký |
| **Quán** | Giao diện kinh doanh trong app + Partner web | **Mã nhân viên** + **mật khẩu riêng** (không có mật khẩu quán) | **Admin cấp** sau khi duyệt hồ sơ |
| **Admin** | Web nội bộ | Tài khoản riêng + **bắt buộc 2 lớp** | Owner |

- App có **một màn đăng nhập**; ô đầu nhận email/SĐT hoặc mã nhân viên, hệ thống chọn luồng.
- Tài khoản quán là **tài khoản riêng**, không gộp vào tài khoản cá nhân của người đăng ký; người đó vẫn dùng app khách như cũ.
- **Email đăng ký là email khôi phục**: dùng để đặt lại mật khẩu và xác định tài khoản thuộc về ai; **không** dùng để đăng nhập tài khoản quán. Quán có thể khai thêm **email nhận thông báo đơn** riêng (tuỳ chọn).

### Quên mật khẩu

| Ai | Cách |
|---|---|
| Khách, Admin | Email/SMS như tài khoản thường (Admin: qua Owner/2 lớp) |
| Mật khẩu **Chủ quán** | “Quên mật khẩu” → nhập mã Chủ quán (`LL-CQ001`) → liên kết đặt lại (một lần, hạn 30 phút) gửi tới **email khôi phục**; đặt xong thì đăng xuất mọi thiết bị và gửi email báo |
| **Quản lý/Nhân viên** | Chủ quán (hoặc Quản lý với Nhân viên chi nhánh mình) bấm “Cấp lại” → mật khẩu tạm một lần, hạn 24 giờ, phải đổi ngay |
| **Mất cả email** | Liên hệ Admin, đối chiếu giấy tờ hộ kinh doanh rồi cấp lại liên kết đặt mật khẩu; Admin không thấy mật khẩu |
| **Đổi email khôi phục** | Xác nhận ở email cũ **và** Admin duyệt lại |

- Sai 5 lần → khoá **15 phút** (theo thiết bị/IP và tài khoản), không khoá vĩnh viễn; gửi email báo khi có nhiều lần thất bại. Email báo Chủ quán khi đổi mật khẩu, đổi email, hoặc có mã đăng nhập trên thiết bị mới (Chủ quán bật được chế độ phải duyệt máy mới, tắt mặc định).

## 2. Vai trò trong quán và giao diện

| Vai trò | Mã | Thấy gì |
|---|---|---|
| **Chủ quán** | `CQ` | **Đầy đủ tab:** đơn, lịch, hồ sơ quán, chi nhánh, nhân viên, thiết bị, báo cáo, cấu hình |
| **Quản lý** | `QL` (theo chi nhánh) | Đơn, lịch, hồ sơ chi nhánh mình; quản lý nhân viên chi nhánh mình |
| **Nhân viên** | `NV` (theo chi nhánh) | Đơn và lịch chi nhánh mình: xác nhận, đổi giờ, từ chối, đã đến |

- **Quyền Chủ quán chỉ có khi vào bằng mã `CQ` + mật khẩu Chủ quán.** Mã `NV`/`QL` không bao giờ mở tab Chủ quán.
- **Chi nhánh:** mỗi chi nhánh có địa chỉ, giờ mở, khung đặt trước và đơn riêng. Chủ quán thấy tất cả chi nhánh.
- **Nhiều thiết bị cùng lúc:** Chủ quán xem danh sách thiết bị (tên máy, lần dùng cuối, chi nhánh) và đăng xuất từng máy hoặc tất cả. Giới hạn số thiết bị/chi nhánh/nhân viên theo gói (`10-business-model.md`, `ASSUMED`).
- Hai thiết bị thao tác một đơn cùng lúc: người đến sau thấy “đơn vừa được cập nhật” và tải lại.
- Mỗi thao tác ghi **mã nhân viên** đang đăng nhập (đối chiếu theo ca, nhật ký). Dữ liệu theo người là để **truy vết**, chưa phải chứng cứ pháp lý. Không có mật khẩu dùng chung: nhân viên nghỉ thì khoá mã của họ.

## 2b. Mã nhân viên và CRUD nhân viên

### Mã nhân viên
- Dạng `<mã tài khoản>-<vai trò><số thứ tự>`: `LL-CQ001`, `LL-QL002`, `LL-NV007`. Số thứ tự tăng theo quán, **không tái sử dụng** sau khi xoá; mã không đổi khi đổi vai trò (quyền theo vai trò hiện tại).
- **Chỉ để định danh và phân quyền**, không dùng một mình để đăng nhập; bí mật nằm ở mật khẩu riêng (6–128 ký tự, cho phép 6 chữ số, chặn mật khẩu phổ biến, không hiện lại).

### CRUD nhân viên (giao diện Chủ quán)
| Việc | Ai làm |
|---|---|
| Thêm (họ tên, SĐT, vai trò, chi nhánh, ca làm; hệ thống sinh mã) | **Chủ quán**; Quản lý chỉ thêm Nhân viên chi nhánh mình |
| Xem danh sách, trạng thái, lần đăng nhập gần nhất | Chủ quán; Quản lý (chi nhánh mình) |
| Sửa, đổi chi nhánh/vai trò | Chủ quán; Quản lý chỉ sửa Nhân viên chi nhánh mình |
| Cấp lại mật khẩu, khoá/mở khoá | Chủ quán; Quản lý với Nhân viên chi nhánh mình |
| Xoá (giữ tên và mã trong nhật ký cũ) | **Chủ quán** |
| Nhật ký thao tác theo người | Chủ quán; Quản lý (chi nhánh mình) |
| **Toàn quyền** mọi tài khoản quán, nhân viên, nhật ký | **Admin** |

- Quản lý không tạo/đổi được Chủ quán hay Quản lý khác (chống leo quyền).
- SĐT, ca làm của nhân viên là dữ liệu cá nhân: chỉ Chủ quán/Quản lý xem; có chính sách xoá khi nghỉ việc.

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

- **Khi duyệt hồ sơ:** tạo tài khoản quán (mã quán, doanh nghiệp, chi nhánh đầu tiên, gói, Chủ quán `LL-CQ001`), gửi mã Chủ quán và **liên kết đặt mật khẩu một lần (24 giờ)** tới email đăng ký. Không gửi mật khẩu dạng chữ.
- **Chi tiết đối tác:** chi nhánh, nhân viên, thiết bị, nhật ký; Admin có **toàn quyền**: cấp lại liên kết đặt mật khẩu, khoá, đổi gói, đổi email khôi phục (sau khi đối chiếu giấy tờ), gỡ.
- Trạng thái quán: Chờ kích hoạt → Hoạt động → (Đã khoá).

## 6. App: giao diện quán

Quán đăng nhập **cùng màn đăng nhập của app** bằng mã nhân viên; giao diện kinh doanh thay cho giao diện khách. Thanh điều hướng riêng (tối) để không nhầm với chế độ khách.

- **Đăng nhập:** mã nhân viên + mật khẩu riêng → quyền theo vai trò (mục 2). Nhiều thiết bị cùng lúc.
- **Máy dùng chung của quán:** nút **“Đổi người đứng ca”** đăng xuất người hiện tại và mở lại ô đăng nhập; máy nhớ vài mã gần đây để chọn nhanh, vẫn phải nhập mật khẩu.
- Màn **Thiết bị đang đăng nhập** (Chủ quán): đăng xuất từng máy hoặc tất cả.

| Việc hằng ngày (mọi vai trò theo quyền) | Chỉ Chủ quán |
|---|---|
| Đơn đặt trước hôm nay: xác nhận, từ chối (có lý do), đổi giờ, đã đến / không đến | Nhân viên (CRUD), chi nhánh, thiết bị |
| Bật/tắt nhận khách, tạm ngưng, chặn giờ | Hồ sơ quán, khung giờ, cấu hình |
| Món đặt trước: Còn bán / Tạm hết, sửa giá nhanh | Báo cáo theo ca và theo người |

Đơn chưa xác nhận: nhắc sau 30 phút; tự huỷ sau 2 giờ *(đề xuất, `ASSUMED`)*.

## 7. Món đặt trước

- Mỗi món: **tên, giá, ảnh (bắt buộc)**; mô tả, nhóm (tuỳ chọn); **Còn bán / Tạm hết**. Thực đơn đầy đủ/nhập Excel không làm ở giai đoạn 1.
- Thêm, sửa, xoá (có hỏi lại), nhân bản. Xoá món không ảnh hưởng đơn cũ: đơn giữ tên và giá lúc đặt.
- **Không làm:** tuỳ chọn món (size/topping), giảm giá theo món, tồn kho.

## 8. Đặt trước theo giai đoạn

- **Giai đoạn 1 (không thu tiền):** khách chọn quán/chi nhánh, ngày, giờ (theo khung quán mở), số người; tuỳ chọn món đặt trước (trả tại quán). Hệ thống gửi **email xác nhận** cho khách và quán; quán xác nhận, đổi giờ hoặc từ chối → email kết quả; email nhắc lịch; khách tự huỷ. **Không cọc, không hoàn cọc, không đối soát tiền.**
- **Bắt buộc tài khoản:** chỉ khách **đã đăng nhập** (email đã xác minh) mới đặt được; **không có đặt vãng lai**, quán không tự tạo đơn thay khách. Khách chưa đăng nhập thấy bảng “Đăng nhập để đặt trước” (bản đồ và tìm tuyến vẫn dùng được).
- **Chống đặt ảo:** giới hạn số đơn mỗi khách mỗi ngày, quán đánh dấu khách không đến, tạm chặn khi không đến nhiều lần (`ASSUMED`, chưa chốt).
- **Đối chiếu theo ca:** mỗi thao tác gắn mã nhân viên; quán xem đơn theo ca/người; **Admin** xem tất cả quán.
- **Giai đoạn 2:** gói theo tháng (nhiều chi nhánh, thiết bị, nhân viên, báo cáo).
- **Giai đoạn 3:** thanh toán trong app qua cổng có giấy phép, hoa hồng; màn cọc, hoàn cọc, đối soát đã vẽ được giữ lại cho giai đoạn này, đánh dấu “Sắp có”. Chi tiết: `10-business-model.md`.

## 9. Dashboard quán và món đặt nhiều

- **Ai thấy:** Chủ quán thấy mọi chi nhánh; Quản lý chỉ chi nhánh mình; Nhân viên **không** thấy dashboard. Có ở **Partner web** (Tổng quan) và **app** (tab Thống kê).
- **Nội dung (giai đoạn 1, không có tiền):** **lọc theo tháng** (tháng, quý, cả năm); số đơn, đã đến, không đến, thời gian xác nhận trung bình, “tháng tốt nhất năm”; **biểu đồ cột 12 tháng** (đơn và không đến, chạm cột để chọn tháng); **giờ đông nhất** trong tháng; **5 món đặt nhiều** của tháng kèm thay đổi so với tháng trước (tăng, giảm, mới vào); số đơn xác nhận theo ca/mã nhân viên; xuất Excel. Số trong bản vẽ là `DEMO`.
- **Tên gọi:** phía quán dùng **“Món đặt nhiều”**, không dùng chữ “Top”; phía khách chỉ hiện nhãn **“Hot”**.
- **Món đặt nhiều:** đếm lượt món trong các đơn **đã xác nhận hoặc khách đã đến** (không đếm đơn huỷ hay không đến), theo tháng đang chọn, chỉ từ khách có tài khoản.
- **Cho khách:** nhãn **“Hot”** cho các món đặt nhiều nhất (tối đa 5) và xếp lên đầu danh sách món đặt trước khi quán có từ **20 đơn** trong cửa sổ (`ASSUMED`); chưa đủ dữ liệu thì không hiện gì, không hiện số giả.
- **Tên gọi trung thực:** giai đoạn 1 nói “đặt nhiều”/“Hot”, không gọi “bán chạy” (không có doanh thu thật). Nếu sau này có xếp hạng trả tiền thì phải ghi “Được tài trợ”, không lẫn với top tự nhiên (nguyên tắc 3).
- **Chống thổi phồng:** không tính đơn không có tài khoản hợp lệ; giới hạn đơn mỗi khách mỗi ngày.

## Câu hỏi mở

1. ~~Hai lớp hay một mật khẩu~~: **đã chốt** một mật khẩu riêng mỗi người. ~~Duyệt máy mới~~: **đã chốt tắt mặc định** (18-11-2026), chỉ gửi email báo Chủ quán; bật theo từng quán là tuỳ chọn.
2. Giới hạn thiết bị, chi nhánh, nhân viên theo gói và các con số `ASSUMED` (khoá 15 phút, hết hạn liên kết, tự huỷ đơn, xoá giấy tờ 30 ngày).
3. Affiliate với ai (link ngoài hay đối tác trực tiếp); mức chống đặt ảo.
4. Cổng thanh toán và luật sư cho giai đoạn 3; nhà cung cấp dịch vụ gửi email và SMS (chưa chọn).
5. **Phí đối tác và gói — chưa chốt, cần nghiên cứu** (xem `10-business-model.md`, `../06-research/README.md` mục “Phí đối tác”). Số trong bản vẽ chỉ là `DEMO`.
6. Luật lao động, quyền riêng tư dữ liệu nhân viên (SĐT, ca) và DPIA — **chưa có luật sư rà**.
