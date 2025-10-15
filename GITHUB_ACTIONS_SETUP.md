# 🚀 GitHub Actions Auto Deploy Setup

This guide will help you set up automatic deployment for your church website using GitHub Actions.

## 📋 Prerequisites

Before setting up auto-deploy, you need:

1. ✅ **GitHub Repository** (already done)
2. ✅ **Production Server** (VPS, AWS, DigitalOcean, etc.)
3. ✅ **Domain Name** (optional but recommended)
4. ✅ **SSH Access** to your server

## 🔧 Setup Steps

### Step 1: Configure GitHub Secrets

Go to your GitHub repository → **Settings** → **Secrets and variables** → **Actions**

Add these secrets:

```
DEPLOY_HOST=your-server-ip-or-domain.com
DEPLOY_USER=your-username
DEPLOY_KEY=your-private-ssh-key
DEPLOY_PATH=/var/www/church-website
```

#### How to get these values:

**DEPLOY_HOST:**

- Your server's IP address or domain name
- Example: `123.456.789.0` or `church.yourdomain.com`

**DEPLOY_USER:**

- Username for SSH access to your server
- Example: `ubuntu`, `root`, or `deploy`

**DEPLOY_KEY:**

- Private SSH key for server access
- Generate with: `ssh-keygen -t rsa -b 4096`
- Copy the **private** key content (not public)

**DEPLOY_PATH:**

- Directory on server where app will be deployed
- Example: `/var/www/church-website` or `/home/user/church-app`

### Step 2: Prepare Your Server

#### Install Required Software:

```bash
# Update system
sudo apt update && sudo apt upgrade -y

# Install Node.js
curl -fsSL https://deb.nodesource.com/setup_18.x | sudo -E bash -
sudo apt-get install -y nodejs

# Install Go
wget https://go.dev/dl/go1.21.6.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.21.6.linux-amd64.tar.gz
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
source ~/.bashrc

# Install PostgreSQL
sudo apt install postgresql postgresql-contrib -y
sudo systemctl start postgresql
sudo systemctl enable postgresql

# Install Nginx (for serving the app)
sudo apt install nginx -y
```

#### Setup Database:

```bash
# Create database user
sudo -u postgres psql
CREATE USER church_user WITH PASSWORD 'your_secure_password';
CREATE DATABASE church_db OWNER church_user;
GRANT ALL PRIVILEGES ON DATABASE church_db TO church_user;
\q
```

#### Setup Nginx:

```nginx
# /etc/nginx/sites-available/church-website
server {
    listen 80;
    server_name your-domain.com www.your-domain.com;

    # Serve frontend files
    location / {
        root /var/www/church-website/frontend/dist;
        try_files $uri $uri/ /index.html;
    }

    # Proxy API requests to Go backend
    location /api/ {
        proxy_pass http://localhost:8080;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
    }
}
```

### Step 3: Create Deployment Script

Create a deployment script on your server:

```bash
# /var/www/church-website/deploy.sh
#!/bin/bash

echo "🚀 Starting deployment..."

# Stop the application
sudo systemctl stop church-website

# Backup current version
cp -r /var/www/church-website /var/www/church-website.backup.$(date +%Y%m%d_%H%M%S)

# Extract new version
tar -xzf church-app.tar.gz

# Install dependencies and build
cd /var/www/church-website
go mod download
cd frontend && npm install && npm run build
cd ..

# Set permissions
chmod +x church-app
chown -R www-data:www-data /var/www/church-website

# Start the application
sudo systemctl start church-website

echo "✅ Deployment completed!"
```

### Step 4: Create System Service

Create a systemd service file:

```ini
# /etc/systemd/system/church-website.service
[Unit]
Description=Church Website
After=network.target

[Service]
Type=simple
User=www-data
WorkingDirectory=/var/www/church-website
ExecStart=/var/www/church-website/church-app
Restart=always
Environment=DATABASE_URL=postgres://church_user:your_secure_password@localhost:5432/church_db?sslmode=disable
Environment=JWT_SECRET=your_jwt_secret_here
Environment=GIN_MODE=release

[Install]
WantedBy=multi-user.target
```

Enable the service:

```bash
sudo systemctl daemon-reload
sudo systemctl enable church-website
sudo systemctl start church-website
```

## 🎯 How It Works

### Automatic Deployment Flow:

1. ✅ **Push to main branch** → Triggers GitHub Actions
2. ✅ **Run tests** → Ensures code quality
3. ✅ **Build application** → Creates production files
4. ✅ **Deploy to server** → Uploads and restarts service
5. ✅ **Notify success** → Confirms deployment

### Manual Deployment:

You can also trigger deployment manually:

1. Go to **Actions** tab in GitHub
2. Select **Deploy Church Website** workflow
3. Click **Run workflow**
4. Choose branch and click **Run workflow**

## 🔍 Monitoring Deployments

### Check Deployment Status:

- **GitHub Actions tab** → View all deployments
- **Server logs** → `sudo journalctl -u church-website -f`
- **Website** → Visit your domain to verify

### Common Issues:

**Deployment Fails:**

- Check GitHub Actions logs
- Verify server SSH access
- Ensure all secrets are correct

**Website Not Loading:**

- Check if service is running: `sudo systemctl status church-website`
- Check Nginx status: `sudo systemctl status nginx`
- Verify database connection

**SSL Certificate (Optional):**

```bash
# Install Certbot for free SSL
sudo apt install certbot python3-certbot-nginx -y
sudo certbot --nginx -d your-domain.com
```

## 🎉 Success!

Once setup is complete:

- ✅ **Push code** → Auto-deploys to production
- ✅ **Website updates** → Live in minutes
- ✅ **No manual work** → Fully automated
- ✅ **Professional setup** → Production-ready

Your church website now has **professional-grade auto-deployment**! 🚀

## 📞 Support

If you need help with setup:

1. Check GitHub Actions logs
2. Verify server configuration
3. Test manual deployment first
4. Contact your server provider if needed

**Happy deploying!** ✨
