# 04 — Showcase Web

- **Trạng thái:** Draft
- **Ngày:** 18-11-2026
- **Chủ sở hữu:** chủ dự án
- **Liên quan:** `../00-foundation/01-vision.md`, `../00-foundation/02-scope.md`, `../03-decisions/ADR-0001-web-stack.md`

## TL;DR

GoWander cần một **trang web showcase** công khai, không chỉ app. Trang này giới thiệu dự án, dịch vụ, giá, chính sách, cách tải app và cách liên hệ/trở thành đối tác.

Đề xuất: **Showcase Web là phần marketing của Client Web**, dùng **Next.js + TypeScript** (ADR-0001), tách route group `(marketing)` khỏi app `(app)`.

- **R1:** Home, About, Services, Pricing, Partners (form), Download, FAQ, Contact + 5 trang policy.
- **Sau R1:** blog/updates, partner directory, case study, status page, CMS.
- **Đa ngôn ngữ VI/EN/ZH**, SEO/SSG, accessibility, không dùng số liệu/testimonial giả.
- **Thiết kế:** light/airy, **Sexy Blue** + orange accent nhỏ, có 3D hero; xem `05-brand-ui-direction.md`.

## Quyết định từ chủ dự án (18-11-2026)

| Vấn đề | Quyết định |
|---|---|
| Tên miền | **Chưa có**; chọn sau. Tạm dùng domain/subdomain khi có; không chặn docs |
| Dịch vụ/giá | **Chưa chốt**; R1 không bán/không thu tiền. Pricing ghi “Beta / Liên hệ”, không bịa số |
| Người viết nội dung | **Chủ dự án** (founder) |
| Ngôn ngữ | **VI + EN + ZH**; EN/VI là chính, ZH cho khách Trung |
| Form partner | **Có trong R1**; chỉ thu quan tâm, chưa phải hợp đồng |
| CMS | **MDX/JSON trong repo cho R1**; CMS sau khi có nhiều editor (xem dưới) |
| Trang status | **Chưa làm R1**; giải thích ở dưới; làm khi có dịch vụ trả phí/SLA |
| Pháp nhân/hợp đồng/hoá đơn | **Chưa có**; không checkout/không hoá đơn/không hợp đồng ràng buộc trong R1 |

## Mục tiêu

| Mục tiêu | Cách đo |
|---|---|
| Giải thích GoWander là gì trong 30 giây | Tỉ lệ scroll/CTA; hỏi người thử |
| Tạo niềm tin trước khi tải app | Tải app/đăng ký thử; liên hệ |
| Giới thiệu dịch vụ + giá rõ ràng | Xem Pricing → Contact/Đăng ký partner |
| Công bố chính sách minh bạch | Trang policy có ngày cập nhật, phiên bản |
| Thu hút đối tác | Form partner/booking |

## Người đọc

- Du khách/người dùng cuối: tải app, hiểu tính năng, đọc chính sách.
- Đối tác: quán ăn, khách sạn, tour, dịch vụ; xem gói/giá/liên hệ.
- Nhà đầu tư/báo chí: hiểu dự án, đội, giai đoạn, khác biệt.
- Cơ quan/quản lý: chính sách, pháp lý, bản đồ, dữ liệu.
- Ứng viên: văn hóa/đội ngũ (sau).

## Sơ đồ trang

| Route | Trang | R1 | Nội dung chính |
|---|---|---|---|
| `/` | Home | ✅ | Hero, vấn đề, giải pháp, tính năng, cách hoạt động, CTA |
| `/about` | Về GoWander | ✅ | Sứ mệnh, đội, giai đoạn, nguyên tắc dữ liệu |
| `/services` | Dịch vụ | ✅ | Dịch vụ cho người dùng/đối tác |
| `/pricing` | Giá | ✅ | Gói, cách tính, VAT, dùng thử/huỷ |
| `/partners` | Đối tác | ✅ | Chương trình partner, form đăng ký quan tâm |
| `/download` | Tải app | ✅ | Android/iOS/TestFlight, yêu cầu |
| `/faq` | FAQ | ✅ | Câu hỏi thường gặp |
| `/contact` | Liên hệ | ✅ | Form, email, địa chỉ (nếu có) |
| `/policies/privacy` | Quyền riêng tư | ✅ | Dữ liệu, vị trí, cookie, quyền |
| `/policies/terms` | Điều khoản | ✅ | Điều kiện dùng |
| `/policies/cookies` | Cookie | ✅ | Loại cookie, mục đích, lựa chọn |
| `/policies/refunds` | Hoàn tiền/huỷ | ⏳ | Khi có booking/dịch vụ trả phí |
| `/policies/content` | Nội dung/ảnh | ✅ | License, credit, nội dung người dùng |
| `/policies/community` | Cộng đồng | ⏳ R4 | Quy tắc post/comment |
| `/policies/map` | Bản đồ/dữ liệu | ✅ | Nguồn, chủ quyền, độ chính xác |
| `/policies/accessibility` | Tiếp cận | ⏳ | Cam kết và kênh phản hồi |
| `/policies/security` | Bảo mật | ⏳ | Cách báo lỗi, phạm vi |
| `/blog` | Cập nhật | ⏳ | Bài viết, changelog |
| `/status` | Trạng thái | ❌ R1 | Health/incident công khai; làm khi có dịch vụ trả phí/SLA |

## Home — bố cục đề xuất

1. **Hero**: “Move Smart. Explore More.” + một câu định vị + CTA Tải app / Xem dịch vụ.
2. **Vấn đề**: đi bus/metro ở TP.HCM khó; thông tin rải rác; mưa/nắng; ngôn ngữ.
3. **Giải pháp**: Wander Mode, route, discover, plan, weather, Journey Memory.
4. **Cách hoạt động**: 3–4 bước (`Discover → Wander Mode → Explore → Memory`).
5. **Tính năng**: card ngắn, có screenshot thật (không mock nếu chưa có).
6. **Dịch vụ**: 3 nhóm — người dùng, đối tác, doanh nghiệp.
7. **Giá**: bảng tóm tắt + link `/pricing`.
8. **Niềm tin**: dữ liệu ASSUMED, license, privacy, không bán dữ liệu.
9. **Đối tác**: logo/quote có phép; nếu chưa có, bỏ trống — không bịa.
10. **FAQ** + CTA cuối.
11. **Footer**: chính sách, liên hệ, ngôn ngữ, mạng xã hội (nếu có).

## Dịch vụ (catalog ban đầu)

> Giá thực tế **chưa chốt**; bảng dưới là khung để chủ dự án điền. Không publish số khi chưa duyệt.

| Dịch vụ | Đối tượng | Mô tả | Giá |
|---|---|---|---|
| **GoWander Free** | Người dùng | Map, route, discover, plan cơ bản | Miễn phí |
| **Business Basic** | Đối tác | Hồ sơ địa điểm, thông tin, ảnh, link | TBD |
| **Business Verified** | Đối tác | Xác minh, badge, ưu tiên hiển thị trong ràng buộc | TBD |
| **Sponsored placement** | Đối tác | Vị trí quảng cáo có nhãn, tách ranking | TBD (CPM/CPC) |
| **Affiliate** | Đối tác | Referral/commission theo chuyển đổi | TBD (%) |
| **Booking referral** | Đối tác | Chuyển booking, đối soát | TBD |
| **Content sponsorship** | Đối tác | Nội dung/guide có tài trợ, gắn nhãn | TBD |
| **API/data licensing** | Doanh nghiệp | API/tích hợp dữ liệu (sau) | TBD |
| **GoWander Plus** | Người dùng | Tính năng nâng cao (sau) | TBD |

### Nguyên tắc giá

- **Minh bạch**: giá niêm yết, VAT rõ, thời hạn, cách huỷ.
- **Không bán thứ hạng**: ranking tự nhiên tách khỏi ads/verified.
- **Không phí ẩn**: phí setup, hoa hồng, hoàn tiền ghi rõ.
- **Thử nghiệm**: gói trial/giảm giá phải có điều kiện.
- **Đồng tiền**: VND; nếu có USD, ghi tỷ giá/nguồn.
- **Hoá đơn/thanh toán**: theo luật Việt Nam; cần tư vấn kế toán/thuế.

## Chính sách bắt buộc

| Chính sách | Nội dung tối thiểu | R1 |
|---|---|---|
| **Privacy** | Dữ liệu thu thập, mục đích, lưu trữ, chia sẻ, quyền, liên hệ, cookie | ✅ |
| **Terms** | Điều kiện dùng, tài khoản, nội dung, giới hạn trách nhiệm | ✅ |
| **Cookies** | Loại cookie, mục đích, opt-in/opt-out | ✅ |
| **Content/photo** | License, credit, nội dung người dùng, báo cáo vi phạm | ✅ |
| **Map/data** | Nguồn, ASSUMED/ước tính, chủ quyền, giới hạn | ✅ |
| **Refund/cancel** | Khi có dịch vụ trả phí/booking | ⏳ |
| **Community** | Khi có post/comment | ⏳ R4 |
| **Accessibility** | Cam kết, kênh phản hồi | ⏳ |
| **Security** | Kênh báo lỗi, phạm vi | ⏳ |
| **Partner/ads** | Điều kiện đối tác, quảng cáo, dữ liệu | ⏳ R6 |

Tất cả policy phải có **ngày hiệu lực, phiên bản, ngôn ngữ VI/EN/ZH**; bản ZH là bản dịch tham khảo, bản VI/EN có hiệu lực nếu có tranh chấp (ghi rõ).

## Kiến trúc

```text
web/apps/client-web (Next.js + TypeScript)
  app/
    (marketing)/        # Showcase: SSG/ISR, SEO
      page.tsx          # /
      about/page.tsx
      services/page.tsx
      pricing/page.tsx
      partners/page.tsx
      download/page.tsx
      faq/page.tsx
      contact/page.tsx
      policies/[slug]/page.tsx
    (app)/              # Client Web app (sau R1)
      ...
```

- R1 chỉ dựng `(marketing)`; `(app)` để sau.
- Nội dung R1: **MDX/JSON trong repo** (stable, versioned, không phụ thuộc dịch vụ ngoài); CMS sau khi có nhiều editor.
- SEO: metadata, OG image, sitemap, robots, JSON-LD (Organization/Product/FAQ), hreflang VI/EN/ZH.
- Analytics: privacy-friendly; consent trước cookie không cần thiết.
- Form contact/partner: gửi API hoặc email; chống spam; không thu PII thừa.
- Không dùng ảnh/logo đối tác khi chưa có phép.

## R1 — phạm vi Showcase

**Có:**

- Home, About, Services, Pricing (Beta/Liên hệ), Partners (form), Download, FAQ, Contact.
- Privacy, Terms, Cookies, Content, Map.
- **VI/EN/ZH**; responsive; accessibility cơ bản.
- SEO/OG/sitemap; `hreflang` cho 3 ngôn ngữ.
- MDX/JSON trong repo; schema nội dung có validate.
- Screenshot app thật (nếu chưa có, dùng placeholder có nhãn “demo”).
- 3D hero tĩnh (WebP/AVIF) + fallback 2D; không WebGL nặng ở R1.

**Chưa có:**

- Blog/updates, partner directory, case study, status page.
- CMS/Admin content (dùng MDX/JSON ở R1).
- Booking/thanh toán.
- Đa ngôn ngữ thứ ba.

## Gate R1 cho Showcase

- [ ] Nội dung dịch vụ/giá được chủ dự án duyệt; không có số bịa.
- [ ] Privacy/Terms/Cookies/Content/Map được luật sư rà.
- [ ] Không có testimonial/logo đối tác khi chưa xin phép.
- [ ] Lighthouse mobile ≥ 90 (performance/SEO/accessibility/best practices) — mục tiêu, không phải cam kết tuyệt đối.
- [ ] Kiểm tra VI/EN/ZH, link, form, 404, sitemap.
- [ ] Không thu thập dữ liệu trước consent.

## Giải thích: trang status/incident là gì

Là trang công khai cho biết hệ thống đang **ổn định / suy giảm / gián đoạn**, có lịch sử sự cố và bảo trì. Thường dùng khi có dịch vụ trả phí/SLA hoặc đối tác cần theo dõi. R1 chưa cần; có thể dùng Better Stack/Instatus hoặc một trang tĩnh đơn giản ở R6.

## CMS hay MDX/JSON?

**R1: MDX/JSON trong repo.** Lý do:

- Ổn định, không phụ thuộc dịch vụ ngoài; nội dung nằm cùng code, có version/rollback.
- Next.js build SSG nhanh; SEO tốt; không thêm auth/hosting CMS.
- Phù hợp khi chủ dự án là người viết chính, có dev hỗ trợ.

**Khi nào chuyển CMS:**

- Có ≥ 2 editor thường xuyên, cần publish/schedule không qua dev.
- Cần localized content phức tạp, workflow duyệt.
- Chọn **Directus** (self-host, Postgres, ổn định) hoặc **Payload CMS** (TypeScript/Next.js, cùng hệ sinh thái) — đánh giá ở R2/R6, không đưa vào R1.

## Câu hỏi mở

1. Tên miền: chọn `gowander.vn` / `gowander.com` / `gowander.app`? Cần kiểm tra availability/trademark.
2. Dịch vụ/giá: gói nào beta, gói nào mở bán sau khi có pháp nhân?
3. Pháp nhân/hợp đồng/hoá đơn: khi nào sẵn sàng? Trước đó không checkout/không hoá đơn.
4. Trang status: R6 hay sớm hơn nếu có partner?
5. Ai dịch ZH: người bản ngữ hay dịch máy + review?
