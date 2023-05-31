create table users_auths
(
    id              bigint auto_increment comment '主键id',
    user_id         bigint       default 0  not null comment '用户id',
    auth_type       tinyint      default 0  not null comment '认证类型(1-手机密码;2-手机验证码;3-邮箱密码)',
    auth_key        varchar(100) default '' not null comment '认证标识',
    auth_credential char(60)     default '' not null comment '认证凭据',
    created_at     timestamp               null comment '创建时间',
    updated_at     timestamp               null comment '更新时间',
    deleted_at     timestamp               null comment '删除时间',
    primary key(id),
    constraint users_auths_uni_user_id_auth_key_auth_type
        unique (user_id, auth_key, auth_type)
)
    comment '用户认证信息表';

