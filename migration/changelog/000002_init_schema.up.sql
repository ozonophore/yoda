create table "dl"."order_delivered_arch"
(
    "order_date"          date                                    not null,
    "transaction_id"      integer references "ml"."transaction" ("id") not null,
    "source"              varchar(20)                             not null,
    "owner_code"          varchar(20) references "ml"."owner" ("code") not null,
    "supplier_article"    varchar(75)                             not null,
    "warehouse"           varchar(50)                             not null,
    "barcode"             varchar(30),
    "external_code"       varchar(50)                             not null,
    "total_price"         numeric(10, 2),
    "price_with_discount" numeric(10, 2),
    "quantity"            numeric(10)                             not null
) PARTITION BY RANGE("order_date");

comment
on table "dl"."order_delivered_arch" is 'Архив заказов поставщиков';
comment
on column "dl"."order_delivered_arch"."order_date" is 'Дата заказа';
comment
on column "dl"."order_delivered_arch"."transaction_id" is 'Идентификатор транзакции';
comment
on column "dl"."order_delivered_arch"."source" is 'Источник заказа';
comment
on column "dl"."order_delivered_arch"."owner_code" is 'Код владельца';
comment
on column "dl"."order_delivered_arch"."supplier_article" is 'Артикул поставщика';
comment
on column "dl"."order_delivered_arch"."warehouse" is 'Склад';
comment
on column "dl"."order_delivered_arch"."barcode" is 'Штрихкод';
comment
on column "dl"."order_delivered_arch"."external_code" is 'Внешний код';
comment
on column "dl"."order_delivered_arch"."total_price" is 'Сумма заказа';
comment
on column "dl"."order_delivered_arch"."price_with_discount" is 'Сумма заказа с учетом скидки';
comment
on column "dl"."order_delivered_arch"."quantity" is 'Количество';

CREATE TABLE "dl"."order_delivered_arch_def" PARTITION OF "dl"."order_delivered_arch" DEFAULT;

create
or replace procedure "ml".partition_for_order_delivered_arch(IN start_date date, IN end_date date)
    language plpgsql
as
$$
declare
v_table varchar(50);
v_from
varchar(10);
v_to
varchar(10);
begin
for v_table, v_from, v_to in SELECT 'order_delivered_arch_' || to_char(day::date,'YYYYMM'),
                                    to_char(day::date,'YYYY-MM-DD'),
                                    to_char(day::date + interval '1 month','YYYY-MM-DD')
                             FROM generate_series(start_date, end_date, '1 month') day
            loop
            EXECUTE format(
                'CREATE TABLE IF NOT EXISTS "dl"."%s" PARTITION OF "dl"."order_delivered_arch" FOR VALUES FROM(''%s'') TO (''%s'')',
                v_table,
                v_from,
                v_to
             );
end loop;
end;$$;

call "ml".partition_for_order_delivered_arch(to_date('2023-01-01', 'yyyy-MM-dd'), to_date('2024-01-01', 'yyyy-MM-dd'));

create table "dl"."order_delivered"
(
    "order_date"          date                                    not null,
    "source"              varchar(20)                             not null,
    "owner_code"          varchar(20) references "ml"."owner" ("code") not null,
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

comment
on table "dl"."order_delivered" is 'Таблица для хранения отгруженных заказов';
comment
on column "dl"."order_delivered"."order_date" is 'Дата заказа';
comment
on column "dl"."order_delivered"."source" is 'Источник заказа';
comment
on column "dl"."order_delivered"."owner_code" is 'Код владельца';
comment
on column "dl"."order_delivered"."supplier_article" is 'Артикул поставщика';
comment
on column "dl"."order_delivered"."warehouse" is 'Склад';
comment
on column "dl"."order_delivered"."barcode" is 'Штрихкод';
comment
on column "dl"."order_delivered"."total_price" is 'Сумма заказа';
comment
on column "dl"."order_delivered"."price_with_discount" is 'Сумма заказа с учетом скидки';
comment
on column "dl"."order_delivered"."warehouse" is 'Название склада';
comment
on column "dl"."order_delivered"."quantity" is 'Количество';
comment
on column "dl"."order_delivered"."external_code" is 'Внешний код';

CREATE TABLE "dl"."order_delivered_def" PARTITION OF "dl"."order_delivered" DEFAULT;

create
or replace procedure "ml".partition_for_order_delivered(IN start_date date, IN end_date date)
    language plpgsql
as
$$
declare
v_table varchar(50);
v_from
varchar(10);
v_to
varchar(10);
begin
for v_table, v_from, v_to in
SELECT 'order_delivered_' || to_char(day::date,'YYYYMM'),
       to_char(day::date,'YYYY-MM-DD'),
       to_char(day::date + interval '1 month','YYYY-MM-DD')
FROM generate_series(start_date, end_date, '1 month') day
            loop
            EXECUTE format(
                'CREATE TABLE IF NOT EXISTS "dl"."%s" PARTITION OF "dl"."order_delivered" FOR VALUES FROM(''%s'') TO (''%s'')',
                v_table,
                v_from,
                v_to
             );
end loop;
end;$$;

call "ml".partition_for_order_delivered(to_date('2023-01-01', 'yyyy-MM-dd'), to_date('2024-01-01', 'yyyy-MM-dd'));

create table "ml"."order_delivered_log"
(
    "transaction_id" bigint references "ml"."transaction" ("id") not null primary key,
    "created_at"     timestamp                              not null,
    "added_rows"     integer                                not null
);

comment
on table "ml"."order_delivered_log" is 'Таблица для хранения логов отгруженных заказов';
comment
on column "ml"."order_delivered_log"."transaction_id" is 'Идентификатор транзакции';
comment
on column "ml"."order_delivered_log"."created_at" is 'Дата создания записи';
