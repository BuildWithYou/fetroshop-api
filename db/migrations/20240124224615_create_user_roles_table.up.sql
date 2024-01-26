/*==============================================================*/
/* Table: user_roles                                            */
/*==============================================================*/
create table user_roles (
   role_id              INT4                 not null,
   user_id              INT4                 not null,
   created_at           TIMESTAMP            not null,
   updated_at           TIMESTAMP            not null,
   deleted_at           TIMESTAMP            null,
   constraint PK_USER_ROLES primary key (role_id, user_id)
);

alter table user_roles
   add constraint FK_USER_ROLES_TO_ROLES foreign key (role_id)
      references roles (id)
      on delete restrict on update cascade;

alter table user_roles
   add constraint FK_USER_ROLES_TO_USERS foreign key (user_id)
      references users (id)
      on delete restrict on update cascade;