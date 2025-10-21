<template>
  <div id="app" class="min-h-screen bg-gray-50">
    <!-- Navigation (hidden on admin pages) -->
    <NavbarComponent
      v-if="!isAdminRoute"
      :church-name="displayChurchName"
      :nav-pages="navPages"
      :current-page="currentPage"
      :mobile-menu-open="mobileMenuOpen"
      @toggle-mobile-menu="mobileMenuOpen = !mobileMenuOpen"
    />

    <!-- Main Content routed -->
    <main class="pt-0">
      <router-view v-slot="{ Component }">
        <component
          :is="Component"
          :events="events"
          :churchSettings="churchSettings"
          :serviceTimes="serviceTimes"
        />
      </router-view>
    </main>

    <!-- Footer (hidden on admin pages) -->
    <FooterComponent
      v-if="!isAdminRoute"
      :church-name="displayChurchName"
      :church-settings="churchSettings"
    />
  </div>
</template>

<script>
import axios from 'axios';
// Layout Components only (routes import their own views)
import NavbarComponent from './components/NavbarComponent.vue';
import FooterComponent from './components/FooterComponent.vue';

const API_URL = import.meta.env.VITE_API_URL || 'https://vichetkeo.com';

export default {
  name: 'App',
  components: {
    NavbarComponent,
    FooterComponent,
  },
  data() {
    return {
      currentPage: 'home',
      mobileMenuOpen: false,
      churchName: 'Muang Thai Korat Church',
      churchSettings: {
        church_name: '',
        church_address: '',
        church_phone: '',
        church_email: '',
      },
      events: [],
      serviceTimes: [
        { day: 'Sunday', time: '10:00 AM', service: 'Main Worship Service' },
        { day: 'Wednesday', time: '7:00 PM', service: 'Bible Study & Prayer' },
        {
          day: 'Friday',
          time: '6:30 PM',
          service: 'Youth Activities & Fellowship',
        },
      ],
      navPages: [
        { id: 'home', name: 'Home' },
        { id: 'about', name: 'About' },
        { id: 'our-ministries', name: 'Our Ministries' },
        { id: 'events', name: 'Events' },
        { id: 'contact', name: 'Contact' },
      ],
    };
  },
  computed: {
    displayChurchName() {
      return (
        this.churchSettings.church_name ||
        this.churchName ||
        'Muang Thai Korat Church'
      );
    },
    isAdminRoute() {
      return this.$route.path.startsWith('/admin');
    },
  },
  methods: {
    async loadPageData() {
      try {
        const page =
          this.$route && this.$route.name ? this.$route.name : 'home';
        this.currentPage = page;
        const response = await axios.get(`${API_URL}/api/${page}`);
        if (page === 'home') {
          this.events = response.data.events || [];
          this.churchSettings = response.data.settings || {};
          this.churchName = this.churchSettings.church_name || this.churchName;
          this.parseServiceTimes();
        } else if (page === 'about') {
          this.churchSettings = response.data.settings || {};
          this.churchName = this.churchSettings.church_name || this.churchName;
          this.parseServiceTimes();
        } else if (page === 'events') {
          this.events = response.data.events || [];
          this.churchSettings = response.data.settings || {};
          this.churchName = this.churchSettings.church_name || this.churchName;
        } else if (page === 'contact') {
          this.churchSettings = response.data.settings || {};
        } else {
          // For all other pages (gallery, give, our-ministries), load settings
          if (response.data.settings) {
            this.churchSettings = response.data.settings || {};
            this.churchName =
              this.churchSettings.church_name || this.churchName;
          }
        }
      } catch (error) {
        // Handle error silently for production
      }
    },
    parseServiceTimes() {
      try {
        if (this.churchSettings.service_times) {
          this.serviceTimes = JSON.parse(this.churchSettings.service_times);
        }
      } catch (error) {
        this.serviceTimes = [];
      }
    },
  },
  watch: {
    '$route.name'() {
      this.loadPageData();
    },
  },
  async mounted() {
    await this.loadPageData();
  },
};
</script>

<style>
@import url('https://fonts.googleapis.com/css2?family=Inter:wght@300;400;500;600;700;800;900&display=swap');

* {
  font-family: 'Inter', sans-serif;
}

html {
  scroll-behavior: smooth;
}
</style>
