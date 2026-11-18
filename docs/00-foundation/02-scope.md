# 02 — Scope & Release Plan

- **Trạng thái:** Draft
- **Ngày:** 18-11-2026
- **Chủ sở hữu:** chủ dự án
- **Liên quan:** `01-vision.md`, `03-principles.md`, `../02-architecture/02-containers.md`

## TL;DR

GoWander là **nền tảng khám phá và di chuyển thông minh**: pilot sâu tại TP.HCM, bản đồ/discover phủ Việt Nam, metro TP.HCM trước; bus và các mode khác mở rộng theo dữ liệu.

- **2 app mobile mục tiêu:** GoWander Client (du khách) và GoWander Business (Partner + Admin).
- **3 khu vực web:** Client Web (Showcase marketing + app), Partner Web, Admin Web.
- **1 backend dùng chung:** modular monolith, các domain Identity · Mobility · Discover · Journey · AI · Community · Ads · Affiliate · Booking.
- **Hero feature:** **Wander Mode** — lập hành trình theo thời gian, sở thích, điểm xuất phát; có Plan B và giải thích gợi ý.
- **Bản release đầu tiên (R1):** Client Mobile + Admin Web + **Showcase Web** (marketing), bản đồ, Discover/Guide, đi bộ/bus/metro/xe đạp, plan cá nhân cơ bản, **lưu địa điểm local**, song ngữ, dữ liệu và kiểm thử đạt gate.
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

## Phạm vi địa lý

GoWander không phủ đều mọi nơi ngay từ R1. Có 2 lớp:

| Lớp | Phạm vi R1 | Ghi chú |
|---|---|---|
| **Pilot sâu** | **TP.HCM** | Bus + Metro L1 + đi bộ + xe đạp; dữ liệu transit giả định/ASSUMED; Discover/Guide TP.HCM |
| **Toàn quốc** | Bản đồ, tìm địa điểm, Discover/Guide, lưu địa điểm | Đi bộ/xe đạp best-effort theo OSM; bus/metro chỉ khi dữ liệu đủ |

### Ma trận coverage R1

| Khu vực | Map/Search | Discover/Guide | Đi bộ/Xe đạp | Bus | Metro |
|---|---|---|---|---|---|
| **TP.HCM** | ✅ | ✅ sâu | ✅ | ✅ `ASSUMED` | ✅ L1 `ASSUMED` |
| **Hà Nội** | ✅ | ⏳ biên tập sau | ✅ best-effort | ⏳ data gate | ⏳ sau |
| **Đà Nẵng/Hội An/Đà Lạt/Phú Quốc** | ✅ | ⏳ biên tập sau | ✅ best-effort | ⏳ data gate | ❌ |
| **Tỉnh/thành khác** | ✅ | ⏳ | ✅ best-effort | ⏳ data gate | ❌ |

Quy tắc:

- **Metro:** chỉ TP.HCM (L1) trong R1. Hà Nội/metro khác là phase sau, cần graph/GTFS riêng.
- **Bus:** TP.HCM là pilot. Tỉnh/thành khác chỉ bật routing khi có GTFS/OSM đủ tốt; nếu không, chỉ map/discover.
- **Đi bộ/xe đạp:** best-effort toàn quốc theo OSM; không hứa đủ đường/hẻm.
- **Ô tô/xe máy:** chưa làm; cần engine/audit riêng.
- **Ngoài coverage:** app vẫn xem map/place; tab tuyến hiện **“Chưa hỗ trợ tìm tuyến tại khu vực này”**, không bịa route.
- **Không gộp OTP toàn quốc vào một graph HCMC.** Mỗi vùng có graph/coverage riêng khi đủ dữ liệu.

### Trạng thái dữ liệu

| Status | Nghĩa | Dùng ở đâu |
|---|---|---|
| `OFFICIAL` | Nguồn chính thức | Khi có GTFS/GTFS-RT chính thức |
| `ASSUMED` | Lịch/route sinh cho pilot, không phải chính thức | TP.HCM R1 |
| `BEST_EFFORT` | OSM/geometry, không có timetable đầy đủ | Walk/bike/bus ngoài HCMC khi bật |
| `UNSUPPORTED` | Chưa hỗ trợ routing | Khu vực chưa có graph/dữ liệu |

## Vòng lặp sản phẩm

```text
Discover → Wander Mode → Explore → Journey Memory → Community → Discover Again
```

| Giai đoạn | Vòng lặp | Release |
|---|---|---|
| Design gate | Figma design system + luồng R1 + prototype | R0.5 |
| Foundation / Pilot | Discover + plan cơ bản + Explore | R1 |
| Identity & Contributions | Tài khoản client + đề xuất địa điểm + moderation | R1.5 |
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
| **R0.5** | **Design gate** | Figma trước khi code UI | Design tokens (Sexy Blue), components, màn R1 (Client/Admin/Showcase), prototype, accessibility | Màn R2/R4, AI, booking | Figma + prototype được duyệt; dev feasible |
| **R1** | **Pilot Foundation** | Người dùng đầu tiên khám phá và đi lại có giá trị | Client Mobile + Admin Web + Showcase Web (marketing); map; discover/guide; walk/bus/metro/bicycle; plan cá nhân cơ bản; **lưu địa điểm local**; VI/EN; dữ liệu + test | Community, AI, Ads, Affiliate, Booking, Partner Web, Business Mobile, Client Web app, **thêm địa điểm** | 20–30 user thử; route/admin/UI test; không P0 |
| **R1.5** | **Identity & Contributions** | Người dùng có tài khoản và đóng góp địa điểm có kiểm duyệt | Client account (email/password + Google, verified); form đề xuất địa điểm + **1–3 ảnh**; Admin queue; trạng thái submission; audit | Ảnh bình luận, collections, community, OSM export, AI, Ads | Auth + moderation + photo/duplicate test; legal terms |
| **R2** | **Wander Mode** | Hero feature: lập hành trình theo thời gian/sở thích/điểm xuất phát | Itinerary nhiều điểm; giải thích gợi ý; Plan B; recommendation; weather-aware | Community, AI, Ads | Usability test đạt; plan quality |
| **R3** | **Retention** | Người dùng quay lại và lưu kỷ niệm | Journey Memory; Storybook; lưu/mở lại; chia sẻ riêng tư | Community đầy đủ, AI | Retention test |
| **R4** | **Community Beta** | Nội dung người dùng và vòng lặp xã hội | Post, Follow, Comment, Post-to-Journey, moderation | Ads, Booking | Legal + moderation gate |
| **R5** | **AI Beta** | AI hỗ trợ lập kế hoạch có kiểm soát | AI Companion; nguồn dữ liệu giới hạn; cost cap; eval | Community mở rộng, Booking | Eval + cost gate |
| **R6** | **Business Beta** | Doanh thu và đối tác | Partner Web; Ads; Affiliate; analytics; đối soát | Booking sâu | Partner + ad policy gate |
| **R7** | **Booking** | Đặt dịch vụ | External booking → Partner API → tích hợp sâu | — | Hợp đồng + pháp lý |
| **R8** | **Expansion** | Mở rộng | Business Mobile; địa lý; cá nhân hóa nâng cao | — | Gate theo thị trường |

## R0.5 — Design gate (Figma)

- **Mục tiêu:** chốt giao diện R1 trong Figma trước khi viết UI code.
- **Phải có:**
  - Design tokens: Sexy Blue, typography VI/EN/ZH, spacing, radius, shadow.
  - Components: button, input, select, card, chip, table, dialog, toast, empty/loading/error, map overlay.
  - Client Mobile R1: onboarding, map, search, place detail, Discover, Guide, route results/detail, navigation, plan, saved places, settings/language.
  - Admin Web R1: login, dashboard, places, guide, feedback, media; trạng thái bảng/form.
  - Showcase Web R1: Home, Services, Pricing, Partners, Download, FAQ, Contact, Policies.
  - Prototype cho 3 luồng chính: tìm tuyến, lưu địa điểm, admin duyệt nội dung (R1).
  - Accessibility: contrast, focus, touch target, text size; copy VI/EN/ZH thật.
  - Handoff: component names, states, spacing, responsive breakpoints.
- **Không có:** màn R2/R4, community, AI, booking, ảnh comment.
- **Gate:** chủ dự án duyệt Figma/prototype; dev xác nhận khả thi; token/component dùng được trong code.

## R1 — Bản release đầu tiên

### Người dùng và mục tiêu

- 20–30 người thử tại TP.HCM: du khách + người mới đi bus/metro.
- Mục tiêu: từ mở app đến hoàn thành một chuyến/plan mà không cần trợ giúp.

### Luồng end-to-end bắt buộc

1. Mở app, chọn **Tiếng Việt / English**.
2. Xem bản đồ, tìm địa điểm, xem chi tiết + nguồn/giấy phép.
3. Tìm tuyến **đi bộ / bus / metro / xe đạp** tại TP.HCM; ngoài TP.HCM xem map/discover, tuyến chỉ khi có coverage.
4. Khám phá Discover/Guide, lưu địa điểm, thêm địa điểm vào plan.
5. Tạo plan nhiều điểm, xem thời gian/đường đi, lưu và mở lại.
6. Admin thêm/sửa địa điểm hoặc nội dung guide; thay đổi hiện trong app.
7. Mở Showcase Web, đọc dịch vụ/giá/chính sách và tải app.
8. Mọi dữ liệu giả định/ước tính được ghi rõ.

### Phạm vi R1

| Khu vực | Phải có |
|---|---|
| **Client Mobile** | Android + iOS; onboarding; map 2D/3D/satellite; search; place detail; Discover; Guide; plan cơ bản; Settings/Language; feedback |
| **Smart Mobility** | TP.HCM: đi bộ, bus, metro, xe đạp; so sánh tuyến; chỉ dẫn từng bước; nhãn `ASSUMED`; ngoài TP.HCM: map/discover, route chỉ khi coverage; lỗi mạng rõ ràng |
| **Discover & Guide** | Địa điểm, category, tìm kiếm, guide TP.HCM, câu chuyện; nguồn/giấy phép |
| **Plan** | Tạo plan, thêm/xoá/sắp xếp điểm, thời gian ước tính, lưu trên máy, mở lại |
| **Saved places** | Lưu/bỏ lưu địa điểm; danh sách Đã lưu local; mở place/chỉ đường/thêm vào plan |
| **Admin Web** | Đăng nhập; quản lý places/guide/feedback/media; phân quyền cơ bản |
| **Showcase Web** | Home, About, Services, Pricing, Download, FAQ, Contact, Policies; VI/EN; SEO |
| **Data** | OSM/GTFS giả định; `dataStatus`; seed TP.HCM; test route case |
| **Quality** | Unit/integration/UI test; accessibility cơ bản; crash-free; evidence |

### Không có trong R1

- Community (post/follow/comment).
- **Routing bus/metro toàn quốc** (chỉ TP.HCM R1; vùng khác theo coverage).
- **Metro Hà Nội/thành phố khác** (sau).
- **Bus liên tỉnh** (không thuộc R1).
- **Thêm địa điểm** (R1.5; cần tài khoản + moderation).
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
- [ ] Saved places: lưu/bỏ lưu/mở lại; không mất khi restart.
- [ ] Mobile build Android/iOS; UI/accessibility test.
- [ ] 20–30 user thử; ≥ 70% hoàn thành luồng.
- [ ] Showcase Web: nội dung/giá/policy được duyệt; không số liệu giả; VI/EN; SEO/accessibility cơ bản.
- [ ] Coverage: TP.HCM route đủ 4 mode; ngoài TP.HCM hiện đúng `BEST_EFFORT`/`UNSUPPORTED`, không bịa route.
- [ ] Không P0; P1 có kế hoạch.

## R1.5 — Identity & Contributions

- **Mục tiêu:** người dùng có tài khoản client và có thể đề xuất địa điểm; Admin duyệt trước khi publish.
- **Phải có:**
  - Client account: email/password + Google; email verified (hoặc Google).
  - Đăng nhập/đăng ký/khôi phục; token Keychain/Keystore.
  - Form đề xuất địa điểm (tên, category, vị trí, địa chỉ, mô tả, **1–3 ảnh optional**, đồng ý điều khoản).
  - Trạng thái `pending/approved/rejected/merged/withdrawn`; “Đề xuất của tôi”.
  - Admin queue: map preview, duplicate candidates, approve/reject/merge, audit.
  - Rate limit theo account; anti-spam; duplicate check PostGIS.
  - Ảnh: upload, strip EXIF, resize/thumb, duyệt từng ảnh; license/credit; lưu media.
  - Legal: điều khoản đóng góp, privacy, content policy, ảnh.
- **Không có:** ảnh bình luận, collection, community, OSM export, AI, Ads.
- **Gate:** auth flow; submission + 1–3 ảnh → moderation → publish; duplicate/merge; photo reject/EXIF; rate limit; legal terms; 10 submission test.

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
- **Auth:** role `client`, `partner`, `admin`, `editor`; R1 dùng `admin`/`editor`; `client` bắt đầu R1.5.
- **Mobile:** 2 app target; Business app dùng chung codebase, khác entry/navigation theo role.
- **Web:** 3 khu vực; R1 dùng Admin + Showcase marketing; Client Web app/Partner Web sau.
- **API:** `/api/v1`; module mới thêm path riêng, không phá contract cũ.
- **Analytics:** event schema từ R1 để R6 đối soát.
- **Ads:** ad slot tách khỏi ranking; sponsored label bắt buộc.

## Đã chốt trong phiên 18-11-2026

- **Xe đạp = xe đạp**, có trong **R1** (kèm gate test OSM/OTP).
- **Client Web app không có trong R1**; Showcase marketing có trong R1.
- **Client guest-first**: không bắt buộc tài khoản; plan/saved places lưu trên máy.
- Tên release: **R1 = Pilot Foundation**, **R1.5 = Identity & Contributions**, **R2 = Wander Mode**.
- **Thêm địa điểm bắt buộc tài khoản**; lưu địa điểm guest-first; thêm địa điểm ở R1.5.
- **Địa lý:** pilot sâu TP.HCM (metro + bus); map/discover toàn quốc; bus/metro vùng khác theo coverage, không hứa toàn quốc R1.
- Account: email/password + Google, verify; rate limit 3/ngày; contributor hiện tên; **1–3 ảnh** khi đề xuất; comment text-only ở R4, ảnh comment sau.

## Câu hỏi mở

1. AI provider/chi phí: chưa chốt; R5 mới làm.
2. Community: có cần cho pilot không? (đề xuất không)
3. Booking: đối tác nào? (R7)
4. Public store release: R1 hay sau usability? (đề xuất pilot build trước)

## Việc tiếp theo

- [ ] Duyệt `01-vision.md`, file này và `03-principles.md`.
- [ ] Duyệt `04-glossary.md`.
- [ ] Duyệt `06-saved-and-contributed-places.md`.
- [ ] Làm Figma R0.5 (design tokens + màn R1 + prototype).
- [ ] Duyệt `02-architecture/01-system-context.md` (đã viết).
- [ ] Duyệt `02-architecture/02-containers.md` (đã viết).
- [ ] Duyệt `01-product/04-showcase-web.md` (đã viết).
- [x] ADR-0001 web stack — đã chốt.
- [ ] Viết ADR-0002 API/auth + API contract v0.
- [ ] Lập data/legal risk register cho R1.
