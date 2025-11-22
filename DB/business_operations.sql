CREATE TABLE business_operations (
    id SERIAL PRIMARY KEY ,
    date TIMESTAMP DEFAULT CURRENT_TIMESTAMP ,
    operation_type VARCHAR NOT NULL ,
    operation_id INTEGER NOT NULL ,
    goods_count SMALLINT NOT NULL
)