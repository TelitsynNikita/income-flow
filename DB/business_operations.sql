CREATE TABLE business_operations (
    id SERIAL PRIMARY KEY ,
    date TIMESTAMP DEFAULT CURRENT_TIMESTAMP ,
    operation_type VARCHAR NOT NULL ,
    operation_id INTEGER NOT NULL ,
    FOREIGN KEY (operation_id) REFERENCES income (id) ,
    FOREIGN KEY (operation_id) REFERENCES outflow (id)
)