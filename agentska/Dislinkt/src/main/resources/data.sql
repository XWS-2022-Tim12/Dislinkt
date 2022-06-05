delete from companies;
delete from users;

insert into users (id, first_name, password, email, birth_day, mobile_number, username, gender, role)
values ('62fghcc3a34d25d8567f9f82', 'Prvi', 'prvasifra', 'prvi@gmail.com', DATE('2022-03-05'), '0645851733', 'mico', 0, 2);
insert into users (id, first_name, password, email, birth_day, mobile_number, username, gender, role)
values ('h348ch3748f34834hcu34hf8', 'Drugi', 'drugasifra', 'drugi@gmail.com', DATE('2000-02-07'), '067841554', 'drugi', 0, 1);
insert into users (id, first_name, password, email, birth_day, mobile_number, username, gender, role)
values ('ujf5524fj2458fj2458fj458', 'Treci', 'trecasifra', 'treci@gmail.com', DATE('1975-07-04'), '063425454', 'treci', 0, 0);

insert into companies (id, name, address, phone_number, description, owner_id, approved)
values ('we78ycq7834qcy3487cyq347', 'First company', 'Address', '061490448', 'Description', '62fghcc3a34d25d8567f9f82', true);
insert into companies (id, name, address, phone_number, description, owner_id, approved)
values ('f45ujf854u8f4582f458845j', 'Second company', 'Address', '064490448', 'Description', 'h348ch3748f34834hcu34hf8', false);