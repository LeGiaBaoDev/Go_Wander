# ADR-0002 — Một kiểu tài khoản, quyền theo vai trò ở từng quán; Admin cấp tài khoản quán

- **Trạng thái:** Proposed
- **Ngày:** 18-11-2026
- **Người quyết định:** chủ dự án (chờ chốt)
- **Liên quan:** `../01-product/08-accounts-roles-partners.md`, `../00-foundation/02-scope.md`, `ADR-0001-web-stack.md`

## Bối cảnh

- Partner (R6–R7) cần quản lý đơn đặt bàn, thực đơn, cọc. Chủ quán và nhân viên không phải lúc nào cũng ngồi web.
- Scope cũ để **Business Mobile ở R8** như một app riêng.
- Chủ dự án muốn: vẫn một app; quyền do tài khoản quyết định; quán không tự đăng ký tự do mà phải nộp giấy tờ để xác minh; Admin cấp tài khoản.

## Quyết định

1. **Một kiểu tài khoản người dùng** cho app và Partner web. Quyền là **vai trò gắn với (tài khoản, quán)**: Chủ quán / Quản lý / Nhân viên. Một tài khoản có thể có nhiều vai trò ở nhiều quán.
2. **App khách có “chế độ quán”** cho việc hằng ngày (đơn, trạng thái, bật/tắt món), chỉ hiện khi tài khoản có vai trò. Cấu hình (khung giờ, cọc, thực đơn đầy đủ, nhân viên, đối soát) ở **Partner web**. Không làm app Business riêng trong giai đoạn này.
3. **Không đăng ký Partner tự do.** Quán nộp **hồ sơ hợp tác kèm giấy tờ**; Admin xác minh; **Admin cấp tài khoản Chủ quán** qua lời mời dùng 1 lần.
4. **Chủ quán tự mời** Quản lý/Nhân viên *(đề xuất; chủ dự án chưa chốt giữa “chủ quán mời” và “chỉ Admin cấp”)*.
5. **Admin GoWander là tài khoản riêng**, chỉ web, do Owner cấp, **bắt buộc 2 lớp**; vai trò: Owner, Vận hành, Biên tập, Duyệt đề xuất, Kế toán, Hỗ trợ Partner.
6. **Kiểm quyền ở server** (RBAC theo quán); client chỉ ẩn/hiện. Mọi thao tác ghi nhật ký.

## Phương án đã cân nhắc

| Phương án | Vì sao không chọn |
|---|---|
| App Business riêng (R8 cũ) | Hai app, hai lần đăng nhập; nhân viên thường đã có app khách |
| Partner chỉ dùng web | Nhân viên giữa ca không ngồi máy tính |
| Partner tự đăng ký, duyệt sau | Rủi ro quán giả, sai chủ quán, vướng nghĩa vụ xác minh |
| Admin cấp cả tài khoản nhân viên | Admin quá tải; quán mất chủ động (vẫn là lựa chọn B nếu chủ dự án muốn) |

## Hệ quả

- Data model cần: `account`, `shop`, `shop_membership(account, shop, role)`, `shop_invite`, `partner_application`, `application_document`, `admin_account`, `admin_role`, nhật ký thao tác.
- API contract (ADR-0003, dự kiến) phải mô tả scope theo quán và lỗi thiếu quyền.
- Scope: bỏ “Business Mobile” khỏi R8, thay bằng chế độ quán trong app (R7).
- Pháp lý: lưu giấy tờ là dữ liệu cá nhân; cần chính sách lưu trữ/xoá và DPIA (xem `06-research`).

## Việc tiếp theo

- Chủ dự án chốt điểm 4.
- Viết `02-architecture/03-data-model.md` với các entity trên.
- Luật sư rà danh sách giấy tờ và nghĩa vụ TMĐT trước khi bật thu cọc.
