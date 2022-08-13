-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS public.reminders (
    id serial PRIMARY KEY,
    date date NOT NULL,
    text text NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS public.reminders;
-- +goose StatementEnd
