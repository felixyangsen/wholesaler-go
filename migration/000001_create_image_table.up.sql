BEGIN;

CREATE TABLE IF NOT EXISTS "image" (
    id SERIAL PRIMARY KEY,
    path VARCHAR(255) NOT NULL,
    link VARCHAR(255) NOT NULL
);

COMMIT;