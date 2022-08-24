create table orders(
    order_id uuid not null primary key,
    book_id uuid not null,
    description text
);