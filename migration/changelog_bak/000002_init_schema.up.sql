create or replace procedure partitionByDay(IN start_date date, IN end_date date, IN table_name varchar)
    language plpgsql
as
$$
declare
v_table varchar(8);
v_from varchar(10);
v_to varchar(10);
begin
for v_table, v_from, v_to in SELECT to_char(day::date,'YYYYMMDD'),to_char(day::date,'YYYY-MM-DD'),to_char(day::date+1,'YYYY-MM-DD')
                             FROM generate_series(start_date, end_date, '1 day') day
            loop
            EXECUTE format(
                'CREATE TABLE IF NOT EXISTS "%s_%s" PARTITION OF %s FOR VALUES FROM(''%s'') TO (''%s'')',
                table_name,
                v_table,
                table_name,
                v_from,
                v_to
             );
end loop;
end;$$;