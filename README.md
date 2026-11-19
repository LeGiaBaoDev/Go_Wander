# GoWander v2 — rebuild docs-first

Ứng dụng di chuyển và khám phá cho TP.HCM: tìm tuyến **bus / metro / đi bộ / xe đạp**, khám phá địa điểm, lập kế hoạch, dẫn đường và quản trị nội dung.

> **Repo:** `git@github.com:LeGiaBaoDev/Go_Wander.git`
>
> **Trạng thái:** khởi tạo lại từ đầu. Giai đoạn hiện tại là **docs-first**: chốt lại vì sao, phạm vi, kiến trúc và quy ước trước khi viết code.
>
> Repo cũ `../GoWander` được giữ làm **nguồn tham chiếu read-only** (code, bằng chứng, bài học). Không copy máy móc; chỉ chép những gì đã được kiểm chứng và ghi rõ nguồn.

## Vì sao có folder này

Repo cũ đã đi qua nhiều milestone, nhiều quyết định và nhiều thay đổi liên tiếp. Tài liệu bị phân tán, một số kết luận đã bị thay thế, một số con số là mock/ASSUMED. Folder mới cho phép:

- viết lại **một nguồn sự thật** cho sản phẩm;
- tách rõ **đã kiểm chứng** / **giả định** / **chưa làm**;
- thiết kế lại kiến trúc trước khi cam kết code;
- giữ lịch sử thay đổi ngay từ ngày đầu.

## Cấu trúc hiện tại

```text
GoWander-v2/
  README.md
  AGENTS.md
  .env.example
  .gitignore
  api/                 # openapi.yaml (nguồn sự thật hợp đồng)
  backend/             # Go API: cmd/api, cmd/migrate, internal/<module>, migrations/
  mobile/              # KMP + Compose: shared/ (Android + iOS), androidApp/
  infra/               # docker-compose.yml (Postgres + PostGIS)
  .github/workflows/   # CI: backend, mobile, OpenAPI
  docs/
    00-foundation/     # vì sao, phạm vi, nguyên tắc, glossary
    01-product/        # người dùng, tính năng, UX
    02-architecture/   # hệ thống, container, dữ liệu, API, bảo mật
    03-decisions/      # ADR
    04-development/    # toolchain, quy ước, test, thứ tự đọc
    05-operations/     # môi trường, runbook
    06-research/       # nguồn dữ liệu, transit, pháp lý
    07-changes/        # change record theo ngày
    templates/         # ADR, change record, design doc
    evidence/          # ảnh/log/biên bản kiểm chứng
```

> Khung code R1 đã có: backend health API + migration bootstrap, mobile Compose rỗng, compose Postgres/PostGIS, CI. **Chưa có tính năng sản phẩm** (place/route/geocode) — các slice đó làm sau.

## Cách chạy local

Yêu cầu: Go 1.27+, JDK 21, Android SDK (API 37 + build-tools 36), Docker Desktop (cho DB), Xcode 27 (tuỳ chọn iOS). Phiên bản đã kiểm 19-11-2026 ở [`docs/04-development/01-toolchain.md`](./docs/04-development/01-toolchain.md).

```bash
cp .env.example .env
docker compose -f infra/docker-compose.yml up -d

cd backend
go run ./cmd/migrate up
go run ./cmd/api
```

Kiểm tra API:

```bash
curl http://localhost:8080/health/live   # 200 {"status":"ok"}
curl http://localhost:8080/health/ready  # 200 nếu DB sẵn sàng; 503 nếu chưa
```

Backend:

```bash
cd backend
go test ./...
go vet ./...
golangci-lint run ./...
```

Mobile:

```bash
cd mobile
./gradlew :androidApp:assembleDebug
./gradlew ktlintCheck detekt
./gradlew :shared:linkDebugFrameworkIosSimulatorArm64
```

APK ở `mobile/androidApp/build/outputs/apk/debug/androidApp-debug.apk`. Build Android local cần `mobile/local.properties` với `sdk.dir=...` (đã gitignore) hoặc biến `ANDROID_HOME`.

OpenAPI:

```bash
npx --yes @redocly/cli@2.80.0 lint api/openapi.yaml
```

> 19-11-2026: Docker Desktop đã khởi động lại; PostGIS 17-3.5 chạy qua `linux/amd64` (Rosetta) vì image chính thức chưa có arm64. Migration up/down/up và `/health/ready` với DB đã chạy. Android đã chạy trên emulator API 37.0; iOS mới link framework, chưa có app Xcode.

## Đọc trước

1. [`docs/README.md`](./docs/README.md) — bản đồ tài liệu và thứ tự đọc.
2. [`docs/00-foundation/01-vision.md`](./docs/00-foundation/01-vision.md) — sản phẩm giải quyết vấn đề gì.
3. [`docs/00-foundation/02-scope.md`](./docs/00-foundation/02-scope.md) — pilot làm gì, không làm gì.
4. [`docs/02-architecture/01-system-context.md`](./docs/02-architecture/01-system-context.md) — hệ thống gồm những khối nào.

## Quy ước làm việc

- **Docs trước, code sau.** Một tính năng chưa có docs/ADR thì chưa code.
- **Không bịa dữ liệu.** Dữ liệu giả định phải ghi `ASSUMED`/`DEMO`; không được hiển thị như dữ liệu chính thức.
- **Mọi thay đổi có change record** trong `docs/07-changes/`.
- **Quyết định lớn có ADR** trong `docs/03-decisions/`.
- **Không commit secret**, `.env`, dữ liệu thô, DB volume, backup.
- **Tiếng Việt** là ngôn ngữ tài liệu; identifier kỹ thuật giữ tiếng Anh.

## Việc tiếp theo

- [ ] Duyệt vision/scope/principles/glossary.
- [ ] Duyệt system context + containers.
- [x] ADR-0001 web stack: React/Next/Vite.
- [x] ADR-0002 Accepted cho R1 (ba loại tài khoản; Admin cấp tài khoản quán, Chủ quán tự CRUD nhân viên); ADR-0003 API/auth/RBAC + ADR-0004 data/PostGIS + ADR-0005 tìm đường vẫn Proposed.
- [ ] R1.5: client account + place submissions (account required, chỉ chữ); R1.6: ảnh đóng góp sau cổng pháp lý.
- [ ] Duyệt `docs/00-foundation/05-legacy-carryover.md` (kế thừa từ repo cũ).
- [ ] R0.5: Figma design tokens + R1 screens + prototype before UI code.
- [ ] Duyệt Showcase Web + brand + saved/contributed places (`docs/01-product/04`, `05`, `06`).
- [x] Chốt API contract v0 (`docs/02-architecture/04-api-contract.md` + `api/openapi.yaml`).
- [ ] Chốt data source inventory + trạng thái pháp lý.
- [x] Dựng repo code + CI + môi trường local (migration up/down/up đã chạy 19-11-2026; CI chưa chạy trên GitHub).
