<template>
  <div id="app" class="min-h-screen bg-gray-50">
    <!-- Navigation -->
    <NavbarComponent
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

    <!-- Footer -->
    <FooterComponent
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

const API_URL = import.meta.env.VITE_API_URL || 'http://localhost:8081';

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
      serviceTimes: [],
      navPages: [
        { id: 'home', name: 'Home' },
        { id: 'about', name: 'About' },
        { id: 'services', name: 'Services' },
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
        } else if (page === 'contact') {
          this.churchSettings = response.data.settings || {};
        }
      } catch (error) {
        console.error(`Error loading ${this.currentPage} data:`, error);
      }
    },
    parseServiceTimes() {
      try {
        if (this.churchSettings.service_times) {
          this.serviceTimes = JSON.parse(this.churchSettings.service_times);
        }
      } catch (error) {
        console.error('Error parsing service times:', error);
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
