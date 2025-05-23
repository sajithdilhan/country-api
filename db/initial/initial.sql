CREATE SCHEMA country_infos AUTHORIZATION postgres;

-- country_infos.countries definition

-- Drop table

-- DROP TABLE country_infos.countries;

CREATE TABLE country_infos.countries (
	id int4 GENERATED BY DEFAULT AS IDENTITY NOT NULL,
	"name" varchar NOT NULL,
	code varchar NOT NULL,
	continent varchar NULL,
	official_language varchar NULL,
	capital varchar NULL,
	area_sq_km int8 NULL,
	population int8 NULL,
	currency varchar NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	CONSTRAINT country_pk PRIMARY KEY (id),
	CONSTRAINT country_unique_code UNIQUE (code),
	CONSTRAINT country_unique_name UNIQUE (name)
);

-- country_infos.economies definition

-- Drop table

-- DROP TABLE country_infos.economies;

CREATE TABLE country_infos.economies (
	id int4 GENERATED BY DEFAULT AS IDENTITY NOT NULL,
	country_id int4 NULL,
	gdp_nominal numeric NULL,
	gdp_ppp numeric NULL,
	inflation_rate numeric NULL,
	unemployment_rate numeric NULL,
	major_industries text NULL,
	created_at timestamptz NULL,
	CONSTRAINT economy_pk PRIMARY KEY (id),
	CONSTRAINT economy_economy_fk FOREIGN KEY (country_id) REFERENCES country_infos.economies(id)
);

