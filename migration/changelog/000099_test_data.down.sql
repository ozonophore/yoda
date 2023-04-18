create table "1c_org"
(
    "id"        integer primary key,
    "name"      varchar(200),
    "create_at" timestamp not null
);

comment
on table "1c_org" is 'Таблица для хранения организаций из 1С';
comment
on column "1c_org"."id" is 'Идентификатор организации';
comment
on column "1c_org"."name" is 'Наименование организации';
comment
on column "1c_org"."create_at" is 'Дата и время создания записи';

create table "1c_marketplace"
(
    "id"        integer primary key,
    "name"      varchar(200),
    "create_at" timestamp not null
);

comment
on table "1c_marketplace" is 'Таблица для хранения маркетплейсов из 1С';
comment
on column "1c_marketplace"."id" is 'Идентификатор маркетплейса';
comment
on column "1c_marketplace"."name" is 'Наименование маркетплейса';
comment
on column "1c_marketplace"."create_at" is 'Дата и время создания записи';

create table "1c_items"
(
    "code1c"    varchar(50) primary key,
    "name"      varchar(200),
    "article"   varchar(75),
    "character" varchar(200),
    "create_at" timestamp not null
);

comment
on table "1c_items" is 'Таблица для хранения товаров из 1С';
comment
on column "1c_items"."code1c" is 'Код товара в 1С';
comment
on column "1c_items"."name" is 'Наименование товара';
comment
on column "1c_items"."article" is 'Артикул товара';
comment
on column "1c_items"."character" is 'Характеристика товара';
comment
on column "1c_items"."create_at" is 'Дата и время создания записи';

create table "1c_barcode"
(
    "code1c"    varchar(50) not null references "1c_items" ("code1c"),
    "barcode"   varchar(50) not null,
    "org_id"    integer     not null references "1c_org" ("id"),
    "market_id" integer     not null references "1c_marketplace" ("id"),
    "article"   varchar(75),
    "create_at" timestamp   not null,
    primary key ("org_id", "market_id", "barcode")
);

comment
on table "1c_barcode" is 'Таблица для хранения штрихкодов товаров из 1С';
comment
on column "1c_barcode"."code1c" is 'Код товара в 1С';
comment
on column "1c_barcode"."barcode" is 'Штрихкод товара';
comment
on column "1c_barcode"."org_id" is 'Идентификатор организации';
comment
on column "1c_barcode"."market_id" is 'Идентификатор маркетплейса';
comment
on column "1c_barcode"."article" is 'Артикул товара';
comment
on column "1c_barcode"."create_at" is 'Дата и время создания записи';

create table "1c_stock"
(
    "code1c"    varchar(50) not null,
    "character" varchar(200),
    "quantity"  integer     not null,
    "update_at" timestamp   not null
);


comment
on table "1c_stock" is 'Таблица для хранения остатков товаров из 1С';
comment
on column "1c_stock"."code1c" is 'Код товара в 1С';
comment
on column "1c_stock"."character" is 'Характеристика товара';
comment
on column "1c_stock"."quantity" is 'Количество товара';
comment
on column "1c_stock"."update_at" is 'Дата и время обновления остатков';

create table "1c_warehouse"
(
    "id"        integer primary key,
    "org_id"    integer   not null references "1c_org" ("id"),
    "market_id" integer   not null references "1c_marketplace" ("id"),
    "name"      varchar(200),
    "create_at" timestamp not null
);

comment on table "1c_warehouse" is 'Таблица для хранения складов из 1С';
comment on column "1c_warehouse"."id" is 'Идентификатор склада';
comment on column "1c_warehouse"."org_id" is 'Идентификатор организации';
comment on column "1c_warehouse"."market_id" is 'Идентификатор маркетплейса';
comment on column "1c_warehouse"."name" is 'Наименование склада';
comment on column "1c_warehouse"."create_at" is 'Дата и время создания записи';
