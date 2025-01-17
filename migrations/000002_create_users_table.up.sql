BEGIN;

CREATE TABLE "sso"."users" (
  "id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  "created_at" TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT now(),
  "updated_at" TIMESTAMP WITHOUT TIME ZONE,
  "deleted_at" TIMESTAMP WITHOUT TIME ZONE,
  "is_deleted" BOOLEAN NOT NULL DEFAULT false,
  "password" CHARACTER VARYING NOT NULL,
  "login" CHARACTER VARYING NOT NUll,
  CONSTRAINT "UK_user_login" UNIQUE("login")
);

CREATE TYPE "sso"."user_contact_type" AS ENUM ('EMAIL', 'PHONE_NUMBER');

CREATE TABLE "sso"."user_contacts" (
  "id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  "created_at" TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT now(),
  "updated_at" TIMESTAMP WITHOUT TIME ZONE,
  "deleted_at" TIMESTAMP WITHOUT TIME ZONE,
  "is_deleted" BOOLEAN NOT NULL DEFAULT false,
  "_value" CHARACTER VARYING NOT NULL,
  "_type" "sso"."user_contact_type" NOT NULL,
  "user_id" UUID NOT NULL REFERENCES "sso"."users"("id")
);

COMMIT;
