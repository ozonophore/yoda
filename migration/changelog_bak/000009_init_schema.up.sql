create or replace procedure calc_stock_daily(IN p_id integer)
    language plpgsql
as
$$
declare
v_status "transaction"."status"%type;
begin
select "status" into v_status from "transaction" where "id" = p_id;
if v_status not in ('COMPLETED', 'BEGIN') then
        raise exception 'Transaction status is not completed';
end if;

merge into "stock_daily" ss
    using (select "transaction_date",
                  "source",
                  "owner_code",
                  "warehouse_name",
                  "external_code",
                  "barcode",
                  sum("quantity") quantity,
                  sum("quantity_full") quantity_full,
                  case when sum("quantity") = 0 then 0 else sum("price") / sum("quantity") end as "price",
                  case
                      when sum("quantity") = 0 then 0
                      else sum("price_after_discount") / sum("quantity") end as "price_with_discount"
           from "stock"
           where "transaction_id" = p_id and "price" is not null
           group by "source", "owner_code", "warehouse_name", "external_code", "barcode",
                    "transaction_date") as data
    on (ss."source" = data."source"
        and ss."owner_code" = data."owner_code"
        and ss."warehouse" = data."warehouse_name"
        and ss."external_code" = data."external_code"
        and ss."stock_date" = data."transaction_date")
    when matched then
        update
            set "attemption"           = "attemption" + 1,
                "quantity"            = (ss."quantity" * ss."attemption" + data."quantity") / (ss."attemption" + 1),
                "quantity_full"       = (ss."quantity_full" * ss."attemption" + data."quantity") / (ss."attemption" + 1),
                "price"               = case when (ss."quantity" + data."quantity") = 0 then 0 else (ss."price" * ss."quantity" + data."price") / (ss."quantity" + data."quantity") end,
                "price_with_discount" = case when (ss."quantity" + data."quantity") = 0 then 0 else (ss."price_with_discount" * ss."quantity" + data."price_with_discount") /
                                                                                                    (ss."quantity" + data."quantity") end,
                "update_at"           = current_date
    when not matched then
        insert ("stock_date", "source", "owner_code", "warehouse", "external_code", "barcode", "quantity",
                "quantity_full", "price", "price_with_discount", "create_at", "update_at", "attemption")
            values (data."transaction_date", data."source", data."owner_code", data."warehouse_name", data."external_code",
                    data."barcode", data."quantity", data."quantity_full", data."price", data."price_with_discount", current_date, current_date, 1);
end
$$;
