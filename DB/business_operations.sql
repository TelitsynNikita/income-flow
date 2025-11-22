CREATE TABLE business_operations (
    id SERIAL PRIMARY KEY ,
    date TIMESTAMP DEFAULT CURRENT_TIMESTAMP ,
    operation_type VARCHAR NOT NULL ,
    contractors_id INTEGER NOT NULL ,
    FOREIGN KEY (contractors_id) REFERENCES contractors (id)
)