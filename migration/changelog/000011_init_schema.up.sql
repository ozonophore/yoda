create table dl.stock_def30 (
                                stock_date date,
                                source varchar(20),
                                owner_code varchar(20),
                                supplier_article varchar(75),
                                warehouse varchar(50),
                                barcode varchar(30),
                                external_code varchar(50),
                                marketplace_id varchar(36),
                                barcode_id varchar(36),
                                item_id varchar(36),
                                org_id varchar(36),
                                item_name varchar(255),
                                def30 integer,
                                days_in_stock integer,
                                avg_price numeric(19, 2),
                                min_price numeric(19, 2),
                                max_price numeric(19, 2),
                                create_at timestamp default now()
) PARTITION BY RANGE("stock_date");

comment on table dl.stock_def30 is 'Таблица с данными по дефициту 30 дней';
comment on column dl.stock_def30.stock_date is 'Дата';
comment on column dl.stock_def30.source is 'Источник';
comment on column dl.stock_def30.owner_code is 'Код владельца';
comment on column dl.stock_def30.supplier_article is 'Артикул поставщика';
comment on column dl.stock_def30.warehouse is 'Склад';
comment on column dl.stock_def30.barcode is 'Штрихкод';
comment on column dl.stock_def30.external_code is 'Внешний код';
comment on column dl.stock_def30.marketplace_id is 'Идентификатор маркетплейса';
comment on column dl.stock_def30.barcode_id is 'Идентификатор штрихкода';
comment on column dl.stock_def30.item_id is 'Идентификатор товара';
comment on column dl.stock_def30.org_id is 'Идентификатор организации';
comment on column dl.stock_def30.item_name is 'Наименование товара';
comment on column dl.stock_def30.def30 is 'Дефицит 30 дней';
comment on column dl.stock_def30.days_in_stock is 'Дней в наличии';
comment on column dl.stock_def30.avg_price is 'Средняя цена';
comment on column dl.stock_def30.min_price is 'Минимальная цена';
comment on column dl.stock_def30.max_price is 'Максимальная цена';
comment on column dl.stock_def30.create_at is 'Дата создания записи';

CREATE TABLE "dl"."stock_def30_def" PARTITION OF "dl"."stock_def30" DEFAULT;

create index IF NOT EXISTS stock_def30_def_stock_date_idx on dl.stock_def30(stock_date);

create
or replace procedure "dl".partition_for_stock_def30(IN start_date date, IN end_date date)
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
for v_table, v_from, v_to in SELECT 'stock_def30_' || to_char(day::date,'YYYYMM'),
                                    to_char(day::date,'YYYY-MM-DD'),
                                    to_char(day::date + interval '1 month','YYYY-MM-DD')
                             FROM generate_series(start_date, end_date, '1 month') day
            loop
            EXECUTE format(
                'CREATE TABLE IF NOT EXISTS "dl"."%s" PARTITION OF "dl"."stock_def30" FOR VALUES FROM(''%s'') TO (''%s'')',
                v_table,
                v_from,
                v_to
             );
end loop;
end;$$;

call "dl".partition_for_stock_def30(to_date('2023-01-01', 'yyyy-MM-dd'), to_date('2024-01-01', 'yyyy-MM-dd'));
