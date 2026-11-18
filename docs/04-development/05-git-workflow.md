# 05 — Quy trình Git, PR và review

- **Trạng thái:** Proposed
- **Ngày:** 18-11-2026
- **Chủ sở hữu:** chủ dự án
- **Liên quan:** `../../AGENTS.md` (nguồn sự thật; file này chỉ giải thích và bổ sung)

> Nếu file này mâu thuẫn với `AGENTS.md`, **`AGENTS.md` thắng.**

## TL;DR

- **Một task = một nhánh = một PR.** Nhánh gốc là `develop`; `main` chỉ nhận merge từ `develop` khi phát hành.
- **Chủ dự án ghi lịch sử Git** (`add`, `commit`, `push`, merge, tag). Agent chỉ sửa file và tạo/chuyển nhánh cục bộ.
- Commit theo Conventional Commits, **không** ghi công AI trong message.

## Vòng đời một task

1. `git fetch`, kiểm tra working tree sạch, tạo nhánh từ `origin/develop`:
   ```bash
   git switch -c <type>/<scope>-<mô-tả-ngắn> origin/develop
   ```
2. Làm nhỏ, theo `02-conventions.md`; docs/ADR trước khi code nếu là tính năng hoặc kiến trúc mới.
3. Chạy lệnh kiểm tra (`01-toolchain.md`), ghi bằng chứng.
4. Thêm mục vào `docs/07-changes/` (một file theo ngày, thêm mục ở cuối, không sửa mục cũ).
5. Chủ dự án xem diff, `add` + `commit` + `push`, mở PR vào `develop`.
6. Review, sửa theo góp ý bằng commit mới, merge, xoá nhánh.

Task phát sinh giữa chừng (ngoài phạm vi) → ghi lại và hỏi, **không** làm chung nhánh. Nhánh dài ngày bị cấm (nguyên tắc 12).

## Đặt tên nhánh

```text
<type>/<scope>-<mô-tả-ngắn>
```

- Chữ thường, kebab-case, tiếng Anh, không dấu, tối đa khoảng 50 ký tự.
- `type`: `feat`, `fix`, `docs`, `refactor`, `test`, `chore`, `ci`, `build`, `perf`.
- `scope`: `foundation`, `product`, `architecture`, `adr`, `research`, `changes`, `repo`, `mobile`, `backend`, `admin-web`, `showcase`, `data`, `infra`, `api`.
- Ví dụ đúng: `feat/mobile-saved-places`, `fix/backend-route-outside-coverage`, `docs/development-guidelines`.

## Commit message

```text
<type>(<scope>): <subject>

<body — tuỳ chọn>
```

- Tiếng Anh, thức mệnh lệnh, chữ thường ở đầu, không dấu chấm cuối, subject ≤ 72 ký tự.
- Body giải thích **vì sao**, ngắt dòng ≤ 100 ký tự.
- Phá vỡ tương thích: `feat(api)!: …` và dòng `BREAKING CHANGE:` trong body.
- Một commit = một thay đổi hợp lý; không trộn tính năng, refactor và format.
- Tham chiếu ADR/change record trong body khi liên quan (`Refs: ADR-0002`).
- **Cấm:** `Co-Authored-By` trỏ tới AI, “Generated with…”, emoji ghi công AI, email/đường dẫn máy/token trong message.

### Danh tính tác giả

Dự án dùng một danh tính tác giả thống nhất (tài khoản GitHub của chủ dự án). Đặt theo repo, **không** đặt toàn máy nếu bạn dùng nhiều tài khoản:

```bash
git config --local user.name  "<tên tài khoản GitHub>"
git config --local user.email "<mail đã thêm vào tài khoản đó>"
```

Kiểm tra trước khi push: `git log -1 --format='%an <%ae>'`. Agent không đổi `user.name`/`user.email` và không dùng `--author`.

## Pull request

- Tên PR giống subject của commit chính (Conventional Commits). PR trỏ vào `develop`.
- Mô tả có: **làm gì**, **vì sao**, **cách kiểm chứng** (lệnh và kết quả), **chưa kiểm chứng**, **rủi ro**, link docs/ADR/change record.
- PR nhỏ, một mục đích. PR lớn trộn nhiều việc sẽ bị yêu cầu tách.
- Không có ghi công AI trong mô tả PR.

## Review

Người review kiểm theo thứ tự:

1. **Đúng phạm vi:** đúng task, không lẫn việc khác.
2. **Docs/ADR:** hành vi mới có docs; quyết định lớn có ADR; change record đủ.
3. **Đúng quy ước:** `02-conventions.md` (lớp, lỗi, quyền, nhãn dữ liệu, trợ năng).
4. **An toàn:** kiểm tra đầu vào, quyền ở server, không lộ PII/bí mật.
5. **Test và bằng chứng:** có test cho logic mới; bằng chứng thật, phần chưa kiểm chứng được nói rõ.
6. **Đơn giản:** có cách ngắn hơn mà vẫn đủ không? Có phần nào không ai yêu cầu không?

## Điều không làm (trừ khi chủ dự án ra lệnh rõ)

`--force` (nếu cần sửa nhánh cá nhân chưa ai dùng: `--force-with-lease`), `--no-verify`, `git reset`, `git clean`, xoá nhánh remote, sửa lịch sử đã merge, tạo tag release. “OK” cho một việc không mở rộng sang việc khác: được commit không có nghĩa được push; được push không có nghĩa được mở PR hay merge.
