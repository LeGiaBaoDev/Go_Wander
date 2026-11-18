# 03 — Principles

- **Trạng thái:** Draft
- **Ngày:** 18-11-2026
- **Chủ sở hữu:** chủ dự án
- **Liên quan:** `01-vision.md`, `02-scope.md`, `../04-development/02-conventions.md`

Các nguyên tắc dưới đây dùng để **ra quyết định**, không phải khẩu hiệu. Khi hai lựa chọn xung đột, nguyên tắc đứng trước thứ tự càng nhỏ càng quan trọng.

## 1. Docs-first, quyết định trước khi code

**Nghĩa là:** một tính năng/kiến trúc mới phải có docs hoặc ADR đủ rõ trước khi viết code. Code không được là nơi “thử xem có gì”.

**Không được:** bắt tay code một domain mới chỉ vì “làm nhanh trước”; để docs tụt lại sau code.

**Cách kiểm tra:** PR có link tới docs/ADR/change record; người review hiểu “vì sao” mà không cần hỏi tác giả.

## 2. Sự thật có bằng chứng, dữ liệu giả định phải ghi rõ

**Nghĩa là:** mọi tuyên bố “đã chạy/đã đo/đã kiểm chứng” phải có lệnh, log, ảnh hoặc số liệu. Dữ liệu mô phỏng/ước tính phải hiện `ASSUMED`, `DEMO` hoặc “ước tính”.

**Không được:** bịa test result; hiển thị GTFS giả định như lịch chính thức; ghi “đã test trên iOS” khi chưa mở iOS.

**Cách kiểm tra:** change record có mục **Bằng chứng**; UI có nhãn tin cậy; docs có ngày kiểm chứng.

## 3. Người dùng có giá trị trước, doanh thu sau

**Nghĩa là:** R1 phải giúp người dùng khám phá/đi lại thật, trước khi nghĩ tới ads, affiliate, booking.

**Không được:** thêm quảng cáo/affiliate làm thay đổi thứ hạng gợi ý hoặc tuyến; chặn tính năng cơ bản sau đăng nhập/trả phí khi chưa có giá trị.

**Cách kiểm tra:** ads có nhãn “Được tài trợ”; ranking tự nhiên tách khỏi vị trí trả tiền; R1 không phụ thuộc doanh thu.

## 4. Wander Mode là tính năng chủ lực, nhưng không được bỏ quên nền tảng

**Nghĩa là:** mọi quyết định sản phẩm phải trả lời: “cái này giúp Wander Mode tốt hơn, hay chỉ là tính năng phụ?”. Đồng thời R1 vẫn phải có map/discover/route/plan đủ tốt để người dùng đầu tiên dùng được.

**Không được:** làm Community/AI/Ads trước khi Wander Mode có dữ liệu/route/place đủ tốt; hoặc ngược lại, chỉ làm bản đồ mà không có lớp hành trình.

**Cách kiểm tra:** roadmap R1→R2; mỗi feature có “phục vụ hero feature nào”.

## 5. Một backend, modular monolith; 2 mobile app, 3 web

**Nghĩa là:** backend là một khối Go với các module domain rõ ràng; mobile Client/Business dùng chung codebase nhưng khác entry/navigation theo role; web có Client/Partner/Admin.

**Không được:** tách microservice sớm; tạo app mobile thứ ba; để mobile gọi thẳng OTP/DB; nhân bản logic giữa các app.

**Cách kiểm tra:** sơ đồ container; module boundary; API contract; role/permission matrix.

## 6. API contract là hợp đồng, không phải gợi ý

**Nghĩa là:** OpenAPI là nguồn sự thật; mobile/Admin chỉ phụ thuộc API, không phụ thuộc implementation.

**Không được:** đổi response/field mà không cập nhật spec; trả lỗi khác format; để client tự đoán trạng thái.

**Cách kiểm tra:** contract test; OpenAPI lint; client mock từ spec.

## 7. Local-first, guest-first, privacy by default

**Nghĩa là:** R1 không bắt buộc tài khoản; plan/saved places lưu trên máy; chỉ thu thập dữ liệu tối thiểu; vị trí chỉ dùng khi cần.

**Không được:** bắt đăng nhập để xem bản đồ/tìm tuyến; gửi vị trí lên server mặc định; log token/vị trí/query nhạy cảm; bán dữ liệu.

**Cách kiểm tra:** privacy checklist; data flow diagram; log review; quyền truy cập.

## 8. Bảo mật và pháp lý là điều kiện, không phải việc dọn sau

**Nghĩa là:** trước khi release phải rà: dữ liệu cá nhân/vị trí, bản đồ chủ quyền, ảnh/giấy phép, nội dung cộng đồng, TMĐT/booking, app store.

**Không được:** dùng ảnh Google/Street View; dùng tile không rõ license; mở community/booking khi chưa có tư vấn; giấu rủi ro.

**Cách kiểm tra:** legal checklist; source/license register; DPIA sơ bộ; security review.

## 9. Song ngữ và tiếp cận từ đầu

**Nghĩa là:** VI/EN là một phần của thiết kế, không phải lớp dịch sau. Accessibility (contrast, touch target, screen reader, không chỉ dùng màu) là tiêu chí release.

**Không được:** hardcode chuỗi; dùng màu làm tín hiệu duy nhất; nút < 48 dp; text bị cắt trên màn nhỏ.

**Cách kiểm tra:** chuỗi qua `Strings`/i18n; contrast test; UI test trên màn nhỏ; screen reader check.

## 10. Test và bằng chứng là một phần của “xong”

**Nghĩa là:** một task chỉ “xong” khi có test/build/evidence tương ứng; nếu chưa kiểm chứng, ghi rõ “chưa kiểm chứng”.

**Không được:** nói “đã xong” khi chưa chạy; merge khi build đỏ; bỏ qua test vì “chắc là được”.

**Cách kiểm tra:** definition of done; CI xanh; change record có lệnh/kết quả.

## 11. Đơn giản, vận hành được bởi đội nhỏ

**Nghĩa là:** ưu tiên giải pháp ít thành phần, managed service khi hợp lý; không thêm hạ tầng nếu chưa có nhu cầu.

**Không được:** Kubernetes/microservices/event bus khi chưa cần; tự viết auth/booking nếu mua được; tối ưu sớm.

**Cách kiểm tra:** ADR có phần “đánh đổi/vận hành”; số thành phần trong sơ đồ; chi phí ước tính.

## 12. Mọi thay đổi có dấu vết

**Nghĩa là:** commit Conventional Commits; ADR cho quyết định lớn; change record theo ngày; issue/PR template; tag release.

**Không được:** commit trộn nhiều tính năng; nhánh dài ngày; sửa docs lịch sử mà không ghi đè; xoá bằng chứng.

**Cách kiểm tra:** git history sạch; `docs/07-changes/`; ADR index; release notes.

## Quy tắc ra quyết định nhanh

Khi phân vân, hỏi theo thứ tự:

1. Có phục vụ **R1 end-to-end** hoặc **Wander Mode** không? Nếu không → defer.
2. Dữ liệu có **thật/đủ tin** không? Nếu không → label hoặc không hiển thị.
3. Có chạm **PII/vị trí/pháp lý** không? Nếu có → privacy/legal review trước.
4. Có làm **ranking/tuyến bị lệch vì tiền** không? Nếu có → tách ads khỏi kết quả.
5. Có thể **test/đo** không? Nếu không → chưa “xong”.
6. Có làm **kiến trúc phức tạp hơn** không? Nếu có → ADR + lý do.

## Điều không được làm

- Không bịa dữ liệu, không bịa bằng chứng.
- Không dùng ảnh/bản đồ không rõ license.
- Không để ads/affiliate thay đổi thứ hạng tự nhiên.
- Không bắt đăng nhập cho tính năng cơ bản.
- Không thêm mode/dịch vụ mới khi chưa có dữ liệu/engine.
- Không merge khi build/test đỏ.
