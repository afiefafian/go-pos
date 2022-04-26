-- +goose Up
-- +goose StatementBegin
CREATE TABLE payments (
    id int NOT NULL AUTO_INCREMENT,
    name varchar(100) NOT NULL,
    type varchar(20) NOT NULL,
    logo varchar(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE payments;
-- +goose StatementEnd
