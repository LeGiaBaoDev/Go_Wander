# 06 — Saved Places & User Contributions

- **Trạng thái:** Draft
- **Ngày:** 18-11-2026
- **Chủ sở hữu:** chủ dự án
- **Liên quan:** `../00-foundation/02-scope.md`, `../02-architecture/02-containers.md`, `../03-decisions/ADR-0001-web-stack.md`

## TL;DR

Hai chức năng, hai mức độ:

1. **Lưu địa điểm** — R1, local-first, guest dùng được, không cần tài khoản.
2. **Thêm địa điểm** — **bắt buộc tài khoản**, đi qua moderation, không publish trực tiếp.

Thêm địa điểm là **R1.5** vì phụ thuộc vào account client (R1 đang guest-first) và cần Admin moderation queue.

**R1.5 chỉ nhận chữ, không nhận ảnh.** Ảnh do người dùng gửi chuyển sang **R1.6**, sau khi qua cổng pháp lý (bản quyền, quyền riêng tư, Nghị định 147). Các phần về ảnh bên dưới là thiết kế cho R1.6, chưa phải phạm vi R1.5. (Quyết định 18-11-2026.)

## Phạm vi địa lý

- **Lưu địa điểm:** toàn quốc; không giới hạn TP.HCM.
- **Đề xuất địa điểm:** toàn quốc; TP.HCM ưu tiên duyệt trong pilot.
- **Routing:** theo coverage; đề xuất place ngoài TP.HCM vẫn có thể duyệt, nhưng app chỉ tìm tuyến khi vùng đó có graph/dữ liệu.
- Không dùng đề xuất place để “mở khoá” routing bus/metro ngoài coverage.

```text
User đã đăng nhập
  → chọn điểm trên map
  → form đề xuất
  → kiểm tra trùng
  → gửi (pending)
  → Admin duyệt / sửa / gộp / từ chối
  → mới hiện trên map
```

## 1. Lưu địa điểm (R1)

### Yêu cầu

- Nút **Lưu** trên place card, place detail, kết quả tìm kiếm.
- Danh sách **Đã lưu** trong Hồ sơ/Discover.
- Lưu **local-first** trên máy; guest vẫn dùng.
- Không public; không ảnh hưởng map.
- Có thể mở place, chỉ đường, thêm vào plan từ danh sách đã lưu.

### Sau này

- R3: account sync, collection (nhóm), sort/filter.
- Không cần server ở R1.

### Dữ liệu local

```text
SavedPlace {
  placeId, name, category, lat, lon, savedAt, note?
}
```

- Lưu bằng Prefs/JSON, giới hạn hợp lý (ví dụ 500 địa điểm); nâng cấp lên SQLite khi vượt giới hạn. Nguồn: `../GoWander/docs/decisions/ADR-036-on-device-storage.md` (repo cũ, đọc 18-11-2026): đã chấp nhận cách này, kèm điều kiện nâng cấp và giới hạn “gỡ app/đổi máy là mất”.
- `placeId` phải xác định được cả khi địa điểm đến từ geocoder mà chưa có trong bảng `places` (chiến lược định danh place: quyết định trong ADR data model).
- Xoá từng địa điểm; clear all trong Settings (nếu cần).
- Không chứa PII.

## 2. Thêm địa điểm (R1.5, cần tài khoản)

### Vì sao cần tài khoản

- Chống spam/trùng lặp/phá hoại.
- Có người chịu trách nhiệm nội dung.
- Liên hệ khi cần xác minh, từ chối, hoặc xoá.
- Theo dõi trạng thái submission.
- Hạn chế rate limit theo account.

**Không cho guest gửi địa điểm.** Nếu chưa đăng nhập, hiện prompt:

> “Đăng nhập để đề xuất địa điểm. Điều này giúp chúng tôi chống spam và liên hệ khi cần xác minh.”

### Điều kiện tài khoản

- Đã đăng nhập client account.
- **Email/password hoặc Google Sign-In**; cả hai đều dùng được.
- Email phải **xác minh** (Google Sign-In được coi là đã xác minh).
- Account không bị suspend.
- **Rate limit: 3 submission/ngày/account** (cấu hình; admin có thể giảm/nâng).
- Account mới có thể bị giới hạn thấp hơn; admin có thể nâng.

### Flow người dùng

1. Đang ở Map/Discover → **“Đề xuất địa điểm”**.
2. Nếu chưa đăng nhập → sign-in/register.
3. Chọn vị trí:
   - chạm bản đồ, hoặc
   - tìm địa chỉ/place hiện có.
4. Điền form.
5. *(R1.6)* Đính kèm **tối đa 3 ảnh** (optional) — ảnh cũng phải chờ duyệt. R1.5 bỏ bước này.
6. App kiểm tra trùng sơ bộ (cùng category, gần, tên tương tự) → cảnh báo “có thể đã tồn tại”.
7. Gửi → `pending`.
8. User thấy **“Đang chờ duyệt”** trong “Đề xuất của tôi”.
9. Có thể sửa/xoá khi còn `pending`.
10. Khi được duyệt → thấy trạng thái trong “Đề xuất của tôi” (thông báo đẩy/email chưa nằm trong kiến trúc R1.5; xem câu hỏi mở); place xuất hiện với nhãn nguồn (ảnh từ R1.6).
11. Khi bị từ chối → thấy lý do; có thể sửa và gửi lại (nếu policy cho phép).

### Form tối thiểu

| Trường | Bắt buộc | Ghi chú |
|---|---|---|
| Tên địa điểm | ✅ | 2–120 ký tự |
| Category | ✅ | Chọn từ danh mục |
| Vị trí | ✅ | Chọn trên map; lưu lat/lon |
| Địa chỉ/mô tả vị trí | ✅ | Text; có thể dùng reverse geocode gợi ý |
| Mô tả ngắn | ✅ | 20–500 ký tự |
| Giờ mở cửa | ⏳ | Text có cấu trúc sau |
| Điện thoại/website | ⏳ | Validate nhẹ |
| Ảnh | R1.6 | **Tối đa 3 ảnh**, optional; JPEG/PNG/WebP ≤ 5 MB; chờ duyệt; cần license/consent. Không có ở R1.5 |
| Ghi chú cho admin | ⏳ | Không public |
| Đồng ý điều khoản đóng góp | ✅ | Checkbox, lưu version |

### Trạng thái

```text
pending → approved
        → changes_requested (admin yêu cầu sửa) → pending (user sửa và gửi lại)
        → rejected (reason)
        → merged (vào place có sẵn)
        → withdrawn (user xoá)
```

- `pending`: user sửa/xoá được.
- `changes_requested`: admin ghi yêu cầu; user sửa rồi gửi lại thành `pending`.
- `approved`: tạo place mới hoặc cập nhật place (nếu admin chọn).
- `rejected`: lý do; có thể gửi lại sau cooldown.
- `merged`: gộp vào place khác; user thấy place đích.
- `withdrawn`: user tự rút.

### Admin moderation

Queue trong Admin Web:

- Danh sách `pending`, sort theo thời gian/ưu tiên.
- Map preview + toạ độ.
- Địa điểm gần nhất + gợi ý trùng (ST_DWithin ~50 m + tên tương tự).
- Lịch sử contributor: số submission, approval rate, account age, bị reject trước đó.
- Hành động:
  - **Approve**: tạo place mới (source `user`, verification `unverified`).
  - **Approve & merge**: gộp vào place có sẵn; cập nhật field nếu cần.
  - **Reject**: lý do (spam, trùng, sai vị trí, không đủ thông tin, vi phạm).
  - **Request changes**: gửi lại user (nếu có notification).
  - **Suspend contributor**: nếu spam/abuse.
- Audit log: ai, khi nào, hành động, lý do.
- SLA mục tiêu: 48 giờ làm việc (chưa cam kết công khai).

### Nguồn & độ tin cậy

| Field | Giá trị |
|---|---|
| `source` | `user` |
| `verification` | `unverified` mặc định; `verified` sau khi admin kiểm tra |
| `contributor` | Account ID (không public); credit tên nếu user chọn |
| `publishedAt` | Thời điểm approve |
| `license` | Theo điều khoản đóng góp của GoWander |

Không tự động đẩy lên OSM. Nếu sau này có flow OSM, phải riêng, có consent và theo ODbL.

### Chống spam/trùng lặp

- Bắt buộc account + email verified/Google.
- Rate limit theo account/IP.
- CAPTCHA nhẹ ở form (nếu cần).
- Duplicate check:
  - PostGIS `ST_DWithin` ~50 m.
  - So sánh tên (trigram) + category.
  - Hiện candidate cho user/admin.
- Admin có thể merge thay vì tạo mới.
- Từ chối các submission có link/quảng cáo/quá ngắn.
- Có report/remove flow cho place sai.

### Dữ liệu dự kiến

```text
place_submissions
  id, account_id, name, category_id,
  location geography(Point,4326),
  address, description, hours, phone, website,
  status, rejection_reason, admin_note,
  created_at, updated_at, reviewed_by, reviewed_at,
  published_place_id

places
  ... existing ...
  source (osm/admin/user/partner), verification,
  contributor_account_id nullable, published_at

moderation_actions
  id, submission_id, actor_id, action, reason, created_at

place_submission_photos   -- R1.6
  id, submission_id, url, thumb_url, width, height,
  credit_name, license, status (pending/approved/rejected),
  exif_stripped, created_at

place_photos
  id, place_id, submission_photo_id, url, credit_name,
  license, approved_at
```

### API dự kiến

```text
# Người dùng — mọi đường dẫn dưới /me (owner-only)
POST   /api/v1/me/place-submissions                 (auth, verified)
GET    /api/v1/me/place-submissions
GET    /api/v1/me/place-submissions/{id}
PATCH  /api/v1/me/place-submissions/{id}            (pending | changes_requested)
DELETE /api/v1/me/place-submissions/{id}            (pending | changes_requested)

# R1.6 — ảnh
POST   /api/v1/me/place-submissions/{id}/photos     (multipart, max 3)
DELETE /api/v1/me/place-submissions/{id}/photos/{photoId}

# Admin
GET    /api/v1/admin/place-submissions?status=pending
POST   /api/v1/admin/place-submissions/{id}/approve
POST   /api/v1/admin/place-submissions/{id}/reject
POST   /api/v1/admin/place-submissions/{id}/merge
POST   /api/v1/admin/place-submissions/{id}/request-changes
POST   /api/v1/admin/accounts/{id}/suspend
# R1.6: POST /api/v1/admin/place-submissions/{id}/photos/{photoId}/(approve|reject)
```

Lỗi:

- `AUTH_REQUIRED`
- `EMAIL_NOT_VERIFIED`
- `ACCOUNT_SUSPENDED`
- `RATE_LIMITED`
- `DUPLICATE_SUSPECTED` (cảnh báo, không nhất thiết chặn)
- `SUBMISSION_NOT_EDITABLE`

### Legal/privacy

- **Điều khoản đóng góp:** user xác nhận thông tin đúng, có quyền với nội dung/ảnh; cấp GoWander quyền không độc quyền để dùng/hiển thị/chỉnh sửa/phân phối.
- **Không auto OSM:** dữ liệu user lưu ở DB GoWander; nếu muốn đóng góp OSM, flow riêng + ODbL.
- **PII:** email/account ID không public; **tên contributor (display name) hiện công khai** trên place/photo đã duyệt; user có thể đổi display name.
- **Account deletion:** contributions đã publish có thể giữ nhưng ẩn danh; policy ghi rõ.
- **Ảnh (R1.6, sau cổng pháp lý; checklist cũ xếp “ảnh người dùng” vào nhóm chặn mở tính năng):** user xác nhận có quyền; license cho GoWander; tối đa 3 ảnh; admin duyệt từng ảnh; **không ảnh cận mặt/biển số** nếu chưa consent; strip EXIF; có thể xoá ảnh khi place bị gỡ.
- **Map disclaimer:** place user có thể sai; báo cáo/chỉnh sửa.
- **Nội dung cấm:** spam, quảng cáo trá hình, thông tin sai, xúc phạm, vi phạm pháp luật.

### Metrics

| Metric | Mục tiêu sơ bộ |
|---|---|
| Submission/tuần | 5–20 trong pilot (giả thuyết) |
| Approval rate | ≥ 60% |
| Duplicate rate | ≤ 20% |
| Time-to-review | ≤ 48h |
| Active contributors | ≥ 5 trong pilot |
| Photos/submission (R1.6) | ≥ 1 ảnh trung bình (nếu user chọn) |
| Photo rejection rate (R1.6) | ≤ 20% |
| Saved places/user | ≥ 3 |

## Phase

| Phase | Chức năng |
|---|---|
| **R1** | Lưu địa điểm local; guest |
| **R1.5** | Client account (email/password + Google) + đề xuất địa điểm (chỉ chữ) + Admin queue |
| **R1.6** | Ảnh đóng góp (1–3 ảnh) sau cổng pháp lý |
| **R2** | Wander Mode |
| **R3** | Saved places sync + collection |
| **R4** | Community comments (text-only), reputation, public profile |
| **Sau** | Ảnh trong comment, partner verification, OSM export |

### Bình luận & ảnh bình luận

- **R4 launch:** comment **text-only**, cần account, rate limit, report/moderation.
- **Chưa có ảnh trong comment ở R4 launch.** Ảnh comment làm tăng khối lượng/chi phí kiểm duyệt; để **R4.1** sau khi có tool moderation và policy ảnh.
- Khi bật ảnh comment: cùng quy tắc ảnh place (quyền, license, 3 ảnh? 1 ảnh, kích thước, strip EXIF, duyệt).
- Không cho comment ẩn danh; display name công khai; email/ID private.

> R1.5 là phase mới, cần cập nhật `02-scope.md` và roadmap. Vì R1 guest-first, **không thể có add-place trước khi có account**.

## Đã chốt trong phiên 18-11-2026

- Account client: **email/password + Google**; email phải verify; Google đủ.
- Rate limit: **3 submission/ngày/account**.
- Contributor: **hiện display name** trên place/ảnh đã duyệt; email/ID private.
- Ảnh khi đề xuất địa điểm: tối đa 3 ảnh, optional, chờ duyệt — **dời sang R1.6** (quyết định 18-11-2026, sau cổng pháp lý).
- Ảnh trong comment: **chưa có ở R4 launch**; text-only; ảnh comment để R4.1.
- Account deletion: place đã duyệt có thể giữ nhưng ẩn danh; user có thể yêu cầu gỡ.
- Request changes: **có**; admin có thể yêu cầu sửa thay vì reject.
- Moderator: chủ dự án + role Admin/Editor.
- OSM: **không auto**; flow riêng sau.

## Câu hỏi mở

1. Ảnh comment: 1 ảnh hay tối đa 3? (đề xuất 1; R4.1)
2. Thông báo cho contributor (push/email) cần ở R1.5 không? Kiến trúc hiện chưa có hạ tầng push/email cho client.
3. Có cần “Đề xuất của tôi” trên web, hay chỉ app?
4. Thời hạn lưu ảnh bị reject? (đề xuất xoá sau 30 ngày)
5. Có cần email thông báo khi duyệt/từ chối không?
