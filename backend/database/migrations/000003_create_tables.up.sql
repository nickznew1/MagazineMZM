CREATE TABLE customer
(
    id       SERIAL PRIMARY KEY,
    login    VARCHAR(50) UNIQUE NOT NULL,
    password VARCHAR(30)  NOT NULL,
    email    VARCHAR(50) UNIQUE NOT NULL,
    user_role VARCHAR(50) NOT NULL DEFAULT 'ordinary',
    registration_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE customer_personal_info
(
    id INT,
    company TEXT,
    first_name Text,
    second_name TEXT,
    FOREIGN KEY(id) REFERENCES customer(id)
);


CREATE TABLE customer_delivery_info (
                                        id INT,
                                        phone_number TEXT,
                                        city TEXT,
                                        address TEXT,
                                        FOREIGN KEY(id) REFERENCES customer(id)
);


CREATE TABLE item
(
    id    SERIAL PRIMARY KEY,
    name  TEXT UNIQUE NOT NULL,
    price INT,
    item_type TEXT NOT NULL,
    secondary_type TEXT NOT NULL,
    item_picture TEXT,
    item_description TEXT NOT NULL,
    item_short_description TEXT NOT NULL,
    article INT,
    visible bool NOT NULL DEFAULT false
);

CREATE TABLE item_spec_files
(
    id INT PRIMARY KEY,p
    name TEXT NOT NULL,
    link TEXT NOT NULL,
    picture TEXT NOT NULL,
    FOREIGN KEY (id) REFERENCES item(id) ON DELETE CASCADE
);


CREATE TABLE item_properties
(
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL
);


CREATE TABLE item_properties_values (
                                        item_id INT,
                                        property_id INT,
                                        value varchar(500) NOT NULL,
                                        PRIMARY KEY(item_id, property_id),
                                        FOREIGN KEY (item_id) REFERENCES item(id) ON DELETE CASCADE,
                                        FOREIGN KEY (property_id) REFERENCES item_properties(id)
);


CREATE TABLE customer_item
(
    item_id INT,
    count INT,
    customer_id INT,
    props JSONB,
    /*describe_item */
    FOREIGN KEY (customer_id) REFERENCES customer (id),
    FOREIGN KEY (item_id) REFERENCES item (id) ON DELETE CASCADE
);

CREATE TABLE spec
(
    id SERIAL PRIMARY KEY,
    item_id INT,
    item_patch TEXT UNIQUE NOT NULL,
    name TEXT UNIQUE NOT NULL,
    /*misc?*/
    FOREIGN KEY (item_id) REFERENCES item (id) ON DELETE CASCADE
);


CREATE TABLE user_applications
(
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    email VARCHAR (50) NOT NULL,
    first_name VARCHAR (30) NOT NULL,
    second_name VARCHAR (30) NOT NULL,
    login VARCHAR(30) NOT NULL,
    phone_number VARCHAR(30) NOT NULL,
    company VARCHAR(50) NOT NULL,
    address VARCHAR(80) NOT NULL,
    city VARCHAR(30) NOT NULL,
    order_date DATE,
    items jsonb,
    order_status VARCHAR (20) NOT NULL DEFAULT 'В обработке',
    FOREIGN KEY (user_id) REFERENCES customer(id)
);