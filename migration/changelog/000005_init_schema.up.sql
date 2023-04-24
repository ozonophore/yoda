create table "order_delivered_arch"
(
    "order_date"          date                                    not null,
    "transaction_id"      integer references "transaction" ("id") not null,
    "source"              varchar(20)                             not null,
    "owner_code"          varchar(20) references "owner" ("code") not null,
    "supplier_article"    varchar(75)                             not null,
    "warehouse"           varchar(50)                             not null,
    "barcode"             varchar(30),
    "external_code"       varchar(50)                             not null,
    "total_price"         numeric(10, 2),
    "price_with_discount" numeric(10, 2),
    "quantity"            numeric(10)                             not null
) PARTITION BY RANGE("order_date");

CREATE TABLE order_delivered_arch_def PARTITION OF order_delivered_arch DEFAULT;
CREATE TABLE order_delivered_arch_202304 PARTITION OF order_delivered_arch
    FOR VALUES FROM
(
    '2023-04-01'
) TO
(
    '2023-05-01'
);
CREATE TABLE order_delivered_arch_202305 PARTITION OF order_delivered_arch
    FOR VALUES FROM
(
    '2023-05-01'
) TO
(
    '2023-06-01'
);
CREATE TABLE order_delivered_arch_202306 PARTITION OF order_delivered_arch
    FOR VALUES FROM
(
    '2023-06-01'
) TO
(
    '2023-07-01'
);
CREATE TABLE order_delivered_arch_202307 PARTITION OF order_delivered_arch
    FOR VALUES FROM
(
    '2023-07-01'
) TO
(
    '2023-08-01'
);
CREATE TABLE order_delivered_arch_202308 PARTITION OF order_delivered_arch
    FOR VALUES FROM
(
    '2023-08-01'
) TO
(
    '2023-09-01'
);
CREATE TABLE order_delivered_arch_202309 PARTITION OF order_delivered_arch
    FOR VALUES FROM
(
    '2023-09-01'
) TO
(
    '2023-10-01'
);
CREATE TABLE order_delivered_arch_202310 PARTITION OF order_delivered_arch
    FOR VALUES FROM
(
    '2023-10-01'
) TO
(
    '2023-11-01'
);
CREATE TABLE order_delivered_arch_202311 PARTITION OF order_delivered_arch
    FOR VALUES FROM
(
    '2023-11-01'
) TO
(
    '2023-12-01'
);
CREATE TABLE order_delivered_arch_202312 PARTITION OF order_delivered_arch
    FOR VALUES FROM
(
    '2023-12-01'
) TO
(
    '2024-01-01'
);

comment
on table "order_delivered_arch" is 'Архив заказов поставщиков';
comment
on column "order_delivered_arch"."order_date" is 'Дата заказа';
comment
on column "order_delivered_arch"."transaction_id" is 'Идентификатор транзакции';
comment
on column "order_delivered_arch"."source" is 'Источник заказа';
comment
on column "order_delivered_arch"."owner_code" is 'Код владельца';
comment
on column "order_delivered_arch"."supplier_article" is 'Артикул поставщика';
comment
on column "order_delivered_arch"."warehouse" is 'Склад';
comment
on column "order_delivered_arch"."barcode" is 'Штрихкод';
comment
on column "order_delivered_arch"."external_code" is 'Внешний код';
comment
on column "order_delivered_arch"."total_price" is 'Сумма заказа';
comment
on column "order_delivered_arch"."price_with_discount" is 'Сумма заказа с учетом скидки';
comment
on column "order_delivered_arch"."quantity" is 'Количество';

create table "order_delivered"
(
    "order_date"          date                                    not null,
    "source"              varchar(20)                             not null,
    "owner_code"          varchar(20) references "owner" ("code") not null,
    "supplier_article"    varchar(75)                             not null,
    "warehouse"           varchar(50)                             not null,
    "barcode"             varchar(30),
    "external_code"       varchar(50)                             not null,
    "total_price"         numeric(10, 2),
    "price_with_discount" numeric(10, 2),
    "quantity"            numeric(10)                             not null,
    "create_at"           timestamp with time zone                not null default now(),
    "updated_at"          timestamp with time zone                not null default now()
) PARTITION BY RANGE("order_date");

CREATE TABLE order_delivered_def PARTITION OF order_delivered DEFAULT;
CREATE TABLE order_delivered_202304 PARTITION OF order_delivered
    FOR VALUES FROM
(
    '2023-04-01'
) TO
(
    '2023-05-01'
);
CREATE TABLE order_delivered_202305 PARTITION OF order_delivered
    FOR VALUES FROM
(
    '2023-05-01'
) TO
(
    '2023-06-01'
);
CREATE TABLE order_delivered_202306 PARTITION OF order_delivered
    FOR VALUES FROM
(
    '2023-06-01'
) TO
(
    '2023-07-01'
);
CREATE TABLE order_delivered_202307 PARTITION OF order_delivered
    FOR VALUES FROM
(
    '2023-07-01'
) TO
(
    '2023-08-01'
);
CREATE TABLE order_delivered_202308 PARTITION OF order_delivered
    FOR VALUES FROM
(
    '2023-08-01'
) TO
(
    '2023-09-01'
);
CREATE TABLE order_delivered_202309 PARTITION OF order_delivered
    FOR VALUES FROM
(
    '2023-09-01'
) TO
(
    '2023-10-01'
);
CREATE TABLE order_delivered_202310 PARTITION OF order_delivered
    FOR VALUES FROM
(
    '2023-10-01'
) TO
(
    '2023-11-01'
);
CREATE TABLE order_delivered_202311 PARTITION OF order_delivered
    FOR VALUES FROM
(
    '2023-11-01'
) TO
(
    '2023-12-01'
);
CREATE TABLE order_delivered_202312 PARTITION OF order_delivered
    FOR VALUES FROM
(
    '2023-12-01'
) TO
(
    '2024-01-01'
);

comment
on table "order_delivered" is 'Таблица для хранения отгруженных заказов';
comment
on column "order_delivered"."order_date" is 'Дата заказа';
comment
on column "order_delivered"."source" is 'Источник заказа';
comment
on column "order_delivered"."owner_code" is 'Код владельца';
comment
on column "order_delivered"."supplier_article" is 'Артикул поставщика';
comment
on column "order_delivered"."warehouse" is 'Склад';
comment
on column "order_delivered"."barcode" is 'Штрихкод';
comment
on column "order_delivered"."total_price" is 'Сумма заказа';
comment
on column "order_delivered"."price_with_discount" is 'Сумма заказа с учетом скидки';
comment
on column "order_delivered"."warehouse" is 'Название склада';
comment
on column "order_delivered"."quantity" is 'Количество';
comment
on column "order_delivered"."external_code" is 'Внешний код';

create table "order_delivered_log"
(
    "transaction_id" bigint references "transaction" ("id") not null primary key,
    "created_at"     timestamp                              not null,
    "added_rows"     integer                                not null
);

comment
on table "order_delivered_log" is 'Таблица для хранения логов отгруженных заказов';
comment
on column "order_delivered_log"."transaction_id" is 'Идентификатор транзакции';
comment
on column "order_delivered_log"."created_at" is 'Дата создания записи';
