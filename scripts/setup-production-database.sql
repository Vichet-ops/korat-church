-- Production Database Setup Script
-- This script ensures all required tables and admin user exist for production deployment

-- Create admins table if it doesn't exist
CREATE TABLE IF NOT EXISTS admins (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    first_name VARCHAR(100),
    last_name VARCHAR(100),
    is_active BOOLEAN DEFAULT true,
    last_login TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Create contact_messages table if it doesn't exist
CREATE TABLE IF NOT EXISTS contact_messages (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    subject VARCHAR(500) NOT NULL,
    message TEXT NOT NULL,
    status VARCHAR(20) DEFAULT 'new' CHECK (status IN ('new', 'read', 'replied')),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Create church_settings table if it doesn't exist
CREATE TABLE IF NOT EXISTS church_settings (
    id SERIAL PRIMARY KEY,
    church_name VARCHAR(255) NOT NULL,
    church_address TEXT,
    church_phone VARCHAR(50),
    church_email VARCHAR(255),
    service_times TEXT,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Create events table if it doesn't exist
CREATE TABLE IF NOT EXISTS events (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    date TIMESTAMP WITH TIME ZONE NOT NULL,
    time VARCHAR(50),
    location VARCHAR(255),
    is_public BOOLEAN DEFAULT true,
    created_by INTEGER,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE
);

-- Create members table if it doesn't exist
CREATE TABLE IF NOT EXISTS members (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE,
    phone VARCHAR(50),
    address TEXT,
    date_of_birth DATE,
    join_date TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    is_active BOOLEAN DEFAULT true,
    notes TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE
);

-- Create indexes for performance
CREATE INDEX IF NOT EXISTS idx_admins_username ON admins(username);
CREATE INDEX IF NOT EXISTS idx_admins_email ON admins(email);
CREATE INDEX IF NOT EXISTS idx_admins_is_active ON admins(is_active);
CREATE INDEX IF NOT EXISTS idx_contact_messages_status ON contact_messages(status);
CREATE INDEX IF NOT EXISTS idx_contact_messages_created_at ON contact_messages(created_at DESC);
CREATE INDEX IF NOT EXISTS idx_events_date ON events(date);
CREATE INDEX IF NOT EXISTS idx_events_is_public ON events(is_public);

-- Insert default church settings if none exist
INSERT INTO church_settings (church_name, church_address, church_phone, church_email, service_times) 
SELECT 'Muang Thai Korat Church', '123 Main St, Nakhon Ratchasima, Thailand 30000', '(555) 123-4567', 'info@muangthaikoratchurch.com', '[{"day":"Sunday","time":"10:00 AM","service":"Main Worship Service"},{"day":"Wednesday","time":"7:00 PM","service":"Bible Study & Prayer"},{"day":"Friday","time":"6:30 PM","service":"Youth Activities & Fellowship"}]'
WHERE NOT EXISTS (SELECT 1 FROM church_settings);

-- Create or update admin user with the correct password hash for 'ChurchPass123!'
-- This will either insert a new admin or update the existing one
INSERT INTO admins (username, email, password, first_name, last_name, is_active) 
VALUES ('churchadmin', 'churchadmin@muangthaikoratchurch.com', '$2a$10$iPOgWX6Kk/lqh5OvscNhPe/TL/lPsg0zpsXG05qzlxS6oe77AcEaO', 'Church', 'Administrator', true)
ON CONFLICT (username) 
DO UPDATE SET 
    password = EXCLUDED.password,
    email = EXCLUDED.email,
    first_name = EXCLUDED.first_name,
    last_name = EXCLUDED.last_name,
    is_active = EXCLUDED.is_active,
    updated_at = CURRENT_TIMESTAMP;

-- Insert sample events if none exist
INSERT INTO events (title, description, date, time, location, is_public) 
SELECT 'Christmas Eve Service', 'Join us for a special candlelight service celebrating the birth of Jesus Christ.', '2024-12-24 19:00:00', '7:00 PM', 'Main Sanctuary', true
WHERE NOT EXISTS (SELECT 1 FROM events);

INSERT INTO events (title, description, date, time, location, is_public) 
SELECT 'Community Potluck', 'Share a meal and fellowship with our church family after Sunday service.', '2024-12-29 12:00:00', '12:00 PM', 'Fellowship Hall', true
WHERE NOT EXISTS (SELECT 1 FROM events WHERE title = 'Community Potluck');

INSERT INTO events (title, description, date, time, location, is_public) 
SELECT 'Bible Study Series', 'Join our new Bible study series exploring the Book of Romans.', '2025-01-15 19:00:00', '7:00 PM', 'Conference Room', true
WHERE NOT EXISTS (SELECT 1 FROM events WHERE title = 'Bible Study Series');

-- Verify setup
SELECT 'Database setup completed successfully!' as status;
SELECT 'Admin user: churchadmin' as admin_info;
SELECT 'Password: ChurchPass123!' as password_info;