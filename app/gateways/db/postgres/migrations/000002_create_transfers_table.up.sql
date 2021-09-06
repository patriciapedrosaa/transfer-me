CREATE TABLE IF NOT EXISTS "transfers" (
    "id" uuid PRIMARY KEY,
    "origin_account_id" uuid NOT NULL REFERENCES "accounts"("id"),
    "destination_account_id" uuid NOT NULL REFERENCES "accounts"("id"),
    "amount" bigint NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now())
);
