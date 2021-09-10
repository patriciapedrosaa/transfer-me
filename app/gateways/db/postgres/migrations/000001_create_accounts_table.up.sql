CREATE TABLE IF NOT EXISTS "accounts" (
  "id" uuid PRIMARY KEY,
  "name" varchar NOT NULL,
  "cpf" varchar NOT NULL,
  "secret" varchar NOT NULL,
  "balance" bigint NOT NULL,
  "created_at" timestamptz NOT NULL
);