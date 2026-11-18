# Architecture Decision Records (ADR)

- **Trạng thái:** Active
- **Ngày:** 18-11-2026
- **Chủ sở hữu:** chủ dự án

ADR ghi lại **quyết định kiến trúc** và lý do. Một ADR đã Accepted thì code/docs phải theo; muốn đổi thì tạo ADR mới thay thế, không sửa lịch sử.

## Trạng thái

| Trạng thái | Nghĩa |
|---|---|
| **Proposed** | Đã đủ để review, chờ chốt |
| **Accepted** | Đã chốt; có hiệu lực |
| **Superseded** | Bị ADR khác thay thế; giữ để lưu vết |
| **Rejected** | Đã cân nhắc và không chọn |
| **Deprecated** | Không còn khuyến nghị nhưng chưa thay thế |

## Danh sách

| ADR | Tiêu đề | Trạng thái | Ngày |
|---|---|---|---|
| [ADR-0001](./ADR-0001-web-stack.md) | Web stack: React + TypeScript, Next.js cho Client Web, Vite cho Admin/Partner | Accepted | 18-11-2026 |
| [ADR-0002](./ADR-0002-account-roles.md) | Tài khoản khách, tài khoản quán do Admin cấp, đăng nhập bằng mã nhân viên + mật khẩu riêng | Accepted (viết lại 18-11-2026) | 18-11-2026 |
| [ADR-0003](./ADR-0003-api-contract-auth-rbac.md) | Hợp đồng API, đăng nhập, phiên và RBAC (đã viết lại theo ADR-0002 mới) | Proposed | 18-11-2026 |
| [ADR-0004](./ADR-0004-data-model-postgis.md) | Mô hình dữ liệu PostgreSQL/PostGIS (đã viết lại theo ADR-0002 mới) | Proposed | 18-11-2026 |
| [ADR-0005](./ADR-0005-routing-algorithm.md) | Thuật toán tìm đường: lọc đúng luật, đa tiêu chí Pareto, bốn nhánh; OTP/RAPTOR/A*, transit/OSM giả định | Proposed | 18-11-2026 |

> **ADR-0002 đã viết lại (18-11-2026) và Accepted:** ba loại đăng nhập (cá nhân, quán do Admin cấp, Admin); đăng nhập quán bằng mã nhân viên + mật khẩu riêng (bỏ mật khẩu quán); nhiều thiết bị và chi nhánh; Chủ quán tự CRUD nhân viên. **ADR-0003/0004 còn mô tả mô hình cũ (`shop_membership`) và phải viết lại theo ADR-0002** (xem ghi chú đầu mỗi file); chúng vẫn **Proposed**.

### Quan hệ giữa các ADR

| ADR | Phụ thuộc / liên quan | Ghi chú |
|---|---|---|
| ADR-0003 | ADR-0001, ADR-0002 | Hợp đồng API, realm xác thực, RBAC, OpenAPI |
| ADR-0004 | ADR-0003 | Bảng account/session/RBAC, mô hình dữ liệu và PostGIS |
| ADR-0005 | ADR-0003, ADR-0004 | Route API, coverage/graph metadata, thuật toán và nhãn dữ liệu |

> Các ADR **0003/0004/0005** đang **Proposed**: chưa phải luật. Số liệu/thời hạn/trọng số trong đó là đề xuất, phần lớn gắn `ASSUMED`; chủ dự án chốt mới có hiệu lực.

## Quy ước

- Tên file: `ADR-XXXX-ten-ngan.md`, số tăng dần, không tái sử dụng.
- Mỗi ADR có: bối cảnh, quyết định, lý do, phương án đã cân nhắc, hệ quả, việc tiếp theo.
- ADR không thay thế change record; nó giải thích **vì sao**, change record ghi **đã làm gì**.
- Khi dùng lại quyết định từ repo cũ, phải viết ADR mới và ghi nguồn tham khảo.

## Liên quan

- `../02-architecture/01-system-context.md`
- `../02-architecture/02-containers.md`
- `../04-development/02-conventions.md`
