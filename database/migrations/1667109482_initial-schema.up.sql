CREATE EXTENSION IF NOT EXISTS pgcrypto;

DROP TABLE IF EXISTS example_entity;

CREATE TABLE IF NOT EXISTS example_entity (
  id          uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  title       varchar(255) NOT NULL,
  description text,
  deleted     boolean NOT NULL DEFAULT false,
  enabled     boolean NOT NULL DEFAULT true,
  status      integer,

  created_on  timestamptz NOT NULL DEFAULT (now() at time zone 'utc'),
  created_by  integer NOT NULL,
  modified_on timestamptz,
  modified_by integer
);

CREATE INDEX example_entity_status_idx ON example_entity (status);
