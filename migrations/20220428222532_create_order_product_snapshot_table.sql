-- +goose Up
-- +goose StatementBegin
CREATE TABLE order_product_snapshots (
    id int NOT NULL AUTO_INCREMENT,
    order_id int NOT NULL,
    product_id int NOT NULL,
    category_id int NULL,
    discount_id int NULL,
    name varchar(100) NOT NULL,
    stock int NOT NULL,
    price bigint NOT NULL,
    image varchar(255) NOT NULL,
    category text NULL,
    discount text NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE order_product_snapshots;
-- +goose StatementEnd
