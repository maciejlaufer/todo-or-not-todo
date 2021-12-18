CREATE TABLE "users" (
  "id" uuid PRIMARY KEY NOT NULL DEFAULT (gen_random_uuid()),
  "email" varchar UNIQUE NOT NULL,
  "password" varchar NOT NULL,
  "first_name" varchar,
  "last_name" varchar,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "lists" (
  "id" uuid PRIMARY KEY NOT NULL DEFAULT (gen_random_uuid()),
  "name" varchar NOT NULL,
  "creator_id" uuid NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "users_lists" (
  "user_id" uuid NOT NULL,
  "list_id" uuid NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  PRIMARY KEY ("user_id", "list_id")
);

CREATE TABLE "tasks" (
  "id" uuid PRIMARY KEY NOT NULL DEFAULT (gen_random_uuid()),
  "name" varchar NOT NULL,
  "description" varchar,
  "list_id" uuid NOT NULL,
  "completed" boolean NOT NULL DEFAULT false,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "lists" ADD FOREIGN KEY ("creator_id") REFERENCES "users" ("id");

ALTER TABLE "users_lists" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "users_lists" ADD FOREIGN KEY ("list_id") REFERENCES "lists" ("id");

ALTER TABLE "tasks" ADD FOREIGN KEY ("list_id") REFERENCES "lists" ("id");

CREATE INDEX ON "users" ("email");

CREATE INDEX ON "users" ("first_name");

CREATE INDEX ON "users" ("last_name");

CREATE INDEX ON "users" ("first_name", "last_name");

CREATE INDEX ON "lists" ("name");

CREATE INDEX ON "lists" ("creator_id");

CREATE INDEX ON "users_lists" ("user_id", "list_id");

CREATE INDEX ON "users_lists" ("list_id", "user_id");

CREATE INDEX ON "tasks" ("name");

CREATE INDEX ON "tasks" ("list_id");

CREATE INDEX ON "tasks" ("completed");
