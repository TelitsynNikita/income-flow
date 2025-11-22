CREATE TABLE remains (
    id SERIAL NOT NULL ,
    goods_id INTEGER NOT NULL ,
    section_id INTEGER NOT NULL ,
    goods_count SMALLINT NOT NULL ,
    FOREIGN KEY (goods_id) REFERENCES goods (id) ,
    FOREIGN KEY (section_id) REFERENCES sections (id)
);