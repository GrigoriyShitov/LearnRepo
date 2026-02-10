-- +goose Up
-- +goose StatementBegin
ALTER TABLE goose
    OWNER TO "awesome-backend";
SET ROLE "awesome-backend";
-- +goose StatementEnd
