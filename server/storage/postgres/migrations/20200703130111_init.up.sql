create table events
(
    id            uuid         not null,
    streamer_name varchar(255) not null,
    event_type    varchar(20)  not null,
    created_at    timestamp    not null,
    constraint events_pk primary key (id)
);