-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

CREATE TABLE country_phone_formats (
    id SERIAL PRIMARY KEY,
    country_code VARCHAR(2) UNIQUE NOT NULL,
    country_name VARCHAR(100) NOT NULL,
    phone_regex TEXT NOT NULL,
    phone_mask VARCHAR(50) NOT NULL,
    min_length INTEGER DEFAULT 0,
    max_length INTEGER DEFAULT 20,
    is_active BOOLEAN DEFAULT true,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO country_phone_formats (country_code, country_name, phone_regex, phone_mask, min_length, max_length) VALUES
('RU', 'Russia', '^(\+7|8)[\s\-]?\(?[489][0-9]{2}\)?[\s\-]?[0-9]{3}[\s\-]?[0-9]{2}[\s\-]?[0-9]{2}$', '+7 (XXX) XXX-XX-XX', 11, 12),
('US', 'United States', '^\+1[2-9]\d{2}[2-9](?!11)\d{6}$', '+1 (XXX) XXX-XXXX', 11, 11),
('GB', 'United Kingdom', '^(\+44|0)7\d{9}$', '+44 XXXX XXXXXX', 10, 12),
('DE', 'Germany', '^(\+49|0)[1-9]\d{1,4}[\s\-]?\d{3,13}$', '+49 XXX XXXXXXX', 10, 15),
('FR', 'France', '^(\+33|0)[1-9](\d{2}){4}$', '+33 X XX XX XX XX', 9, 12),
('IT', 'Italy', '^(\+39)?[3]\d{8,9}$', '+39 XXX XXXXXX', 9, 12),
('UA', 'Ukraine', '^(\+38)?0\d{9}$', '+380 XX XXX XX XX', 10, 12),
('KZ', 'Kazakhstan', '^(\+7|8)[\s\-]?\(?[67][0-9]{2}\)?[\s\-]?[0-9]{3}[\s\-]?[0-9]{2}[\s\-]?[0-9]{2}$', '+7 (XXX) XXX-XX-XX', 11, 12),
('BY', 'Belarus', '^(\+375)?(24|25|29|33|44)\d{7}$', '+375 (XX) XXX-XX-XX', 9, 12);


-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd

DROP TABLE country_phone_formats;
