
/*==============================================================*/
/* Table: customers                                             */
/*==============================================================*/
create table customers (
   id                   SERIAL not null,
   username             VARCHAR(32)          not null,
   phone                VARCHAR(32)          null,
   email                VARCHAR(32)          null,
   full_name            VARCHAR(255)         null,
   password             VARCHAR(255)         null,
   created_at           TIMESTAMP            not null,
   updated_at           TIMESTAMP            not null,
   deleted_at           TIMESTAMP            null,
   constraint PK_CUSTOMERS primary key (id),
   constraint AK_USERNAME_UNIQUE_CUSTOMER unique (username),
   constraint AK_PHONE_UNIQUE_CUSTOMER unique (phone),
   constraint AK_EMAIL_UNIQUE_CUSTOMER unique (email)
);

/* customers table data init */
   INSERT INTO customers VALUES (1, 'testerwebapi', '081234567890', 'tester@mail.com', 'Tester Web API', '$2a$10$UQy61eWcNaOfzrINghwLO.DHmIsTTgWidWKmDziberHOVUi4NLV4W', '2024-01-18 09:45:12.736703', '2024-01-18 09:45:12.736703', NULL);
   SELECT setval('customers_id_seq', (SELECT MAX(id) FROM customers));
