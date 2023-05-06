create table stock_daily
(
    "stock_date" date not null,
    "source" varchar(20) not null,
    "owner_code" varchar(20) not null,
    "supplier_article" varchar(75),
    "barcode" varchar(30),
    "external_code" varchar(50) not null,
    "name" varchar(200),
    "subject" varchar(200),
    "category" varchar(50),
    "brand" varchar(50),
    "warehouse" varchar(50) not null,
    "create_at" date not null,
    "update_at" date,
    "quantity" numeric(10,2) not null,
    "quantity_full" numeric(10,2) not null,
    "attention" int not null,
    "price" numeric(10,2) not null,
    "price_with_discount" numeric(10,2) not null
) PARTITION BY RANGE("stock_date");

CREATE TABLE stock_daily_def PARTITION OF stock_daily DEFAULT;
create index stock_daily_def_soec_idx on
    stock_daily_def ("source", "owner_code", "warehouse", "external_code");

CREATE TABLE stock_daily_202304 PARTITION OF stock_daily
    FOR VALUES FROM
(
    '2023-04-01'
) TO
(
    '2023-05-01'
);
create index stock_daily_202304_soec_idx on
    stock_daily_202304 ("source", "owner_code", "warehouse", "external_code");

CREATE TABLE stock_daily_202305 PARTITION OF stock_daily
    FOR VALUES FROM
(
    '2023-05-01'
) TO
(
    '2023-06-01'
);
create index stock_daily_202305_soec_idx on
    stock_daily_202305 ("source", "owner_code", "warehouse", "external_code");

CREATE TABLE stock_daily_202306 PARTITION OF stock_daily
    FOR VALUES FROM
(
    '2023-06-01'
) TO
(
    '2023-07-01'
);
create index stock_daily_202306_soec_idx on
    stock_daily_202306 ("source", "owner_code", "warehouse", "external_code");

CREATE TABLE stock_daily_202307 PARTITION OF stock_daily
    FOR VALUES FROM
(
    '2023-07-01'
) TO
(
    '2023-08-01'
);
create index stock_daily_202307_soec_idx on
    stock_daily_202307 ("source", "owner_code", "warehouse", "external_code");

CREATE TABLE stock_daily_202308 PARTITION OF stock_daily
    FOR VALUES FROM
(
    '2023-08-01'
) TO
(
    '2023-09-01'
);
create index stock_daily_202308_soec_idx on
    stock_daily_202308 ("source", "owner_code", "warehouse", "external_code");

CREATE TABLE stock_daily_202309 PARTITION OF stock_daily
    FOR VALUES FROM
(
    '2023-09-01'
) TO
(
    '2023-10-01'
);
create index stock_daily_202309_soec_idx on
    stock_daily_202309 ("source", "owner_code", "warehouse", "external_code");

CREATE TABLE stock_daily_202310 PARTITION OF stock_daily
    FOR VALUES FROM
(
    '2023-10-01'
) TO
(
    '2023-11-01'
);
create index stock_daily_202310_soec_idx on
    stock_daily_202310 ("source", "owner_code", "warehouse", "external_code");

CREATE TABLE stock_daily_202311 PARTITION OF stock_daily
    FOR VALUES FROM
(
    '2023-11-01'
) TO
(
    '2023-12-01'
);
create index stock_daily_202311_soec_idx on
    stock_daily_202311 ("source", "owner_code", "warehouse", "external_code");

CREATE TABLE stock_daily_202312 PARTITION OF stock_daily
    FOR VALUES FROM
(
    '2023-12-01'
) TO
(
    '2024-01-01'
);
create index stock_daily_202312_soec_idx on
    stock_daily_202312 ("source", "owner_code", "warehouse", "external_code");

create index stock_soec_idx on "stock" ("transaction_id", "source", "owner_code", "warehouse_name", "external_code");
