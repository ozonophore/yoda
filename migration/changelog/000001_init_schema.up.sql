create schema if not exists "dl";
create schema if not exists "ml";

set "ml".search_path to dl;
set search_path to dl,ml,public;

create table "ml"."owner"
(
    "code"        varchar(20) primary key,
    "name"        varchar(100) not null,
    "create_date" timestamp default now(),
    "is_deleted"  boolean      not null
);

create table "ml"."job"
(
    "id"          integer primary key,
    "create_date" timestamp default now(),
    "is_active"   boolean     not null,
    "description" varchar(100),
    "week_days"   varchar(200),
    "at_time"     varchar(200),
    "interval"    integer,
    "max_runs"    integer,
    "next_run"    timestamp,
    "last_run"    timestamp,
    "type"        varchar(20) not null
);

comment
on column "ml"."job"."is_active" is 'Активен ли';
comment
on column "ml"."job"."description" is 'Описание';
comment
on column "ml"."job"."week_days" is 'Дни недели monday | tuesday | wednesday | thursday | friday | saturday | sunday';
comment
on column "ml"."job"."at_time" is 'Время в формате 8:04;16:00';
comment
on column "ml"."job"."interval" is 'Интервал в секундах';
comment
on column "ml"."job"."max_runs" is 'Максимальное количество запусков';

create table "ml"."job_owner"
(
    "job_id"     integer references "ml"."job" ("id")         not null,
    "owner_code" varchar(20) references "ml"."owner" ("code") not null,
    primary key ("job_id", "owner_code")
);

create table "ml"."owner_marketplace"
(
    "owner_code" varchar(20) references "ml"."owner" ("code") not null,
    "source"     varchar(20)                             not null,
    "host"       varchar(200)                            not null,
    "password"   varchar(200),
    "client_id"  varchar(50),
    primary key ("owner_code", "source")
);

create table "ml"."transaction"
(
    "id"         serial primary key,
    "job_id"     integer references "ml"."job" ("id") not null,
    "start_date" timestamp                       not null,
    "end_date"   timestamp,
    "status"     varchar(20),
    "message"    varchar(256)
);

create table "dl"."stock"
(
    "id"                   serial primary key,
    "transaction_date"     date                                    not null,
    "transaction_id"       integer references "ml"."transaction" ("id") not null,
    "source"               varchar(20)                             not null,
    "owner_code"           varchar(20) references "ml"."owner" ("code") not null,
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
    "price"                numeric(10, 2),
    "discount"             numeric(10, 2),
    "price_after_discount" numeric(10, 2),
    "card_created"         date                                    not null,
    "item_id"              varchar(36),
    "barcode_id"           varchar(36),
    "message"              varchar(250)
);

create index stock_soec_idx on "dl"."stock" ("transaction_id", "source", "owner_code", "warehouse_name", "external_code");

comment
on column "dl"."stock"."last_change_date" is 'Дата обновления информации в сервисе';
comment
on column "dl"."stock"."last_change_time" is 'Время обновления информации в сервисе';
comment
on column "dl"."stock"."supplier_article" is 'Артикул поставщика';
comment
on column "dl"."stock"."tech_size" is 'Размер';
comment
on column "dl"."stock"."barcode" is 'Бар-код';
comment
on column "dl"."stock"."quantity" is 'Количество, доступное для продажи';
comment
on column "dl"."stock"."is_supply" is 'Договор поставки';
comment
on column "dl"."stock"."is_realization" is 'Договор реализации';
comment
on column "dl"."stock"."quantity_full" is 'Полное (непроданное) количество, которое числится за складом (= quantity + в пути)';
comment
on column "dl"."stock"."quantity_promised" is 'Количество товара, указанное в подтверждённых будущих поставках';
comment
on column "dl"."stock"."warehouse_name" is 'Название склада';
comment
on column "dl"."stock"."external_code" is 'Код WB/OZON';
comment
on column "dl"."stock"."name" is 'Наименование продукта';
comment
on column "dl"."stock"."subject" is 'Предмет';
comment
on column "dl"."stock"."category" is 'Категория';
comment
on column "dl"."stock"."days_on_site" is 'Количество дней на сайте';
comment
on column "dl"."stock"."brand" is 'Бренд';
comment
on column "dl"."stock"."SCCode" is 'Код контракта';
comment
on column "dl"."stock"."price" is 'Цена';
comment
on column "dl"."stock"."discount" is 'Скидка';
comment
on column "dl"."stock"."price_after_discount" is 'Цена после скидки';
comment
on column "dl"."stock"."card_created" is 'Дата создания карточки';

create table "dl"."sale"
(
    "id"                  serial primary key,
    "transaction_date"    date                                    not null,
    "transaction_id"      integer references "ml"."transaction" ("id") not null,
    "source"              varchar(20)                             not null,
    "owner_code"          varchar(20) references "ml"."owner" ("code") not null,
    "last_change_date"    date                                    not null,
    "last_change_time"    time                                    not null,
    "sale_date"           date                                    not null,
    "sale_time"           time                                    not null,
    "supplier_article"    varchar(75),
    "tech_size"           varchar(30),
    "barcode"             varchar(30),
    "total_price"         numeric(10, 2),
    "discount_percent"    numeric(10, 2),
    "discount_value"      numeric(10, 2),
    "is_supply"           boolean,
    "is_realization"      boolean,
    "promo_code_discount" float,
    "warehouse_name"      varchar(50),
    "country_name"        varchar(200),
    "oblast_okrug_name"   varchar(200),
    "region_name"         varchar(200),
    "income_id"           bigint,
    "sale_id"             varchar(15),
    "odid"                bigint,
    "spp"                 float,
    "for_pay"             float,
    "finished_price"      numeric(10, 2),
    "price_with_disc"     numeric(10, 2),
    "external_code"       varchar(50),
    "subject"             varchar(50),
    "category"            varchar(50),
    "brand"               varchar(50),
    "is_storno"           boolean,
    "g_number"            varchar(50),
    "sticker"             varchar(200),
    "srid"                varchar(50)
);

create index sls__idxotrsrcodid on "dl"."sale" ("transaction_id", "odid");

comment
on column "dl"."sale"."last_change_date" is 'Дата обновления информации в сервисе';
comment
on column "dl"."sale"."last_change_time" is 'Время обновления информации в сервисе';
comment
on column "dl"."sale"."sale_date" is 'Дата продажи';
comment
on column "dl"."sale"."sale_time" is 'Время продажи';
comment
on column "dl"."sale"."supplier_article" is 'Артикул поставщика';
comment
on column "dl"."sale"."tech_size" is 'Размер';
comment
on column "dl"."sale"."barcode" is 'Бар-код';
comment
on column "dl"."sale"."total_price" is 'Цена до согласованной итоговой скидки/промо/спп. Для получения цены со скидкой можно воспользоваться формулой priceWithDiscount = totalPrice * (1 - discountPercent/100)';
comment
on column "dl"."sale"."discount_percent" is 'Согласованный итоговый дисконт(процент)';
comment
on column "dl"."sale"."discount_value" is 'Согласованный итоговый дисконт(значение)';
comment
on column "dl"."sale"."is_supply" is 'Договор поставки';
comment
on column "dl"."sale"."is_realization" is 'Договор реализации';
comment
on column "dl"."sale"."promo_code_discount" is 'Скидка по промокоду';
comment
on column "dl"."sale"."warehouse_name" is 'Название склада отгрузки';
comment
on column "dl"."sale"."country_name" is 'Страна';
comment
on column "dl"."sale"."oblast_okrug_name" is 'Область';
comment
on column "dl"."sale"."region_name" is 'Регион';
comment
on column "dl"."sale"."income_id" is 'Номер поставки (от продавца на склад)';
comment
on column "dl"."sale"."sale_id" is 'Уникальный идентификатор продажи/возврата.';
comment
on column "dl"."sale"."odid" is 'Уникальный идентификатор позиции заказа';
comment
on column "dl"."sale"."spp" is 'Согласованная скидка постоянного покупателя';
comment
on column "dl"."sale"."for_pay" is 'К перечислению поставщику';
comment
on column "dl"."sale"."finished_price" is 'Фактическая цена заказа с учетом всех скидок';
comment
on column "dl"."sale"."price_with_disc" is 'Цена, от которой считается вознаграждение поставщика forpay';
comment
on column "dl"."sale"."external_code" is 'Код WB';
comment
on column "dl"."sale"."subject" is 'Предмет';
comment
on column "dl"."sale"."category" is 'Категория';
comment
on column "dl"."sale"."brand" is 'Бренд';
comment
on column "dl"."sale"."is_storno" is 'Для сторно-операций 1, для остальных 0';
comment
on column "dl"."sale"."g_number" is 'Номер заказа. Объединяет все позиции одного заказа.';
comment
on column "dl"."sale"."sticker" is 'Цифровое значение стикера';
comment
on column "dl"."sale"."srid" is 'Уникальный идентификатор заказа';

create table "dl"."order"
(
    "id"                  serial primary key,
    "transaction_date"    date                                    not null,
    "transaction_id"      integer references "ml"."transaction" ("id") not null,
    "source"              varchar(20)                             not null,
    "owner_code"          varchar(20) references "ml"."owner" ("code") not null,
    "last_change_date"    date                                    not null,
    "last_change_time"    time                                    not null,
    "order_date"          date                                    not null,
    "order_time"          time                                    not null,
    "supplier_article"    varchar(75),
    "tech_size"           varchar(30),
    "barcode"             varchar(30),
    "total_price"         numeric(10, 2),
    "discount_percent"    numeric(10, 2),
    "discount_value"      numeric(10, 2),
    "price_with_discount" numeric(10, 2),
    "warehouse_name"      varchar(50),
    "oblast"              varchar(200),
    "income_id"           bigint,
    "external_code"       varchar(50),
    "odid"                bigint,
    "subject"             varchar(200),
    "category"            varchar(200),
    "brand"               varchar(100),
    "is_cancel"           boolean,
    "status"              varchar(50),
    "cancel_dt"           timestamp,
    "g_number"            varchar(50),
    "sticker"             varchar(200),
    "srid"                varchar(50),
    "quantity"            integer                                 not null,
    "item_id"             varchar(36),
    "barcode_id"          varchar(36),
    "message"             varchar(250)
);

create index ordr__indx_tridsrc on "dl"."order" ("transaction_id", "source", "owner_code", "warehouse_name", "external_code", "srid");

comment
on column "dl"."order"."last_change_date" is 'Дата обновления информации в сервисе';
comment
on column "dl"."order"."last_change_time" is 'Время обновления информации в сервисе';
comment
on column "dl"."order"."order_date" is 'Дата заказа';
comment
on column "dl"."order"."order_time" is 'Время заказа';
comment
on column "dl"."order"."supplier_article" is 'Артикул поставщика';
comment
on column "dl"."order"."tech_size" is 'Размер';
comment
on column "dl"."order"."barcode" is 'Бар-код';
comment
on column "dl"."order"."total_price" is 'Цена до согласованной итоговой скидки/промо/спп. Для получения цены со скидкой можно воспользоваться формулой priceWithDiscount = totalPrice * (1 - discountPercent/100)';
comment
on column "dl"."order"."discount_percent" is 'Согласованный итоговый дисконт(процент)';
comment
on column "dl"."order"."discount_value" is 'Согласованный итоговый дисконт(значение)';
comment
on column "dl"."order"."price_with_discount" is 'Цена, priceWithDiscount = totalPrice * (1 - discountPercent/100)';
comment
on column "dl"."order"."warehouse_name" is 'Название склада отгрузки';
comment
on column "dl"."order"."oblast" is 'Область';
comment
on column "dl"."order"."income_id" is 'Номер поставки (от продавца на склад)';
comment
on column "dl"."order"."external_code" is 'Код WB';
comment
on column "dl"."order"."odid" is 'Уникальный идентификатор позиции заказа';
comment
on column "dl"."order"."subject" is 'Предмет';
comment
on column "dl"."order"."category" is 'Категория';
comment
on column "dl"."order"."brand" is 'Бренд';
comment
on column "dl"."order"."is_cancel" is 'Отмена заказа. true - заказ отменен до оплаты';
comment
on column "dl"."order"."status" is 'Статус заказа';
comment
on column "dl"."order"."cancel_dt" is 'Дата и время отмены заказа';
comment
on column "dl"."order"."g_number" is 'Номер заказа. Объединяет все позиции одного заказа.';
comment
on column "dl"."order"."sticker" is 'Цифровое значение стикера';
comment
on column "dl"."order"."srid" is 'Уникальный идентификатор заказа';
comment
on column "dl"."order"."quantity" is 'Количество';

create table "ml"."log_load"
(
    "transaction_id" integer references "ml"."transaction" ("id"),
    "owner_code"     varchar(20) references "ml"."owner" ("code"),
    "source"         varchar(20) not null,
    "description"    varchar(200),
    "status"         varchar(20) not null,
    primary key ("transaction_id", "owner_code", "source")
);

comment
on table "ml"."log_load" is 'Таблица для хранения логов загрузки данных';
comment
on column "ml"."log_load"."transaction_id" is 'Идентификатор транзакции';
comment
on column "ml"."log_load"."owner_code" is 'Код владельца';
comment
on column "ml"."log_load"."source" is 'Источник данных';
comment
on column "ml"."log_load"."description" is 'Описание';
comment
on column "ml"."log_load"."status" is 'Статус BEGIN|COMPLETED|ERROR';

create table "ml"."scheduler"
(
    "code"        varchar(20) primary key,
    "description" varchar(200),
    "status"      varchar(20) not null,
    "update_at"   timestamp   not null
);

create table "dl"."item"
(
    "id"        varchar(36)  not null primary key,
    "name"      varchar(200) not null,
    "update_at" timestamp
);

comment
on table "dl"."item" is 'Таблица для хранения справочника товаров';
    comment
on column "dl"."item"."id" is 'Идентификатор';
    comment
on column "dl"."item"."name" is 'Наименование';
    comment
on column "dl"."item"."update_at" is 'Дата обновления';

insert into "ml"."job"("id","is_active","description","week_days","at_time","type") values (1, true, 'Test job','monday,tuesday,wednesday,thursday,friday,saturday,sunday','06:00,12:00,16:00,20:00','REGULAR');
commit;
