<template>
  <div>
    <div
      class="flex flex-col sm:flex-row sm:items-center sm:justify-between mt-8 mb-8 gap-4"
    >
      <h1 class="text-2xl sm:text-3xl font-bold text-gray-900">
        Contact Messages
      </h1>
      <div class="flex flex-wrap gap-2">
        <button
          @click="filterStatus = 'all'"
          :class="
            filterStatus === 'all'
              ? 'bg-blue-600 text-white'
              : 'bg-gray-200 text-gray-700'
          "
          class="px-4 py-2 rounded-lg font-medium text-sm transition-colors cursor-pointer"
        >
          All ({{ stats.total }})
        </button>
        <button
          @click="filterStatus = 'new'"
          :class="
            filterStatus === 'new'
              ? 'bg-green-600 text-white'
              : 'bg-gray-200 text-gray-700'
          "
          class="px-4 py-2 rounded-lg font-medium text-sm transition-colors cursor-pointer"
        >
          New ({{ stats.new }})
        </button>
        <button
          @click="filterStatus = 'read'"
          :class="
            filterStatus === 'read'
              ? 'bg-blue-600 text-white'
              : 'bg-gray-200 text-gray-700'
          "
          class="px-4 py-2 rounded-lg font-medium text-sm transition-colors cursor-pointer"
        >
          Read ({{ stats.read }})
        </button>
        <button
          @click="filterStatus = 'replied'"
          :class="
            filterStatus === 'replied'
              ? 'bg-purple-600 text-white'
              : 'bg-gray-200 text-gray-700'
          "
          class="px-4 py-2 rounded-lg font-medium text-sm transition-colors cursor-pointer"
        >
          Replied ({{ stats.replied }})
        </button>
      </div>
    </div>

    <div v-if="loading" class="bg-white rounded-lg shadow-lg p-12 text-center">
      <div
        class="inline-block animate-spin rounded-full h-12 w-12 border-b-2 border-blue-600"
      ></div>
      <p class="mt-4 text-gray-600">Loading messages...</p>
    </div>

    <div
      v-else-if="filteredMessages.length === 0"
      class="bg-white rounded-lg shadow-lg p-12 text-center"
    >
      <svg
        class="mx-auto h-12 w-12 text-gray-400"
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
      <p class="mt-4 text-gray-600">No messages found</p>
    </div>

    <div v-else>
      <!-- Mobile Card View (Hidden on MD and up) -->
      <div class="md:hidden space-y-4">
        <div
          v-for="message in filteredMessages"
          :key="message.id"
          class="bg-white rounded-lg shadow-lg p-4"
        >
          <div class="flex items-center justify-between mb-3">
            <span
              :class="{
                'bg-green-100 text-green-800': message.status === 'new',
                'bg-blue-100 text-blue-800': message.status === 'read',
                'bg-gray-100 text-gray-800': message.status === 'replied',
              }"
              class="px-2 py-1 text-xs font-semibold rounded-full"
            >
              {{ message.status }}
            </span>
            <span class="text-xs text-gray-500">
              {{ formatDate(message.created_at) }}
            </span>
          </div>

          <h3 class="font-semibold text-gray-900 mb-1">{{ message.name }}</h3>
          <p class="text-sm text-gray-600 mb-2">{{ message.email }}</p>
          <p class="text-sm font-medium text-gray-700 mb-3">
            {{ message.subject }}
          </p>

          <div class="flex gap-2">
            <button
              @click="viewMessage(message)"
              class="flex-1 px-3 py-2 bg-blue-600 text-white text-sm rounded-lg hover:bg-blue-700 cursor-pointer"
            >
              View
            </button>
            <button
              @click="deleteMessage(message.id)"
              class="px-3 py-2 bg-red-600 text-white text-sm rounded-lg hover:bg-red-700 cursor-pointer"
            >
              Delete
            </button>
          </div>
        </div>
      </div>

      <!-- Desktop Table View (Hidden on SM and below) -->
      <div
        class="hidden md:block bg-white rounded-lg shadow-lg overflow-hidden"
      >
        <div class="overflow-x-auto">
          <table class="min-w-full divide-y divide-gray-200">
            <thead class="bg-gray-50">
              <tr>
                <th
                  class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
                >
                  Status
                </th>
                <th
                  class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
                >
                  Name
                </th>
                <th
                  class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
                >
                  Email
                </th>
                <th
                  class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
                >
                  Subject
                </th>
                <th
                  class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
                >
                  Date
                </th>
                <th
                  class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
                >
                  Actions
                </th>
              </tr>
            </thead>
            <tbody class="bg-white divide-y divide-gray-200">
              <tr
                v-for="message in filteredMessages"
                :key="message.id"
                class="hover:bg-gray-50"
              >
                <td class="px-6 py-4 whitespace-nowrap">
                  <span
                    :class="{
                      'bg-green-100 text-green-800': message.status === 'new',
                      'bg-blue-100 text-blue-800': message.status === 'read',
                      'bg-gray-100 text-gray-800': message.status === 'replied',
                    }"
                    class="px-2 inline-flex text-xs leading-5 font-semibold rounded-full"
                  >
                    {{ message.status }}
                  </span>
                </td>
                <td class="px-6 py-4 whitespace-nowrap">
                  <div class="text-sm font-medium text-gray-900">
                    {{ message.name }}
                  </div>
                </td>
                <td class="px-6 py-4 whitespace-nowrap">
                  <div class="text-sm text-gray-500">{{ message.email }}</div>
                </td>
                <td class="px-6 py-4">
                  <div class="text-sm text-gray-900">{{ message.subject }}</div>
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                  {{ formatDate(message.created_at) }}
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                  <button
                    @click="viewMessage(message)"
                    class="text-blue-600 hover:text-blue-900 mr-3 cursor-pointer"
                  >
                    View
                  </button>
                  <button
                    @click="deleteMessage(message.id)"
                    class="text-red-600 hover:text-red-900 cursor-pointer"
                  >
                    Delete
                  </button>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </div>

    <!-- Message Detail Modal -->
    <div
      v-if="selectedMessage"
      class="fixed inset-0 bg-gray-600 bg-opacity-50 overflow-y-auto h-full w-full z-50"
      @click="closeModal"
    >
      <div
        class="relative top-20 mx-auto p-8 border w-full max-w-2xl shadow-lg rounded-lg bg-white"
        @click.stop
      >
        <div class="flex items-center justify-between mb-6">
          <h3 class="text-2xl font-bold text-gray-900">Message Details</h3>
          <button
            @click="closeModal"
            class="text-gray-400 hover:text-gray-500 cursor-pointer"
          >
            <svg
              class="h-6 w-6"
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

        <div class="space-y-4">
          <div>
            <label class="text-sm font-medium text-gray-500">From</label>
            <p class="text-lg font-semibold text-gray-900">
              {{ selectedMessage.name }}
            </p>
            <p class="text-sm text-gray-600">{{ selectedMessage.email }}</p>
          </div>

          <div>
            <label class="text-sm font-medium text-gray-500">Subject</label>
            <p class="text-lg text-gray-900">{{ selectedMessage.subject }}</p>
          </div>

          <div>
            <label class="text-sm font-medium text-gray-500">Message</label>
            <p class="text-gray-900 whitespace-pre-wrap">
              {{ selectedMessage.message }}
            </p>
          </div>

          <div>
            <label class="text-sm font-medium text-gray-500">Date</label>
            <p class="text-sm text-gray-600">
              {{ formatDate(selectedMessage.created_at) }}
            </p>
          </div>

          <div>
            <label class="text-sm font-medium text-gray-500 mb-2 block"
              >Status</label
            >
            <select
              v-model="selectedMessage.status"
              @change="updateStatus"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none"
            >
              <option value="new">New</option>
              <option value="read">Read</option>
              <option value="replied">Replied</option>
            </select>
          </div>

          <div class="pt-4">
            <a
              :href="`mailto:${selectedMessage.email}?subject=Re: ${selectedMessage.subject}`"
              class="inline-block px-6 py-3 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors"
            >
              Reply via Email
            </a>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'MessagesPage',
  data() {
    return {
      loading: true,
      messages: [],
      filterStatus: 'all',
      selectedMessage: null,
      stats: {
        total: 0,
        new: 0,
        read: 0,
        replied: 0,
      },
    };
  },
  computed: {
    filteredMessages() {
      if (this.filterStatus === 'all') {
        return this.messages;
      }
      return this.messages.filter((m) => m.status === this.filterStatus);
    },
  },
  methods: {
    async loadMessages() {
      try {
        const token = localStorage.getItem('admin_token');
        const apiUrl = import.meta.env.VITE_API_URL || 'https://vichetkeo.com';

        const response = await fetch(`${apiUrl}/api/admin/messages`, {
          headers: {
            Authorization: `Bearer ${token}`,
          },
        });

        if (!response.ok) {
          throw new Error('Failed to load messages');
        }

        const data = await response.json();
        this.messages = data.messages;
        this.stats.total = data.count;
        this.stats.new = data.messages.filter((m) => m.status === 'new').length;
        this.stats.read = data.messages.filter(
          (m) => m.status === 'read'
        ).length;
        this.stats.replied = data.messages.filter(
          (m) => m.status === 'replied'
        ).length;
      } catch (error) {
        console.error('Error loading messages:', error);
        alert('Failed to load messages');
      } finally {
        this.loading = false;
      }
    },
    async deleteMessage(id) {
      if (!confirm('Are you sure you want to delete this message?')) {
        return;
      }

      try {
        const token = localStorage.getItem('admin_token');
        const apiUrl = import.meta.env.VITE_API_URL || 'https://vichetkeo.com';

        const response = await fetch(`${apiUrl}/api/admin/messages/${id}`, {
          method: 'DELETE',
          headers: {
            Authorization: `Bearer ${token}`,
          },
        });

        if (!response.ok) {
          throw new Error('Failed to delete message');
        }

        this.messages = this.messages.filter((m) => m.id !== id);
        this.stats.total--;
        alert('Message deleted successfully');
      } catch (error) {
        console.error('Error deleting message:', error);
        alert('Failed to delete message');
      }
    },
    async updateStatus() {
      try {
        const token = localStorage.getItem('admin_token');
        const apiUrl = import.meta.env.VITE_API_URL || 'https://vichetkeo.com';

        const response = await fetch(
          `${apiUrl}/api/admin/messages/${this.selectedMessage.id}/status`,
          {
            method: 'PATCH',
            headers: {
              'Content-Type': 'application/json',
              Authorization: `Bearer ${token}`,
            },
            body: JSON.stringify({
              status: this.selectedMessage.status,
            }),
          }
        );

        if (!response.ok) {
          throw new Error('Failed to update status');
        }

        // Update the message in the list
        const index = this.messages.findIndex(
          (m) => m.id === this.selectedMessage.id
        );
        if (index !== -1) {
          this.messages[index].status = this.selectedMessage.status;
        }

        // Recalculate stats
        this.stats.new = this.messages.filter((m) => m.status === 'new').length;
        this.stats.read = this.messages.filter(
          (m) => m.status === 'read'
        ).length;
        this.stats.replied = this.messages.filter(
          (m) => m.status === 'replied'
        ).length;

        this.selectedMessage = null;
        alert('Status updated successfully');
      } catch (error) {
        console.error('Error updating status:', error);
        alert('Failed to update status');
      }
    },
    viewMessage(message) {
      this.selectedMessage = { ...message };
      // Auto-mark as read if it's new
      if (message.status === 'new') {
        this.selectedMessage.status = 'read';
        this.updateStatus();
      }
    },
    closeModal() {
      this.selectedMessage = null;
    },
    formatDate(dateString) {
      return new Date(dateString).toLocaleString();
    },
  },
  mounted() {
    this.loadMessages();
  },
};
</script>
