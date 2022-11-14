CREATE EXTENSION IF NOT EXISTS pgcrypto;

DROP TABLE IF EXISTS "episode";

CREATE TABLE IF NOT EXISTS "episode" (
  id uuid     PRIMARY KEY DEFAULT gen_random_uuid(),
  season_id   uuid NOT NULL,
  title       varchar(255) NOT NULL,
  description text,
  year        integer,
  director    varchar(255),
  deleted     boolean NOT NULL DEFAULT false,
  enabled     boolean NOT NULL DEFAULT true,
  status      integer,

  created_on  timestamptz NOT NULL DEFAULT (now() at time zone 'utc'),
  created_by  integer NOT NULL,
  modified_on timestamptz,
  modified_by integer

  CREATE INDEX episode_season_id_idx ON episode(season_id);
  CREATE INDEX episode_status_idx ON episode(status);
);

DROP TABLE IF EXISTS "season";

CREATE TABLE IF NOT EXISTS "season" (
  id uuid     PRIMARY KEY DEFAULT gen_random_uuid(),
  title       varchar(255) NOT NULL,
  description text,
  deleted     boolean NOT NULL DEFAULT false,
  enabled     boolean NOT NULL DEFAULT true,
  status      integer,

  created_on  timestamptz NOT NULL DEFAULT (now() at time zone 'utc'),
  created_by  integer NOT NULL,
  modified_on timestamptz,
  modified_by integer

  CREATE INDEX season_status_idx ON season(status);
);
