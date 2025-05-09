/*==============================================================*/
/* Table: stores                                                */
/*==============================================================*/
create table stores (
   id                   SERIAL               not null,
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
   add constraint FK_STORES_REFERENCE_CITIES foreign key (city_id)
      references cities (id)
      on delete restrict on update cascade;

alter table stores
   add constraint FK_STORES_REFERENCE_DISTRICT foreign key (district_id)
      references districts (id)
      on delete restrict on update cascade;

alter table stores
   add constraint FK_STORES_REFERENCE_SUBDISTR foreign key (subdistrict_id)
      references subdistricts (id)
      on delete restrict on update cascade;

alter table stores
   add constraint FK_STORES_REFERENCE_USERS foreign key (user_id)
      references users (id)
      on delete restrict on update cascade;

alter table stores
   add constraint FK_STORES_REFERENCE_PROVINCE foreign key (province_id)
      references provinces (id)
      on delete restrict on update cascade;