create table billings
(
    id          bigint auto_increment
        primary key,
    loan_id     bigint      null,
    user_id     bigint      null,
    created_at  datetime(3) null,
    amount      double      null,
    status      longtext    null,
    approved_at datetime(3) null
);

create table loans
(
    id                 bigint auto_increment
        primary key,
    user_id            bigint      null,
    loan_no            longtext    null,
    otr                double      null,
    admin_fee          double      null,
    installment_amount double      null,
    interest_amount    double      null,
    asset_name         longtext    null,
    status             longtext    null,
    tenor              bigint      null,
    created_at         datetime(3) null,
    approved_at        datetime(3) null
);

create table master_tenor
(
    id       bigint auto_increment
        primary key,
    user_id  bigint null,
    tenor    bigint null,
    amount   double null,
    interest double null
);

create table users
(
    id                   bigint auto_increment
        primary key,
    nik                  longtext    null,
    full_name            longtext    null,
    legal_name           longtext    null,
    birthplace           longtext    null,
    birthdate            datetime(3) null,
    salary               double      null,
    identification_photo longtext    null,
    photo_selfie         longtext    null,
    status               longtext    null,
    created_at           datetime(3) null
);

