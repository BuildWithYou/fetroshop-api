
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

/*==============================================================*/
/* Table: categories                                            */
/*==============================================================*/
create table categories (
   id                   SERIAL not null,
   parent_id            INT4                 null,
   code                 VARCHAR(64)          not null,
   name                 VARCHAR(64)          not null,
   is_active            BOOL                 not null default false,
   icon                 VARCHAR(255)         null,
   display_order        INT4                 null,
   created_at           TIMESTAMP            not null,
   updated_at           TIMESTAMP            not null,
   deleted_at           TIMESTAMP            null,
   constraint PK_CATEGORIES primary key (id),
   constraint AK_CATEGORY_CODE_UNIQ_CATEGORI unique (code)
);

-- tables alter
alter table user_accesses
   add constraint FK_USER_ACC_REFERENCE_USERS foreign key (user_id)
      references users (id)
      on delete cascade on update cascade;

alter table customer_accesses
   add constraint FK_CUSTOMER_REFERENCE_CUSTOMER foreign key (customer_id)
      references customers (id)
      on delete cascade on update cascade;

alter table user_roles
   add constraint FK_USER_ROLES_TO_ROLES foreign key (role_id)
      references roles (id)
      on delete restrict on update cascade;

alter table user_roles
   add constraint FK_USER_ROLES_TO_USERS foreign key (user_id)
      references users (id)
      on delete restrict on update cascade;


-- data initialization
/* categories table data init */
INSERT INTO categories (id, parent_id, code, name, is_active, icon, display_order, created_at, updated_at) VALUES
(1, null, 'kebutuhan-dapur', 'Kebutuhan Dapur', true, null, 1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(2, null, 'kebutuhan-ibu-anak', 'Kebutuhan Ibu & Anak', true, null, 2, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(3, null, 'kebutuhan-rumah', 'Kebutuhan Rumah', true, null, 3, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(4, null, 'makanan', 'Makanan', true, null, 4, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(5, null, 'minuman', 'Minuman', true, null, 5, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(6, null, 'produk-segar-beku', 'Produk Segar & Beku', true, null, 6, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(7, null, 'personal-care', 'Personal Care', true, null, 7, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(8, null, 'kebutuhan-kesehatan', 'Kebutuhan Kesehatan', true, null, 8, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(9, null, 'lifestyle', 'Lifestyle', true, null, 9, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(10, null, 'pet-foods', 'Pet Foods', true, null, 10, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(11, 1, 'perlengkapan-dapur-ruang-makan', 'Perlengkapan Dapur & Ruang Makan', true, null, 11, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(12, 1, 'bahan-masakan', 'Bahan Masakan', true, null, 12, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(13, 1, 'bahan-roti-kue', 'Bahan Roti & Kue', true, null, 13, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(14, 1, 'bahan-puding-agar-agar', 'Bahan Puding & Agar-Agar', true, null, 14, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(15, 2, 'makanan-bayi-anak', 'Makanan Bayi & Anak', true, null, 15, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(16, 2, 'susu-formula-bayi-anak', 'Susu Formula Bayi & Anak', true, null, 16, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(17, 2, 'perlengkapan-mandi-perawatan-anak', 'Perlengkapan Mandi & Perawatan Anak', true, null, 17, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(18, 2, 'susu-ibu-hamil-menyusui', 'Susu Ibu Hamil & Menyusui', true, null, 18, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(19, 2, 'popok-bayi-anak', 'Popok Bayi & Anak', true, null, 19, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(20, 2, 'pembersih-pakaian-perlengkapan-anak', 'Pembersih Pakaian & Perlengkapan Anak', true, null, 20, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(21, 2, 'perlengkapan-makan-minum-anak', 'Perlengkapan Makan & Minum Anak', true, null, 21, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(22, 2, 'kebutuhan-ibu-bayi-anak-lainnya', 'Kebutuhan Ibu, Bayi, & Anak Lainnya', true, null, 22, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(23, 3, 'perlengkapan-kamar-mandi', 'Perlengkapan Kamar Mandi', true, null, 23, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(24, 3, 'perawatan-pembersih', 'Perawatan & Pembersih', true, null, 24, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(25, 3, 'perlengkapan-rumah', 'Perlengkapan Rumah', true, null, 25, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(26, 3, 'tisu', 'Tisu', true, null, 26, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(27, 3, 'pengharum-ruangan-anti-lembab', 'Pengharum Ruangan & Anti Lembab', true, null, 27, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(28, 3, 'pembasmi-hama-serangga', 'Pembasmi Hama & Serangga', true, null, 28, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(29, 4, 'makanan-ringan', 'Makanan Ringan', true, null, 29, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(30, 4, 'makanan-instan', 'Makanan Instan', true, null, 30, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(31, 4, 'roti-selai-sereal', 'Roti, Selai, & Sereal', true, null, 31, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(32, 5, 'minuman-ringan', 'Minuman Ringan', true, null, 32, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(33, 5, 'produk-olahan-susu', 'Produk Olahan Susu', true, null, 33, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(34, 5, 'minuman-instan', 'Minuman Instan', true, null, 34, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(35, 6, 'makanan-siap-saji', 'Makanan Siap Saji', true, null, 35, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(36, 6, 'minuman-siap-saji', 'Minuman Siap Saji', true, null, 36, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(37, 6, 'makanan-segar', 'Makanan Segar', true, null, 37, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(38, 6, 'makanan-beku', 'Makanan Beku', true, null, 38, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(39, 6, 'es-krim', 'Es Krim', true, null, 39, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(40, 6, 'makanan-kemasan', 'Makanan Kemasan', true, null, 40, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(41, 7, 'peralatan-rias-tata-rambut', 'Peralatan Rias & Tata Rambut', true, null, 41, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(42, 7, 'perawatan-tubuh', 'Perawatan Tubuh', true, null, 42, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(43, 7, 'perawatan-pria', 'Perawatan Pria', true, null, 43, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(44, 7, 'perawatan-rambut', 'Perawatan Rambut', true, null, 44, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(45, 7, 'perawatan-wajah', 'Perawatan Wajah', true, null, 45, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(46, 7, 'riasan-wajah-tubuh', 'Riasan Wajah & Tubuh', true, null, 46, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(47, 7, 'pembalut-popok-dewasa', 'Pembalut & Popok Dewasa', true, null, 47, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(48, 7, 'perawatan-gigi-mulut', 'Perawatan Gigi & Mulut', true, null, 48, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(49, 7, 'parfum-cologne', 'Parfum & Cologne', true, null, 49, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(50, 8, 'obat-obatan', 'Obat-Obatan', true, null, 50, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(51, 8, 'vitamin-suplemen', 'Vitamin & Suplemen', true, null, 51, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(52, 8, 'alat-kesehatan', 'Alat Kesehatan', true, null, 52, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(53, 8, 'produk-higienis', 'Produk Higienis', true, null, 53, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(54, 9, 'mainan-hiburan', 'Mainan & Hiburan', true, null, 54, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(55, 9, 'fashion', 'Fashion', true, null, 55, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(56, 9, 'perlengkapan-kantor-sekolah', 'Perlengkapan Kantor & Sekolah', true, null, 56, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(57, 9, 'perlengkapan-otomotif', 'Perlengkapan Otomotif', true, null, 57, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(58, 9, 'rokok-korek', 'Rokok & Korek', true, null, 58, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(59, 10, 'makanan-anjing', 'Makanan Anjing', true, null, 59, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(60, 10, 'makanan-kucing', 'Makanan Kucing', true, null, 60, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(61, 10, 'makanan-hewan-peliharan-lainnya', 'Makanan Hewan Peliharan Lainnya', true, null, 61, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
ON CONFLICT (id) DO UPDATE SET
   code = EXCLUDED.code,
   parent_id = EXCLUDED.parent_id,
   name = EXCLUDED.name,
   is_active = EXCLUDED.is_active,
   icon = EXCLUDED.icon,
   display_order = EXCLUDED.display_order,
   updated_at = EXCLUDED.updated_at;