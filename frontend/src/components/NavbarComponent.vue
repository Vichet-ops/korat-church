<template>
  <nav
    class="fixed top-0 left-0 w-full z-50 bg-[#0f2744] shadow-lg border-b border-white/10 transition-all duration-300"
  >
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
      <div
        class="flex justify-between items-center h-28 lg:h-32 transition-all duration-300"
      >
        <!-- Logo -->
        <div class="flex items-center">
          <a
            href="/"
            @click.prevent="onLogoClick"
            class="hover:opacity-80 transition duration-300 flex items-center"
          >
            <img
              src="/images/cross-logo.png"
              alt="Muang Thai Korat Church Logo"
              class="h-16 lg:h-28 w-auto mr-2"
            />
            <span
              class="text-white font-semibold text-lg lg:text-2xl whitespace-nowrap"
            >
              {{ churchName }}
            </span>
          </a>
        </div>

        <!-- Desktop Navigation -->
        <div class="hidden md:flex items-center space-x-2 ml-12">
          <router-link
            v-for="page in navPages"
            :key="page.id"
            :to="getRoutePath(page.id)"
            :class="[
              'px-3 py-3 text-base font-medium transition duration-300 relative group',
              $route.path === getRoutePath(page.id)
                ? 'text-blue-400'
                : 'text-white hover:text-blue-400',
            ]"
          >
            <span class="relative inline-block">
              {{ page.name }}
              <span
                :class="
                  $route.path === getRoutePath(page.id)
                    ? 'w-full'
                    : 'w-0 group-hover:w-full'
                "
                class="absolute -bottom-0.5 left-0 h-0.5 bg-blue-400 transition-all duration-300 ease-out"
              ></span>
            </span>
          </router-link>
          <router-link
            to="/give"
            class="bg-blue-600 text-white px-4 py-2.5 rounded-md text-base font-medium hover:bg-blue-700 transition duration-300 shadow-md ml-2"
          >
            Give
          </router-link>
        </div>

        <!-- Mobile menu button -->
        <div class="md:hidden">
          <button
            @click="$emit('toggle-mobile-menu')"
            class="text-white hover:text-blue-200 focus:outline-none transition-colors duration-300 p-2 rounded-md hover:bg-white/10"
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

      <!-- Backdrop to close menu on outside click -->
      <div
        v-if="mobileMenuOpen"
        class="fixed inset-0 md:hidden z-40"
        @click="$emit('toggle-mobile-menu')"
      ></div>

      <!-- Mobile Navigation Menu -->
      <div
        v-if="mobileMenuOpen"
        class="md:hidden absolute left-0 right-0 px-2 pt-2 pb-3 space-y-1 bg-[#0f2744] rounded-lg mt-4 shadow-lg border border-[#1a3a5c] transition-all duration-300 z-50"
      >
        <router-link
          v-for="page in navPages"
          :key="page.id"
          :to="getRoutePath(page.id)"
          :class="mobileNavClass(page.id)"
          class="hover:text-blue-400 block px-3 py-2 text-base font-medium rounded-md transition duration-300 relative group"
          @click="$emit('toggle-mobile-menu')"
        >
          <span class="relative inline-block">
            {{ page.name }}
            <span
              :class="
                currentPage === page.id ? 'w-full' : 'w-0 group-hover:w-full'
              "
              class="absolute -bottom-0.5 left-0 h-0.5 bg-blue-400 transition-all duration-300 ease-out"
            ></span>
          </span>
        </router-link>
        <router-link
          to="/give"
          class="bg-blue-600 text-white block px-3 py-2 text-base font-medium rounded-md hover:bg-blue-700 transition duration-300"
          @click="$emit('toggle-mobile-menu')"
        >
          Give
        </router-link>
      </div>
    </div>
  </nav>
</template>

<script>
export default {
  name: 'NavbarComponent',
  props: {
    churchName: {
      type: String,
      default: 'Muang Thai Korat Church',
    },
    navPages: {
      type: Array,
      default: () => [],
    },
    currentPage: {
      type: String,
      default: 'home',
    },
    mobileMenuOpen: {
      type: Boolean,
      default: false,
    },
  },
  methods: {
    getRoutePath(pageId) {
      const routeMap = {
        home: '/',
        about: '/about',
        'our-ministries': '/our-ministries',
        events: '/events',
        contact: '/contact',
        give: '/give',
      };
      return routeMap[pageId] || '/';
    },
    navClass(pageId) {
      const currentPath = this.$route.path;
      const pagePath = this.getRoutePath(pageId);
      return currentPath === pagePath ? 'text-white' : 'text-gray-300';
    },
    mobileNavClass(pageId) {
      const currentPath = this.$route.path;
      const pagePath = this.getRoutePath(pageId);
      return currentPath === pagePath ? 'text-blue-400' : 'text-white';
    },
    onLogoClick() {
      if (window.location.pathname !== '/') {
        window.location.href = '/';
        return;
      }
      window.scrollTo({ top: 0, behavior: 'smooth' });
    },
  },
};
</script>
