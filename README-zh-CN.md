# 药品管理系统

本系统为中小型药店设计一个药品管理系统

1. 要求实现药品管理功能，包括入库、分类销售管理等功能
2. 价格管理功能，包括定价、优惠券、满减等功能
3. 会员管理功能，包括积分、优惠政策、药品提醒推荐等功能

## 功能设计

- 登录页
  - 使用收银员的工号或其他信息来登录
- 首页
  - 提供了以下页面的入口，并显示欢迎语与时间
    - 收银
    - 订单
    - 药品管理
      - 库存管理
      - 价格管理
      - 分类管理
    - 会员管理
    - 优惠券管理
    - 系统管理（需要管理员权限）
      - 用户管理
      - 系统日志

### 后端API请求设计

#### 登录API

- `POST /login`：用于验证收银员的工号和密码，返回登录令牌（token）

#### 用户API

- `GET /users/getinfo`：根据用户传入的token，获取用户的头像和用户名
- `POST /users/create`：创建新用户，并加密密码后存储
- `DEL /users/delete/:id`：删除对应id的用户
- `GET /users/get/:id`：查询对应id的详细信息
- `GET /users/getall`：获取所有用户信息的列表
- `GET /users/getbyusername`：根据用户名获取用户详细信息
- `PUT /users/update/password/:id`：修改对应id的用户，并将密码加密存储
- `PUT /users/update/isadmin/:id`：修改对应id的用户，重新设置权限

#### 药品API

- `DEL /drugs/delete/:id`：删除对应id的药品
- `POST /drugs/create`：创建新药品
- `GET /drugs/get/:id`：查询对应id的药品详细信息
- `GET /drugs/getall`: 获取所有药品信息的列表
- `PUT /drugs/update`：修改药品

- `DEL /categoties/delete/:id`：删除对应id的分类
- `POST /categoties/create`：创建新分类
- `GET /categoties/get/:id`：查询对应id的分类详细信息
- `GET /categoties/getall`: 获取所有分类信息的列表
- `PUT /categoties/update`：修改分类

