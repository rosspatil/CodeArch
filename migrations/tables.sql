
CREATE EXTENSION "uuid-ossp";

CREATE TABLE users (
  id UUID DEFAULT uuid_generate_v4(),
  name TEXT ,
  email TEXT primary key,
  password TEXT,
  created_at INTEGER DEFAULT extract(epoch from now() AT time zone 'utc')
)