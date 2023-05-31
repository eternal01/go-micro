create table system_region
(
    id        smallint auto_increment comment '主键id',
    parent_id smallint     default 0  not null comment '父级id',
    name      varchar(120) default '' not null comment '区域id',
    primary key(id)
);

create index system_region_idx_parent_id
    on system_region (parent_id);

