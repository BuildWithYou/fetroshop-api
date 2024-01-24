/*==============================================================*/
/* Table: stores                                                */
/*==============================================================*/
create table stores (
   id                   INT4                 not null,
   user_id              INT4                 not null,
   code                 VARCHAR(64)          not null,
   name                 VARCHAR(64)          not null,
   is_active            BOOL                 not null default false,
   icon                 VARCHAR(255)         null,
   latitude             VARCHAR(64)          null,
   longitude            VARCHAR(64)          null,
   address              VARCHAR(255)         null,
   province_id          INT4                 null,
   city_id              INT4                 null,
   district_id          INT4                 null,
   subdistrict_id       INT4                 null,
   postal_code          VARCHAR(16)          null,
   created_at           TIMESTAMP            not null,
   updated_at           TIMESTAMP            not null,
   deleted_at           TIMESTAMP            null,
   constraint PK_STORES primary key (id),
   constraint AK_STORE_CODE_UNIQUE_STORES unique (code)
);

alter table stores
   add constraint FK_STORES_REFERENCE_USERS foreign key (user_id)
      references users (id)
      on delete restrict on update cascade;