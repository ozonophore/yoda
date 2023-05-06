create table "organisation"
(
    "id"       varchar(36) primary key,
    "name"     varchar(200),
    "inn"      varchar(15),
    "kpp"      varchar(15),
    "updateAt" timestamp not null
);

comment
on table "organisation" is 'Таблица для хранения организаций';
    comment
on column "organisation"."id" is 'Идентификатор организации';
    comment
on column "organisation"."name" is 'Наименование организации';
    comment
on column "organisation"."inn" is 'ИНН организации';
    comment
on column "organisation"."kpp" is 'КПП организации';
    comment
on column "organisation"."updateAt" is 'Дата и время обновления записи';

alter table "owner" add column "organisation_id" varchar(36) references "organisation" ("id");
create unique index ownr__indx_unq on "owner" ("code", "organisation_id");