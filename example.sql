CREATE TABLE users (
    user_id INT AUTO_INCREMENT PRIMARY KEY,  
    username VARCHAR(50) NOT NULL COMMENT 'Username of the user',
    email VARCHAR(100) NOT NULL UNIQUE COMMENT 'Email address of the user (must be unique)',
    password_hash VARCHAR(255) NOT NULL COMMENT 'Hashed password of the user',
    registration_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,  
    last_login TIMESTAMP,  
    is_admin BOOLEAN DEFAULT FALSE  
);

CREATE TABLE products (
    product_id INT AUTO_INCREMENT PRIMARY KEY,  
    product_name VARCHAR(100) NOT NULL,  
    description TEXT,  
    price DECIMAL(10, 2) NOT NULL COMMENT 'Price of the product',
    stock_quantity INT NOT NULL DEFAULT 0  
);

CREATE TABLE orders (
    order_id INT AUTO_INCREMENT PRIMARY KEY,  
    user_id INT NOT NULL,  
    order_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,  
    total_amount DECIMAL(10, 2) NOT NULL COMMENT 'Total amount of the order',
    FOREIGN KEY (user_id) REFERENCES users(user_id)
);

CREATE TABLE order_details (
    detail_id INT AUTO_INCREMENT PRIMARY KEY,  
    order_id INT NOT NULL,  
    product_id INT NOT NULL,  
    quantity INT NOT NULL DEFAULT 1 COMMENT 'Quantity of the product being ordered, defaults to 1',
    subtotal DECIMAL(10, 2) NOT NULL,
    FOREIGN KEY (order_id) REFERENCES orders(order_id),
    FOREIGN KEY (product_id) REFERENCES products(product_id)
);

CREATE TABLE categories (
    category_id INT AUTO_INCREMENT PRIMARY KEY,  
    category_name VARCHAR(50) NOT NULL,  
    description TEXT  
);

CREATE TABLE reviews (
    review_id INT AUTO_INCREMENT PRIMARY KEY,  
    product_id INT NOT NULL,  
    user_id INT NOT NULL,  
    rating INT NOT NULL COMMENT 'Rating given by the user for the product',
    comment TEXT,  
    review_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (product_id) REFERENCES products(product_id),
    FOREIGN KEY (user_id) REFERENCES users(user_id)
);

CREATE TABLE invoices (
    invoice_id INT AUTO_INCREMENT PRIMARY KEY,  
    order_id INT NOT NULL,  
    amount DECIMAL(10, 2) NOT NULL,  
    invoice_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (order_id) REFERENCES orders(order_id)
);

CREATE TABLE regions (
    region_id INT AUTO_INCREMENT PRIMARY KEY,  
    region_name VARCHAR(50) NOT NULL,  
    country VARCHAR(50) NOT NULL  
);

CREATE TABLE shippers (
    shipper_id INT AUTO_INCREMENT PRIMARY KEY,  
    shipper_name VARCHAR(100) NOT NULL,  
    phone VARCHAR(20) NOT NULL COMMENT 'Contact phone number for the shipper'
);

CREATE TABLE employees (
    employee_id INT AUTO_INCREMENT PRIMARY KEY,  
    last_name VARCHAR(50) NOT NULL,  
    first_name VARCHAR(50) NOT NULL,  
    birth_date DATE NOT NULL,  
    hire_date DATE NOT NULL,  
    address VARCHAR(100),  
    city VARCHAR(50),  
    region_id INT,  
    postal_code VARCHAR(20),  
    country VARCHAR(50) NOT NULL,  
    phone VARCHAR(20),  
    photo BLOB,  
    notes TEXT,  
    manager_id INT,
    FOREIGN KEY (region_id) REFERENCES regions(region_id),
    FOREIGN KEY (manager_id) REFERENCES employees(employee_id)
);

CREATE TABLE carts (
    cart_id INT AUTO_INCREMENT PRIMARY KEY,  
    user_id INT NOT NULL,  
    product_id INT NOT NULL,  
    quantity INT NOT NULL DEFAULT 1 COMMENT 'Quantity of the product being ordered, defaults to 1',
    added_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(user_id),
    FOREIGN KEY (product_id) REFERENCES products(product_id)
);
