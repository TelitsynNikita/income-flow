CREATE TABLE outflow (
    id SERIAL PRIMARY KEY ,
    goods_id INTEGER NOT NULL ,
    goods_count SMALLINT NOT NULL ,
    FOREIGN KEY (goods_id) REFERENCES goods (id)
);