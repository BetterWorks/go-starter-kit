CREATE EXTENSION IF NOT EXISTS pgcrypto;

DROP TABLE IF EXISTS "entity";

CREATE TABLE IF NOT EXISTS "entity" (
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  org_id  integer NOT NULL,
  name    varchar(64) NOT NULL,
  enabled boolean NOT NULL DEFAULT true,
  deleted boolean NOT NULL DEFAULT false,
  status  integer,

  created_on  timestamptz NOT NULL DEFAULT (now() at time zone 'utc'),
  created_by  integer NOT NULL,
  modified_on timestamptz,
  modified_by integer

  CREATE INDEX entity_org_id_idx ON entity(org_id);
);
