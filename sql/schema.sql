-- ============================================================
-- 苍穹外卖 Go 版数据库脚本（PostgreSQL）
-- 对应 Java 版 MySQL 数据库 sky_take_out
--
-- 使用方式：
--   createdb sky_take_out
--   psql -d sky_take_out -f sql/schema.sql
--
-- 注意：Java 版中的 `user` 表在 PostgreSQL 中是保留字，
--       这里改名为 `users`。
-- ============================================================

-- 员工表
create table employee (
    id          bigserial primary key,
    username    varchar(50)  not null unique,
    name        varchar(50)  not null,
    password    varchar(100) not null,
    phone       varchar(20)  not null,
    sex         varchar(2)   not null,
    id_number   varchar(30)  not null,
    status      int          not null default 1,             -- 0禁用 1启用
    create_time timestamp    not null default now(),
    update_time timestamp    not null default now(),
    create_user bigint,
    update_user bigint
);

-- 分类表（type: 1菜品分类 2套餐分类）
create table category (
    id          bigserial primary key,
    type        int         not null,
    name        varchar(32) not null,
    sort        int         not null default 0,
    status      int         not null default 1,              -- 0禁用 1启用
    create_time timestamp   not null default now(),
    update_time timestamp   not null default now(),
    create_user bigint,
    update_user bigint
);

-- 菜品表
create table dish (
    id          bigserial primary key,
    name        varchar(32)  not null,
    category_id bigint       not null,
    price       numeric(10,2) not null,
    image       varchar(255),
    description varchar(255),
    status      int          not null default 1,             -- 0停售 1起售
    create_time timestamp    not null default now(),
    update_time timestamp    not null default now(),
    create_user bigint,
    update_user bigint
);

-- 菜品口味表（value 为口味JSON字符串，如 '["微辣","中辣"]'）
create table dish_flavor (
    id      bigserial primary key,
    dish_id bigint       not null,
    name    varchar(32),
    value   varchar(255)
);

-- 套餐表
create table setmeal (
    id          bigserial primary key,
    category_id bigint       not null,
    name        varchar(32)  not null,
    price       numeric(10,2) not null,
    status      int          not null default 1,             -- 0停用 1启用
    description varchar(255),
    image       varchar(255),
    create_time timestamp    not null default now(),
    update_time timestamp    not null default now(),
    create_user bigint,
    update_user bigint
);

-- 套餐菜品关系表
create table setmeal_dish (
    id        bigserial primary key,
    setmeal_id bigint       not null,
    dish_id   bigint       not null,
    name      varchar(32),
    price     numeric(10,2),
    copies    int
);

-- 用户表（微信用户）
create table users (
    id          bigserial primary key,
    openid      varchar(64) not null unique,
    name        varchar(32),
    phone       varchar(20),
    sex         varchar(2),
    id_number   varchar(30),
    avatar      varchar(255),
    create_time timestamp   not null default now()
);

-- 地址簿表
create table address_book (
    id            bigserial primary key,
    user_id       bigint       not null,
    consignee     varchar(32)  not null,
    phone         varchar(20)  not null,
    sex           varchar(2)   not null,                     -- '0'女 '1'男
    province_code varchar(12),
    province_name varchar(32),
    city_code     varchar(12),
    city_name     varchar(32),
    district_code varchar(12),
    district_name varchar(32),
    detail        varchar(255) not null,
    label         varchar(16),
    is_default    int          not null default 0            -- 0否 1是
);

-- 购物车表
create table shopping_cart (
    id          bigserial primary key,
    name        varchar(32),
    user_id     bigint       not null,
    dish_id     bigint,
    setmeal_id  bigint,
    dish_flavor varchar(255),
    number      int          not null default 1,
    amount      numeric(10,2) not null,
    image       varchar(255),
    create_time timestamp    not null default now()
);

-- 订单表
create table orders (
    id                     bigserial primary key,
    number                 varchar(64)   not null unique,    -- 订单号
    status                 int           not null default 1, -- 1待付款 2待接单 3已接单 4派送中 5已完成 6已取消
    user_id                bigint        not null,
    address_book_id        bigint        not null,
    order_time             timestamp     not null default now(),
    checkout_time          timestamp,
    pay_method             int,                              -- 1微信 2支付宝
    pay_status             int           not null default 0, -- 0未支付 1已支付 2退款
    amount                 numeric(10,2) not null,
    remark                 varchar(255),
    user_name              varchar(32)   not null,
    phone                  varchar(20)   not null,
    address                varchar(255)  not null,
    consignee              varchar(32)   not null,
    cancel_reason          varchar(255),
    rejection_reason       varchar(255),
    cancel_time            timestamp,
    estimated_delivery_time timestamp,
    delivery_status        int,                              -- 1立即送出 0选择具体时间
    delivery_time          timestamp,
    pack_amount            int           not null default 0, -- 打包费
    tableware_number       int           not null default 1, -- 餐具数量
    tableware_status       int                                -- 1按餐量提供 0选择具体数量
);

-- 订单明细表
create table order_detail (
    id          bigserial primary key,
    name        varchar(32),
    order_id    bigint        not null,
    dish_id     bigint,
    setmeal_id  bigint,
    dish_flavor varchar(255),
    number      int           not null default 1,
    amount      numeric(10,2) not null,
    image       varchar(255)
);

-- ============================================================
-- 初始数据：管理员账号
-- 注意：基础版代码中密码校验为明文比对（对应 Java 版的 TODO），
--       所以这里先存明文密码 123456。
--       当你实现 MD5 加密比对后，请将其改为：
--       e10adc3949ba59abbe56e057f20f883e
-- ============================================================
insert into employee (username, name, password, phone, sex, id_number, status)
values ('admin', '管理员', '123456', '13812312312', '1', '110101199001010047', 1);
