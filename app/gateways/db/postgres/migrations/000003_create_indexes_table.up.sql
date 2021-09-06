CREATE INDEX IF NOT EXISTS "id" ON "accounts"("id") ;
CREATE INDEX IF NOT EXISTS "cpf" ON "accounts" ("cpf");
CREATE INDEX IF NOT EXISTS "origin_account_id" ON "transfers" ("origin_account_id");