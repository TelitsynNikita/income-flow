CREATE TABLE outflow (
    id SERIAL PRIMARY KEY ,
    goods_id INTEGER NOT NULL ,
    goods_count SMALLINT NOT NULL ,
    contractors_id INTEGER NOT NULL ,
    FOREIGN KEY (goods_id) REFERENCES goods (id) ,
    FOREIGN KEY (contractors_id) REFERENCES contractors (id)
);