CREATE EXTENSION IF NOT EXISTS pgcrypto;

-- Resource Entity
DROP TABLE IF EXISTS resource_entity;

CREATE TABLE IF NOT EXISTS resource_entity (
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

CREATE INDEX resource_entity_status_idx ON resource_entity (status);
