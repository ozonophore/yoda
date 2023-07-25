create table dl.sales_stock(
                               report_date date,
                               source varchar(5),
                               owner_code varchar(20),
                               supplier_article varchar(75),
                               barcode varchar(50),
                               warehouse_name varchar(50),
                               external_code varchar(50),
                               item_id varchar(36),
                               item_name varchar(255),
                               marketplace_id varchar(36),
                               org_id varchar(36),
                               def30 numeric(10),
                               days_in_stock30 numeric(10),
                               def5 numeric(10),
                               days_in_stock5 numeric(10),
                               avg_price numeric(10,2),
                               min_price numeric(10,2),
                               max_price numeric(10,2),
                               quantity30 numeric(10),
                               quantity5 numeric(10),
                               order_by_day30 numeric(10,2),
                               forecast_order30 numeric(10,2),
                               order_by_day5 numeric(10,2),
                               forecast_order5 numeric(10,2),
                               is_excluded boolean default false
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
comment on column dl.sales_stock.order_by_day30 is 'Заказов в день за 30 дней';
comment on column dl.sales_stock.forecast_order30 is 'Прогноз заказов за 30 дней';
comment on column dl.sales_stock.order_by_day5 is 'Заказов в день за 5 дней';
comment on column dl.sales_stock.forecast_order5 is 'Прогноз заказов за 5 дней';

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

create procedure calc_sales_stock_by_day(IN p_day date)
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
insert into dl.sales_stock(report_date, source, owner_code, supplier_article, barcode, warehouse_name, external_code, def30, days_in_stock30, def5, days_in_stock5, avg_price, min_price, max_price, quantity30,
                           quantity5, order_by_day30, forecast_order30, order_by_day5, forecast_order5,quantity
    ,item_id, item_name, marketplace_id, org_id, is_excluded)
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
     ,case when sd.def30 = 30 then 0 else oo.quantity30 / (30 - sd.def30) end order_by_day30
     ,case when sd.def30 = 30 then 0 else (oo.quantity30 * 30) / (30 - sd.def30) end forecast_order30
     ,case when sd.def5 = 5 then 0 else oo.quantity5 / (5 - sd.def5)  end order_by_day5
     ,case when sd.def5 = 5 then 0 else (oo.quantity5 * 5) / (5 - sd.def5)  end forecast_order5
     ,sdd.quantity quantity
     ,s.item_id item_id
     ,s.name item_name
     ,s.marketplace_id marketplace_id
     ,s.organisation_id org_id
     ,case when ei.barcode is null then false else true end is_exclud
from dl.stock_def sd
    left outer join (select o.code owner_code, m.code source, item_id, barcode_id, barcode, b.organisation_id,b.marketplace_id, i.name from dl.barcode b
    inner join dl.item i on i.id = b.item_id
    inner join ml.owner o on o.organisation_id = b.organisation_id
    inner join ml.marketplace m on m.marketplace_id = b.marketplace_id) s on s.source = sd.source and s.owner_code = sd.owner_code and s.barcode = sd.barcode
    left outer join dl.stock_daily sdd on sdd.warehouse = sd.warehouse and sdd.owner_code = sd.owner_code and sdd.source = sd.source and sdd.external_code = sd.external_code and sdd.stock_date = sd.stock_date
    left outer join
    (select owner_code, source, warehouse_name, external_code, sum(total_price) total_price, sum(price_with_discount) price_with_discount,
    sum(quantity) quantity30
    ,count(distinct order_date) order_date30
    ,sum(case when order_date > v_day - INTERVAL '5 day' then quantity else 0 end) quantity5
    ,count(distinct case when order_date > v_day - INTERVAL '5 day' then order_date end) order_date5
    from dl."order" o
    where o.transaction_id=v_t_id and o.is_cancel != true and order_date > v_day - INTERVAL '30 day' and order_date <= v_day
    group by owner_code, source, warehouse_name, external_code) oo on oo.external_code=sd.external_code and upper(oo.warehouse_name)=upper(sd.warehouse) and oo.owner_code=sd.owner_code and oo.source = sd.source
    left outer join dl.exclud_item ei on ei.source = sd.source and ei.owner_code = sd.owner_code and ei.barcode = sd.barcode
where sd.stock_date=v_day;
end
$$;

create index stock_daily_dwosec on dl.stock_daily(stock_date,source,owner_code,warehouse,external_code);

create table ml.settings
(
    code varchar(50) not null,
    value varchar(100),
    description varchar(100)
);

create table dl.exclud_item
(
    name varchar(100),
    source varchar(5),
    org_name text,
    barcode varchar(20),
    article varchar(20),
    owner_code varchar(20)
);

alter table exclud_item owner to "user";

create unique index idx_exclud_item
    on exclud_item (source, owner_code, barcode);

create table dl.warehouse (
  code varchar(50) not null primary key,
  cluster varchar(50) not null,
  source varchar(10) not null
);