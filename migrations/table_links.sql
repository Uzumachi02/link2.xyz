CREATE TABLE "public"."links" (
  "id" serial4 NOT NULL,
  "type_id" int4 NOT NULL,
  "url" varchar COLLATE "pg_catalog"."default" NOT NULL,
  "target_url" varchar COLLATE "pg_catalog"."default" NOT NULL,
  "viewed" int4 DEFAULT 0,
  "created_at" timestamp(6) DEFAULT now(),
  "updated_at" timestamp(6) DEFAULT now(),
  PRIMARY KEY ("id")
)
;

ALTER TABLE "public"."links"
  OWNER TO "postgres";
