### Auth
```sh
GET : / - Vào trang chủ
```
```sh
POST : /api/register - Đăng ký user mới
{
    "name": "",
    "email": "",
    "password": "",
    "confirm": ""
}
```
```sh
POST : /api/login - Đăng nhập
{
    "email":"",
    "password": ""
}
```
### Product
```sh
POST : /api/products - tạo sản phẩm mới
{
    "id": "",
    "name":"",
    "desc":"",
    "price":""
}
```
```sh
GET : /api/products - Lấy tất cả sản phẩm
```
```sh
GET : /api/products/:id - Lấy sản phẩm theo id
```
```sh
PUT : /api/products/:id - Chỉnh sửa sản phẩm theo id
```
```sh
DELETE : /api/products/:id - Xóa sản phẩm theo id
```
### Orders // giỏ hàng
```sh
POST : /api/orders - tạo order mới
{
    user_id: "",
    product_id: "",
    quantity: number
}
```
```sh
GET : /api/orders/:id - Lấy order theo id
```
```sh
PUT : /api/orders/:id - Chỉnh sửa order theo id
```
```sh
DELETE : /api/orders/:id - Xóa order theo id
```
