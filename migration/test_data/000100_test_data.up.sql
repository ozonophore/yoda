BEGIN

insert into "owner"("code","name", "is_deleted") values('OWNER','TEST OWNER', false);
insert into "job_owner"("owner_code", "job_id") values ('OWNER', 1);

insert into "owner_marketplace"("owner_code",   "source", "host", "client_id", "password")
values ('OWNER', 'WB', 'http://localhost:1080/wb', null, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NJRCI6IjFiMzVmODljLTMyNGYtNGM3OS05NzhhLTkwMmYwODk3Mjc4YiJ9.WeYv1vqA46_9D5up2LRUeSBZCXxSBNcmH8lUhG9Jii0');

insert into "owner_marketplace"("owner_code",   "source", "host", "client_id", "password")
values ('OWNER', 'OZON', 'http://localhost:1080/ozon', '538358', '8539be7e-a37f-4b4f-b5e1-3879e5f1738c');

commit;