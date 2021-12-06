CREATE TABLE public."character"
(
    id integer NOT NULL,
    created timestamp without time zone,
    updated time without time zone,
    name character varying(255) COLLATE pg_catalog."default" NOT NULL,
    description character varying COLLATE pg_catalog."default",
    rarity integer,
    birthday timestamp without time zone,
    unique_key integer,
    deleted_at timestamp without time zone,
    CONSTRAINT character_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE public."character"
    OWNER to postgres;