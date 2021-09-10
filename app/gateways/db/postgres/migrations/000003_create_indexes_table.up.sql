CREATE INDEX IF NOT EXISTS "accounts_id_idx" ON "accounts"("id") ;
CREATE INDEX IF NOT EXISTS "accounts_cpf_idx" ON "accounts" ("cpf");
CREATE INDEX IF NOT EXISTS "transfers_origin_account_id_idx" ON "transfers" ("origin_account_id");