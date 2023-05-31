create table system_stage
(
    id          int auto_increment comment '主键id',
    name        varchar(100) default '' not null comment '阶段名称',
    parent_id   int       default 0  not null comment '父级id',
    description varchar(255) default '' not null comment '阶段描述',
    primary key(id)
)
    comment '分类表';

create index system_stage_idx_parent_id
    on system_stage (parent_id);