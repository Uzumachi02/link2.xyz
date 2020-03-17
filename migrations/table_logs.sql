CREATE TABLE "public"."logs" (
  "id" serial4 NOT NULL,
  "link_id" int4 NOT NULL,
  "ip_id" int4 NOT NULL,
  "agent_id" int4 NOT NULL,
  "others" text COLLATE "pg_catalog"."default",
  "created_at" timestamp(6) DEFAULT now(),
  PRIMARY KEY ("id")
)
;

ALTER TABLE "public"."logs"
  OWNER TO "postgres";
