create table news
(
    id         bigint auto_increment,
    title      varchar(150) default '' not null comment '标题',
    img        varchar(255) default '' not null comment '首图',
    content    blob                    null comment '内容',
    status     tinyint      default 0  not null comment '状态',
    user_id    bigint       default 0  not null comment '用户id',
    auditor_id bigint       default 0  not null comment '审核者id',
    created_at  timestamp    default CURRENT_TIMESTAMP not null comment '创建时间',
    updated_at  timestamp    default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '更新时间',
    audited_at timestamp               null comment '审核时间',
    deleted_at timestamp               null comment '删除时间',
    constraint news_pk
        primary key (id)
)
    comment '新闻表';

create index news_idx_auditor_id
    on news (auditor_id);

create index news_idx_title
    on news (title);

create index news_idx_user_id
    on news (user_id);