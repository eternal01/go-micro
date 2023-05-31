create table users
(
    id          bigint       auto_increment comment '主键id',
    user_name        varchar(255) default '' not null comment '用户名',
    nick_name        varchar(255) default '' not null comment '用户昵称',
    avatar      varchar(255) default '' not null comment '头像',
    gender      tinyint      default 0  not null comment '用户性别;0-保密;1-男;2-女',
    mobile      varchar(20)  default '' not null comment '用户手机号',
    email       varchar(50)  default '' not null comment '用户邮箱地址',
    status      tinyint      default 0  not null comment '状态(1-启用2-禁用)',
    created_at     timestamp               null comment '创建时间',
    updated_at     timestamp               null comment '更新时间',
    deleted_at     timestamp               null comment '删除时间',
    primary key(id),
    constraint users_uni_email
        unique (email),
    constraint users_uni_mobile
        unique (mobile)
)
    comment '用户表';
