<template>
  <div class="min-h-screen bg-gray-100 overflow-x-hidden flex">
    <!-- Sidebar -->
    <div
      :class="
        sidebarOpen ? 'translate-x-0' : '-translate-x-full lg:translate-x-0'
      "
      class="fixed inset-y-0 left-0 w-64 bg-white shadow-lg transform transition-transform duration-300 ease-in-out lg:static lg:inset-0 z-30"
      style="will-change: transform"
    >
      <div
        class="flex items-center justify-between h-16 px-6"
        style="background: linear-gradient(to right, #1e4d7a, #2d5a8f)"
      >
        <h1 class="text-xl font-bold text-white">Admin Panel</h1>
        <button @click="sidebarOpen = false" class="lg:hidden text-white">
          <svg
            class="w-6 h-6"
            fill="none"
            stroke="currentColor"
            viewBox="0 0 24 24"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M6 18L18 6M6 6l12 12"
            />
          </svg>
        </button>
      </div>

      <nav class="mt-12 px-4 pb-24">
        <router-link
          v-for="item in menuItems"
          :key="item.path"
          :to="item.path"
          @click="handleMenuClick"
          class="flex items-center px-4 py-3 mb-2 text-gray-700 rounded-lg hover:bg-blue-50 hover:text-blue-600 transition-colors cursor-pointer"
          exact-active-class="bg-blue-50 text-blue-600 font-semibold"
        >
          <svg
            class="w-5 h-5 mr-3"
            fill="none"
            stroke="currentColor"
            viewBox="0 0 24 24"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              :d="item.icon"
            />
          </svg>
          {{ item.name }}
        </router-link>

        <button
          @click="handleLogout"
          class="w-full flex items-center px-4 py-3 mt-4 text-red-600 rounded-lg hover:bg-red-50 transition-colors cursor-pointer"
        >
          <svg
            class="w-5 h-5 mr-3"
            fill="none"
            stroke="currentColor"
            viewBox="0 0 24 24"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M17 16l4-4m0 0l-4-4m4 4H7m6 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h4a3 3 0 013 3v1"
            />
          </svg>
          Logout
        </button>
      </nav>

      <div
        class="absolute bottom-0 left-0 right-0 p-4 border-t border-gray-200 bg-white"
      >
        <div class="text-sm text-gray-600">
          <p class="font-semibold">
            {{ adminInfo.first_name }} {{ adminInfo.last_name }}
          </p>
          <p class="text-xs truncate">{{ adminInfo.email }}</p>
        </div>
      </div>
    </div>

    <!-- Mobile Overlay -->
    <div
      v-if="sidebarOpen"
      @click="sidebarOpen = false"
      class="fixed inset-0 z-20 bg-gray-600 bg-opacity-50 lg:hidden"
    ></div>

    <!-- Main Content -->
    <div class="flex-1 flex flex-col min-w-0">
      <!-- Top Bar -->
      <header class="bg-white shadow-sm sticky top-0 z-10">
        <div class="flex items-center justify-between px-4 sm:px-6 py-4">
          <button
            @click="sidebarOpen = !sidebarOpen"
            class="lg:hidden text-gray-800 hover:bg-gray-100 cursor-pointer p-2 rounded-lg transition-colors"
          >
            <svg
              class="w-6 h-6"
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M4 6h16M4 12h16M4 18h16"
              />
            </svg>
          </button>

          <div class="flex items-center gap-4">
            <span class="text-gray-600 text-sm">
              {{ currentDate }}
            </span>
            <router-link
              to="/"
              target="_blank"
              class="text-blue-600 hover:text-blue-700 text-sm font-medium"
            >
              View Website →
            </router-link>
          </div>
        </div>
      </header>

      <!-- Page Content -->
      <main class="p-4 sm:p-6 flex-1 overflow-auto">
        <router-view></router-view>
      </main>
    </div>
  </div>
</template>

<script>
export default {
  name: 'AdminLayout',
  data() {
    return {
      sidebarOpen: false,
      adminInfo: {},
      menuItems: [
        {
          name: 'Dashboard',
          path: '/admin/dashboard',
          icon: 'M3 12l2-2m0 0l7-7 7 7M5 10v10a1 1 0 001 1h3m10-11l2 2m-2-2v10a1 1 0 01-1 1h-3m-6 0a1 1 0 001-1v-4a1 1 0 011-1h2a1 1 0 011 1v4a1 1 0 001 1m-6 0h6',
        },
        {
          name: 'Messages',
          path: '/admin/messages',
          icon: 'M3 8l7.89 5.26a2 2 0 002.22 0L21 8M5 19h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z',
        },
      ],
    };
  },
  computed: {
    currentDate() {
      return new Date().toLocaleDateString('en-US', {
        weekday: 'long',
        year: 'numeric',
        month: 'long',
        day: 'numeric',
      });
    },
  },
  methods: {
    handleMenuClick() {
      // Close sidebar on mobile after clicking a menu item
      if (window.innerWidth < 1024) {
        this.sidebarOpen = false;
      }
    },
    handleResize() {
      // Update sidebar state based on screen size
      if (window.innerWidth >= 1024) {
        this.sidebarOpen = true;
      } else {
        this.sidebarOpen = false;
      }
    },
    handleLogout() {
      localStorage.removeItem('admin_token');
      localStorage.removeItem('admin_info');
      this.$router.push('/admin/login');
    },
  },
  mounted() {
    const adminInfoStr = localStorage.getItem('admin_info');
    if (adminInfoStr) {
      this.adminInfo = JSON.parse(adminInfoStr);
    }

    // Open sidebar by default on desktop (lg breakpoint is 1024px)
    if (window.innerWidth >= 1024) {
      this.sidebarOpen = true;
    }

    // Listen for window resize
    window.addEventListener('resize', this.handleResize);
  },
  beforeUnmount() {
    // Clean up event listener
    window.removeEventListener('resize', this.handleResize);
  },
};
</script>
