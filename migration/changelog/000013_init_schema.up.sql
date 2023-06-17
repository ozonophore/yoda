create or replace procedure dl.calc_stock_def30_by_day(IN p_day date)
    language plpgsql
as
$$
declare
v_count integer;
    v_day   date := date_trunc('day', p_day);
begin
    select count(1) into v_count from dl.stock_def30 where stock_date = v_day;
    if v_count > 0 then
            return;
    end if;
    insert into dl.stock_def30(stock_date, source, owner_code, supplier_article, barcode, warehouse, external_code, marketplace_id, barcode_id, item_id, org_id, item_name, def30,  days_in_stock, avg_price, min_price, max_price)
    select v_day, source, owner_code, supplier_article, sd.barcode, warehouse, sd.external_code,
        mp.marketplace_id, b.barcode_id
         , i.id item_id, org.id org_id, i.name
         ,sum(case when sd.quantity = 0 then 1 else 0 end ) def30
         ,count(distinct sd.stock_date) days_in_stock
         ,sum(sd.price) / count(sd.stock_date) avg_price
         ,min(sd.price) min_price
         ,max(sd.price) max_price
    from dl.stock_daily sd
             left outer join ml.owner o on o.code = sd.owner_code
             left outer join dl.organisation org on org.id = o.organisation_id
             left outer join ml.marketplace mp on mp.code = sd.source
             left outer join dl.barcode b on b.barcode = sd.barcode and b.organisation_id = org.id and b.marketplace_id = mp.marketplace_id
             left outer join dl.item i on i.id = b.item_id
    where sd.stock_date between v_day - INTERVAL '30 day' and v_day
    group by source, owner_code, supplier_article, sd.barcode, warehouse, sd.external_code, mp.marketplace_id, b.barcode_id, i.id, org.id, i.name;
end
$$;

alter procedure dl.calc_stock_def30_by_day(date) owner to "user";