# 01 — Toolchain

- **Trạng thái:** Proposed
- **Ngày:** 18-11-2026
- **Chủ sở hữu:** chủ dự án
- **Liên quan:** `02-conventions.md`, `03-testing.md`, `../02-architecture/02-containers.md`, `../03-decisions/ADR-0001-web-stack.md`

## TL;DR

Chọn công cụ theo stack đã chốt (ADR-0001, `02-containers.md`). Phiên bản **chưa khoá**: ghi `ASSUMED` cho đến khi dựng khung dự án và chạy được trên máy thật; khi đó khoá bằng file cấu hình (`go.mod`, `package.json` + lockfile, `libs.versions.toml`) và cập nhật bảng dưới kèm ngày kiểm chứng.

## Bảng công cụ

| Khối | Ngôn ngữ / framework | Công cụ chính | Trạng thái phiên bản |
|---|---|---|---|
| Backend (Go API) | Go, modular monolith | `chi` (router), `pgx` (Postgres), `sqlc` hoặc truy vấn tay có test, `golangci-lint`, `goose` hoặc tương đương cho migration | ASSUMED, chưa khoá |
| Client Mobile | Kotlin Multiplatform + Compose | MapLibre, Ktor client, Gradle version catalog, `ktlint`/`detekt` | ASSUMED, chưa khoá |
| Admin Web, Partner Web | React + TypeScript + Vite | TanStack Router/Query/Table, React Hook Form + Zod, Tailwind + shadcn/ui, ESLint, Vitest, Playwright | ASSUMED, chưa khoá |
| Client Web / Showcase | Next.js (App Router) + TypeScript | Static export ở R1; cùng bộ lint/test với web khác | ASSUMED, chưa khoá |
| API client | OpenAPI → TypeScript | `openapi-typescript` hoặc Orval; sinh từ spec, không viết tay | ASSUMED |
| Dữ liệu | PostgreSQL + PostGIS | Migration có số thứ tự, seed `DEMO` tách khỏi dữ liệu thật | ASSUMED |
| Tìm tuyến | OpenTripPlanner | Graph TP.HCM (OSM + GTFS có nhãn nguồn/giả định) | Chờ ADR-0005 |
| Geocoding, bản đồ | Vietmap API, Photon, OpenFreeMap | Khoá API giữ ở server, không đưa vào client hay git | Dùng lại từ repo cũ, xem dưới |
| Hạ tầng dev | Docker Compose | Một lệnh dựng API + DB + OTP + proxy | ASSUMED |

> Không thêm thư viện nằm ngoài bảng khi chưa có lý do ghi trong PR (và ADR nếu ảnh hưởng nhiều khối). “Đã có sẵn” không phải lý do.

## Tái dùng từ repo cũ (`../GoWander`)

Chủ dự án quyết định ngày 18-11-2026: **được dùng lại dữ liệu đã làm và các API (Vietmap, …) của repo cũ.**

- Trước khi dùng một mục, ghi nó vào `../00-foundation/05-legacy-carryover.md` (giữ ý, đổi, hay bỏ) kèm nguồn, giấy phép, ngày kiểm chứng. Quyết định của repo cũ không mặc nhiên kế thừa.
- **Code cũ chỉ để tham chiếu;** v2 viết lại, không sao chép nguyên khối.
- Dữ liệu cũ vẫn mang nhãn tin cậy (`ASSUMED`/`DEMO`/có nguồn) như ở `02-conventions.md`.
- **Khoá API** (Vietmap, …) lấy từ biến môi trường; không chép từ repo cũ vào git. Mẫu ở `.env.example`, không giá trị thật.
- Kiểm tra điều khoản dùng của từng nhà cung cấp (giới hạn gọi, lưu cache, ghi công) trước khi bật ở pilot.

## Giao diện (UI)

Nguồn thiết kế giao diện là canvas **GoWander Screens**: <https://claude.ai/artifact/Wt2LzogXbLm7msPKHBaUjY>. Mục lục mọi màn: `../01-product/07-screen-map.html`; quy ước: `../01-product/07-design-canvas.md`. Code UI bám canvas và token palette v4 (`../01-product/05-brand-ui-direction.md`); lệch thiết kế phải ghi lý do.

## Máy dev

- Một lệnh để chạy được toàn bộ (`make dev` hoặc `docker compose up`); tên lệnh chốt khi dựng khung. Ai clone về phải chạy được mà không hỏi ai.
- Biến môi trường mẫu ở `.env.example` (không giá trị thật). **Không commit** `.env`, khoá, dữ liệu thô, DB volume, backup (AGENTS.md mục 5).
- Seed mặc định là dữ liệu `DEMO`; dữ liệu thật nạp bằng lệnh riêng, ghi nguồn và giấy phép.

## Kiểm tra cục bộ trước khi đẩy

Mỗi khối có đúng một lệnh “xanh hay đỏ” (tên lệnh chốt khi dựng khung):

| Khối | Phải chạy sạch |
|---|---|
| Backend | format, lint, test, build, kiểm tra migration lên/xuống trên DB trống |
| Web | typecheck, lint, test, build |
| Mobile | lint, test đơn vị, build debug |
| OpenAPI | lint spec, sinh lại client không có diff chưa commit |

CI chạy lại đúng các lệnh này; PR đỏ thì không merge.

## Việc chưa chốt

- Chọn `sqlc` hay truy vấn tay; công cụ migration cụ thể.
- Monorepo task runner (npm workspaces / pnpm / Turborepo) cho `web/`.
- Cách chạy OTP trong dev (container riêng, graph nhỏ để test nhanh).
- Khoá phiên bản thực tế sau lần dựng khung đầu tiên.
- Danh sách dữ liệu và API cụ thể của repo cũ được mang sang (làm ở `05-legacy-carryover.md`).
