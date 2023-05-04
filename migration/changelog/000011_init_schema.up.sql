create table "marketplace"
(
    "id"         varchar(36)  not null primary key,
    "name"       varchar(255) not null,
    "article"    varchar(255),
    "updated_at" timestamp    not null
);

comment
on table "marketplace" is 'Marketplaces';
comment
on column "marketplace"."id" is 'ID';
comment
on column "marketplace"."name" is 'Name';
comment
on column "marketplace"."article" is 'Article';
comment
on column "marketplace"."updated_at" is 'Updated at';

create table "barcode"
(
    "id"              varchar(36)  not null primary key,
    "barcode"         varchar(255) not null,
    "organisation_id" varchar(36)  not null references "organisation" ("id"),
    "marketplace_id"  varchar(36)  not null references "marketplace" ("id"),
    "article"         varchar(255),
    "updated_at"      timestamp    not null
);

comment
on table "barcode" is 'Barcodes';
comment
on column "barcode"."id" is 'ID';
comment
on column "barcode"."barcode" is 'Barcode';
comment
on column "barcode"."organisation_id" is 'Organization ID';
comment
on column "barcode"."marketplace_id" is 'Marketplace ID';
comment
on column "barcode"."article" is 'Article';
comment
on column "barcode"."updated_at" is 'Updated at';

