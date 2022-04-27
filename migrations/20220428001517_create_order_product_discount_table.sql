-- +goose Up
-- +goose StatementBegin
CREATE TABLE order_product_discounts (
    id int NOT NULL AUTO_INCREMENT,
    type varchar(20) NOT NULL,
    qty int NOT NULL,
    result bigint NOT NULL,
    expired_at bigint NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE order_product_discounts;
-- +goose StatementEnd
