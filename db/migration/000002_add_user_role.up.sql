ALTER TABLE "users" ADD "role" varchar NOT NULL DEFAULT 'USER';

CREATE INDEX ON "users" ("role");