# ADR-0003 — Hợp đồng API, đăng nhập, phiên và RBAC

- **Trạng thái:** Proposed — chờ chủ dự án chốt
- **Ngày:** 18-11-2026
- **Người quyết định:** chủ dự án (chờ chốt)
- **Liên quan:** `ADR-0001-web-stack.md`, `ADR-0002-account-roles.md`, `../02-architecture/01-system-context.md`, `../02-architecture/02-containers.md`, `../01-product/08-accounts-roles-partners.md`, `../04-development/02-conventions.md`

> ADR-0002 đang **Proposed**. ADR-0003 giả định hướng của ADR-0002 (một kiểu tài khoản, quyền theo quán, Admin là tài khoản riêng) sẽ được chấp nhận. Nếu ADR-0002 đổi, phần quyền trong ADR này phải sửa theo.

## Bối cảnh

GoWander có ba loại client với nhu cầu xác thực khác nhau:

| Client | R1 | Cách dùng | Ràng buộc |
|---|---|---|---|
| Client Mobile (KMP) | Guest-first; tài khoản từ R1.5 | Bearer token | Token trong Keychain/Keystore; không lưu plan/saved trên server mặc định |
| Admin Web (React/Vite) | Có | Session cookie | RBAC, bắt buộc 2 lớp, chống CSRF, không token trong `localStorage` |
| Partner Web (React/Vite) | R6 | Session cookie | Quyền theo từng quán; nhân viên dùng chung tài khoản với app |
| Showcase Web (Next.js, static) | Có | Public | Không form thu PII ở R1; không cần đăng nhập |

Các quyết định đã có:

- Một backend Go (modular monolith), client không gọi DB/OTP/nhà cung cấp ngoài (`01-system-context.md`).
- OpenAPI là nguồn sự thật của hợp đồng API (`03-principles.md` nguyên tắc 6; `02-conventions.md`).
- Lỗi có một dạng thống nhất, có `requestId` (`02-conventions.md`).
- Ngoài vùng hỗ trợ phải trả `UNSUPPORTED`, không bịa tuyến (`02-scope.md`).
- Guest-first, local-first, privacy by default (`03-principles.md` nguyên tắc 7).
- Admin là tài khoản riêng, bắt buộc hai lớp; quyền quán theo vai trò (`ADR-0002`, `08-accounts-roles-partners.md`).

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
| `client` | Người dùng app (R1.5) | Bearer token | Token trong Keychain/Keystore | Không bắt buộc |
| `admin` | Admin/Editor GoWander | Session cookie | `__Host-gw_admin` | **Bắt buộc** |
| `partner` | Chủ quán/Quản lý/Nhân viên (R6) | Session cookie | `__Host-gw_partner` | Không bắt buộc mặc định |

Quy tắc:

- **Không dùng chung cookie/token giữa các realm.** Cookie Admin không mở được endpoint Partner và ngược lại; token client không gọi được `/admin/*` hay `/partner/*`.
- Admin là **tài khoản riêng**, không phải account client được gắn role. Partner dùng **cùng kiểu account** với client, quyền đến từ `shop_membership` (ADR-0002).
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

#### 3.3 Partner Web

- Session cookie `__Host-gw_partner`: cùng thuộc tính bảo vệ như Admin.
- Idle timeout **8 giờ**, absolute timeout **7 ngày**. *(Đề xuất; chờ chốt. Nhân viên giữa ca không muốn đăng nhập lại liên tục.)*
- Không bắt buộc 2 lớp mặc định; Chủ quán có thể bật, và GoWander có thể bắt buộc khi có thu cọc hoặc đối soát tiền.
- Mọi request đến `/partner/shops/{shopId}/...` phải kiểm `shop_membership` của account đang đăng nhập; không tin `shopId` do client gửi.

#### 3.4 Quy tắc chung cho cookie

- Chỉ phát cookie qua HTTPS; `Secure` bắt buộc ở production.
- Ở dev trên `localhost`, trình duyệt có thể không nhận cookie `__Host-`/`Secure`; dùng tên cookie dev riêng (ví dụ `gw_admin_dev`) và bật HTTPS cục bộ khi test luồng đăng nhập thật. Không nới `Secure` ở production.
- `SameSite=Lax`; nếu sau này cần nhúng cross-site thì mở ADR riêng, không nới `SameSite=None` mặc định.
- Không đặt session ID trong URL; không log cookie, `Set-Cookie`, `Authorization`.
- CORS chỉ mở cho origin cấu hình (Showcase/Client Web); Admin/Partner chạy cùng origin qua reverse proxy nên không cần CORS rộng.

### 4. Mật khẩu, xác minh và khôi phục

- Băm mật khẩu bằng **Argon2id** (`golang.org/x/crypto/argon2`), tham số khởi điểm OWASP: 19 MiB, 2 vòng, 1 luồng, muối 16 byte; tham số lưu kèm hash để nâng cấp dần. *(Chờ chốt; không dùng bcrypt vì giới hạn 72 byte.)*
- Mật khẩu client: 10–128 ký tự, có ít nhất một chữ và một số; không trùng 3 mật khẩu gần nhất (lưu hash lịch sử). Admin: tối thiểu 12 ký tự. *(Đề xuất; chờ chốt.)*
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
| Partner | `owner`, `manager`, `staff` | Theo `shop_id` trong `shop_membership` |
| Admin | `owner`, `ops`, `editor`, `moderator`, `accountant`, `partner_support` | Toàn hệ thống hoặc theo quyền được cấp |

- Một account có thể có nhiều vai trò ở nhiều quán; mỗi quán **đúng một** `owner` đang hoạt động.
- Admin không tự động có quyền Partner; Partner không tự động có quyền Admin.
- Danh sách quyền Admin lấy từ `08-accounts-roles-partners.md`; tên quyền kỹ thuật dạng `resource:action` (ví dụ `place:write`, `story:publish`, `submission:review`, `shop:payout:read`, `admin:user:manage`, `backup:restore`).

#### 5.2 Cách kiểm

- **Deny by default.** Route middleware chỉ kiểm realm; service kiểm quyền cụ thể và phạm vi đối tượng.
- Kiểm quyền theo đối tượng, không chỉ theo route: ví dụ `PATCH /partner/shops/{shopId}/menu/{itemId}` phải kiểm membership của `shopId`, item thuộc `shopId`, và role được sửa menu.
- Không tin `role`, `shopId`, `accountId` do client gửi; lấy từ session/token và DB.
- Mọi thao tác ghi của Admin/Partner ghi audit trong cùng transaction (xem mục 9).
- Có test cho **đường từ chối**: thiếu token, sai realm, thiếu role, sai `shopId`, đối tượng không tồn tại, role bị khoá.
- UI chỉ ẩn/hiện; ẩn nút không phải bảo vệ.

#### 5.3 Ma trận quyền tối thiểu

| Hành động | Client | Staff | Manager | Owner | Admin |
|---|:-:|:-:|:-:|:-:|:-:|
| Xem place/story/route công khai | ✓ | ✓ | ✓ | ✓ | ✓ |
| Lưu plan/saved (local) | ✓ | ✓ | ✓ | ✓ | — |
| Đề xuất địa điểm (R1.5) | ✓ | ✓ | ✓ | ✓ | — |
| Xem đơn hôm nay của quán | — | ✓ | ✓ | ✓ | ✓ |
| Bật/tắt món, sửa giá nhanh | — | ✓ | ✓ | ✓ | — |
| Sửa thực đơn/khung giờ/cọc | — | — | ✓ | ✓ | — |
| Đối soát/tiền | — | — | — | ✓ | ✓ (kế toán) |
| Mời/gỡ nhân viên | — | — | — | ✓ | ✓ |
| Duyệt đề xuất/đối tác | — | — | — | — | ✓ |
| Sao lưu/khôi phục | — | — | — | — | ✓ (Owner/ops) |

> Bảng này là mức tối thiểu; chi tiết cuối cùng ở `03-data-model.md` và `04-api-contract.md` khi viết. Các ô “✓” vẫn phải kiểm ở server.

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
| Partner (R6) | `/api/v1/partner/auth/*`, `/api/v1/partner/shops/{shopId}/...` | Partner |

- `GET /api/v1/app/config` không trả bí mật; chỉ trả URL tile/style, feature flag công khai, phiên bản tối thiểu, locale mặc định.
- Feedback guest có thể gửi không cần account; nếu có account thì lấy email/tên từ account, không tin client.
- Các endpoint Admin/Partner chỉ được thêm khi phase tương ứng mở; không tạo endpoint rỗng “chuẩn bị trước”.

## Lý do

1. **Một hợp đồng, nhiều client.** OpenAPI + contract test giảm lệch giữa KMP, Admin, Partner và tránh lặp lại việc app phụ thuộc schema OTP.
2. **Tách realm là yêu cầu bảo mật.** Admin có quyền cao; dùng chung cookie/token với client/partner làm tăng blast radius. Cookie riêng + 2FA + step-up giảm rủi ro.
3. **Opaque session thay JWT.** Cần thu hồi ngay khi đổi mật khẩu/xoá tài khoản; JWT stateless khó thu hồi và dễ bị dùng sai. DB lookup mỗi request chấp nhận được ở quy mô pilot; có thể cache sau.
4. **Session cookie cho web, bearer cho mobile.** Web cần chống XSS bằng `HttpOnly`; mobile không có cookie jar ổn định và cần Keychain/Keystore.
5. **RBAC theo đối tượng.** Quyền quán là quan hệ (account, shop, role); kiểm ở route là không đủ. Deny by default + audit là mức tối thiểu cho dữ liệu tiền và giấy tờ.
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
- RBAC theo quán mở đường cho R6/R7 mà không phải đổi kiến trúc.
- OpenAPI + contract test tạo nền cho client sinh tự động và mock.

### Tiêu cực / chi phí

- Phải xây module Identity, session store, middleware RBAC, CSRF, 2FA, rate limit, audit.
- Session lookup DB mỗi request; cần index và có thể cần cache khi tải tăng.
- OpenAPI + sinh client + contract test là thêm CI và kỷ luật; đội nhỏ phải duy trì.
- Mobile phải xử lý token hết hạn/ bị thu hồi, đăng nhập lại, nhiều thiết bị.
- 2FA Admin làm chậm thao tác khẩn cấp; cần quy trình mã dự phòng và khôi phục.
- Các con số timeout/rate limit là đề xuất, cần đo và chỉnh.

### Rủi ro

- ADR-0002 chưa Accepted; nếu mô hình quyền quán đổi (ví dụ chỉ Admin cấp nhân viên), ma trận quyền và bảng `shop_membership` phải sửa.
- Google/Apple sign-in có thể trễ R1.5; không được chặn đăng nhập email/password.
- 2FA TOTP có thể mất thiết bị; quy trình recovery phải được viết và thử trước khi mở Admin ra Internet.
- Rate limit theo IP có thể chặn nhầm người dùng sau NAT; cần theo dõi và có ngoại lệ.
- Audit log có thể chứa PII nếu không che; review schema/log trước release.
- Chưa có tải thật để chọn session store/cache; có thể phải thêm Redis sau (ADR mới).

## Việc tiếp theo

- [ ] Chủ dự án chốt các con số đề xuất: thời hạn phiên, 2FA (TOTP/passkey), thời gian step-up, rate limit, độ dài mật khẩu.
- [ ] Chủ dự án chốt ADR-0002 (đặc biệt điểm 4: chủ quán mời nhân viên hay chỉ Admin cấp).
- [ ] Viết `../02-architecture/04-api-contract.md` với đặc tả chi tiết và kế hoạch OpenAPI.
- [ ] Viết `api/openapi.yaml` v0 cho các endpoint R1; lint + mock.
- [ ] Viết `../02-architecture/03-data-model.md` cho bảng session/account/RBAC (khớp ADR-0004).
- [ ] Viết `../02-architecture/05-security-privacy.md` cho threat model, CSRF/XSS, 2FA, retention.
- [ ] Chọn thư viện Argon2id/TOTP và viết test vector trước khi code auth.
- [ ] Quyết định có dùng passkey/WebAuthn ở R1 hay R6.

## Tham khảo

- `ADR-0001-web-stack.md` — cookie httpOnly + CSRF, không token trong localStorage.
- `ADR-0002-account-roles.md` — một kiểu account, quyền theo quán, Admin riêng + 2FA.
- `../01-product/08-accounts-roles-partners.md` — luồng đăng nhập, quên mật khẩu, vai trò, hồ sơ đối tác.
- `../04-development/02-conventions.md` — dạng lỗi, quy ước backend/web/mobile.
- Repo cũ `../GoWander`: `ADR-022` (đăng nhập/ngôn ngữ), `ADR-027` (API Go, mã lỗi), `ADR-039` (Admin + Argon2id), `ADR-040` (token app 60 ngày, không JWT). Chỉ tham khảo; quyết định v2 là quyết định mới.
