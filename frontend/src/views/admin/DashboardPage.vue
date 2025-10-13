<template>
  <div class="pt-6">
    <h1 class="text-3xl font-bold text-gray-900 mb-8">Dashboard</h1>

    <!-- Stats Grid -->
    <div class="grid grid-cols-1 md:grid-cols-3 gap-6 mb-8">
      <!-- New Messages -->
      <div
        class="rounded-lg shadow-lg p-6 text-white"
        style="background: linear-gradient(to bottom right, #1e4d7a, #2d5a8f)"
      >
        <div>
          <p class="text-blue-200 text-sm font-medium">New Messages</p>
          <p class="text-3xl font-bold mt-2">{{ stats.newMessages }}</p>
        </div>
      </div>

      <!-- Total Messages -->
      <div
        class="rounded-lg shadow-lg p-6 text-white"
        style="background: linear-gradient(to bottom right, #1e4d7a, #2d5a8f)"
      >
        <div>
          <p class="text-blue-200 text-sm font-medium">Total Messages</p>
          <p class="text-3xl font-bold mt-2">{{ stats.totalMessages }}</p>
        </div>
      </div>

      <!-- Admin Info -->
      <div
        class="rounded-lg shadow-lg p-6 text-white"
        style="background: linear-gradient(to bottom right, #1e4d7a, #2d5a8f)"
      >
        <div>
          <p class="text-blue-200 text-sm font-medium">Logged in as</p>
          <p class="text-xl font-bold mt-2">{{ adminInfo.username }}</p>
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
      // Parse the date string
      const date = new Date(dateString);

      // Check if date is valid
      if (isNaN(date.getTime())) {
        return 'Invalid date';
      }

      const now = new Date();
      const diff = now - date;
      const minutes = Math.floor(diff / (1000 * 60));
      const hours = Math.floor(diff / (1000 * 60 * 60));
      const days = Math.floor(hours / 24);

      if (minutes < 1) return 'Just now';
      if (minutes < 60) return `${minutes}m ago`;
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
