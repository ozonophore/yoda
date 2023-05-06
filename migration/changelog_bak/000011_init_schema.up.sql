create table "marketplace"
(
    "id"         varchar(36)  not null primary key,
    "name"       varchar(255) not null,
    "updated_at" timestamp    not null
);

comment
on table "marketplace" is 'Marketplaces';
comment
on column "marketplace"."id" is 'ID';
comment
on column "marketplace"."name" is 'Name';
comment
on column "marketplace"."updated_at" is 'Updated at';

create table "barcode"
(
    "item_id"         varchar(36)  not null references "item" ("id") primary key,
    "barcode_id"      varchar(36)  not null,
    "barcode"         varchar(255) not null,
    "organisation_id" varchar(36)  not null references "organisation" ("id"),
    "marketplace_id"  varchar(36)  not null references "marketplace" ("id"),
    "updated_at"      timestamp    not null
);

comment
on table "barcode" is 'Barcodes';
comment
on column "barcode"."barcode_id" is 'ID';
comment
on column "barcode"."item_id" is 'Item ID';
comment
on column "barcode"."barcode" is 'Barcode';
comment
on column "barcode"."organisation_id" is 'Organisation ID';
comment
on column "barcode"."marketplace_id" is 'Marketplace ID';

create table "sale_agr"
(
    "transaction_id"          varchar(50)    not null,
    "order_date"              date           not null,
    "owner_code"              varchar(20)    not null references "owner" ("code"),
    "source"                  varchar(20)    not null,
    "supplier_article"        varchar(75),
    "warehouse_name"          varchar(50)    not null,
    "barcode"                 varchar(30),
    "external_code"           varchar(50)    not null,
    "quantity"                numeric(10)    not null,
    "sum_total_price"         numeric(10, 2) not null,
    "sum_price_with_discount" numeric(10, 2) not null,
    primary key ("transaction_id", "order_date", "owner_code", "source", "warehouse_name", "external_code")
);

comment on table "sale_agr" is 'Таблица с данными по продажам';
comment on column "sale_agr"."transaction_id" is 'Идентификатор транзакции';
comment on column "sale_agr"."order_date" is 'Дата заказа';
comment on column "sale_agr"."owner_code" is 'Код владельца';
comment on column "sale_agr"."source" is 'Источник';
comment on column "sale_agr"."supplier_article" is 'Артикул поставщика';
comment on column "sale_agr"."warehouse_name" is 'Название склада';
comment on column "sale_agr"."barcode" is 'Штрихкод';
comment on column "sale_agr"."external_code" is 'Внешний код';
comment on column "sale_agr"."quantity" is 'Количество';
comment on column "sale_agr"."sum_total_price" is 'Цкнва';
comment on column "sale_agr"."sum_price_with_discount" is 'Цена с учетом скидки';
