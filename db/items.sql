CREATE TABLE auth-skm.items (
    id bigint NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL,
    deleted_at timestamp with time zone,
    sku character varying(255) NOT NULL,
    name character varying(255) DEFAULT ''::character varying NOT NULL,
    price float DEFAULT 0 NOT NULL,
    inventory_qty integer DEFAULT 0 NOT NULL
);

CREATE SEQUENCE auth-skm.items_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE auth-skm.items_id_seq OWNED BY auth-skm.items.id;

ALTER TABLE ONLY auth-skm.items ALTER COLUMN id SET DEFAULT nextval('auth-skm.items_id_seq'::regclass);

ALTER TABLE ONLY auth-skm.items
    ADD CONSTRAINT items_pkey PRIMARY KEY (id);

ALTER TABLE ONLY auth-skm.items
    ADD CONSTRAINT items_sku_key UNIQUE (sku);

CREATE INDEX itemsindex ON auth-skm.items USING btree (deleted_at, sku, name);


INSERT INTO auth-skm.items (sku, name, price, inventory_qty)
VALUES ('120P90', 'Google Home', 49.99, 10);
INSERT INTO auth-skm.items (sku, name, price, inventory_qty)
VALUES ('43N23P', 'Macbook Pro', 5399.99, 5);
INSERT INTO auth-skm.items (sku, name, price, inventory_qty)
VALUES ('A304SD', 'Alexa Speaker', 109.50, 10);
INSERT INTO auth-skm.items (sku, name, price, inventory_qty)
VALUES ('234234', 'Raspberry Pi B', 30, 2);