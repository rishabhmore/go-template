-- +migrate Up
CREATE TABLE public.notes (
    id SERIAL UNIQUE PRIMARY KEY,
    user_id int NOT NULL REFERENCES users(id),
    first_name TEXT,
    last_name TEXT,
    title TEXT,
    note TEXT,
    created_at TIMESTAMP WITH TIME ZONE,
	updated_at TIMESTAMP WITH TIME ZONE,
	deleted_at TIMESTAMP WITH TIME ZONE
);
CREATE INDEX notes_user_id_idx ON notes(user_id);

-- +migrate Down
DROP TABLE notes;