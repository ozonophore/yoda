create sequence ml.transaction_id_seq
    as integer
    start with 90;

alter sequence ml.transaction_id_seq owner to "user";

create sequence ml.log_id_seq;

alter sequence ml.log_id_seq owner to "user";

create table if not exists dl.barcode
(
    item_id         varchar(36)  not null,
    barcode_id      varchar(36)  not null
    primary key,
    barcode         varchar(255) not null,
    organisation_id varchar(36)  not null,
    marketplace_id  varchar(36)  not null,
    updated_at      timestamp    not null
    );

comment on table dl.barcode is 'Barcodes';

comment on column dl.barcode.item_id is 'Item ID';

comment on column dl.barcode.barcode_id is 'ID';

comment on column dl.barcode.barcode is 'Barcode';

comment on column dl.barcode.organisation_id is 'Organisation ID';

comment on column dl.barcode.marketplace_id is 'Marketplace ID';

alter table dl.barcode
    owner to "user";

grant select on dl.barcode to readonly;

create table if not exists dl.item
(
    id        varchar(36)  not null
    primary key,
    name      varchar(200) not null,
    update_at timestamp
    );

comment on table dl.item is 'Таблица для хранения справочника товаров';

comment on column dl.item.id is 'Идентификатор';

comment on column dl.item.name is 'Наименование';

comment on column dl.item.update_at is 'Дата обновления';

alter table dl.item
    owner to "user";

grant select on dl.item to readonly;

create table if not exists dl.marketplace
(
    id         varchar(36)  not null
    primary key,
    name       varchar(255) not null,
    updated_at timestamp    not null
    );

comment on table dl.marketplace is 'Marketplaces';

comment on column dl.marketplace.id is 'ID';

comment on column dl.marketplace.name is 'Name';

comment on column dl.marketplace.updated_at is 'Updated at';

alter table dl.marketplace
    owner to "user";

grant select on dl.marketplace to readonly;

create table if not exists dl."order"
(
    id                  serial,
    transaction_date    date        not null,
    transaction_id      integer     not null,
    source              varchar(20) not null,
    owner_code          varchar(20) not null,
    last_change_date    date        not null,
    last_change_time    time        not null,
    order_date          date        not null,
    order_time          time        not null,
    supplier_article    varchar(75),
    tech_size           varchar(30),
    barcode             varchar(30),
    total_price         numeric(18, 2),
    discount_percent    numeric(18, 2),
    discount_value      numeric(18, 2),
    price_with_discount numeric(18, 2),
    warehouse_name      varchar(50),
    oblast              varchar(200),
    income_id           bigint,
    external_code       varchar(50),
    odid                bigint,
    subject             varchar(200),
    category            varchar(200),
    brand               varchar(100),
    is_cancel           boolean,
    status              varchar(50),
    cancel_dt           timestamp,
    g_number            varchar(50),
    sticker             varchar(200),
    srid                varchar(50),
    quantity            numeric(18) not null,
    item_id             varchar(36),
    barcode_id          varchar(36),
    message             varchar(250)
    );

comment on column dl."order".last_change_date is 'Дата обновления информации в сервисе';

comment on column dl."order".last_change_time is 'Время обновления информации в сервисе';

comment on column dl."order".order_date is 'Дата заказа';

comment on column dl."order".order_time is 'Время заказа';

comment on column dl."order".supplier_article is 'Артикул поставщика';

comment on column dl."order".tech_size is 'Размер';

comment on column dl."order".barcode is 'Бар-код';

comment on column dl."order".total_price is 'Цена до согласованной итоговой скидки/промо/спп. Для получения цены со скидкой можно воспользоваться формулой priceWithDiscount = totalPrice * (1 - discountPercent/100)';

comment on column dl."order".discount_percent is 'Согласованный итоговый дисконт(процент)';

comment on column dl."order".discount_value is 'Согласованный итоговый дисконт(значение)';

comment on column dl."order".price_with_discount is 'Цена, priceWithDiscount = totalPrice * (1 - discountPercent/100)';

comment on column dl."order".warehouse_name is 'Название склада отгрузки';

comment on column dl."order".oblast is 'Область';

comment on column dl."order".income_id is 'Номер поставки (от продавца на склад)';

comment on column dl."order".external_code is 'Код WB';

comment on column dl."order".odid is 'Уникальный идентификатор позиции заказа';

comment on column dl."order".subject is 'Предмет';

comment on column dl."order".category is 'Категория';

comment on column dl."order".brand is 'Бренд';

comment on column dl."order".is_cancel is 'Отмена заказа. true - заказ отменен до оплаты';

comment on column dl."order".status is 'Статус заказа';

comment on column dl."order".cancel_dt is 'Дата и время отмены заказа';

comment on column dl."order".g_number is 'Номер заказа. Объединяет все позиции одного заказа.';

comment on column dl."order".sticker is 'Цифровое значение стикера';

comment on column dl."order".srid is 'Уникальный идентификатор заказа';

comment on column dl."order".quantity is 'Количество';

alter table dl."order"
    owner to "user";

create index if not exists order_report_date_idx
    on dl."order" (transaction_id, source, owner_code, warehouse_name, barcode);

grant select on dl."order" to readonly;

create table if not exists dl.order_delivered
(
    order_date          date                                   not null,
    source              varchar(20)                            not null,
    owner_code          varchar(20)                            not null,
    supplier_article    varchar(75)                            not null,
    warehouse           varchar(50)                            not null,
    barcode             varchar(30),
    external_code       varchar(50)                            not null,
    total_price         numeric(10, 2),
    price_with_discount numeric(10, 2),
    quantity            numeric(10)                            not null,
    create_at           timestamp with time zone default now() not null,
    updated_at          timestamp with time zone default now() not null
    )
    partition by RANGE (order_date);

comment on table dl.order_delivered is 'Таблица для хранения отгруженных заказов';

comment on column dl.order_delivered.order_date is 'Дата заказа';

comment on column dl.order_delivered.source is 'Источник заказа';

comment on column dl.order_delivered.owner_code is 'Код владельца';

comment on column dl.order_delivered.supplier_article is 'Артикул поставщика';

comment on column dl.order_delivered.warehouse is 'Название склада';

comment on column dl.order_delivered.barcode is 'Штрихкод';

comment on column dl.order_delivered.external_code is 'Внешний код';

comment on column dl.order_delivered.total_price is 'Сумма заказа';

comment on column dl.order_delivered.price_with_discount is 'Сумма заказа с учетом скидки';

comment on column dl.order_delivered.quantity is 'Количество';

alter table dl.order_delivered
    owner to "user";

grant select on dl.order_delivered to readonly;

create table if not exists dl.order_delivered_202301
    partition of dl.order_delivered
    FOR VALUES FROM ('2023-01-01') TO ('2023-02-01');

alter table dl.order_delivered_202301
    owner to "user";

grant select on dl.order_delivered_202301 to readonly;

create table if not exists dl.order_delivered_202302
    partition of dl.order_delivered
    FOR VALUES FROM ('2023-02-01') TO ('2023-03-01');

alter table dl.order_delivered_202302
    owner to "user";

grant select on dl.order_delivered_202302 to readonly;

create table if not exists dl.order_delivered_202303
    partition of dl.order_delivered
    FOR VALUES FROM ('2023-03-01') TO ('2023-04-01');

alter table dl.order_delivered_202303
    owner to "user";

grant select on dl.order_delivered_202303 to readonly;

create table if not exists dl.order_delivered_202304
    partition of dl.order_delivered
    FOR VALUES FROM ('2023-04-01') TO ('2023-05-01');

alter table dl.order_delivered_202304
    owner to "user";

grant select on dl.order_delivered_202304 to readonly;

create table if not exists dl.order_delivered_202305
    partition of dl.order_delivered
    FOR VALUES FROM ('2023-05-01') TO ('2023-06-01');

alter table dl.order_delivered_202305
    owner to "user";

grant select on dl.order_delivered_202305 to readonly;

create table if not exists dl.order_delivered_202306
    partition of dl.order_delivered
    FOR VALUES FROM ('2023-06-01') TO ('2023-07-01');

alter table dl.order_delivered_202306
    owner to "user";

grant select on dl.order_delivered_202306 to readonly;

create table if not exists dl.order_delivered_202307
    partition of dl.order_delivered
    FOR VALUES FROM ('2023-07-01') TO ('2023-08-01');

alter table dl.order_delivered_202307
    owner to "user";

grant select on dl.order_delivered_202307 to readonly;

create table if not exists dl.order_delivered_202308
    partition of dl.order_delivered
    FOR VALUES FROM ('2023-08-01') TO ('2023-09-01');

alter table dl.order_delivered_202308
    owner to "user";

grant select on dl.order_delivered_202308 to readonly;

create table if not exists dl.order_delivered_202309
    partition of dl.order_delivered
    FOR VALUES FROM ('2023-09-01') TO ('2023-10-01');

alter table dl.order_delivered_202309
    owner to "user";

grant select on dl.order_delivered_202309 to readonly;

create table if not exists dl.order_delivered_202310
    partition of dl.order_delivered
    FOR VALUES FROM ('2023-10-01') TO ('2023-11-01');

alter table dl.order_delivered_202310
    owner to "user";

grant select on dl.order_delivered_202310 to readonly;

create table if not exists dl.order_delivered_202311
    partition of dl.order_delivered
    FOR VALUES FROM ('2023-11-01') TO ('2023-12-01');

alter table dl.order_delivered_202311
    owner to "user";

grant select on dl.order_delivered_202311 to readonly;

create table if not exists dl.order_delivered_202312
    partition of dl.order_delivered
    FOR VALUES FROM ('2023-12-01') TO ('2024-01-01');

alter table dl.order_delivered_202312
    owner to "user";

grant select on dl.order_delivered_202312 to readonly;

create table if not exists dl.order_delivered_202401
    partition of dl.order_delivered
    FOR VALUES FROM ('2024-01-01') TO ('2024-02-01');

alter table dl.order_delivered_202401
    owner to "user";

grant select on dl.order_delivered_202401 to readonly;

create table if not exists dl.order_delivered_arch
(
    order_date          date        not null,
    transaction_id      integer     not null,
    source              varchar(20) not null,
    owner_code          varchar(20) not null,
    supplier_article    varchar(75) not null,
    warehouse           varchar(50) not null,
    barcode             varchar(30),
    external_code       varchar(50) not null,
    total_price         numeric(10, 2),
    price_with_discount numeric(10, 2),
    quantity            numeric(10) not null
    )
    partition by RANGE (order_date);

comment on table dl.order_delivered_arch is 'Архив заказов поставщиков';

comment on column dl.order_delivered_arch.order_date is 'Дата заказа';

comment on column dl.order_delivered_arch.transaction_id is 'Идентификатор транзакции';

comment on column dl.order_delivered_arch.source is 'Источник заказа';

comment on column dl.order_delivered_arch.owner_code is 'Код владельца';

comment on column dl.order_delivered_arch.supplier_article is 'Артикул поставщика';

comment on column dl.order_delivered_arch.warehouse is 'Склад';

comment on column dl.order_delivered_arch.barcode is 'Штрихкод';

comment on column dl.order_delivered_arch.external_code is 'Внешний код';

comment on column dl.order_delivered_arch.total_price is 'Сумма заказа';

comment on column dl.order_delivered_arch.price_with_discount is 'Сумма заказа с учетом скидки';

comment on column dl.order_delivered_arch.quantity is 'Количество';

alter table dl.order_delivered_arch
    owner to "user";

grant select on dl.order_delivered_arch to readonly;

create table if not exists dl.order_delivered_arch_202301
    partition of dl.order_delivered_arch
    FOR VALUES FROM ('2023-01-01') TO ('2023-02-01');

alter table dl.order_delivered_arch_202301
    owner to "user";

grant select on dl.order_delivered_arch_202301 to readonly;

create table if not exists dl.order_delivered_arch_202302
    partition of dl.order_delivered_arch
    FOR VALUES FROM ('2023-02-01') TO ('2023-03-01');

alter table dl.order_delivered_arch_202302
    owner to "user";

grant select on dl.order_delivered_arch_202302 to readonly;

create table if not exists dl.order_delivered_arch_202303
    partition of dl.order_delivered_arch
    FOR VALUES FROM ('2023-03-01') TO ('2023-04-01');

alter table dl.order_delivered_arch_202303
    owner to "user";

grant select on dl.order_delivered_arch_202303 to readonly;

create table if not exists dl.order_delivered_arch_202304
    partition of dl.order_delivered_arch
    FOR VALUES FROM ('2023-04-01') TO ('2023-05-01');

alter table dl.order_delivered_arch_202304
    owner to "user";

grant select on dl.order_delivered_arch_202304 to readonly;

create table if not exists dl.order_delivered_arch_202305
    partition of dl.order_delivered_arch
    FOR VALUES FROM ('2023-05-01') TO ('2023-06-01');

alter table dl.order_delivered_arch_202305
    owner to "user";

grant select on dl.order_delivered_arch_202305 to readonly;

create table if not exists dl.order_delivered_arch_202306
    partition of dl.order_delivered_arch
    FOR VALUES FROM ('2023-06-01') TO ('2023-07-01');

alter table dl.order_delivered_arch_202306
    owner to "user";

grant select on dl.order_delivered_arch_202306 to readonly;

create table if not exists dl.order_delivered_arch_202307
    partition of dl.order_delivered_arch
    FOR VALUES FROM ('2023-07-01') TO ('2023-08-01');

alter table dl.order_delivered_arch_202307
    owner to "user";

grant select on dl.order_delivered_arch_202307 to readonly;

create table if not exists dl.order_delivered_arch_202308
    partition of dl.order_delivered_arch
    FOR VALUES FROM ('2023-08-01') TO ('2023-09-01');

alter table dl.order_delivered_arch_202308
    owner to "user";

grant select on dl.order_delivered_arch_202308 to readonly;

create table if not exists dl.order_delivered_arch_202309
    partition of dl.order_delivered_arch
    FOR VALUES FROM ('2023-09-01') TO ('2023-10-01');

alter table dl.order_delivered_arch_202309
    owner to "user";

grant select on dl.order_delivered_arch_202309 to readonly;

create table if not exists dl.order_delivered_arch_202310
    partition of dl.order_delivered_arch
    FOR VALUES FROM ('2023-10-01') TO ('2023-11-01');

alter table dl.order_delivered_arch_202310
    owner to "user";

grant select on dl.order_delivered_arch_202310 to readonly;

create table if not exists dl.order_delivered_arch_202311
    partition of dl.order_delivered_arch
    FOR VALUES FROM ('2023-11-01') TO ('2023-12-01');

alter table dl.order_delivered_arch_202311
    owner to "user";

grant select on dl.order_delivered_arch_202311 to readonly;

create table if not exists dl.order_delivered_arch_202312
    partition of dl.order_delivered_arch
    FOR VALUES FROM ('2023-12-01') TO ('2024-01-01');

alter table dl.order_delivered_arch_202312
    owner to "user";

grant select on dl.order_delivered_arch_202312 to readonly;

create table if not exists dl.order_delivered_arch_202401
    partition of dl.order_delivered_arch
    FOR VALUES FROM ('2024-01-01') TO ('2024-02-01');

alter table dl.order_delivered_arch_202401
    owner to "user";

grant select on dl.order_delivered_arch_202401 to readonly;

create table if not exists dl.order_delivered_arch_def
    partition of dl.order_delivered_arch
    DEFAULT;

alter table dl.order_delivered_arch_def
    owner to "user";

grant select on dl.order_delivered_arch_def to readonly;

create table if not exists dl.order_delivered_def
    partition of dl.order_delivered
    DEFAULT;

alter table dl.order_delivered_def
    owner to "user";

grant select on dl.order_delivered_def to readonly;

create table if not exists dl.organisation
(
    id        varchar(36) not null,
    name      varchar(200),
    inn       varchar(15),
    kpp       varchar(15),
    update_at timestamp   not null
    );

comment on table dl.organisation is 'Таблица для хранения организаций';

comment on column dl.organisation.id is 'Идентификатор организации';

comment on column dl.organisation.name is 'Наименование организации';

comment on column dl.organisation.inn is 'ИНН организации';

comment on column dl.organisation.kpp is 'КПП организации';

comment on column dl.organisation.update_at is 'Дата и время обновления записи';

alter table dl.organisation
    owner to "user";

grant select on dl.organisation to readonly;

create table if not exists dl.report_detail_by_period
(
    id                          serial
    primary key,
    transaction_id              integer     not null,
    source                      varchar(50) not null,
    owner_code                  varchar(50) not null,
    acquiring_bank              varchar(100),
    acquiring_fee               numeric(18, 2),
    additional_payment          numeric(18, 2),
    barcode                     varchar(30),
    bonus_type_name             varchar(100),
    brand_name                  varchar(100),
    commission_percent          numeric(18),
    create_dt                   date,
    date_from                   date,
    date_to                     date,
    declaration_number          varchar(100),
    delivery_amount             numeric(18),
    delivery_rub                numeric(18, 2),
    doc_number                  varchar(50),
    doc_type_name               varchar(100),
    gi_box_type_name            varchar(100),
    gi_id                       numeric(18),
    kiz                         varchar(100),
    nm_id                       numeric(18),
    office_name                 varchar(100),
    order_dt                    date,
    penalty                     numeric(18, 2),
    ppvz_for_pay                numeric(18, 2),
    ppvz_inn                    varchar(20),
    ppvz_kvw_prc                numeric(10, 2),
    ppvz_kvw_prc_base           numeric(18, 2),
    ppvz_office_id              numeric(18),
    ppvz_office_name            varchar(100),
    ppvz_reward                 numeric(18, 2),
    ppvz_sales_commission       numeric(18, 2),
    ppvz_spp_prc                numeric(18, 2),
    ppvz_supplier_id            numeric(18),
    ppvz_supplier_name          varchar(100),
    ppvz_vw                     numeric(18, 2),
    ppvz_vw_nds                 numeric(18, 2),
    product_discount_for_report numeric(18, 2),
    quantity                    numeric(18),
    realizationreport_id        numeric(18),
    retail_amount               numeric(18, 2),
    retail_price                numeric(18, 2),
    retail_price_withdisc_rub   numeric(18, 2),
    return_amount               numeric(18),
    rid                         numeric(20),
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
    supplier_promo              numeric(18, 2),
    suppliercontract_code       varchar(100),
    ts_name                     varchar(100),
    item_id                     varchar(36),
    barcode_id                  varchar(36),
    message                     varchar(250)
    );

alter table dl.report_detail_by_period
    owner to "user";

create index if not exists report_detail_by_period_transaction_id_idx
    on dl.report_detail_by_period (transaction_id, source, owner_code, barcode);

grant select on dl.report_detail_by_period to readonly;

create table if not exists dl.sale
(
    id                  serial,
    transaction_date    date        not null,
    transaction_id      integer     not null,
    source              varchar(20) not null,
    owner_code          varchar(20) not null,
    last_change_date    date        not null,
    last_change_time    time        not null,
    sale_date           date        not null,
    sale_time           time        not null,
    supplier_article    varchar(75),
    tech_size           varchar(30),
    barcode             varchar(30),
    total_price         numeric(10, 2),
    discount_percent    numeric(10, 2),
    discount_value      numeric(10, 2),
    is_supply           boolean,
    is_realization      boolean,
    promo_code_discount double precision,
    warehouse_name      varchar(50),
    country_name        varchar(200),
    oblast_okrug_name   varchar(200),
    region_name         varchar(200),
    income_id           bigint,
    sale_id             varchar(15),
    odid                bigint,
    spp                 double precision,
    for_pay             double precision,
    finished_price      numeric(10, 2),
    price_with_disc     numeric(10, 2),
    external_code       varchar(50),
    subject             varchar(50),
    category            varchar(50),
    brand               varchar(50),
    is_storno           boolean,
    g_number            varchar(50),
    sticker             varchar(200),
    srid                varchar(50)
    );

comment on column dl.sale.last_change_date is 'Дата обновления информации в сервисе';

comment on column dl.sale.last_change_time is 'Время обновления информации в сервисе';

comment on column dl.sale.sale_date is 'Дата продажи';

comment on column dl.sale.sale_time is 'Время продажи';

comment on column dl.sale.supplier_article is 'Артикул поставщика';

comment on column dl.sale.tech_size is 'Размер';

comment on column dl.sale.barcode is 'Бар-код';

comment on column dl.sale.total_price is 'Цена до согласованной итоговой скидки/промо/спп. Для получения цены со скидкой можно воспользоваться формулой priceWithDiscount = totalPrice * (1 - discountPercent/100)';

comment on column dl.sale.discount_percent is 'Согласованный итоговый дисконт(процент)';

comment on column dl.sale.discount_value is 'Согласованный итоговый дисконт(значение)';

comment on column dl.sale.is_supply is 'Договор поставки';

comment on column dl.sale.is_realization is 'Договор реализации';

comment on column dl.sale.promo_code_discount is 'Скидка по промокоду';

comment on column dl.sale.warehouse_name is 'Название склада отгрузки';

comment on column dl.sale.country_name is 'Страна';

comment on column dl.sale.oblast_okrug_name is 'Область';

comment on column dl.sale.region_name is 'Регион';

comment on column dl.sale.income_id is 'Номер поставки (от продавца на склад)';

comment on column dl.sale.sale_id is 'Уникальный идентификатор продажи/возврата.';

comment on column dl.sale.odid is 'Уникальный идентификатор позиции заказа';

comment on column dl.sale.spp is 'Согласованная скидка постоянного покупателя';

comment on column dl.sale.for_pay is 'К перечислению поставщику';

comment on column dl.sale.finished_price is 'Фактическая цена заказа с учетом всех скидок';

comment on column dl.sale.price_with_disc is 'Цена, от которой считается вознаграждение поставщика forpay';

comment on column dl.sale.external_code is 'Код WB';

comment on column dl.sale.subject is 'Предмет';

comment on column dl.sale.category is 'Категория';

comment on column dl.sale.brand is 'Бренд';

comment on column dl.sale.is_storno is 'Для сторно-операций 1, для остальных 0';

comment on column dl.sale.g_number is 'Номер заказа. Объединяет все позиции одного заказа.';

comment on column dl.sale.sticker is 'Цифровое значение стикера';

comment on column dl.sale.srid is 'Уникальный идентификатор заказа';

alter table dl.sale
    owner to "user";

create index if not exists sale_transaction_id_idx
    on dl.sale (transaction_id, odid);

grant select on dl.sale to readonly;

create table if not exists dl.sale_agr
(
    transaction_id          varchar(50)    not null,
    order_date              date           not null,
    owner_code              varchar(20)    not null,
    source                  varchar(20)    not null,
    supplier_article        varchar(75),
    warehouse_name          varchar(50)    not null,
    barcode                 varchar(30),
    external_code           varchar(50)    not null,
    quantity                numeric(10)    not null,
    sum_total_price         numeric(10, 2) not null,
    sum_price_with_discount numeric(10, 2) not null
    );

comment on table dl.sale_agr is 'Таблица с данными по продажам';

comment on column dl.sale_agr.transaction_id is 'Идентификатор транзакции';

comment on column dl.sale_agr.order_date is 'Дата заказа';

comment on column dl.sale_agr.owner_code is 'Код владельца';

comment on column dl.sale_agr.source is 'Источник';

comment on column dl.sale_agr.supplier_article is 'Артикул поставщика';

comment on column dl.sale_agr.warehouse_name is 'Название склада';

comment on column dl.sale_agr.barcode is 'Штрихкод';

comment on column dl.sale_agr.external_code is 'Внешний код';

comment on column dl.sale_agr.quantity is 'Количество';

comment on column dl.sale_agr.sum_total_price is 'Цкнва';

comment on column dl.sale_agr.sum_price_with_discount is 'Цена с учетом скидки';

alter table dl.sale_agr
    owner to "user";

grant select on dl.sale_agr to readonly;

create table if not exists dl.stock
(
    id                   serial,
    transaction_date     date        not null,
    transaction_id       integer     not null,
    source               varchar(20) not null,
    owner_code           varchar(20) not null,
    last_change_date     date        not null,
    last_change_time     time        not null,
    supplier_article     varchar(75),
    tech_size            varchar(30),
    barcode              varchar(30),
    quantity             integer     not null,
    is_supply            boolean,
    is_realization       boolean,
    quantity_full        integer     not null,
    quantity_promised    integer,
    warehouse_name       varchar(50) not null,
    external_code        varchar(50),
    name                 varchar(200),
    subject              varchar(200),
    category             varchar(50),
    days_on_site         integer,
    brand                varchar(50),
    "SCCode"             varchar(50),
    price                numeric(10, 2),
    discount             numeric(10, 2),
    price_after_discount numeric(10, 2),
    card_created         date        not null,
    item_id              varchar(36),
    barcode_id           varchar(36),
    message              varchar(250)
    );

comment on column dl.stock.last_change_date is 'Дата обновления информации в сервисе';

comment on column dl.stock.last_change_time is 'Время обновления информации в сервисе';

comment on column dl.stock.supplier_article is 'Артикул поставщика';

comment on column dl.stock.tech_size is 'Размер';

comment on column dl.stock.barcode is 'Бар-код';

comment on column dl.stock.quantity is 'Количество, доступное для продажи';

comment on column dl.stock.is_supply is 'Договор поставки';

comment on column dl.stock.is_realization is 'Договор реализации';

comment on column dl.stock.quantity_full is 'Полное (непроданное) количество, которое числится за складом (= quantity + в пути)';

comment on column dl.stock.quantity_promised is 'Количество товара, указанное в подтверждённых будущих поставках';

comment on column dl.stock.warehouse_name is 'Название склада';

comment on column dl.stock.external_code is 'Код WB/OZON';

comment on column dl.stock.name is 'Наименование продукта';

comment on column dl.stock.subject is 'Предмет';

comment on column dl.stock.category is 'Категория';

comment on column dl.stock.days_on_site is 'Количество дней на сайте';

comment on column dl.stock.brand is 'Бренд';

comment on column dl.stock."SCCode" is 'Код контракта';

comment on column dl.stock.price is 'Цена';

comment on column dl.stock.discount is 'Скидка';

comment on column dl.stock.price_after_discount is 'Цена после скидки';

comment on column dl.stock.card_created is 'Дата создания карточки';

alter table dl.stock
    owner to "user";

grant select on dl.stock to readonly;

create table if not exists dl.stock_daily
(
    stock_date          date           not null,
    source              varchar(20)    not null,
    owner_code          varchar(20)    not null,
    supplier_article    varchar(75),
    barcode             varchar(30),
    external_code       varchar(50)    not null,
    name                varchar(200),
    subject             varchar(200),
    category            varchar(50),
    brand               varchar(50),
    warehouse           varchar(50)    not null,
    create_at           date           not null,
    update_at           date,
    quantity            numeric(10, 2) not null,
    quantity_full       numeric(10, 2) not null,
    attemption          integer        not null,
    price               numeric(10, 2) not null,
    price_with_discount numeric(10, 2) not null
    )
    partition by RANGE (stock_date);

alter table dl.stock_daily
    owner to "user";

grant select on dl.stock_daily to readonly;

create table if not exists dl.stock_daily_202301
    partition of dl.stock_daily
    FOR VALUES FROM ('2023-01-01') TO ('2023-02-01');

alter table dl.stock_daily_202301
    owner to "user";

grant select on dl.stock_daily_202301 to readonly;

create table if not exists dl.stock_daily_202302
    partition of dl.stock_daily
    FOR VALUES FROM ('2023-02-01') TO ('2023-03-01');

alter table dl.stock_daily_202302
    owner to "user";

grant select on dl.stock_daily_202302 to readonly;

create table if not exists dl.stock_daily_202303
    partition of dl.stock_daily
    FOR VALUES FROM ('2023-03-01') TO ('2023-04-01');

alter table dl.stock_daily_202303
    owner to "user";

grant select on dl.stock_daily_202303 to readonly;

create table if not exists dl.stock_daily_202304
    partition of dl.stock_daily
    FOR VALUES FROM ('2023-04-01') TO ('2023-05-01');

alter table dl.stock_daily_202304
    owner to "user";

grant select on dl.stock_daily_202304 to readonly;

create table if not exists dl.stock_daily_202305
    partition of dl.stock_daily
    FOR VALUES FROM ('2023-05-01') TO ('2023-06-01');

alter table dl.stock_daily_202305
    owner to "user";

grant select on dl.stock_daily_202305 to readonly;

create table if not exists dl.stock_daily_202306
    partition of dl.stock_daily
    FOR VALUES FROM ('2023-06-01') TO ('2023-07-01');

alter table dl.stock_daily_202306
    owner to "user";

grant select on dl.stock_daily_202306 to readonly;

create table if not exists dl.stock_daily_202307
    partition of dl.stock_daily
    FOR VALUES FROM ('2023-07-01') TO ('2023-08-01');

alter table dl.stock_daily_202307
    owner to "user";

grant select on dl.stock_daily_202307 to readonly;

create table if not exists dl.stock_daily_202308
    partition of dl.stock_daily
    FOR VALUES FROM ('2023-08-01') TO ('2023-09-01');

alter table dl.stock_daily_202308
    owner to "user";

grant select on dl.stock_daily_202308 to readonly;

create table if not exists dl.stock_daily_202309
    partition of dl.stock_daily
    FOR VALUES FROM ('2023-09-01') TO ('2023-10-01');

alter table dl.stock_daily_202309
    owner to "user";

grant select on dl.stock_daily_202309 to readonly;

create table if not exists dl.stock_daily_202310
    partition of dl.stock_daily
    FOR VALUES FROM ('2023-10-01') TO ('2023-11-01');

alter table dl.stock_daily_202310
    owner to "user";

grant select on dl.stock_daily_202310 to readonly;

create table if not exists dl.stock_daily_202311
    partition of dl.stock_daily
    FOR VALUES FROM ('2023-11-01') TO ('2023-12-01');

alter table dl.stock_daily_202311
    owner to "user";

grant select on dl.stock_daily_202311 to readonly;

create table if not exists dl.stock_daily_202312
    partition of dl.stock_daily
    FOR VALUES FROM ('2023-12-01') TO ('2024-01-01');

alter table dl.stock_daily_202312
    owner to "user";

grant select on dl.stock_daily_202312 to readonly;

create table if not exists dl.stock_daily_202401
    partition of dl.stock_daily
    FOR VALUES FROM ('2024-01-01') TO ('2024-02-01');

alter table dl.stock_daily_202401
    owner to "user";

grant select on dl.stock_daily_202401 to readonly;

create table if not exists dl.stock_daily_def
    partition of dl.stock_daily
    DEFAULT;

alter table dl.stock_daily_def
    owner to "user";

create index if not exists stock_daily_dwosec
    on dl.stock_daily (stock_date, source, owner_code, warehouse, external_code);

grant select on dl.stock_daily_def to readonly;

create table if not exists ml.job
(
    id          integer     not null,
    create_date timestamp default now(),
    is_active   boolean     not null,
    description varchar(100),
    week_days   varchar(200),
    at_time     varchar(200),
    interval    integer,
    max_runs    integer,
    next_run    timestamp,
    last_run    timestamp,
    type        varchar(20) not null
    );

comment on column ml.job.is_active is 'Активен ли';

comment on column ml.job.description is 'Описание';

comment on column ml.job.week_days is 'Дни недели monday | tuesday | wednesday | thursday | friday | saturday | sunday';

comment on column ml.job.at_time is 'Время в формате 8:04;16:00';

comment on column ml.job.interval is 'Интервал в секундах';

comment on column ml.job.max_runs is 'Максимальное количество запусков';

alter table ml.job
    owner to "user";

create table if not exists ml.job_owner
(
    job_id     integer     not null,
    owner_code varchar(20) not null
    );

alter table ml.job_owner
    owner to "user";

create table if not exists ml.log_load
(
    transaction_id integer     not null,
    owner_code     varchar(20) not null,
    source         varchar(20) not null,
    description    varchar(200),
    status         varchar(20) not null
    );

comment on table ml.log_load is 'Таблица для хранения логов загрузки данных';

comment on column ml.log_load.transaction_id is 'Идентификатор транзакции';

comment on column ml.log_load.owner_code is 'Код владельца';

comment on column ml.log_load.source is 'Источник данных';

comment on column ml.log_load.description is 'Описание';

comment on column ml.log_load.status is 'Статус BEGIN|COMPLETED|ERROR';

alter table ml.log_load
    owner to "user";

create table if not exists ml.marketplace
(
    code           varchar(15) not null
    primary key,
    marketplace_id varchar(36) not null
    );

alter table ml.marketplace
    owner to "user";

create table if not exists ml.order_delivered_log
(
    transaction_id bigint    not null,
    created_at     timestamp not null,
    added_rows     integer   not null
);

comment on table ml.order_delivered_log is 'Таблица для хранения логов отгруженных заказов';

comment on column ml.order_delivered_log.transaction_id is 'Идентификатор транзакции';

comment on column ml.order_delivered_log.created_at is 'Дата создания записи';

alter table ml.order_delivered_log
    owner to "user";

create table if not exists ml.owner
(
    code            varchar(20)  not null
    primary key,
    name            varchar(100) not null,
    create_date     timestamp default now(),
    is_deleted      boolean      not null,
    organisation_id varchar(36)
    );

alter table ml.owner
    owner to "user";

create table if not exists ml.owner_marketplace
(
    owner_code varchar(20)  not null,
    source     varchar(20)  not null,
    host       varchar(200) not null,
    password   varchar(200),
    client_id  varchar(50)
    );

alter table ml.owner_marketplace
    owner to "user";

create table if not exists ml.scheduler
(
    code        varchar(20) not null,
    description varchar(200),
    status      varchar(20) not null,
    update_at   timestamp   not null
    );

alter table ml.scheduler
    owner to "user";

create table if not exists ml.transaction
(
    id         integer default nextval('ml.transaction_id_seq'::regclass) not null,
    job_id     integer                                                    not null,
    start_date timestamp                                                  not null,
    end_date   timestamp,
    status     varchar(20),
    message    varchar(256)
    );

alter table ml.transaction
    owner to "user";

alter sequence ml.transaction_id_seq owned by ml.transaction.id;

create table if not exists dl.stock_def30
(
    stock_date       date,
    source           varchar(20),
    owner_code       varchar(20),
    supplier_article varchar(75),
    warehouse        varchar(50),
    barcode          varchar(30),
    external_code    varchar(50),
    marketplace_id   varchar(36),
    barcode_id       varchar(36),
    item_id          varchar(36),
    org_id           varchar(36),
    item_name        varchar(255),
    def30            integer,
    days_in_stock    integer,
    avg_price        numeric(19, 2),
    min_price        numeric(19, 2),
    max_price        numeric(19, 2),
    create_at        timestamp default now()
    )
    partition by RANGE (stock_date);

alter table dl.stock_def30
    owner to "user";

grant select on dl.stock_def30 to readonly;

create table if not exists dl.stock_def30_def
    partition of dl.stock_def30
    DEFAULT;

alter table dl.stock_def30_def
    owner to "user";

grant select on dl.stock_def30_def to readonly;

create table if not exists dl.stock_def30_202301
    partition of dl.stock_def30
    FOR VALUES FROM ('2023-01-01') TO ('2023-02-01');

alter table dl.stock_def30_202301
    owner to "user";

grant select on dl.stock_def30_202301 to readonly;

create table if not exists dl.stock_def30_202302
    partition of dl.stock_def30
    FOR VALUES FROM ('2023-02-01') TO ('2023-03-01');

alter table dl.stock_def30_202302
    owner to "user";

grant select on dl.stock_def30_202302 to readonly;

create table if not exists dl.stock_def30_202303
    partition of dl.stock_def30
    FOR VALUES FROM ('2023-03-01') TO ('2023-04-01');

alter table dl.stock_def30_202303
    owner to "user";

grant select on dl.stock_def30_202303 to readonly;

create table if not exists dl.stock_def30_202304
    partition of dl.stock_def30
    FOR VALUES FROM ('2023-04-01') TO ('2023-05-01');

alter table dl.stock_def30_202304
    owner to "user";

grant select on dl.stock_def30_202304 to readonly;

create table if not exists dl.stock_def30_202305
    partition of dl.stock_def30
    FOR VALUES FROM ('2023-05-01') TO ('2023-06-01');

alter table dl.stock_def30_202305
    owner to "user";

grant select on dl.stock_def30_202305 to readonly;

create table if not exists dl.stock_def30_202306
    partition of dl.stock_def30
    FOR VALUES FROM ('2023-06-01') TO ('2023-07-01');

alter table dl.stock_def30_202306
    owner to "user";

grant select on dl.stock_def30_202306 to readonly;

create table if not exists dl.stock_def30_202307
    partition of dl.stock_def30
    FOR VALUES FROM ('2023-07-01') TO ('2023-08-01');

alter table dl.stock_def30_202307
    owner to "user";

grant select on dl.stock_def30_202307 to readonly;

create table if not exists dl.stock_def30_202308
    partition of dl.stock_def30
    FOR VALUES FROM ('2023-08-01') TO ('2023-09-01');

alter table dl.stock_def30_202308
    owner to "user";

grant select on dl.stock_def30_202308 to readonly;

create table if not exists dl.stock_def30_202309
    partition of dl.stock_def30
    FOR VALUES FROM ('2023-09-01') TO ('2023-10-01');

alter table dl.stock_def30_202309
    owner to "user";

grant select on dl.stock_def30_202309 to readonly;

create table if not exists dl.stock_def30_202310
    partition of dl.stock_def30
    FOR VALUES FROM ('2023-10-01') TO ('2023-11-01');

alter table dl.stock_def30_202310
    owner to "user";

grant select on dl.stock_def30_202310 to readonly;

create table if not exists dl.stock_def30_202311
    partition of dl.stock_def30
    FOR VALUES FROM ('2023-11-01') TO ('2023-12-01');

alter table dl.stock_def30_202311
    owner to "user";

grant select on dl.stock_def30_202311 to readonly;

create table if not exists dl.stock_def30_202312
    partition of dl.stock_def30
    FOR VALUES FROM ('2023-12-01') TO ('2024-01-01');

alter table dl.stock_def30_202312
    owner to "user";

grant select on dl.stock_def30_202312 to readonly;

create table if not exists dl.stock_def30_202401
    partition of dl.stock_def30
    FOR VALUES FROM ('2024-01-01') TO ('2024-02-01');

alter table dl.stock_def30_202401
    owner to "user";

create index if not exists stock_def30_def_stock_date_idx
    on dl.stock_def30 (stock_date);

grant select on dl.stock_def30_202401 to readonly;

create table if not exists dl.stock_def
(
    stock_date       date,
    source           varchar(5),
    owner_code       varchar(20),
    supplier_article varchar(75),
    barcode          varchar(50),
    warehouse        varchar(50),
    external_code    varchar(50),
    marketplace_id   varchar(36),
    barcode_id       varchar(36),
    item_id          varchar(36),
    org_id           varchar(36),
    item_name        varchar(255),
    def30            numeric(10),
    days_in_stock30  numeric(10),
    def5             numeric(10),
    days_in_stock5   numeric(10),
    avg_price        numeric(10, 2),
    min_price        numeric(10, 2),
    max_price        numeric(10, 2)
    )
    partition by RANGE (stock_date);

comment on table dl.stock_def is 'Таблица с данными о деффектуре по дням';

comment on column dl.stock_def.stock_date is 'Дата';

comment on column dl.stock_def.source is 'Источник';

comment on column dl.stock_def.owner_code is 'Код владельца';

comment on column dl.stock_def.supplier_article is 'Артикул поставщика';

comment on column dl.stock_def.barcode is 'Штрихкод';

comment on column dl.stock_def.warehouse is 'Склад';

comment on column dl.stock_def.external_code is 'Внешний код';

comment on column dl.stock_def.marketplace_id is 'Идентификатор маркетплейса';

comment on column dl.stock_def.barcode_id is 'Идентификатор штрихкода';

comment on column dl.stock_def.item_id is 'Идентификатор товара';

comment on column dl.stock_def.org_id is 'Идентификатор организации';

comment on column dl.stock_def.item_name is 'Наименование товара';

comment on column dl.stock_def.def30 is 'Деффектура за 30 дней';

comment on column dl.stock_def.days_in_stock30 is 'Количество дней в наличии за 30 дней';

comment on column dl.stock_def.def5 is 'Деффектура за 5 дней';

comment on column dl.stock_def.days_in_stock5 is 'Количество дней в наличии за 5 дней';

comment on column dl.stock_def.avg_price is 'Средняя цена';

comment on column dl.stock_def.min_price is 'Минимальная цена';

comment on column dl.stock_def.max_price is 'Максимальная цена';

alter table dl.stock_def
    owner to "user";

grant select on dl.stock_def to readonly;

create table if not exists dl.stock_def_def
    partition of dl.stock_def
    DEFAULT;

alter table dl.stock_def_def
    owner to "user";

grant select on dl.stock_def_def to readonly;

create table if not exists dl.stock_def_202301
    partition of dl.stock_def
    FOR VALUES FROM ('2023-01-01') TO ('2023-02-01');

alter table dl.stock_def_202301
    owner to "user";

grant select on dl.stock_def_202301 to readonly;

create table if not exists dl.stock_def_202302
    partition of dl.stock_def
    FOR VALUES FROM ('2023-02-01') TO ('2023-03-01');

alter table dl.stock_def_202302
    owner to "user";

grant select on dl.stock_def_202302 to readonly;

create table if not exists dl.stock_def_202303
    partition of dl.stock_def
    FOR VALUES FROM ('2023-03-01') TO ('2023-04-01');

alter table dl.stock_def_202303
    owner to "user";

grant select on dl.stock_def_202303 to readonly;

create table if not exists dl.stock_def_202304
    partition of dl.stock_def
    FOR VALUES FROM ('2023-04-01') TO ('2023-05-01');

alter table dl.stock_def_202304
    owner to "user";

grant select on dl.stock_def_202304 to readonly;

create table if not exists dl.stock_def_202305
    partition of dl.stock_def
    FOR VALUES FROM ('2023-05-01') TO ('2023-06-01');

alter table dl.stock_def_202305
    owner to "user";

grant select on dl.stock_def_202305 to readonly;

create table if not exists dl.stock_def_202306
    partition of dl.stock_def
    FOR VALUES FROM ('2023-06-01') TO ('2023-07-01');

alter table dl.stock_def_202306
    owner to "user";

grant select on dl.stock_def_202306 to readonly;

create table if not exists dl.stock_def_202307
    partition of dl.stock_def
    FOR VALUES FROM ('2023-07-01') TO ('2023-08-01');

alter table dl.stock_def_202307
    owner to "user";

grant select on dl.stock_def_202307 to readonly;

create table if not exists dl.stock_def_202308
    partition of dl.stock_def
    FOR VALUES FROM ('2023-08-01') TO ('2023-09-01');

alter table dl.stock_def_202308
    owner to "user";

grant select on dl.stock_def_202308 to readonly;

create table if not exists dl.stock_def_202309
    partition of dl.stock_def
    FOR VALUES FROM ('2023-09-01') TO ('2023-10-01');

alter table dl.stock_def_202309
    owner to "user";

grant select on dl.stock_def_202309 to readonly;

create table if not exists dl.stock_def_202310
    partition of dl.stock_def
    FOR VALUES FROM ('2023-10-01') TO ('2023-11-01');

alter table dl.stock_def_202310
    owner to "user";

grant select on dl.stock_def_202310 to readonly;

create table if not exists dl.stock_def_202311
    partition of dl.stock_def
    FOR VALUES FROM ('2023-11-01') TO ('2023-12-01');

alter table dl.stock_def_202311
    owner to "user";

grant select on dl.stock_def_202311 to readonly;

create table if not exists dl.stock_def_202312
    partition of dl.stock_def
    FOR VALUES FROM ('2023-12-01') TO ('2024-01-01');

alter table dl.stock_def_202312
    owner to "user";

grant select on dl.stock_def_202312 to readonly;

create table if not exists dl.stock_def_202401
    partition of dl.stock_def
    FOR VALUES FROM ('2024-01-01') TO ('2024-02-01');

alter table dl.stock_def_202401
    owner to "user";

create index if not exists stock_def_stock_date_idx
    on dl.stock_def (stock_date);

grant select on dl.stock_def_202401 to readonly;

create table if not exists dl.sales_stock
(
    report_date      date,
    source           varchar(5),
    owner_code       varchar(20),
    supplier_article varchar(75),
    barcode          varchar(50),
    warehouse_name   varchar(50),
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
    is_excluded      boolean,
    cluster          varchar(20)
    )
    partition by RANGE (report_date);

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

alter table dl.sales_stock
    owner to "user";

grant select on dl.sales_stock to readonly;

create table if not exists dl.sales_stock_def
    partition of dl.sales_stock
    DEFAULT;

alter table dl.sales_stock_def
    owner to "user";

grant select on dl.sales_stock_def to readonly;

create table if not exists dl.sales_stock_202301
    partition of dl.sales_stock
    FOR VALUES FROM ('2023-01-01') TO ('2023-02-01');

alter table dl.sales_stock_202301
    owner to "user";

grant select on dl.sales_stock_202301 to readonly;

create table if not exists dl.sales_stock_202302
    partition of dl.sales_stock
    FOR VALUES FROM ('2023-02-01') TO ('2023-03-01');

alter table dl.sales_stock_202302
    owner to "user";

grant select on dl.sales_stock_202302 to readonly;

create table if not exists dl.sales_stock_202303
    partition of dl.sales_stock
    FOR VALUES FROM ('2023-03-01') TO ('2023-04-01');

alter table dl.sales_stock_202303
    owner to "user";

grant select on dl.sales_stock_202303 to readonly;

create table if not exists dl.sales_stock_202304
    partition of dl.sales_stock
    FOR VALUES FROM ('2023-04-01') TO ('2023-05-01');

alter table dl.sales_stock_202304
    owner to "user";

grant select on dl.sales_stock_202304 to readonly;

create table if not exists dl.sales_stock_202305
    partition of dl.sales_stock
    FOR VALUES FROM ('2023-05-01') TO ('2023-06-01');

alter table dl.sales_stock_202305
    owner to "user";

grant select on dl.sales_stock_202305 to readonly;

create table if not exists dl.sales_stock_202306
    partition of dl.sales_stock
    FOR VALUES FROM ('2023-06-01') TO ('2023-07-01');

alter table dl.sales_stock_202306
    owner to "user";

grant select on dl.sales_stock_202306 to readonly;

create table if not exists dl.sales_stock_202307
    partition of dl.sales_stock
    FOR VALUES FROM ('2023-07-01') TO ('2023-08-01');

alter table dl.sales_stock_202307
    owner to "user";

grant select on dl.sales_stock_202307 to readonly;

create table if not exists dl.sales_stock_202308
    partition of dl.sales_stock
    FOR VALUES FROM ('2023-08-01') TO ('2023-09-01');

alter table dl.sales_stock_202308
    owner to "user";

grant select on dl.sales_stock_202308 to readonly;

create table if not exists dl.sales_stock_202309
    partition of dl.sales_stock
    FOR VALUES FROM ('2023-09-01') TO ('2023-10-01');

alter table dl.sales_stock_202309
    owner to "user";

grant select on dl.sales_stock_202309 to readonly;

create table if not exists dl.sales_stock_202310
    partition of dl.sales_stock
    FOR VALUES FROM ('2023-10-01') TO ('2023-11-01');

alter table dl.sales_stock_202310
    owner to "user";

grant select on dl.sales_stock_202310 to readonly;

create table if not exists dl.sales_stock_202311
    partition of dl.sales_stock
    FOR VALUES FROM ('2023-11-01') TO ('2023-12-01');

alter table dl.sales_stock_202311
    owner to "user";

grant select on dl.sales_stock_202311 to readonly;

create table if not exists dl.sales_stock_202312
    partition of dl.sales_stock
    FOR VALUES FROM ('2023-12-01') TO ('2024-01-01');

alter table dl.sales_stock_202312
    owner to "user";

grant select on dl.sales_stock_202312 to readonly;

create table if not exists dl.sales_stock_202401
    partition of dl.sales_stock
    FOR VALUES FROM ('2024-01-01') TO ('2024-02-01');

alter table dl.sales_stock_202401
    owner to "user";

grant select on dl.sales_stock_202401 to readonly;

create table if not exists ml.settings
(
    code        varchar(75) not null
    primary key,
    value       varchar(150),
    description varchar(100)
    );

alter table ml.settings
    owner to "user";

create table if not exists ml.tg_group
(
    user_name  varchar(75) not null,
    group_name varchar(75) not null,
    chat_id    numeric(10) not null
    constraint pk_tg_group
    primary key
    );

alter table ml.tg_group
    owner to "user";

create table if not exists dl.exclud_item
(
    name        varchar(100),
    source      varchar(5),
    org_name    text,
    barcode     varchar(20),
    article     varchar(20),
    owner_code  varchar(20),
    create_date date default CURRENT_DATE
    );

alter table dl.exclud_item
    owner to "user";

create unique index if not exists idx_exclud_item
    on dl.exclud_item (source, owner_code, barcode);

grant select on dl.exclud_item to readonly;

create table if not exists dl.warehouse
(
    code    varchar(50) not null
    constraint pk_warehouse
    primary key,
    cluster varchar(50) not null,
    source  varchar(10) not null
    );

alter table dl.warehouse
    owner to "user";

grant select on dl.warehouse to readonly;

create table if not exists dl.stock_cluster_def
(
    stock_date       date,
    source           varchar(5),
    owner_code       varchar(20),
    supplier_article varchar(75),
    barcode          varchar(50),
    cluster          varchar(50),
    external_code    varchar(50),
    marketplace_id   varchar(36),
    barcode_id       varchar(36),
    item_id          varchar(36),
    org_id           varchar(36),
    item_name        varchar(255),
    def30            numeric(10),
    days_in_stock30  numeric(10),
    def5             numeric(10),
    days_in_stock5   numeric(10),
    avg_price        numeric(10, 2),
    min_price        numeric(10, 2),
    max_price        numeric(10, 2)
    )
    partition by RANGE (stock_date);

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

alter table dl.stock_cluster_def
    owner to "user";

create table if not exists dl.stock_cluster_def_def
    partition of dl.stock_cluster_def
    DEFAULT;

alter table dl.stock_cluster_def_def
    owner to "user";

create table if not exists dl.stock_cluster_def_202301
    partition of dl.stock_cluster_def
    FOR VALUES FROM ('2023-01-01') TO ('2023-02-01');

alter table dl.stock_cluster_def_202301
    owner to "user";

create table if not exists dl.stock_cluster_def_202302
    partition of dl.stock_cluster_def
    FOR VALUES FROM ('2023-02-01') TO ('2023-03-01');

alter table dl.stock_cluster_def_202302
    owner to "user";

create table if not exists dl.stock_cluster_def_202303
    partition of dl.stock_cluster_def
    FOR VALUES FROM ('2023-03-01') TO ('2023-04-01');

alter table dl.stock_cluster_def_202303
    owner to "user";

create table if not exists dl.stock_cluster_def_202304
    partition of dl.stock_cluster_def
    FOR VALUES FROM ('2023-04-01') TO ('2023-05-01');

alter table dl.stock_cluster_def_202304
    owner to "user";

create table if not exists dl.stock_cluster_def_202305
    partition of dl.stock_cluster_def
    FOR VALUES FROM ('2023-05-01') TO ('2023-06-01');

alter table dl.stock_cluster_def_202305
    owner to "user";

create table if not exists dl.stock_cluster_def_202306
    partition of dl.stock_cluster_def
    FOR VALUES FROM ('2023-06-01') TO ('2023-07-01');

alter table dl.stock_cluster_def_202306
    owner to "user";

create table if not exists dl.stock_cluster_def_202307
    partition of dl.stock_cluster_def
    FOR VALUES FROM ('2023-07-01') TO ('2023-08-01');

alter table dl.stock_cluster_def_202307
    owner to "user";

create table if not exists dl.stock_cluster_def_202308
    partition of dl.stock_cluster_def
    FOR VALUES FROM ('2023-08-01') TO ('2023-09-01');

alter table dl.stock_cluster_def_202308
    owner to "user";

create table if not exists dl.stock_cluster_def_202309
    partition of dl.stock_cluster_def
    FOR VALUES FROM ('2023-09-01') TO ('2023-10-01');

alter table dl.stock_cluster_def_202309
    owner to "user";

create table if not exists dl.stock_cluster_def_202310
    partition of dl.stock_cluster_def
    FOR VALUES FROM ('2023-10-01') TO ('2023-11-01');

alter table dl.stock_cluster_def_202310
    owner to "user";

create table if not exists dl.stock_cluster_def_202311
    partition of dl.stock_cluster_def
    FOR VALUES FROM ('2023-11-01') TO ('2023-12-01');

alter table dl.stock_cluster_def_202311
    owner to "user";

create table if not exists dl.stock_cluster_def_202312
    partition of dl.stock_cluster_def
    FOR VALUES FROM ('2023-12-01') TO ('2024-01-01');

alter table dl.stock_cluster_def_202312
    owner to "user";

create table if not exists dl.stock_cluster_def_202401
    partition of dl.stock_cluster_def
    FOR VALUES FROM ('2024-01-01') TO ('2024-02-01');

alter table dl.stock_cluster_def_202401
    owner to "user";

create index if not exists stock_def_stock_cluster_date_idx
    on dl.stock_cluster_def (stock_date);

create table if not exists dl.report_by_cluster
(
    report_date      date,
    source           varchar(5),
    owner_code       varchar(20),
    supplier_article varchar(75),
    cluster          varchar(20),
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
    is_excluded      boolean,
    transaction_id   numeric(10)
    )
    partition by RANGE (report_date);

comment on table dl.report_by_cluster is 'Отчет по кластерам';

comment on column dl.report_by_cluster.report_date is 'Дата отчета';

comment on column dl.report_by_cluster.source is 'Источник';

comment on column dl.report_by_cluster.owner_code is 'Код владельца';

comment on column dl.report_by_cluster.supplier_article is 'Артикул поставщика';

comment on column dl.report_by_cluster.cluster is 'Кластер';

comment on column dl.report_by_cluster.barcode is 'Штрихкод';

comment on column dl.report_by_cluster.external_code is 'Внешний код';

comment on column dl.report_by_cluster.item_id is 'Идентификатор товара';

comment on column dl.report_by_cluster.item_name is 'Наименование товара';

comment on column dl.report_by_cluster.marketplace_id is 'Идентификатор маркетплейса';

comment on column dl.report_by_cluster.org_id is 'Идентификатор организации';

comment on column dl.report_by_cluster.def30 is 'Дефицит за 30 дней';

comment on column dl.report_by_cluster.days_in_stock30 is 'Дней в наличии за 30 дней';

comment on column dl.report_by_cluster.def5 is 'Дефицит за 5 дней';

comment on column dl.report_by_cluster.days_in_stock5 is 'Дней в наличии за 5 дней';

comment on column dl.report_by_cluster.avg_price is 'Средняя цена';

comment on column dl.report_by_cluster.min_price is 'Минимальная цена';

comment on column dl.report_by_cluster.max_price is 'Максимальная цена';

comment on column dl.report_by_cluster.quantity30 is 'Количество за 30 дней';

comment on column dl.report_by_cluster.quantity5 is 'Количество за 5 дней';

comment on column dl.report_by_cluster.order_by_day30 is 'Заказов в день за 30 дней';

comment on column dl.report_by_cluster.forecast_order30 is 'Прогноз заказов за 30 дней';

comment on column dl.report_by_cluster.order_by_day5 is 'Заказов в день за 5 дней';

comment on column dl.report_by_cluster.forecast_order5 is 'Прогноз заказов за 5 дней';

comment on column dl.report_by_cluster.quantity is 'Количество';

comment on column dl.report_by_cluster.is_excluded is 'Исключен';

alter table dl.report_by_cluster
    owner to "user";

create table if not exists dl.report_by_cluster_def
    partition of dl.report_by_cluster
    DEFAULT;

alter table dl.report_by_cluster_def
    owner to "user";

create table if not exists dl.report_by_cluster_202301
    partition of dl.report_by_cluster
    FOR VALUES FROM ('2023-01-01') TO ('2023-02-01');

alter table dl.report_by_cluster_202301
    owner to "user";

create table if not exists dl.report_by_cluster_202302
    partition of dl.report_by_cluster
    FOR VALUES FROM ('2023-02-01') TO ('2023-03-01');

alter table dl.report_by_cluster_202302
    owner to "user";

create table if not exists dl.report_by_cluster_202303
    partition of dl.report_by_cluster
    FOR VALUES FROM ('2023-03-01') TO ('2023-04-01');

alter table dl.report_by_cluster_202303
    owner to "user";

create table if not exists dl.report_by_cluster_202304
    partition of dl.report_by_cluster
    FOR VALUES FROM ('2023-04-01') TO ('2023-05-01');

alter table dl.report_by_cluster_202304
    owner to "user";

create table if not exists dl.report_by_cluster_202305
    partition of dl.report_by_cluster
    FOR VALUES FROM ('2023-05-01') TO ('2023-06-01');

alter table dl.report_by_cluster_202305
    owner to "user";

create table if not exists dl.report_by_cluster_202306
    partition of dl.report_by_cluster
    FOR VALUES FROM ('2023-06-01') TO ('2023-07-01');

alter table dl.report_by_cluster_202306
    owner to "user";

create table if not exists dl.report_by_cluster_202307
    partition of dl.report_by_cluster
    FOR VALUES FROM ('2023-07-01') TO ('2023-08-01');

alter table dl.report_by_cluster_202307
    owner to "user";

create table if not exists dl.report_by_cluster_202308
    partition of dl.report_by_cluster
    FOR VALUES FROM ('2023-08-01') TO ('2023-09-01');

alter table dl.report_by_cluster_202308
    owner to "user";

create table if not exists dl.report_by_cluster_202309
    partition of dl.report_by_cluster
    FOR VALUES FROM ('2023-09-01') TO ('2023-10-01');

alter table dl.report_by_cluster_202309
    owner to "user";

create table if not exists dl.report_by_cluster_202310
    partition of dl.report_by_cluster
    FOR VALUES FROM ('2023-10-01') TO ('2023-11-01');

alter table dl.report_by_cluster_202310
    owner to "user";

create table if not exists dl.report_by_cluster_202311
    partition of dl.report_by_cluster
    FOR VALUES FROM ('2023-11-01') TO ('2023-12-01');

alter table dl.report_by_cluster_202311
    owner to "user";

create table if not exists dl.report_by_cluster_202312
    partition of dl.report_by_cluster
    FOR VALUES FROM ('2023-12-01') TO ('2024-01-01');

alter table dl.report_by_cluster_202312
    owner to "user";

create table if not exists dl.report_by_cluster_202401
    partition of dl.report_by_cluster
    FOR VALUES FROM ('2024-01-01') TO ('2024-02-01');

alter table dl.report_by_cluster_202401
    owner to "user";

create table if not exists ml.notification
(
    id         serial,
    type       varchar(50) not null,
    message    varchar(1000),
    created_at timestamp default now(),
    is_sent    boolean     not null,
    sender     varchar(70)
    );

alter table ml.notification
    owner to "user";

create table if not exists dl.stock1c
(
    stock_date date        not null,
    item_id    varchar(20) not null,
    quantity   numeric(10) not null,
    primary key (stock_date, item_id)
    );

alter table dl.stock1c
    owner to "user";

create table if not exists dl.stock_item_def
(
    stock_date       date,
    source           varchar(5),
    owner_code       varchar(20),
    supplier_article varchar(75),
    barcode          varchar(50),
    external_code    varchar(50),
    marketplace_id   varchar(36),
    barcode_id       varchar(36),
    item_id          varchar(36),
    org_id           varchar(36),
    item_name        varchar(255),
    def30            numeric(10),
    days_in_stock30  numeric(10),
    def5             numeric(10),
    days_in_stock5   numeric(10),
    avg_price        numeric(10, 2),
    min_price        numeric(10, 2),
    max_price        numeric(10, 2)
    )
    partition by RANGE (stock_date);

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

alter table dl.stock_item_def
    owner to "user";

create table if not exists dl.stock_item_def_def
    partition of dl.stock_item_def
    DEFAULT;

alter table dl.stock_item_def_def
    owner to "user";

create table if not exists dl.stock_item_def_202301
    partition of dl.stock_item_def
    FOR VALUES FROM ('2023-01-01') TO ('2023-02-01');

alter table dl.stock_item_def_202301
    owner to "user";

create table if not exists dl.stock_item_def_202302
    partition of dl.stock_item_def
    FOR VALUES FROM ('2023-02-01') TO ('2023-03-01');

alter table dl.stock_item_def_202302
    owner to "user";

create table if not exists dl.stock_item_def_202303
    partition of dl.stock_item_def
    FOR VALUES FROM ('2023-03-01') TO ('2023-04-01');

alter table dl.stock_item_def_202303
    owner to "user";

create table if not exists dl.stock_item_def_202304
    partition of dl.stock_item_def
    FOR VALUES FROM ('2023-04-01') TO ('2023-05-01');

alter table dl.stock_item_def_202304
    owner to "user";

create table if not exists dl.stock_item_def_202305
    partition of dl.stock_item_def
    FOR VALUES FROM ('2023-05-01') TO ('2023-06-01');

alter table dl.stock_item_def_202305
    owner to "user";

create table if not exists dl.stock_item_def_202306
    partition of dl.stock_item_def
    FOR VALUES FROM ('2023-06-01') TO ('2023-07-01');

alter table dl.stock_item_def_202306
    owner to "user";

create table if not exists dl.stock_item_def_202307
    partition of dl.stock_item_def
    FOR VALUES FROM ('2023-07-01') TO ('2023-08-01');

alter table dl.stock_item_def_202307
    owner to "user";

create table if not exists dl.stock_item_def_202308
    partition of dl.stock_item_def
    FOR VALUES FROM ('2023-08-01') TO ('2023-09-01');

alter table dl.stock_item_def_202308
    owner to "user";

create table if not exists dl.stock_item_def_202309
    partition of dl.stock_item_def
    FOR VALUES FROM ('2023-09-01') TO ('2023-10-01');

alter table dl.stock_item_def_202309
    owner to "user";

create table if not exists dl.stock_item_def_202310
    partition of dl.stock_item_def
    FOR VALUES FROM ('2023-10-01') TO ('2023-11-01');

alter table dl.stock_item_def_202310
    owner to "user";

create table if not exists dl.stock_item_def_202311
    partition of dl.stock_item_def
    FOR VALUES FROM ('2023-11-01') TO ('2023-12-01');

alter table dl.stock_item_def_202311
    owner to "user";

create table if not exists dl.stock_item_def_202312
    partition of dl.stock_item_def
    FOR VALUES FROM ('2023-12-01') TO ('2024-01-01');

alter table dl.stock_item_def_202312
    owner to "user";

create table if not exists dl.stock_item_def_202401
    partition of dl.stock_item_def
    FOR VALUES FROM ('2024-01-01') TO ('2024-02-01');

alter table dl.stock_item_def_202401
    owner to "user";

create index if not exists stock_item_date_idx
    on dl.stock_item_def (stock_date);

create table if not exists dl.report_by_item
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
    is_excluded      boolean,
    quantity1c       numeric(10)
    )
    partition by RANGE (report_date);

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

alter table dl.report_by_item
    owner to "user";

create table if not exists dl.report_by_item_def
    partition of dl.report_by_item
    DEFAULT;

alter table dl.report_by_item_def
    owner to "user";

create table if not exists dl.report_by_item_202301
    partition of dl.report_by_item
    FOR VALUES FROM ('2023-01-01') TO ('2023-02-01');

alter table dl.report_by_item_202301
    owner to "user";

create table if not exists dl.report_by_item_202302
    partition of dl.report_by_item
    FOR VALUES FROM ('2023-02-01') TO ('2023-03-01');

alter table dl.report_by_item_202302
    owner to "user";

create table if not exists dl.report_by_item_202303
    partition of dl.report_by_item
    FOR VALUES FROM ('2023-03-01') TO ('2023-04-01');

alter table dl.report_by_item_202303
    owner to "user";

create table if not exists dl.report_by_item_202304
    partition of dl.report_by_item
    FOR VALUES FROM ('2023-04-01') TO ('2023-05-01');

alter table dl.report_by_item_202304
    owner to "user";

create table if not exists dl.report_by_item_202305
    partition of dl.report_by_item
    FOR VALUES FROM ('2023-05-01') TO ('2023-06-01');

alter table dl.report_by_item_202305
    owner to "user";

create table if not exists dl.report_by_item_202306
    partition of dl.report_by_item
    FOR VALUES FROM ('2023-06-01') TO ('2023-07-01');

alter table dl.report_by_item_202306
    owner to "user";

create table if not exists dl.report_by_item_202307
    partition of dl.report_by_item
    FOR VALUES FROM ('2023-07-01') TO ('2023-08-01');

alter table dl.report_by_item_202307
    owner to "user";

create table if not exists dl.report_by_item_202308
    partition of dl.report_by_item
    FOR VALUES FROM ('2023-08-01') TO ('2023-09-01');

alter table dl.report_by_item_202308
    owner to "user";

create table if not exists dl.report_by_item_202309
    partition of dl.report_by_item
    FOR VALUES FROM ('2023-09-01') TO ('2023-10-01');

alter table dl.report_by_item_202309
    owner to "user";

create table if not exists dl.report_by_item_202310
    partition of dl.report_by_item
    FOR VALUES FROM ('2023-10-01') TO ('2023-11-01');

alter table dl.report_by_item_202310
    owner to "user";

create table if not exists dl.report_by_item_202311
    partition of dl.report_by_item
    FOR VALUES FROM ('2023-11-01') TO ('2023-12-01');

alter table dl.report_by_item_202311
    owner to "user";

create table if not exists dl.report_by_item_202312
    partition of dl.report_by_item
    FOR VALUES FROM ('2023-12-01') TO ('2024-01-01');

alter table dl.report_by_item_202312
    owner to "user";

create table if not exists dl.report_by_item_202401
    partition of dl.report_by_item
    FOR VALUES FROM ('2024-01-01') TO ('2024-02-01');

alter table dl.report_by_item_202401
    owner to "user";

create or replace procedure ml.calc_order_delivered(IN p_id integer)
    language plpgsql
as
$$
declare
v_id "ml"."transaction".id%type;
v_count integer;
begin
select coalesce(max(transaction_id), 0) into v_id from "dl"."order_delivered_log";
if p_id <= v_id then
        raise exception 'transaction_id % is already processed', p_id;
end if;
---- CALCULATE DELIVERED ORDERS ARCH
insert into "dl"."order_delivered_arch" ("transaction_id",
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
FROM "dl"."order" o
         left outer join "dl"."order_delivered_arch" oda on
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
                 FROM "dl"."order" o2
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
MERGE INTO "dl"."order_delivered" AS od
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
           FROM "dl"."order" o
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

insert into "ml"."order_delivered_log" ("transaction_id", "created_at", "added_rows")
values (p_id, now(), v_count);
end;
$$;

alter procedure ml.calc_order_delivered(integer) owner to "user";

create or replace procedure ml.calc_stock_daily(IN p_id integer)
    language plpgsql
as
$$
declare
v_status "ml"."transaction"."status"%type;
begin
select "status" into v_status from "ml"."transaction" where "id" = p_id;
if v_status not in ('COMPLETED', 'BEGIN') then
        raise exception 'Transaction status is not completed';
end if;

merge into "dl"."stock_daily" ss
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
           from "dl"."stock"
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

alter procedure ml.calc_stock_daily(integer) owner to "user";

create or replace procedure ml.partition_for_order_delivered(IN start_date date, IN end_date date)
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
for v_table, v_from, v_to in
SELECT 'order_delivered_' || to_char(day::date,'YYYYMM'),
       to_char(day::date,'YYYY-MM-DD'),
       to_char(day::date + interval '1 month','YYYY-MM-DD')
FROM generate_series(start_date, end_date, '1 month') day
            loop
            EXECUTE format(
                'CREATE TABLE IF NOT EXISTS "dl"."%s" PARTITION OF "dl"."order_delivered" FOR VALUES FROM(''%s'') TO (''%s'')',
                v_table,
                v_from,
                v_to
             );
end loop;
end;$$;

alter procedure ml.partition_for_order_delivered(date, date) owner to "user";

create or replace procedure ml.partition_for_order_delivered_arch(IN start_date date, IN end_date date)
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
for v_table, v_from, v_to in SELECT 'order_delivered_arch_' || to_char(day::date,'YYYYMM'),
                                    to_char(day::date,'YYYY-MM-DD'),
                                    to_char(day::date + interval '1 month','YYYY-MM-DD')
                             FROM generate_series(start_date, end_date, '1 month') day
            loop
            EXECUTE format(
                'CREATE TABLE IF NOT EXISTS "dl"."%s" PARTITION OF "dl"."order_delivered_arch" FOR VALUES FROM(''%s'') TO (''%s'')',
                v_table,
                v_from,
                v_to
             );
end loop;
end;$$;

alter procedure ml.partition_for_order_delivered_arch(date, date) owner to "user";

create or replace procedure ml.partition_for_stock_daily(IN start_date date, IN end_date date)
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
for v_table, v_from, v_to in
SELECT 'stock_daily_' || to_char(day::date,'YYYYMM'),
       to_char(day::date,'YYYY-MM-DD'),
       to_char(day::date + interval '1 month','YYYY-MM-DD')
FROM generate_series(start_date, end_date, '1 month') day
            loop
            EXECUTE format(
                'CREATE TABLE IF NOT EXISTS "dl"."%s" PARTITION OF "dl"."stock_daily" FOR VALUES FROM(''%s'') TO (''%s'')',
                v_table,
                v_from,
                v_to
             );
EXECUTE format(
        'create index %s_soec_idx on "dl"."%s" ("source", "owner_code", "warehouse", "external_code")',
        v_table,
        v_table
    );
end loop;
end;$$;

alter procedure ml.partition_for_stock_daily(date, date) owner to "user";

create or replace procedure dl.calc_stock_daily_by_day(IN p_day date)
    language plpgsql
as
$$
declare
v_count integer;
    v_day   date := date_trunc('day', p_day);
begin
select count(1) into v_count from dl.stock_daily where stock_date = v_day;
if v_count > 0 then
        return;
end if;

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
     , v_count                                                attention
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
         s.warehouse_name, s.external_code;
end
$$;

alter procedure dl.calc_stock_daily_by_day(date) owner to "user";

create or replace procedure dl.partition_for_stock_def30(IN start_date date, IN end_date date)
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
for v_table, v_from, v_to in SELECT 'stock_def30_' || to_char(day::date,'YYYYMM'),
                                    to_char(day::date,'YYYY-MM-DD'),
                                    to_char(day::date + interval '1 month','YYYY-MM-DD')
                             FROM generate_series(start_date, end_date, '1 month') day
            loop
            EXECUTE format(
                'CREATE TABLE IF NOT EXISTS "dl"."%s" PARTITION OF "dl"."stock_def30" FOR VALUES FROM(''%s'') TO (''%s'')',
                v_table,
                v_from,
                v_to
             );
--             EXECUTE format(
--                 'CREATE INDEX IF NOT EXISTS "dl"."%s_stock_date_idx" ON "dl"."%s" ("stock_date")',
--                 v_table,
--                 v_table
--              );
end loop;
end;$$;

alter procedure dl.partition_for_stock_def30(date, date) owner to "user";

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

create or replace procedure dl.partition_for_stock_def(IN start_date date, IN end_date date)
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
for v_table, v_from, v_to in SELECT 'stock_def_' || to_char(day::date,'YYYYMM'),
                                    to_char(day::date,'YYYY-MM-DD'),
                                    to_char(day::date + interval '1 month','YYYY-MM-DD')
                             FROM generate_series(start_date, end_date, '1 month') day
            loop
            EXECUTE format(
                'CREATE TABLE IF NOT EXISTS "dl"."%s" PARTITION OF "dl"."stock_def" FOR VALUES FROM(''%s'') TO (''%s'')',
                v_table,
                v_from,
                v_to
             );
end loop;
end;$$;

alter procedure dl.partition_for_stock_def(date, date) owner to "user";

create or replace procedure dl.calc_stock_def_by_day(IN p_day date)
    language plpgsql
as
$$
declare
v_count int;
    v_day   date := date_trunc('day', p_day);
begin
select count(1) into v_count from dl.stock_def where stock_date = v_day;
if v_count > 0 then
        return;
end if;
insert into dl.stock_def(
    stock_date, source, owner_code, supplier_article, barcode, warehouse, external_code, marketplace_id, barcode_id, item_id, org_id, item_name, def30, days_in_stock30, def5, days_in_stock5, avg_price, min_price, max_price
)
select v_day, source, owner_code, supplier_article, sd.barcode, warehouse, sd.external_code,
    mp.marketplace_id, b.barcode_id
     , i.id item_id, org.id org_id, i.name
     ,30 - sum(case when sd.quantity > 0 then 1 else 0 end ) def30
     ,count(distinct sd.stock_date) days_in_stock_30
     ,5 - sum(case when (sd.stock_date >  v_day - INTERVAL '5 day' and sd.quantity > 0) then 1 else 0 end ) def5
     ,count(distinct case when (sd.stock_date > v_day - INTERVAL '5 day') then sd.stock_date end) days_in_stock_5
     ,sum(sd.price) / count(sd.stock_date) avg_price
     ,min(sd.price) min_price
     ,max(sd.price) max_price
from dl.stock_daily sd
         left outer join ml.owner o on o.code = sd.owner_code
         left outer join dl.organisation org on org.id = o.organisation_id
         left outer join ml.marketplace mp on mp.code = sd.source
         left outer join dl.barcode b on b.barcode = sd.barcode and b.organisation_id = org.id and b.marketplace_id = mp.marketplace_id
         left outer join dl.item i on i.id = b.item_id
where sd.stock_date > v_day - INTERVAL '30 day' and sd.stock_date <= v_day
group by source, owner_code, supplier_article, sd.barcode, warehouse, sd.external_code, mp.marketplace_id, b.barcode_id, i.id, org.id, i.name;
end
$$;

alter procedure dl.calc_stock_def_by_day(date) owner to "user";

create or replace procedure dl.partition_for_sales_stock(IN start_date date, IN end_date date)
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

alter procedure dl.partition_for_sales_stock(date, date) owner to "user";

create or replace procedure dl.calc_sales_stock_by_day(IN p_day date)
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
select max(id) into v_t_id from ml.transaction where date_trunc('day', "start_date") = v_day and status='COMPLETED';

insert into dl.sales_stock(report_date, source, owner_code, supplier_article, barcode, warehouse_name, external_code, def30, days_in_stock30, def5, days_in_stock5, avg_price, min_price, max_price, quantity30,
                           quantity5, order_by_day30, forecast_order30, order_by_day5, forecast_order5,quantity
    ,item_id, item_name, marketplace_id, org_id, is_excluded, cluster)
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
     ,w.cluster
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
         left outer join dl.warehouse w on w.code = sd.warehouse and w.source = sd.source
where sd.stock_date=v_day;
end
$$;

alter procedure dl.calc_sales_stock_by_day(date) owner to "user";

create or replace procedure dl.partition_for_stock_cluster_def(IN start_date date, IN end_date date)
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

alter procedure dl.partition_for_stock_cluster_def(date, date) owner to "user";

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
     , s.source
     , s.owner_code
     , s.supplier_article
     , s.barcode
     , s.cluster
     , s.external_code
     , s.marketplace_id
     , s.barcode_id
     , s.item_id
     , s.org_id
     , s.name
     , 30 - sum(case when s.def > 0 then 1 else 0 end) def30
     , count(distinct s.stock_date)                    days_in_stock_30
     , 5 - sum(case
                   when (s.stock_date > v_day - INTERVAL '5 day' and s.def > 0) then 1
                   else 0 end)                         def5
     , count(distinct case
                          when (s.stock_date > v_day - INTERVAL '5 day')
                              then s.stock_date end)   days_in_stock_5
     , sum(s.avg_price) / count(s.stock_date)          avg_price
     , min(s.min_price)                                min_price
     , max(s.max_price)                                max_price
from (select sd.source
           , owner_code
           , supplier_article
           , sd.barcode
           , wh.cluster
           , sd.external_code
           , mp.marketplace_id
           , b.barcode_id
           , i.id                                             item_id
           , org.id                                           org_id
           , i.name
           , sd.stock_date
           , case when sum(sd.quantity) > 0 then 1 else 0 end def
           , sum(sd.price) / count(sd.stock_date)             avg_price
           , min(sd.price)                                    min_price
           , max(sd.price)                                    max_price
      from dl.stock_daily sd
               left outer join dl.warehouse wh on wh.source = sd.source and wh.code = sd.warehouse
               left outer join ml.owner o on o.code = sd.owner_code
               left outer join dl.organisation org on org.id = o.organisation_id
               left outer join ml.marketplace mp on mp.code = sd.source
               left outer join dl.barcode b on b.barcode = sd.barcode and b.organisation_id = org.id and
                                               b.marketplace_id = mp.marketplace_id
               left outer join dl.item i on i.id = b.item_id
      where sd.stock_date > v_day - INTERVAL '30 day'
        and sd.stock_date <= v_day
      group by sd.stock_date, sd.source, owner_code, supplier_article, sd.barcode, wh.cluster, sd.external_code,
          mp.marketplace_id,
          b.barcode_id, i.id, org.id, i.name) s
group by s.source
       , s.owner_code
       , s.supplier_article
       , s.barcode
       , s.cluster
       , s.external_code
       , s.marketplace_id
       , s.barcode_id
       , s.item_id
       , s.org_id
       , s.name;
end
$$;

alter procedure dl.calc_stock_cluster_def_by_day(date) owner to "user";

create or replace procedure dl.partition_for_report_by_cluster(IN start_date date, IN end_date date)
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
for v_table, v_from, v_to in SELECT 'report_by_cluster_' || to_char(day::date,'YYYYMM'),
                                    to_char(day::date,'YYYY-MM-DD'),
                                    to_char(day::date + interval '1 month','YYYY-MM-DD')
                             FROM generate_series(start_date, end_date, '1 month') day
            loop
            EXECUTE format(
                'CREATE TABLE IF NOT EXISTS "dl"."%s" PARTITION OF "dl"."report_by_cluster" FOR VALUES FROM(''%s'') TO (''%s'')',
                v_table,
                v_from,
                v_to
             );
end loop;
end;$$;

alter procedure dl.partition_for_report_by_cluster(date, date) owner to "user";

create or replace procedure dl.calc_report_by_cluster(IN p_day date)
    language plpgsql
as
$$
declare
v_t_id  ml.transaction.id%type;
    v_count int;
    v_day   date := date_trunc('day', p_day);
begin
select count(1) into v_count from dl.report_by_cluster where report_date = v_day;
if v_count > 0 then
        return;
end if;
select max(id) into v_t_id from ml.transaction where date_trunc('day', "start_date") = v_day and status='COMPLETED';
--select id into v_t_id from ml.transaction where id = v_t_id;

insert into dl.report_by_cluster(report_date, source, owner_code, supplier_article, barcode, external_code, def30,
                                 days_in_stock30, def5, days_in_stock5, avg_price, min_price, max_price, quantity30,
                                 quantity5, order_by_day30, forecast_order30, order_by_day5, forecast_order5,
                                 quantity
    , item_id, item_name, marketplace_id, org_id, is_excluded, cluster, transaction_id)
select sd.stock_date                                                                  report_date
     , sd.source
     , sd.owner_code
     , sd.supplier_article
     , sd.barcode
     , sd.external_code
     , sd.def30
     , sd.days_in_stock30
     , sd.def5
     , sd.days_in_stock5
     , sd.avg_price
     , sd.min_price
     , sd.max_price
     , oo.quantity30
     , oo.quantity5
     , case when sd.def30 = 30 then 0 else oo.quantity30 / (30 - sd.def30) end        order_by_day30
     , case when sd.def30 = 30 then 0 else (oo.quantity30 * 30) / (30 - sd.def30) end forecast_order30
     , case when sd.def5 = 5 then 0 else oo.quantity5 / (5 - sd.def5) end             order_by_day5
     , case when sd.def5 = 5 then 0 else (oo.quantity5 * 5) / (5 - sd.def5) end       forecast_order5
     , sdd.quantity                                                                   quantity
     , s.item_id                                                                      item_id
     , s.name                                                                         item_name
     , s.marketplace_id                                                               marketplace_id
     , s.organisation_id                                                              org_id
     , case when ei.barcode is null then false else true end                          is_exclud
     , sd.cluster
     , v_t_id
from dl.stock_cluster_def sd
         left outer join (select o.code owner_code,
                                 m.code source,
                                 item_id,
                                 barcode_id,
                                 barcode,
                                 b.organisation_id,
                                 b.marketplace_id,
                                 i.name
                          from dl.barcode b
                                   inner join dl.item i on i.id = b.item_id
                                   inner join ml.owner o on o.organisation_id = b.organisation_id
                                   inner join ml.marketplace m on m.marketplace_id = b.marketplace_id) s
                         on s.source = sd.source and s.owner_code = sd.owner_code and s.barcode = sd.barcode
         left outer join (select ssd.source,
                                 wh.cluster,
                                 ssd.owner_code,
                                 ssd.external_code,
                                 sum(ssd.quantity) quantity
                          from dl.stock_daily ssd
                                   left outer join dl.warehouse wh on wh.source = ssd.source and wh.code = ssd.warehouse
                          where ssd.stock_date = v_day
                          group by ssd.source, wh.cluster, ssd.owner_code, ssd.external_code) sdd
                         on sdd.cluster = sd.cluster and sdd.owner_code = sd.owner_code and
                            sdd.source = sd.source and sdd.external_code = sd.external_code
         left outer join
     (select owner_code
           , o.source
           , cluster
           , external_code
           , sum(total_price)                            total_price
           , sum(price_with_discount)                    price_with_discount
           , sum(quantity)                               quantity30
           , count(distinct order_date)                  order_date30
           , sum(case
                     when order_date > v_day - INTERVAL '5 day' then quantity
                 else 0 end)                         quantity5
           , count(distinct case
                                when order_date > v_day - INTERVAL '5 day'
                   then order_date end) order_date5
      from dl."order" o
               left outer join dl.warehouse w on w.source = o.source and w.code = o.warehouse_name
      where o.transaction_id = v_t_id
        and o.is_cancel != true
            and order_date > v_day - INTERVAL '30 day'
            and order_date <= v_day
      group by owner_code, o.source, w.cluster, external_code) oo
     on oo.external_code = sd.external_code and oo.cluster = sd.cluster and
        oo.owner_code = sd.owner_code and oo.source = sd.source
         left outer join dl.exclud_item ei
                         on ei.source = sd.source and ei.owner_code = sd.owner_code and ei.barcode = sd.barcode
where sd.stock_date = v_day;
end
$$;

alter procedure dl.calc_report_by_cluster(date) owner to "user";

create or replace procedure dl.partition_for_stock_item_def(IN start_date date, IN end_date date)
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

alter procedure dl.partition_for_stock_item_def(date, date) owner to "user";

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

alter procedure dl.calc_stock_item_def_by_day(date) owner to "user";

create or replace procedure dl.partition_for_report_by_item(IN start_date date, IN end_date date)
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

alter procedure dl.partition_for_report_by_item(date, date) owner to "user";

create or replace procedure dl.calc_report_by_item(IN p_day date)
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
    ,item_id, item_name, marketplace_id, org_id, is_excluded, quantity1c)

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
     ,s1c.quantity quantity1c
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
         left outer join dl.stock1c s1c on s1c.item_id = s.item_id and s1c.stock_date = v_day
where sd.stock_date=v_day;
end
$$;

alter procedure dl.calc_report_by_item(date) owner to "user";

insert into "ml"."job"("id","is_active","description","week_days","at_time","type") values (1, true, 'Test job','monday,tuesday,wednesday,thursday,friday,saturday,sunday','06:00,12:00,16:00,20:00','REGULAR');
commit;
