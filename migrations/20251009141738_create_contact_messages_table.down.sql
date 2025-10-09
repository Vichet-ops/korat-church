-- Drop indexes
DROP INDEX IF EXISTS idx_contact_messages_created_at;
DROP INDEX IF EXISTS idx_contact_messages_status;

-- Drop contact_messages table
DROP TABLE IF EXISTS contact_messages;

