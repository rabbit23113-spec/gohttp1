CREATE TABLE users (
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    username VARCHAR(255) NOT NULL UNIQUE,
    hash_password VARCHAR(255) NOT NULL
);