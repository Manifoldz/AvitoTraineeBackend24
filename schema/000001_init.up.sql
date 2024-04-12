CREATE TABLE banners
(
    id SERIAL PRIMARY KEY,
	content JSONB NOT NULL,
	is_active  boolean NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE features
(
    id int PRIMARY KEY,
);

CREATE TABLE tags
(
    id int PRIMARY KEY,
);

CREATE TABLE bannerFeatureTag
(
    banner_id int NOT NULL REFERENCES banners (id) ON DELETE CASCADE,
    feature_id int NOT NULL REFERENCES features (id) ON DELETE CASCADE,
    tag_id int NOT NULL REFERENCES tags (id) ON DELETE CASCADE,
    PRIMARY KEY (banner_id, feature_id, tag_id)
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
