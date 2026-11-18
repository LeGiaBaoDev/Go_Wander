# 04 — Showcase Web

- **Trạng thái:** Draft
- **Ngày:** 18-11-2026 (rút gọn R1 cùng ngày)
- **Chủ sở hữu:** chủ dự án
- **Liên quan:** `../00-foundation/01-vision.md`, `../00-foundation/02-scope.md`, `../03-decisions/ADR-0001-web-stack.md`

## TL;DR

GoWander cần một **trang web showcase** công khai, nhưng R1 chỉ cần đủ để: giải thích GoWander là gì, tải app, liên hệ, và công bố chính sách bắt buộc.

Đề xuất: **Showcase Web là phần marketing của Client Web**, dùng **Next.js + TypeScript** (ADR-0001), route group `(marketing)` tách khỏi `(app)`, **build tĩnh (static export)**.

- **R1 (6 trang):** Home, Download, Contact, Privacy, Terms, Map/Data.
- **Sau R1:** About, Services, Pricing, Partners (form), FAQ, Cookies, Content policy, blog, partner directory, status page, CMS.
- **Ngôn ngữ R1: VI + EN.** ZH hoãn sau R1.
- **Không** form thu PII, không giá, không dịch vụ trả phí, không số liệu/testimonial giả.
- **Thiết kế:** light/airy, **Sexy Blue**, 3D hero tĩnh; xem `05-brand-ui-direction.md`.

## Quyết định từ chủ dự án

| Vấn đề | Quyết định |
|---|---|
| Tên miền | **Chưa có**; chọn sau. Không chặn docs |
| Dịch vụ/giá | **Chưa chốt**; R1 không bán, không thu tiền, **không có trang Pricing/Services** |
| Người viết nội dung | **Chủ dự án** (founder) |
| Ngôn ngữ | **VI + EN** ở R1 (18-11-2026). ZH hoãn sau R1 |
| Phạm vi trang R1 | **Rút gọn còn 6 trang** (18-11-2026) |
| Form partner/liên hệ | **Không form ở R1**; Contact là địa chỉ email (không thu PII qua form, không cần endpoint/spam filter) |
| CMS | **MDX/JSON trong repo cho R1**; CMS sau khi có nhiều editor |
| Trang status | Chưa làm R1; làm khi có dịch vụ trả phí/SLA |
| Pháp nhân/hợp đồng/hoá đơn | **Chưa có**; không checkout/hoá đơn/hợp đồng ràng buộc trong R1 |

> **Phụ thuộc chặn:** Privacy và Terms phải nêu rõ bên chịu trách nhiệm xử lý dữ liệu và liên hệ. Khi chưa có pháp nhân/tên miền, hai trang này chỉ là bản nháp; không publish trước khi chủ dự án chốt bên chịu trách nhiệm.

## Mục tiêu

| Mục tiêu | Cách đo |
|---|---|
| Giải thích GoWander là gì trong 30 giây | Hỏi người thử; tỉ lệ scroll/CTA |
| Tạo niềm tin trước khi tải app | Lượt bấm Tải app |
| Công bố chính sách minh bạch | Policy có ngày hiệu lực, phiên bản |

## Người đọc

- Du khách/người dùng cuối: tải app, hiểu tính năng, đọc chính sách.
- Cơ quan/quản lý, store reviewer: privacy, terms, map/data.
- Đối tác, nhà đầu tư/báo chí: sau R1 (About/Partners).

## Sơ đồ trang

| Route | Trang | R1 | Nội dung chính |
|---|---|---|---|
| `/` | Home | ✅ | Hero, vấn đề, giải pháp **chỉ gồm tính năng R1**, cách hoạt động, CTA |
| `/download` | Tải app | ✅ | Android/iOS/TestFlight, yêu cầu |
| `/contact` | Liên hệ | ✅ | Email liên hệ (không form) |
| `/policies/privacy` | Quyền riêng tư | ✅ | Dữ liệu, vị trí, quyền, liên hệ; gồm phần cookie nếu có |
| `/policies/terms` | Điều khoản | ✅ | Điều kiện dùng |
| `/policies/map` | Bản đồ/dữ liệu | ✅ | Nguồn, `ASSUMED`/ước tính, chủ quyền, giấy phép/credit ảnh và nội dung, giới hạn |
| `/about` | Về GoWander | ⏳ | Sứ mệnh, đội |
| `/services`, `/pricing` | Dịch vụ, Giá | ⏳ | Khi có dịch vụ/giá/pháp nhân |
| `/partners` | Đối tác (form) | ⏳ R6 | Khi có Partner Web/leads |
| `/faq` | FAQ | ⏳ | Sau usability test |
| `/policies/cookies`, `/policies/content` | Cookie, Nội dung | ⏳ | Khi có cookie không thiết yếu / nội dung người dùng (R1.5+) |
| `/policies/community` | Cộng đồng | ⏳ R4 | Quy tắc post/comment |
| `/policies/refunds`, `/accessibility`, `/security` | Hoàn tiền, Tiếp cận, Bảo mật | ⏳ | Khi cần |
| `/blog`, `/status` | Cập nhật, Trạng thái | ❌ R1 | Sau |

## Home — bố cục R1

1. **Hero**: “Move Smart. Explore More.” + một câu định vị + CTA Tải app.
2. **Vấn đề**: đi bus/metro ở TP.HCM khó; thông tin rải rác; ngôn ngữ.
3. **Giải pháp (chỉ tính năng R1)**: bản đồ, tìm tuyến đi bộ/bus/metro/xe đạp tại TP.HCM, Discover/Guide, plan cá nhân, lưu địa điểm, VI/EN.
4. **Cách hoạt động**: 3 bước (`Tìm → Chọn tuyến → Đi`), không nhắc Wander Mode/Journey Memory.
5. **Tính năng**: card ngắn, screenshot thật (nếu chưa có, mock có nhãn “demo”).
6. **Niềm tin**: dữ liệu `ASSUMED`/ước tính, license, privacy, không bán dữ liệu.
7. **CTA cuối + Footer**: chính sách, liên hệ, ngôn ngữ.

> **Không đưa vào Home R1:** Wander Mode, Journey Memory, thời tiết, quảng cáo/affiliate/booking, giá, logo/quote đối tác. Có thể thêm “Sắp có” sau khi chủ dự án duyệt, ghi rõ chưa có.

## Chính sách bắt buộc

| Chính sách | Nội dung tối thiểu | R1 |
|---|---|---|
| **Privacy** | Dữ liệu thu thập (đặc biệt vị trí), mục đích, lưu trữ, chia sẻ, quyền, liên hệ, cookie (nếu có) | ✅ |
| **Terms** | Điều kiện dùng, giới hạn trách nhiệm; phần tài khoản bổ sung ở R1.5 | ✅ |
| **Map/data** | Nguồn, ASSUMED/ước tính, chủ quyền, giấy phép/credit ảnh và nội dung, giới hạn | ✅ |
| Cookies, Content/photo | Khi có cookie không thiết yếu / nội dung người dùng | ⏳ R1.5+ |
| Refund/cancel, Community, Accessibility, Security, Partner/ads | Khi cần | ⏳ |

Mỗi policy có **ngày hiệu lực, phiên bản**; ngôn ngữ VI/EN, bản VI có hiệu lực nếu có tranh chấp (ghi rõ). Khi thêm ZH sau R1, bản ZH là bản dịch tham khảo.

## Kiến trúc

```text
web/apps/client-web (Next.js + TypeScript, static export ở R1)
  app/
    (marketing)/        # Showcase: SSG, SEO
      page.tsx          # /
      download/page.tsx
      contact/page.tsx
      policies/[slug]/page.tsx   # privacy, terms, map
    (app)/              # Client Web app (sau R1)
```

- R1 chỉ dựng `(marketing)`; `(app)` để sau.
- Build tĩnh, host riêng (CDN/static hosting); không chạy Node chung host với OTP/DB.
- Nội dung R1: **MDX/JSON trong repo** (versioned, không phụ thuộc dịch vụ ngoài); CMS sau.
- SEO: metadata, OG image, sitemap, robots, JSON-LD (Organization), hreflang VI/EN.
- Analytics: không dùng cookie theo dõi ở R1; nếu thêm, consent trước.
- Không form ở R1 → không PII qua web, không cần endpoint/anti-spam cho Showcase.
- Không dùng ảnh/logo đối tác khi chưa có phép.

## R1 — phạm vi Showcase

**Có:**

- Home, Download, Contact (email), Privacy, Terms, Map/Data.
- **VI/EN**; responsive; accessibility cơ bản (contrast AA kiểm bằng số).
- SEO/OG/sitemap; `hreflang` VI/EN.
- MDX/JSON trong repo; schema nội dung có validate.
- Screenshot app thật (nếu chưa có, placeholder có nhãn “demo”).
- 3D hero tĩnh (WebP/AVIF) + fallback 2D; không WebGL ở R1.

**Chưa có:**

- About, Services, Pricing, Partners, FAQ, Cookies/Content policy, blog, status, CMS.
- ZH và mọi ngôn ngữ thứ ba.
- Form (liên hệ/partner), booking/thanh toán.

## Gate R1 cho Showcase

- [ ] Chủ dự án chốt bên chịu trách nhiệm (pháp nhân hoặc cá nhân) và tên miền cho Privacy/Terms/Contact.
- [ ] Privacy/Terms/Map-Data được luật sư rà.
- [ ] Không có số liệu/testimonial/logo đối tác khi chưa xin phép; không hứa tính năng chưa có.
- [ ] Lighthouse mobile ≥ 90 (mục tiêu, không phải cam kết tuyệt đối).
- [ ] Kiểm tra VI/EN, link, 404, sitemap.
- [ ] Không thu thập dữ liệu trước consent.

## Sau R1 (khung nháp — chưa công bố)

> Phần này là khung để chủ dự án điền khi có pháp nhân/giá. **Không publish số khi chưa duyệt.**

### Catalog dịch vụ

| Dịch vụ | Đối tượng | Mô tả | Giá |
|---|---|---|---|
| **GoWander Free** | Người dùng | Map, route, discover, plan cơ bản | Miễn phí |
| **Business Basic** | Đối tác | Hồ sơ địa điểm, thông tin, ảnh, link | TBD |
| **Business Verified** | Đối tác | Xác minh, badge | TBD |
| **Sponsored placement** | Đối tác | Vị trí quảng cáo có nhãn, tách ranking (R6) | TBD |
| **Affiliate / Booking referral** | Đối tác | Referral/commission, đối soát (R6–R7) | TBD |
| **Content sponsorship** | Đối tác | Nội dung/guide có tài trợ, gắn nhãn | TBD |
| **API/data licensing** | Doanh nghiệp | API/tích hợp dữ liệu | TBD |
| **GoWander Plus** | Người dùng | Tính năng nâng cao | TBD |

### Nguyên tắc giá

- **Minh bạch**: giá niêm yết, VAT rõ, thời hạn, cách huỷ.
- **Không bán thứ hạng**: ranking tự nhiên tách khỏi ads/verified.
- **Không phí ẩn**; trial/giảm giá có điều kiện.
- **Đồng tiền**: VND; hoá đơn/thanh toán theo luật Việt Nam, cần tư vấn kế toán/thuế.

### Trang status/incident

Trang công khai cho biết hệ thống đang ổn định/suy giảm/gián đoạn. Cần khi có dịch vụ trả phí/SLA; có thể dùng Better Stack/Instatus hoặc trang tĩnh ở R6.

### CMS hay MDX/JSON?

**R1: MDX/JSON trong repo** (ổn định, có version/rollback, SSG nhanh, không thêm auth/hosting CMS). Chuyển CMS khi có ≥ 2 editor thường xuyên hoặc cần workflow duyệt/publish; đánh giá Directus hoặc Payload ở R2/R6.

## Câu hỏi mở

1. Tên miền: `gowander.vn` / `.com` / `.app`? Kiểm tra availability/trademark.
2. Ai là bên chịu trách nhiệm trong Privacy/Terms khi chưa có pháp nhân?
3. Khi nào bổ sung About/Pricing/Partners (sau R1 nào)?
4. ZH: khi nào thêm, ai dịch (người bản ngữ hay dịch máy + review)?
