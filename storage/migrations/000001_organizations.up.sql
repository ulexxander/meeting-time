
CREATE TABLE organizations (
  id bigint PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  name text NOT NULL,
  createdAt timestamptz NOT NULL DEFAULT NOW(),
  updatedAt timestamptz
);
