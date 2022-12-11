CREATE TABLE products (
    id          BIGSERIAL NOT NULL,
    category    VARCHAR (64) NOT NULL,
    name        VARCHAR (255) NOT NULL,
    description TEXT DEFAULT NULL,

    CONSTRAINT pk_products__id PRIMARY KEY (id)
);
