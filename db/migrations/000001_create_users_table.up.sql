CREATE TABLE users(
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    full_name VARCHAR(55) NOT NULL,
    email VARCHAR(55) NOT NULL UNIQUE,
    password VARCHAR(20) NOT NULL
);
