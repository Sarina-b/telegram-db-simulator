create table if not exists account_stickers
(
    used_at         TIMESTAMP DEFAULT NULL ,
    usage_count     bigint,

    account_id      bigint not null
    constraint fk_account_stickers_account
    references accounts on delete cascade,
    sticker_id bigint not null
    constraint fk_account_stickers_sticker
    references stickers on delete cascade,
    primary key (account_id, sticker_id)
    );

