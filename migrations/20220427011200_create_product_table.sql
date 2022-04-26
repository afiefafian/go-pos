-- +goose Up
-- +goose StatementBegin
CREATE TABLE products (
    id int NOT NULL AUTO_INCREMENT,
    category_id int NOT NULL,
    discount_id int NULL,
    name varchar(100) NOT NULL,
    stock int NOT NULL,
    price bigint NOT NULL,
    image varchar(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE products;
-- +goose StatementEnd
