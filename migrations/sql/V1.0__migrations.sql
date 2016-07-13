-----------
-- Users --
-----------

CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE users (
  id uuid NOT NULL PRIMARY KEY DEFAULT gen_random_uuid(),
  email varchar NOT NULL CHECK (length(email) > 0 AND length(email) < 255),
  password_hash varchar NOT NULL CHECK (length(password_hash) > 0 AND length(password_hash) < 255),
  created_at timestamp without time zone NOT NULL,
  updated_at timestamp without time zone NOT NULL,
  deleted_at timestamp without time zone
);

CREATE INDEX user_pk_idx ON users USING btree (id);
CREATE INDEX user_deleted_at_idx ON users USING btree (deleted_at);

----------------------
