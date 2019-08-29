# Api Document
## 目錄
1. [玩家登入](#玩家登入)

## 共用參數
### Errors
#### Response 
code | description | data content
---- | --- | ---
1 | 必要參數 | ["branch_id"] 
2 | 登入失敗 |  
3 | 參數不合法 | ["branch_id", ...] 
4 | Token 過期 |
5 | Token 不合法 |
6 | 新增失敗：紀錄已存在 |
7 | 查無紀錄 |
8 | 權限不足 |
9 | 時間區間不合法 |

#### Example
```json
{
    "code": 1,
    "data": [
        "branch_id"
    ]
}
```

### 分頁
#### Response 
name | type | description
---- | --- | --- 
current_page | int | 當前頁數
last_page | int | 最後頁數
per_page | int | 每頁最多比數
total | int |  總比數
from | int | 本頁從第 n 筆開始
to | int | 本頁到第 n 筆結束

#### Example
```json
{
    "current_page": 1,
    "first_page_url": "http:\/\/localhost\/api\/admin\/admin-user\/list?page=1",
    "from": 1,
    "last_page": 1,
    "last_page_url": "http:\/\/localhost\/api\/admin\/admin-user\/list?page=1",
    "next_page_url": null,
    "path": "http:\/\/localhost\/api\/admin\/admin-user\/list",
    "per_page": 20,
    "prev_page_url": null,
    "to": 3,
    "total": 3
}
```    


1. ## 玩家登入

    ### Request
    ```
    POST /api/v1/user/login
    ```
    
    #### Params
    name | type | description | is required | 
    ---- | --- | --- | ---
    username | string | | V
    password | string | | V
    
    ### Response
    name | type | description
    ---- | --- | --- 
    token | string | 
    
    #### Example
    
    ```json
    {
        "token": "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpYXQiOjE1NjM0MzAwMzgsIm5iZiI6MTU2MzQzMDAzOCwiZXhwIjoxNTYzODYyMDM4LCJ1c2VybmFtZSI6ImFkbWluIiwicGFzc3dvcmQiOiIzMjMxMWQ2ZTVmNzI1MTBhOWQ0MjNhYzU5NzBjZDI4MyIsImJyYW5jaF9pZCI6MX0.wEOGBR8kMW4aVE0EzJbnSWBdo6qbMteqYp0VvvY9VYI"
    }
    ```
    
    #### test token: 
    
    ```
    GET /api/v1/user/auth-test
    ```

    #### Errors
    
    code | description | data content
    ---- | --- | ---
    1 | 必要參數 | ```  ["branch_id"] ```
    2 | 登入失敗 |  