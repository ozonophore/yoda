create table dl."stock_def30"(
                                 stock_date date not null,
                                 source varchar(20) not null,
                                 owner_code varchar(20) not null,
                                 supplier_article varchar(75),
                                 barcode varchar(30),
                                 external_code varchar(50) not null,
                                 warehouse varchar(50) not null,
                                 quantity numeric(18, 2) not null,
                                 quantity_full numeric(18, 2) not null,
                                 attempt integer not null,
                                 price numeric(18, 2) not null,
                                 price_with_discount numeric(18, 2) not null,
                                 create_at timestamp not null,
                                 barcode_id varchar(36),
                                 item_id varchar(36),
                                 org_id varchar(36),
                                 org_name varchar(255),
                                 marketplace_id varchar(36),
                                 item_name varchar(255)
) PARTITION BY RANGE("stock_date");

comment on table dl."stock_def30" is 'Таблица с данными по дефектуре за 30 дней';
comment on column dl."stock_def30".stock_date is 'Дата';
comment on column dl."stock_def30".source is 'Источник';
comment on column dl."stock_def30".owner_code is 'Код владельца';
comment on column dl."stock_def30".supplier_article is 'Артикул поставщика';
comment on column dl."stock_def30".barcode is 'Штрихкод';
comment on column dl."stock_def30".external_code is 'Внешний код';
comment on column dl."stock_def30".warehouse is 'Склад';
comment on column dl."stock_def30".quantity is 'Количество';
comment on column dl."stock_def30".quantity_full is 'Количество полное';
comment on column dl."stock_def30".attempt is 'Количество попыток';
comment on column dl."stock_def30".price is 'Цена';
comment on column dl."stock_def30".price_with_discount is 'Цена со скидкой';
comment on column dl."stock_def30".create_at is 'Дата создания';


CREATE TABLE "dl"."stock_def30_def" PARTITION OF "dl"."stock_def30" DEFAULT;

create
or replace procedure "ml".partition_for_stock_def30(IN start_date date, IN end_date date)
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

call "ml".partition_for_stock_def30(to_date('2023-01-01', 'yyyy-MM-dd'), to_date('2024-01-01', 'yyyy-MM-dd'));
