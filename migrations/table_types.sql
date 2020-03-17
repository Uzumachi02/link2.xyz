CREATE TABLE "public"."types" (
  "id" serial4 NOT NULL,
  "name" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "status" int2 DEFAULT 1,
  "created_at" timestamp(6) DEFAULT now(),
  PRIMARY KEY ("id")
)
;

ALTER TABLE "public"."types"
  OWNER TO "postgres";
