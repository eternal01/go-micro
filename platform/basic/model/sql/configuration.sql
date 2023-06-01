create table system_configuration
(
    id          int auto_increment comment '主键id',
    name        varchar(50)   default '' not null comment '名称',
    description varchar(100)  default '' not null comment '描述',
    content     varchar(8000) default '' not null comment '内容',
    primary key(id),
    constraint system_configuration_uni_name
        unique (name)
)
    comment '配置表';

