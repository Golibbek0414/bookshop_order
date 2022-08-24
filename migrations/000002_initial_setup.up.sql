ALTER TABLE orders ADD COLUMN created_at timestamp default current_timestamp;
ALTER TABLE orders ADD COLUMN updated_at timestamp;
ALTER TABLE orders ADD COLUMN deleted_at timestamp;