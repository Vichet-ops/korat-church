<template>
  <div id="app" class="min-h-screen bg-gray-50">
    <!-- Navigation -->
    <nav
      class="fixed top-0 left-0 w-full z-50 bg-gradient-to-b from-[#0b1f3a] to-[#1b2f58] shadow-md pt-4 pb-4"
    >
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="flex justify-between items-center h-20 md:h-24 lg:h-28">
          <!-- Logo -->
          <div class="flex items-center">
            <a
              href="#"
              @click.prevent="currentPage = 'home'"
              class="hover:opacity-80 transition duration-300 flex items-center"
            >
              <img
                src="/images/cross-logo.png"
                alt="Muang Thai Korat Church Logo"
                class="h-12 md:h-20 lg:h-24 w-auto"
              />
              <span
                class="ml-3 text-white font-semibold text-base md:text-lg"
                >{{ churchName }}</span
              >
            </a>
          </div>

          <!-- Desktop Navigation -->
          <div class="hidden md:flex items-center space-x-4 ml-16 lg:ml-20">
            <a
              v-for="page in navPages"
              :key="page.id"
              href="#"
              @click.prevent="currentPage = page.id"
              :class="navClass(page.id)"
              class="px-4 py-3 text-base font-medium transition duration-300 relative group"
            >
              <span class="relative inline-block">
                {{ page.name }}
                <span
                  :class="
                    currentPage === page.id
                      ? 'w-full'
                      : 'w-0 group-hover:w-full'
                  "
                  class="absolute -bottom-0.5 left-0 h-0.5 bg-white transition-all duration-300 ease-out"
                ></span>
              </span>
            </a>
            <a
              href="#"
              @click.prevent="currentPage = 'give'"
              class="bg-blue-600 text-white px-6 py-3 rounded-md text-base font-medium hover:bg-blue-700 transition duration-300 shadow-md"
            >
              Give
            </a>
          </div>

          <!-- Mobile menu button -->
          <div class="md:hidden">
            <button
              @click="mobileMenuOpen = !mobileMenuOpen"
              class="text-white hover:text-blue-200 focus:outline-none"
            >
              <svg
                class="h-6 w-6"
                fill="none"
                viewBox="0 0 24 24"
                stroke="currentColor"
              >
                <path
                  v-if="!mobileMenuOpen"
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M4 6h16M4 12h16M4 18h16"
                />
                <path
                  v-else
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M6 18L18 6M6 6l12 12"
                />
              </svg>
            </button>
          </div>
        </div>

        <!-- Mobile Navigation Menu -->
        <div
          v-if="mobileMenuOpen"
          class="md:hidden px-2 pt-2 pb-3 space-y-1 bg-[#0b1f3a] rounded-lg mt-4 shadow-lg border border-[#1b2f58]"
        >
          <a
            v-for="page in navPages"
            :key="page.id"
            href="#"
            @click.prevent="
              currentPage = page.id;
              mobileMenuOpen = false;
            "
            :class="currentPage === page.id ? 'text-blue-200' : 'text-white'"
            class="hover:text-blue-200 block px-3 py-2 text-base font-medium rounded-md transition duration-300"
          >
            {{ page.name }}
          </a>
          <a
            href="#"
            @click.prevent="
              currentPage = 'give';
              mobileMenuOpen = false;
            "
            class="bg-blue-600 text-white block px-3 py-2 text-base font-medium rounded-md hover:bg-blue-700 transition duration-300"
          >
            Give
          </a>
        </div>
      </div>
    </nav>

    <!-- Main Content -->
    <main class="pt-20 md:pt-24 lg:pt-28">
      <!-- Home Page -->
      <div v-if="currentPage === 'home'">
        <HomePage
          :events="events"
          :churchSettings="churchSettings"
          :serviceTimes="serviceTimes"
          @navigate="currentPage = $event"
        />
      </div>

      <!-- About Page -->
      <div v-else-if="currentPage === 'about'">
        <AboutPage
          :churchSettings="churchSettings"
          :serviceTimes="serviceTimes"
          @navigate="currentPage = $event"
        />
      </div>

      <!-- Services Page -->
      <div v-else-if="currentPage === 'services'">
        <ServicesPage @navigate="currentPage = $event" />
      </div>

      <!-- Events Page -->
      <div v-else-if="currentPage === 'events'">
        <EventsPage :events="events" @navigate="currentPage = $event" />
      </div>

      <!-- Gallery Page -->
      <div v-else-if="currentPage === 'gallery'">
        <GalleryPage />
      </div>

      <!-- Contact Page -->
      <div v-else-if="currentPage === 'contact'">
        <ContactPage
          :churchSettings="churchSettings"
          @navigate="currentPage = $event"
        />
      </div>

      <!-- Give Page -->
      <div v-else-if="currentPage === 'give'">
        <GivePage @navigate="currentPage = $event" />
      </div>
    </main>

    <!-- Footer -->
    <footer class="bg-gray-900 text-white py-12">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="grid grid-cols-1 md:grid-cols-4 gap-8">
          <!-- Church Info -->
          <div class="col-span-1 md:col-span-2">
            <div class="flex items-center mb-4">
              <a
                href="#"
                @click.prevent="currentPage = 'home'"
                class="flex items-center hover:opacity-80 transition duration-300"
              >
                <img
                  src="/images/cross-logo.png"
                  alt="Muang Thai Korat Church Logo"
                  class="h-12 w-auto mr-3"
                />
                <h3 class="text-lg font-semibold text-white">
                  {{ churchName }}
                </h3>
              </a>
            </div>
          </div>

          <!-- Quick Links -->
          <div>
            <h4 class="text-lg font-semibold text-white mb-4">Quick Links</h4>
            <ul class="space-y-2">
              <li>
                <a
                  href="#"
                  @click.prevent="currentPage = 'about'"
                  class="text-gray-300 hover:text-white transition duration-300"
                  >About Us</a
                >
              </li>
              <li>
                <a
                  href="#"
                  @click.prevent="currentPage = 'services'"
                  class="text-gray-300 hover:text-white transition duration-300"
                  >Services</a
                >
              </li>
              <li>
                <a
                  href="#"
                  @click.prevent="currentPage = 'events'"
                  class="text-gray-300 hover:text-white transition duration-300"
                  >Events</a
                >
              </li>
              <li>
                <a
                  href="#"
                  @click.prevent="currentPage = 'contact'"
                  class="text-gray-300 hover:text-white transition duration-300"
                  >Contact</a
                >
              </li>
              <li>
                <a
                  href="#"
                  @click.prevent="currentPage = 'give'"
                  class="text-gray-300 hover:text-white transition duration-300"
                  >Give</a
                >
              </li>
            </ul>
          </div>

          <!-- Contact Info -->
          <div>
            <h4 class="text-lg font-semibold text-white mb-4">Contact Info</h4>
            <div class="space-y-2 text-gray-300">
              <div class="flex items-start">
                <svg
                  class="w-5 h-5 mr-2 mt-1 flex-shrink-0"
                  fill="currentColor"
                  viewBox="0 0 20 20"
                >
                  <path
                    fill-rule="evenodd"
                    d="M5.05 4.05a7 7 0 119.9 9.9L10 18.9l-4.95-4.95a7 7 0 010-9.9zM10 11a2 2 0 100-4 2 2 0 000 4z"
                    clip-rule="evenodd"
                  />
                </svg>
                <span>{{ churchSettings.church_address }}</span>
              </div>
              <div class="flex items-center">
                <svg
                  class="w-5 h-5 mr-2 flex-shrink-0"
                  fill="currentColor"
                  viewBox="0 0 20 20"
                >
                  <path
                    d="M2.003 5.884L10 9.882l7.997-3.998A2 2 0 0016 4H4a2 2 0 00-1.997 1.884z"
                  />
                  <path
                    d="M18 8.118l-8 4-8-4V14a2 2 0 002 2h12a2 2 0 002-2V8.118z"
                  />
                </svg>
                <span>{{ churchSettings.church_email }}</span>
              </div>
              <div class="flex items-center">
                <svg
                  class="w-5 h-5 mr-2 flex-shrink-0"
                  fill="currentColor"
                  viewBox="0 0 20 20"
                >
                  <path
                    d="M2 3a1 1 0 011-1h2.153a1 1 0 01.986.836l.74 4.435a1 1 0 01-.54 1.06l-1.548.773a11.037 11.037 0 006.105 6.105l.774-1.548a1 1 0 011.059-.54l4.435.74a1 1 0 01.836.986V17a1 1 0 01-1 1h-2C7.82 18 2 12.18 2 5V3z"
                  />
                </svg>
                <span>{{ churchSettings.church_phone }}</span>
              </div>
            </div>
          </div>
        </div>

        <!-- Bottom Bar -->
        <div class="border-t border-gray-800 mt-8 pt-8 text-center">
          <p class="text-gray-400">
            © {{ currentYear }} {{ churchName }}. All rights reserved.
          </p>
        </div>
      </div>
    </footer>
  </div>
</template>

<script>
import axios from 'axios';
import HomePage from './components/HomePage.vue';
import AboutPage from './components/AboutPage.vue';
import ServicesPage from './components/ServicesPage.vue';
import EventsPage from './components/EventsPage.vue';
import GalleryPage from './components/GalleryPage.vue';
import ContactPage from './components/ContactPage.vue';
import GivePage from './components/GivePage.vue';

const API_URL = import.meta.env.VITE_API_URL || 'http://localhost:8081';

export default {
  name: 'App',
  components: {
    HomePage,
    AboutPage,
    ServicesPage,
    EventsPage,
    GalleryPage,
    ContactPage,
    GivePage,
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
        { id: 'gallery', name: 'Gallery' },
        { id: 'contact', name: 'Contact' },
      ],
    };
  },
  computed: {
    currentYear() {
      return new Date().getFullYear();
    },
  },
  methods: {
    navClass(page) {
      return this.currentPage === page ? 'text-blue-200' : 'text-white';
    },
    async loadPageData() {
      try {
        const response = await axios.get(`${API_URL}/api/${this.currentPage}`);
        if (this.currentPage === 'home') {
          this.events = response.data.events || [];
          this.churchSettings = response.data.settings || {};
          this.churchName = this.churchSettings.church_name || 'Church';
          this.parseServiceTimes();
        } else if (this.currentPage === 'about') {
          this.churchSettings = response.data.settings || {};
          this.churchName = this.churchSettings.church_name || 'Church';
          this.parseServiceTimes();
        } else if (this.currentPage === 'events') {
          this.events = response.data.events || [];
        } else if (this.currentPage === 'contact') {
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
    currentPage() {
      this.loadPageData();
      this.mobileMenuOpen = false;
      window.scrollTo({ top: 0, behavior: 'smooth' });
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
