CREATE TABLE "public"."agents" (
  "id" serial4 NOT NULL,
  "hash" varchar COLLATE "pg_catalog"."default" NOT NULL,
  "user_agent" varchar COLLATE "pg_catalog"."default" NOT NULL,
  "platform" varchar COLLATE "pg_catalog"."default",
  "browser" varchar COLLATE "pg_catalog"."default",
  "version" varchar COLLATE "pg_catalog"."default",
  "created_at" timestamp(6) DEFAULT now(),
  PRIMARY KEY ("id")
)
;

ALTER TABLE "public"."agents"
  OWNER TO "postgres";
