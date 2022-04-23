
CREATE TABLE meetings (
  id bigint PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  schedule_id bigint NOT NULL
    REFERENCES schedules(id) ON DELETE CASCADE,
  started_at timestamptz NOT NULL,
  ended_at timestamptz NOT NULL,
  created_at timestamptz NOT NULL DEFAULT NOW(),
  updated_at timestamptz
);
