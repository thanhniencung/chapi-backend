-- +migrate Up
CREATE TABLE "cate"
(
    "cate_id" text NOT NULL PRIMARY KEY,
    "cate_name" text NOT NULL,
    "created_at" timestamp with time zone,
    "deleted_at" time with time zone,
    "updated_at" time with time zone
)

CREATE TABLE "product"
(
    "product_id" text NOT NULL PRIMARY KEY,
    "product_name" text NOT NULL,
    "quantity" integer NOT NULL,
    "sold_items" integer NOT NULL,
    "price" numeric NOT NULL,
    "cate_id" text NOT NULL,
    "product_image" text NOT NULL,
    "deleted_at" timestamp with time zone,
    "created_at" timestamp with time zone,
    "updated_at" timestamp with time zone,
    "user_id" text NOT NULL
)

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE cate;
DROP TABLE product;



