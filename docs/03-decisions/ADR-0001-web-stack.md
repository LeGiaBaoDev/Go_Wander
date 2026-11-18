# ADR-0001 — Web stack: React + TypeScript, Next.js cho Client Web, Vite cho Admin/Partner

- **Trạng thái:** Accepted
- **Ngày:** 18-11-2026
- **Người quyết định:** chủ dự án
- **Liên quan:** `../02-architecture/01-system-context.md`, `../02-architecture/02-containers.md`, `../00-foundation/02-scope.md`

## Bối cảnh

GoWander có 3 khu vực web:

- **Client Web** — người dùng cuối: landing, địa điểm, cẩm nang, chia sẻ hành trình, bản đồ, plan. Cần SEO/SSR/SSG và preview khi chia sẻ.
- **Admin Web** — vận hành: CRUD địa điểm/guide, upload ảnh, block editor, dashboard, drag/drop, phân quyền.
- **Partner Web** (R6) — dashboard đối tác, khuyến mãi, ads/affiliate/booking; dùng chung component với Admin.

Ràng buộc:

- Backend là **Go API**; web chỉ gọi API, không nhân bản business logic.
- Đội nhỏ; cần một ngôn ngữ/stack web dễ bảo trì, dễ tuyển người.
- Admin có màn tương tác phức tạp; htmx thuần sẽ phải nhúng nhiều JS rời.
- Client Web cần SEO; SPA thuần không đủ.
- R1 cần **Admin Web** + **Showcase Web** (marketing); Client Web app/Partner Web mở sau.

## Quyết định

Dùng **React + TypeScript** cho toàn bộ web, chia theo kiểu render:

| Web | Stack | Lý do |
|---|---|---|
| **Client Web** | **Next.js + TypeScript** (App Router) | Showcase marketing R1; app/SEO/share sau |
| **Admin Web** | **React + TypeScript + Vite** | SPA cho editor/dashboard/drag-drop; không cần SEO |
| **Partner Web** | **React + TypeScript + Vite** | Dùng chung component/table/chart/form với Admin |
| **Shared** | `packages/ui`, `packages/api-client`, `packages/config` | Một design system, một API client, một cấu hình |
| **Backend** | Go API (modular monolith) | Không đổi; không dùng Node backend |

### Stack chi tiết

| Thành phần | Chọn |
|---|---|
| Ngôn ngữ | TypeScript (strict) |
| Client Web | Next.js (App Router) |
| Admin/Partner | Vite + React |
| Router | TanStack Router (Vite) / Next router |
| Data/state | TanStack Query |
| Form | React Hook Form + Zod |
| Table | TanStack Table |
| UI | Tailwind CSS + shadcn/ui (Radix primitives) |
| Editor | TipTap (React) hoặc Editor.js |
| Drag/drop | dnd-kit |
| Chart | Recharts |
| Map | MapLibre GL JS |
| API client | OpenAPI → TypeScript (`openapi-typescript`/Orval), dùng chung |
| Test | Vitest + Testing Library + Playwright |
| Auth | Session cookie httpOnly + CSRF, RBAC; Go API |
| Deploy | Admin/Partner static build; Client Web Next server; cùng origin qua reverse proxy |

### Cấu trúc repo dự kiến

```text
web/
  apps/
    client-web/     Next.js + TypeScript
    admin-web/      Vite + React + TypeScript
    partner-web/    Vite + React + TypeScript
  packages/
    ui/             component + design tokens
    api-client/     TypeScript sinh từ OpenAPI
    config/         eslint, tsconfig, tailwind
```

## Lý do

1. **SEO cho Client Web**: Next.js cho SSR/SSG, metadata, share preview; React SPA thuần không đủ.
2. **Admin phức tạp**: editor, dashboard, drag/drop, nhiều state → React SPA sạch hơn htmx.
3. **Một ngôn ngữ**: TypeScript cho cả 3 web; một đội đọc/sửa được.
4. **Chia sẻ**: `ui` + `api-client` dùng chung giữa Admin/Partner/Client.
5. **Hệ sinh thái**: MapLibre, TipTap/Editor.js, dnd-kit, TanStack, Recharts, shadcn.
6. **Tuyển người**: React/TS phổ biến hơn Vue/Svelte/htmx ở Việt Nam.
7. **Backend độc lập**: Go API vẫn là nơi duy nhất chứa business logic; web thay được mà không đổi API.

## Phương án đã cân nhắc

| Phương án | Vì sao không chọn |
|---|---|
| **Go + Templ + htmx cho tất cả** | Nhanh cho CRUD nhưng yếu với editor/dashboard/drag-drop; phải nhúng nhiều JS rời; khó dùng chung Partner Web |
| **Next.js cho tất cả** | Vẫn tốt, nhưng Admin/Partner không cần SSR; thêm server runtime/deployment phức tạp hơn Vite |
| **Vue/Nuxt** | Tốt, nhưng hệ sinh thái component/map/editor cho React phong phú hơn; tuyển người React dễ hơn |
| **SvelteKit** | Nhẹ/đẹp, nhưng đội chưa chắc quen; hệ sinh thái nhỏ hơn |
| **Astro + React islands** | Rất tốt cho content/SEO, nhưng Client Web có map/planner app-like; Next phù hợp hơn |
| **React SPA cho tất cả** | Client Web thiếu SEO/SSR; không chọn |

## Hệ quả

### Tích cực

- Client Web có SEO/share tốt.
- Admin/Partner có UI phức tạp, dễ test, dễ mở rộng.
- Một design system/API client; giảm lệch giữa các web.
- Không phụ thuộc Node backend; Go API giữ nguyên.

### Tiêu cực / chi phí

- R1 Admin + Showcase chậm hơn so với chỉ dùng htmx/static HTML: build, auth, routing, component library.
- Cần CI cho web: lint, typecheck, test, build.
- Cần kỷ luật component/API client để không phân mảnh.
- Đội phải biết React/TS; nếu chưa, cần học trước.
- Next.js + Vite là **hai kiểu render**, nhưng cùng React/TS; không phải hai ngôn ngữ.

## Ghi chú triển khai

- **Auth:** session cookie httpOnly + SameSite + CSRF; RBAC `admin`/`editor`/`partner`; không lưu token trong localStorage.
- **Same origin:** reverse proxy phục vụ web + `/api` để cookie/CORS đơn giản.
- **API client:** sinh từ OpenAPI; không viết tay DTO.
- **MapLibre:** client component, dynamic import, không SSR.
- **CSP:** không `unsafe-eval`; kiểm tra thư viện editor/map.
- **Design tokens:** chia sẻ màu/spacing/typography với mobile ở mức token JSON; không chia component React ↔ Compose.
- **Test:** Playwright cho luồng Admin; Vitest cho component/logic.

## Việc tiếp theo

- [ ] Tạo `web/` monorepo + package config.
- [ ] Chốt ADR API/auth (ADR-0002).
- [ ] Chốt data model (ADR-0003).
- [ ] Sinh API client từ OpenAPI.
- [ ] Dựng Showcase marketing (Home/Services/Pricing/Policies/Download).
- [ ] Dựng design system tối thiểu (button, input, table, form, dialog).
- [ ] CI web: lint/typecheck/test/build.

## Tham khảo

- Quyết định cũ (repo `../GoWander`): Admin Go templates + htmx; chỉ dùng làm bài học, không mặc nhiên kế thừa.
