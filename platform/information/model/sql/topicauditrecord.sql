create table topic_audit_record
(
    id         bigint auto_increment comment '主键id',
    topic_id   bigint       default 0                 not null comment '文章id',
    status     tinyint      default 0                 not null comment '审核状态',
    remark     varchar(255) default ''                not null comment '备注',
    auditor_id bigint       default 0                 not null comment '审核者id',
    created_at timestamp    default CURRENT_TIMESTAMP not null comment '创建时间',
    updated_at timestamp    default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '更新时间',
    deleted_at timestamp                              null comment '删除时间',
    primary key(id)
)
    comment '文章审核记录表';

create index topic_audit_record_idx_topic_id
    on topic_audit_record (topic_id);

create index topic_audit_record_idx_auditor_id
    on topic_audit_record (auditor_id);
