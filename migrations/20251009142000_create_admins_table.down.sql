-- Drop indexes
DROP INDEX IF EXISTS idx_admins_is_active;
DROP INDEX IF EXISTS idx_admins_email;
DROP INDEX IF EXISTS idx_admins_username;

-- Drop admins table
DROP TABLE IF EXISTS admins;

