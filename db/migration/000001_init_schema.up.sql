CREATE TABLE "user_account" (
  "id" bigserial PRIMARY KEY,
  "nickname" varchar (100) NOT NULL UNIQUE,
  "email" varchar NOT NULL,
  "balance" bigint NOT NULL,
  "hashed_password" varchar NOT NULL,
  "resp" boolean NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "transfers" (
  "id" bigserial PRIMARY KEY,
  "from_account_id" bigint NOT NULL,
  "to_account_id" bigint NOT NULL,
  "amount" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "entries" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigint NOT NULL,
  "amount" bigint NOT NULL,
  "messages" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "specialization" (
  "id" bigserial PRIMARY KEY,
  "branch" bigint NOT NULL,
  "user_id" bigint NOT NULL,
  "spec_name" varchar NOT NULL,
  "descr" varchar NOT NULL,
  "is_online" boolean DEFAULT false
);

CREATE TABLE "messages" (
  "id" bigserial PRIMARY KEY,
  "who_ask_id" bigint NOT NULL,
  "who_answer_id" bigint NOT NULL,
  "specialization" bigint NOT NULL,
  "message_text" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "sessions" (
  "id" uuid PRIMARY KEY,
  "nickname" varchar NOT NULL,
  "refresh_token" varchar NOT NULL,
  "user_agent" varchar NOT NULL,
  "client_ip" varchar NOT NULL,
  "is_blocked" boolean NOT NULL DEFAULT false,
  "expires_at" timestamptz NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "branches" (
  "id" bigserial PRIMARY KEY,
  "branch_name" varchar (100) NOT NULL
);

CREATE TYPE status AS ENUM ('init', 'processed', 'completed');

CREATE TABLE "purchase" (
  "id" bigserial PRIMARY KEY,
  "from_account_id" bigint NOT NULL,
  "amount_fiat" float8 NOT NULL,
  "amount_coins" bigint NOT NULL,
  "status_p" status NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "withdraw" (
  "id" bigserial PRIMARY KEY,
  "from_account_id" bigint NOT NULL,
  "amount_fiat" float8 NOT NULL,
  "amount_coins" bigint NOT NULL,
  "status_w" status NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "settings" (
  "id" bigserial PRIMARY KEY,
  "rate" float8 NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "user_account" ("nickname");

CREATE INDEX ON "transfers" ("from_account_id");

CREATE INDEX ON "transfers" ("to_account_id");

CREATE INDEX ON "transfers" ("from_account_id", "to_account_id");

CREATE INDEX ON "entries" ("user_id");

CREATE INDEX ON "messages" ("who_answer_id");

CREATE INDEX ON "messages" ("who_ask_id");

COMMENT ON COLUMN "entries"."amount" IS 'can be negative or positive';

ALTER TABLE "transfers" ADD FOREIGN KEY ("from_account_id") REFERENCES "user_account" ("id");

ALTER TABLE "transfers" ADD FOREIGN KEY ("to_account_id") REFERENCES "user_account" ("id");

ALTER TABLE "entries" ADD FOREIGN KEY ("user_id") REFERENCES "user_account" ("id");

ALTER TABLE "specialization" ADD FOREIGN KEY ("branch") REFERENCES "branches" ("id");

ALTER TABLE "specialization" ADD FOREIGN KEY ("user_id") REFERENCES "user_account" ("id");

ALTER TABLE "messages" ADD FOREIGN KEY ("who_ask_id") REFERENCES "user_account" ("id");

ALTER TABLE "messages" ADD FOREIGN KEY ("who_answer_id") REFERENCES "user_account" ("id");

ALTER TABLE "entries" ADD FOREIGN KEY ("messages") REFERENCES "messages" ("id");

ALTER TABLE "sessions" ADD FOREIGN KEY ("nickname") REFERENCES "user_account" ("nickname");

ALTER TABLE "messages" ADD FOREIGN KEY ("specialization") REFERENCES "specialization" ("id");

ALTER TABLE "purchase" ADD FOREIGN KEY ("from_account_id") REFERENCES "user_account" ("id");

ALTER TABLE "withdraw" ADD FOREIGN KEY ("from_account_id") REFERENCES "user_account" ("id");