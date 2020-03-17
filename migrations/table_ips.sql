CREATE TABLE "public"."ips" (
  "id" serial4 NOT NULL,
  "ip" varchar COLLATE "pg_catalog"."default" NOT NULL,
  "country_code" varchar COLLATE "pg_catalog"."default",
  "country_name" varchar COLLATE "pg_catalog"."default",
  "created_at" timestamp(6) DEFAULT now(),
  PRIMARY KEY ("id")
)
;

ALTER TABLE "public"."ips"
  OWNER TO "postgres";
