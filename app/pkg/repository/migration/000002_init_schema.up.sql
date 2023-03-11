insert into "owner"("code","name") values('OWNER','TEST OWNER');
insert into "job"("owner_code", "is_active") values ('OWNER', true);
commit;