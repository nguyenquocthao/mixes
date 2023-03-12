create table if not exists folders(
    id integer primary key,
    parent_id integer references folders(id)
);

create or replace function check_folder_cycle( parentID int )
returns boolean
as
$$
declare 
    fastp int;
    slowp int;
begin
    fastp := parentID;
    slowp := parentID;
    loop
        select parent_id into slowp from folders where id=slowp;
        select t2.parent_id into fastp from folders t1 join folders t2 on t1.parent_id = t2.id
        where t1.id = fastp;
        if fastp is null then return false; end if;
        if slowp = fastp then return true; end if;
    end loop;
end; 
$$ language plpgsql;

CREATE or replace FUNCTION trigger_folder_parent_id() RETURNS TRIGGER AS $$
BEGIN
    if new.parent_id is null then return new; end if;
    if check_folder_cycle(new.parent_id) then raise exception 'circular referencing parent_id'; end if;
    return new;
END;
$$ language plpgsql;

create trigger folder_parent_id after insert or update on folders for each row execute procedure trigger_folder_parent_id();