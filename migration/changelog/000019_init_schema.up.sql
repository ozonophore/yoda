create table dl.report_by_item
(
    report_date      date,
    source           varchar(5),
    owner_code       varchar(20),
    supplier_article varchar(75),
    barcode          varchar(50),
    external_code    varchar(50),
    item_id          varchar(36),
    item_name        varchar(255),
    marketplace_id   varchar(36),
    org_id           varchar(36),
    def30            numeric(10),
    days_in_stock30  numeric(10),
    def5             numeric(10),
    days_in_stock5   numeric(10),
    avg_price        numeric(10, 2),
    min_price        numeric(10, 2),
    max_price        numeric(10, 2),
    quantity30       numeric(10),
    quantity5        numeric(10),
    order_by_day30   numeric(10, 2),
    forecast_order30 numeric(10, 2),
    order_by_day5    numeric(10, 2),
    forecast_order5  numeric(10, 2),
    quantity         numeric(10),
    is_excluded      boolean
) partition by RANGE (report_date);

comment on table dl.report_by_item is 'Отчет по кластерам';
comment on column dl.report_by_item.report_date is 'Дата отчета';
comment on column dl.report_by_item.source is 'Источник';
comment on column dl.report_by_item.owner_code is 'Код владельца';
comment on column dl.report_by_item.supplier_article is 'Артикул поставщика';
comment on column dl.report_by_item.barcode is 'Штрихкод';
comment on column dl.report_by_item.external_code is 'Внешний код';
comment on column dl.report_by_item.item_id is 'Идентификатор товара';
comment on column dl.report_by_item.item_name is 'Наименование товара';
comment on column dl.report_by_item.marketplace_id is 'Идентификатор маркетплейса';
comment on column dl.report_by_item.org_id is 'Идентификатор организации';
comment on column dl.report_by_item.def30 is 'Дефицит за 30 дней';
comment on column dl.report_by_item.days_in_stock30 is 'Дней в наличии за 30 дней';
comment on column dl.report_by_item.def5 is 'Дефицит за 5 дней';
comment on column dl.report_by_item.days_in_stock5 is 'Дней в наличии за 5 дней';
comment on column dl.report_by_item.avg_price is 'Средняя цена';
comment on column dl.report_by_item.min_price is 'Минимальная цена';
comment on column dl.report_by_item.max_price is 'Максимальная цена';
comment on column dl.report_by_item.quantity30 is 'Количество за 30 дней';
comment on column dl.report_by_item.quantity5 is 'Количество за 5 дней';
comment on column dl.report_by_item.order_by_day30 is 'Заказов в день за 30 дней';
comment on column dl.report_by_item.forecast_order30 is 'Прогноз заказов за 30 дней';
comment on column dl.report_by_item.order_by_day5 is 'Заказов в день за 5 дней';
comment on column dl.report_by_item.forecast_order5 is 'Прогноз заказов за 5 дней';
comment on column dl.report_by_item.quantity is 'Количество';
comment on column dl.report_by_item.is_excluded is 'Исключен';


CREATE TABLE "dl"."report_by_item_def" PARTITION OF "dl"."report_by_item" DEFAULT;

create
or replace procedure "dl".partition_for_report_by_item(IN start_date date, IN end_date date)
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
for v_table, v_from, v_to in SELECT 'report_by_item_' || to_char(day::date,'YYYYMM'),
                                    to_char(day::date,'YYYY-MM-DD'),
                                    to_char(day::date + interval '1 month','YYYY-MM-DD')
                             FROM generate_series(start_date, end_date, '1 month') day
            loop
            EXECUTE format(
                'CREATE TABLE IF NOT EXISTS "dl"."%s" PARTITION OF "dl"."report_by_item" FOR VALUES FROM(''%s'') TO (''%s'')',
                v_table,
                v_from,
                v_to
             );
end loop;
end;$$;

call "dl".partition_for_report_by_item(to_date('2023-01-01', 'yyyy-MM-dd'), to_date('2024-01-01', 'yyyy-MM-dd'));

create procedure dl.calc_report_by_item(IN p_day date)
    language plpgsql
as
$$
declare
v_t_id  ml.transaction.id%type;
    v_count int;
    v_day   date := date_trunc('day', p_day);
begin
select count(1) into v_count from dl.report_by_item where report_date = v_day;
if v_count > 0 then
        return;
end if;
select max(id) into v_t_id from ml.transaction where date_trunc('day', "start_date") = v_day;
insert into dl.report_by_item(report_date, source, owner_code, supplier_article, barcode, external_code,
                              def30, days_in_stock30, def5, days_in_stock5, avg_price, min_price, max_price, quantity30,
                              quantity5, order_by_day30, forecast_order30, order_by_day5, forecast_order5,quantity
    ,item_id, item_name, marketplace_id, org_id, is_excluded)

select
    sd.stock_date report_date
     ,sd.source
     ,sd.owner_code
     ,sd.supplier_article
     ,sd.barcode
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
from dl.stock_item_def sd
         left outer join (select o.code owner_code, m.code source, item_id, barcode_id, barcode, b.organisation_id,b.marketplace_id, i.name from dl.barcode b
                                                                                                                                                     inner join dl.item i on i.id = b.item_id
                                                                                                                                                     inner join ml.owner o on o.organisation_id = b.organisation_id
                                                                                                                                                     inner join ml.marketplace m on m.marketplace_id = b.marketplace_id) s
                         on s.source = sd.source and s.owner_code = sd.owner_code and s.barcode = sd.barcode
         left outer join (
    select
        owner_code, source, external_code, sum(quantity) quantity
    from dl.stock_daily sdd
    where stock_date=v_day
    group by owner_code, source, external_code
) sdd on sdd.owner_code = sd.owner_code and sdd.source = sd.source and sdd.external_code = sd.external_code
         left outer join
     (select owner_code, source, external_code, sum(total_price) total_price, sum(price_with_discount) price_with_discount,
          sum(quantity) quantity30
           ,count(distinct order_date) order_date30
           ,sum(case when order_date > v_day - INTERVAL '5 day' then quantity else 0 end) quantity5
           ,count(distinct case when order_date > v_day - INTERVAL '5 day' then order_date end) order_date5
      from dl."order" o
      where o.transaction_id=v_t_id and o.is_cancel != true
        and order_date > v_day - INTERVAL '30 day' and order_date <= v_day
      group by owner_code, source, external_code) oo on oo.external_code=sd.external_code and oo.owner_code=sd.owner_code and oo.source = sd.source
         left outer join dl.exclud_item ei on ei.source = sd.source and ei.owner_code = sd.owner_code and ei.barcode = sd.barcode
where sd.stock_date=v_day;
end
$$;
