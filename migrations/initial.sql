create table if not exists organization
(
    organization_id varchar(36),
    label           varchar(256) not null,
    is_deleted      boolean default false,
    created_at      timestamp without time zone,
    updated_at      timestamp without time zone,

    primary key (organization_id)
);

create table if not exists "user"
(
    user_id         varchar(36)  not null,
    username        varchar(256) not null,
    first_name      varchar(256) not null,
    last_name       varchar(256) not null,
    patronymic      varchar(256),
    enabled         boolean                     default true,
    last_login      timestamp without time zone default null,
    role            int                         default 2,
    organization_id varchar(36)  not null,
    created_at      timestamp without time zone,
    updated_at      timestamp without time zone,

    primary key (user_id),
    foreign key (organization_id) references organization (organization_id)
);

create table if not exists "table"
(
    table_id        varchar(36)  not null,
    label           varchar(256) not null,
    is_deleted      boolean                     default false,
    organization_id varchar(36)  not null,
    created_at      timestamp without time zone,
    updated_at      timestamp without time zone,

    primary key (table_id),
    foreign key (organization_id) references organization (organization_id)
);

create table if not exists table_header
(
    table_header_id varchar(36)  not null,
    label           varchar(256) not null,
    is_deleted      boolean default false,
    table_id        varchar(36)  not null,
    parent_id       varchar(36)  not null,
    created_at      timestamp without time zone,
    updated_at      timestamp without time zone,

    primary key (table_header_id),
    foreign key (table_id) references "table" (table_id),
    foreign key (parent_id) references table_header (table_header_id)
);

create table if not exists table_row
(
    table_row_id varchar(36)  not null,
    label        varchar(256) not null,
    is_deleted boolean default false,
    table_id varchar(36)  not null,
    created_at      timestamp without time zone,
    updated_at      timestamp without time zone,

    primary key (table_row_id),
    foreign key (table_id) references "table" (table_id)


);


create table if not exists table_cell
(
    table_cell_id   varchar(36)  not null,
    table_header_id varchar(36)  not null,
    table_row_id    varchar(36)  not null,
    table_id        varchar(36)  not null,
    value           varchar(256) not null,
    type            int          not null,
    created_at      timestamp without time zone,
    updated_at      timestamp without time zone,

    primary key (table_cell_id),
    foreign key (table_id) references "table" (table_id),

    constraint constr_table_cell unique (table_header_id, table_row_id)
)
