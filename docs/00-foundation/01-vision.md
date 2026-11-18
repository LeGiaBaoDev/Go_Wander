# 01 — Vision

- **Trạng thái:** Draft
- **Ngày:** 18-11-2026
- **Chủ sở hữu:** chủ dự án
- **Liên quan:** `02-scope.md`, `03-principles.md`, `../01-product/01-users.md`

> **Move Smart. Explore More.**

## TL;DR

GoWander là **nền tảng khám phá và di chuyển thông minh**: pilot sâu tại TP.HCM (tìm tuyến + Discover/Guide), bản đồ và tìm kiếm phủ Việt Nam, không chỉ là app tìm đường.

- **2 app mobile mục tiêu:** GoWander Client (du khách) và GoWander Business (Partner + Admin).
- **3 khu vực web:** Client Web, Partner Web, Admin Web.
- **1 backend dùng chung:** modular monolith.
- **Hero feature:** **Wander Mode** — lập hành trình theo thời gian, sở thích, điểm xuất phát; có Plan B và giải thích gợi ý.
- **Bản release đầu tiên (R1):** Client Mobile + Admin Web + **Showcase Web** (marketing), map, Discover/Guide, đi bộ/bus/metro/xe đạp **tại TP.HCM**, plan cá nhân cơ bản, song ngữ.
- **Địa lý:** pilot sâu TP.HCM (tìm tuyến + Discover/Guide); map/search toàn quốc; tìm tuyến vùng khác chỉ khi có graph và dữ liệu riêng (sau R1).

Điểm khác biệt cốt lõi: GoWander không chỉ giúp bạn **tìm đường tới một địa điểm**, mà còn giúp bạn **khám phá nơi đáng đến, xây dựng chuyến đi, điều chỉnh kế hoạch, lưu giữ kỷ niệm và chia sẻ trải nghiệm**.

## Vấn đề

Một du khách hoặc người mới chuyển đến TP.HCM thường gặp:

- Bus có hàng trăm tuyến, nhiều tuyến đổi lộ trình, tên trạm không thống nhất; Google Maps đôi khi thiếu tuyến hoặc không đúng giờ.
- Không biết nên đi bus nào, xuống trạm nào, đi bộ bao xa, có phải đổi tuyến không.
- Thông tin địa điểm, món ăn, câu chuyện lịch sử nằm rải rác, không gắn với tuyến đang đi.
- Rào cản ngôn ngữ: du khách cần tiếng Anh; người địa phương cần tiếng Việt.
- Thời tiết, mưa, nắng nóng ảnh hưởng trực tiếp đến lịch trình đi bộ/đi bus.
- Đi sai đường, rẽ nhầm một chiều, gặp hẻm quá nhỏ hoặc cầu/hầm không phù hợp phương tiện.
- Lên kế hoạch một ngày phải mở nhiều app: bản đồ, thời tiết, địa điểm, ghi chú, ảnh.

## Giải pháp

GoWander là một hệ sinh thái gồm:

1. **GoWander Client** — Android/iOS; Client Web (Showcase marketing R1, app sau): bản đồ, tìm tuyến, Discover/Guide, Wander Mode, Journey Memory.
2. **GoWander Business** — Partner + Admin (Mobile + Web): quản lý nội dung, địa điểm, ảnh, feedback, đối tác, quảng cáo, booking.
3. **Go Backend** — modular monolith: Identity · Mobility · Discover · Journey · AI · Community · Ads · Affiliate · Booking.
4. **Data & Integrations** — PostgreSQL/PostGIS, OTP, Photon, MapLibre, media storage, đối tác ngoài.

R1 chỉ cần **Client Mobile + Admin Web** (cùng Showcase Web rút gọn); geo: TP.HCM là pilot sâu, map/search toàn quốc. Các phần còn lại có vị trí trong kiến trúc nhưng không chặn release đầu tiên.

## Vòng lặp sản phẩm

```text
Discover → Wander Mode → Explore → Journey Memory → Community → Discover Again
```

| Tầng | Vai trò |
|---|---|
| **Discover** | Tìm địa điểm, món ăn, câu chuyện, hidden gems |
| **Wander Mode** | Biến sở thích + thời gian + vị trí thành hành trình có lý do |
| **Explore** | Đi theo tuyến, xem bản đồ, chỉ dẫn, thời tiết |
| **Journey Memory** | Lưu plan, ảnh, ghi chú, storybook riêng tư |
| **Community** | Chia sẻ, follow, comment, post-to-journey |
| **Discover Again** | Nội dung và trải nghiệm mới kéo người dùng quay lại |

Ads, Affiliate và Booking tạo doanh thu **xung quanh** vòng lặp này, không được làm sai lệch đề xuất tự nhiên hoặc kết quả tuyến.

## Người dùng mục tiêu

| Nhóm | Nhu cầu chính |
|---|---|
| Du khách quốc tế | Đi từ sân bay/khách sạn đến điểm tham quan bằng bus/metro; hiểu rõ chỉ dẫn bằng tiếng Anh |
| Người địa phương mới đi bus/metro | Tuyến nhanh, ít đi bộ, ít đổi tuyến, biết trạm và giờ |
| Người khám phá | Tìm quán ăn, địa điểm, câu chuyện gần tuyến hoặc gần mình |
| Cặp nhóm/du khách tự túc | Lên lịch trình nhiều điểm, chia sẻ, lưu kỷ niệm |
| Đối tác/Admin | Quản lý nội dung, địa điểm, ảnh, feedback, khuyến mãi, booking |

## Lời hứa sản phẩm

- **Rõ ràng:** mỗi chặng nói rõ đi gì, từ đâu, đến đâu, mất bao lâu, tốn bao nhiêu.
- **Trung thực:** dữ liệu giả định được ghi rõ; không giả vờ có dữ liệu thật.
- **An toàn:** cảnh báo đi sai, tránh loại đường không phù hợp phương tiện.
- **Địa phương:** nội dung Việt Nam, tiếng Việt là công dân hạng nhất, tiếng Anh đầy đủ.
- **Riêng tư:** vị trí chỉ dùng khi cần; plan lưu trên máy; không bán dữ liệu.

## Khác Google Maps ở đâu

| Google Maps | GoWander |
|---|---|
| Trả lời “Từ A đến B thế nào?” | Trả lời “Hôm nay đi đâu, bằng gì, mất bao lâu, tốn bao nhiêu, nếu mưa thì sao?” |
| Tối ưu thời gian | Tối ưu thời gian / chi phí / ít đi bộ / tránh mưa / tránh đông |
| POI + review đại trà | Nội dung địa phương tuyển chọn, gắn với tuyến |
| Một tuyến A→B | Plan nhiều điểm, có giờ, thứ tự, Plan B |
| Dữ liệu hiển thị “như thật” | Ghi rõ `ASSUMED` / ước tính / chưa chính thức |
| Toàn cầu, hời hợt | TP.HCM sâu trước; map/search toàn quốc; hiểu hẻm, bus, metro, mùa mưa |
| Tài khoản + quảng cáo + tracking | Guest-first, local-first, tối thiểu dữ liệu |

Không cạnh tranh với Google ở: độ phủ bản đồ, realtime toàn cầu, Street View, review số lượng lớn, traffic toàn cầu. GoWander thắng ở **một job hẹp: lên và sống một ngày ở Sài Gòn**.

## Thước đo thành công giai đoạn pilot

| Thước đo | Mục tiêu sơ bộ |
|---|---|
| Người dùng thử | 20–30 người |
| Tìm được tuyến cho các cặp điểm phổ biến | ≥ 90% test case |
| Hoàn thành một chuyến thử (không cần trợ giúp) | ≥ 70% người thử |
| Hiểu chỉ dẫn (không hỏi lại) | ≥ 80% |
| Lỗi chặn (crash, không tìm được tuyến) | 0 trong buổi thử |
| Thời gian tải tuyến | ≤ 3 giây (mạng tốt) |

Các số này là **giả thuyết**, sẽ hiệu chỉnh sau usability test.

## Ưu tiên theo release

- **R1 — Pilot Foundation:** map/search toàn quốc; Discover/Guide + route đi bộ/bus/metro/xe đạp tại TP.HCM; plan cơ bản; lưu địa điểm; Admin Web; Showcase Web rút gọn.
- **R2 — Wander Mode:** hero feature; lập hành trình, Plan B, giải thích, weather-aware.
- **R3 — Retention:** Journey Memory, Storybook.
- **R4 — Community Beta:** post, follow, comment, moderation.
- **R5 — AI Beta:** AI companion có giới hạn nguồn/chi phí/kiểm chứng.
- **R6 — Business Beta:** Partner Web, Ads, Affiliate, analytics.
- **R7 — Booking:** external → Partner API → tích hợp sâu.
- **R8 — Expansion:** địa lý, cá nhân hóa (giao diện quán trong app thay Business Mobile, xem ADR-0002).

Chi tiết: [`02-scope.md`](./02-scope.md).

## Không phải điều kiện của R1

- Community, AI, Ads, Affiliate, Booking.
- Partner Web, Business Mobile, Client Web app (Showcase marketing có trong R1).
- Tìm tuyến ngoài TP.HCM (mọi mode); metro Hà Nội/thành phố khác; bus liên tỉnh.
- Ô tô, xe máy, phà.
- Realtime GTFS/GTFS-RT.
- Review/rating cộng đồng đầy đủ.

Các phần này **có chỗ trong kiến trúc** (module/role/feature flag), nhưng không được chặn người dùng đầu tiên.

## Câu hỏi mở

- Mở rộng bus/metro tỉnh/thành nào trước: Hà Nội, Đà Nẵng, hay theo đối tác?
- Ngôn ngữ thứ ba (Trung/Hàn/Nhật): đã hoãn sau R1 (18-11-2026); chọn ngôn ngữ nào sau pilot?
- AI provider/chi phí và mức kiểm chứng nào là đủ cho R5?
- Mô hình doanh thu nào được thử trong R6 (affiliate, quảng cáo, gói trả phí)?
- Ai chịu trách nhiệm nội dung/ảnh và kiểm duyệt cộng đồng?
