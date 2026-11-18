# 04 — Thứ tự đọc trước khi code

- **Trạng thái:** Proposed
- **Ngày:** 18-11-2026
- **Chủ sở hữu:** chủ dự án
- **Liên quan:** `../README.md`, `../01-product/07-screen-map.html`

## TL;DR

Đọc **docs trước, code sau**. Người mới (hoặc agent) mất khoảng một buổi đọc là bắt đầu được; đừng đoán thay bằng việc đọc.

## Thứ tự

1. `AGENTS.md` — quy tắc làm việc, Git, commit (thắng mọi mặc định).
2. `docs/00-foundation/` — mục tiêu, phạm vi R1, 12 nguyên tắc, thuật ngữ.
3. `docs/01-product/07-screen-map.html` — **mục lục mọi màn hình** trên canvas thiết kế, cụm màn, luồng chính. Mở file này trước khi đọc từng màn.
4. `docs/01-product/` — thương hiệu (`05`), tài khoản và vai trò (`08`), câu chuyện/khám phá/tìm kiếm (`09`).
5. `docs/02-architecture/` — bối cảnh hệ thống, container, (sau này) dữ liệu, API, bảo mật.
6. `docs/03-decisions/` — đọc các ADR **Accepted** trước; ADR **Proposed** là đề xuất, chưa phải luật.
7. `docs/04-development/` — toolchain, quy ước, test, Git.
8. Chỉ khi cần: `docs/06-research/` (dữ liệu, pháp lý, transit), `docs/07-changes/` (lịch sử và lý do).

## Đọc theo việc

| Việc | Đọc thêm |
|---|---|
| Tìm đường | `02-conventions.md` mục Thuật toán, ADR-0005 (khi có), canvas `RouteBranches`/`RouteLogic` |
| Tài khoản, quyền | ADR-0002, `01-product/08-accounts-roles-partners.md` |
| Câu chuyện, khám phá, tìm kiếm | `01-product/09-stories-explore-search.md` |
| Màn hình / UI | `01-product/05-brand-ui-direction.md`, `07-design-canvas.md`, canvas `GoWander Screens` |
| Đối tác, đặt bàn, cọc | `01-product/08-accounts-roles-partners.md`, `06-research/README.md` (pháp lý, phí) |
| Dữ liệu giả định | `00-foundation/02-scope.md` (nhãn tin cậy), `00-foundation/05-legacy-carryover.md` |

## Giao diện

Nguồn thiết kế UI: canvas **GoWander Screens** <https://claude.ai/artifact/Wt2LzogXbLm7msPKHBaUjY> (mục lục: `07-screen-map.html`). Code UI bám canvas.

## Repo cũ (`../GoWander`)

Code cũ chỉ để **đọc tham khảo**; v2 viết lại. Chủ dự án cho phép **dùng lại dữ liệu và API (Vietmap, …)** từ repo cũ (18-11-2026), nhưng mỗi mục phải ghi vào `00-foundation/05-legacy-carryover.md` trước khi dùng; quyết định cũ không mặc nhiên kế thừa. Khoá API không chép vào git (`01-toolchain.md`).

## Khi mâu thuẫn

Thứ tự ưu tiên: `AGENTS.md` → ADR Accepted → tài liệu Accepted → tài liệu Proposed/Draft → canvas → chat. Gặp mâu thuẫn thì **dừng và hỏi chủ dự án**, đừng tự chốt; sửa tài liệu cho nhất quán thay vì để hai nơi nói khác nhau.
