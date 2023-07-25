create table "dl"."marketplace"
(
    "id"         varchar(36)  not null primary key,
    "name"       varchar(255) not null,
    "updated_at" timestamp    not null
);

comment
on table "dl"."marketplace" is 'Marketplaces';
comment
on column "dl"."marketplace"."id" is 'ID';
comment
on column "dl"."marketplace"."name" is 'Name';
comment
on column "dl"."marketplace"."updated_at" is 'Updated at';

create table "dl"."barcode"
(
    "item_id"         varchar(36)  not null references "dl"."item" ("id") primary key,
    "barcode_id"      varchar(36)  not null,
    "barcode"         varchar(255) not null,
    "organisation_id" varchar(36)  not null references "dl"."organisation" ("id"),
    "marketplace_id"  varchar(36)  not null references "dl"."marketplace" ("id"),
    "updated_at"      timestamp    not null
);

comment
on table "dl"."barcode" is 'Barcodes';
comment
on column "dl"."barcode"."barcode_id" is 'ID';
comment
on column "dl"."barcode"."item_id" is 'Item ID';
comment
on column "dl"."barcode"."barcode" is 'Barcode';
comment
on column "dl"."barcode"."organisation_id" is 'Organisation ID';
comment
on column "dl"."barcode"."marketplace_id" is 'Marketplace ID';

create table "dl"."sale_agr"
(
    "transaction_id"          varchar(50)    not null,
    "order_date"              date           not null,
    "owner_code"              varchar(20)    not null references "ml"."owner" ("code"),
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

comment on table "dl"."sale_agr" is 'Таблица с данными по продажам';
comment on column "dl"."sale_agr"."transaction_id" is 'Идентификатор транзакции';
comment on column "dl"."sale_agr"."order_date" is 'Дата заказа';
comment on column "dl"."sale_agr"."owner_code" is 'Код владельца';
comment on column "dl"."sale_agr"."source" is 'Источник';
comment on column "dl"."sale_agr"."supplier_article" is 'Артикул поставщика';
comment on column "dl"."sale_agr"."warehouse_name" is 'Название склада';
comment on column "dl"."sale_agr"."barcode" is 'Штрихкод';
comment on column "dl"."sale_agr"."external_code" is 'Внешний код';
comment on column "dl"."sale_agr"."quantity" is 'Количество';
comment on column "dl"."sale_agr"."sum_total_price" is 'Цкнва';
comment on column "dl"."sale_agr"."sum_price_with_discount" is 'Цена с учетом скидки';
