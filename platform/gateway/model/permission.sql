create table permissions
(
    id          bigint auto_increment comment '主键id',
    name        varchar(255) default '' not null comment '权限名称',
    method      varchar(10)  default '' not null comment '方法',
    url         varchar(255) default '' not null comment '路径',
    description varchar(255) default '' not null comment '描述',
    create_time timestamp               null comment '创建时间',
    update_time timestamp               null comment '更新时间',
    delete_time timestamp               null comment '删除时间',
    primary key(id),
    constraint permissions_uni_name
        unique (name)
)
    comment '权限表';

