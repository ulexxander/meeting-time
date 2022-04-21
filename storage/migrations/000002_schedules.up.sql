
CREATE TABLE schedules (
  id bigint PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  teamID bigint
    REFERENCES teams(id) ON DELETE CASCADE,
  name text NOT NULL,
  startsAt timestamptz NOT NULL,
  endsAt timestamptz NOT NULL,
  createdAt timestamptz NOT NULL DEFAULT NOW(),
  updatedAt timestamptz
);
