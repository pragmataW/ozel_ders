-- Veritabanı oluşturuluyor
CREATE DATABASE authentication;

-- Veritabanına bağlanılıyor
\c authentication;

-- Kullanıcı tablosu oluşturuluyor
CREATE TABLE users (
    id SERIAL PRIMARY KEY,                   -- Otomatik artan birincil anahtar
    user_name VARCHAR(255) UNIQUE NOT NULL,  -- unique ve NOT NULL olan user_name
    password VARCHAR(255) NOT NULL,          -- NOT NULL olan password
    created_at TIMESTAMPTZ DEFAULT NOW(),    -- Oluşturulma zamanı
    updated_at TIMESTAMPTZ DEFAULT NOW(),    -- Güncellenme zamanı
    deleted_at TIMESTAMPTZ                  -- Silinme zamanı (soft delete için)
);

-- Örnek veri eklemek isterseniz, bunu yapabilirsiniz
-- INSERT INTO users (user_name, password) VALUES ('example_user', 'example_password');
