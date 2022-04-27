-- +goose Up
-- +goose StatementBegin
CREATE TABLE orders (
    id int NOT NULL AUTO_INCREMENT,
    cashier_id int NOT NULL,
    payment_id int NULL,
    total_price bigint NOT NULL,
    total_paid bigint NOT NULL,
    total_return bigint NOT NULL,
    receipt_id varchar(50) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE orders;
-- +goose StatementEnd
