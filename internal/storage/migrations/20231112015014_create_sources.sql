-- +goose Up
-- +goose StatementBegin
    create table sources(
        id SERIAL PRIMARY KEY,
        source_name VARCHAR(255) NOT NULL,
        feed_url VARCHAR(255) NOT NULL,
        created_at TIMESTAMP NOT NULL DEFAULT NOW(),
        updated_at TIMESTAMP NOT NULL DEFAULT NOW()
    );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
    drop table if exists sources;
-- +goose StatementEnd
