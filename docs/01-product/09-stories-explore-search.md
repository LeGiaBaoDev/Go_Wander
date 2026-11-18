# 09 — Câu chuyện, Khám phá Việt Nam, tìm theo loại

- **Trạng thái:** Proposed — chờ chủ dự án duyệt
- **Ngày:** 18-11-2026
- **Chủ sở hữu:** chủ dự án
- **Liên quan:** `07-design-canvas.md`, `05-brand-ui-direction.md` (ảnh và bản quyền), `../00-foundation/02-scope.md` (R1: Discover & Guide), `../06-research/README.md` (bản đồ, chủ quyền)

## TL;DR

- **Câu chuyện:** bài về khu lịch sử và công trình nổi bật, có ảnh, ảnh so sánh xưa/nay, dòng thời gian, **audio Gô đọc** và **video**, nút “Đi tới đây”. Admin soạn bằng **trình soạn theo khối (Editor.js)**.
- **Khám phá Việt Nam:** bản đồ cả nước, **có Hoàng Sa và Trường Sa ghi rõ thuộc Việt Nam**, chọn vùng → danh sách nổi bật.
- **Tìm theo loại:** bấm “Quán ăn” thì bản đồ chỉ hiện quán ăn kèm danh sách; “Bệnh viện” thì hiện bệnh viện + nút gọi 115.
- Đường vào nằm ở tab **Khám phá** và chip dưới ô tìm của **Bản đồ**.

## 1. Đường vào (app)

| Tính năng | Mở từ |
|---|---|
| Câu chuyện | Khám phá → ô “Câu chuyện Sài Gòn”; danh mục “Lịch sử”; thẻ “Có câu chuyện” trên địa điểm; vùng trong Khám phá Việt Nam |
| Khám phá Việt Nam | Khám phá → ô “Khám phá Việt Nam” |
| Tìm theo loại | Bản đồ → chip dưới ô tìm (Quán ăn, Cà phê, Bệnh viện, Ga metro, ATM); Tìm kiếm → “Tìm theo loại”; Khám phá → danh mục |

Trên iPad: cùng nội dung, bố cục 2 cột (danh sách + bản đồ/ảnh).

## 2. Câu chuyện

### Phía khách

- Danh sách lọc theo: Tất cả / Khu lịch sử / Công trình / Gần bạn / Có audio. Thẻ hiện thời gian đọc, thời lượng audio, cách đi tới (metro/bus/đi bộ).
- Màn đọc: ảnh bìa có ghi công, tiêu đề, trình phát **Gô kể chuyện** (tốc độ 1×…), đoạn văn, ảnh so sánh, dòng thời gian, video có phụ đề, nút **Đi tới đây** và **Giờ mở cửa**; lưu, chia sẻ.
- Gợi ý “đi bộ một vòng N câu chuyện” mở hành trình (R2 Wander).
- Nội dung lịch sử phải biên tập và kiểm chứng; bản vẽ ghi “Nội dung mẫu”.

### Phía Admin (soạn bài)

- Trình soạn **theo khối** kiểu Editor.js, lưu **JSON** để app hiển thị native; **không cho dán HTML tự do**.
- Khối: đoạn văn, tiêu đề, ảnh / so sánh ảnh, trích dẫn, dòng thời gian, **thẻ địa điểm + “Đi tới đây”**, **audio** (Gô đọc tự động hoặc tải file), **video** (tải lên hoặc YouTube).
- Thiết lập: loại (Khu lịch sử / Công trình), gắn địa điểm, mô tả ngắn ≤ 140 ký tự, thẻ, nguồn tham khảo, lịch đăng, VI/EN (có % dịch).
- Quy trình: Nháp → Gửi duyệt → (người thứ hai duyệt) → Lên lịch / Đã đăng.
- **Bản quyền media bắt buộc:** mỗi ảnh/video có tác giả, giấy phép (Tự quay/sở hữu, CC BY, CC BY-SA, CC0, Phạm vi công cộng, Được cấp phép + đính kèm), link nguồn. **Thiếu giấy phép → không đăng được.** App tự hiện dòng ghi công.

## 3. Khám phá Việt Nam

- Bản đồ cả nước; **quần đảo Hoàng Sa (Đà Nẵng) và Trường Sa (Khánh Hoà) ghi rõ thuộc Việt Nam**; có nhãn Biển Đông.
- Chọn vùng (Bắc / Trung / Nam) → thẻ tỉnh/thành nổi bật; vùng có tìm tuyến (R1: chỉ TP.HCM) có nhãn “Có tìm tuyến”, vùng khác “Chỉ khám phá”.
- Bản vẽ dùng đường biên đơn giản hoá từ GeoJSON công khai (`johan/world.geo.json`) và **ghi “Ranh giới: minh hoạ”**. Sản phẩm thật phải dùng bản đồ đúng chủ quyền, biên giới theo quy định (xem `06-research`, mục Bản đồ).

## 4. Tìm theo loại

- Bản đồ chỉ hiện **ghim của loại đã chọn**; ghim khác mờ đi; ghim đầu tiên nổi bật kèm nhãn.
- Sheet dưới: số kết quả trong bán kính, sắp xếp (Gần nhất / Đang mở / Đánh giá cao), mỗi dòng: tên, loại, khoảng cách, đi bộ hoặc tuyến, trạng thái mở cửa, mức giá; nút **Đường đi** (hoặc **Gọi** với bệnh viện).
- **Bệnh viện:** thanh “Khẩn cấp? Gọi cấp cứu 115” cố định trên cùng danh sách; nhãn “Cấp cứu 24/7”.
- Loại nhiều hơn 5 → “Thêm ▾” (quy tắc dropdown ở `05-brand-ui-direction.md`).

## Câu hỏi mở

1. Câu chuyện thuộc R1 (Discover & Guide đã có “câu chuyện”) hay tách phase? Audio/video có ở R1 không?
2. Nguồn dữ liệu bệnh viện, ATM, nhà thuốc (OSM? danh sách Sở Y tế?) và nhãn `dataStatus`.
3. Nhà cung cấp TTS tiếng Việt chạy trên máy cho “Gô đọc”.
4. Bản đồ nền đáp ứng chủ quyền cho Khám phá Việt Nam.
