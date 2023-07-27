CREATE SCHEMA meeting;

CREATE TYPE "public"."status" AS ENUM ('to_do', 'in_progress', 'done', 'canceled');


CREATE TABLE "meeting".meetings (
    "id" uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    "title" TEXT NOT NULL,
    "detail" TEXT NOT NULL,
    "status" "public"."status"  NOT NULL DEFAULT 'to_do', 
    "created_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "created_by" uuid ,
    "updated_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (created_by) REFERENCES "user"."users" (user_id)
)