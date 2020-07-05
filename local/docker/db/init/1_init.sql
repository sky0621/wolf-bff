
CREATE TYPE content_type AS ENUM('Text', 'Image', 'Voice', 'Movie', 'Other');

CREATE TABLE wht (
  id bigserial NOT NULL,
  recordDate timestamp NOT NULL,
  title varchar(256),
  content_type content_type NOT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE content_text (
  id bigserial NOT NULL,
  wht_id bigint REFERENCES wht(id) NOT NULL,
  name varchar(256),
  text text NOT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE content_image (
  id bigserial NOT NULL,
  wht_id bigint REFERENCES wht(id) NOT NULL,
  name varchar(256),
  path varchar(1024) NOT NULL,
  PRIMARY KEY (id)
);
