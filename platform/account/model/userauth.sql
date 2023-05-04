create table users_auths
(
    id              bigint auto_increment comment '主键id',
    user_id         bigint       default 0  not null comment '用户id',
    auth_type       tinyint      default 0  not null comment '认证类型(0-平台)',
    auth_key        varchar(100) default '' not null comment '认证标识',
    auth_credential char(60)     default '' not null comment '认证凭据',
    create_time     timestamp               null comment '创建时间',
    update_time     timestamp               null comment '更新时间',
    delete_time     timestamp               null comment '删除时间',
    primary key(id),
    constraint users_auths_auth_key
        unique (auth_key),
    constraint users_auths_uni_user_id_auth_key
        unique (user_id, auth_key)
)
    comment '用户认证信息表';

