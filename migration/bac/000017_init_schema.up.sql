
create table dl.stock_cluster_def(
                                     stock_date date,
                                     source varchar(5),
                                     owner_code varchar(20),
                                     supplier_article varchar(75),
                                     barcode varchar(50),
                                     cluster varchar(50),
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

CREATE TABLE "dl"."stock_def_def" PARTITION OF "dl"."stock_cluster_def" DEFAULT;

comment on table dl.stock_cluster_def is 'Stock Def by cluster';
comment on column dl.stock_cluster_def.stock_date is 'Stock Date';
comment on column dl.stock_cluster_def.source is 'Source';
comment on column dl.stock_cluster_def.owner_code is 'Owner Code';
comment on column dl.stock_cluster_def.supplier_article is 'Supplier Article';
comment on column dl.stock_cluster_def.barcode is 'Barcode';
comment on column dl.stock_cluster_def.cluster is 'Cluster';
comment on column dl.stock_cluster_def.external_code is 'External Code';
comment on column dl.stock_cluster_def.marketplace_id is 'Marketplace Id';
comment on column dl.stock_cluster_def.barcode_id is 'Barcode Id';
comment on column dl.stock_cluster_def.item_id is 'Item Id';
comment on column dl.stock_cluster_def.org_id is 'Org Id';
comment on column dl.stock_cluster_def.item_name is 'Item Name';
comment on column dl.stock_cluster_def.def30 is 'Def30';
comment on column dl.stock_cluster_def.days_in_stock30 is 'Days In Stock 30';
comment on column dl.stock_cluster_def.def5 is 'Def5';
comment on column dl.stock_cluster_def.days_in_stock5 is 'Days In Stock 5';
comment on column dl.stock_cluster_def.avg_price is 'Avg Price';
comment on column dl.stock_cluster_def.min_price is 'Min Price';
comment on column dl.stock_cluster_def.max_price is 'Max Price';

create index stock_def_stock_cluster_date_idx on dl.stock_cluster_def(stock_date);


create
or replace procedure "dl".partition_for_stock_cluster_def(IN start_date date, IN end_date date)
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
for v_table, v_from, v_to in SELECT 'stock_cluster_def_' || to_char(day::date,'YYYYMM'),
                                    to_char(day::date,'YYYY-MM-DD'),
                                    to_char(day::date + interval '1 month','YYYY-MM-DD')
                             FROM generate_series(start_date, end_date, '1 month') day
            loop
            EXECUTE format(
                'CREATE TABLE IF NOT EXISTS "dl"."%s" PARTITION OF "dl"."stock_cluster_def" FOR VALUES FROM(''%s'') TO (''%s'')',
                v_table,
                v_from,
                v_to
             );
end loop;
end;$$;

call "dl".partition_for_stock_cluster_def(to_date('2023-01-01', 'yyyy-MM-dd'), to_date('2024-01-01', 'yyyy-MM-dd'));

create or replace procedure dl.calc_stock_cluster_def_by_day(IN p_day date)
    language plpgsql
as
$$
declare
v_count int;
    v_day   date := date_trunc('day', p_day);
begin
select count(1) into v_count from dl.stock_cluster_def where stock_date = v_day;
if v_count > 0 then
        return;
end if;
insert into dl.stock_cluster_def(stock_date, source, owner_code, supplier_article, barcode, cluster, external_code,
                                 marketplace_id, barcode_id, item_id, org_id, item_name, def30, days_in_stock30,
                                 def5, days_in_stock5, avg_price, min_price, max_price)
select v_day
     , sd.source
     , owner_code
     , supplier_article
     , sd.barcode
     , wh.cluster
     , sd.external_code
     , mp.marketplace_id
     , b.barcode_id
     , i.id                                                                                                item_id
     , org.id                                                                                              org_id
     , i.name
     , 30 - sum(case when sd.quantity > 0 then 1 else 0 end)                                               def30
     , count(distinct sd.stock_date)                                                                       days_in_stock_30
     , 5 - sum(case when (sd.stock_date > v_day - INTERVAL '5 day' and sd.quantity > 0) then 1 else 0 end) def5
     , count(distinct case when (sd.stock_date > v_day - INTERVAL '5 day') then sd.stock_date end)         days_in_stock_5
     , sum(sd.price) / count(sd.stock_date)                                                                avg_price
     , min(sd.price)                                                                                       min_price
     , max(sd.price)                                                                                       max_price
from dl.stock_daily sd
         left outer join dl.warehouse wh on wh.source = sd.source and wh.code = sd.warehouse
         left outer join ml.owner o on o.code = sd.owner_code
         left outer join dl.organisation org on org.id = o.organisation_id
         left outer join ml.marketplace mp on mp.code = sd.source
         left outer join dl.barcode b on b.barcode = sd.barcode and b.organisation_id = org.id and
                                         b.marketplace_id = mp.marketplace_id
         left outer join dl.item i on i.id = b.item_id
where sd.stock_date between v_day - INTERVAL '30 day' and v_day
group by sd.source, owner_code, supplier_article, sd.barcode, wh.cluster, sd.external_code, mp.marketplace_id,
    b.barcode_id, i.id, org.id, i.name;
end
$$;