alter table dl."order" add column item_id varchar(36);
alter table dl."order" add column barcode_id varchar(36);
alter table dl."order" add column message varchar(250);

alter table dl."stock" add column item_id varchar(36);
alter table dl."stock" add column barcode_id varchar(36);
alter table dl."stock" add column message varchar(250);