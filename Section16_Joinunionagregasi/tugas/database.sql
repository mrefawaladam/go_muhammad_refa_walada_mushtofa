CREATE TABLE product_types (
  id INT(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(255),
  created_at TIMESTAMP,
  updated_at TIMESTAMP
);

CREATE TABLE operators (
  id INT(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(255),
  created_at TIMESTAMP,
  updated_at TIMESTAMP
);

CREATE TABLE product_descriptions (
  id INT(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
  description TEXT, 
  created_at TIMESTAMP,
  updated_at TIMESTAMP
);

CREATE TABLE products (
  id INT(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
  product_type_id INT(11),
  operator_id INT(11),
  code VARCHAR(50),
  name VARCHAR(100),
  status SMALLINT,
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  FOREIGN KEY (product_type_id) REFERENCES product_types(id),
  FOREIGN KEY (operator_id) REFERENCES operators(id)
);

CREATE TABLE payment_methods (
  id INT(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(255),
  status SMALLINT,
  created_at TIMESTAMP,
  updated_at TIMESTAMP
);

CREATE TABLE users (
  id INT(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
  status SMALLINT,
  dob DATE,
  gender CHAR(1),
  created_at TIMESTAMP,
  updated_at TIMESTAMP
);

CREATE TABLE transactions (
  id INT(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
  user_id INT(11),
  payment_method_id INT(11),
  status VARCHAR(10),
  total_qty INT(11),
  total_price NUMERIC(25,2),
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  FOREIGN KEY (user_id) REFERENCES users(id),
  FOREIGN KEY (payment_method_id) REFERENCES payment_methods(id)
);

CREATE TABLE transaction_details (
  id INT(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
  transaction_id INT(11),
  product_id INT(11),
  status VARCHAR(10),
  qty INT(11),
  price NUMERIC(25,2),
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  FOREIGN KEY (transaction_id) REFERENCES transactions(id),
  FOREIGN KEY (product_id) REFERENCES products(id)
);
