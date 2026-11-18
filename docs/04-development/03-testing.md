# 03 — Kiểm thử và bằng chứng

- **Trạng thái:** Proposed
- **Ngày:** 18-11-2026
- **Chủ sở hữu:** chủ dự án
- **Liên quan:** `01-toolchain.md`, `02-conventions.md`, `../00-foundation/03-principles.md` (nguyên tắc 2 và 10)

## TL;DR

“Xong” nghĩa là **có bằng chứng**: lệnh đã chạy, kết quả, và những gì **chưa** kiểm chứng. Không ghi “đã test” khi chưa chạy (nguyên tắc 2).

## Tầng kiểm thử

| Tầng | Công cụ (đề xuất) | Kiểm gì | Khi nào bắt buộc |
|---|---|---|---|
| Đơn vị | Go `testing`; Kotlin test; Vitest | Hàm/logic thuần, quy tắc nghiệp vụ | Mọi logic mới có nhánh, vòng lặp, tiền, bảo mật |
| Tích hợp | Go + DB thật (container); OTP graph nhỏ | SQL/PostGIS, migration lên/xuống, adapter OTP | Mọi thay đổi SQL hoặc migration |
| Hợp đồng | OpenAPI lint + test so khớp response với spec | Không lệch giữa spec, server, client | Mọi thay đổi API |
| Giao diện web | Testing Library; Playwright | Luồng chính, phân quyền, trạng thái lỗi/rỗng | Luồng Admin/Partner quan trọng |
| Mobile | Test đơn vị + thử trên thiết bị thật | Luồng chính; offline; quyền vị trí | Trước mỗi bản thử; ghi rõ máy/OS đã thử |
| Trợ năng | Axe (web) + kiểm tay (VoiceOver/TalkBack) | Tương phản, nhãn, thứ tự focus | Màn mới |
| Hiệu năng | Đo có số | Phản hồi tìm tuyến, thời gian tải | Trước pilot |

## Quy tắc

1. **Test đi cùng logic mới** (nhánh, vòng lặp, parser, tiền, bảo mật). Thay đổi tầm thường không cần test.
2. **Test cả đường từ chối:** thiếu quyền, đầu vào sai, ngoài vùng hỗ trợ, dữ liệu rỗng.
3. **Test tìm đường dùng dữ liệu cố định** nhỏ, có đáp án đã biết: đường một chiều bị loại, nhánh “nhanh/gần/rẻ” khác nhau đúng chỗ, ngoài TP.HCM trả `UNSUPPORTED`.
4. **Không phụ thuộc mạng ngoài** trong test tự động; giả lập nhà cung cấp (thời tiết, geocoding).
5. **Không để test phụ thuộc thứ tự hay giờ máy;** cố định thời gian trong test.
6. **Test lỗi thì sửa gốc,** không xoá test hay nới điều kiện để qua.
7. Dữ liệu test là dữ liệu tổng hợp; **không dùng dữ liệu thật hay cá nhân.**

## Bằng chứng khi báo “xong”

Báo cáo cuối task theo khuôn ở AGENTS.md mục 6, phần **Kiểm chứng** ghi:

- Lệnh đã chạy và kết quả (rút gọn: số test, thời gian, lỗi nếu có).
- Môi trường (máy, OS, phiên bản công cụ) khi liên quan.
- **Chưa kiểm chứng:** liệt kê rõ; ví dụ “chưa thử trên iOS”, “chưa chạy với dữ liệu thật”.

Bằng chứng nặng (ảnh, log dài) để ở `../evidence/`, link từ mục trong `../07-changes/`.

## Cổng chất lượng trước khi merge

- Lệnh kiểm tra ở `01-toolchain.md` đều xanh.
- Có test cho logic mới; test chạy ổn định (không chập chờn).
- Spec OpenAPI và client sinh ra khớp nhau.
- Migration lên và xuống được trên DB trống.
- Không có bí mật, `.env`, dữ liệu thô trong diff.
- Docs/ADR/change record cập nhật nếu hành vi hoặc quyết định đổi.

## Việc chưa chốt

- Ngưỡng độ phủ (nếu dùng); đề xuất không đặt số cứng, ưu tiên test đúng chỗ.
- Cách tạo OTP graph nhỏ cho test tích hợp.
- Thiết bị tối thiểu cho kiểm tay mobile (Android/iOS đời nào).
