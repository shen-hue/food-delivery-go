# 苍穹外卖 Go 版（sky-take-out-go）

这是 [sky-take-out copy](../sky-take-out%20copy)（Java 基础版）的 Go 语言复刻。
与 Java 基础版一样，**只实现了员工登录/退出**，其余功能留给你参照
[food-delivery-j](../food-delivery-j)（Java 完成版）自己动手实现，达到练习 Go 的目的。

## 技术选型（与 Java 版的对应关系）

| Java 版 | Go 版 | 说明 |
|---|---|---|
| Spring Boot 2.7 / Spring MVC | **net/http**（标准库，Go 1.22 路由） | `mux.HandleFunc("POST /path", h)` 对应 `@PostMapping` |
| MyBatis | **pgx/v5**（手写SQL） | 实体 `db` 标签 + `RowToStructByName` 自动扫描，相当于 map-underscore-to-camel-case |
| MySQL | **PostgreSQL** | 你选择了 pgx 驱动；建表脚本见 `sql/schema.sql`（`user` 表因保留字改名为 `users`） |
| jjwt | **golang-jwt/jwt/v5** | 见 `internal/common/utils/jwt.go` |
| Spring 依赖注入 | **手工组装** | 见 `cmd/server/main.go`（构造函数注入） |
| @RestControllerAdvice | **handler.Wrap** | 见 `internal/server/handler/global_exception_handler.go` |
| HandlerInterceptor | **中间件函数** | 见 `internal/server/interceptor/jwt_token_admin_interceptor.go` |
| ThreadLocal BaseContext | **context.Context** | 见 `internal/common/basecontext` |
| 异常类继承体系 | **error 值 + 具名构造函数** | 见 `internal/common/exception`，`exception.NewAccountNotFoundException(msg)` ≈ `throw new AccountNotFoundException(msg)` |
| Lombok @Data | **结构体 + json/db 标签** | json标签与Jackson输出一致；db标签供pgx扫描 |
| AOP @AutoFill 自动填充 | **无（Go没有AOP）** | 练习时在mapper/service里直接设置审计字段，或写个小工具函数 |
| knife4j 接口文档 | 未包含 | 可以后续自行接入 swaggo/swag |
| WeChatPayUtil 微信支付 | 未包含 | 练习支付功能时建议做"模拟支付"即可 |

## 目录结构（与 Java 版模块一一对应）

```
sky-take-out-go/
├── cmd/server/main.go                 # SkyApplication 启动类
├── config/
│   ├── application.yaml               # application.yml
│   └── application-dev.yaml           # application-dev.yml（数据源）
├── sql/schema.sql                     # PostgreSQL建表脚本+种子数据
└── internal/
    ├── common/                        # ≈ sky-common 模块
    │   ├── constant/                  # MessageConstant/StatusConstant等
    │   ├── basecontext/               # BaseContext（ThreadLocal → context.Context）
    │   ├── enumeration/               # OperationType
    │   ├── exception/                 # BaseException及各业务异常构造函数
    │   ├── jsonutil/                  # JSON读写 + DateTime（yyyy-MM-dd HH:mm:ss）
    │   ├── properties/                # Jwt/AliOss/WeChat配置 + 配置加载
    │   ├── result/                    # Result/PageResult
    │   └── utils/                     # JwtUtil/HttpClientUtil/AliOssUtil
    ├── pojo/                          # ≈ sky-pojo 模块
    │   ├── entity/                    # 11个实体
    │   ├── dto/                       # 21个DTO
    │   └── vo/                        # 17个VO
    └── server/                        # ≈ sky-server 模块
        ├── config/router.go           # WebMvcConfiguration（路由+拦截器注册）
        ├── controller/admin/          # 目前只有EmployeeController
        ├── handler/                   # GlobalExceptionHandler
        ├── interceptor/               # JwtTokenAdminInterceptor
        ├── mapper/                    # 目前只有EmployeeMapper
        └── service/                   # EmployeeService + impl
```

## 运行

```bash
# 1. 准备PostgreSQL数据库
createdb sky_take_out
psql -d sky_take_out -f sql/schema.sql

# 2. 按需修改 config/application-dev.yaml 中的数据库连接信息

# 3. 启动（profile默认dev，可通过环境变量 SKY_PROFILE 切换）
go run ./cmd/server
```

## 验证已实现的功能

```bash
# 登录（返回token）
curl -X POST http://localhost:8080/admin/employee/login \
  -H 'Content-Type: application/json' \
  -d '{"username":"admin","password":"123456"}'

# 退出（需要token；不带token会返回401，说明jwt拦截器生效）
curl -X POST http://localhost:8080/admin/employee/logout \
  -H 'token: <上一步返回的token>'
```

注意：基础版密码为**明文比对**（与Java版一致，service里有TODO），
种子数据中密码是 `123456`。实现MD5后记得把 `sql/schema.sql` 中
管理员密码改为 `e10adc3949ba59abbe56e057f20f883e`。

## 练习路线图

参照 `../food-delivery-j` 完成版，建议按以下顺序实现（每一步都遵循
Controller → Service → Mapper 的三层结构）：

1. **员工管理**：新增员工、分页查询（`PageResult`）、启用禁用、编辑、修改密码
   - 练习点：POST/GET/PUT路由、URL查询参数解析、`LIMIT/OFFSET`分页、MD5加密
2. **分类管理**：CRUD、启停、删除保护（分类下有菜品/套餐时不能删）
   - 练习点：事务（`pool.Begin` + `defer rollback`）
3. **菜品管理**：带口味的新增/编辑、批量删除、起售停售
   - 练习点：一对多保存（dish + dish_flavor）、事务
4. **套餐管理**：带菜品关联的CRUD、起售校验
5. **公共接口**：文件上传（`utils.AliOssUtil` 已备好）、店铺营业状态（可用Redis或数据库）
6. **用户端**：微信登录（可先用openid模拟code2session）、地址簿、购物车、菜品/套餐浏览
   - 练习点：`JwtTokenUserInterceptor`（用户端拦截器）、`basecontext` 传递userId
7. **订单**：下单、模拟支付、历史订单、管理端订单状态流转（接单/拒单/派送/完成）、催单提醒
   - 练习点：复杂事务、订单号生成、定时任务（`time.Ticker`代替@Scheduled取消超时订单）
8. **数据统计**：营业额/用户/订单统计、销量Top10、工作台
   - 练习点：`date_trunc`按天分组聚合、字符串拼接（对应Java版逗号拼接VO）

## 常见对照示例

**Java Controller方法 → Go**

```java
@PostMapping("/status/{status}")
public Result<String> startOrStop(@PathVariable Integer status, Long id) { ... }
```

```go
// 需要注册到 adminMux：adminMux.HandleFunc("POST /admin/employee/status/{status}", handler.Wrap(c.StartOrStop))
func (c *EmployeeController) StartOrStop(w http.ResponseWriter, r *http.Request) error {
	status := r.PathValue("status") // Go 1.22 路径参数
	...
}
```

**Java 抛业务异常 → Go**

```java
throw new AccountNotFoundException(MessageConstant.ACCOUNT_NOT_FOUND);
```

```go
return nil, exception.NewAccountNotFoundException(constant.AccountNotFound)
// handler.Wrap 会把它转换成 {"code":0,"msg":"账号不存在","data":null}
```

**MyBatis 查询 → pgx**

```java
@Select("select * from employee where id = #{id}")
Employee getById(Long id);
```

```go
func (m *EmployeeMapper) GetByID(ctx context.Context, id int64) (*entity.Employee, error) {
	rows, err := m.pool.Query(ctx, "select * from employee where id = $1", id)
	if err != nil {
		return nil, err
	}
	employee, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[entity.Employee])
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &employee, nil
}
```
