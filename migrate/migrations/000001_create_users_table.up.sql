CREATE TABLE IF NOT EXISTS users(
   id serial PRIMARY KEY,
   created_at timestamp with time zone NOT NULL default now(),
   removed_at timestamp with time zone NULL,
   updated_at timestamp with time zone NOT NULL default now(),
   email VARCHAR (300) UNIQUE NOT NULL,
   password VARCHAR (50) NOT NULL,
   last_name VARCHAR (50) NOT NULL,
   first_name VARCHAR (50) NOT NULL
);
