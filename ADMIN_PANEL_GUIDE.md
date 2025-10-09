# 🎨 Admin Panel - Complete Guide

## ✅ Admin Panel is READY!

Your admin panel is now fully built and working at: **http://localhost:3000/admin**

---

## 🚀 How to Access

### **Step 1: Create Your Admin Account**

Run this command (or use the API):

```bash
curl -X POST http://localhost:8081/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "username": "admin",
    "email": "admin@church.com",
    "password": "yourpassword123",
    "first_name": "Your",
    "last_name": "Name"
  }'
```

### **Step 2: Access the Admin Panel**

1. Open your browser
2. Go to: **http://localhost:3000/admin**
3. You'll see the login page
4. Enter your username and password
5. Click "Sign In"
6. You're in! 🎉

---

## 📱 Admin Panel Features

### **1. Login Page** (`/admin/login`)

- Secure authentication
- Beautiful gradient design
- Automatic redirect if already logged in

### **2. Dashboard** (`/admin/dashboard`)

- **Quick Stats:**
  - New Messages count
  - Total Messages count
  - Your admin info
- **Recent Messages:** See latest 5 contact submissions
- **Quick Actions:** Links to view all messages

### **3. Messages Management** (`/admin/messages`)

- **View All Messages:** See all contact form submissions
- **Filter by Status:** All / New / Read
- **Message Details:** Click "View" to see full message
- **Change Status:** Mark as New / Read / Replied
- **Delete Messages:** Remove spam or old messages
- **Reply via Email:** Click to open your email client

---

## 🎯 What You Can Do

### **View Contact Messages**

1. Click "Messages" in sidebar
2. See table of all submissions
3. Filter by status (New, Read, Replied)

### **View Message Details**

1. Click "View" on any message
2. See full details in popup
3. Read name, email, subject, message
4. See timestamp

### **Mark as Read/Replied**

1. Open message details
2. Change status dropdown
3. Select: New / Read / Replied
4. Automatically saves

### **Delete Message**

1. Click "Delete" button
2. Confirm deletion
3. Message removed permanently

### **Reply to Sender**

1. Open message details
2. Click "Reply via Email"
3. Opens your default email client
4. Pre-fills recipient and subject

### **Logout**

1. Click "Logout" in sidebar
2. Confirm
3. Redirected to login page

---

## 🔐 Security Features

✅ **Protected Routes** - Can't access admin pages without login  
✅ **JWT Authentication** - Secure token-based auth  
✅ **Auto Redirect** - Not logged in? Go to login page  
✅ **Token Expiration** - Tokens expire after 24h  
✅ **Hidden from Public** - No links to /admin on public site

---

## 📱 Responsive Design

✅ **Mobile Friendly** - Works on phones and tablets  
✅ **Sidebar Toggle** - Collapsible menu on mobile  
✅ **Touch Optimized** - Easy to use on touchscreens

---

## 🎨 Admin Panel Structure

```
/admin/login           → Login page (public)
/admin                 → Redirects to dashboard
/admin/dashboard       → Main dashboard (protected)
/admin/messages        → Messages management (protected)
```

---

## 💡 Tips & Tricks

### **Stay Logged In**

- Token is stored in browser
- Lasts 24 hours
- Refresh page? Still logged in!

### **View Website**

- Click "View Website →" in top bar
- Opens public site in new tab

### **Quick Navigation**

- Use sidebar menu to switch pages
- Dashboard shows recent messages
- Click "View All" to see everything

### **Keyboard Shortcuts**

- Tab through form fields
- Enter to submit forms
- Escape to close modals

---

## 🔧 Troubleshooting

### **Can't Login?**

- Check username/password
- Make sure admin account is created
- Check backend is running: `docker-compose ps`

### **Not Redirecting After Login?**

- Clear browser cache
- Check browser console for errors
- Verify token is saved: Check localStorage

### **Messages Not Loading?**

- Check backend is running
- Verify API URL in `.env`
- Check browser console for errors

### **Logout Not Working?**

- Hard refresh: Ctrl+Shift+R (Windows) / Cmd+Shift+R (Mac)
- Clear localStorage manually
- Restart browser

---

## 📊 Files Created

**New Admin Components:**

- `frontend/src/views/admin/LoginPage.vue`
- `frontend/src/views/admin/DashboardPage.vue`
- `frontend/src/views/admin/MessagesPage.vue`
- `frontend/src/components/AdminLayout.vue`

**Updated Files:**

- `frontend/src/router/index.js` - Added admin routes

---

## 🎯 What's Next?

You now have:

- ✅ Working admin panel
- ✅ Secure login system
- ✅ Messages management
- ✅ Dashboard with stats
- ✅ Responsive design

### **Optional Enhancements:**

1. Add events management (CRUD)
2. Add settings editor
3. Add gallery upload
4. Add multiple admin users
5. Add activity logs

---

## 🚀 Production Deployment

### **Before Going Live:**

1. **Change JWT Secret**

   ```bash
   # In .env, set a strong secret
   JWT_SECRET=your-super-secret-random-string-here
   ```

2. **Update API URL**

   ```bash
   # In frontend/.env
   VITE_API_URL=https://your-api-domain.com
   ```

3. **Test Everything**

   - Login/logout
   - View messages
   - Update status
   - Delete messages

4. **Deploy!**
   - Public site: https://yourchurch.com
   - Admin panel: https://yourchurch.com/admin

---

## 📚 Quick Reference

### **Login Credentials**

- Create via API or registration endpoint
- Username + Password required
- First admin must be created via API

### **Access URLs**

- **Login:** http://localhost:3000/admin/login
- **Dashboard:** http://localhost:3000/admin/dashboard
- **Messages:** http://localhost:3000/admin/messages

### **API Endpoints Used**

- `POST /api/auth/login` - Authentication
- `GET /api/admin/messages` - Get all messages
- `PATCH /api/admin/messages/:id/status` - Update status
- `DELETE /api/admin/messages/:id` - Delete message

---

## 🎉 You're All Set!

Your admin panel is **100% functional**.

You can now:

- ✅ Login securely
- ✅ View contact messages
- ✅ Manage message status
- ✅ Reply to senders
- ✅ Delete spam

**No more TablePlus needed for managing messages!** 🚀

---

## 📞 Need Help?

Check the following files for reference:

- `CRITICAL_FIXES_COMPLETED.md` - Backend API docs
- `QUICK_START_GUIDE.md` - Setup instructions
- `README.md` - Project overview

Enjoy your new admin panel! 🎨✨
