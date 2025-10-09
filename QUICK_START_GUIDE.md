# 🚀 Quick Start Guide - New Features

## What's New? ✨

Your church website now has:

1. ✅ **Working Contact Form** - Saves to database + sends emails
2. ✅ **Admin Authentication** - Secure login with JWT
3. ✅ **Contact Message Management** - View & manage form submissions

---

## 🎯 Quick Setup (3 Steps)

### Step 1: Create Your First Admin Account

Open your browser and use this tool or terminal:

```bash
curl -X POST http://localhost:8081/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "username": "admin",
    "email": "your-email@church.com",
    "password": "YourSecurePassword123",
    "first_name": "Your",
    "last_name": "Name"
  }'
```

**Response:**

```json
{
  "message": "Admin account created successfully",
  "token": "eyJhbGc...",
  "admin": {
    "id": 1,
    "username": "admin",
    "email": "your-email@church.com"
  }
}
```

💾 **Save the token** - you'll need it to access admin features!

---

### Step 2: Test the Contact Form

Go to your website: **http://localhost:3000/contact**

Fill out the form and submit. You should see:

- ✅ Success message
- ✅ Message saved to database
- ✅ Console log in backend (check: `docker-compose logs backend`)

---

### Step 3: View Contact Messages

Use your JWT token from Step 1:

```bash
curl -X GET http://localhost:8081/api/admin/messages \
  -H "Authorization: Bearer YOUR_TOKEN_HERE"
```

**Response:**

```json
{
  "count": 2,
  "messages": [
    {
      "id": 1,
      "name": "John Doe",
      "email": "john@example.com",
      "subject": "Question about services",
      "message": "What time is Sunday service?",
      "status": "new",
      "created_at": "2025-10-09T..."
    }
  ]
}
```

---

## 📧 Email Setup (Optional)

Want to receive email notifications? Configure SMTP in `.env`:

### Option 1: Gmail

1. Enable 2-Factor Authentication in your Google account
2. Generate an App Password: https://myaccount.google.com/apppasswords
3. Update `.env`:

```bash
SMTP_HOST=smtp.gmail.com
SMTP_PORT=587
SMTP_USER=your-email@gmail.com
SMTP_PASS=your-16-digit-app-password
SMTP_FROM=noreply@yourchurch.com
ADMIN_EMAIL=admin@yourchurch.com
```

### Option 2: SendGrid (Recommended for Production)

1. Sign up: https://sendgrid.com (Free tier: 100 emails/day)
2. Create API Key
3. Update `.env`:

```bash
SMTP_HOST=smtp.sendgrid.net
SMTP_PORT=587
SMTP_USER=apikey
SMTP_PASS=YOUR_SENDGRID_API_KEY
SMTP_FROM=noreply@yourchurch.com
ADMIN_EMAIL=admin@yourchurch.com
```

4. Restart backend:

```bash
docker-compose restart backend
```

---

## 🔐 All Available API Endpoints

### Public Endpoints (No Auth Required)

| Method | Endpoint             | Description          |
| ------ | -------------------- | -------------------- |
| POST   | `/api/contact/send`  | Submit contact form  |
| POST   | `/api/auth/login`    | Login to admin panel |
| POST   | `/api/auth/register` | Register new admin   |

### Protected Endpoints (JWT Required)

| Method | Endpoint                         | Description              |
| ------ | -------------------------------- | ------------------------ |
| GET    | `/api/admin/profile`             | Get your profile         |
| POST   | `/api/admin/change-password`     | Change password          |
| GET    | `/api/admin/messages`            | Get all contact messages |
| PATCH  | `/api/admin/messages/:id/status` | Update message status    |
| DELETE | `/api/admin/messages/:id`        | Delete message           |

---

## 💻 Using the API

### Login and Get Token

```bash
curl -X POST http://localhost:8081/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{"username":"admin","password":"YourPassword123"}'
```

### Use Token for Protected Endpoints

```bash
# Get all messages
curl -X GET http://localhost:8081/api/admin/messages \
  -H "Authorization: Bearer YOUR_TOKEN"

# Update message status
curl -X PATCH http://localhost:8081/api/admin/messages/1/status \
  -H "Authorization: Bearer YOUR_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"status":"read"}'

# Delete message
curl -X DELETE http://localhost:8081/api/admin/messages/1 \
  -H "Authorization: Bearer YOUR_TOKEN"
```

---

## 🛠️ Common Tasks

### Change Your Password

```bash
curl -X POST http://localhost:8081/api/admin/change-password \
  -H "Authorization: Bearer YOUR_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "current_password": "OldPassword123",
    "new_password": "NewSecurePassword456"
  }'
```

### Filter Messages by Status

```bash
# Get only new messages
curl -X GET "http://localhost:8081/api/admin/messages?status=new" \
  -H "Authorization: Bearer YOUR_TOKEN"

# Available statuses: new, read, replied
```

---

## 🐛 Troubleshooting

### "Invalid or expired token"

- Your JWT token expired (default: 24h)
- Solution: Login again to get a new token

### "Authorization header required"

- You forgot to include the Bearer token
- Solution: Add `-H "Authorization: Bearer YOUR_TOKEN"`

### "Email not configured"

- SMTP settings are empty (that's OK!)
- Messages still save to database
- Check backend logs for submissions

### Contact form not working

- Check frontend console for errors
- Verify backend is running: `docker-compose ps`
- Test API directly with curl

---

## 📱 Next Steps

### Build Admin Dashboard (Recommended)

Create a simple admin panel in Vue.js:

1. **Login Page** - Use `/api/auth/login`
2. **Messages Page** - Display contact submissions
3. **Message Detail** - View and update status

### Example Vue Component:

```vue
<script setup>
import { ref, onMounted } from 'vue';

const messages = ref([]);
const token = ref(localStorage.getItem('admin_token'));

async function fetchMessages() {
  const response = await fetch('http://localhost:8081/api/admin/messages', {
    headers: {
      Authorization: `Bearer ${token.value}`,
    },
  });
  const data = await response.json();
  messages.value = data.messages;
}

onMounted(() => {
  fetchMessages();
});
</script>

<template>
  <div>
    <h1>Contact Messages</h1>
    <div v-for="msg in messages" :key="msg.id">
      <h3>{{ msg.subject }}</h3>
      <p>From: {{ msg.name }} ({{ msg.email }})</p>
      <p>{{ msg.message }}</p>
      <span>Status: {{ msg.status }}</span>
    </div>
  </div>
</template>
```

---

## 📚 Additional Resources

- **Backend API:** http://localhost:8081
- **Frontend:** http://localhost:3000
- **Health Check:** http://localhost:8081/api/health
- **Environment Variables:** See `env.example`
- **Full Documentation:** See `CRITICAL_FIXES_COMPLETED.md`

---

## 🎉 You're All Set!

Your church website now has:

- ✅ Functional contact form
- ✅ Secure admin authentication
- ✅ Message management API
- ✅ Email notification system (when configured)

**Ready for production deployment!**

Just need to:

1. Configure SMTP for email sending
2. Build an admin dashboard UI (optional)
3. Set strong JWT_SECRET in production
4. Deploy to your hosting platform

Need help? Check the full documentation or ask for assistance!
