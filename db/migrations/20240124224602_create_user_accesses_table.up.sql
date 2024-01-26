/*==============================================================*/
/* Table: user_accesses                                     */
/*==============================================================*/
create table user_accesses (
   id                   SERIAL not null,
   user_id              INT4                 not null,
   key                  VARCHAR(255)         not null,
   platform             VARCHAR(32)          null,
   user_agent           VARCHAR(255)         null,
   expired_at           TIMESTAMP            not null,
   created_at           TIMESTAMP            not null,
   updated_at           TIMESTAMP            not null,
   deleted_at           TIMESTAMP            null,
   constraint PK_USER_ACCESSES primary key (id)
);

alter table user_accesses
   add constraint FK_USER_ACC_REFERENCE_USERS foreign key (user_id)
      references users (id)
      on delete cascade on update cascade;