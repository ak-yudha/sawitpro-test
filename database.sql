/**
  This is the SQL script that will be used to initialize the database schema.
  We will evaluate you based on how well you design your database.
  1. How you design the tables.
  2. How you choose the data types and keys.
  3. How you name the fields.
  In this assignment we will use PostgreSQL as the database.
  */

/** This is test table. Remove this table and replace with your own tables. */
CREATE TABLE users (
                       id            serial primary key,
                       full_name     varchar(200)      not null,
                       phone_number  varchar(200)      not null unique,
                       password      varchar(200)      not null,
                       counter integer default 0 not null
);

INSERT INTO users (full_name, phone_number, password, login_counter) VALUES ('1', '+628561119023', '111', 1);
INSERT INTO users (full_name, phone_number, password, login_counter) VALUES ('2', '+625693213213', '222', 0);
