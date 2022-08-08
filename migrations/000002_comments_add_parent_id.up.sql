-- add parent_id
ALTER TABLE comments
ADD parent_id INT;
-- add parent_id foreign key
ALTER TABLE comments
ADD CONSTRAINT comments_parent_id_fk FOREIGN KEY (parent_id) REFERENCES comments(id) ON DELETE CASCADE ON UPDATE CASCADE;