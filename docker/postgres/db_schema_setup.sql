CREATE TABLE "public"."tale" (
  "tale_id" uuid NOT NULL DEFAULT NULL,
  "parent" uuid DEFAULT NULL,
  "content" varchar(255) COLLATE "pg_catalog"."default" NOT NULL DEFAULT NULL,
  CONSTRAINT "tale_pkey" PRIMARY KEY ("tale_id"),
  CONSTRAINT "tale_parent_fkey" FOREIGN KEY ("parent") REFERENCES "public"."tale" ("tale_id") ON DELETE NO ACTION ON UPDATE NO ACTION,
  CONSTRAINT "tale_tale_id_parent_key" UNIQUE ("tale_id", "parent")
);
