-- +goose Up
CREATE TABLE users (
    user_id UUID PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    phone VARCHAR(20),
    balance NUMERIC DEFAULT 0,
    total_spent NUMERIC DEFAULT 0,
    total_earned NUMERIC DEFAULT 0,
    first_name VARCHAR(100),
    last_name VARCHAR(100),
    middle_name VARCHAR(100),
    date_of_birth DATE,
    avatar_url TEXT,
    country VARCHAR(100),
    city VARCHAR(100),
    address TEXT,
    postal_code VARCHAR(20),
    is_verified BOOLEAN DEFAULT FALSE,
    is_active BOOLEAN DEFAULT TRUE,
    is_seller BOOLEAN DEFAULT FALSE,
    rating NUMERIC DEFAULT 0,
    review_count INT DEFAULT 0,
    last_password_change TIMESTAMP,
    failed_login_attempts INT DEFAULT 0,
    is_locked BOOLEAN DEFAULT FALSE,
    email_notifications BOOLEAN DEFAULT TRUE,
    sms_notifications BOOLEAN DEFAULT FALSE,
    push_notifications BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now(),
    last_login_at TIMESTAMP,
    deleted_at TIMESTAMP
);

-- +goose Down
DROP TABLE users;
