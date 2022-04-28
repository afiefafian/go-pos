-- +goose Up
-- +goose StatementBegin
CREATE TABLE order_products (
    id int NOT NULL AUTO_INCREMENT,
    order_id int NOT NULL,
    product_id int NOT NULL,
    discount_id bigint NOT NULL,
    qty int NOT NULL,
    price bigint NOT NULL,
    total_normal_price bigint NOT NULL,
    total_final_price bigint NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE order_products;
-- +goose StatementEnd
