create index ordr__indx_tridsrc on "order" ("transaction_id", "source");

create index sls__idxotrsrcodid on "sale" ("transaction_id", "source", "odid");