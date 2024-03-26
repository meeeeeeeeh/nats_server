DROP TABLE orders CASCADE;
DROP TABLE wb_l0 CASCADE;
DROP TABLE delivery CASCADE;
DROP TABLE item CASCADE;
DROP TABLE payment CASCADE



SELECT COUNT(*) FROM order_info

INSERT INTO delivery (name, phone, zip, city, address, region, email) 
VALUES ()

DROP DATABASE IF EXISTS order_info

SELECT * FROM delivery

TRUNCATE delivery

INSERT INTO order_info (order_uid, track_number, locale, internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard) VALUES ("b563feb7b2b84b6test", "WBILMTESTTRACK", "WBIL", "en", "", "test", "meest", "9", 99,  "2021-11-26T06:22:19Z", "1")



SELECT * FROM delivery