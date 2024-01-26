/*==============================================================*/
/* Table: customer_accesses                                     */
/*==============================================================*/
create table customer_accesses (
   id                   SERIAL not null,
   customer_id          INT4                 not null,
   key                  VARCHAR(255)          not null,
   platform             VARCHAR(32)          null,
   user_agent           VARCHAR(255)         null,
   expired_at           TIMESTAMP            not null,
   created_at           TIMESTAMP            not null,
   updated_at           TIMESTAMP            not null,
   deleted_at           TIMESTAMP            null,
   constraint PK_CUSTOMER_ACCESSES primary key (id)
);

alter table customer_accesses
   add constraint FK_CUSTOMER_REFERENCE_CUSTOMER foreign key (customer_id)
      references customers (id)
      on delete cascade on update cascade;