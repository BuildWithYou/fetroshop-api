/*==============================================================*/
/* Table: products                                              */
/*==============================================================*/
create table products (
   id                   SERIAL not null,
   store_id             INT4                 not null,
   brand_id             INT4                 not null,
   slug                 VARCHAR(255)         not null,
   name                 VARCHAR(255)         not null,
   is_active            BOOL                 not null default false,
   price                INT4                 not null,
   description          text                 null,
   minimum_purchase     INT4                 not null,
   varian_code          VARCHAR(64)          not null,
   sku                  VARCHAR(64)          not null,
   has_multiple_varian  BOOL                 not null,
   short_description    text                 null,
   weight               INT4                 not null,
   quantity             INT4                 not null,
   virtual_quantity     INT4                 not null,
   created_at           TIMESTAMP            not null,
   updated_at           TIMESTAMP            not null,
   deleted_at           TIMESTAMP            null,
   constraint PK_PRODUCTS primary key (id),
   constraint AK_PRODUCT_SLUG_UNIQU_PRODUCTS unique (slug)
);

alter table products
   add constraint FK_PRODUCTS_REFERENCE_STORES foreign key (store_id)
      references stores (id)
      on delete restrict on update cascade;

alter table products
   add constraint FK_PRODUCTS_REFERENCE_BRANDS foreign key (brand_id)
      references brands (id)
      on delete restrict on update cascade;
