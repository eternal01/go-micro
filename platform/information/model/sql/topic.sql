create table topic
(
    id          bigint auto_increment,
    title       varchar(150) default '' not null comment '标题',
    img         varchar(255) default '' not null comment '首图',
    content     blob                    null comment '内容',
    status      tinyint      default 0  not null comment '状态',
    stage_id    int          default 0  not null comment '阶段id',
    category_id int          default 0  not null comment '类别id',
    user_id     bigint       default 0  not null comment '用户id',
    auditor_id  bigint       default 0  not null comment '审核者id',
    created_at  timestamp    default CURRENT_TIMESTAMP not null comment '创建时间',
    updated_at  timestamp    default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '更新时间',
    audited_at  timestamp               null comment '审核时间',
    deleted_at  timestamp               null comment '删除时间',
    primary key(id)
)
    comment '文章表';

create index topic_idx_auditor_id
    on topic (auditor_id);

create index topic_idx_title
    on topic (title);

create index topic_idx_user_id
    on topic (user_id);

