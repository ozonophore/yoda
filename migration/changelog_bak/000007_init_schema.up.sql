create index ordr__indx_tridsrc on "order" ("transaction_id", "source", "owner_code", "warehouse_name", "external_code", "srid");

create index sls__idxotrsrcodid on "sale" ("transaction_id", "odid");