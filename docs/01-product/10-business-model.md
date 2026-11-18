# 10 — Mô hình kinh doanh theo giai đoạn

- **Trạng thái:** Proposed — chủ dự án đã chốt hướng theo giai đoạn (18-11-2026); số liệu và pháp lý **chưa kiểm chứng**
- **Ngày:** 18-11-2026
- **Chủ sở hữu:** chủ dự án
- **Liên quan:** `08-accounts-roles-partners.md`, `../03-decisions/ADR-0002-account-roles.md`, `../00-foundation/02-scope.md`, `../06-research/README.md`

## TL;DR

Doanh thu **xếp chồng theo ba giai đoạn**, mỗi giai đoạn thêm một nguồn thu và **không bỏ** nguồn cũ:

1. **Giai đoạn 1:** đặt trước **không thu tiền** + xác nhận qua email + **affiliate**. Quán dùng miễn phí.
2. **Giai đoạn 2:** **gói theo tháng** cho quán (miễn phí cơ bản, trả phí khi cần nhiều chi nhánh, thiết bị, nhân viên, báo cáo).
3. **Giai đoạn 3:** **hoa hồng theo đơn** khi khách **trả tiền trong app** (qua cổng thanh toán có giấy phép), kiểu GrabFood nhưng **không giao hàng**.

```text
GĐ 1: CÓ KHÁCH, CÓ QUÁN       GĐ 2: BÁN DỊCH VỤ            GĐ 3: HOA HỒNG
Khách ─đặt trước (miễn phí)─▶  Quán trả phí tháng           Khách ─trả tiền trong app─▶ GoWander
Quán xác nhận qua email        để có thêm tính năng                          │ giữ % hoa hồng
                                                                              ▼
Thu: affiliate                 Thu: affiliate + gói tháng   Quán nhận phần còn lại
                                                             Thu: affiliate + gói tháng + hoa hồng (+ quảng cáo sau)
```

## Vì sao không bắt đầu bằng hoa hồng

- Hoa hồng kiểu Shopee, TikTok Shop, GrabFood thu được vì **tiền đi qua nền tảng**, nên đo được từng đơn. Ở giai đoạn 1 khách trả tại quán; nền tảng **không thấy doanh thu thật**, hoa hồng theo phần trăm gần như không thu được.
- Hoa hồng kéo theo: cổng thanh toán có giấy phép và chia tiền, đối soát và trả tiền cho quán, hoàn tiền và tranh chấp, nghĩa vụ sàn thương mại điện tử, hoá đơn điện tử và thuế, chăm sóc khách hàng, chống gian lận. Các nghĩa vụ này **chưa có luật sư rà**.
- Hoa hồng chỉ đáng kể khi **đã có nhiều khách đặt qua app**. Chưa có khách mà đã làm thanh toán thì chi phí và rủi ro đã gánh trước khi có doanh thu.
- **Giao hàng không làm** (cần tài xế và vận hành giao nhận, vượt sức đội nhỏ). Chỉ **đặt món trước, đến ăn hoặc lấy tại quán**.

## Từng giai đoạn

| | Giai đoạn 1 | Giai đoạn 2 | Giai đoạn 3 |
|---|---|---|---|
| **Quán trả** | Không | Phí tháng | Phí tháng (có thể giảm) + % mỗi đơn |
| **Quán được** | Lên bản đồ, nhận đặt trước, email xác nhận | Thêm chi nhánh, thiết bị, nhân viên, báo cáo, nổi bật | Khách trả sẵn, giảm khách không đến, nhiều đơn |
| **Khách trả** | Không (trả tại quán) | Không | Trả trong app |
| **GoWander cần thêm** | Email, lịch hẹn, xác nhận | Thu tiền của quán, hoá đơn điện tử | Cổng thanh toán, đối soát, pháp lý |
| **Mức khó** | Thấp | Trung bình | Cao |

### Giai đoạn 1: đặt trước không thu tiền
- Khách chọn quán, ngày giờ, số người, tuỳ chọn món đặt trước (tên, giá, **ảnh**; trả tại quán).
- Email xác nhận cho khách và quán; quán xác nhận, đổi giờ hoặc từ chối; email nhắc lịch; khách tự huỷ.
- **Không cọc, không hoàn cọc, không đối soát tiền.** “Đối soát” ở giai đoạn này là **đối chiếu đơn theo ca và theo mã nhân viên** (đã xác nhận, đã đến, không đến), không có tiền.
- **Bắt buộc tài khoản, không đặt vãng lai:** chỉ khách đã đăng nhập (email đã xác minh) đặt được.
- **Chống đặt ảo** (không có cọc): giới hạn số đơn mỗi khách mỗi ngày, quán đánh dấu khách không đến, tạm chặn khi không đến nhiều lần. Mức cụ thể `ASSUMED`, chưa chốt.
- **Dashboard quán:** đơn, đã đến, không đến, lọc theo tháng, biểu đồ 12 tháng, **món đặt nhiều**, theo ca (`08`, mục 9); khách thấy nhãn **“Hot”**.
- **Affiliate:** link đối tác ngoài (vé, tour, đặt phòng) có nhãn “Được tài trợ”; không đổi thứ hạng gợi ý hay tuyến (nguyên tắc 3). Đối tác cụ thể chưa chốt.

### Giai đoạn 2: gói theo tháng (miễn phí cơ bản)
- **Miễn phí:** hiện quán trên bản đồ, nhận đặt trước cơ bản, 1 chi nhánh, 1 thiết bị, ít nhân viên.
- **Trả phí theo tháng:** nhiều chi nhánh, nhiều thiết bị đăng nhập cùng lúc, nhiều nhân viên, báo cáo theo ca và theo người, xếp hạng nổi bật.
- Giới hạn từng gói **chưa chốt**; đề xuất ban đầu cho gói miễn phí: 1 chi nhánh, 1 thiết bị, 2 nhân viên (`ASSUMED`, cần đối chiếu thị trường).
- Giai đoạn này **GoWander thu tiền của quán** (hoá đơn điện tử, thuế); chưa có luật sư hay kế toán rà.
- Mức giá **chưa có**: bản vẽ chỉ giữ chỗ `DEMO`. Cần nghiên cứu (`../06-research/README.md`, mục “Phí đối tác”).

### Giai đoạn 3: hoa hồng, thanh toán trong app
- Khách **trả tiền món đặt trước trong app**; GoWander giữ phần trăm, chuyển phần còn lại cho quán theo chu kỳ.
- Dùng **cổng thanh toán có giấy phép** với chia tiền cho quán (ví dụ VNPay, MoMo, ZaloPay: **chưa chọn, chưa kiểm chứng**). Không tự giữ tiền hộ khi chưa rõ nghĩa vụ giấy phép.
- Hoàn tiền, huỷ đơn, tranh chấp và chăm sóc khách là quy trình phải có trước khi bật.
- Tỷ lệ hoa hồng **chưa có**; cần khảo sát thị trường. Không lấy số của Shopee, TikTok Shop, Grab khi chưa có nguồn.
- Phần lớn màn đã vẽ sẵn trên canvas (cọc, hoàn cọc, đối soát, `PartnerPayout`, huỷ đặt bàn) được **giữ lại làm thiết kế cho giai đoạn này**, đánh dấu “Sắp có” cho tới khi mở.

## Khi nào chuyển giai đoạn

| Chuyển | Điều kiện (chủ dự án đặt ngưỡng; chưa có số liệu) |
|---|---|
| 1 → 2 | Có quán dùng thật, và có quán cần thêm tính năng (ví dụ nhiều chi nhánh) |
| 2 → 3 | Số đơn đặt trước hằng tháng đủ lớn để hoa hồng đáng kể; pháp lý thanh toán xong; có đối tác cổng thanh toán |

- **Không ép quán cũ khi đổi giai đoạn:** quán đang dùng miễn phí giữ quyền lợi một thời gian, báo trước, không đổi giá bất ngờ. Cần chính sách chuyển đổi rõ để quán không bỏ đi.

## Ảnh hưởng tới thiết kế và code

- **Dữ liệu ngay từ đầu:** đơn đặt trước có **trạng thái** (chờ xác nhận, đã xác nhận, đã đến, không đến, đã huỷ) và có chỗ cho **thanh toán** sau này (khoá ngoại tới bảng thanh toán, trạng thái thanh toán, hoa hồng) để giai đoạn 3 chỉ là thêm, không làm lại.
- **Tài khoản quán** (`ADR-0002`): giới hạn thiết bị, chi nhánh, nhân viên theo gói.
- **Admin tài chính:** giai đoạn 1 chỉ có doanh thu affiliate; giai đoạn 2 thêm doanh thu gói; giai đoạn 3 thêm hoa hồng.
- **Ảnh món:** bắt buộc khi quán thêm món đặt trước.

## Câu hỏi mở

1. Affiliate với ai: link ngoài (Klook, Agoda, Traveloka…) hay đối tác trực tiếp? (đề xuất: bắt đầu bằng link ngoài).
2. Giới hạn từng gói và mức phí tháng.
3. Ngưỡng chuyển 1 → 2 và 2 → 3.
4. Đối tác cổng thanh toán và luật sư cho giai đoạn 3.
5. Quảng cáo/nổi bật trên bản đồ: sau giai đoạn 3 hay cùng giai đoạn 2?
6. Ánh xạ giai đoạn sang R-phase trong `02-scope.md` (đề xuất: giai đoạn 1 cùng lần đầu mở Partner; chưa sửa roadmap).
