create procedure calc_order_delivered(IN p_id integer)
    language plpgsql
as
$$
declare
v_id "transaction".id%type;
v_count integer;
begin
select coalesce(max(transaction_id), 0) into v_id from "order_delivered_log";
if p_id <= v_id then
        raise exception 'transaction_id % is already processed', p_id;
end if;
---- CALCULATE DELIVERED ORDERS ARCH
insert into "order_delivered_arch" ("transaction_id",
                                    "order_date",
                                    "owner_code",
                                    "source",
                                    "supplier_article",
                                    "warehouse",
                                    "barcode",
                                    "external_code",
                                    "quantity",
                                    "total_price",
                                    "price_with_discount")
SELECT p_id,
       o."order_date",
       o."owner_code",
       o."source",
       o."supplier_article",
       o."warehouse_name",
       o."barcode",
       o."external_code",
       sum(o."quantity") + coalesce(sum(oda."quantity"), 0) as "quantity",
       (sum(coalesce(oda."total_price", 0) * coalesce(oda."quantity", 0)) + sum(o."total_price")) /
       (sum(coalesce(oda.quantity, 0)) + sum(o.quantity))   as "total_price",
       (sum(coalesce(oda."price_with_discount", 0) * coalesce(oda."quantity", 0)) + sum(o."price_with_discount")) /
       (sum(coalesce(oda.quantity, 0)) + sum(o.quantity))   as "price_with_discount"
FROM "order" o
         left outer join "order_delivered_arch" oda on
            oda."order_date" = o."order_date"
        and oda."order_date" <= current_date
        and oda."order_date" > current_date - 25
        and oda."owner_code" = o."owner_code"
        and oda."source" = o."source"
        and oda."warehouse" = o."warehouse_name"
        and oda."external_code" = o."external_code"
        and oda."transaction_id" = v_id
WHERE o."transaction_id" = p_id
  AND o."status" = 'delivered'
  AND NOT EXISTS(SELECT *
                 FROM "order" o2
                 WHERE o2."transaction_id" = v_id
                   AND o2."source" = o."source"
                   and o2."owner_code" = o."owner_code"
                   AND o2."status" = o."status"
                   and o2."warehouse_name" = o."warehouse_name"
                   and o2."external_code" = o."external_code"
                   AND o2."srid" = o."srid")
group by o."order_date",
         o."owner_code",
         o."source",
         o."supplier_article",
         o."warehouse_name",
         o."barcode",
         o."external_code";

--- CALCULATE DELIVERED ORDERS
MERGE INTO "order_delivered" AS od
    USING (SELECT o."order_date",
                  o."owner_code",
                  o."source",
                  o."supplier_article",
                  o."warehouse_name",
                  o."barcode",
                  o."external_code",
                  sum(o."quantity")                                as "quantity",
                  sum(o."total_price")         as "sum_total_price",
                  sum(o."price_with_discount") as "sum_price_with_discount"
           FROM "order" o
           WHERE o."transaction_id" = p_id
             AND o."status" = 'delivered'
             AND NOT EXISTS(SELECT *
                            FROM "order" o2
                            WHERE o2."transaction_id" = v_id
                              AND o2."srid" = o."srid"
                              AND o2."source" = o."source"
                              AND o2."status" = o."status"
                              AND o2."owner_code" = o."owner_code"
                              and o2."warehouse_name" = o."warehouse_name"
                              and o2."external_code" = o."external_code"
               )
           GROUP BY o."order_date",
                    o."owner_code",
                    o."source",
                    o."supplier_article",
                    o."warehouse_name",
                    o."barcode",
                    o."external_code") AS data
    ON (od."order_date" = data."order_date" AND od."owner_code" = data."owner_code" AND od."source" = data."source" AND
        od."supplier_article" = data."supplier_article" AND od."warehouse" = data."warehouse_name" AND
        od."external_code" = data."external_code")
    WHEN MATCHED THEN
        UPDATE
            SET "quantity" = data."quantity" + od."quantity",
                "total_price" = (od."total_price" * od."quantity" + data."sum_total_price") / (data."quantity" + od."quantity"),
                "price_with_discount" = (od."price_with_discount" * od."quantity" + data."sum_price_with_discount") / (data."quantity" + od."quantity")
    WHEN NOT MATCHED THEN
        INSERT ("order_date", "owner_code", "source", "supplier_article", "warehouse", "barcode", "external_code",
                "quantity",
                "total_price",
                "price_with_discount")
            VALUES (data."order_date", data."owner_code", data."source", data."supplier_article", data."warehouse_name",
                    data."barcode", data."external_code", data."quantity", data."sum_total_price" / data."quantity",
                    data."sum_price_with_discount" / data."quantity");

GET DIAGNOSTICS v_count = ROW_COUNT;

insert into "order_delivered_log" ("transaction_id", "created_at", "added_rows")
values (p_id, now(), v_count);
end;
$$;
