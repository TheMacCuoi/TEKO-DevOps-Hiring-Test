# Rate Limiting Core Function

## **Mô tả**
Chương trình thực hiện chức năng giới hạn tần suất request (**Rate Limiting**) để đảm bảo rằng một client không gửi quá nhiều request vượt mức cho phép trong một khoảng thời gian nhất định. Chương trình sử dụng thuật toán **Sliding Window Log** để xác định xem mỗi request có được chấp nhận hay từ chối.

---

## **Hướng dẫn cài đặt**

### **Yêu cầu hệ thống**
- Ngôn ngữ lập trình: **Go (Golang)**

### **Cài đặt**
1. **Cài đặt Go**:
   - Nếu chưa có Go, tải và cài đặt tại: [https://golang.org/dl/](https://golang.org/dl/).
   - Phiên bản được sử dụng trong bài này là Go 1.22.5
2. **Cài đặt các thành phần cần thiết**:
   ```bash
   go mod tidy
   ```

---

## **Hướng dẫn sử dụng**

### **Chạy chương trình**

- **Cách 1:** 
   - **Bước 1:** Nhập các giá trị đầu vào vào file ![test_input](./test_input.txt) theo định dạng:
      - **Dòng 1**: `N R`
         - `N`: Tổng số request cần kiểm tra.
         - `R`: Giới hạn số request cho phép trong 1 giờ (rate limit).
      - **Dòng 2 đến N+1**: Timestamps của các request, theo định dạng ISO-8601 (`YYYY-MM-DDTHH:mm:ssZ`).
   - **Bước 2:** Chạy lệnh 
      ```bash
      go run main.go test_input.txt
      ```

- **Cách 2:** 
   - **Bước 1:** Chạy lệnh 
      ```bash
      go run main.go
      ``` 
   - **Bước 2:** Nhập các giá trị đầu vào theo định dạng:
   - **Dòng 1**: `N R`
      - `N`: Tổng số request cần kiểm tra.
      - `R`: Giới hạn số request cho phép trong 1 giờ (rate limit).
   - **Dòng 2 đến N+1**: Timestamps của các request, theo định dạng ISO-8601 (`YYYY-MM-DDTHH:mm:ssZ`).

### **Output**
- Chương trình trả về **true/false** cho từng request:
  - **true**: Request được chấp nhận (trong giới hạn).
  - **false**: Request bị từ chối (vượt giới hạn).

---

### **Input/Output mẫu**

**Input**:
```
10 3
2022-01-20T00:13:05Z
2022-01-20T00:27:31Z
2022-01-20T00:45:27Z
2022-01-20T00:49:00Z
2022-01-20T01:15:45Z
2022-01-20T01:20:01Z
2022-01-20T01:50:09Z
2022-01-20T01:52:15Z
2022-01-20T01:54:00Z
2022-01-20T02:00:00Z
```

**Output**:
```
true
true
true
false
true
false
true
true
false
false
```

---

## **Cách hoạt động**

Chương trình sử dụng thuật toán **Sliding Window Log** để kiểm tra rate limit:
1. **Xử lý từng request**:
   - Tìm tất cả các request trong khung thời gian 1 giờ trước thời điểm request hiện tại.
   - Loại bỏ các request đã hết hạn (quá 1 giờ).
2. **Kiểm tra giới hạn**:
   - Nếu số lượng request còn lại nhỏ hơn giới hạn `R`, request được chấp nhận (**true**).
   - Nếu vượt quá giới hạn, request bị từ chối (**false**).
3. **Cập nhật trạng thái**:
   - Lưu lại timestamp của request hiện tại (nếu được chấp nhận).

---

## **Cấu trúc file**
```plaintext
project-folder/
├── main.go        # File chính chứa code xử lý
├── README.md      # Tài liệu hướng dẫn
└── test_input.txt # File chứa input mẫu
```

---

## **Ghi chú**
- **Định dạng timestamp**: Sử dụng chuẩn ISO-8601.

