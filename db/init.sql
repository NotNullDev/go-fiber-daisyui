create table products
(
  id    integer,
  name  text    not null,
  price integer not null
);

create table sqlite_master
(
  type     TEXT,
  name     TEXT,
  tbl_name TEXT,
  rootpage INT,
  sql      TEXT
);
