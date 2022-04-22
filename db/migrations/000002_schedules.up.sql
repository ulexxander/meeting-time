
CREATE TABLE schedules (
  id bigint PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  team_id bigint NOT NULL
    REFERENCES teams(id) ON DELETE CASCADE,
  name text NOT NULL,
  starts_at timestamptz NOT NULL,
  ends_at timestamptz NOT NULL,
  created_at timestamptz NOT NULL DEFAULT NOW(),
  updated_at timestamptz
);
