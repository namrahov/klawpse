-- +migrate Up
create table if not exists application
(
    id                 bigserial    not null primary key,
    request_id         bigserial    not null,
    checked_id         bigserial    not null,
    person             varchar(16)  not null,
    customer_type      varchar(16)  not null,
    customer_name      varchar(32)  not null,
    file_path          varchar(32)  not null,
    court_name         varchar(32)  not null,
    judge_name         varchar(32)  not null,
    decision_number    varchar(32)  not null,
    decision_date      varchar(32)  not null,
    is_checked         varchar(32)  not null,
    status             varchar(32)  not null,
    status_history_id  bigserial    not null,
    created_at         timestamp    not null default now(),
    updated_at         timestamp    not null default current_timestamp
);
