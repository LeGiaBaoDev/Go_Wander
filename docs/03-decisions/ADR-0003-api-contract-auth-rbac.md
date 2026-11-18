# ADR-0003 — Hợp đồng API, đăng nhập, phiên và RBAC

- **Trạng thái:** Proposed — chờ chủ dự án chốt
- **Ngày:** 18-11-2026 (thân viết lại theo ADR-0002 mới: tài khoản quán do Admin cấp, mã nhân viên + mật khẩu riêng)
- **Người quyết định:** chủ dự án (chờ chốt)
- **Liên quan:** `ADR-0001-web-stack.md`, `ADR-0002-account-roles.md`, `../02-architecture/01-system-context.md`, `../02-architecture/02-containers.md`, `../01-product/08-accounts-roles-partners.md`, `../04-development/02-conventions.md`

> **Viết lại 18-11-2026.** Bản trước dựa trên `shop_membership` (một kiểu tài khoản + vai trò) và mật khẩu quán dùng chung. ADR-0002 mới thay bằng: tài khoản quán do Admin cấp; **mỗi người đăng nhập bằng mã nhân viên + mật khẩu riêng** (không có mật khẩu quán); nhiều thiết bị cùng lúc; chi nhánh; email chỉ để khôi phục. ADR này viết lại realm xác thực, phiên và RBAC theo đó.

## Bối cảnh

GoWander có ba loại client với nhu cầu xác thực khác nhau:

| Client | R1 | Cách dùng | Ràng buộc |
|---|---|---|---|
| Client Mobile (KMP) | Guest-first; tài khoản khách từ R1.5; **cùng app có đăng nhập quán** | Bearer token (một realm cho khách, một realm cho quán) | Token trong Keychain/Keystore; không lưu plan/saved trên server mặc định |
| Admin Web (React/Vite) | Có | Session cookie | RBAC, bắt buộc 2 lớp, chống CSRF, không token trong `localStorage` |
| Partner Web (React/Vite) | R6 | Session cookie | Đăng nhập bằng mã nhân viên; quyền theo (quán, chi nhánh, vai trò) |
| Showcase Web (Next.js, static) | Có | Public | Không form thu PII ở R1; không cần đăng nhập |

Các quyết định đã có:

- Một backend Go (modular monolith), client không gọi DB/OTP/nhà cung cấp ngoài (`01-system-context.md`).
- OpenAPI là nguồn sự thật của hợp đồng API (`03-principles.md` nguyên tắc 6; `02-conventions.md`).
- Lỗi có một dạng thống nhất, có `requestId` (`02-conventions.md`).
- Ngoài vùng hỗ trợ phải trả `UNSUPPORTED`, không bịa tuyến (`02-scope.md`).
- Guest-first, local-first, privacy by default (`03-principles.md` nguyên tắc 7).
- Admin là tài khoản riêng, bắt buộc hai lớp. Tài khoản quán do Admin cấp; mỗi người một mã nhân viên + mật khẩu riêng (`ADR-0002`, `08-accounts-roles-partners.md`).

Chưa có quyết định cho: phiên dài bao lâu, token hay cookie, RBAC kiểm ở đâu, mã lỗi nào là ổn định, làm sao giữ OpenAPI và client khớp nhau. Nếu để mỗi client tự đoán, sẽ lặp lại lỗi cũ: app phụ thuộc schema OTP, lỗi mỗi nơi một kiểu, quyền chỉ ẩn ở UI.

## Quyết định

### 1. REST/JSON, OpenAPI 3.1, phiên bản trong URL

- API là **REST/JSON** trên HTTPS; mọi endpoint sản phẩm nằm dưới `/api/v1`.
- Endpoint vận hành `/health/live`, `/health/ready` không thuộc `/api/v1` và không trả dữ liệu nghiệp vụ.
- **OpenAPI 3.1** tại `api/openapi.yaml` là nguồn sự thật. Code server, client web và client mobile phải khớp spec; không có DTO viết tay ngoài spec.
- Thay đổi trong `/api/v1` chỉ được **thêm** (additive). Bỏ field/đổi nghĩa/đổi kiểu là breaking change:
  - Thêm endpoint/field mới: giữ `/api/v1`.
  - Breaking change: mở `/api/v2`, giữ `/api/v1` trong thời gian chuyển; ghi `Sunset`/deprecation trong spec và changelog.
- Không dùng GraphQL, gRPC hay RPC kiểu `/doSomething` cho API sản phẩm ở R1. Lý do: client KMP và web đều dùng HTTP/JSON đơn giản; OpenAPI sinh client được; đội nhỏ.

### 2. Ba realm xác thực tách biệt

| Realm | Đối tượng | Cơ chế | Cookie/token | 2 lớp |
|---|---|---|---|---|
| `client` | Khách dùng app (R1.5) | Bearer token | Token trong Keychain/Keystore | Không bắt buộc |
| `partner` | Chủ quán/Quản lý/Nhân viên (R6) | App: bearer token theo thiết bị. Web: session cookie | Keychain/Keystore; `__Host-gw_partner` | Khuyến nghị cho Chủ quán; bắt buộc nếu có thu tiền (giai đoạn 3) |
| `admin` | Admin/Editor GoWander | Session cookie | `__Host-gw_admin` | **Bắt buộc** |

Quy tắc:

- **Không dùng chung cookie/token giữa các realm.** Cookie Admin không mở được endpoint Partner và ngược lại; token `client` không gọi được `/admin/*` hay `/partner/*`.
- Danh tính realm `partner` là **`staff`** (một người trong quán, định danh bằng mã nhân viên), **không** phải `account` cá nhân được gắn vai trò. Một người vừa là khách vừa là nhân viên có hai danh tính độc lập và hai token riêng.
- Admin là tài khoản riêng. Admin không đăng nhập vào realm `partner`; Admin can thiệp tài khoản quán qua `/admin/partners/*` và mọi thao tác được ghi nhật ký.
- Không lưu token/session trong `localStorage` hoặc `sessionStorage`. Web dùng cookie `HttpOnly`; mobile dùng Keychain/Keystore.
- Mọi endpoint ghi rõ realm và quyền cần có trong spec (`security`, mô tả operation).

### 3. Phiên và token

#### 3.1 Client Mobile

- Token là chuỗi ngẫu nhiên **256-bit** (32 byte), không phải JWT; server chỉ lưu **SHA-256** của token, không lưu token gốc.
- Một token cho một thiết bị/phiên; thời hạn **60 ngày trượt** (mỗi lần dùng hợp lệ gia hạn), thu hồi được. *(Đề xuất kế thừa ADR-040 repo cũ; chờ chốt.)*
- Đăng xuất thu hồi token hiện tại; đổi/đặt lại mật khẩu thu hồi **mọi** token của account; xoá tài khoản thu hồi mọi token và anonymize dữ liệu theo chính sách.
- Không dùng refresh token riêng ở R1.5. Nếu sau này cần access token ngắn + refresh token, đó là ADR mới.
- Token gửi qua `Authorization: Bearer <token>`; không đặt trong query string, URL, log hay ảnh.

#### 3.2 Admin Web

- Session cookie `__Host-gw_admin`: `HttpOnly`, `Secure`, `SameSite=Lax`, `Path=/`, không `Domain`.
- **Idle timeout 30 phút**, absolute timeout **12 giờ**. *(Đề xuất; chờ chốt.)*
- Bắt buộc 2 lớp: TOTP (RFC 6238) là mặc định; passkey/WebAuthn là lựa chọn bổ sung sau. Mã dự phòng dùng một lần, lưu dạng băm.
- Hành động nhạy cảm (khôi phục sao lưu, đổi quyền Admin, duyệt hồ sơ đối tác, đổi chủ quán) yêu cầu **step-up**: xác thực lại mật khẩu + 2FA trong vòng 10 phút trước khi thực hiện. *(Đề xuất.)*
- Sai mật khẩu 5 lần liên tiếp → khoá tạm 15 phút; sai 2FA 5 lần → khoá, chỉ Owner mở. *(Đề xuất kế thừa ADR-039/08; chờ chốt.)*
- Cookie CSRF: token double-submit trong cookie không `HttpOnly` riêng (`gw_csrf`) + header `X-CSRF-Token`; kiểm cả `Origin`/`Referer` cho request ghi.

#### 3.3 Quán (app và Partner web)

- **Danh tính:** `staff` đăng nhập bằng **mã nhân viên** (ví dụ `LL-NV007`, đã chứa mã quán) + **mật khẩu riêng**. Không có mật khẩu quán, không đăng nhập bằng email (ADR-0002 mục 3).
- **Phiên theo thiết bị:** mỗi lần đăng nhập tạo một `device_session`. Token là chuỗi ngẫu nhiên 256-bit, opaque, server chỉ lưu SHA-256. App gửi `Authorization: Bearer`; web dùng cookie `__Host-gw_partner` (cùng thuộc tính bảo vệ như Admin). Một nhân viên có thể có nhiều phiên cùng lúc.
- Idle timeout **8 giờ**, absolute **7 ngày**. *(Đề xuất; chờ chốt. Nhân viên giữa ca không muốn đăng nhập lại liên tục.)*
- **“Đổi người đứng ca”** thu hồi phiên hiện tại rồi mở lại ô mã nhân viên + mật khẩu.
- **Chủ quán quản lý thiết bị:** xem danh sách phiên (tên máy, lần dùng cuối, chi nhánh), thu hồi từng phiên hoặc tất cả. Đổi/đặt lại mật khẩu, khoá hoặc xoá một `staff` thu hồi **mọi phiên của `staff` đó**.
- **Thiết bị mới:** lần đầu một `staff` đăng nhập trên `device_id` chưa từng thấy, hệ thống ghi `outbox_event` gửi email báo Chủ quán. Nếu quán bật `require_device_approval` (tắt mặc định), phiên ở trạng thái `pending_approval` và chỉ gọi được endpoint xem trạng thái cho đến khi Chủ quán duyệt.
- **Khoá:** 5 lần sai → khoá 15 phút, tính theo `staff` và theo thiết bị/IP; tự mở, **không khoá vĩnh viễn** (tránh kẻ xấu cố tình khoá quán). Nhiều lần thất bại thì gửi email báo.
- **2 lớp:** không bắt buộc mặc định; khuyến nghị bật cho Chủ quán vì mật khẩu nhân viên chỉ tối thiểu 6 ký tự (mục 4). GoWander có thể bắt buộc khi có thu tiền.
- Mọi request `/partner/shops/{shopId}/...` lấy quán, chi nhánh và vai trò từ phiên và DB, **không tin `shopId`/`branchId` do client gửi** (xem mục 5).
- **Ghi đồng thời:** ghi đơn/thực đơn gửi kèm `version` (hoặc `If-Match`); sai phiên bản → `409 CONFLICT`, client tải lại (ADR-0002 mục 5).

#### 3.4 Quy tắc chung cho cookie

- Chỉ phát cookie qua HTTPS; `Secure` bắt buộc ở production.
- Ở dev trên `localhost`, trình duyệt có thể không nhận cookie `__Host-`/`Secure`; dùng tên cookie dev riêng (ví dụ `gw_admin_dev`) và bật HTTPS cục bộ khi test luồng đăng nhập thật. Không nới `Secure` ở production.
- `SameSite=Lax`; nếu sau này cần nhúng cross-site thì mở ADR riêng, không nới `SameSite=None` mặc định.
- Không đặt session ID trong URL; không log cookie, `Set-Cookie`, `Authorization`.
- CORS chỉ mở cho origin cấu hình (Showcase/Client Web); Admin/Partner chạy cùng origin qua reverse proxy nên không cần CORS rộng.

### 4. Mật khẩu, xác minh và khôi phục

- Băm mật khẩu bằng **Argon2id** (`golang.org/x/crypto/argon2`), tham số khởi điểm OWASP: 19 MiB, 2 vòng, 1 luồng, muối 16 byte; tham số lưu kèm hash để nâng cấp dần. *(Chờ chốt; không dùng bcrypt vì giới hạn 72 byte.)*
- Mật khẩu client: 10–128 ký tự, có ít nhất một chữ và một số; không trùng 3 mật khẩu gần nhất (lưu hash lịch sử). Admin: tối thiểu 12 ký tự. *(Đề xuất; chờ chốt.)*
- **Nhân viên quán (Chủ quán/Quản lý/Nhân viên):** 6–128 ký tự; cho phép **6 chữ số**; chặn mật khẩu phổ biến (`123456`, `000000`, `111111`, `password`…). Ngoại lệ có chủ đích theo ADR-0002 mục 6; bù bằng khoá 15 phút, rate limit, email báo, nhật ký theo mã nhân viên và khuyến nghị 2FA cho Chủ quán.
- **Đặt mật khẩu lần đầu:** hệ thống gửi **liên kết dùng một lần (hạn 24 giờ)** tới email đăng ký; không bao giờ gửi mật khẩu dạng chữ.
- **Quên mật khẩu Chủ quán:** nhập mã Chủ quán → liên kết đặt lại một lần (hạn **30 phút**) tới `recovery_email`; đặt xong thì thu hồi mọi phiên và gửi email báo.
- **Quên mật khẩu Quản lý/Nhân viên:** người có quyền (Chủ quán; Quản lý với Nhân viên chi nhánh mình) bấm “Cấp lại” → **mật khẩu tạm một lần**, hạn 24 giờ, bắt đổi ngay, chỉ hiện đúng một lần cho người cấp, không xem được mật khẩu cũ. Không gửi mật khẩu tạm qua email/SMS; người cấp chuyển trực tiếp.
- **Đổi `recovery_email`:** xác nhận ở email cũ **và** Admin duyệt (đối chiếu giấy tờ); gửi email báo cả hai địa chỉ.
- Không tiết lộ tài khoản có tồn tại hay không: login sai, quên mật khẩu, đăng ký trùng đều trả thông báo trung tính; vẫn băm thử một lần khi account không tồn tại để thời gian phản hồi tương đương.
- Mã xác minh email/khôi phục: 6 số, hết hạn **10 phút**, tối đa **5 lần gửi/giờ/tài khoản**, lưu dạng băm, dùng một lần. *(Đề xuất; chờ chốt.)*
- Đăng nhập Google (R1.5) và Apple (trước App Store, nếu có Google) là **identity provider**, không phải realm riêng: sau khi xác minh, cấp cùng token client. Chỉ liên kết account khi email đã được xác minh; không tự động gộp theo email chưa xác minh.
- Đổi mật khẩu/đặt lại mật khẩu: thu hồi mọi phiên khác; gửi thông báo qua email (không kèm token).
- Xoá tài khoản trong app (R1.5): xác nhận, thu hồi phiên, anonymize dữ liệu công khai (tên hiển thị, ảnh đại diện), giữ dữ liệu kiểm toán cần thiết theo chính sách. Không xoá cứng ngay nếu có nghĩa vụ lưu trữ; ghi rõ trong chính sách.

### 5. RBAC: kiểm ở server, theo đối tượng

#### 5.1 Vai trò

| Realm | Vai trò | Phạm vi |
|---|---|---|
| Client | `client` (ngầm định) | Dữ liệu của chính mình |
| Partner | `owner` (mã `CQ`), `manager` (`QL`), `staff` (`NV`) | `owner`: mọi chi nhánh của tài khoản quán. `manager`/`staff`: chi nhánh được gán (`branch_id`) |
| Admin | `owner`, `ops`, `editor`, `moderator`, `accountant`, `partner_support` | Toàn hệ thống hoặc theo quyền được cấp |

- **Quyền suy từ `role` hiện tại**, không suy từ chữ trong mã nhân viên (mã không đổi khi đổi vai trò).
- Quyền Chủ quán chỉ có khi vào bằng mã `CQ` + mật khẩu của Chủ quán; mã `NV`/`QL` không bao giờ mở được thao tác Chủ quán.
- Quản lý **không** tạo, sửa hay xoá được Chủ quán hoặc Quản lý khác (chống leo quyền). Luôn còn ít nhất một Chủ quán `active` trong mỗi tài khoản quán; service kiểm trong cùng transaction.
- **Admin toàn quyền** với mọi tài khoản quán, nhân viên và nhật ký, qua realm `admin` và có ghi nhật ký. Admin không tự động có quyền realm `partner` và ngược lại.
- Tên quyền kỹ thuật dạng `resource:action` (ví dụ `booking:confirm`, `menu:write`, `staff:manage`, `device:revoke`, `place:write`, `story:publish`, `submission:review`, `backup:restore`).

#### 5.2 Cách kiểm

- **Deny by default.** Middleware chỉ xác định realm và phiên; **một `Authorizer` duy nhất** (`Can(principal, action, resource)`) quyết định quyền, nhận resource đã nạp từ DB (có `shop_account_id`, `branch_id`, `role` của người bị tác động). Không rải `if role == ...` trong handler.
- Kiểm theo đối tượng, không chỉ theo route: `PATCH /partner/shops/{shopId}/menu/{itemId}` phải kiểm phiên thuộc `shopId`, item thuộc `shopId`, chi nhánh trong phạm vi, vai trò được sửa menu.
- Không tin `role`, `shopId`, `branchId`, `staffId`, `accountId` do client gửi; lấy từ phiên và DB.
- Mọi thao tác ghi của Admin/quán ghi audit **trong cùng transaction** với thay đổi (mục 9, mẫu Outbox ở mục 13).
- Có test cho **đường từ chối**: thiếu token, sai realm, thiếu vai trò, sai `shopId`, sai chi nhánh, đối tượng không tồn tại, nhân viên bị khoá/xoá, phiên bị thu hồi, Quản lý cố cấp quyền cho Quản lý/Chủ quán.
- UI chỉ ẩn/hiện; ẩn nút không phải bảo vệ.

#### 5.3 Ma trận quyền tối thiểu

| Hành động | Client | Nhân viên | Quản lý | Chủ quán | Admin |
|---|:-:|:-:|:-:|:-:|:-:|
| Xem place/story/route công khai | ✓ | ✓ | ✓ | ✓ | ✓ |
| Lưu plan/saved (local), đề xuất địa điểm (R1.5) | ✓ | — | — | — | — |
| Xem đơn | — | chi nhánh mình | chi nhánh mình | mọi chi nhánh | ✓ |
| Xác nhận, đổi giờ, từ chối, đã đến | — | chi nhánh mình | chi nhánh mình | mọi chi nhánh | ✓ |
| Bật/tắt món, sửa giá nhanh | — | ✓ | ✓ | ✓ | — |
| Sửa thực đơn, khung giờ đặt trước, hồ sơ chi nhánh | — | — | chi nhánh mình | ✓ | ✓ |
| Thêm/sửa Nhân viên; cấp lại mật khẩu; khoá/mở khoá | — | — | Nhân viên chi nhánh mình | ✓ | ✓ |
| Xoá Nhân viên | — | — | — | ✓ | ✓ |
| Tạo/sửa Quản lý, Chủ quán | — | — | — | ✓ | ✓ |
| Chi nhánh, thiết bị (xem, đăng xuất), cấu hình, báo cáo toàn quán | — | — | — | ✓ | ✓ |
| Nhật ký thao tác theo người | — | — | chi nhánh mình | ✓ | ✓ |
| Đổi `recovery_email` | — | — | — | ✓ + Admin duyệt | ✓ |
| Đối soát/tiền (giai đoạn 3) | — | — | — | ✓ | ✓ (kế toán) |
| Duyệt hồ sơ, cấp tài khoản quán, duyệt đề xuất | — | — | — | — | ✓ |
| Sao lưu/khôi phục | — | — | — | — | ✓ (Owner/ops) |

> Mức tối thiểu; chi tiết cuối cùng ở `03-data-model.md` và `04-api-contract.md` khi viết. Mọi ô “✓” vẫn phải kiểm ở server.

### 6. Quy ước hợp đồng

| Chủ đề | Quy ước |
|---|---|
| Đường dẫn | Danh từ số nhiều, chữ thường, gạch nối: `/places`, `/stories`, `/routes/plan`; không động từ trong path trừ action rõ nghĩa (`/auth/login`, `/routes/plan`) |
| ID | UUID dạng chuỗi; không lộ số thứ tự nội bộ |
| Thời gian | ISO 8601 có múi giờ; DB/API dùng UTC, UI đổi múi giờ; route request nhận `departAt` có offset |
| Toạ độ | WGS84 (`lat`, `lon`), 6 chữ số thập phân khi trả công khai; validate khoảng hợp lệ |
| Tiền | Số nguyên VND (`bigint`), không dấu thập phân; field `currency: "VND"`; không dùng float |
| Ngôn ngữ | `Accept-Language` hoặc `locale=vi` / `locale=en`; fallback `vi`; không trả chuỗi đã dịch cứng trong code |
| Phân trang | Public: cursor (`limit`, `cursor`, `nextCursor`); Admin: `page`, `pageSize`, `total` |
| Lọc/sắp xếp | Tham số có tên, whitelist; không nhận SQL/sort tự do |
| Idempotency | `Idempotency-Key` cho POST tạo feedback/submission/booking/payment; lưu 24 giờ; cùng key khác payload → `409` |
| Tải file | `multipart/form-data`; giới hạn loại/kích thước; trả metadata + URL, không trả nội dung file trong JSON |
| Log | `requestId` trả trong header `X-Request-Id` và error envelope; log theo `requestId`, đã che PII |

### 7. Dạng lỗi và mã lỗi

Mọi lỗi:

```json
{
  "error": {
    "code": "VALIDATION_FAILED",
    "message": "Toạ độ không hợp lệ.",
    "requestId": "req_01J...",
    "details": { "field": "origin.lat" }
  }
}
```

- `code` ổn định, tiếng Anh, viết hoa gạch dưới; `message` tiếng Việt/Anh theo locale, không lộ stack/SQL/tên bảng.
- `details` chỉ chứa dữ liệu an toàn (field lỗi, danh sách mã), không chứa token/PII.
- `requestId` có trong mọi response lỗi; thành công cũng trả `X-Request-Id`.

| Mã | HTTP | Khi nào |
|---|---|---|
| `VALIDATION_FAILED` | 400 | JSON sai, thiếu field, toạ độ ngoài khoảng, thời gian sai định dạng |
| `UNAUTHENTICATED` | 401 | Thiếu/hết hạn token/cookie, sai thông tin đăng nhập |
| `FORBIDDEN` | 403 | Đúng realm nhưng thiếu quyền/không thuộc quán |
| `NOT_FOUND` | 404 | Không có tài nguyên hoặc không được phép biết là có |
| `CONFLICT` | 409 | Trùng email, trùng trạng thái, `Idempotency-Key` khác payload |
| `PAYLOAD_TOO_LARGE` | 413 | Body/file quá giới hạn |
| `UNSUPPORTED_MEDIA_TYPE` | 415 | Loại file không cho phép |
| `UNSUPPORTED` | 422 | Ngoài vùng phủ hoặc mode/constraint không hỗ trợ; `details.reason` = `OUTSIDE_COVERAGE`, `MODE_NOT_SUPPORTED`, `CONSTRAINT_NOT_SUPPORTED` |
| `NO_ROUTE` | 422 | Trong vùng phủ nhưng không có tuyến thỏa ràng buộc |
| `DATA_UNAVAILABLE` | 422 | Ngoài giờ chạy/không có lịch cho ngày đó |
| `RATE_LIMITED` | 429 | Vượt giới hạn |
| `INTERNAL` | 500 | Lỗi không lường trước; không trả chi tiết nội bộ |
| `DEPENDENCY_UNAVAILABLE` | 503 | DB/OTP/geocoder/media không sẵn sàng |
| `ROUTING_TIMEOUT` | 504 | OTP quá hạn |

- `UNSUPPORTED` là mã chung cho “chưa hỗ trợ”; client đọc `details.reason` để hiển thị đúng. Không dùng `200` + danh sách rỗng để che lỗi.
- Lỗi xác thực luôn cùng một thông báo, không phân biệt “sai email” và “sai mật khẩu”.

### 8. Route response và nhãn dữ liệu

- Route response có `dataStatus` (`OFFICIAL|ASSUMED|BEST_EFFORT|UNSUPPORTED`), `coverage` và `graphVersion`.
- Mỗi route/branch có `legality.status` (`FILTERED|PARTIAL|UNKNOWN`) và `legality.unverifiedRules[]`; không được nói “đúng luật tuyệt đối” khi `PARTIAL/UNKNOWN` (chi tiết ADR-0005).
- Không trả route ngoài coverage; trả `UNSUPPORTED` trước khi gọi engine.
- Không log toạ độ chính xác của người dùng; nếu cần điều tra, dùng `requestId` và mẫu đã làm tròn.

### 9. Audit và nhật ký bảo mật

- Bảng audit append-only (chi tiết ADR-0004): ai (realm/id), làm gì, đối tượng nào, trước/sau (đã che), `requestId`, thời gian, `ip_hash`.
- Ghi audit cho: đăng nhập Admin/Partner (thành công/thất bại), đổi quyền, duyệt/từ chối hồ sơ, publish/unpublish nội dung, sửa giá/cọc, khôi phục sao lưu, xoá tài khoản.
- Không ghi mật khẩu, token, cookie, mã 2FA, số giấy tờ đầy đủ, toạ độ người dùng.
- Log ứng dụng (slog JSON) có `requestId`, route, status, latency; không có body mặc định; body chỉ log khi có cờ debug ở môi trường dev và đã che PII.

### 10. Giới hạn tần suất

| Nhóm | Giới hạn đề xuất | Ghi chú |
|---|---|---|
| Đăng nhập | 5 lần/15 phút/tài khoản; 20 lần/giờ/IP | Trả `429`, không tiết lộ tài khoản |
| Đăng nhập quán | 5 lần/15 phút/mã nhân viên và 5 lần/15 phút/thiết bị-IP | Khoá tự mở, email báo; không khoá vĩnh viễn |
| Quên mật khẩu | 5 lần/giờ/tài khoản; 20 lần/giờ/IP | Mã 10 phút |
| Đăng ký | 5 lần/giờ/IP | Chống spam |
| Đề xuất địa điểm | 3/ngày/account (đã chốt ở scope) | R1.5 |
| Route | 30 lần/phút/IP; 120 lần/giờ/token | Tránh lạm dụng OTP |
| Geocode | 60 lần/phút/IP | Không log query |
| Feedback | 5/ngày/IP | Có captcha/honeypot nếu bị spam |

> Các con số là **đề xuất**, chưa đo tải thật; cấu hình có tên, không hard-code trong handler.

### 11. Quy trình OpenAPI và client

1. Viết/sửa `api/openapi.yaml` trước khi sửa handler.
2. Lint spec (Spectral hoặc tương đương) + kiểm breaking change trong CI.
3. Sinh TypeScript client cho `web/packages/api-client` (`openapi-typescript` hoặc Orval); không sửa file sinh.
4. Mobile dùng client Ktor mỏng viết theo spec; DTO/route được kiểm bằng contract test với mock server sinh từ spec. Không sinh Kotlin client tự động ở R1 nếu công cụ không hỗ trợ KMP tốt; quyết định lại khi dựng khung.
5. Contract test ở Go: response thật phải validate theo schema; có test cho error envelope và các mã lỗi chính.
6. Mock server từ spec cho web/mobile khi backend chưa xong; fixture đánh dấu `DEMO`.
7. Spec có `operationId` ổn định, ví dụ request/response, mô tả quyền và realm cho từng operation.

### 12. Bề mặt API v0 (R1)

Chỉ các endpoint cần cho R1; phần còn lại là seam, chưa mở:

| Nhóm | Endpoint | Realm |
|---|---|---|
| Health | `GET /health/live`, `GET /health/ready` | Public nội bộ |
| Cấu hình | `GET /api/v1/app/config` (feature flag, phiên bản tối thiểu, coverage summary) | Public |
| Discover | `GET /api/v1/categories`, `GET /api/v1/places`, `GET /api/v1/places/{id}`, `GET /api/v1/places/nearby` | Public |
| Story/Guide | `GET /api/v1/stories`, `GET /api/v1/stories/{slug}` | Public |
| Geocode | `GET /api/v1/geocode`, `GET /api/v1/geocode/reverse` | Public (rate limit) |
| Route | `POST /api/v1/routes/plan` | Public (rate limit) |
| Feedback | `POST /api/v1/feedback` | Public/guest |
| Auth client (R1.5) | `/api/v1/auth/register`, `/login`, `/google`, `/logout`, `/verify-email`, `/password/forgot`, `/password/reset` | Public |
| Me (R1.5) | `GET/DELETE /api/v1/me`, `GET /api/v1/me/submissions`, `POST /api/v1/me/submissions`, `POST /api/v1/me/submissions/{id}/withdraw` | Client |
| Admin | `/api/v1/admin/auth/*`, `/api/v1/admin/places`, `/stories`, `/media`, `/feedback`, `/submissions`, `/partners`, `/users`, `/audit` | Admin |
| Quán (R6) | `POST /api/v1/partner/auth/login` (mã nhân viên + mật khẩu), `/logout`, `/switch-staff`, `/password/forgot`, `/password/set`; `/api/v1/partner/shops/{shopId}/branches`, `/staff`, `/devices`, `/bookings`, `/menu`; `POST .../staff/{staffId}/reset-password`; `DELETE .../devices/{deviceId}` | Partner |

- `GET /api/v1/app/config` không trả bí mật; chỉ trả URL tile/style, feature flag công khai, phiên bản tối thiểu, locale mặc định.
- Feedback guest có thể gửi không cần account; nếu có account thì lấy email/tên từ account, không tin client.
- Các endpoint Admin/Partner chỉ được thêm khi phase tương ứng mở; không tạo endpoint rỗng “chuẩn bị trước”.

### 13. Mẫu thiết kế áp dụng

Mỗi mẫu chỉ dùng khi có vấn đề thật ở cột “Vấn đề”; không thêm lớp trừu tượng “để sau này”.

| Vấn đề | Mẫu | Cách áp dụng |
|---|---|---|
| Ba realm, mỗi realm xác minh khác nhau | **Strategy** (`CredentialVerifier` theo realm) | Mỗi realm một bộ xác minh (khách: email/SĐT; quán: mã nhân viên; Admin: email + 2 lớp) cùng một interface; không viết một hàm login có `switch` dài |
| Xử lý request nhiều bước | **Chain of Responsibility** (middleware) | `request-id → realm → phiên → rate limit → authorize`; mỗi bước làm một việc, thứ tự cố định |
| Quyền rải rác | **Policy** (một `Authorizer`) | Mục 5.2; thêm quyền = thêm một dòng policy và một test từ chối |
| Ghi dữ liệu và gửi email/thông báo | **Transactional Outbox** | Thay đổi nghiệp vụ, audit và `outbox_event` ghi **cùng một transaction**; worker gửi email, thông báo, đếm “món đặt nhiều”. Không gọi SMTP trong request, không ghi kép |
| Hai thiết bị sửa một đơn | **Optimistic locking** | Cột `version`; `UPDATE ... WHERE id = :id AND version = :v`; sai → `409` |
| Retry do mạng | **Idempotency-Key** | Đã có ở mục 6 |
| Nhà cung cấp ngoài (Vietmap, OTP, email) | **Port/Adapter (anti-corruption layer)** | Domain chỉ biết interface của mình; schema nhà cung cấp không rò vào domain; thay nhà cung cấp chỉ đổi adapter |
| Trạng thái đơn, nhân viên, phiên thiết bị | **State machine tường minh** | Bảng chuyển hợp lệ trong code và test; chuyển sai trả `CONFLICT` |
| Chống dò mật khẩu, lạm dụng | **Token bucket / sliding window** | Cấu hình có tên (mục 10), không hard-code trong handler |

## Lý do

1. **Một hợp đồng, nhiều client.** OpenAPI + contract test giảm lệch giữa KMP, Admin, Partner và tránh lặp lại việc app phụ thuộc schema OTP.
2. **Tách realm là yêu cầu bảo mật.** Admin có quyền cao; dùng chung cookie/token với client/partner làm tăng blast radius. Cookie riêng + 2FA + step-up giảm rủi ro.
3. **Opaque session thay JWT.** Cần thu hồi ngay khi đổi mật khẩu/xoá tài khoản; JWT stateless khó thu hồi và dễ bị dùng sai. DB lookup mỗi request chấp nhận được ở quy mô pilot; có thể cache sau.
4. **Session cookie cho web, bearer cho mobile.** Web cần chống XSS bằng `HttpOnly`; mobile không có cookie jar ổn định và cần Keychain/Keystore.
5. **RBAC theo đối tượng.** Quyền quán gắn (tài khoản quán, chi nhánh, vai trò) của người đăng nhập; kiểm ở route là không đủ. Một `Authorizer` deny-by-default + audit là mức tối thiểu cho dữ liệu tiền, giấy tờ và dữ liệu nhân viên.
6. **Lỗi có mã ổn định.** Client hiển thị đúng “ngoài vùng phủ” khác “không có tuyến” khác “hết giờ chạy”; không đoán qua HTTP status.
7. **Idempotency cho tiền và nội dung.** Feedback/submission/booking có thể retry do mạng; thiếu idempotency tạo bản ghi trùng.

## Phương án đã cân nhắc

| Phương án | Vì sao không chọn |
|---|---|
| **JWT access token + refresh token** | Thu hồi phức tạp, dễ lộ lâu; chưa cần scale stateless ở pilot. Có thể xem lại khi API nhiều instance và session lookup thành điểm nghẽn |
| **Một cookie/session cho cả Admin, Partner, Client** | Blast radius lớn; lỗi phân quyền một chỗ ảnh hưởng toàn hệ thống; không hợp ADR-0002 (Admin tài khoản riêng) |
| **Auth0/Keycloak/Clerk** | Thêm nhà cung cấp, chi phí, phụ thuộc ngoài; yêu cầu dữ liệu/giấy tờ và RBAC theo quán vẫn phải tự làm. Xem lại nếu cần SSO doanh nghiệp |
| **OIDC/Google là cơ chế đăng nhập duy nhất** | Không phù hợp Admin (chủ dự án muốn tài khoản riêng) và người dùng không có Google; Google chỉ là identity provider phụ |
| **GraphQL** | Linh hoạt cho client nhưng khó chặn truy vấn nặng, khó sinh client KMP, không cần cho phạm vi R1 |
| **gRPC** | Tốt cho nội bộ/hiệu năng, nhưng web/mobile/OpenAPI phức tạp hơn; REST/JSON đủ |
| **Cookie cho mobile** | Không ổn định trên KMP, khó kiểm soát lưu trữ an toàn; bearer + Keychain/Keystore phù hợp hơn |
| **Lưu token trong localStorage** | XSS đọc được; vi phạm quy ước web |

## Hệ quả

### Tích cực

- Client biết chính xác endpoint, mã lỗi, kiểu dữ liệu; giảm công đoán và lỗi tích hợp.
- Admin/Partner có mô hình phiên riêng, 2FA, step-up, audit; phù hợp dữ liệu tiền và giấy tờ.
- RBAC theo (quán, chi nhánh, vai trò) mở đường cho R6 và giai đoạn 3 mà không phải đổi kiến trúc.
- OpenAPI + contract test tạo nền cho client sinh tự động và mock.

### Tiêu cực / chi phí

- Phải xây module Identity, session store, middleware RBAC, CSRF, 2FA, rate limit, audit.
- Session lookup DB mỗi request; cần index và có thể cần cache khi tải tăng.
- OpenAPI + sinh client + contract test là thêm CI và kỷ luật; đội nhỏ phải duy trì.
- Mobile phải xử lý token hết hạn/ bị thu hồi, đăng nhập lại, nhiều thiết bị.
- 2FA Admin làm chậm thao tác khẩn cấp; cần quy trình mã dự phòng và khôi phục.
- Mật khẩu nhân viên quán tối thiểu 6 ký tự/6 số yếu hơn khách/Admin; phải bù bằng khoá 15 phút, rate limit, email báo thiết bị mới, audit theo mã nhân viên.
- Các con số timeout/rate limit là đề xuất, cần đo và chỉnh.

### Rủi ro

- Mã nhân viên có quy luật nên tên đăng nhập đoán được; an toàn dựa vào mật khẩu riêng, khoá và rate limit. Nếu các lớp này bị vô hiệu thì mật khẩu 6 số rất dễ đoán; cần test và giám sát.
- Email khôi phục là điểm tấn công lớn nhất; bắt buộc Admin duyệt khi đổi email (ADR-0002 mục 6).
- Dữ liệu nhân viên (họ tên, SĐT, ca làm) chưa có DPIA hay luật sư rà.
- Google/Apple sign-in có thể trễ R1.5; không được chặn đăng nhập email/password.
- 2FA TOTP có thể mất thiết bị; quy trình recovery phải được viết và thử trước khi mở Admin ra Internet.
- Rate limit theo IP có thể chặn nhầm người dùng sau NAT; cần theo dõi và có ngoại lệ.
- Audit log có thể chứa PII nếu không che; review schema/log trước release.
- Chưa có tải thật để chọn session store/cache; có thể phải thêm Redis sau (ADR mới).

## Việc tiếp theo

- [ ] Chủ dự án chốt các con số đề xuất: thời hạn phiên (khách, quán, Admin), 2FA (TOTP/passkey), thời gian step-up, rate limit.
- [x] ADR-0002 viết lại 18-11-2026: tài khoản quán do Admin cấp, mã nhân viên + mật khẩu riêng, nhiều thiết bị, chi nhánh.
- [ ] Viết `../02-architecture/04-api-contract.md` và `api/openapi.yaml` v0 (R1); thêm nhóm `/partner/auth/*` khi mở R6.
- [ ] Viết `../02-architecture/03-data-model.md` khớp ADR-0004 (đã viết lại: `shop_account`, `branch`, `staff`, `device_session`).
- [ ] Viết `../02-architecture/05-security-privacy.md`: threat model, CSRF/XSS, 2FA, retention, dữ liệu nhân viên.
- [ ] Chọn thư viện Argon2id/TOTP và viết test vector trước khi code auth.
- [ ] Test: khoá 15 phút, thiết bị mới gửi email, thu hồi phiên khi đổi mật khẩu/xoá nhân viên, Quản lý không cấp được cho Quản lý/Chủ quán, mã `NV` không vào được thao tác Chủ quán.
- [ ] Quyết định passkey/WebAuthn ở R1 hay R6.

## Tham khảo

- `ADR-0001-web-stack.md` — cookie httpOnly + CSRF, không token trong localStorage.
- `ADR-0002-account-roles.md` — ba loại đăng nhập; tài khoản quán do Admin cấp; mã nhân viên + mật khẩu riêng; Admin riêng + 2FA.
- `../01-product/08-accounts-roles-partners.md` — luồng đăng nhập, quên mật khẩu, vai trò, hồ sơ đối tác.
- `../04-development/02-conventions.md` — dạng lỗi, quy ước backend/web/mobile.
- Repo cũ `../GoWander`: `ADR-022` (đăng nhập/ngôn ngữ), `ADR-027` (API Go, mã lỗi), `ADR-039` (Admin + Argon2id), `ADR-040` (token app 60 ngày, không JWT). Chỉ tham khảo; quyết định v2 là quyết định mới.
