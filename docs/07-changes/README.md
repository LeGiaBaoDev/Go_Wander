# Change records

Mỗi thay đổi trong repo — code, docs, config, schema, hạ tầng, dữ liệu — phải có một mục trong thư mục này.

## Quy ước

- **Một file mỗi ngày:** `change-DD-MM-YYYY.md` theo ngày ở Asia/Ho_Chi_Minh.
- Nếu file của ngày đã có, **thêm mục mới** ở cuối; không sửa mục cũ.
- Ngôn ngữ: tiếng Việt; identifier kỹ thuật giữ tiếng Anh.
- Không ghi “đã test” nếu chưa chạy; ghi “chưa kiểm chứng” khi chưa có bằng chứng.
- Mỗi mục trả lời: **thay đổi gì, ở đâu, làm gì, vì sao, bằng chứng, ảnh hưởng/rủi ro, việc tiếp theo**.

## Mẫu

```markdown
## N. Tiêu đề ngắn

- **Milestone / nhánh:** ...
- **Người / agent thực hiện:** ...
- **Thay đổi gì / ở đâu:** ...
- **Làm gì:** ...
- **Vì sao:** ...
- **Bằng chứng / kiểm thử:** ...
- **Ảnh hưởng / rủi ro:** ...
- **Việc tiếp theo:** ...
```

## Liên quan

- `../03-decisions/` — quyết định kiến trúc (ADR).
- `../templates/` — mẫu tài liệu.
- `../04-development/` — quy ước Git/commit (sẽ viết).
