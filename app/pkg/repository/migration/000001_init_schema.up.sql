create table "owner"
(
    "code"        varchar(20) primary key,
    "name"        varchar(100) not null,
    "create_date" timestamp default now()
);

create table "job"
(
    "id"          serial primary key,
    "owner_code"  varchar(20) references "owner" ("code") not null,
    "create_date" timestamp default now(),
    "is_active"   boolean                                 not null,
    "description" varchar(100)
);

create table "job_marketplace"
(
    "job_id"    integer references "job" ("id") not null,
    "source"    varchar(20)                     not null,
    "host"      varchar(200)                    not null,
    "password"  varchar(200),
    "client_id" varchar(50)
);

create table "transaction"
(
    "id"         serial primary key,
    "owner_code" varchar(20) references "owner" ("code") not null,
    "job_id"     integer references "job" ("id")         not null,
    "start_date" timestamp                               not null,
    "end_date"   timestamp,
    "source"     varchar(20),
    "status"     varchar(20)
);

create table "stock"
(
    "id"                   serial primary key,
    "transaction_id"       integer references "transaction" ("id") not null,
    "source"               varchar(20)                             not null,
    "last_change_date"     date                                    not null,
    "last_change_time"     time                                    not null,
    "supplier_article"     varchar(75),
    "tech_size"            varchar(30),
    "barcode"              varchar(30),
    "quantity"             integer                                 not null,
    "is_supply"            boolean,
    "is_realization"       boolean,
    "quantity_full"        integer                                 not null,
    "quantity_promised"    integer,
    "warehouse_name"       varchar(50)                             not null,
    "external_code"        varchar(50),
    "name"                 varchar(200),
    "subject"              varchar(200),
    "category"             varchar(50),
    "days_on_site"         integer,
    "brand"                varchar(50),
    "SCCode"               varchar(50),
    "price"                float,
    "discount"             float,
    "price_after_discount" float
);

comment
on column "stock"."last_change_date" is 'Дата обновления информации в сервисе';
comment
on column "stock"."last_change_time" is 'Время обновления информации в сервисе';
comment
on column "stock"."supplier_article" is 'Артикул поставщика';
comment
on column "stock"."tech_size" is 'Размер';
comment
on column "stock"."barcode" is 'Бар-код';
comment
on column "stock"."quantity" is 'Количество, доступное для продажи';
comment
on column "stock"."is_supply" is 'Договор поставки';
comment
on column "stock"."is_realization" is 'Договор реализации';
comment
on column "stock"."quantity_full" is 'Полное (непроданное) количество, которое числится за складом (= quantity + в пути)';
comment
on column "stock"."quantity_promised" is 'Количество товара, указанное в подтверждённых будущих поставках';
comment
on column "stock"."warehouse_name" is 'Название склада';
comment
on column "stock"."external_code" is 'Код WB/OZON';
comment
on column "stock"."name" is 'Наименование продукта';
comment
on column "stock"."subject" is 'Предмет';
comment
on column "stock"."category" is 'Категория';
comment
on column "stock"."days_on_site" is 'Количество дней на сайте';
comment
on column "stock"."brand" is 'Бренд';
comment
on column "stock"."SCCode" is 'Код контракта';
comment
on column "stock"."price" is 'Цена';
comment
on column "stock"."discount" is 'Скидка';
comment
on column "stock"."price_after_discount" is 'Цена после скидки';

create table "sale"
(
    "id"                  serial primary key,
    "transaction_id"      integer references "transaction" ("id") not null,
    "source"              varchar(20)                             not null,
    "last_change_date"    date                                    not null,
    "last_change_time"    time                                    not null,
    "sale_date"           date                                    not null,
    "sale_time"           time                                    not null,
    "supplier_article"    varchar(75),
    "tech_size"           varchar(30),
    "barcode"             varchar(30),
    "total_price"         money,
    "discount_percent"    integer,
    "is_supply"           boolean,
    "is_realization"      boolean,
    "promo_code_discount" float4,
    "warehouse_name"      varchar(50),
    "country_name"        varchar(200),
    "oblast_okrug_name"   varchar(200),
    "region_name"         varchar(200),
    "income_id"           integer,
    "sale_id"             varchar(15),
    "odid"                integer,
    "spp"                 float,
    "for_pay"             float,
    "finished_price"      money,
    "price_with_disc"     money,
    "external_code"       varchar(50),
    "subject"             varchar(50),
    "category"            varchar(50),
    "brand"               varchar(50),
    "is_storno"           boolean,
    "g_number"            varchar(50),
    "sticker"             varchar(200),
    "srid"                varchar(50)
);

comment
on column "sale"."last_change_date" is 'Дата обновления информации в сервисе';
comment
on column "sale"."last_change_time" is 'Время обновления информации в сервисе';
comment
on column "sale"."sale_date" is 'Дата продажи';
comment
on column "sale"."sale_time" is 'Время продажи';
comment
on column "sale"."supplier_article" is 'Артикул поставщика';
comment
on column "sale"."tech_size" is 'Размер';
comment
on column "sale"."barcode" is 'Бар-код';
comment
on column "sale"."total_price" is 'Цена до согласованной итоговой скидки/промо/спп. Для получения цены со скидкой можно воспользоваться формулой priceWithDiscount = totalPrice * (1 - discountPercent/100)';
comment
on column "sale"."discount_percent" is 'Согласованный итоговый дисконт';
comment
on column "sale"."is_supply" is 'Договор поставки';
comment
on column "sale"."is_realization" is 'Договор реализации';
comment
on column "sale"."promo_code_discount" is 'Скидка по промокоду';
comment
on column "sale"."warehouse_name" is 'Название склада отгрузки';
comment
on column "sale"."country_name" is 'Страна';
comment
on column "sale"."oblast_okrug_name" is 'Область';
comment
on column "sale"."region_name" is 'Регион';
comment
on column "sale"."income_id" is 'Номер поставки (от продавца на склад)';
comment
on column "sale"."sale_id" is 'Уникальный идентификатор продажи/возврата.';
comment
on column "sale"."odid" is 'Уникальный идентификатор позиции заказа';
comment
on column "sale"."spp" is 'Согласованная скидка постоянного покупателя';
comment
on column "sale"."for_pay" is 'К перечислению поставщику';
comment
on column "sale"."finished_price" is 'Фактическая цена заказа с учетом всех скидок';
comment
on column "sale"."price_with_disc" is 'Цена, от которой считается вознаграждение поставщика forpay';
comment
on column "sale"."external_code" is 'Код WB';
comment
on column "sale"."subject" is 'Предмет';
comment
on column "sale"."category" is 'Категория';
comment
on column "sale"."brand" is 'Бренд';
comment
on column "sale"."is_storno" is 'Для сторно-операций 1, для остальных 0';
comment
on column "sale"."g_number" is 'Номер заказа. Объединяет все позиции одного заказа.';
comment
on column "sale"."sticker" is 'Цифровое значение стикера';
comment
on column "sale"."srid" is 'Уникальный идентификатор заказа';

create table "order"
(
    "id"               serial primary key,
    "transaction_id"   integer references "transaction" ("id") not null,
    "source"           varchar(20)                             not null,
    "last_change_date" date                                    not null,
    "last_change_time" time                                    not null,
    "order_date"       date                                    not null,
    "order_time"       time                                    not null,
    "supplier_article" varchar(75),
    "tech_size"        varchar(30),
    "barcode"          varchar(30),
    "total_price"      money,
    "discount_percent" integer,
    "warehouse_name"   varchar(50),
    "oblast"           varchar(200),
    "income_id"        integer,
    "external_code"    varchar(50),
    "odid"             integer,
    "subject"          varchar(50),
    "category"         varchar(50),
    "brand"            varchar(50),
    "is_cancel"        boolean,
    "cancel_dt"        timestamp,
    "g_number"         varchar(50),
    "sticker"          varchar(200),
    "srid"             varchar(50)
);

comment
on column "order"."last_change_date" is 'Дата обновления информации в сервисе';
comment
on column "order"."last_change_time" is 'Время обновления информации в сервисе';
comment
on column "order"."order_date" is 'Дата заказа';
comment
on column "order"."order_time" is 'Время заказа';
comment
on column "order"."supplier_article" is 'Артикул поставщика';
comment
on column "order"."tech_size" is 'Размер';
comment
on column "order"."barcode" is 'Бар-код';
comment
on column "order"."total_price" is 'Цена до согласованной итоговой скидки/промо/спп. Для получения цены со скидкой можно воспользоваться формулой priceWithDiscount = totalPrice * (1 - discountPercent/100)';
comment
on column "order"."discount_percent" is 'Согласованный итоговый дисконт';
comment
on column "order"."warehouse_name" is 'Название склада отгрузки';
comment
on column "order"."oblast" is 'Область';
comment
on column "order"."income_id" is 'Номер поставки (от продавца на склад)';
comment
on column "order"."external_code" is 'Код WB';
comment
on column "order"."odid" is 'Уникальный идентификатор позиции заказа';
comment
on column "order"."subject" is 'Предмет';
comment
on column "order"."category" is 'Категория';
comment
on column "order"."brand" is 'Бренд';
comment
on column "order"."is_cancel" is 'Отмена заказа. true - заказ отменен до оплаты';
comment
on column "order"."cancel_dt" is 'Дата и время отмены заказа';
comment
on column "order"."g_number" is 'Номер заказа. Объединяет все позиции одного заказа.';
comment
on column "order"."sticker" is 'Цифровое значение стикера';
comment
on column "order"."srid" is 'Уникальный идентификатор заказа';
