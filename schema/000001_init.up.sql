CREATE TABLE banners
(
    id SERIAL PRIMARY KEY,
	feature_id int NOT NULL,
	content_id int NOT NULL,
	is_active  boolean NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (feature_id) REFERENCES features (id) ON DELETE CASCADE,
    FOREIGN KEY (content_id) REFERENCES contents (id) ON DELETE CASCADE,
);

CREATE TABLE features
(
    id SERIAL PRIMARY KEY,
	name_feature varchar(255) NOT NULL
);

CREATE TABLE tags
(
    id SERIAL PRIMARY KEY,
	name_tag varchar(255) NOT NULL
);

CREATE TABLE contents
(
    id SERIAL PRIMARY KEY,
	title varchar(255) NOT NULL,
	text_content  text NOT NULL,
	url_content   varchar(4096) NOT NULL
);


CREATE TABLE BannerTags
(
	banner_id  int NOT NULL,
	tag_id     int NOT NULL,
    PRIMARY KEY (banner_id, tag_id),
    FOREIGN KEY (banner_id) REFERENCES banners (id) ON DELETE CASCADE,
    FOREIGN KEY (tag_id) REFERENCES tags (id) ON DELETE CASCADE
);

/* В PostgreSQL нет встроенного функционала ON UPDATE, поэтому для автоматического обновления поля updated_at при каждом изменении записи необходимо создать триггер:
CREATE OR REPLACE FUNCTION update_modified_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_updated_at BEFORE UPDATE
ON banners FOR EACH ROW
EXECUTE FUNCTION update_modified_column(); */
