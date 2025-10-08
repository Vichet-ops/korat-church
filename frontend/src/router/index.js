import { createRouter, createWebHistory } from 'vue-router';

// Views
import HomePage from '../views/HomePage.vue';
import AboutPage from '../views/AboutPage.vue';
import ServicesPage from '../views/ServicesPage.vue';
import EventsPage from '../views/EventsPage.vue';
import GalleryPage from '../views/GalleryPage.vue';
import ContactPage from '../views/ContactPage.vue';
import GivePage from '../views/GivePage.vue';

const routes = [
  { path: '/', name: 'home', component: HomePage },
  { path: '/about', name: 'about', component: AboutPage },
  { path: '/our-ministries', name: 'our-ministries', component: ServicesPage },
  { path: '/events', name: 'events', component: EventsPage },
  { path: '/gallery', name: 'gallery', component: GalleryPage },
  { path: '/contact', name: 'contact', component: ContactPage },
  { path: '/give', name: 'give', component: GivePage },
  // Fallback to home
  { path: '/:pathMatch(.*)*', redirect: '/' },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
  scrollBehavior(to, from, savedPosition) {
    if (savedPosition) {
      return { ...savedPosition, behavior: 'smooth' };
    }
    return { left: 0, top: 0, behavior: 'smooth' };
  },
});

export default router;
