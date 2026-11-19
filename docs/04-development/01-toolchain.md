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
| Backend (Go API) | Go 1.27.4, modular monolith | `chi` v5.5.0 · `pgx` v5.13.0 · `goose` v3.30.1 · `golangci-lint` v2.15.0 | Đã khoá trong `backend/go.mod` + `backend/.golangci.yml`; kiểm chứng 19-11-2026 |
| Client Mobile | Kotlin Multiplatform + Compose | Kotlin 2.4.21 · Compose Multiplatform 1.12.2 · AGP 9.5.1 · Gradle 9.7.1 · Ktor 3.6.1 · JDK 21 · `ktlint` 1.8.0 · `detekt` 1.23.8 | Đã khoá trong `mobile/gradle/libs.versions.toml` + `gradle-wrapper.properties`; Android và iOS framework đã build 19-11-2026 |
| Admin Web, Partner Web | React + TypeScript + Vite | TanStack Router/Query/Table, React Hook Form + Zod, Tailwind + shadcn/ui, ESLint, Vitest, Playwright | ASSUMED, chưa khoá (chưa dựng web) |
| Client Web / Showcase | Next.js (App Router) + TypeScript | Static export ở R1; cùng bộ lint/test với web khác | ASSUMED, chưa khoá (chưa dựng web) |
| API client | OpenAPI 3.1 → TypeScript | `openapi-typescript` hoặc Orval; lint bằng Redocly CLI 2.80.0 | Lint đã chạy 19-11-2026; client TS/Kotlin chưa sinh |
| Dữ liệu | PostgreSQL 17 + PostGIS 3.5 | Image `postgis/postgis:17-3.5`; migration goose 0001 (extension + schema + role) | Ảnh khoá trong `infra/docker-compose.yml`; migration up/down/up đã chạy 19-11-2026 |
| Tìm tuyến | OpenTripPlanner | Graph TP.HCM (OSM + GTFS có nhãn nguồn/giả định) | Chờ ADR-0005; chưa dựng container OTP |
| Geocoding, bản đồ | Vietmap API, Photon, OpenFreeMap | Khoá API giữ ở server, không đưa vào client hay git | Dùng lại từ repo cũ, xem dưới |
| Hạ tầng dev | Docker Compose v5.3.1 | `infra/docker-compose.yml` (Postgres + PostGIS; chưa có OTP) | `docker compose up -d` đã chạy 19-11-2026; PostGIS dùng `linux/amd64` (Rosetta) |

> Không thêm thư viện nằm ngoài bảng khi chưa có lý do ghi trong PR (và ADR nếu ảnh hưởng nhiều khối). “Đã có sẵn” không phải lý do.

### Bằng chứng phiên bản (19-11-2026)

| Lệnh | Kết quả |
|---|---|
| `go version` | `go1.27.4 darwin/arm64` |
| `cd backend && go test ./...` | pass; `internal/httpapi` ok |
| `cd backend && go vet ./...` | không lỗi |
| `cd backend && go build ./...` | ok |
| `cd backend && golangci-lint run ./...` | `0 issues` (v2.15.0) |
| `cd mobile && ./gradlew --version` | Gradle 9.7.1; JVM 21.0.8 |
| `cd mobile && ./gradlew :androidApp:assembleDebug` | BUILD SUCCESSFUL; có APK debug |
| `adb install` + `am start` trên emulator API 37.0 | app chạy, hiện “GoWander” (`docs/evidence/scaffold-android-19-11-2026.png`) |
| `cd mobile && ./gradlew ktlintCheck detekt` | BUILD SUCCESSFUL |
| `cd mobile && ./gradlew :shared:linkDebugFrameworkIosSimulatorArm64` | BUILD SUCCESSFUL; framework tĩnh |
| `npx @redocly/cli@2.80.0 lint api/openapi.yaml` | valid; 3 warnings |
| `docker compose -f infra/docker-compose.yml config` | hợp lệ |
| `docker compose -f infra/docker-compose.yml up -d` | PostGIS 17-3.5 healthy (linux/amd64/Rosetta) |
| `cd backend && go run ./cmd/migrate up` | OK 0001; `down` OK; `up` lại OK |
| `cd backend && go run ./cmd/migrate status` | 0001 applied |
| `psql` qua container | 6 schema `gowander_owner`; role `gowander_owner`/`gowander_app` NOLOGIN; extension `postgis`/`pg_trgm`/`citext` |
| `GET /health/ready` với DB | 200 `{"checks":{"database":"ok"},"status":"ok"}` |

- iOS: mới link framework; **chưa** build/run app Xcode, chưa thử simulator.
- CI: file đã tạo, **chưa** chạy trên GitHub.

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

- Lệnh chạy local (19-11-2026): `cp .env.example .env`; `docker compose -f infra/docker-compose.yml up -d`; `cd backend && go run ./cmd/migrate up && go run ./cmd/api`; mobile `cd mobile && ./gradlew :androidApp:assembleDebug`. Chi tiết ở `README.md`.
- Docker Desktop đã restart ngày 19-11-2026; image `postgis/postgis:17-3.5` chỉ có `linux/amd64`, nên compose khai báo `platform: linux/amd64` (Rosetta trên Apple Silicon).
- Biến môi trường mẫu ở `.env.example` (không giá trị thật). **Không commit** `.env`, khoá, dữ liệu thô, DB volume, backup (AGENTS.md mục 5).
- Seed mặc định là dữ liệu `DEMO`; dữ liệu thật nạp bằng lệnh riêng, ghi nguồn và giấy phép.

## Kiểm tra cục bộ trước khi đẩy

Mỗi khối có đúng một lệnh “xanh hay đỏ” (tên lệnh chốt khi dựng khung):

| Khối | Phải chạy sạch |
|---|---|
| Backend | `gofmt -l .` rỗng, `go vet ./...`, `golangci-lint run ./...`, `go test ./...`, `go build ./...`; migration up/down trên DB trống (đã chạy 19-11-2026) |
| Web | chưa có code |
| Mobile | `./gradlew ktlintCheck detekt :androidApp:assembleDebug`; unit test chưa có |
| OpenAPI | `npx @redocly/cli@2.80.0 lint api/openapi.yaml`; sinh client chưa có |

CI chạy lại đúng các lệnh này; PR đỏ thì không merge.

## Việc chưa chốt

- Chọn `sqlc` hay truy vấn tay (scaffold dùng `pgx` trực tiếp); migration đã chọn `goose`.
- Monorepo task runner (npm workspaces / pnpm / Turborepo) cho `web/`.
- Cách chạy OTP trong dev (container riêng, graph nhỏ để test nhanh).
- iOS app target + Xcode project; hiện chỉ link framework từ `shared`.
- Web scaffold (Admin/Showcase) và sinh API client.
- Chạy migration trên DB trống/CI; Docker Desktop trên máy dev đang lỗi.
- Danh sách dữ liệu và API cụ thể của repo cũ được mang sang (làm ở `05-legacy-carryover.md`).
