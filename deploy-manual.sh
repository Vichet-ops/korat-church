#!/bin/bash

# 🚀 Manual Deployment Script for Church Website
# This script deploys your church website to production

set -e  # Exit on any error

echo "🚀 Starting manual deployment to vichetkeo.com..."

# Configuration
SERVER_IP="167.172.210.140"
SERVER_USER="root"
SSH_KEY="~/.ssh/vichet_server"
APP_DIR="/var/www/church-website"
WEB_DIR="/var/www/vichetkeo.com/html"

# Colors for output
GREEN='\033[0;32m'
BLUE='\033[0;34m'
RED='\033[0;31m'
NC='\033[0m' # No Color

log() {
    echo -e "${BLUE}[$(date +'%H:%M:%S')]${NC} $1"
}

success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

error() {
    echo -e "${RED}[ERROR]${NC} $1"
    exit 1
}

# Step 1: Build frontend
log "Building frontend..."
cd frontend
npm run build
cd ..

# Step 2: Build Go application for Linux
log "Building Go application for Linux..."
GOOS=linux GOARCH=amd64 go build -o church-app main.go

# Step 3: Create deployment package
log "Creating deployment package..."
tar -czf church-app.tar.gz church-app frontend/dist

# Step 4: Upload to server
log "Uploading to server..."
scp -i $SSH_KEY church-app.tar.gz $SERVER_USER@$SERVER_IP:/tmp/

# Step 5: Deploy on server
log "Deploying on server..."
ssh -i $SSH_KEY $SERVER_USER@$SERVER_IP "
    echo '📦 Extracting deployment package...'
    cd $APP_DIR
    tar -xzf /tmp/church-app.tar.gz
    
    echo '🌐 Updating website files...'
    cp -r frontend/dist/* $WEB_DIR/
    
    echo '🔧 Starting backend service...'
    pkill -f church-app || true
    sleep 2
    
    chmod +x church-app
    DATABASE_URL='postgres://church_user:church123@localhost:5432/church_db?sslmode=disable' JWT_SECRET='your-jwt-secret-key-here' GIN_MODE=release nohup ./church-app > church-app.log 2>&1 &
    
    echo '⏳ Waiting for service to start...'
    sleep 5
    
    if curl -f http://localhost:8080/api/health > /dev/null 2>&1; then
        echo '✅ Backend service is running!'
    else
        echo '⚠️ Backend service may not be running'
    fi
    
    echo '✅ Deployment completed!'
"

# Step 6: Cleanup
log "Cleaning up..."
rm church-app.tar.gz
rm church-app

success "🎉 Deployment completed successfully!"
success "🌐 Your website is live at: https://vichetkeo.com"
success "🔧 Admin panel: https://vichetkeo.com/admin/login"
success "📊 Username: admin, Password: admin123"
