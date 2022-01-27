CREATE TABLE IF NOT EXISTS "user" (
  "full_name" varchar(50) NOT NULL,
  "email" varchar(50) PRIMARY KEY,
  "passwd" varchar(50) NOT NULL,
  "mobile_no" varchar(10) NOT NULL constraint Check_Digits check ("mobile_no" NOT LIKE '%[^0-9]%'),
  "source" varchar(50) NOT NULL,
  "agreement" boolean NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE INDEX ON "user" ("email");