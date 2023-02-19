CREATE TABLE "user_account" (
  "id" bigserial PRIMARY KEY,
  "nickname" varchar (100) NOT NULL,
  "email" varchar NOT NULL,
  "balance" bigint NOT NULL,
  "hashed_password" varchar NOT NULL,
  "resp" boolean NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "transfers" (
  "id" bigserial PRIMARY KEY,
  "from_account_id" bigint,
  "to_account_id" bigint,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "entries" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigint,
  "amount" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "specialization" (
  "id" bigserial PRIMARY KEY,
  "branch" bigint,
  "name" varchar NOT NULL,
  "description" varchar NOT NULL,
  "online" boolean DEFAULT false
);

CREATE TABLE "messages" (
  "id" bigserial PRIMARY KEY,
  "who_ask_id" bigint,
  "who_answer_id" bigint,
  "specialization" bigint,
  "message_text" varchar NOT NULL,
  "entries" bigint,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "Branches" (
  "id" bigserial PRIMARY KEY,
  "name" varchar (100) NOT NULL
);

CREATE TABLE "purchase" (
  "id" bigserial PRIMARY KEY,
  "from_account_id" bigint,
  "amount_fiat" float8,
  "amount_coins" bigint,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "withdraw" (
  "id" bigserial PRIMARY KEY,
  "from_account_id" bigint,
  "amount_fiat" float8,
  "amount_coins" bigint,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "settings" (
  "rate" float8
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

ALTER TABLE "specialization" ADD FOREIGN KEY ("branch") REFERENCES "Branches" ("id");

ALTER TABLE "messages" ADD FOREIGN KEY ("who_ask_id") REFERENCES "user_account" ("id");

ALTER TABLE "messages" ADD FOREIGN KEY ("who_answer_id") REFERENCES "user_account" ("id");

ALTER TABLE "messages" ADD FOREIGN KEY ("entries") REFERENCES "entries" ("id");

ALTER TABLE "messages" ADD FOREIGN KEY ("specialization") REFERENCES "specialization" ("id");

ALTER TABLE "purchase" ADD FOREIGN KEY ("from_account_id") REFERENCES "user_account" ("id");

ALTER TABLE "withdraw" ADD FOREIGN KEY ("from_account_id") REFERENCES "user_account" ("id");