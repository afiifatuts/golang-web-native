ALTER TABLE categories ALTER COLUMN id SET DATA TYPE SERIAL;

CREATE SEQUENCE my_table_id_seq2;
ALTER TABLE products ALTER COLUMN id SET DEFAULT NEXTVAL('my_table_id_seq2');

ALTER TABLE products
ADD CONSTRAINT products_category_id_fkey
FOREIGN KEY (category_id)
REFERENCES categories(id);

ALTER TABLE products
DROP CONSTRAINT products_category_id_fkey;