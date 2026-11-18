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
| [ADR-0002](./ADR-0002-account-roles.md) | Một kiểu tài khoản, quyền theo vai trò ở từng quán; Admin cấp tài khoản quán | Proposed | 18-11-2026 |
| ADR-0003 | API contract & auth (dự kiến) | Planned | — |
| ADR-0004 | Data model & PostGIS (dự kiến) | Planned | — |
| ADR-0005 | Thuật toán tìm đường (lọc đúng luật → đa tiêu chí → 4 nhánh), transit/OTP & dữ liệu giả định (dự kiến; bản giải thích: canvas `RouteLogic`) | Planned | — |

## Quy ước

- Tên file: `ADR-XXXX-ten-ngan.md`, số tăng dần, không tái sử dụng.
- Mỗi ADR có: bối cảnh, quyết định, lý do, phương án đã cân nhắc, hệ quả, việc tiếp theo.
- ADR không thay thế change record; nó giải thích **vì sao**, change record ghi **đã làm gì**.
- Khi dùng lại quyết định từ repo cũ, phải viết ADR mới và ghi nguồn tham khảo.

## Liên quan

- `../02-architecture/01-system-context.md`
- `../02-architecture/02-containers.md`
- `../04-development/02-conventions.md`
