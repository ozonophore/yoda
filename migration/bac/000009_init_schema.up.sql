create table dl."report_detail_by_period"
(
    id                          serial      not null primary key,
    transaction_id              integer     not null,
    source                      varchar(50) not null,
    owner_code                  varchar(50) not null,
    acquiring_bank              varchar(100),
    acquiring_fee               numeric(10, 2),
    additional_payment          numeric(10, 2),
    barcode                     varchar(30),
    bonus_type_name             varchar(100),
    brand_name                  varchar(100),
    commission_percent          numeric(18),
    create_dt                   date,
    date_from                   date,
    date_to                     date,
    declaration_number          varchar(100),
    delivery_amount             numeric(18),
    delivery_rub                numeric(10, 2),
    doc_number                  varchar(50),
    doc_type_name               varchar(100),
    gi_box_type_name            varchar(100),
    gi_id                       numeric(18),
    kiz                         varchar(100),
    nm_id                       numeric(18),
    office_name                 varchar(100),
    order_dt                    date,
    penalty                     numeric(10, 2),
    ppvz_for_pay                numeric(10, 2),
    ppvz_inn                    varchar(20),
    ppvz_kvw_prc                numeric(10, 2),
    ppvz_kvw_prc_base           numeric(10, 2),
    ppvz_office_id              numeric(18),
    ppvz_office_name            varchar(100),
    ppvz_reward                 numeric(10, 2),
    ppvz_sales_commission       numeric(10, 2),
    ppvz_spp_prc                numeric(10, 2),
    ppvz_supplier_id            numeric(18),
    ppvz_supplier_name          varchar(100),
    ppvz_vw                     numeric(10, 2),
    ppvz_vw_nds                 numeric(10, 2),
    product_discount_for_report numeric(10, 2),
    quantity                    numeric(18),
    realizationreport_id        numeric(18),
    retail_amount               numeric(10, 2),
    retail_price                numeric(10, 2),
    retail_price_withdisc_rub   numeric(10, 2),
    return_amount               numeric(18),
    rid                         varchar(50),
    rr_dt                       date,
    rrd_id                      numeric(18),
    sa_name                     varchar(100),
    sale_dt                     date,
    sale_percent                numeric(18),
    shk_id                      numeric(18),
    site_country                varchar(100),
    srid                        varchar(100),
    sticker_id                  varchar(50),
    subject_name                varchar(200),
    supplier_oper_name          varchar(200),
    supplier_promo              numeric(10, 2),
    suppliercontract_code       varchar(100),
    ts_name                     varchar(100),
    item_id                     varchar(36),
    barcode_id                  varchar(36),
    message                     varchar(250)
);

create index report_detail_by_period_transaction_id_idx on "dl"."report_detail_by_period"(transaction_id, source, owner_code, barcode);