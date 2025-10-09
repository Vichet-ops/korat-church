<template>
  <div class="pt-6">
    <h1 class="text-3xl font-bold text-gray-900 mb-8">Dashboard</h1>

    <!-- Stats Grid -->
    <div class="grid grid-cols-1 md:grid-cols-3 gap-6 mb-8">
      <!-- New Messages -->
      <div
        class="bg-gradient-to-br from-blue-500 to-blue-600 rounded-lg shadow-lg p-6 text-white"
      >
        <div class="flex items-center justify-between">
          <div>
            <p class="text-blue-100 text-sm font-medium">New Messages</p>
            <p class="text-3xl font-bold mt-2">{{ stats.newMessages }}</p>
          </div>
          <div class="bg-blue-400 bg-opacity-30 rounded-full p-3">
            <svg
              class="w-8 h-8"
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M3 8l7.89 5.26a2 2 0 002.22 0L21 8M5 19h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z"
              />
            </svg>
          </div>
        </div>
      </div>

      <!-- Total Messages -->
      <div
        class="bg-gradient-to-br from-green-500 to-green-600 rounded-lg shadow-lg p-6 text-white"
      >
        <div class="flex items-center justify-between">
          <div>
            <p class="text-green-100 text-sm font-medium">Total Messages</p>
            <p class="text-3xl font-bold mt-2">{{ stats.totalMessages }}</p>
          </div>
          <div class="bg-green-400 bg-opacity-30 rounded-full p-3">
            <svg
              class="w-8 h-8"
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M7 8h10M7 12h4m1 8l-4-4H5a2 2 0 01-2-2V6a2 2 0 012-2h14a2 2 0 012 2v8a2 2 0 01-2 2h-3l-4 4z"
              />
            </svg>
          </div>
        </div>
      </div>

      <!-- Admin Info -->
      <div
        class="bg-gradient-to-br from-purple-500 to-purple-600 rounded-lg shadow-lg p-6 text-white"
      >
        <div class="flex items-center justify-between">
          <div>
            <p class="text-purple-100 text-sm font-medium">Logged in as</p>
            <p class="text-xl font-bold mt-2">{{ adminInfo.username }}</p>
          </div>
          <div class="bg-purple-400 bg-opacity-30 rounded-full p-3">
            <svg
              class="w-8 h-8"
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"
              />
            </svg>
          </div>
        </div>
      </div>
    </div>

    <!-- Recent Messages -->
    <div class="bg-white rounded-lg shadow-lg p-6">
      <div class="flex items-center justify-between mb-6">
        <h2 class="text-xl font-bold text-gray-900">Recent Messages</h2>
        <router-link
          to="/admin/messages"
          class="text-blue-600 hover:text-blue-700 font-medium text-sm"
        >
          View All →
        </router-link>
      </div>

      <div v-if="loading" class="text-center py-8">
        <div
          class="inline-block animate-spin rounded-full h-8 w-8 border-b-2 border-blue-600"
        ></div>
        <p class="mt-2 text-gray-600">Loading messages...</p>
      </div>

      <div
        v-else-if="recentMessages.length === 0"
        class="text-center py-8 text-gray-500"
      >
        No messages yet
      </div>

      <div v-else class="space-y-4">
        <div
          v-for="message in recentMessages"
          :key="message.id"
          class="border border-gray-200 rounded-lg p-4 hover:border-blue-300 transition-colors"
        >
          <div class="flex items-start justify-between">
            <div class="flex-1">
              <div class="flex items-center gap-3 mb-2">
                <h3 class="font-semibold text-gray-900">{{ message.name }}</h3>
                <span
                  :class="{
                    'bg-green-100 text-green-800': message.status === 'new',
                    'bg-blue-100 text-blue-800': message.status === 'read',
                    'bg-gray-100 text-gray-800': message.status === 'replied',
                  }"
                  class="px-2 py-1 rounded-full text-xs font-medium"
                >
                  {{ message.status }}
                </span>
              </div>
              <p class="text-sm text-gray-600 mb-1">{{ message.email }}</p>
              <p class="text-sm font-medium text-gray-700 mb-2">
                {{ message.subject }}
              </p>
              <p class="text-sm text-gray-600 line-clamp-2">
                {{ message.message }}
              </p>
            </div>
            <span class="text-xs text-gray-500 ml-4">
              {{ formatDate(message.created_at) }}
            </span>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'DashboardPage',
  data() {
    return {
      loading: true,
      stats: {
        newMessages: 0,
        totalMessages: 0,
      },
      recentMessages: [],
      adminInfo: {},
    };
  },
  methods: {
    async loadDashboardData() {
      try {
        const token = localStorage.getItem('admin_token');
        const apiUrl = import.meta.env.VITE_API_URL || 'http://localhost:8081';

        const response = await fetch(`${apiUrl}/api/admin/messages`, {
          headers: {
            Authorization: `Bearer ${token}`,
          },
        });

        if (!response.ok) {
          throw new Error('Failed to load messages');
        }

        const data = await response.json();
        this.stats.totalMessages = data.count;
        this.stats.newMessages = data.messages.filter(
          (m) => m.status === 'new'
        ).length;
        this.recentMessages = data.messages.slice(0, 5);
      } catch (error) {
        console.error('Error loading dashboard:', error);
      } finally {
        this.loading = false;
      }
    },
    formatDate(dateString) {
      const date = new Date(dateString);
      const now = new Date();
      const diff = now - date;
      const hours = Math.floor(diff / (1000 * 60 * 60));
      const days = Math.floor(hours / 24);

      if (hours < 1) return 'Just now';
      if (hours < 24) return `${hours}h ago`;
      if (days < 7) return `${days}d ago`;
      return date.toLocaleDateString();
    },
  },
  mounted() {
    const adminInfoStr = localStorage.getItem('admin_info');
    if (adminInfoStr) {
      this.adminInfo = JSON.parse(adminInfoStr);
    }
    this.loadDashboardData();
  },
};
</script>
