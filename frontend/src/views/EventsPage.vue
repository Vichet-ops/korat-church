<template>
  <div>
    <!-- Hero Section -->
    <div
      class="relative h-[28rem] md:h-[500px] lg:h-[600px] flex items-center justify-center pt-24 lg:pt-32"
      style="
        background: linear-gradient(rgba(0, 0, 0, 0.5), rgba(0, 0, 0, 0.5)),
          url('/images/christmas_eve.jpg') no-repeat center 65%;
        background-size: cover;
      "
    >
      <div
        class="relative max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 text-center flex flex-col items-center justify-center h-full"
      >
        <h1
          class="text-5xl md:text-7xl xl:text-8xl font-extrabold text-white tracking-tight mb-6"
        >
          Our Church Events
        </h1>
        <p class="text-xl md:text-2xl text-white max-w-3xl mx-auto">
          Join us for special events, celebrations, and community gatherings.
        </p>
      </div>
    </div>

    <!-- Events Section -->
    <section class="py-20 bg-white">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="text-center mb-16">
          <h2
            class="text-3xl md:text-4xl lg:text-5xl font-bold text-gray-900 mb-6"
          >
            Upcoming <span class="text-blue-600">Events</span>
          </h2>
          <p class="text-xl text-gray-600 max-w-3xl mx-auto">
            Join us for these special occasions and community gatherings
            throughout the year.
          </p>
        </div>
        <!-- Show API events if available, otherwise show sample events -->
        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
          <div
            v-for="(event, index) in displayEvents"
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
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- Call to Action -->
    <section class="py-20 bg-blue-50">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 text-center">
        <h2
          class="text-3xl md:text-4xl lg:text-5xl font-bold text-gray-900 mb-6"
        >
          Stay <span class="text-blue-600">Connected</span>
        </h2>
        <p class="text-xl text-gray-600 mb-8 max-w-3xl mx-auto">
          Want to stay updated on upcoming events? Contact us to join our
          mailing list.
        </p>
        <button
          @click="$router.push('/contact')"
          class="border border-blue-600 text-blue-600 px-6 py-3 rounded-lg text-base font-semibold hover:bg-blue-600 hover:text-white hover:shadow-lg transition-all duration-300 cursor-pointer"
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
  data() {
    return {
      sampleEvents: [
        {
          id: 1,
          title: 'Bible Study Group',
          description:
            "Deepen your understanding of God's word in our weekly Bible study sessions.",
          date: '2024-01-16',
          time: '7:00 PM',
          location: 'Fellowship Hall',
        },
        {
          id: 2,
          title: 'Youth Fellowship',
          description:
            'A fun and engaging gathering for our youth community with games, discussions, and snacks.',
          date: '2024-01-18',
          time: '6:30 PM',
          location: 'Youth Room',
        },
        {
          id: 3,
          title: 'Community Outreach',
          description:
            'Join us as we serve our local community through various outreach programs.',
          date: '2024-01-20',
          time: '9:00 AM',
          location: 'Various Locations',
        },
        {
          id: 4,
          title: 'Prayer Meeting',
          description:
            'Come together in prayer for our church, community, and world.',
          date: '2024-01-22',
          time: '7:30 PM',
          location: 'Prayer Room',
        },
        {
          id: 5,
          title: 'Christmas Celebration',
          description:
            'Join us for a beautiful Christmas service celebrating the birth of our Savior with traditional carols, nativity story, and fellowship.',
          date: '2024-12-25',
          time: '10:00 AM',
          location: 'Main Sanctuary',
        },
        {
          id: 6,
          title: 'Easter Sunday Service',
          description:
            'Celebrate the resurrection of our Lord Jesus Christ with joyful worship, special music, and the message of hope and new life.',
          date: '2024-03-31',
          time: '10:00 AM',
          location: 'Main Sanctuary',
        },
      ],
    };
  },
  computed: {
    displayEvents() {
      // Show API events if available, otherwise show sample events
      return this.events && this.events.length > 0
        ? this.events
        : this.sampleEvents;
    },
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
        'bible_study.jpg',
        'community-outreach.jpg',
        'prayer-ministry.jpg',
        'christmas_eve.jpg',
      ];
      return images[index % images.length];
    },
  },
};
</script>
