CREATE TABLE IF NOT EXISTS delivery (
	id BIGSERIAL PRIMARY KEY,
	name VARCHAR(128),
	phone VARCHAR(128),
	zip VARCHAR(128),
	city VARCHAR(128),
	address VARCHAR(128),
	region VARCHAR(128),
	email VARCHAR(128)
);

CREATE TABLE IF NOT EXISTS payment (
	id BIGSERIAL PRIMARY KEY,
	trasaction VARCHAR(128),
	request_id INTEGER,
	currency VARCHAR(128),
	provider VARCHAR(128),
	amount INTEGER,
	payment_dt INTEGER,
	bank VARCHAR(128),
	delivery_cost INTEGER,
	goods_total INTEGER,
	custom_fee INTEGER
);

CREATE TABLE IF NOT EXISTS item (
	id BIGSERIAL PRIMARY KEY,
	chrt_id INTEGER,
    track_number VARCHAR(128),
    price INTEGER,
    rid VARCHAR(128),
    name VARCHAR(128),
    sale INTEGER,
    size VARCHAR(128),
    total_price INTEGER,
    nm_id INTEGER,
    brand VARCHAR(128),
    status INTEGER
);

CREATE TABLE IF NOT EXISTS orders (
	id BIGSERIAL PRIMARY KEY,
	order_uid VARCHAR(128),
	track_number VARCHAR (128),
	entry VARCHAR (128),
	delivery_id INTEGER REFERENCES delivery(id),
	payment_id INTEGER REFERENCES payment (id),
	item_id INTEGER REFERENCES item(id),
	locale VARCHAR (128),
	internal_signature VARCHAR (128),
	customer_id VARCHAR (128),
	delivery_service VARCHAR (128),
	shardkey INTEGER,
	sm_id INTEGER,
	date_created TIMESTAMP,
	oof_shard VARCHAR (128)
)