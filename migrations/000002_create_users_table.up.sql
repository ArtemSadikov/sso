BEGIN;

CREATE TABLE "sso"."users" (
  "id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  "created_at" TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT now(),
  "updated_at" TIMESTAMP WITHOUT TIME ZONE,
  "deleted_at" TIMESTAMP WITHOUT TIME ZONE,

  "password" CHARACTER VARYING NOT NULL
);


CREATE TYPE "sso"."user_contact_type" AS ENUM ('LOGIN');
CREATE TABLE "sso"."user_contacts" (
  "id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  "created_at" TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT now(),
  "updated_at" TIMESTAMP WITHOUT TIME ZONE,
  "deleted_at" TIMESTAMP WITHOUT TIME ZONE,

  "_value" CHARACTER VARYING NOT NULL,
  "_type" "sso"."user_contact_type" NOT NULL,

  "user_id" UUID NOT NULL REFERENCES "sso"."users"("id")
);

COMMIT;
