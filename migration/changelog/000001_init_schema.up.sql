create table "owner"
(
    "code"        varchar(20) primary key,
    "name"        varchar(100) not null,
    "create_date" timestamp default now(),
    "is_deleted"  boolean      not null
);

create table "job"
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
on column "job"."is_active" is 'Активен ли';
comment
on column "job"."description" is 'Описание';
comment
on column "job"."week_days" is 'Дни недели monday | tuesday | wednesday | thursday | friday | saturday | sunday';
comment
on column "job"."at_time" is 'Время в формате 8:04;16:00';
comment
on column "job"."interval" is 'Интервал в секундах';
comment
on column "job"."max_runs" is 'Максимальное количество запусков';

create table "job_owner"
(
    "job_id"     integer references "job" ("id")         not null,
    "owner_code" varchar(20) references "owner" ("code") not null,
    primary key ("job_id", "owner_code")
);

create table "owner_marketplace"
(
    "owner_code" varchar(20) references "owner" ("code") not null,
    "source"     varchar(20)                             not null,
    "host"       varchar(200)                            not null,
    "password"   varchar(200),
    "client_id"  varchar(50),
    primary key ("owner_code", "source")
);

create table "transaction"
(
    "id"         serial primary key,
    "job_id"     integer references "job" ("id") not null,
    "start_date" timestamp                       not null,
    "end_date"   timestamp,
    "status"     varchar(20)
);

create table "stock"
(
    "id"                   serial primary key,
    "transaction_date"     date                                    not null,
    "transaction_id"       integer references "transaction" ("id") not null,
    "source"               varchar(20)                             not null,
    "owner_code"           varchar(20) references "owner" ("code") not null,
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
    "card_created"         date                                    not null
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
comment
on column "stock"."card_created" is 'Дата создания карточки';

create table "sale"
(
    "id"                  serial primary key,
    "transaction_date"    date                                    not null,
    "transaction_id"      integer references "transaction" ("id") not null,
    "source"              varchar(20)                             not null,
    "owner_code"          varchar(20) references "owner" ("code") not null,
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
on column "sale"."discount_percent" is 'Согласованный итоговый дисконт(процент)';
comment
on column "sale"."discount_value" is 'Согласованный итоговый дисконт(значение)';
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
    "id"                  serial primary key,
    "transaction_date"    date                                    not null,
    "transaction_id"      integer references "transaction" ("id") not null,
    "source"              varchar(20)                             not null,
    "owner_code"          varchar(20) references "owner" ("code") not null,
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
    "quantity"            integer                                 not null
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
on column "order"."discount_percent" is 'Согласованный итоговый дисконт(процент)';
comment
on column "order"."discount_value" is 'Согласованный итоговый дисконт(значение)';
comment
on column "order"."price_with_discount" is 'Цена, priceWithDiscount = totalPrice * (1 - discountPercent/100)';
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
on column "order"."status" is 'Статус заказа';
comment
on column "order"."cancel_dt" is 'Дата и время отмены заказа';
comment
on column "order"."g_number" is 'Номер заказа. Объединяет все позиции одного заказа.';
comment
on column "order"."sticker" is 'Цифровое значение стикера';
comment
on column "order"."srid" is 'Уникальный идентификатор заказа';
comment
on column "order"."quantity" is 'Количество';

create table "tlg_queue"
(
    "id"            serial primary key,
    "chat_id"       bigint      not null,
    "chat_type"     varchar(20) not null,
    "message_id"    bigint      not null,
    "message"       varchar(100),
    "message_date"  timestamp   not null,
    "response_date" timestamp,
    "status"        varchar(15) not null
);
comment
on table "tlg_queue" is 'Таблица для хранения сообщений, отправленных ботом в Telegram';
comment
on column "tlg_queue"."chat_id" is 'Идентификатор чата';
comment
on column "tlg_queue"."chat_type" is 'Тип чата';
comment
on column "tlg_queue"."message_id" is 'Идентификатор сообщения';
comment
on column "tlg_queue"."message" is 'Текст сообщения';
comment
on column "tlg_queue"."message_date" is 'Дата и время отправки сообщения';
comment
on column "tlg_queue"."response_date" is 'Дата и время получения ответа на сообщение';
comment
on column "tlg_queue"."status" is 'Статус сообщения(CREATED|PENDING|COMPLETED|ERROR)';

create table "tlg_event"
(
    "chat_id"   bigint      not null,
    "data_type" varchar(20) not null,
    primary key ("chat_id", "data_type")
);

create table "log_load"
(
    "transaction_id" integer references "transaction" ("id"),
    "owner_code"     varchar(20) references "owner" ("code"),
    "source"         varchar(20) not null,
    "description"    varchar(200),
    "status"         varchar(20) not null,
    primary key ("transaction_id", "owner_code", "source")
);

comment
on table "log_load" is 'Таблица для хранения логов загрузки данных';
comment
on column "log_load"."transaction_id" is 'Идентификатор транзакции';
comment
on column "log_load"."owner_code" is 'Код владельца';
comment
on column "log_load"."source" is 'Источник данных';
comment
on column "log_load"."description" is 'Описание';
comment
on column "log_load"."status" is 'Статус BEGIN|COMPLETED|ERROR';

create table "scheduler"
(
    "code"        varchar(20) primary key,
    "description" varchar(200),
    "status"      varchar(20) not null,
    "update_at"   timestamp   not null
);

create table "dict_items"
(
    "supplier_article" varchar(75),
    "barcode"          varchar(30),
    "name"             varchar(200),
    "subject"          varchar(200),
    "category"         varchar(200),
    "brand"            varchar(200),
    primary key ("supplier_article", "barcode")
);

comment
on table "dict_items" is 'Справочник товаров';
comment
on column "dict_items"."supplier_article" is 'Артикул поставщика';
comment
on column "dict_items"."barcode" is 'Штрихкод';
comment
on column "dict_items"."name" is 'Наименование';
comment
on column "dict_items"."subject" is 'Предмет';
comment
on column "dict_items"."category" is 'Категория';
comment
on column "dict_items"."brand" is 'Бренд';
