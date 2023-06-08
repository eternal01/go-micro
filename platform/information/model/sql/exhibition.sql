create table exhibition
(
    id          bigint auto_increment comment '主键id',
    title       varchar(150) default '' not null comment '标题',
    img         varchar(255) default '' not null comment '首图',
    content     blob                    null comment '内容',
    status      tinyint      default 0  not null comment '状态',
    stage_id    int          default 0  not null comment '阶段id',
    category_id int          default 0  not null comment '类别id',
    country_id  int          default 0  not null comment '国家id',
    province_id int          default 0  not null comment '省id',
    city_id     int          default 0  not null comment '市id',
    area_id     int          default 0  not null comment '区id',
    address     varchar(255) default '' not null comment '地址',
    user_id     bigint       default 0  not null comment '用户id',
    auditor_id  bigint       default 0  not null comment '审核者id',
    created_at  timestamp    default CURRENT_TIMESTAMP not null comment '创建时间',
    updated_at  timestamp    default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '更新时间',
    audited_at  timestamp               null comment '审核时间',
    deleted_at  timestamp               null comment '删除时间',
    primary key(id)
)
    comment '展会表';

create index exhibition_idx_auditor_id
    on exhibition (auditor_id);

create index exhibition_idx_title
    on exhibition (title);

create index exhibition_idx_user_id
    on exhibition (user_id);

