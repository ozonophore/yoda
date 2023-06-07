create or replace procedure dl.calc_stock_daily_by_day(IN p_day date)
    language plpgsql
as
$$
declare
v_count integer;
    v_day   date := date_trunc('day', p_day);
begin
select count(*) into v_count from "ml"."transaction"
where date_trunc('day', "start_date") = v_day
  and "status" = 'COMPLETED';
insert into dl.stock_daily(stock_date, source, owner_code, supplier_article, barcode, external_code, warehouse, quantity, quantity_full, attemption, price, price_with_discount, create_at, update_at)
select s.transaction_date
     , s.source
     , s.owner_code
     , s.supplier_article
     , s.barcode
     , s.external_code
     , s.warehouse_name
     , sum(s.quantity) / v_count                              quantity
     , sum(s.quantity_full) / v_count                          quantity_full
     , v_count                                                attemption
     , case
           when sum(s.quantity) / v_count = 0
               then sum(coalesce(s.price, 0)) / count(1)
           else sum(coalesce(s.price, 0)) / (sum(s.quantity) / v_count) end price
     , case
           when sum(s.quantity) / v_count = 0
               then sum(coalesce(s.price_after_discount, 0))
           else sum(coalesce(s.price_after_discount, 0)) /
                (sum(s.quantity) / v_count) end                price_after_discount
     , now()                                                                    create_at
     , now()                                                                    update_at
from dl.stock s
         inner join ml.transaction t on t.id = s.transaction_id
where s.transaction_id in (select id from ml.transaction t
                           where date_trunc('day', t.start_date) = v_day
                             and t.status = 'COMPLETED')
group by s.transaction_date, s.source, s.owner_code, s.supplier_article, s.barcode,
         s.barcode_id, s.item_id, s.warehouse_name, s.external_code;
end
$$;
