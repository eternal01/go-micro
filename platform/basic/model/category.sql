create table system_category
(
    id          int auto_increment comment '主键id',
    name        varchar(100) default '' not null comment '类别名称',
    parent_id   int       default 0  not null comment '父级id',
    description varchar(255) default '' not null comment '类别描述',
    primary key(id)
)
    comment '类别表';

create index system_category_idx_parent_id
    on system_category (parent_id);