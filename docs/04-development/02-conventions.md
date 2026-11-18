# 02 — Quy ước code

- **Trạng thái:** Proposed
- **Ngày:** 18-11-2026
- **Chủ sở hữu:** chủ dự án
- **Liên quan:** `01-toolchain.md`, `03-testing.md`, `05-git-workflow.md`, `../00-foundation/03-principles.md`, `../03-decisions/ADR-0002-account-roles.md`

## TL;DR

Mười điều cần nhớ khi viết code GoWander:

1. **Docs trước, code sau.** Tính năng/kiến trúc mới cần docs hoặc ADR trước (nguyên tắc 1).
2. **Business logic chỉ ở Go API.** Mobile và web không nhân bản logic, không gọi thẳng DB hay OTP.
3. **OpenAPI là hợp đồng.** Đổi API thì đổi spec trước, sinh lại client, không sửa tay file sinh.
4. **Mọi quyền kiểm tra ở server** theo vai trò (ADR-0002). Ẩn nút trên UI không phải là bảo vệ.
5. **Không bịa dữ liệu.** Dữ liệu giả định mang nhãn `ASSUMED`/`DEMO` ở cả DB, API và UI.
6. **Đường đi luôn đúng luật.** Cạnh sai luật bị loại trước khi tính; không bao giờ “phạt” rồi vẫn trả về.
7. **Lỗi có dạng chuẩn,** không để client tự đoán.
8. **Không log dữ liệu cá nhân,** token, vị trí chính xác.
9. **Một thay đổi hợp lý một commit;** một task một nhánh.
10. **Test và bằng chứng là một phần của “xong”** (`03-testing.md`).

## Cấu trúc repo

Chốt khi dựng khung. Hướng dự kiến (từ ADR-0001 và `02-containers.md`):

```text
backend/        Go API (modular monolith)
mobile/         Kotlin Multiplatform + Compose
web/            apps/{client-web,admin-web,partner-web}, packages/{ui,api-client,config}
api/            openapi.yaml (nguồn sự thật hợp đồng)
infra/          docker compose, proxy, OTP graph config
data/           seed DEMO + script nạp (không commit dữ liệu thô)
docs/           tài liệu (nguồn sự thật)
```

Quy tắc chung:

- Mỗi module domain một thư mục, giao tiếp qua interface/hàm công khai; không “với tay” vào bảng của module khác.
- Tên thư mục, file, biến, hàm, endpoint, bảng, mã lỗi: **tiếng Anh** (AGENTS.md mục 4). UI copy và tài liệu: tiếng Việt, UI có thêm EN.
- Không để file nhảy cóc: một file làm một việc; quá khoảng 400 dòng thì tách, trừ file sinh tự động.

## Backend (Go)

| Chủ đề | Quy ước |
|---|---|
| Module | Theo `02-containers.md`: Identity, Mobility, Discover, Journey, Media, Admin, Weather…; module chưa tới phase thì chưa tạo thư mục rỗng |
| Lớp | handler (HTTP, validate, map lỗi) → service (nghiệp vụ) → repository (SQL). Handler không chứa nghiệp vụ; repository không biết HTTP |
| Context | Mọi hàm I/O nhận `context.Context`; có timeout cho DB, OTP, dịch vụ ngoài |
| Lỗi | Trả `error` có ngữ cảnh (`fmt.Errorf("...: %w", err)`); không `panic` cho luồng thường; không nuốt lỗi |
| Cấu hình | Từ biến môi trường, đọc một nơi; không đọc env rải rác |
| Phụ thuộc | Qua interface nhỏ ở nơi dùng; dễ giả lập khi test |
| Migration | Một file một thay đổi, có thứ tự, chạy lên/xuống được; không sửa migration đã áp dụng |
| SQL | Tham số hoá; không nối chuỗi; index cho truy vấn PostGIS thường dùng |
| Giờ | UTC trong DB và API (ISO 8601); đổi múi giờ ở UI |

### Dạng lỗi API

Mọi lỗi trả cùng một dạng (chi tiết trong `04-api-contract.md` khi viết):

```json
{ "error": { "code": "ROUTE_OUTSIDE_COVERAGE", "message": "…", "requestId": "…" } }
```

- `code` ổn định, tiếng Anh, viết hoa gạch dưới; `message` cho người đọc, không lộ chi tiết nội bộ.
- Có `requestId` để truy vết; log theo `requestId`.
- Ngoài vùng hỗ trợ trả mã `UNSUPPORTED` (xem `02-scope.md`), không trả tuyến “đoán”.

## Thuật toán tìm đường (Mobility)

Chi tiết sẽ nằm ở ADR-0005; bản giải thích hiện có ở canvas (`RouteLogic`). Quy ước code:

- **Thứ tự cố định:** lọc cạnh sai luật → tìm đa tiêu chí (thời gian, quãng đường, tiền) → chọn đại diện → giải thích. Không đảo thứ tự.
- Lọc “đúng luật” là **ràng buộc cứng**; thiếu dữ liệu thì nhãn “chưa xác minh”, không khẳng định đúng luật.
- Trọng số và ngưỡng nằm ở **cấu hình có tên**, kèm ghi chú nguồn; mọi con số giả định ghi `ASSUMED`.
- Mỗi nhánh trả về kèm `dataStatus` và lý do chọn; kết quả tính toán phải **tái lập được** với cùng đầu vào (test bằng dữ liệu cố định).
- Quảng cáo/khuyến mãi không được đổi thứ hạng hay lựa chọn tuyến (nguyên tắc 3).

## Web (React + TypeScript)

- TypeScript `strict`; không dùng `any` khi chưa có lý do (ghi comment một dòng).
- Gọi API qua `packages/api-client` (sinh từ OpenAPI) và TanStack Query; không `fetch` tay rải rác.
- Form: React Hook Form + Zod; schema Zod khớp spec API.
- Component dùng chung ở `packages/ui`; không sao chép component giữa các app.
- Màu, chữ, khoảng cách lấy từ token (palette v4, `../01-product/05-brand-ui-direction.md`); không hard-code mã màu trong component.
- **Màn nhỏ:** không lướt ngang. Bảng thành card, nhiều hơn 4 card thì phân trang; chip dài thành dropdown; nút sửa/xoá dạng icon có `aria-label` (`../01-product/07-design-canvas.md`).
- Admin/Partner: cookie httpOnly + CSRF; **không** lưu token trong `localStorage`.

## Mobile (Kotlin Multiplatform + Compose)

- Plan và địa điểm đã lưu **lưu trên máy** (local-first, nguyên tắc 7); không gửi server mặc định ở R1.
- Mobile chỉ nói chuyện với Go API; không gọi OTP, DB, nhà cung cấp geocoding trực tiếp.
- Token trong Keychain/Keystore; không ghi vào log hay file thường.
- Vị trí chỉ xin khi cần; dùng cho gợi ý gần đây thì không lưu lại.
- Chuỗi giao diện đưa vào tài nguyên song ngữ VI/EN, không hard-code trong composable.

## Dữ liệu và nhãn tin cậy

| Nhãn | Nghĩa | Ở đâu bắt buộc hiện |
|---|---|---|
| `DEMO` | Số liệu mẫu cho thiết kế/seed | UI, seed, tài liệu |
| `ASSUMED` | Giả định (lịch, giờ, giá, trọng số) | UI (giờ ước tính), cấu hình, docs |
| `BEST_EFFORT` | Có hình học nhưng thiếu timetable | UI |
| `UNSUPPORTED` | Ngoài vùng hỗ trợ | API + UI |

- Mọi dữ liệu thật ghi **nguồn, giấy phép, ngày** (trong DB hoặc docs).
- Ảnh có ghi công đúng giấy phép (CC BY, CC BY-SA…); không dùng ảnh không rõ giấy phép.

## Tài khoản và quyền

Theo ADR-0002: ba loại đăng nhập trong một app — cá nhân (khách), quán (Admin cấp; đăng nhập bằng mã nhân viên + mật khẩu riêng), Admin (tài khoản riêng, bắt buộc hai lớp).

- Mọi endpoint ghi rõ vai trò được gọi; kiểm tra ở server, có test cho cả trường hợp bị từ chối.
- Vai trò quán: Chủ quán / Quản lý / Nhân viên. Việc tiền (đối soát) chỉ Chủ quán.
- Mã nhân viên chỉ định danh, **không dùng một mình để đăng nhập**; mật khẩu thành viên quán tối thiểu 6 ký tự/6 số (ADR-0002); cấp lại mật khẩu tạm 24 giờ, bắt đổi, audit.
- Hành động nhạy cảm (khôi phục sao lưu, đổi quyền, duyệt hồ sơ đối tác) ghi nhật ký kiểm toán: ai, khi nào, làm gì.

## Bảo mật và riêng tư (tối thiểu)

- Kiểm tra đầu vào ở mọi biên tin cậy (API, upload, form); giới hạn kích thước và loại file.
- Giấy tờ của đối tác là dữ liệu cá nhân: lưu có kiểm soát truy cập, có chính sách xoá; không đưa vào log.
- Căn cứ pháp lý dữ liệu cá nhân: Luật 91/2025/QH15 và Nghị định 356/2025 (hiệu lực 01-01-2026) — **chưa có luật sư rà**; đừng khẳng định “đã tuân thủ”.
- Không commit bí mật; phát hiện thì dừng và báo chủ dự án.

## Comment và tài liệu trong code

- Comment giải thích **vì sao**, không nhắc lại **cái gì**.
- Một lối tắt có giới hạn: `// shortcut: <giới hạn>, <khi nào nâng cấp>`.
- Hàm công khai có mô tả ngắn khi hành vi không hiển nhiên.
- Quyết định lớn → ADR; thay đổi → mục trong `../07-changes/`.
