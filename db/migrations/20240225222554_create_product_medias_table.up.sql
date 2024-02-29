/*==============================================================*/
/* Table: product_medias                                        */
/*==============================================================*/
create table product_medias (
   id                   INT4                 not null,
   product_id           INT4                 not null,
   file                 TEXT                 not null,
   media_type           VARCHAR(16)          not null,
   size                 VARCHAR(1)           not null,
   created_at           TIMESTAMP            not null,
   updated_at           TIMESTAMP            not null,
   deleted_at           TIMESTAMP            null,
   constraint PK_PRODUCT_MEDIAS primary key (id)
);

comment on column product_medias.media_type is
'Available options: image_file, image_url, video_file, video_url';

comment on column product_medias.size is
'Available options: S, M, L, XL';

alter table product_medias
   add constraint FK_PRODUCT__REFERENCE_PRODUCTS foreign key (product_id)
      references products (id)
      on delete restrict on update cascade;