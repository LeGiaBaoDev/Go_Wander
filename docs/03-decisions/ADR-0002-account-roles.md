# ADR-0002 — Tài khoản khách, tài khoản quán do Admin cấp, mỗi người đăng nhập bằng mã nhân viên + mật khẩu riêng

- **Trạng thái:** Accepted (chủ dự án chốt 18-11-2026) cho phần tài khoản quán; các mục ghi “chưa chốt” bên dưới vẫn mở
- **Ngày:** 18-11-2026 (viết lại; thay bản nháp trước dùng “một kiểu tài khoản cá nhân + vai trò”)
- **Người quyết định:** chủ dự án
- **Liên quan:** `../01-product/08-accounts-roles-partners.md`, `../01-product/10-business-model.md`, `../00-foundation/02-scope.md`, `ADR-0001-web-stack.md`, `ADR-0003-api-contract-auth-rbac.md`, `ADR-0004-data-model-postgis.md`

## Bối cảnh

- Partner cần nhận đặt trước, quản lý quán, và (giai đoạn sau) thanh toán. Chủ quán và nhân viên không phải lúc nào cũng ngồi trước máy tính.
- Chủ dự án muốn: **một app**; quán **không tự đăng ký tự do** mà nộp giấy tờ để xác minh; sau khi duyệt, hệ thống **cấp tài khoản cho quán**; quán đăng nhập **cùng app** và thấy giao diện kinh doanh theo quyền được cấp.
- Chủ quán cần **nhiều thiết bị đăng nhập cùng lúc** (chủ và các chi nhánh) và **tự tạo, sửa, xoá nhân viên**.
- Mỗi nhân viên cần **mã nhân viên** (định danh) và **mật khẩu riêng**; quyền hiển thị theo người đó.
- Email dùng để **khôi phục mật khẩu và xác định tài khoản thuộc về ai**, **không** dùng để đăng nhập tài khoản quán.
- Hình thức kinh doanh theo giai đoạn (đặt trước miễn phí → gói tháng → hoa hồng) nằm ở `../01-product/10-business-model.md`; ADR này chỉ lo tài khoản và quyền.

## Quyết định

### 1. Ba loại đối tượng đăng nhập, một app

| Loại | Dùng cho | Đăng nhập bằng | Cấp bởi |
|---|---|---|---|
| **Tài khoản cá nhân** (khách) | App khách (bản đồ, lưu, đặt trước) | Email hoặc số điện thoại + mật khẩu | Tự đăng ký |
| **Tài khoản quán** | Giao diện kinh doanh trong app và Partner web | **Mã nhân viên** (chứa mã quán) + **mật khẩu riêng** (mục 3) | **Admin cấp** sau khi duyệt hồ sơ |
| **Tài khoản Admin** | Web nội bộ GoWander | Tài khoản riêng + **bắt buộc 2 lớp** | Owner cấp |

- App chỉ có **một màn đăng nhập**. Ô đầu nhận email/SĐT (khách) hoặc mã nhân viên (quán); hệ thống chọn đúng luồng.
- Tài khoản quán **không** là vai trò gắn vào tài khoản cá nhân. Người đăng ký doanh nghiệp vẫn giữ tài khoản cá nhân của mình, không bị đổi hay gộp.

### 2. Từ đăng ký đến tài khoản quán

1. Khách (đang dùng tài khoản cá nhân) bấm “Đăng ký hợp tác”, điền thông tin quán và nộp giấy tờ (`08`, mục 4).
2. Admin xác minh trong 1–3 ngày làm việc, báo kết quả qua email.
3. Khi duyệt, hệ thống tạo **tài khoản quán** gồm: **mã quán** (do Admin đặt, 2–6 chữ in hoa, duy nhất, ví dụ `LL`), doanh nghiệp, chi nhánh đầu tiên, gói, và **Chủ quán đầu tiên** với mã `LL-CQ001`.
4. Hệ thống gửi **mã Chủ quán** (`LL-CQ001`) và **liên kết đặt mật khẩu dùng một lần (hạn 24 giờ)** tới **email đăng ký**. **Không gửi mật khẩu dạng chữ.**
5. Email đăng ký trở thành **email khôi phục** của tài khoản quán. Quán có thể khai thêm **email nhận thông báo đơn** riêng (tuỳ chọn) để người nhận đơn không có quyền khôi phục tài khoản.

### 3. Đăng nhập một mật khẩu và quyền (chốt 18-11-2026)

- Mỗi người trong quán đăng nhập bằng **mã nhân viên** (đã chứa mã quán, ví dụ `LL-NV007`) + **mật khẩu riêng của người đó**. **Không có mật khẩu quán dùng chung.**
- Đăng nhập được trên **nhiều thiết bị cùng lúc**. Trên máy dùng chung của quán, nút **“Đổi người đứng ca”** đăng xuất người hiện tại và mở lại ô mã nhân viên + mật khẩu.
- Quyền hiển thị theo vai trò của mã:

| Vai trò | Thấy gì |
|---|---|
| **Chủ quán** (`CQ`) | **Đầy đủ tab:** đơn, lịch, hồ sơ quán, chi nhánh, nhân viên, thiết bị, báo cáo, cấu hình |
| **Quản lý** (`QL`, theo chi nhánh) | Đơn, lịch, hồ sơ chi nhánh mình; quản lý nhân viên chi nhánh mình |
| **Nhân viên** (`NV`, theo chi nhánh) | Đơn và lịch chi nhánh mình: xác nhận, đổi giờ, từ chối, đã đến |

- **Quyền Chủ quán chỉ có khi vào bằng mã `CQ` + mật khẩu của Chủ quán.** Mã `NV`/`QL` không bao giờ mở tab Chủ quán. Khuyến nghị Chủ quán bật 2FA cho việc nhạy cảm (nhân viên, cấu hình).
- **Thiết bị mới:** lần đầu một mã đăng nhập trên máy lạ, hệ thống **gửi email báo Chủ quán** (mặc định). Chủ quán có thể bật thêm chế độ **phải duyệt máy mới** (tuỳ chọn, tắt mặc định).
- Mỗi thao tác ghi **mã nhân viên** đang đăng nhập (đối soát đặt trước và nhật ký theo người).

### 4. Mã nhân viên

- Dạng `<mã tài khoản>-<vai trò><số thứ tự>`: `LL-CQ001`, `LL-QL002`, `LL-NV007`. Số thứ tự tăng theo quán, cấp phát trong transaction, **không tái sử dụng** sau khi xoá.
- Mã **không đổi khi đổi vai trò**; quyền luôn theo vai trò hiện tại; **không suy quyền từ chữ trong mã**.
- Mã có quy luật nên đoán được; vì vậy bí mật nằm ở mật khẩu và các giới hạn mục 6.

### 5. Chi nhánh và nhiều thiết bị

- Một doanh nghiệp có **một hoặc nhiều chi nhánh**. Mỗi chi nhánh có **địa chỉ, giờ mở, khung giờ đặt trước và đơn riêng**.
- Chủ quán thấy mọi chi nhánh; Quản lý và Nhân viên chỉ thấy chi nhánh được gán.
- **Nhiều thiết bị cùng lúc:** Chủ quán xem danh sách thiết bị đang đăng nhập (tên máy, lần dùng cuối, chi nhánh) và **đăng xuất từng máy hoặc tất cả**. Giới hạn số thiết bị và chi nhánh theo gói (`10-business-model.md`).
- Hai thiết bị thao tác một đơn cùng lúc: đơn có phiên bản (optimistic locking); người đến sau thấy “đơn vừa được cập nhật” và tải lại.

### 6. Quên mật khẩu, khoá, an toàn

| Trường hợp | Cách xử lý |
|---|---|
| Quên mật khẩu **Chủ quán** | “Quên mật khẩu” → nhập mã Chủ quán (`LL-CQ001`) → liên kết đặt lại (một lần, hạn 30 phút) gửi tới **email khôi phục**; đặt xong thì đăng xuất mọi thiết bị và gửi email báo |
| Quên mật khẩu **Quản lý/Nhân viên** | Chủ quán (hoặc Quản lý, với Nhân viên chi nhánh mình) bấm “Cấp lại” → **mật khẩu tạm một lần**, hiện đúng một lần cho người cấp, hạn 24 giờ; người đó phải đổi ngay. Người cấp **không xem được** mật khẩu cũ |
| **Mất cả email khôi phục** | Liên hệ Admin; Admin đối chiếu giấy tờ hộ kinh doanh rồi cấp lại liên kết đặt mật khẩu. Admin **không bao giờ thấy** mật khẩu |
| **Đổi email khôi phục** | Xác nhận ở email cũ **và** Admin duyệt lại (đối chiếu giấy tờ); gửi email báo cả hai địa chỉ |
| **Nhập sai nhiều lần** | 5 lần sai → khoá **15 phút** theo thiết bị/IP và theo tài khoản (không khoá vĩnh viễn, tránh bị cố tình khoá quán); gửi email báo khi có nhiều lần thất bại. Nhân viên bị khoá lặp lại: Quản lý/Chủ quán mở |

- Mật khẩu: 6–128 ký tự, cho phép 6 chữ số; chặn mật khẩu phổ biến; băm Argon2id; không hiện lại, không ghi log. Mật khẩu ngắn là mức tối thiểu thấp; bù bằng khoá, giới hạn tốc độ, nhật ký, email báo và khuyến nghị 2FA cho Chủ quán trên Partner web.
- Gửi email báo Chủ quán khi **đổi mật khẩu, đổi email, có mã đăng nhập trên thiết bị mới**.

### 7. CRUD nhân viên và quyền của Admin

| Việc | Ai làm |
|---|---|
| Thêm nhân viên (họ tên, SĐT, vai trò, chi nhánh, ca làm; hệ thống sinh mã) | **Chủ quán**; Quản lý chỉ thêm Nhân viên ở chi nhánh mình |
| Xem danh sách, trạng thái, lần đăng nhập gần nhất | Chủ quán; Quản lý (chi nhánh mình) |
| Sửa thông tin, đổi chi nhánh/vai trò | Chủ quán; Quản lý chỉ sửa Nhân viên chi nhánh mình |
| Cấp lại mật khẩu; khoá/mở khoá | Chủ quán; Quản lý với Nhân viên chi nhánh mình |
| Xoá | **Chủ quán** |
| Xem nhật ký thao tác theo người | Chủ quán; Quản lý (chi nhánh mình) |
| **Toàn quyền** với mọi tài khoản quán, nhân viên, nhật ký; cấp lại, khoá, gỡ | **Admin** |

- Quản lý **không** tạo hay đổi được Chủ quán hoặc Quản lý khác (chống leo quyền).
- Xoá nhân viên: chặn đăng nhập, **giữ tên và mã trong nhật ký cũ**; mã không cấp lại.
- Mọi thao tác của Admin lên tài khoản quán ghi nhật ký (ai, lúc nào, làm gì).

### 8. Kiểm quyền ở server, nhật ký

- **Kiểm quyền ở server** (theo tài khoản quán, chi nhánh, vai trò); client chỉ ẩn/hiện. Gọi API thiếu quyền phải bị chặn.
- Hành động nhạy cảm ghi nhật ký: đổi thông tin quán, đổi giờ/khung đặt trước, xác nhận/từ chối/huỷ đơn, thêm/xoá/đổi nhân viên, cấp lại mật khẩu, đổi email, đăng xuất thiết bị, xem báo cáo.

### 9. Giới hạn đã biết của thiết kế

- **Không có mật khẩu dùng chung:** nhân viên nghỉ thì khoá mã của người đó, các người khác không bị ảnh hưởng. Nhật ký theo người **giúp truy vết, chưa phải chứng cứ pháp lý** (người trong quán vẫn có thể lộ mật khẩu cho nhau).
- Mã có quy luật nên **tên đăng nhập đoán được**; bảo vệ bằng mật khẩu riêng, khoá 15 phút, email báo thiết bị mới.
- Quyền Chủ quán được bảo vệ bởi mật khẩu riêng của Chủ quán chứ không phải bởi mã.
- Email khôi phục là điểm tấn công lớn nhất (chiếm email là khôi phục được tài khoản); vì vậy mục 6 bắt buộc Admin duyệt khi đổi email.

### 10. Dữ liệu cá nhân và pháp lý

- Họ tên, SĐT, ca làm của nhân viên là dữ liệu cá nhân: chỉ Chủ quán/Quản lý xem; có mục đích, thời hạn lưu, quyền xoá; cần DPIA.
- Luật lao động (quản lý ca, xử lý dữ liệu nhân viên), nghĩa vụ xác minh đối tác và luật bảo vệ dữ liệu cá nhân **chưa có luật sư rà**; không khẳng định “đã tuân thủ”.
- Không dùng dữ liệu nhân viên cho quảng cáo hay affiliate; không bán.

## Lý do

1. **Tài khoản quán do Admin cấp** giữ được bước xác minh (chặn quán giả, sai chủ) mà vẫn tách khỏi tài khoản cá nhân của người đăng ký.
2. **Một app, giao diện theo quyền** tránh hai app và hai lần đăng nhập; quán dùng chung công cụ với khách.
3. **Một mật khẩu riêng mỗi người** đơn giản cho quán (chủ dự án thấy hai lớp rắc rối), bỏ bí mật dùng chung, và quyền Chủ quán vẫn không phụ thuộc thứ ai cũng đoán được.
4. **Email tách khỏi đăng nhập** cho phép khôi phục an toàn mà không đẩy mật khẩu qua email.
5. **Chủ quán tự quản lý nhân viên** để Admin không phải tạo từng người; Admin giữ toàn quyền để can thiệp khi có sự cố.

## Phương án đã cân nhắc

| Phương án | Vì sao không chọn |
|---|---|
| Một kiểu tài khoản cá nhân, quán chỉ là vai trò (bản nháp trước) | Quán phải có tài khoản cấp riêng sau xác minh; nhân viên không cần tài khoản cá nhân; chủ dự án đổi hướng |
| Một tài khoản quán dùng chung, không phân người | Không biết ai làm gì; không tách được quyền Chủ quán |
| Chỉ dùng mã nhân viên để đăng nhập | Mã có quy luật, đoán được |
| Hai lớp: mật khẩu quán + mật khẩu riêng (bản 18-11 trước) | Nhập hai mật khẩu rắc rối; mật khẩu quán là bí mật dùng chung. Chủ dự án bỏ (18-11-2026) |
| Đăng nhập quán bằng email | Chủ dự án muốn email chỉ để khôi phục và định danh |
| Gửi mật khẩu qua email | Dễ lộ; dùng liên kết đặt mật khẩu một lần |
| Khoá vĩnh viễn sau 5 lần sai | Kẻ xấu cố tình khoá quán; dùng khoá 15 phút |
| App Business riêng | Hai app; trùng đăng nhập |

## Hệ quả

### Tích cực
- Luồng rõ: khách → đăng ký hợp tác → Admin duyệt → tài khoản quán → quán tự quản lý nhân viên.
- Nhiều thiết bị và chi nhánh; truy vết theo mã nhân viên cho đối soát đặt trước.
- Nâng gói (nhiều chi nhánh, thiết bị, nhân viên) khớp mô hình kinh doanh.

### Tiêu cực / chi phí
- Phải làm: tài khoản quán, chi nhánh, nhân viên, thiết bị/phiên, khôi phục qua email, đặt mật khẩu bằng liên kết, khoá và email báo.
- Màn hình phải vẽ lại: đăng nhập quán một mật khẩu, quản lý nhân viên, thiết bị, chi nhánh (`10`, `08`).
- ADR-0003 và ADR-0004 cần cập nhật theo ADR này.

### Rủi ro
- Mật khẩu ngắn + mã có quy luật: dựa vào khoá, giới hạn tốc độ, email báo; cần test và giám sát.
- Chiếm email khôi phục: bắt buộc Admin duyệt khi đổi email.
- Dữ liệu nhân viên là dữ liệu cá nhân; chưa có DPIA hay luật sư.

## Việc tiếp theo

- [ ] Cập nhật `ADR-0003` (đăng nhập bằng mã nhân viên + mật khẩu riêng, phiên nhiều thiết bị, khôi phục qua email) và `ADR-0004` (`shop_account`, `branch`, `staff`, `device_session`, email khôi phục).
- [ ] Viết `../02-architecture/03-data-model.md` theo mô hình này.
- [ ] Vẽ lại canvas: đăng nhập quán một mật khẩu, nhân viên (CRUD), thiết bị, chi nhánh, đặt mật khẩu qua liên kết.
- [x] **Đã chốt:** một mật khẩu riêng mỗi người, bỏ mật khẩu quán; thiết bị mới gửi email báo Chủ quán (duyệt máy là tuỳ chọn).
- [ ] **Chưa chốt:** giới hạn thiết bị/chi nhánh/nhân viên theo gói (xem `10`, số là `ASSUMED`).
- [ ] Luật sư rà dữ liệu nhân viên, luật lao động và DPIA.
- [ ] Test: khoá 15 phút, đặt lại qua email, mật khẩu tạm hết hạn, Quản lý không cấp được cho Quản lý/Chủ quán, mã `NV` không vào được tab Chủ quán.

## Tham khảo

- Trao đổi với chủ dự án 18-11-2026 (mô hình tài khoản quán, nhiều thiết bị và chi nhánh, mã nhân viên, mô hình kinh doanh theo giai đoạn).
- `../01-product/08-accounts-roles-partners.md` — luồng hồ sơ đối tác, giao diện quán.
- `../01-product/10-business-model.md` — gói và hoa hồng theo giai đoạn.
