#!/bin/bash

# Database Setup Script for Production Deployment
# This script runs the database setup SQL to ensure all tables and admin user exist

set -e  # Exit on any error

echo "🗄️  Setting up production database..."

# Database connection details (from deployment environment)
DB_HOST="${DB_HOST:-localhost}"
DB_PORT="${DB_PORT:-5432}"
DB_NAME="${DB_NAME:-church_db}"
DB_USER="${DB_USER:-church_user}"
DB_PASSWORD="${DB_PASSWORD:-church123}"

# Export password for psql
export PGPASSWORD="$DB_PASSWORD"

# Check if database is accessible
echo "🔍 Checking database connection..."
if ! psql -h "$DB_HOST" -p "$DB_PORT" -U "$DB_USER" -d "$DB_NAME" -c "SELECT 1;" > /dev/null 2>&1; then
    echo "❌ Cannot connect to database. Please check your database configuration."
    exit 1
fi

echo "✅ Database connection successful"

# Run the setup SQL script
echo "📝 Running database setup script..."
if psql -h "$DB_HOST" -p "$DB_PORT" -U "$DB_USER" -d "$DB_NAME" -f "$(dirname "$0")/setup-production-database.sql"; then
    echo "✅ Database setup completed successfully!"
    echo "🔐 Admin credentials:"
    echo "   Username: churchadmin"
    echo "   Password: ChurchPass123!"
else
    echo "❌ Database setup failed!"
    exit 1
fi

# Unset password
unset PGPASSWORD

echo "🎉 Database is ready for production deployment!"