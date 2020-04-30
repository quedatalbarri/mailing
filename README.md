# Queda't al Barri

## abeja

Go server. 
Gets data from calendar and sends to mailing list, via mailchimp template

## apuntat

Is the server that recieves all requests to quedat.barcelona. 

Uses config to redirect users from /:barrio to the url for the mailing sign up list for that barrio

## barris-client

React app.
Administration panel for creating a barri.

## barris-server

Go server
"barris-client's" backend.


## postgresql

-- Table: public.barris

CREATE TABLE public.barris
(
    name character varying(50) COLLATE pg_catalog."default" NOT NULL,
    url character varying COLLATE pg_catalog."default",
    admin character varying COLLATE pg_catalog."default",
    telegram_token character varying COLLATE pg_catalog."default",
    CONSTRAINT barris_pkey PRIMARY KEY (name)
)