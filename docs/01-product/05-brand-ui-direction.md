# 05 — Brand & UI Direction

- **Trạng thái:** Draft
- **Ngày:** 18-11-2026
- **Chủ sở hữu:** chủ dự án
- **Liên quan:** `04-showcase-web.md`, `../03-decisions/ADR-0001-web-stack.md`, `../00-foundation/01-vision.md`

## TL;DR

Hướng thị giác: **sáng, thoáng, đáng tin, có chiều sâu**.

- Nền sáng/warm white; chữ navy đậm.
- Màu chính: **Sexy Blue** (xanh dương rực, gradient); màu nhấn nhỏ: **orange** (tùy chọn).
- **3D hero** (bản đồ/địa danh/phương tiện) + floating cards.
- Bo góc lớn, bóng mềm, nhiều khoảng trắng.
- Typography đậm, rõ; VI/EN dùng Be Vietnam Pro (ZH hoãn sau R1).
- 3D chỉ ở hero/section chính; không dùng khắp nơi, không làm nặng trang.

## Ảnh tham khảo — lấy gì, bỏ gì

> **Chưa kiểm chứng được:** 6 ảnh mẫu dưới đây chưa nằm trong repo (`docs/evidence/` chưa có). Cần lưu ảnh (hoặc link nguồn) trước khi duyệt hướng thiết kế. Ghi chú “app có dark” ở câu hỏi mở chưa khớp phạm vi R0.5 (chưa có token dark).

| Mẫu | Lấy | Không lấy |
|---|---|---|
| **1. Mobile travel app** (xanh/tím, núi, route card) | Cảm giác “travel”, card tuyến, illustration mềm | Màu tím làm màu chính; bố cục mobile |
| **2. Map/ad marketplace** (xanh đậm, map + panel) | Map làm hero, panel nổi, cảm giác SaaS chuyên nghiệp | Quá nhiều panel/filter; màu xanh rêu đậm đặc |
| **3. Air quality dashboard** (mint/teal, floating data) | Card dữ liệu nổi trên map, khoảng trắng | Mint/teal làm màu chính; data dày đặc |
| **4. 3D energy dashboard** (cam/trắng, nhà 3D) | 3D hero + floating metric; ánh sáng mềm | Mô hình 3D quá chi tiết; dashboard nặng |
| **5. 3D product showcase** (trắng/xám, phòng 3D) | Hero 3D trung tâm, tối giản, sáng | Nền xám lạnh; sản phẩm không liên quan |
| **6. 3D logistics landing** (cam/trắng, isometric) | Hero 3D + typography lớn + CTA rõ | Copy bố cục; 3D quá nhiều chi tiết |

**Kết luận:** **Sexy Blue của mẫu 2** là màu chính; **nền sáng của 5**, **map/panel của 2**, **card nổi của 3**, **3D của 4/6**, **cảm giác travel của 1**. Xanh lá/teal của mẫu 1 chỉ dùng cho success/data, không phải brand. Không copy brand/layout.

## Palette — Sexy Blue

Ảnh 2 là hướng chính: xanh dương rực, gradient, glow nhẹ, nền sáng. Ảnh 1 (green/teal) chỉ dùng cho success/data, không làm màu brand.

| Token | Hex | Dùng cho |
|---|---|---|
| `--ink` | `#071A33` | Chữ chính/heading |
| `--navy` | `#0A2A5E` | Footer, section tối, chữ trên nền sáng |
| `--blue` | `#007BFF` | **Brand chính** cho nền lớn, icon, minh hoạ, gradient. **Không** đặt chữ trắng cỡ thường lên đây (3.98:1, không đạt AA) |
| `--blue-strong` | `#0062E0` | **Nền nút primary + link/chữ xanh**; chữ trắng trên nền này đạt 5.49:1, chữ này trên nền trắng cũng đạt AA |
| `--blue-bright` | `#00A3FF` | Gradient, hover, chart; không đặt chữ trắng (2.73:1) |
| `--sky` | `#6BC5FF` | Gradient nhẹ, illustration; không đặt chữ trắng (1.9:1) |
| `--ice` | `#E6F4FF` | Surface soft, badge, icon tile |
| `--orange` | `#F97316` | Accent nhỏ (badge “Beta”, một CTA phụ); có thể bỏ. Chữ trên nền cam dùng `--ink` (6.22:1), không dùng chữ trắng (2.8:1) |
| `--coral` | `#FB7185` | Cảnh báo/“hot”, rất ít |
| `--teal` | `#0F766E` | Success/data phụ; không làm brand |
| `--bg` | `#F5FAFF` | Nền trang |
| `--surface` | `#FFFFFF` | Card |
| `--line` | `#DCEBFA` | Viền |
| `--muted` | `#5B7290` | Chữ phụ |
| `--success` | `#16A34A` | Trạng thái tốt |
| `--danger` | `#DC2626` | Lỗi |

### Gradient & glow

```css
--gradient-hero: linear-gradient(135deg, #007BFF 0%, #00A3FF 45%, #6BC5FF 100%);
--gradient-soft: linear-gradient(180deg, #F5FAFF 0%, #E6F4FF 100%);
--glow-blue: 0 20px 60px rgba(0, 123, 255, .25);
```

- Hero có thể dùng **nền sáng + gradient mesh xanh nhạt + 3D xanh** (chữ `--ink`/`--navy`), hoặc **nền `--navy`/`--blue-strong` + chữ trắng** ở section CTA.
- `--gradient-hero` (đến `--sky`) chỉ làm nền trang trí/3D; **không đặt chữ trắng lên đoạn sáng của gradient**.
- Không dùng gradient xanh cho toàn bộ trang; chỉ hero/CTA/footer.
- Mọi cặp chữ/nền phải kiểm bằng số: ≥ 4.5:1 chữ thường, ≥ 3:1 chữ lớn (≥ 24px hoặc ≥ 18.66px đậm) và thành phần UI.

Quy tắc:

- **Một màu chính:** blue; orange chỉ là điểm nhấn nhỏ, không cạnh tranh.
- **Nền sáng là chính**; nền xanh đậm chỉ ở hero/CTA/footer.
- Không dùng quá 3 màu đậm trong một section.
- Màu không phải tín hiệu duy nhất; luôn có icon/text.

## Typography

| Ngôn ngữ | Font | Weight |
|---|---|---|
| VI/EN | **Be Vietnam Pro** | 400/500/600/700/800 |
| ZH (sau R1) | **Noto Sans SC** | 400/500/700 |
| Fallback | system-ui, -apple-system, “Segoe UI”, sans-serif | — |

- Display/heading: Be Vietnam Pro 700/800, tracking hơi âm (`-0.02em`).
- Body: 16–18px, line-height 1.6.
- UI nhỏ: 12–14px; không dưới 12px trên mobile.
- ZH: tăng line-height 1.7; tránh chữ quá đậm ở size nhỏ.

## 3D & illustration

**Dùng 3D khi:**

- Hero showcase: bản đồ/isometric địa danh, xe bus/metro, điện thoại.
- Section “cách hoạt động”: 3–4 object nhỏ (pin, route, plan, memory).
- CTA/partner: object đối tác (quán ăn, khách sạn) cách điệu.

**Không dùng 3D khi:**

- Trang policy/FAQ/contact — chỉ typography.
- Bảng giá/Admin — ưu tiên rõ ràng.
- Mobile low-end — tránh WebGL.

**Định dạng:**

- Hero 3D render sẵn: **AVIF/WebP**, transparent hoặc nền gradient, ≤ 300 KB.
- Fallback: PNG/2D illustration nếu trình duyệt không hỗ trợ.
- Nếu interactive: Spline/Three.js **lazy load**, desktop only, `prefers-reduced-motion`, có nút tắt.
- Ánh sáng mềm, bóng đổ nhẹ; 3D ánh xanh (blue rim light); không photoreal rối.

## Layout showcase

### Hero

```text
[ nav ]
[ trái: eyebrow + H1 + sub + CTA + trust ]
[ phải: 3D map/landmark + 2 floating cards ]
```

- H1 56–72px desktop, 36–44px mobile.
- CTA chính: blue pill (hoặc trắng trên nền xanh); CTA phụ: outline navy/blue.
- Floating cards: route, “3 điểm gần đây”; không bịa số; không card thời tiết (chưa có ở R1).

### Section map/dịch vụ

- Map lớn (như mẫu 2) + panel nổi:
  - “Tuyến của bạn”
  - “Điểm gần đây”
  - “Đối tác”
- Dùng screenshot/map thật khi có; nếu chưa, mock có nhãn “demo”.

### Feature cards

- 3–4 card, icon tile ice/blue, tiêu đề 18–20px, mô tả 14–16px.
- Bo 20–24px, shadow nhẹ, hover nhẹ.

### Pricing (sau R1 — không có trang Pricing ở R1)

- 3 card; gói giữa highlight (viền blue).
- R1: giá “Beta / Liên hệ”; không số bịa.
- Bảng so sánh ngắn; CTA “Liên hệ”.

### Policies

- Trang chữ sạch, max-width 720–780px; mục lục bên trái trên desktop.
- Không 3D, không gradient.

## Components

| Component | Style |
|---|---|
| Button primary | `--blue-strong` fill, white text, pill 999, height 48 |
| Button secondary | Outline navy/blue, transparent |
| Card | White, radius 20–24, border `--line`, shadow mềm |
| Chip | Ice/white, radius 999, icon + text |
| Input | White, border `--line`, focus blue |
| Nav | Trong suốt trên hero, chuyển trắng khi scroll |
| Footer | Navy, chữ trắng, link muted |

## Motion

- Floating card: 6–8s, biên độ 4–8px.
- Parallax hero: nhẹ, không giật.
- Transition: 150–250ms.
- Tôn trọng `prefers-reduced-motion`; không autoplay video có tiếng.

## Accessibility

- Contrast AA (4.5:1 cho text thường).
- Focus ring rõ; không bỏ outline.
- Alt text cho 3D/ảnh; mô tả ngắn.
- Touch target ≥ 48px.
- Không dùng màu làm tín hiệu duy nhất.
- Form có label, error text, `aria`.
- Ngôn ngữ `lang` đúng cho VI/EN (thêm ZH khi mở ZH).

## Performance

- Hero LCP < 2.5s (mục tiêu).
- 3D/ảnh lazy below fold; preload hero.
- Font subset; chỉ load weight cần.
- Không WebGL nặng ở R1.
- Lighthouse mobile ≥ 90 (mục tiêu).

## Nội dung/ảnh

- Chỉ dùng screenshot thật khi có; nếu mock, ghi “demo”.
- Không testimonial/logo đối tác khi chưa xin phép.
- Không hứa tính năng chưa có.
- Credit ảnh/3D nếu dùng nguồn ngoài.

## R1 — phạm vi thiết kế

- Hero 3D tĩnh + 2 floating card.
- Section map (screenshot/map có nhãn).
- 4 feature card (chỉ tính năng R1).
- Trang Download, Contact (email).
- 3 policy page (Privacy, Terms, Map/Data); chỉ typography.
- VI/EN; responsive; accessibility cơ bản.

## Câu hỏi mở

1. Accent phụ: **orange** hay bỏ hẳn? (đề xuất orange rất nhẹ, chỉ badge/CTA phụ)
2. 3D style: isometric sạch hay realistic? (đề xuất isometric sạch, ánh xanh)
3. Có cần dark mode cho showcase không? (đề xuất không; app có dark)
4. Dùng ảnh thật hay illustration cho hero?
5. Có thuê 3D artist/illustrator không, hay dùng asset có license?
6. Font ZH: Noto Sans SC hay system font?
