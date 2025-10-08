<template>
  <div>
    <!-- Hero Section -->
    <div
      class="relative h-64 md:h-80 lg:h-96 flex items-center justify-center"
      style="
        background: linear-gradient(rgba(0, 0, 0, 0.5), rgba(0, 0, 0, 0.5)),
          url('/images/bible_study.jpg') no-repeat center center;
        background-size: cover;
      "
    >
      <div class="relative max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 text-center">
        <h1 class="text-4xl md:text-6xl font-bold text-white mb-6">
          Church Events
        </h1>
        <p class="text-xl md:text-2xl text-gray-200 max-w-3xl mx-auto">
          Join us for special events, celebrations, and community gatherings.
        </p>
      </div>
    </div>

    <!-- Events Section -->
    <section class="pt-24 lg:pt-28 pb-20 bg-gray-50">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div
          v-if="events.length > 0"
          class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8"
        >
          <div
            v-for="(event, index) in events"
            :key="event.id"
            class="bg-white rounded-xl shadow-lg overflow-hidden hover:shadow-xl transition-shadow duration-300"
          >
            <div
              class="h-48 bg-cover bg-center"
              :style="`background-image: url('/images/${getEventImage(
                index
              )}')`"
            ></div>
            <div class="p-6">
              <h3 class="text-2xl font-bold text-gray-900 mb-3">
                {{ event.title }}
              </h3>
              <p class="text-gray-600 mb-4">{{ event.description }}</p>
              <div class="space-y-2">
                <div class="flex items-center text-gray-600">
                  <svg
                    class="w-5 h-5 mr-2"
                    fill="none"
                    stroke="currentColor"
                    viewBox="0 0 24 24"
                  >
                    <path
                      stroke-linecap="round"
                      stroke-linejoin="round"
                      stroke-width="2"
                      d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z"
                    ></path>
                  </svg>
                  <span>{{ formatDate(event.date) }}</span>
                </div>
                <div v-if="event.time" class="flex items-center text-gray-600">
                  <svg
                    class="w-5 h-5 mr-2"
                    fill="none"
                    stroke="currentColor"
                    viewBox="0 0 24 24"
                  >
                    <path
                      stroke-linecap="round"
                      stroke-linejoin="round"
                      stroke-width="2"
                      d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"
                    ></path>
                  </svg>
                  <span>{{ event.time }}</span>
                </div>
                <div
                  v-if="event.location"
                  class="flex items-center text-gray-600"
                >
                  <svg
                    class="w-5 h-5 mr-2"
                    fill="none"
                    stroke="currentColor"
                    viewBox="0 0 24 24"
                  >
                    <path
                      stroke-linecap="round"
                      stroke-linejoin="round"
                      stroke-width="2"
                      d="M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.827 0l-4.244-4.243a8 8 0 1111.314 0z"
                    ></path>
                    <path
                      stroke-linecap="round"
                      stroke-linejoin="round"
                      stroke-width="2"
                      d="M15 11a3 3 0 11-6 0 3 3 0 016 0z"
                    ></path>
                  </svg>
                  <span>{{ event.location }}</span>
                </div>
              </div>
            </div>
          </div>
        </div>
        <div v-else class="text-center py-16">
          <svg
            class="mx-auto h-12 w-12 text-gray-400"
            fill="none"
            viewBox="0 0 24 24"
            stroke="currentColor"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z"
            />
          </svg>
          <h3 class="mt-2 text-xl font-medium text-gray-900">
            No events scheduled
          </h3>
          <p class="mt-1 text-gray-500">Check back soon for upcoming events!</p>
        </div>
      </div>
    </section>

    <!-- Call to Action -->
    <section class="py-20 bg-white">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 text-center">
        <h2 class="text-3xl md:text-4xl font-bold text-gray-900 mb-6">
          Stay <span class="text-blue-600">Connected</span>
        </h2>
        <p class="text-xl text-gray-600 mb-8 max-w-3xl mx-auto">
          Want to stay updated on upcoming events? Contact us to join our
          mailing list.
        </p>
        <button
          @click="$emit('navigate', 'contact')"
          class="bg-blue-600 text-white px-6 py-3 rounded-lg text-base font-semibold hover:bg-blue-700 transition duration-300 transform hover:scale-105 cursor-pointer"
        >
          Contact Us
        </button>
      </div>
    </section>
  </div>
</template>

<script>
export default {
  name: 'EventsPage',
  props: {
    events: Array,
  },
  methods: {
    formatDate(dateString) {
      const date = new Date(dateString);
      return date.toLocaleDateString('en-US', {
        year: 'numeric',
        month: 'long',
        day: 'numeric',
      });
    },
    getEventImage(index) {
      const images = [
        'christmas_eve.jpg',
        'potluck.jpg',
        'bible_study.jpg',
        'Annual Church Picnic.jpg',
        'Christmas Celebration.jpg',
        'Easter Sunday Service.jpg',
      ];
      return images[index % images.length];
    },
  },
};
</script>
