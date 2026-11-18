# 06 — Research (kế thừa từ repo cũ)

- **Trạng thái:** Draft — **chưa kiểm chứng lại ở v2**
- **Ngày:** 18-11-2026
- **Liên quan:** `../00-foundation/05-legacy-carryover.md`, `../00-foundation/02-scope.md`

Các tài liệu dưới đây nằm ở repo cũ và **không được copy nguyên văn**. Mỗi dòng ghi nguồn và ngày đọc theo quy ước (`Nguồn: ../GoWander/<path>, đọc 18-11-2026`). Trước khi dùng làm căn cứ cho gate, phải đọc lại bản gốc và kiểm tra điều khoản hiện hành.

## Pháp lý

- **Nguồn:** `../GoWander/docs/research/legal-compliance-checklist-2026-11-17.md`. Đây là tài liệu nghiên cứu nội bộ, **không phải tư vấn pháp lý**; mọi điểm đánh dấu “Cần luật sư” phải được luật sư xác nhận.
- **Điểm chính cho v2** (theo bản gốc):
  - Luật 91/2025/QH15 và Nghị định 356/2025/NĐ-CP có hiệu lực từ **01-01-2026**, thay Nghị định 13/2023. **Vị trí là dữ liệu cá nhân nhạy cảm**.
  - DPIA và hồ sơ chuyển dữ liệu xuyên biên giới: nộp trong **60 ngày**; miễn trừ cho doanh nghiệp nhỏ có thể không áp dụng khi xử lý dữ liệu nhạy cảm.
  - Luật An ninh mạng 116/2025 và Nghị định 333/2026: nghĩa vụ lưu trữ dữ liệu trong nước có thể áp dụng; app dùng Cloudflare R2 cần rà soát.
  - Bản đồ: phải thể hiện đúng chủ quyền, biên giới; tile OSM/OpenFreeMap **không mặc nhiên** đáp ứng.
  - Nghị định 147/2024: tài khoản + nội dung/bình luận/theo dõi có thể là “mạng xã hội”.
  - TMĐT/booking, quảng cáo, người tiêu dùng, bản quyền ảnh, app store (privacy, xoá tài khoản), accessibility.
  - Phân giai đoạn: **P0** trước pilot (privacy, terms, thông báo vị trí, attribution + chủ quyền bản đồ, ảnh có giấy phép); **P1** mở tài khoản; **P2** cộng đồng/**ảnh người dùng** (chặn mở tính năng); **P3** bán vé/đặt chỗ; **P4** quảng cáo/affiliate.
- **Hệ quả đã áp dụng trong v2 (18-11-2026):** ảnh người dùng chuyển sang R1.6; thêm gate phân loại Nghị định 147 cho R1.5; gate DPIA ở R1.

## Pháp lý đối tác (thêm 18-11-2026, chưa kiểm chứng)

Danh sách giấy tờ đăng ký hợp tác ở `../01-product/08-accounts-roles-partners.md` (mục 4) dựa trên hiểu biết chung, **chưa đối chiếu văn bản gốc và chưa có luật sư xác nhận**:

- Đăng ký hộ kinh doanh / doanh nghiệp: NĐ 01/2021/NĐ-CP.
- An toàn thực phẩm (ăn uống): Luật ATTP 2010, NĐ 15/2018/NĐ-CP (giấy chứng nhận hoặc cam kết khi thuộc diện miễn).
- Lưu trú: điều kiện an ninh, trật tự theo NĐ 96/2016/NĐ-CP.
- Lữ hành: Luật Du lịch 2017.
- Dữ liệu cá nhân trong hồ sơ: Luật 91/2025/QH15 và NĐ 356/2025/NĐ-CP (theo checklist cũ, thay NĐ 13/2023 từ 01-01-2026).
- Nếu GoWander thu cọc hộ quán: có thể là sàn TMĐT, phải đăng ký với Bộ Công Thương (NĐ 52/2013, sửa bởi NĐ 85/2021). **Phải hỏi luật sư trước khi bật.**

## Phí đối tác (cần nghiên cứu, thêm 18-11-2026)

Chủ dự án chưa chốt có thu phí đăng ký / phí tháng / phí % hay không. Cần làm trước khi vẽ màn “Gói & thanh toán”:

- [ ] Khảo sát phí của các nền tảng đặt bàn và giao đồ ăn đang chạy ở TP.HCM (ghi nguồn, ngày xem).
- [ ] Ước chi phí mỗi quán: SMS/OTP, email, phí cổng thanh toán, hỗ trợ, lưu giấy tờ.
- [ ] Hỏi luật sư/kế toán: phí % trên cọc có làm GoWander thành sàn TMĐT không; nghĩa vụ hoá đơn khi thu phí tháng.
- [ ] So sánh 4 hướng: (A) miễn phí đăng ký + Cơ bản 0 ₫ + Pro trả tháng + % trên cọc; (B) chỉ %; (C) chỉ phí tháng; (D) con số do chủ dự án đặt.

## Nguồn dữ liệu

- **Nguồn:** `../GoWander/docs/research/data-sources-inventory.md` (DATA-1, 14-11-2026), kèm `data/sources/data-register.csv`, `source-inventory.csv`.
- **Điểm chính:**
  - Dữ liệu mở dùng được: OSM (ODbL), Wikidata (CC0), Wikimedia Commons (giấy phép từng ảnh), Wikipedia/Wikivoyage (CC BY-SA), Foursquare OS Places/Overture (theo từng bản ghi).
  - Wikidata có 1.910 đối tượng có toạ độ trong khung TP.HCM (đo 14-11-2026).
  - **Không dùng:** dữ liệu app khác (Google, Naver, Kakao, BusMap, Vietmap) khi không có hợp đồng; GADM; Open-Meteo bản miễn phí cho thương mại.
  - Ảnh vệ tinh (MapTiler, Esri, Mapbox, Sentinel-2): gói miễn phí không đủ; **chờ chủ dự án chọn nhà cung cấp**.
- **Giới hạn của chính bản gốc:** tác giả ghi “chỉ đọc kết quả tìm kiếm, chưa đọc trọn trang nào”; điều khoản giấy phép phải đọc lại trang gốc.

## Transit

- **Nguồn:** `../GoWander/docs/research/transit-data-feasibility.md` — **chưa đọc** ở v2. Kết luận cô đọng nằm trong `ADR-020` và `ADR-032` (xem `05-legacy-carryover.md`): không có GTFS chính thức dùng được; feed hiện tại là giả định có nhãn.
- **Việc cần làm:** xin dữ liệu chính thức từ cơ quan quản lý giao thông công cộng và đơn vị vận hành Metro (đã hoãn ở repo cũ, chưa huỷ).

## Việc tiếp theo

- [ ] Đọc `transit-data-feasibility.md` và viết `01-data-sources.md` (kèm trạng thái license/ngày kiểm chứng).
- [ ] Luật sư rà checklist pháp lý trước gate R1 và R1.5.
- [ ] Đối chiếu số điều luật với văn bản gốc trước khi trích dẫn trong policy.
