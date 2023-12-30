# Pharmacy-POS
An Pharmacy POS system by web application based in golang

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
    - 收银页
    - 统计页
    - 库存页
    - 会员页
    - 优惠券页
    - 超级管理页

- 收银页
  - 当有顾客进行购买药品时，收银员通过该页面创建一个订单，并且根据顾客是否是会员、优惠券以及满减等等优惠政策进行计算

- 统计页
  - 收银员也可以查看某一段时间的营业额、某一段时间内某个药品的销售量，药品的数量变化，会员的数量变化，以及优惠券的使用率

- 库存页
  - 用于添加某类药品，也可以删除或修改药品的信息，或者更新药品的库存数量

- 会员页
  - 为收银员提供了创建会员，删除会员，修改会员信息，以及查询会员信息，包括会员积分查询以及优惠券查询的功能

- 优惠券页
  - 可以对优惠券进行创建，提供两个类别的优惠券，比如满100元减20元，或者某类或者某写些指定id的药品打八折，以及优惠券的使用数量与使用率

- 超级管理页
  - 用于对本系统的所有操作进行记录，并记录操作员id，提供添加修改删除用户（收银员）的功能

### 后端API请求设计

#### 登录API

- `POST /login`：用于验证收银员的工号和密码，返回登录令牌（token）.

#### 药品管理API

- `GET /drugs`：获取药品列表.
- `POST /drugs`：添加新的药品信息.
- `PUT /drugs/{id}`：更新特定药品信息.
- `DELETE /drugs/{id}`：删除特定药品信息.
- `GET /drugs/{id}`：获取特定药品的详细信息.

#### 价格管理API

- `GET /prices`：获取所有药品的价格信息.
- `POST /prices`：设置新的药品价格.
- `PUT /prices/{id}`：更新特定药品的价格信息.

#### 优惠券管理API

- `GET /coupons`：获取优惠券列表.
- `POST /coupons`：创建新的优惠券.
- `PUT /coupons/{id}`：更新优惠券信息.
- `DELETE /coupons/{id}`：删除优惠券.

#### 会员管理API

- `GET /members`：获取会员列表.
- `POST /members`：添加新会员.
- `PUT /members/{id}`：更新会员信息.
- `DELETE /members/{id}`：删除会员.
- `GET /members/{id}`：获取特定会员的详细信息.

#### 销售与订单管理API

- `POST /sales`：创建新订单.
- `GET /sales`：获取订单列表.
- `GET /sales/{id}`：获取特定订单的详细信息.
- `PUT /sales/{id}`：更新订单信息（例如，处理退货）.

#### 统计API

- `GET /statistics/sales`：获取销售统计数据.
- `GET /statistics/drugs`：获取药品销售统计数据.
- `GET /statistics/members`：获取会员统计数据.

#### 库存管理API

- `GET /inventory`：获取库存列表.
- `PUT /inventory/{id}`：更新库存数量.

#### 超级管理员API

- `GET /admin/users`：获取用户列表.
- `POST /admin/users`：添加新用户.
- `PUT /admin/users/{id}`：更新用户信息.
- `DELETE /admin/users/{id}`：删除用户.
- `GET /admin/logs`：获取系统操作日志.

### 数据库表单设计

#### 用户表（Users）

- UserID（主键）
- UserName（用户名）
- Password（密码）
- Role（身份）

#### 药品表（Drugs）

- DrugID（主键）
- Name（姓名）
- Category（分类）
- Price（价格）
- Description（描述）
- StockQuantity（剩余数量）

#### 价格表（Prices）

- PriceID（主键）
- DrugID（外键）
- Price（价格）

#### 折扣表（Discounts）

- DiscountID（主键）
- DrugID（外键，关联到药品表）
- DiscountRate（折扣率，0-1之间的小数）
- StartDate（生效时间）
- EndDate（结束时间）

#### 优惠券表（Coupons）

- CouponID（主键）
- Type（例如：满减、折扣）
- DiscountValue（折扣金额）
- MinPurchaseAmount（最低消费金额）
- StartDate（生效日期）
- EndDate（结束日期）

#### 会员表（Members）

- MemberID（主键）
- Name（姓名）
- PhoneNumber（手机号）
- Points（积分）
- JoinDate（加入时间）

#### 订单表（Sales）

- SaleID（主键）
- UserID（外键）
- MemberID（外键）
- TotalAmount（总金额）
- DiscountAmount（折扣金额）
- FinalAmount（实际金额）
- SaleDate（订单日期）

#### 订单详情表（SaleDetails）

- SaleDetailID（主键）
- SaleID（外键）
- DrugID（外键）
- Quantity（数量）
- PricePerUnit（每个的价格）

#### 库存表（Inventory）

- InventoryID（主键）
- DrugID（外键）
- Quantity（数量）
- LastUpdated（上次更新时间）

#### 操作日志表（Logs）

- LogID（主键）
- UserID（外键）
- ActionType（例如：登录、创建、更新、删除）
- Description（描述）
- ActionDate（执行时间）

## 项目结构

