create table "marketplace"
(
    "id"         varchar(36)  not null primary key,
    "name"       varchar(255) not null,
    "updated_at" timestamp    not null
);

comment
on table "marketplace" is 'Marketplaces';
comment
on column "marketplace"."id" is 'ID';
comment
on column "marketplace"."name" is 'Name';
comment
on column "marketplace"."updated_at" is 'Updated at';

create table "barcode"
(
    "item_id"         varchar(36)  not null references "item" ("id") primary key,
    "barcode_id"      varchar(36)  not null,
    "barcode"         varchar(255) not null,
    "organisation_id" varchar(36)  not null references "organisation" ("id"),
    "marketplace_id"  varchar(36)  not null references "marketplace" ("id"),
    "updated_at"      timestamp    not null
);

comment
on table "barcode" is 'Barcodes';
comment
on column "barcode"."barcode_id" is 'ID';
comment
on column "barcode"."item_id" is 'Item ID';
comment
on column "barcode"."barcode" is 'Barcode';
comment
on column "barcode"."organisation_id" is 'Organisation ID';
comment
on column "barcode"."marketplace_id" is 'Marketplace ID';
