CREATE SCHEMA "user";
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE "user"."users" (
    "user_id" uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    "email" VARCHAR(255) NOT NULL,
    "password" varchar(200) NOT NULL, 
    "status" VARCHAR(255) NOT NULL DEFAULT 'active',
    "created_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" timestamptz
);