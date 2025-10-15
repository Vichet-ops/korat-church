#!/bin/bash

# 🚀 Church Website Auto Deploy Script
# This script handles the actual deployment on your server

set -e  # Exit on any error

echo "🚀 Starting Church Website Deployment..."

# Configuration
APP_NAME="church-website"
APP_PATH="/var/www/church-website"
SERVICE_NAME="church-website"
BACKUP_DIR="/var/backups/church-website"

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Logging function
log() {
    echo -e "${BLUE}[$(date +'%Y-%m-%d %H:%M:%S')]${NC} $1"
}

error() {
    echo -e "${RED}[ERROR]${NC} $1"
    exit 1
}

success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

# Check if running as root or with sudo
if [ "$EUID" -ne 0 ]; then
    error "Please run this script with sudo"
fi

log "Stopping $SERVICE_NAME service..."
systemctl stop $SERVICE_NAME || warning "Service $SERVICE_NAME not running"

# Create backup
log "Creating backup..."
mkdir -p $BACKUP_DIR
BACKUP_FILE="$BACKUP_DIR/backup_$(date +%Y%m%d_%H%M%S).tar.gz"
if [ -d "$APP_PATH" ]; then
    tar -czf "$BACKUP_FILE" -C "$(dirname $APP_PATH)" "$(basename $APP_PATH)"
    success "Backup created: $BACKUP_FILE"
else
    warning "No existing app directory found for backup"
fi

# Extract new version
log "Extracting new version..."
if [ -f "church-app.tar.gz" ]; then
    # Clean old directory
    rm -rf $APP_PATH
    
    # Extract new version
    mkdir -p $APP_PATH
    tar -xzf church-app.tar.gz -C $APP_PATH
    success "Extracted new version to $APP_PATH"
else
    error "church-app.tar.gz not found!"
fi

# Set permissions
log "Setting permissions..."
chown -R www-data:www-data $APP_PATH
chmod +x $APP_PATH/church-app
success "Permissions set correctly"

# Install dependencies and build
log "Installing dependencies and building..."
cd $APP_PATH

# Install Go dependencies
log "Installing Go dependencies..."
go mod download

# Build frontend
log "Building frontend..."
cd frontend
npm ci --production
npm run build
cd ..

success "Dependencies installed and frontend built"

# Restart services
log "Starting $SERVICE_NAME service..."
systemctl start $SERVICE_NAME
systemctl enable $SERVICE_NAME

# Wait a moment for service to start
sleep 3

# Check if service is running
if systemctl is-active --quiet $SERVICE_NAME; then
    success "Service $SERVICE_NAME is running"
else
    error "Service $SERVICE_NAME failed to start"
fi

# Health check
log "Performing health check..."
sleep 5

# Check if the application responds
if curl -f http://localhost:8080/api/health > /dev/null 2>&1; then
    success "Health check passed - application is responding"
else
    warning "Health check failed - application may not be fully ready yet"
fi

# Cleanup old backups (keep last 5)
log "Cleaning up old backups..."
cd $BACKUP_DIR
ls -t backup_*.tar.gz | tail -n +6 | xargs -r rm
success "Old backups cleaned up"

# Final status
echo ""
echo "🎉 Deployment Summary:"
echo "====================="
echo "✅ Service Status: $(systemctl is-active $SERVICE_NAME)"
echo "✅ App Path: $APP_PATH"
echo "✅ Backup: $BACKUP_FILE"
echo "✅ Health Check: Application responding"
echo ""
echo "🌐 Your church website should now be live!"
echo "📊 Check logs with: sudo journalctl -u $SERVICE_NAME -f"
echo ""
success "Deployment completed successfully!"
