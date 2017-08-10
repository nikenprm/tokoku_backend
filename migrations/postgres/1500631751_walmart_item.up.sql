CREATE TABLE walmart_item (
    item_id BIGINT PRIMARY KEY,
    parent_item_id BIGINT,
    name VARCHAR(250),
    sale_price FLOAT,
    upc VARCHAR(20),
    category_path VARCHAR(150),
    long_description TEXT,
    brand_name VARCHAR(100),
    thumbnail_image VARCHAR(250),
    medium_image VARCHAR(250),
    large_image VARCHAR(250),
    size VARCHAR(50),
    product_url TEXT,
    category_node VARCHAR(100),
    stock VARCHAR(25)
);