-- migrate:up
CREATE TABLE IF NOT EXISTS books (
  id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
  title varchar(255) DEFAULT '',
  author varchar(255) DEFAULT '',
  isbn varchar(50) DEFAULT '',
  created_time timestamp NOT NULL DEFAULT (now()),
  updated_time timestamp NOT NULL DEFAULT (now())
);


-- migrate:down
DROP TABLE IF EXISTS books;