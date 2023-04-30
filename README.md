# Insured Tree

此專案用於展示一個保戶的二叉樹結構，提供以下兩個API：

1. 生成二元樹資料結構，以輸入的 `id` 為根節點，最多顯示四層15個用戶：
   ```
   [GET] /api/user/tree/{id}
   ```
2. 用來搜尋用戶，帶入 `id` 即可搜尋用戶：
   ```
   [GET] /api/user/{id}
   ```

## 使用方法

1. 複製專案：

   ```bash
   git clone https://github.com/TomasTangisCoding/insured-tree.git
   ```

2. 進入專案目錄：

   ```bash
   cd insured-tree
   ```

3. 根據DB環境設定 `config.yaml` 中的 `dsn` 參數

4. 創建對應的資料庫表格，可以參考 `user.sql` 中的SQL語句

5. 執行服務：

   ```bash
   go run main.go
   ```

6. 在瀏覽器訪問以下URL來使用API
    - 生成用戶樹：`http://localhost:8080/api/user/tree/{id}`
    - 搜尋用戶：`http://localhost:8080/api/user/{id}`