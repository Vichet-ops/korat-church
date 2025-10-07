-- Create members table
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

-- Create events table
CREATE TABLE IF NOT EXISTS events (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    date TIMESTAMP WITH TIME ZONE NOT NULL,
    time VARCHAR(50),
    location VARCHAR(255),
    is_public BOOLEAN DEFAULT true,
    created_by INTEGER REFERENCES users(id),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE
);

-- Create church_settings table
CREATE TABLE IF NOT EXISTS church_settings (
    id SERIAL PRIMARY KEY,
    church_name VARCHAR(255) NOT NULL,
    church_address TEXT,
    church_phone VARCHAR(50),
    church_email VARCHAR(255),
    service_times TEXT, -- JSON string for multiple service times
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Insert default church settings
INSERT INTO church_settings (church_name, church_address, church_phone, church_email, service_times) 
VALUES (
    'Muang Thai Korat Church',
    '123 Main St, Nakhon Ratchasima, Thailand 30000',
    '(555) 123-4567',
    'info@muangthaikoratchurch.com',
    '[{"day":"Sunday","time":"10:00 AM","service":"Main Worship Service"},{"day":"Wednesday","time":"7:00 PM","service":"Bible Study & Prayer"},{"day":"Friday","time":"6:30 PM","service":"Youth Activities & Fellowship"}]'
);

-- Insert sample events
INSERT INTO events (title, description, date, time, location, is_public) VALUES
('Christmas Eve Service', 'Join us for a special candlelight service celebrating the birth of Jesus Christ.', '2024-12-24 19:00:00', '7:00 PM', 'Main Sanctuary', true),
('Community Potluck', 'Share a meal and fellowship with our church family after Sunday service.', '2024-12-29 12:00:00', '12:00 PM', 'Fellowship Hall', true),
('Bible Study Series', 'Join our new Bible study series exploring the Book of Romans.', '2025-01-15 19:00:00', '7:00 PM', 'Conference Room', true);
