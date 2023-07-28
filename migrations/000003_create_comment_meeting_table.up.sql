CREATE TABLE "meeting".comments (
    "id" uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    "detail" TEXT NOT NULL
)

CREATE TABLE "meeting".comment_meetings (
    "id" uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    "meeting_id" uuid NOT NULL,
    "comment_id" uuid NOT NULL,
    "status" "public"."status" NOT NULL DEFAULT 'active',
    "created_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "created_by" uuid ,
    "updated_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (meeting_id) REFERENCES "meeting".meetings (id),
    FOREIGN KEY (comment_id) REFERENCES "meeting".comments (id),
    FOREIGN KEY (created_by) REFERENCES "user"."users" (user_id)
)