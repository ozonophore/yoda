
create table dl.sales_stock(
                               report_date date,
                               source varchar(5),
                               owner_code varchar(20),
                               supplier_article varchar(75),
                               barcode varchar(50),
                               warehouse_name varchar(50),
                               external_code varchar(50),
                               def30 numeric(10),
                               days_in_stock30 numeric(10),
                               def5 numeric(10),
                               days_in_stock5 numeric(10),
                               avg_price numeric(10,2),
                               min_price numeric(10,2),
                               max_price numeric(10,2),
                               quantity30 numeric(10),
                               quantity5 numeric(10)
) partition by range(report_date);

comment on table dl.sales_stock is 'Таблица продаж по складам';
comment on column dl.sales_stock.report_date is 'Дата отчета';
comment on column dl.sales_stock.source is 'Источник';
comment on column dl.sales_stock.owner_code is 'Код владельца';
comment on column dl.sales_stock.supplier_article is 'Артикул поставщика';
comment on column dl.sales_stock.barcode is 'Штрихкод';
comment on column dl.sales_stock.warehouse_name is 'Наименование склада';
comment on column dl.sales_stock.external_code is 'Внешний код';
comment on column dl.sales_stock.def30 is 'Дефицит на 30 дней';
comment on column dl.sales_stock.days_in_stock30 is 'Дней в наличии на 30 дней';
comment on column dl.sales_stock.def5 is 'Дефицит на 5 дней';
comment on column dl.sales_stock.days_in_stock5 is 'Дней в наличии на 5 дней';
comment on column dl.sales_stock.avg_price is 'Средняя цена';
comment on column dl.sales_stock.min_price is 'Минимальная цена';
comment on column dl.sales_stock.max_price is 'Максимальная цена';
comment on column dl.sales_stock.quantity30 is 'Количество продаж за 30 дней';
comment on column dl.sales_stock.quantity5 is 'Количество продаж за 5 дней';

CREATE TABLE "dl"."sales_stock_def" PARTITION OF "dl"."sales_stock" DEFAULT;

create
or replace procedure "dl".partition_for_sales_stock(IN start_date date, IN end_date date)
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
for v_table, v_from, v_to in SELECT 'sales_stock_' || to_char(day::date,'YYYYMM'),
                                    to_char(day::date,'YYYY-MM-DD'),
                                    to_char(day::date + interval '1 month','YYYY-MM-DD')
                             FROM generate_series(start_date, end_date, '1 month') day
            loop
            EXECUTE format(
                'CREATE TABLE IF NOT EXISTS "dl"."%s" PARTITION OF "dl"."sales_stock" FOR VALUES FROM(''%s'') TO (''%s'')',
                v_table,
                v_from,
                v_to
             );
end loop;
end;$$;

call "dl".partition_for_sales_stock(to_date('2023-01-01', 'yyyy-MM-dd'), to_date('2024-01-01', 'yyyy-MM-dd'));

create or replace procedure dl.calc_sales_stock_by_day(IN p_day date)
    language plpgsql
as
$$
declare
v_t_id  ml.transaction.id%type;
    v_count int;
    v_day   date := date_trunc('day', p_day);
begin
select count(1) into v_count from dl.sales_stock where report_date = v_day;
if v_count > 0 then
        return;
end if;
select max(id) into v_t_id from ml.transaction where date_trunc('day', "start_date") = v_day;
insert into dl.sales_stock(report_date, source, owner_code, supplier_article, barcode, warehouse_name, external_code, def30, days_in_stock30, def5, days_in_stock5, avg_price, min_price, max_price, quantity30, quantity5)
select
    sd.stock_date report_date
     ,sd.source
     ,sd.owner_code
     ,sd.supplier_article
     ,sd.barcode
     ,upper(sd.warehouse) warehouse_name
     ,sd.external_code
     ,sd.def30
     ,sd.days_in_stock30
     ,sd.def5
     ,sd.days_in_stock5
     ,sd.avg_price
     ,sd.min_price
     ,sd.max_price
     ,oo.quantity30
     ,oo.quantity5
from dl.stock_def sd
         left outer join
     (select owner_code, source, warehouse_name, external_code, sum(total_price) total_price, sum(price_with_discount) price_with_discount,
          sum(quantity) quantity30
           ,count(distinct order_date) order_date30
           ,sum(case when order_date >= v_day - INTERVAL '5 day' then quantity else 0 end) quantity5
           ,count(distinct case when order_date >= v_day - INTERVAL '10 day' then order_date end) order_date5
      from dl."order" o
      where o.transaction_id=v_t_id and o.status='delivered' and order_date between v_day - INTERVAL '30 day' and v_day
      group by owner_code, source, warehouse_name, external_code) oo on oo.external_code=sd.external_code and upper(oo.warehouse_name)=upper(sd.warehouse) and oo.owner_code=sd.owner_code and oo.source = sd.source
where sd.stock_date=v_day;
end
$$;
