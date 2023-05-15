create table "ml"."transaction_log" (
                                        id integer not null,
                                        status varchar(50) not null,
                                        msg varchar(255),
                                        created_at timestamp not null,
                                        primary key (id, status)
);

comment on table "ml"."transaction_log" is 'Таблица логов транзакций';
comment on column "ml"."transaction_log".id is 'Идентификатор транзакции';
comment on column "ml"."transaction_log".status is 'Статус транзакции';
comment on column "ml"."transaction_log".msg is 'Сообщение об ошибке';
comment on column "ml"."transaction_log".created_at is 'Дата создания записи';

alter table "dl"."item" add primary key (id);
alter table "dl"."marketplace" add primary key (id);
alter table "ml"."owner" add primary key (code);
alter table "ml"."owner_marketplace" add primary key ("owner_code", "source");

create table "ml"."marketplace" (
                                    "code" varchar(15) not null primary key ,
                                    "marketplace_id" varchar(36) not null
);

insert into "ml"."marketplace" ("code", "marketplace_id") values ('WB', '00-00360961');
insert into "ml"."marketplace" ("code", "marketplace_id") values ('OZON', '00-00374442');