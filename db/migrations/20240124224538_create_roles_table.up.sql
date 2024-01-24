/*==============================================================*/
/* Table: roles                                                 */
/*==============================================================*/
create table roles (
   id                   SERIAL not null,
   name                 VARCHAR(32)          not null,
   created_at           TIMESTAMP            not null,
   updated_at           TIMESTAMP            not null,
   deleted_at           TIMESTAMP            null,
   constraint PK_ROLES primary key (id),
   constraint AK_USERNAME_UNIQUE_ROLES unique (name)
);