import { createRouter, createWebHistory } from 'vue-router';

// Public Views
import HomePage from '../views/HomePage.vue';
import AboutPage from '../views/AboutPage.vue';
import MinistriesPage from '../views/MinistriesPage.vue';
import EventsPage from '../views/EventsPage.vue';
import GalleryPage from '../views/GalleryPage.vue';
import ContactPage from '../views/ContactPage.vue';
import GivePage from '../views/GivePage.vue';

// Admin Views
import AdminLayout from '../components/AdminLayout.vue';
import LoginPage from '../views/admin/LoginPage.vue';
import DashboardPage from '../views/admin/DashboardPage.vue';
import MessagesPage from '../views/admin/MessagesPage.vue';

const routes = [
  // Public Routes
  { path: '/', name: 'home', component: HomePage },
  { path: '/about', name: 'about', component: AboutPage },
  {
    path: '/our-ministries',
    name: 'our-ministries',
    component: MinistriesPage,
  },
  { path: '/events', name: 'events', component: EventsPage },
  { path: '/gallery', name: 'gallery', component: GalleryPage },
  { path: '/contact', name: 'contact', component: ContactPage },
  { path: '/give', name: 'give', component: GivePage },

  // Admin Routes
  {
    path: '/admin/login',
    name: 'admin-login',
    component: LoginPage,
  },
  {
    path: '/admin',
    component: AdminLayout,
    meta: { requiresAuth: true },
    children: [
      {
        path: '',
        redirect: '/admin/dashboard',
      },
      {
        path: 'dashboard',
        name: 'admin-dashboard',
        component: DashboardPage,
      },
      {
        path: 'messages',
        name: 'admin-messages',
        component: MessagesPage,
      },
    ],
  },

  // Fallback to home
  { path: '/:pathMatch(.*)*', redirect: '/' },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
  scrollBehavior(to, from, savedPosition) {
    // Always scroll to top on page refresh or navigation
    return { left: 0, top: 0, behavior: 'smooth' };
  },
});

// Navigation guard for admin routes
router.beforeEach((to, from, next) => {
  if (to.matched.some((record) => record.meta.requiresAuth)) {
    const token = localStorage.getItem('admin_token');
    if (!token) {
      next('/admin/login');
    } else {
      next();
    }
  } else {
    next();
  }
});

export default router;
