-- Enable the extension for UUID generation
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Users Table (Assuming this exists, referenced by client_id in orders)
CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    username VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    -- other fields as necessary
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now(),
    deleted_at BIGINT DEFAULT 0
);

-- Products Table
CREATE TABLE IF NOT EXISTS products (
    product_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    description TEXT,
    price NUMERIC(10, 2) NOT NULL,
    stock INT NOT NULL,
    category VARCHAR(100),
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now(),
    deleted_at BIGINT DEFAULT 0
);

-- Orders Table
CREATE TABLE IF NOT EXISTS orders (
    order_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    client_id UUID NOT NULL,
    product_id UUID NOT NULL,
    quantity INT NOT NULL,
    address TEXT,
    order_date TIMESTAMPTZ,
    delivery_date TIMESTAMPTZ,
    status VARCHAR(50),
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now(),
    deleted_at BIGINT DEFAULT 0,
    FOREIGN KEY (product_id) REFERENCES products(product_id),
    FOREIGN KEY (client_id) REFERENCES users(id) -- Reference to the users table
);

-- Offices Table
CREATE TABLE IF NOT EXISTS offices (
    office_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    address TEXT NOT NULL,
    phone_number VARCHAR(20),
    email VARCHAR(255),
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now(),
    deleted_at BIGINT DEFAULT 0
);

-- Carts Table
CREATE TABLE IF NOT EXISTS carts (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL,
    status VARCHAR(50),
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now(),
    deleted_at BIGINT DEFAULT 0,
    FOREIGN KEY (user_id) REFERENCES users(id) -- Reference to the users table
);

-- Cart Items Table
CREATE TABLE IF NOT EXISTS cart_items (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    cart_id UUID NOT NULL,
    product_id UUID NOT NULL,
    quantity INT NOT NULL,
    price NUMERIC(10, 2) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now(),
    deleted_at BIGINT DEFAULT 0,
    FOREIGN KEY (cart_id) REFERENCES carts(id),
    FOREIGN KEY (product_id) REFERENCES products(product_id)
);

