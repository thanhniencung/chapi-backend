CREATE TABLE public.product
(
    product_id text COLLATE pg_catalog."default" NOT NULL,
    product_name text COLLATE pg_catalog."default" NOT NULL,
    quatity integer NOT NULL,
    sold_items integer NOT NULL,
    price numeric NOT NULL,
    cate_id text COLLATE pg_catalog."default" NOT NULL,
    product_image text COLLATE pg_catalog."default",
    delete_at timestamp with time zone,
    create_at timestamp with time zone,
    update_at timestamp with time zone,
    CONSTRAINT product_pkey PRIMARY KEY (product_id)
)