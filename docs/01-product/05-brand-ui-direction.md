# 05 — Brand & UI Direction

- **Trạng thái:** Proposed (palette v4, chờ chủ dự án duyệt trên canvas)
- **Ngày:** 18-11-2026
- **Chủ sở hữu:** chủ dự án
- **Liên quan:** `07-design-canvas.md`, `04-showcase-web.md`, `../03-decisions/ADR-0001-web-stack.md`, `../00-foundation/02-scope.md` (R0.5)

## TL;DR

- **Nền trắng là chính**, chữ gần đen `#0F1E24`, **một màu nhấn duy nhất: teal `#19C4B0`**.
- Nút chính: nền teal, **chữ tối** (7.77:1). **Không bao giờ chữ trắng trên teal** (2.2:1).
- Link/chữ teal dùng `#08766B` (5.51:1 trên trắng).
- Màu khác chỉ để báo trạng thái (thành công, cảnh báo, lỗi, thông tin). Màu tuyến bus/metro chỉ dùng trên chip tuyến.
- Linh vật **Gô** (robot hộp kiểu xe buýt mini) xuất hiện ở trạng thái trống, onboarding, loa đọc chỉ dẫn.
- Có **chế độ tối** cho app (Sáng / Tối / Theo máy).
- Thiết kế mọi phase (R1→R7) ngay ở R0.5 để phase sau không phải vẽ lại; xem `07-design-canvas.md`.

> **Lịch sử:** bản đầu ngày 18-11-2026 dùng “Sexy Blue” (`#007BFF`, `#0062E0`, cam nhấn). Chủ dự án đổi sang palette v4 teal trong cùng ngày sau nhiều vòng thử trên canvas. Các token blue cũ không còn dùng.

## Palette v4

Tỉ lệ tương phản tính bằng công thức WCAG 2.x (script Python, 18-11-2026).

| Token | Hex | Dùng cho | Tương phản |
|---|---|---|---|
| `--ink` | `#0F1E24` | Chữ chính, heading, chữ trên nút teal, nền thanh điều hướng tối | 17.05 trên trắng; 7.77 trên teal |
| `--body` | `#374549` | Chữ thân | 9.95 trên trắng |
| `--muted` | `#5F6E73` | Chữ phụ, nhãn nhỏ | 5.30 trên trắng; 4.83 trên `--bg` |
| `--accent` | `#19C4B0` | Nền nút chính, tab đang chọn, đường tuyến, chấm vị trí | — (không đặt chữ trắng: 2.2) |
| `--accent-text` | `#08766B` | Link, chữ teal | 5.51 trên trắng |
| `--accent-soft` | `#E3F9F5` | Nền chip, icon tile, hàng đang chọn | — |
| `--bg` | `#F2F5F4` | Nền trang web/admin | — |
| `--surface` | `#FFFFFF` | Card, nền app | — |
| `--line` | `#E6EAEC` / `#D5DCDE` | Viền, đường chia / viền input | — |
| `--success` | `#0F766E` trên `#E3FAEA` | Trạng thái tốt | 4.99 |
| `--warn` | `#7A4E0E` trên `#FFF4D6` | Cảnh báo, “Giờ ước tính” | 6.55 |
| `--error` | `#B42318` trên `#FDECEC` | Lỗi, xoá | 5.76 |
| `--info` | `#0B5F8F` trên `#E3F1FA` | Thông tin | 5.98 |

### Chế độ tối (app)

| Token | Hex | Tương phản |
|---|---|---|
| `--bg-dark` | `#0B1417` | — |
| `--surface-dark` | `#15232A` | — |
| `--ink-dark` | `#EAF2F3` | 16.41 trên `--bg-dark` |
| `--muted-dark` | `#9FB1B5` | 7.22 trên `--surface-dark` |
| `--accent-text-dark` | `#5EDBCB` | 9.55 trên `--surface-dark` |

- Nút chính vẫn là teal + chữ `--ink`. Màu tuyến bus/metro giữ nguyên ở cả hai chế độ.
- Người dùng chọn **Sáng / Tối / Theo máy** trong Hồ sơ → Giao diện; tuỳ chọn riêng “Bản đồ tối khi trời tối”.

### Quy tắc

- Một màu nhấn; không thêm màu brand thứ hai.
- Mọi cặp chữ/nền phải kiểm bằng số: ≥ 4.5:1 chữ thường, ≥ 3:1 chữ lớn và thành phần UI.
- Màu không phải tín hiệu duy nhất; luôn có chữ hoặc icon.
- Gradient chỉ nhẹ, trên nút teal hoặc lớp phủ ảnh để chữ đọc được; không gradient cả trang.

## Typography

| Ngôn ngữ | Font | Weight |
|---|---|---|
| VI/EN | **Be Vietnam Pro** | 400/500/600/700/800/900 |
| ZH (sau R1) | **Noto Sans SC** | 400/500/700 |
| Fallback | system-ui, -apple-system, “Segoe UI”, sans-serif | — |

- Heading 800/900, tracking `-0.02em`. Body 14–16px, line-height ~1.5. Không dưới 11px; nhãn 11px chỉ cho tag in đậm.
- Số liệu dùng `font-variant-numeric: tabular-nums`.

## Logo, linh vật, minh hoạ

- **Logo:** giữ logo GoWander cũ (ghim teal-xanh lá có xe buýt); file đã tải lên canvas.
- **Gô:** robot hộp như xe buýt mini (mặt pixel trên màn kem, bảng LED “GÔ 01”, cửa sổ xe ở hai bên, tuyến có trạm trên thân, bánh xe làm chân, ghim bản đồ trên đầu). Có các trạng thái (mood) và cỡ (component `Chip`).
  - Ảnh tham khảo của chủ dự án là một nhân vật hoạt hình có sẵn; **Gô phải khác rõ** để tránh vấn đề bản quyền.
- Minh hoạ phẳng, bo tròn, bóng mềm; không 3D nặng trong app. Showcase có thể dùng hero 3D render sẵn (AVIF/WebP ≤ 300 KB).

## Ảnh và bản quyền

- Ảnh thật phải có **tác giả, giấy phép, link nguồn**; app hiện dòng ghi công dưới ảnh/video.
- Bản vẽ R0.5 dùng ảnh Wikimedia Commons (CC BY-SA 3.0/4.0, CC BY 4.0, CC0, phạm vi công cộng); danh sách và tác giả nằm trong ghi chú `photoCredit` trên canvas.
- Admin không đăng được câu chuyện khi một media còn thiếu giấy phép (xem `09-stories-explore-search.md`).
- Ảnh do quán tải lên: quán chịu trách nhiệm bản quyền (điều khoản đối tác).
- Không testimonial/logo đối tác khi chưa xin phép; mock ghi “DEMO”.

## Layout và responsive

| Bề mặt | Khung thiết kế | Ghi chú |
|---|---|---|
| App điện thoại | 390 × 844 | Thanh điều hướng dưới **chỉ icon** (xem “Điều hướng”); chế độ quán có thanh tối riêng 4 mục |
| App iPad | 1194 × 834 (ngang) | Thanh điều hướng dọc + 2 cột (danh sách/nội dung + bản đồ/ảnh) |
| Web (Admin, Partner, Showcase, đăng ký hợp tác) | 1440 rộng, fluid | Bản iPad 834 và điện thoại 390 phải dùng được |

- Breakpoint web: **≤ 760px** (Partner) / **≤ 700px** (Admin): thanh bên chuyển thành thanh ngang trên cùng; bảng rộng cuộn trong hộp; form một cột.
- Khung responsive trên canvas là **khung thường, không viền thiết bị** (chủ dự án không muốn khung bao).
- Touch target ≥ 44px (app ≥ 48px khi có thể).

## Điều hướng và lựa chọn dài

- **App (điện thoại):** thanh dưới dạng viên thuốc, **chỉ icon**: Bản đồ, Khám phá, nút Tìm tuyến nổi ở giữa, Kế hoạch, Đã lưu, **⋯**. Mục đang mở có nền `--accent-soft`. Tên mục vẫn có cho trình đọc màn hình (`aria-label` + chữ ẩn). Bấm **⋯** → **xổ ra bảng điều hướng** có tên đầy đủ, nhóm Chính / Khám phá thêm (Câu chuyện, Khám phá Việt Nam, Tìm theo loại) / Của tôi (Thông báo, Đặt bàn của tôi, Quán của tôi, Hồ sơ). Màn: `NavMenu`.
  - *Đánh đổi:* icon không chữ khó hiểu hơn với người mới; onboarding nên giới thiệu thanh này một lần.
- **Admin:** thanh bên 5 nhóm: Tổng quan, Tài chính, Nội dung ▸, Partner ▸, Hệ thống ▸; mục con mở bằng flyout. Không để thanh bên dài quá ~6 mục.
- **Partner web:** thanh bên chia nhóm (Vận hành / Thiết lập / Tiền) + ô chọn quán.
- **Nhiều lựa chọn (> ~5) → dropdown, nhóm, hoặc “Thêm ▾”**; không để chip/nút kéo dài nhiều hàng.

## Components

| Component | Style |
|---|---|
| Button primary | Nền `--accent`, chữ `--ink`, pill, cao 44–54 |
| Button secondary | Viền `--line`, nền trắng, chữ `--body` |
| Button danger | Nền `#B42318`, chữ trắng |
| Card | Trắng, bo 20–28, bóng mềm |
| Chip / tag | Pill; trạng thái dùng cặp màu trạng thái ở trên |
| Input / select | Trắng, viền `--line`, cao ≥ 44, có `<label>` |
| Sheet (app) | Bo trên 28–30, tay kéo |
| Toggle | Teal khi bật, xám khi tắt |

## Motion

- Transition 150–250ms; tôn trọng `prefers-reduced-motion`; không autoplay video có tiếng.

## Accessibility

- Contrast AA kiểm bằng số; focus ring rõ.
- Nút chỉ có icon phải có `aria-label`; SVG trang trí `aria-hidden`.
- Form có label, lỗi bằng chữ, `aria-invalid`.
- `lang` đúng cho VI/EN.
- Loa đọc chỉ dẫn (Gô đọc) có 3 mức: Tắt / Cảnh báo / Đầy đủ.

## Performance

- Hero LCP < 2.5s (mục tiêu); ảnh lazy dưới màn đầu; font subset.
- Không WebGL nặng ở R1.

## Câu hỏi mở

1. Chủ dự án duyệt palette v4 + Gô trên canvas (gate R0.5).
2. Cập nhật Design System artifact sang v4 (hiện vẫn token blue cũ).
3. Bộ icon thống nhất: tự vẽ stroke 2px hay dùng bộ có license (ví dụ Lucide, MIT)?
4. Font ZH khi mở ZH.
