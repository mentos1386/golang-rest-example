CREATE TABLE IF NOT EXISTS groups(
  id serial PRIMARY KEY,
  name VARCHAR (300) UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS users(
  id serial PRIMARY KEY,
  email VARCHAR (300) UNIQUE NOT NULL,
  name VARCHAR (100) UNIQUE NOT NULL,

  group_id INT,
  CONSTRAINT fk_users_groups_group_id
    FOREIGN KEY(group_id)
      REFERENCES groups(id)
);
