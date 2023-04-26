create table basic.system_industry
(
    id          int auto_increment comment '主键id'
        primary key,
    industry_id varchar(12)   default '' not null comment '行业id',
    name        varchar(1000) default '' not null comment '行业名称',
    parent_id   varchar(12)   default '' not null comment '父级id',
    level_type  int           default 0  not null comment '等级',
    description varchar(4000) default '' not null comment '描述'
)
    comment '行业表';

