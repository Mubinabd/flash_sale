-- ENUM TYPES
CREATE TYPE role_type AS ENUM ('admin', 'user');
CREATE TYPE status_type AS ENUM ('pending', 'active', 'completed', 'canceled');
CREATE TYPE order_type AS ENUM ('pending', 'confirmed', 'canceled', 'refunded');
CREATE TYPE types AS ENUM ('sms', 'email');
CREATE TYPE notification_type AS ENUM ('pending', 'sent', 'failed', 'read');
CREATE TYPE transaction_type AS ENUM ('debit', 'credit');

-- USERS TABLE
CREATE TABLE IF NOT EXISTS users (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    full_name VARCHAR(100),
    date_of_birth DATE,
    role role_type DEFAULT 'user',
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at BIGINT DEFAULT 0
);

-- SETTINGS TABLE
CREATE TABLE IF NOT EXISTS settings (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users(id),
    privacy_level VARCHAR(50) NOT NULL DEFAULT 'private',
    notification VARCHAR(30) NOT NULL DEFAULT 'on',
    language VARCHAR(255) NOT NULL DEFAULT 'en',
    theme VARCHAR(255) NOT NULL DEFAULT 'light',
    updated_at TIMESTAMP DEFAULT NOW()
);

-- TOKENS TABLE
CREATE TABLE IF NOT EXISTS tokens (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users(id),
    token VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at BIGINT DEFAULT 0
);

-- PRODUCTS TABLE
CREATE TABLE IF NOT EXISTS products (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR NOT NULL,
    description TEXT NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    image_url VARCHAR NOT NULL,
    stock_quantity INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at BIGINT DEFAULT 0
);

-- FLASH SALES TABLE
CREATE TABLE IF NOT EXISTS flash_sales (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    start_time TIMESTAMP NOT NULL,
    end_time TIMESTAMP NOT NULL,
    status status_type NOT NULL,
    latitude DOUBLE PRECISION,
    longitude DOUBLE PRECISION,
    address VARCHAR,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at BIGINT DEFAULT 0
);

-- FLASH SALES PRODUCTS TABLE
CREATE TABLE IF NOT EXISTS flash_sales_products (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    flash_sale_id UUID NOT NULL REFERENCES flash_sales(id),
    product_id UUID NOT NULL REFERENCES products(id),
    discounted_price DECIMAL(10, 2) NOT NULL,
    available_quantity INTEGER,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at BIGINT DEFAULT 0
);

-- ORDERS TABLE
CREATE TABLE IF NOT EXISTS orders (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id),
    flash_sale_id UUID NOT NULL REFERENCES flash_sales(id),
    status order_type NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at BIGINT DEFAULT 0
);

-- NOTIFICATION TABLE
CREATE TABLE IF NOT EXISTS notification (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id),
    type types NOT NULL,
    content TEXT NOT NULL,
    status notification_type NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at BIGINT DEFAULT 0
);

-- REVIEWS TABLE
CREATE TABLE IF NOT EXISTS reviews (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users(id),
    product_id UUID NOT NULL REFERENCES products(id),
    rating INTEGER CHECK (rating >= 1 AND rating <= 5) NOT NULL,
    review_text TEXT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at BIGINT DEFAULT 0
);

-- ORDER STATUS TRACKING TABLE
CREATE TABLE IF NOT EXISTS order_status_tracking (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    order_id UUID NOT NULL REFERENCES orders(id),
    order_status VARCHAR NOT NULL,
    estimated_delivery TIMESTAMP,
    current_location VARCHAR,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP
);

-- REFUNDS TABLE
CREATE TABLE IF NOT EXISTS refunds (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    order_id UUID NOT NULL REFERENCES orders(id),
    refund_status VARCHAR NOT NULL,
    refund_amount DECIMAL(10, 2) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP
);

-- SHARED DEALS TABLE
CREATE TABLE shared_deals (
    id SERIAL PRIMARY KEY,
    user_id VARCHAR(255),
    flash_sale_id VARCHAR(255),
    platform VARCHAR(255),
    message TEXT,
    shared_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- SHARING STATS TABLE
CREATE TABLE sharing_stats (
    flash_sale_id VARCHAR(255) PRIMARY KEY,
    total_shares BIGINT DEFAULT 0,
    shares_by_platform JSONB
);

-- FLASH SALE PRODUCTS TABLE
CREATE TABLE IF NOT EXISTS flash_sale_products (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    flash_sale_id UUID NOT NULL REFERENCES flash_sales(id),
    product_id UUID NOT NULL,
    added_at TIMESTAMP DEFAULT NOW()
);

-- FLASH SALE CANCELLATIONS TABLE
CREATE TABLE IF NOT EXISTS flash_sale_cancellations (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    flash_sale_id UUID NOT NULL REFERENCES flash_sales(id),
    cancellation_status VARCHAR NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

-- STORES TABLE
CREATE TABLE IF NOT EXISTS stores (
    store_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR NOT NULL,
    address VARCHAR NOT NULL,
    latitude DOUBLE PRECISION NOT NULL,
    longitude DOUBLE PRECISION NOT NULL
);
