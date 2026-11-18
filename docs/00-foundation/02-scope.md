# 02 — Scope & Release Plan

- **Trạng thái:** Draft
- **Ngày:** 18-11-2026
- **Chủ sở hữu:** chủ dự án
- **Liên quan:** `01-vision.md`, `03-principles.md`, `../02-architecture/02-containers.md`

## TL;DR

GoWander là **nền tảng khám phá và di chuyển thông minh**: pilot sâu tại TP.HCM (tìm tuyến đi bộ/bus/metro/xe đạp), bản đồ và tìm kiếm phủ Việt Nam; vùng khác chỉ mở tìm tuyến khi có graph và dữ liệu riêng.

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
| **Pilot sâu** | **TP.HCM** | Bus + Metro L1 + đi bộ + xe đạp trên **một graph OTP TP.HCM**; dữ liệu transit giả định/ASSUMED; Discover/Guide TP.HCM |
| **Toàn quốc** | Bản đồ, tìm kiếm địa điểm/địa chỉ, xem place, lưu địa điểm | **Không** tìm tuyến ngoài TP.HCM trong R1 |

> Nguồn đối chiếu: repo cũ chỉ có graph OTP cắt theo khung TP.HCM và có test case “ngoài vùng phủ” kỳ vọng lỗi rõ ràng (`../GoWander/data/test-cases/route-cases.csv` RC-09; `docs/decisions/ADR-032`; đọc 18-11-2026).

### Ma trận coverage R1

| Khu vực | Map/Search | Discover/Guide | Đi bộ/Xe đạp | Bus | Metro |
|---|---|---|---|---|---|
| **TP.HCM** | ✅ | ✅ sâu | ✅ | ✅ `ASSUMED` | ✅ L1 `ASSUMED` |
| **Hà Nội** | ✅ | ⏳ biên tập sau | ❌ `UNSUPPORTED` | ❌ | ❌ |
| **Đà Nẵng/Hội An/Đà Lạt/Phú Quốc** | ✅ | ⏳ biên tập sau | ❌ `UNSUPPORTED` | ❌ | ❌ |
| **Tỉnh/thành khác** | ✅ | ⏳ | ❌ `UNSUPPORTED` | ❌ | ❌ |

Quy tắc:

- **Mọi mode tìm tuyến chỉ có ở TP.HCM trong R1.** Vùng khác mở sau, từng vùng một, khi có graph và dữ liệu riêng (ADR transit sau R1).
- **Metro:** chỉ TP.HCM (L1). Hà Nội/metro khác cần graph/GTFS riêng.
- **Ô tô/xe máy:** chưa làm; cần engine/audit riêng.
- **Ngoài coverage:** app vẫn xem map/place; tab tuyến hiện **“Chưa hỗ trợ tìm tuyến tại khu vực này”**, không bịa route. Điểm đi hoặc đến ngoài khung TP.HCM trả `UNSUPPORTED` (khác “không có tuyến”).
- **Không gộp OTP toàn quốc vào một graph HCMC.** Mỗi vùng có graph/coverage riêng khi đủ dữ liệu.

### Trạng thái dữ liệu

| Status | Nghĩa | Dùng ở đâu |
|---|---|---|
| `OFFICIAL` | Nguồn chính thức | Khi có GTFS/GTFS-RT chính thức |
| `ASSUMED` | Lịch/route sinh cho pilot, không phải chính thức. Ghi rõ phần **có nguồn** (ví dụ giờ đầu/cuối, số lượt Metro L1) và phần **ước tính** (ví dụ tốc độ bus) trong feed | TP.HCM R1 |
| `BEST_EFFORT` | OSM/geometry, không có timetable đầy đủ | Chưa dùng ở R1; để dành cho vùng mở sau |
| `UNSUPPORTED` | Chưa hỗ trợ routing | Mọi khu vực ngoài TP.HCM ở R1 |

## Vòng lặp sản phẩm

```text
Discover → Wander Mode → Explore → Journey Memory → Community → Discover Again
```

| Giai đoạn | Vòng lặp | Release |
|---|---|---|
| Design gate | Figma design system + luồng R1 + prototype | R0.5 |
| Foundation / Pilot | Discover + plan cơ bản + Explore | R1 |
| Identity & Contributions | Tài khoản client + đề xuất địa điểm (chữ) + moderation | R1.5 |
| Photo Contributions | Ảnh đóng góp sau cổng pháp lý | R1.6 |
| Wander Mode | Lập hành trình, giải thích, Plan B | R2 |
| Retention | Journey Memory, Storybook | R3 |
| Community Beta | Post, Follow, Comment, Post-to-Journey | R4 |
| AI Beta | AI Travel Companion có giới hạn | R5 |
| Business Beta | Partner Web, Ads, Affiliate, Analytics | R6 |
| Booking | External → Partner API → deep integration | R7 |
| Expansion | Địa lý, cá nhân hóa (chế độ quán thay Business Mobile, ở R7) | R8 |

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
| **R0.5** | **Design gate** | Thiết kế trước khi code UI | Palette v4 teal, components, **màn mọi phase R1→R7** trên canvas (app điện thoại + iPad, Admin, Partner, Showcase), responsive, accessibility | — | Canvas + docs được duyệt; dev feasible |
| **R1** | **Pilot Foundation** | Người dùng đầu tiên khám phá và đi lại có giá trị | Client Mobile + Admin Web + Showcase Web (marketing); map; discover/guide; walk/bus/metro/bicycle; plan cá nhân cơ bản; **lưu địa điểm local**; VI/EN; dữ liệu + test | Community, AI, Ads, Affiliate, Booking, Partner Web, Business Mobile, Client Web app, **thêm địa điểm** | 20–30 user thử; route/admin/UI test; không P0 |
| **R1.5** | **Identity & Contributions** | Người dùng có tài khoản và đóng góp địa điểm (chỉ chữ) có kiểm duyệt | Client account (email/password + Google, verified; xoá tài khoản trong app); form đề xuất địa điểm **không ảnh**; Admin queue; trạng thái submission; audit | **Ảnh người dùng**, ảnh bình luận, collections, community, OSM export, AI, Ads | Auth + moderation + duplicate test; legal terms; phân loại Nghị định 147 + DPIA đã được luật sư xác nhận |
| **R1.6** | **Photo Contributions** | Cho phép ảnh đóng góp sau khi qua cổng pháp lý | 1–3 ảnh/đề xuất; duyệt từng ảnh; strip EXIF; license/credit | Ảnh bình luận, community | Cổng pháp lý ảnh người dùng (bản quyền, quyền riêng tư, Nghị định 147); photo/EXIF test |
| **R2** | **Wander Mode** | Hero feature: lập hành trình theo thời gian/sở thích/điểm xuất phát | Itinerary nhiều điểm; giải thích gợi ý; Plan B; recommendation; weather-aware | Community, AI, Ads | Usability test đạt; plan quality |
| **R3** | **Retention** | Người dùng quay lại và lưu kỷ niệm | Journey Memory; Storybook; lưu/mở lại; chia sẻ riêng tư | Community đầy đủ, AI | Retention test |
| **R4** | **Community Beta** | Nội dung người dùng và vòng lặp xã hội | Post, Follow, Comment, Post-to-Journey, moderation | Ads, Booking | Legal + moderation gate |
| **R5** | **AI Beta** | AI hỗ trợ lập kế hoạch có kiểm soát | AI Companion; nguồn dữ liệu giới hạn; cost cap; eval | Community mở rộng, Booking | Eval + cost gate |
| **R6** | **Business Beta** | Doanh thu và đối tác | Partner Web; Ads; Affiliate; analytics; đối soát | Booking sâu | Partner + ad policy gate |
| **R7** | **Booking** | Đặt dịch vụ | Đặt bàn + đặt món trước tuỳ chọn + cọc theo chính sách quán; chế độ quán trong app (xem `../01-product/08-accounts-roles-partners.md`) | — | Hợp đồng + pháp lý (TMĐT nếu thu cọc hộ) |
| **R8** | **Expansion** | Mở rộng | Địa lý; cá nhân hóa nâng cao (Business Mobile riêng bỏ, thay bằng chế độ quán ở R7 — ADR-0002) | — | Gate theo thị trường |

## R0.5 — Design gate (Figma)

> **Cập nhật 18-11-2026:** chủ dự án quyết định **thiết kế mọi phase R1→R7 ngay ở R0.5** (kể cả mọi trạng thái: splash, đang tải, lỗi, sheet/form, tính năng “sắp có” làm mờ, Face ID, bong bóng góp ý), vẽ trên **canvas claude.ai** thay cho Figma. Canvas, cụm màn và cách agent đọc: `../01-product/07-design-canvas.md`; sơ đồ màn hình: `../01-product/07-screen-map.html`. Các mục “Phải có / Không có” dưới đây là bản đầu (chỉ R1), giữ để lưu vết; phạm vi vẽ hiện tại rộng hơn.

- **Mục tiêu:** chốt giao diện R1 trong Figma trước khi viết UI code.
- **Phải có:**
  - Design tokens: Sexy Blue, typography VI/EN, spacing, radius, shadow (ZH để sau R1).
  - Components: button, input, select, card, chip, table, dialog, toast, empty/loading/error, map overlay.
  - Client Mobile R1: onboarding, map, search, place detail, Discover, Guide, route results/detail, navigation, plan, saved places, settings/language.
  - Admin Web R1: login, dashboard, places, guide, feedback, media; trạng thái bảng/form.
  - Showcase Web R1: Home, Download, Contact, Privacy, Terms, Map/Data.
  - Prototype cho 3 luồng chính: tìm tuyến, lưu địa điểm, admin duyệt nội dung (R1).
  - Accessibility: contrast AA (kiểm bằng số, không chỉ bằng mắt), focus, touch target, text size; copy VI/EN thật.
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
3. Tìm tuyến **đi bộ / bus / metro / xe đạp** tại TP.HCM; ngoài TP.HCM chỉ xem map/search/place, tab tuyến hiện “Chưa hỗ trợ”.
4. Khám phá Discover/Guide, lưu địa điểm, thêm địa điểm vào plan.
5. Tạo plan nhiều điểm, xem thời gian/đường đi, lưu và mở lại.
6. Admin thêm/sửa địa điểm hoặc nội dung guide; thay đổi hiện trong app.
7. Mở Showcase Web, hiểu GoWander là gì, đọc chính sách và tải app.
8. Mọi dữ liệu giả định/ước tính được ghi rõ.

### Phạm vi R1

| Khu vực | Phải có |
|---|---|
| **Client Mobile** | Android + iOS; onboarding; map 2D/3D (không vệ tinh ở R1); search; place detail; Discover; Guide; plan cơ bản; Settings/Language; feedback |
| **Smart Mobility** | TP.HCM: đi bộ, bus, metro, xe đạp; so sánh tuyến; chỉ dẫn từng bước; nhãn `ASSUMED`; ngoài TP.HCM: map/search, route trả `UNSUPPORTED`; lỗi mạng rõ ràng |
| **Discover & Guide** | Địa điểm, category, tìm kiếm, guide TP.HCM, câu chuyện; nguồn/giấy phép |
| **Plan** | Tạo plan, thêm/xoá/sắp xếp điểm, thời gian ước tính, lưu trên máy, mở lại |
| **Saved places** | Lưu/bỏ lưu địa điểm; danh sách Đã lưu local; mở place/chỉ đường/thêm vào plan |
| **Admin Web** | Đăng nhập; quản lý places/guide/feedback/media; phân quyền cơ bản |
| **Showcase Web** | Home, Download, Contact (email), Privacy, Terms, Map/Data; VI/EN; SEO; không form thu PII; không giá/dịch vụ trả phí |
| **Data** | OSM/GTFS giả định; `dataStatus`; seed TP.HCM; test route case |
| **Quality** | Unit/integration/UI test; accessibility cơ bản; crash-free; evidence |

### Không có trong R1

- Community (post/follow/comment).
- **Tìm tuyến ngoài TP.HCM** (mọi mode; vùng khác mở sau, từng vùng một).
- **Ảnh vệ tinh** (Esri World Imagery cần giấy phép ArcGIS; xem `05-legacy-carryover.md`).
- **Ảnh do người dùng gửi** (R1.6, sau cổng pháp lý).
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
| Ô tô | ❌ | Cần audit one-way/cầu/hầm/cấm rẽ và dữ liệu giao thông. Repo cũ đã có mode CAR (ADR-032); v2 hoãn có chủ đích |
| Xe máy | ❌ | OTP không có profile; cần engine riêng |
| Phà/buýt sông | ❌ | Chờ dữ liệu |
| Ride-hailing | 🔗 | Link-out sau; không tìm tuyến |

> “Bike” trong tài liệu này là **xe đạp**. Xe máy là mode riêng, chưa làm.

### Gate R1

- [ ] Docs/ADR liên quan được duyệt.
- [ ] API contract v0 + OpenAPI chạy.
- [ ] Data source/license/assumption register.
- [ ] Security/privacy checklist + DPIA sơ bộ (Luật 91/2025 + Nghị định 356/2025: vị trí là dữ liệu nhạy cảm; hồ sơ DPIA nộp trong 60 ngày kể từ khi xử lý — luật sư xác nhận).
- [ ] Legal: bản đồ chủ quyền, ảnh, dữ liệu vị trí.
- [ ] **Dữ liệu transit trước khi cho người dùng thật thử:** có dữ liệu được phép dùng hoặc xác nhận bằng văn bản của nhà vận hành; nếu chưa, mọi giờ bus/metro hiện nhãn “ước tính” và biên bản thử ghi rõ giới hạn (điều kiện đã có ở `ADR-020` repo cũ).
- [ ] Tile bản đồ: quyết định tự host hoặc nhà cung cấp có SLA (OpenFreeMap không có SLA; `ADR-025` repo cũ).
- [ ] Route test ≥ 90% case; không lỗi one-way/chỉ dẫn sai nghiêm trọng.
- [ ] Bike test: ≥ 10 route case OSM/OTP; không route vào đường cấm/không phù hợp xe đạp.
- [ ] Admin CRUD + phân quyền test.
- [ ] Saved places: lưu/bỏ lưu/mở lại; không mất khi restart.
- [ ] Mobile build Android/iOS; UI/accessibility test.
- [ ] 20–30 user thử; ≥ 70% hoàn thành luồng.
- [ ] Showcase Web (6 trang R1): nội dung và policy được duyệt; không số liệu giả; VI/EN; SEO/accessibility cơ bản.
- [ ] Coverage: TP.HCM route đủ 4 mode; ngoài TP.HCM trả đúng `UNSUPPORTED`, không bịa route (có test case “ngoài vùng phủ”).
- [ ] Không P0; P1 có kế hoạch.

## R1.5 — Identity & Contributions

- **Mục tiêu:** người dùng có tài khoản client và có thể đề xuất địa điểm (chỉ chữ); Admin duyệt trước khi publish.
- **Phải có:**
  - Client account: email/password + Google; email verified (hoặc Google). **Xoá tài khoản trong app.** Nếu có Google Sign-In thì cân nhắc Sign in with Apple trước khi lên App Store (guideline 4.8; repo cũ `ADR-040`).
  - Đăng nhập/đăng ký/khôi phục; token Keychain/Keystore.
  - Form đề xuất địa điểm (tên, category, vị trí, địa chỉ, mô tả, đồng ý điều khoản); **không ảnh**.
  - Trạng thái `pending/changes_requested/approved/rejected/merged/withdrawn`; “Đề xuất của tôi”.
  - Admin queue: map preview, duplicate candidates, approve/reject/merge/request changes, suspend contributor, audit.
  - Rate limit theo account; anti-spam; duplicate check PostGIS.
  - Legal: điều khoản đóng góp, privacy, content policy; **phân loại Nghị định 147** (tài khoản + nội dung người dùng + tên công khai) do luật sư xác nhận trước khi mở.
- **Không có:** ảnh người dùng (R1.6), ảnh bình luận, collection, community, OSM export, AI, Ads.
- **Gate:** auth flow; submission → moderation → publish; duplicate/merge; request changes; rate limit; legal terms + phân loại NĐ 147; 10 submission test.

## R1.6 — Photo Contributions

- **Điều kiện mở:** qua cổng pháp lý cho ảnh người dùng (bản quyền, quyền riêng tư/người trong ảnh, trách nhiệm nền tảng); checklist cũ xếp “ảnh người dùng” vào nhóm chặn mở tính năng (`../GoWander/docs/research/legal-compliance-checklist-2026-11-17.md`, §0.2 P2).
- **Phải có:** 1–3 ảnh/đề xuất (optional, ≤ 5 MB, JPEG/PNG/WebP), pending, duyệt từng ảnh, strip EXIF, resize/thumb, license/credit, xoá khi place bị gỡ.
- **Gate:** photo reject/EXIF test; license/credit hiển thị; legal xác nhận.

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
- **R8 Expansion:** mở rộng địa lý, cá nhân hóa nâng cao (Business Mobile riêng bỏ — ADR-0002).

## Seam kiến trúc cho các phase sau

- **Backend:** module `identity`, `mobility`, `discover`, `journey`, `community`, `ai`, `ads`, `affiliate`, `booking`; module chưa làm vẫn có interface/feature flag.
- **Auth:** role `client`, `partner`, `admin`, `editor`; R1 dùng `admin`/`editor`; `client` bắt đầu R1.5.
- **Mobile:** 2 app target; Business app dùng chung codebase, khác entry/navigation theo role.
- **Web:** 3 khu vực; R1 dùng Admin + Showcase marketing; Client Web app/Partner Web sau.
- **API:** `/api/v1`; module mới thêm path riêng, không phá contract cũ.
- **Analytics:** event schema từ R1 để R6 đối soát; chỉ bộ đếm gộp, không PII, không lưu sự kiện từng người (hướng của `ADR-038` repo cũ).
- **Ads:** ad slot tách khỏi ranking; sponsored label bắt buộc.

## Đã chốt trong phiên 18-11-2026

- **Xe đạp = xe đạp**, có trong **R1** (kèm gate test OSM/OTP).
- **Client Web app không có trong R1**; Showcase marketing có trong R1.
- **Client guest-first**: không bắt buộc tài khoản; plan/saved places lưu trên máy.
- Tên release: **R1 = Pilot Foundation**, **R1.5 = Identity & Contributions**, **R2 = Wander Mode**.
- **Thêm địa điểm bắt buộc tài khoản**; lưu địa điểm guest-first; thêm địa điểm ở R1.5.
- **Địa lý:** pilot sâu TP.HCM (metro + bus + đi bộ + xe đạp); map/search toàn quốc; **mọi tìm tuyến ngoài TP.HCM mở sau**, không hứa trong R1.
- Account: email/password + Google, verify; rate limit 3/ngày; contributor hiện tên; comment text-only ở R4, ảnh comment sau.

## Đã chốt bổ sung — 18-11-2026 (rà soát nhất quán)

- **ZH hoãn sau R1**: R1 chỉ VI/EN (app và Showcase).
- **Showcase R1 thu gọn:** Home, Download, Contact (email), Privacy, Terms, Map/Data. About, Services, Pricing, Partners, FAQ, Cookies, Content, blog đều sau R1.
- **Tìm tuyến ngoài TP.HCM: bỏ khỏi R1** (khớp graph và test của repo cũ).
- **Admin Web giữ React + Vite** (ADR-0001).
- **Vệ tinh: bỏ khỏi R1** (cần giấy phép nhà cung cấp; quyết định lại khi chọn được nhà cung cấp trả phí).
- **Ảnh người dùng dời sang R1.6**, sau cổng pháp lý; R1.5 chỉ chữ.

## Đã chốt bổ sung — 18-11-2026 (design gate trên canvas)

- **Palette v4 teal** thay Sexy Blue; linh vật **Gô**; có **chế độ tối** cho app (`../01-product/05-brand-ui-direction.md`).
- **App có bản iPad** (ngang); web Admin/Partner/đăng ký hợp tác phải dùng được ở iPad và điện thoại.
- **Câu chuyện** (khu lịch sử, công trình) có audio Gô đọc và video; Admin soạn bằng editor theo khối; media bắt buộc ghi giấy phép. **Khám phá Việt Nam** có Hoàng Sa, Trường Sa. **Tìm theo loại** (quán ăn, bệnh viện…) hiện ghim + danh sách (`../01-product/09-stories-explore-search.md`).
- **Tài khoản:** một kiểu tài khoản, quyền theo vai trò ở từng quán; Admin là tài khoản riêng, Owner cấp, bắt buộc 2 lớp; quên mật khẩu qua email (SMTP) hoặc SMS (ADR-0002, `../01-product/08-accounts-roles-partners.md`).
- **Partner:** không đăng ký tự do; nộp **hồ sơ hợp tác kèm giấy tờ**, duyệt 1–3 ngày làm việc, báo qua email; **Admin cấp tài khoản chủ quán**.
- **Booking (R7):** đặt bàn + **đặt món trước tuỳ chọn (tên + giá)** + cọc theo chính sách quán; thực đơn/dịch vụ có thêm, sửa, xoá.
- **Admin:** có Tài chính (doanh thu, chi phí, lợi nhuận), Trung tâm thông báo, Phản hồi từ app, quản lý đối tác; số trong bản vẽ là DEMO.
- **Lựa chọn dài → dropdown/nhóm**; thanh bên Admin gom nhóm.

## Câu hỏi mở

0. Chưa chốt: mời nhân viên quán do chủ quán (đang vẽ) hay chỉ Admin cấp; câu chuyện/audio/video có ở R1 không; có thu cọc qua GoWander không (pháp lý TMĐT).
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
- [ ] Viết ADR-0003 API/auth + API contract v0.
- [ ] Chốt ADR-0002 (tài khoản và vai trò).
- [ ] Lập data/legal risk register cho R1.
