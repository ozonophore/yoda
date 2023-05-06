create table "dl"."stock_daily"
(
    "stock_date" date not null,
    "source" varchar(20) not null,
    "owner_code" varchar(20) not null,
    "supplier_article" varchar(75),
    "barcode" varchar(30),
    "external_code" varchar(50) not null,
    "name" varchar(200),
    "subject" varchar(200),
    "category" varchar(50),
    "brand" varchar(50),
    "warehouse" varchar(50) not null,
    "create_at" date not null,
    "update_at" date,
    "quantity" numeric(10,2) not null,
    "quantity_full" numeric(10,2) not null,
    "attention" int not null,
    "price" numeric(10,2) not null,
    "price_with_discount" numeric(10,2) not null
) PARTITION BY RANGE("stock_date");

CREATE TABLE "dl"."stock_daily_def" PARTITION OF "dl"."stock_daily" DEFAULT;

create index stock_daily_def_soec_idx on "dl"."stock_daily_def" ("source", "owner_code", "warehouse", "external_code");

create
or replace procedure "ml".partition_for_stock_daily(IN start_date date, IN end_date date)
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
SELECT 'stock_daily_' || to_char(day::date,'YYYYMM'),
       to_char(day::date,'YYYY-MM-DD'),
       to_char(day::date + interval '1 month','YYYY-MM-DD')
FROM generate_series(start_date, end_date, '1 month') day
            loop
            EXECUTE format(
                'CREATE TABLE IF NOT EXISTS "dl"."%s" PARTITION OF "dl"."stock_daily" FOR VALUES FROM(''%s'') TO (''%s'')',
                v_table,
                v_from,
                v_to
             );
EXECUTE format(
        'create index %s_soec_idx on "dl"."%s" ("source", "owner_code", "warehouse", "external_code")',
        v_table,
        v_table
    );
end loop;
end;$$;

call "ml".partition_for_stock_daily(to_date('2023-01-01', 'yyyy-MM-dd'), to_date('2024-01-01', 'yyyy-MM-dd'));
