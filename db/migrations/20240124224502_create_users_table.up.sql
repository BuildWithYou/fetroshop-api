/*==============================================================*/
/* Table: users                                                 */
/*==============================================================*/
create table users (
   id                   SERIAL not null,
   username             VARCHAR(32)          not null,
   phone                VARCHAR(32)          null,
   email                VARCHAR(32)          null,
   full_name            VARCHAR(255)         null,
   password             VARCHAR(255)         null,
   created_at           TIMESTAMP            not null,
   updated_at           TIMESTAMP            not null,
   deleted_at           TIMESTAMP            null,
   constraint PK_USERS primary key (id),
   constraint AK_USERNAME_UNIQUE_USERS unique (username),
   constraint AK_PHONE_UNIQUE_USERS unique (phone),
   constraint AK_EMAIL_UNIQUE_USERS unique (email)
);

/* users table data init */
   INSERT INTO users VALUES (2, 'testercmsapi', '081234567890', 'tester@mail.com', 'Tester CMS API', '$2a$10$x.S5GJvGqw4L5366USju6.I2fISEOEyqPFswmYQiX/fF.ZjOcYChO', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL);
   SELECT setval('users_id_seq', (SELECT MAX(id) FROM users));
