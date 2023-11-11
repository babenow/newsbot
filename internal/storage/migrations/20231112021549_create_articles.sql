-- +goose Up
-- +goose StatementBegin
    create table articles(
        id SERIAL PRIMARY KEY,
        source_id INT NOT NULL,
        title VARCHAR(255) NOT NULL,
        link VARCHAR(255) NOT NULL,
        summary TEXT NOT NULL,
        published_at TIMESTAMP NOT NULL,
        created_at TIMESTAMP NOT NULL,
        posted_at TIMESTAMP
    );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
    drop table if exists articles;
-- +goose StatementEnd