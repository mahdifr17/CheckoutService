CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE public.products (
    id uuid DEFAULT uuid_generate_v4() NOT NULL, 
    created_at timestamptz NULL,
    updated_at timestamptz NULL,
    deleted_at timestamptz NULL,
    CONSTRAINT products_pkey PRIMARY KEY (id)
);
CREATE INDEX idx_products_deleted_at ON public.products USING btree (deleted_at);

