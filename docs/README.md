# Tài liệu GoWander v2

Đây là **nguồn sự thật** cho dự án làm lại. Mọi tài liệu khác (repo cũ, chat, ghi chú) chỉ là tham khảo.

## Bản đồ tài liệu

| Thư mục | Trả lời câu hỏi |
|---|---|
| `00-foundation/` | Sản phẩm là gì, cho ai, phạm vi nào, nguyên tắc nào? |
| `01-product/` | Người dùng, tính năng, luồng UX, ngôn ngữ, accessibility |
| `02-architecture/` | Hệ thống gồm khối nào, dữ liệu nào, API nào, bảo mật ra sao? |
| `03-decisions/` | Quyết định gì đã chốt, vì sao, đánh đổi nào? |
| `04-development/` | Toolchain, quy ước code, test, thứ tự đọc |
| `05-operations/` | Môi trường, runbook, observability |
| `06-research/` | Nguồn dữ liệu, transit, pháp lý |
| `07-changes/` | Thay đổi theo ngày, bằng chứng, việc tiếp theo |
| `templates/` | Mẫu ADR, change record, design doc |
| `evidence/` | Ảnh, log, biên bản kiểm chứng |

## Thứ tự đọc cho người mới

```text
README.md
  → docs/README.md
    → 00-foundation/01-vision.md
    → 00-foundation/02-scope.md
    → 00-foundation/03-principles.md
    → 00-foundation/05-legacy-carryover.md
    → 01-product/04-showcase-web.md
    → 01-product/05-brand-ui-direction.md
    → 01-product/06-saved-and-contributed-places.md
    → 02-architecture/01-system-context.md
    → 02-architecture/02-containers.md
    → 03-decisions/README.md (đọc các ADR đã Accepted)
    → 06-research/README.md
```

Các file sau **chưa viết** (xem “Việc cần viết tiếp”): `01-product/01-users.md`, `02-feature-map.md`, `02-architecture/03-data-model.md`, `04-api-contract.md`, `05-security-privacy.md`, `04-development/*`.

Khi code bắt đầu, `04-development/04-reading-order.md` sẽ mở rộng thành thứ tự đọc code (docs → backend → mobile → test).

## Trạng thái tài liệu

Mỗi tài liệu ghi một trong các trạng thái:

- **Draft** — đang viết, chưa dùng để code.
- **Proposed** — đã đủ để review, chờ chủ dự án chốt.
- **Accepted** — đã chốt; code phải theo.
- **Superseded** — bị thay bởi tài liệu/ADR khác; giữ để lưu vết.
- **Blocked** — chờ dữ liệu/pháp lý/quyết định bên ngoài.

## Quy tắc viết docs

1. Một chủ đề = một file. Không nhân bản nội dung; link tới file gốc.
2. Ghi rõ **nguồn** và **ngày kiểm chứng** cho số liệu.
3. Phân biệt rõ: **fact** (có bằng chứng), **assumption** (giả định), **decision** (đã chốt), **open question** (chưa chốt).
4. Dùng bảng cho so sánh, checklist cho việc cần làm.
5. Tài liệu dài phải có mục lục và “TL;DR” ở đầu.
6. Thay đổi lớn phải có ADR; thay đổi thường ngày có change record.

## Quan hệ với repo cũ

Repo mới: `git@github.com:LeGiaBaoDev/Go_Wander.git`.

Repo cũ: `../GoWander`.

- Được phép đọc để lấy **bài học, số liệu đã đo, bằng chứng**.
- Không copy code/tài liệu mà chưa kiểm chứng lại.
- Khi dùng lại một kết luận, ghi: `Nguồn: ../GoWander/<path> (kiểm chứng DD-MM-YYYY)`.
- Các quyết định của repo cũ **không mặc nhiên** là quyết định của v2; nếu giữ, phải có ADR mới.

## Việc cần viết tiếp

- [ ] Duyệt `00-foundation/01-vision.md` (đã viết).
- [ ] Duyệt `00-foundation/02-scope.md` (đã chốt R1/R2, chờ duyệt).
- [ ] Duyệt `00-foundation/03-principles.md` (đã viết).
- [ ] Duyệt `00-foundation/04-glossary.md` (đã viết).
- [ ] Duyệt `02-architecture/01-system-context.md` (đã viết).
- [ ] Duyệt `02-architecture/02-containers.md` (đã viết).
- [ ] Duyệt `01-product/04-showcase-web.md` (đã viết).
- [ ] Duyệt `01-product/05-brand-ui-direction.md` (đã viết).
- [ ] Duyệt `01-product/06-saved-and-contributed-places.md` (đã viết).
- [ ] Duyệt `00-foundation/05-legacy-carryover.md` (bảng kế thừa từ repo cũ).
- [ ] Làm Figma R0.5: tokens + màn R1 + prototype (trước UI code). Thứ tự gợi ý: Client Mobile (Map, Search, Place, Route, Saved) → Admin Web → Showcase.
- [ ] Viết `01-product/01-users.md`, `02-feature-map.md`.
- [ ] Viết `02-architecture/04-api-contract.md` (kèm `dataStatus`, `UNSUPPORTED`), `05-security-privacy.md`.
- [ ] Viết `02-architecture/03-data-model.md` — chốt entity/ownership.
- [x] `03-decisions/ADR-0001-web-stack.md` — đã chốt.
- [ ] Viết `03-decisions/ADR-0002-api-auth.md` — API/auth.
- [ ] Viết `04-development/01-toolchain.md`, `02-conventions.md`, `03-testing.md`, `04-reading-order.md`.
- [x] `06-research/README.md` — mục lục nghiên cứu kế thừa từ repo cũ (chưa kiểm chứng lại).
- [ ] Kiểm chứng lại license/status nguồn dữ liệu, rồi viết `06-research/01-data-sources.md`.
- [ ] Nhờ luật sư rà checklist pháp lý (bản gốc: `../GoWander/docs/research/legal-compliance-checklist-2026-11-17.md`).
