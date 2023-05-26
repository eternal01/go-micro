create table system_classify
(
    id          int auto_increment comment '主键id',
    name        varchar(100) default '' not null comment '分类名称',
    parent_id   int       default 0  not null comment '父级id',
    description varchar(255) default '' not null comment '分类描述',
    primary key(id)
)
    comment '分类表';

create index system_classify_idx_parent_id
    on system_classify (parent_id);