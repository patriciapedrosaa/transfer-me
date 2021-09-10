CREATE TABLE IF NOT EXISTS "tokens" (
    "id" uuid PRIMARY KEY,
    "name" varchar NOT NULL,
    "subject" varchar NOT NULL,
    "issuer" varchar NOT NULL,
    "issued_at" timestamptz NOT NULL,
    "expired_at" timestamptz NOT NULL
);