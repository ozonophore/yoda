
create table dl.stock_item_def(
                                  stock_date date,
                                  source varchar(5),
                                  owner_code varchar(20),
                                  supplier_article varchar(75),
                                  barcode varchar(50),
                                  external_code varchar(50),
                                  marketplace_id varchar(36),
                                  barcode_id varchar(36),
                                  item_id varchar(36),
                                  org_id varchar(36),
                                  item_name varchar(255),
                                  def30 numeric(10,0),
                                  days_in_stock30 numeric(10,0),
                                  def5 numeric(10,0),
                                  days_in_stock5 numeric(10,0),
                                  avg_price numeric(10,2),
                                  min_price numeric(10,2),
                                  max_price numeric(10,2)
) partition by range (stock_date);

CREATE TABLE "dl"."stock_item_def_def" PARTITION OF "dl"."stock_item_def" DEFAULT;

comment on table dl.stock_item_def is 'Stock Def by cluster';
comment on column dl.stock_item_def.stock_date is 'Stock Date';
comment on column dl.stock_item_def.source is 'Source';
comment on column dl.stock_item_def.owner_code is 'Owner Code';
comment on column dl.stock_item_def.supplier_article is 'Supplier Article';
comment on column dl.stock_item_def.barcode is 'Barcode';
comment on column dl.stock_item_def.external_code is 'External Code';
comment on column dl.stock_item_def.marketplace_id is 'Marketplace Id';
comment on column dl.stock_item_def.barcode_id is 'Barcode Id';
comment on column dl.stock_item_def.item_id is 'Item Id';
comment on column dl.stock_item_def.org_id is 'Org Id';
comment on column dl.stock_item_def.item_name is 'Item Name';
comment on column dl.stock_item_def.def30 is 'Def30';
comment on column dl.stock_item_def.days_in_stock30 is 'Days In Stock 30';
comment on column dl.stock_item_def.def5 is 'Def5';
comment on column dl.stock_item_def.days_in_stock5 is 'Days In Stock 5';
comment on column dl.stock_item_def.avg_price is 'Avg Price';
comment on column dl.stock_item_def.min_price is 'Min Price';
comment on column dl.stock_item_def.max_price is 'Max Price';

create index stock_item_date_idx on dl.stock_item_def(stock_date);


create
or replace procedure "dl".partition_for_stock_item_def(IN start_date date, IN end_date date)
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
for v_table, v_from, v_to in SELECT 'stock_item_def_' || to_char(day::date,'YYYYMM'),
                                    to_char(day::date,'YYYY-MM-DD'),
                                    to_char(day::date + interval '1 month','YYYY-MM-DD')
                             FROM generate_series(start_date, end_date, '1 month') day
            loop
            EXECUTE format(
                'CREATE TABLE IF NOT EXISTS "dl"."%s" PARTITION OF "dl"."stock_item_def" FOR VALUES FROM(''%s'') TO (''%s'')',
                v_table,
                v_from,
                v_to
             );
end loop;
end;$$;

call "dl".partition_for_stock_item_def(to_date('2023-01-01', 'yyyy-MM-dd'), to_date('2024-01-01', 'yyyy-MM-dd'));

create or replace procedure dl.calc_stock_item_def_by_day(IN p_day date)
    language plpgsql
as
$$
declare
v_count int;
    v_day   date := date_trunc('day', p_day);
begin
select count(1) into v_count from dl.stock_item_def where stock_date = v_day;
if v_count > 0 then
            return;
end if;
insert into dl.stock_item_def(stock_date, source, owner_code, supplier_article, barcode, external_code,
                              marketplace_id, barcode_id, item_id, org_id, item_name, def30, days_in_stock30,
                              def5, days_in_stock5, avg_price, min_price, max_price)
select v_day, ss.source, ss.owner_code, ss.supplier_article, ss.barcode, ss.external_code,
    mp.marketplace_id, b.barcode_id
     , i.id item_id, org.id org_id, i.name
     ,30 - sum(case when ss.quantity30 > 0 then 1 else 0 end ) quantity30
     ,count(ss.stock_date) days_in_stock30
     ,5 - sum(case when ss.quantity5 > 0 then 1 else 0 end ) quantity5
     ,count(distinct case when ss.stock_date > v_day - INTERVAL '5 day' then ss.stock_date end) days_in_stock5
     ,sum(ss.avg_price) / count(ss.stock_date) avg_price
     ,min(ss.min_price) min_price
     ,max(ss.max_price) max_price
from (select sd.stock_date
           , source
           , owner_code
           , supplier_article
           , sd.barcode
           , sd.external_code
           , sum(case when sd.quantity > 0 then 1 else 0 end) quantity30
           , sum(case
                     when (sd.stock_date > v_day - INTERVAL '5 day') then sd.quantity
                     else 0 end)                              quantity5
           , sum(sd.price) / count(sd.stock_date)             avg_price
           , sum(sd.price)
           , count(sd.stock_date)
           , min(sd.price)                                    min_price
           , max(sd.price)                                    max_price
      from dl.stock_daily sd
      where sd.stock_date > v_day - INTERVAL '30 day'
        and sd.stock_date <= v_day
      group by stock_date, source, owner_code, supplier_article, sd.barcode, sd.external_code) ss
         inner join ml.marketplace mp on mp.code = ss.source
         inner join ml.owner o on o.code = ss.owner_code
         inner join dl.organisation org on org.id = o.organisation_id
         left outer join dl.barcode b on b.barcode = ss.barcode and b.organisation_id = org.id and
                                         b.marketplace_id = mp.marketplace_id
         left outer join dl.item i on i.id = b.item_id
group by ss.source, ss.owner_code, ss.supplier_article, ss.barcode, ss.external_code,
         mp.marketplace_id, b.barcode_id
        , i.id, org.id, i.name;
end
$$;