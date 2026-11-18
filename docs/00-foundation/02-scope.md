# 02 — Scope & Release Plan

- **Trạng thái:** Draft
- **Ngày:** 18-11-2026
- **Chủ sở hữu:** chủ dự án
- **Liên quan:** `01-vision.md`, `03-principles.md`, `../02-architecture/02-containers.md`

## TL;DR

GoWander là **nền tảng khám phá và di chuyển thông minh** cho TP.HCM, không chỉ là app tìm đường.

- **2 app mobile mục tiêu:** GoWander Client (du khách) và GoWander Business (Partner + Admin).
- **3 khu vực web:** Client Web (Showcase marketing + app), Partner Web, Admin Web.
- **1 backend dùng chung:** modular monolith, các domain Identity · Mobility · Discover · Journey · AI · Community · Ads · Affiliate · Booking.
- **Hero feature:** **Wander Mode** — lập hành trình theo thời gian, sở thích, điểm xuất phát; có Plan B và giải thích gợi ý.
- **Bản release đầu tiên (R1):** Client Mobile + Admin Web + **Showcase Web** (marketing), bản đồ, Discover/Guide, đi bộ/bus/metro/xe đạp, plan cá nhân cơ bản, song ngữ, dữ liệu và kiểm thử đạt gate.
- Community, AI, quảng cáo, affiliate, booking **có chỗ trong kiến trúc nhưng không phải điều kiện** để R1 ra mắt.

## Cấu trúc sản phẩm

```text
GoWander Ecosystem
├── GoWander Client        Android / iOS / Client Web (Showcase R1, app sau)
├── GoWander Business      Partner + Admin (Mobile + Web)
└── Go Backend — Modular Monolith
    ├── Identity
    ├── Mobility
    ├── Discover
    ├── Journey
    ├── AI
    ├── Community
    ├── Ads
    ├── Affiliate
    └── Booking
        │
Data & Integrations
PostgreSQL/PostGIS · OTP · Photon · MapLibre · Media Storage · External Partners
```

Nguyên tắc:

- **Không phải 3 app mobile.** Chỉ 2 app mobile; Business dùng phân quyền để hiện giao diện Partner hoặc Admin.
- **Một backend.** Các domain là module trong monolith, không tách microservice ở giai đoạn này.
- **R1 cần Client Mobile + Admin Web + Showcase Web (marketing).** Client Web app, Partner Web và Business Mobile mở sau theo nhu cầu.
- Các domain chưa làm vẫn có **seam** trong kiến trúc: module, role, feature flag, API versioning, analytics event.

## Vòng lặp sản phẩm

```text
Discover → Wander Mode → Explore → Journey Memory → Community → Discover Again
```

| Giai đoạn | Vòng lặp | Release |
|---|---|---|
| Foundation / Pilot | Discover + plan cơ bản + Explore | R1 |
| Wander Mode | Lập hành trình, giải thích, Plan B | R2 |
| Retention | Journey Memory, Storybook | R3 |
| Community Beta | Post, Follow, Comment, Post-to-Journey | R4 |
| AI Beta | AI Travel Companion có giới hạn | R5 |
| Business Beta | Partner Web, Ads, Affiliate, Analytics | R6 |
| Booking | External → Partner API → deep integration | R7 |
| Expansion | Business Mobile, địa lý, cá nhân hóa | R8 |

## Nguyên tắc xuyên suốt

1. **Wander Mode là tính năng chủ lực**, nhưng R1 phải có giá trị trước: map, discover, guide, plan cơ bản.
2. **Ads/Affiliate/Booking không được làm sai lệch** kết quả tuyến hoặc thứ hạng gợi ý. Quảng cáo phải gắn nhãn; ranking tự nhiên tách khỏi vị trí trả tiền.
3. **Community, AI, Business không chặn R1.** Chúng được thiết kế seam từ đầu, nhưng không xây trước khi có người dùng thật.
4. **AI có giới hạn nguồn dữ liệu, chi phí và kiểm chứng đầu ra.** Không để AI bịa địa điểm/giờ/tuyến.
5. **Dữ liệu giả định phải ghi rõ** (`ASSUMED`/`DEMO`/ước tính), không hiển thị như dữ liệu chính thức.
6. **Privacy by default:** plan và lịch sử lưu trên máy; không cần tài khoản cho R1.
7. **Guest-first:** R1 không bắt buộc tài khoản client; plan và saved places lưu trên máy. Tài khoản chỉ cần cho đồng bộ/community ở release sau.

## Release plan

> **Đã chốt:** R1 = **Pilot Foundation**; R2 = **Wander Mode**. R1 gồm xe đạp và Showcase Web (marketing); Client Web app chưa làm; client dùng guest-first.

| Release | Tên | Mục tiêu | Phải có | Không có trong release | Gate |
|---|---|---|---|---|---|
| **R0** | Docs & Architecture | Chốt lại dự án trước khi code | Vision, scope, principles, ADR stack, API contract v0, data/legal risk register | Code sản phẩm | Docs được duyệt |
| **R1** | **Pilot Foundation** | Người dùng đầu tiên khám phá và đi lại có giá trị | Client Mobile + Admin Web + Showcase Web (marketing); map; discover/guide; walk/bus/metro/bicycle; plan cá nhân cơ bản; VI/EN; dữ liệu + test | Community, AI, Ads, Affiliate, Booking, Partner Web, Business Mobile, Client Web app | 20–30 user thử; route/admin/UI test; không P0 |
| **R2** | **Wander Mode** | Hero feature: lập hành trình theo thời gian/sở thích/điểm xuất phát | Itinerary nhiều điểm; giải thích gợi ý; Plan B; recommendation; weather-aware | Community, AI, Ads | Usability test đạt; plan quality |
| **R3** | **Retention** | Người dùng quay lại và lưu kỷ niệm | Journey Memory; Storybook; lưu/mở lại; chia sẻ riêng tư | Community đầy đủ, AI | Retention test |
| **R4** | **Community Beta** | Nội dung người dùng và vòng lặp xã hội | Post, Follow, Comment, Post-to-Journey, moderation | Ads, Booking | Legal + moderation gate |
| **R5** | **AI Beta** | AI hỗ trợ lập kế hoạch có kiểm soát | AI Companion; nguồn dữ liệu giới hạn; cost cap; eval | Community mở rộng, Booking | Eval + cost gate |
| **R6** | **Business Beta** | Doanh thu và đối tác | Partner Web; Ads; Affiliate; analytics; đối soát | Booking sâu | Partner + ad policy gate |
| **R7** | **Booking** | Đặt dịch vụ | External booking → Partner API → tích hợp sâu | — | Hợp đồng + pháp lý |
| **R8** | **Expansion** | Mở rộng | Business Mobile; địa lý; cá nhân hóa nâng cao | — | Gate theo thị trường |

## R1 — Bản release đầu tiên

### Người dùng và mục tiêu

- 20–30 người thử tại TP.HCM: du khách + người mới đi bus/metro.
- Mục tiêu: từ mở app đến hoàn thành một chuyến/plan mà không cần trợ giúp.

### Luồng end-to-end bắt buộc

1. Mở app, chọn **Tiếng Việt / English**.
2. Xem bản đồ, tìm địa điểm, xem chi tiết + nguồn/giấy phép.
3. Tìm tuyến **đi bộ / bus / metro / xe đạp**, xem các chặng và bản đồ.
4. Khám phá Discover/Guide, thêm địa điểm vào plan.
5. Tạo plan nhiều điểm, xem thời gian/đường đi, lưu và mở lại.
6. Admin thêm/sửa địa điểm hoặc nội dung guide; thay đổi hiện trong app.
7. Mở Showcase Web, đọc dịch vụ/giá/chính sách và tải app.
8. Mọi dữ liệu giả định/ước tính được ghi rõ.

### Phạm vi R1

| Khu vực | Phải có |
|---|---|
| **Client Mobile** | Android + iOS; onboarding; map 2D/3D/satellite; search; place detail; Discover; Guide; plan cơ bản; Settings/Language; feedback |
| **Smart Mobility** | Đi bộ, bus, metro, xe đạp; so sánh tuyến; chỉ dẫn từng bước; tuyến/nhãn `ASSUMED`; lỗi mạng rõ ràng |
| **Discover & Guide** | Địa điểm, category, tìm kiếm, guide TP.HCM, câu chuyện; nguồn/giấy phép |
| **Plan** | Tạo plan, thêm/xoá/sắp xếp điểm, thời gian ước tính, lưu trên máy, mở lại |
| **Admin Web** | Đăng nhập; quản lý places/guide/feedback/media; phân quyền cơ bản |
| **Showcase Web** | Home, About, Services, Pricing, Download, FAQ, Contact, Policies; VI/EN; SEO |
| **Data** | OSM/GTFS giả định; `dataStatus`; seed TP.HCM; test route case |
| **Quality** | Unit/integration/UI test; accessibility cơ bản; crash-free; evidence |

### Không có trong R1

- Community (post/follow/comment).
- AI Companion.
- Ads, Affiliate, Booking.
- Partner Web, Business Mobile.
- Client Web app (Showcase marketing có trong R1; app sau).
- Car, motorbike, ferry.
- Realtime GTFS/GTFS-RT.
- Review/rating cộng đồng đầy đủ.

### Quyết định mode cho R1

| Mode | R1 | Ghi chú |
|---|---|---|
| Đi bộ | ✅ | OTP/OSM; chặng nối |
| Bus | ✅ | GTFS giả định; nhãn `ASSUMED` |
| Metro | ✅ | Metro L1 giả định; nhãn `ASSUMED` |
| Xe đạp | ✅ R1 | OTP/OSM; phải test route case + hạ tầng trước gate R1 |
| Ô tô | ❌ | Cần audit one-way/cầu/hầm/cấm rẽ |
| Xe máy | ❌ | OTP không có profile; cần engine riêng |
| Phà/buýt sông | ❌ | Chờ dữ liệu |
| Ride-hailing | 🔗 | Link-out sau; không tìm tuyến |

> “Bike” trong tài liệu này là **xe đạp**. Xe máy là mode riêng, chưa làm.

### Gate R1

- [ ] Docs/ADR liên quan được duyệt.
- [ ] API contract v0 + OpenAPI chạy.
- [ ] Data source/license/assumption register.
- [ ] Security/privacy checklist + DPIA sơ bộ.
- [ ] Legal: bản đồ chủ quyền, ảnh, dữ liệu vị trí.
- [ ] Route test ≥ 90% case; không lỗi one-way/chỉ dẫn sai nghiêm trọng.
- [ ] Bike test: ≥ 10 route case OSM/OTP; không route vào đường cấm/không phù hợp xe đạp.
- [ ] Admin CRUD + phân quyền test.
- [ ] Mobile build Android/iOS; UI/accessibility test.
- [ ] 20–30 user thử; ≥ 70% hoàn thành luồng.
- [ ] Showcase Web: nội dung/giá/policy được duyệt; không số liệu giả; VI/EN; SEO/accessibility cơ bản.
- [ ] Không P0; P1 có kế hoạch.

## R2 — Wander Mode (hero)

- Input: thời gian có, điểm bắt đầu, sở thích, ngân sách, phương tiện.
- Output: lịch trình nhiều điểm có giờ, tuyến, chi phí, weather, điểm trong nhà.
- Plan B: mưa/trễ/đóng cửa → đổi điểm/thứ tự/giờ.
- Giải thích: vì sao gợi ý điểm này, vì sao tuyến này.
- Recommendation: gần bạn, trên đường, theo chủ đề/thời gian.
- Gate: usability test; plan quality; không bịa giờ/địa điểm.

## R3–R8 — Tóm tắt

- **R3 Retention:** Journey Memory (plan, ảnh, ghi chú), Storybook, lưu/mở lại, chia sẻ riêng tư.
- **R4 Community Beta:** Threads-style feed, post, follow, comment, post-to-journey, moderation; cần legal/Decree 147 review.
- **R5 AI Beta:** AI companion giới hạn nguồn (place/route/guide/weather), cost cap, eval, log; không open web.
- **R6 Business Beta:** Partner Web, Ads có nhãn, Affiliate, analytics/đối soát; ranking tự nhiên tách khỏi ads.
- **R7 Booking:** external link → partner API → deep integration; hợp đồng/pháp lý trước.
- **R8 Expansion:** Business Mobile, mở rộng địa lý, cá nhân hóa nâng cao.

## Seam kiến trúc cho các phase sau

- **Backend:** module `identity`, `mobility`, `discover`, `journey`, `community`, `ai`, `ads`, `affiliate`, `booking`; module chưa làm vẫn có interface/feature flag.
- **Auth:** role `client`, `partner`, `admin`, `editor`; R1 chỉ dùng `client`/`admin`.
- **Mobile:** 2 app target; Business app dùng chung codebase, khác entry/navigation theo role.
- **Web:** 3 khu vực; R1 chỉ Admin.
- **API:** `/api/v1`; module mới thêm path riêng, không phá contract cũ.
- **Analytics:** event schema từ R1 để R6 đối soát.
- **Ads:** ad slot tách khỏi ranking; sponsored label bắt buộc.

## Đã chốt trong phiên 18-11-2026

- **Xe đạp = xe đạp**, có trong **R1** (kèm gate test OSM/OTP).
- **Client Web app không có trong R1**; Showcase marketing có trong R1.
- **Client guest-first**: không bắt buộc tài khoản; plan/saved places lưu trên máy.
- Tên release: **R1 = Pilot Foundation**, **R2 = Wander Mode**.

## Câu hỏi mở

1. AI provider/chi phí: chưa chốt; R5 mới làm.
2. Community: có cần cho pilot không? (đề xuất không)
3. Booking: đối tác nào? (R7)
4. Public store release: R1 hay sau usability? (đề xuất pilot build trước)

## Việc tiếp theo

- [ ] Duyệt `01-vision.md`, file này và `03-principles.md`.
- [ ] Duyệt `04-glossary.md`.
- [ ] Duyệt `02-architecture/01-system-context.md` (đã viết).
- [ ] Duyệt `02-architecture/02-containers.md` (đã viết).
- [ ] Duyệt `01-product/04-showcase-web.md` (đã viết).
- [x] ADR-0001 web stack — đã chốt.
- [ ] Viết ADR-0002 API/auth + API contract v0.
- [ ] Lập data/legal risk register cho R1.
