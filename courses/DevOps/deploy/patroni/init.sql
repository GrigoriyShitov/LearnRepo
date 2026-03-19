CREATE TABLE orders (
  id bigint PRIMARY KEY,
  status varchar(50) NOT NULL DEFAULT 'created',
  description text,
  created_at timestamp NOT NULL DEFAULT now()
);
