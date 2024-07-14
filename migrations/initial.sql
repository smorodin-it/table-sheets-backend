create table if not exists organization
(
    organization_id varchar(36),
    label           varchar(256) not null,
    enabled         boolean default true,
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
    is_deleted      boolean default false,
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
    parent_id       varchar(36),
    lft             int          not null,
    rgt             int          not null,
    created_at      timestamp without time zone,
    updated_at      timestamp without time zone,

    primary key (table_header_id),
    foreign key (table_id) references "table" (table_id),
    foreign key (parent_id) references table_header (table_header_id),
    constraint constr_table_header unique (table_header_id, lft, rgt)
);

create table if not exists table_row
(
    table_row_id varchar(36) not null,
    label        varchar(256) default null,
    is_deleted   boolean      default false,
    table_id     varchar(36) not null,
    created_at   timestamp without time zone,
    updated_at   timestamp without time zone,

    primary key (table_row_id),
    foreign key (table_id) references "table" (table_id)


);


create table if not exists table_cell
(
    table_cell_id   varchar(36) not null,
    table_header_id varchar(36) not null,
    table_row_id    varchar(36) not null,
    table_id        varchar(36) not null,
    value           varchar(256),
    type            int         not null,
    created_at      timestamp without time zone,
    updated_at      timestamp without time zone,

    primary key (table_cell_id),
    foreign key (table_id) references "table" (table_id),

    constraint constr_table_cell unique (table_header_id, table_row_id)
);

create table if not exists table_cell_2_table_cell
(
    table_cell_value_id    varchar(36) not null,
    table_cell_argument_id varchar(36) not null,
    id                     varchar(36) not null,
    created_at             timestamp without time zone,
    updated_at             timestamp without time zone,

    primary key (id),
    foreign key (table_cell_value_id) references table_cell (table_cell_id),
    foreign key (table_cell_argument_id) references table_cell (table_cell_id),

    constraint constr_table_cell_2_table_cell unique (table_cell_value_id, table_cell_argument_id)
);

create or replace procedure add_table_header(
    tableHeaderId table_header.table_header_id%type,
    headerLabel table_header.label%type,
    tableId table_header.table_id%type,
    parentId table_header.table_id%type,
    createdAt table_header.created_at%type,
    updatedAt table_header.updated_at%type
)
    language plpgsql
as
$$
declare
    myRight int;

begin
    --         TODO: IMPLEMENT TABLE LOCK
    select table_header.rgt into myRight from table_header where table_header_id = parentId;

    update table_header set rgt = rgt + 2 where rgt > myRight;
    update table_header set lft = lft + 2 where lft > myRight;

    insert into table_header (table_header_id, label, is_deleted, table_id, parent_id, lft, rgt, created_at, updated_at)
    values (tableHeaderId, headerLabel, false, tableId, parentId, myRight + 1, myRight + 2, createdAt, updatedAt);

--     commit;
end;
$$