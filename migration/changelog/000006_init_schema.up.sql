create table "dl"."organisation"
(
    "id"       varchar(36) primary key,
    "name"     varchar(200),
    "inn"      varchar(15),
    "kpp"      varchar(15),
    "update_at" timestamp not null
);

comment
on table "dl"."organisation" is 'Таблица для хранения организаций';
    comment
on column "dl"."organisation"."id" is 'Идентификатор организации';
    comment
on column "dl"."organisation"."name" is 'Наименование организации';
    comment
on column "dl"."organisation"."inn" is 'ИНН организации';
    comment
on column "dl"."organisation"."kpp" is 'КПП организации';
    comment
on column "dl"."organisation"."update_at" is 'Дата и время обновления записи';

alter table "ml"."owner" add column "organisation_id" varchar(36) references "dl"."organisation" ("id");
create unique index ownr__indx_unq on "ml"."owner" ("code", "organisation_id");