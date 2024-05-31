CREATE TABLE users (
    user_id INT AUTO_INCREMENT PRIMARY KEY COMMENT 'Unique identifier for each user in the system',
    username VARCHAR(50) NOT NULL COMMENT 'User-selected name for logging into the system, must be unique',
    email VARCHAR(100) NOT NULL UNIQUE COMMENT 'User\'s email address, used for notifications and password recovery',
    password_hash VARCHAR(255) NOT NULL COMMENT 'Securely stored hashed password for user authentication',
    registration_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT 'Date and time when the user registered in the system',
    last_login TIMESTAMP COMMENT 'Date and time when the user last logged into the system',
    is_admin BOOLEAN DEFAULT FALSE COMMENT 'Indicates whether the user has administrative privileges (TRUE/FALSE)'
) COMMENT='Stores basic information about users including login credentials and administrative status';

CREATE TABLE products (
    product_id INT AUTO_INCREMENT PRIMARY KEY COMMENT 'Unique identifier for each product in the catalog',
    product_name VARCHAR(100) NOT NULL COMMENT 'Name of the product as displayed to customers',
    description TEXT COMMENT 'Detailed description of the product, including features and specifications',
    price DECIMAL(10, 2) NOT NULL COMMENT 'Retail price of the product in the store\'s currency',
    stock_quantity INT NOT NULL DEFAULT 0 COMMENT 'Current stock level of the product, used for inventory management'
) COMMENT='Stores information about products available for sale, including pricing and stock levels';

CREATE TABLE orders (
    order_id INT AUTO_INCREMENT PRIMARY KEY COMMENT 'Unique identifier for each customer order',
    user_id INT NOT NULL COMMENT 'Identifier of the user who placed the order, references users(user_id)',
    order_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT 'Date and time when the order was created',
    total_amount DECIMAL(10, 2) NOT NULL COMMENT 'Total monetary value of the order, including taxes and shipping',
    FOREIGN KEY (user_id) REFERENCES users(user_id)
) COMMENT='Contains customer order information including order date and total amount';

CREATE TABLE order_details (
    detail_id INT AUTO_INCREMENT PRIMARY KEY COMMENT 'Unique identifier for each order line item',
    order_id INT NOT NULL COMMENT 'Identifier of the order this line item belongs to, references orders(order_id)',
    product_id INT NOT NULL COMMENT 'Identifier of the product in this line item, references products(product_id)',
    quantity INT NOT NULL DEFAULT 1 COMMENT 'Number of units of the product ordered',
    subtotal DECIMAL(10, 2) NOT NULL COMMENT 'Subtotal amount for this line item (quantity * product price)',
    FOREIGN KEY (order_id) REFERENCES orders(order_id),
    FOREIGN KEY (product_id) REFERENCES products(product_id)
) COMMENT='Stores details for each product included in a customer order';

CREATE TABLE categories (
    category_id INT AUTO_INCREMENT PRIMARY KEY COMMENT 'Unique identifier for each product category',
    category_name VARCHAR(50) NOT NULL COMMENT 'Name of the category used to group similar products',
    description TEXT COMMENT 'Detailed description of the product category, including examples and usage'
) COMMENT='Stores information about product categories for organizational purposes';

CREATE TABLE reviews (
    review_id INT AUTO_INCREMENT PRIMARY KEY COMMENT 'Unique identifier for each product review',
    product_id INT NOT NULL COMMENT 'Identifier of the product being reviewed, references products(product_id)',
    user_id INT NOT NULL COMMENT 'Identifier of the user who wrote the review, references users(user_id)',
    rating INT NOT NULL COMMENT 'Numeric rating given by the user to the product, usually from 1 to 5',
    comment TEXT COMMENT 'Textual feedback provided by the user about the product',
    review_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT 'Date and time when the review was submitted',
    FOREIGN KEY (product_id) REFERENCES products(product_id),
    FOREIGN KEY (user_id) REFERENCES users(user_id)
) COMMENT='Stores user-generated reviews for products, including ratings and comments';

CREATE TABLE invoices (
    invoice_id INT AUTO_INCREMENT PRIMARY KEY COMMENT 'Unique identifier for each invoice issued',
    order_id INT NOT NULL COMMENT 'Identifier of the order associated with this invoice, references orders(order_id)',
    amount DECIMAL(10, 2) NOT NULL COMMENT 'Total amount to be paid as per the invoice',
    invoice_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT 'Date and time when the invoice was generated',
    FOREIGN KEY (order_id) REFERENCES orders(order_id)
) COMMENT='Contains invoice information for orders, including total amount and invoice date';

CREATE TABLE regions (
    region_id INT AUTO_INCREMENT PRIMARY KEY COMMENT 'Unique identifier for each geographical region',
    region_name VARCHAR(50) NOT NULL COMMENT 'Name of the geographical region',
    country VARCHAR(50) NOT NULL COMMENT 'Country where the region is located'
) COMMENT='Stores information about different geographical regions, used for shipping and employee assignments';

CREATE TABLE shippers (
    shipper_id INT AUTO_INCREMENT PRIMARY KEY COMMENT 'Unique identifier for each shipping company',
    shipper_name VARCHAR(100) NOT NULL COMMENT 'Name of the shipping company',
    phone VARCHAR(20) NOT NULL COMMENT 'Contact phone number for the shipping company'
) COMMENT='Contains information about shipping companies responsible for delivering orders';

CREATE TABLE employees (
    employee_id INT AUTO_INCREMENT PRIMARY KEY COMMENT 'Unique identifier for each employee',
    last_name VARCHAR(50) NOT NULL COMMENT 'Last name of the employee',
    first_name VARCHAR(50) NOT NULL COMMENT 'First name of the employee',
    birth_date DATE NOT NULL COMMENT 'Birth date of the employee',
    hire_date DATE NOT NULL COMMENT 'Date when the employee was hired',
    address VARCHAR(100) COMMENT 'Residential address of the employee',
    city VARCHAR(50) COMMENT 'City where the employee resides',
    region_id INT COMMENT 'Identifier of the region where the employee works, references regions(region_id)',
    postal_code VARCHAR(20) COMMENT 'Postal code of the employee\'s residential address',
    country VARCHAR(50) NOT NULL COMMENT 'Country where the employee resides',
    phone VARCHAR(20) COMMENT 'Contact phone number of the employee',
    photo BLOB COMMENT 'Photograph of the employee',
    notes TEXT COMMENT 'Additional notes or comments about the employee',
    manager_id INT COMMENT 'Identifier of the employee\'s manager, references employees(employee_id)',
    FOREIGN KEY (region_id) REFERENCES regions(region_id),
    FOREIGN KEY (manager_id) REFERENCES employees(employee_id)
) COMMENT='Stores detailed information about employees, including personal details and job-related information';

CREATE TABLE carts (
    cart_id INT AUTO_INCREMENT PRIMARY KEY COMMENT 'Unique identifier for each shopping cart',
    user_id INT NOT NULL COMMENT 'Identifier of the user who owns the cart, references users(user_id)',
    product_id INT NOT NULL COMMENT 'Identifier of the product in the cart, references products(product_id)',
    quantity INT NOT NULL DEFAULT 1 COMMENT 'Number of units of the product added to the cart',
    added_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT 'Date and time when the product was added to the cart',
    FOREIGN KEY (user_id) REFERENCES users(user_id),
    FOREIGN KEY (product_id) REFERENCES products(product_id)
) COMMENT='Stores information about products added to users\' shopping carts';
