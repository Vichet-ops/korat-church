<template>
  <div>
    <!-- Full Screen Hero Section -->
    <div
      class="relative min-h-screen flex items-center justify-center"
      style="
        background: linear-gradient(rgba(0, 0, 0, 0.4), rgba(0, 0, 0, 0.4)),
          url('/images/church_building_1.jpg') no-repeat center center;
        background-size: cover;
      "
    >
      <!-- Hero Content -->
      <div class="relative z-10 text-center text-white px-4 sm:px-6 lg:px-8">
        <h1 class="text-4xl md:text-6xl font-bold mb-6 leading-tight">
          Welcome to
          <span class="text-blue-300 text-5xl md:text-8xl">{{
            churchSettings.church_name || 'Muang Thai Korat Church'
          }}</span>
        </h1>
        <p
          class="text-xl md:text-2xl text-gray-200 mb-2 max-w-3xl mx-auto leading-relaxed"
        >
          Join us in worship, fellowship, and growing together in faith.
        </p>
        <p
          class="text-lg md:text-xl text-gray-300 mb-8 max-w-3xl mx-auto leading-relaxed"
        >
          Experience the love of Christ in our community.
        </p>
        <div class="flex flex-col sm:flex-row gap-4 justify-center">
          <button
            @click="$emit('navigate', 'services')"
            class="bg-gray-700 text-white px-8 py-4 rounded-lg text-lg font-semibold hover:bg-gray-800 transition duration-300 transform hover:scale-105"
          >
            Join Our Services
          </button>
          <button
            @click="$emit('navigate', 'about')"
            class="border-2 border-white text-white px-8 py-4 rounded-lg text-lg font-semibold hover:bg-white hover:text-gray-900 transition duration-300 transform hover:scale-105"
          >
            Learn More
          </button>
        </div>
      </div>
    </div>

    <!-- Welcome Section -->
    <section class="py-20 relative bg-gray-200">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="text-center">
          <h2 class="text-4xl md:text-5xl font-bold text-gray-800 mb-6">
            Welcome to Our <span class="text-blue-600">Faith Community</span>
          </h2>
          <p
            class="text-xl text-gray-600 max-w-3xl mx-auto leading-relaxed mb-16"
          >
            Join us in worship, fellowship, and growing together in faith.
            <br />
            Experience the love of Christ in our welcoming community where
            everyone belongs.
          </p>
        </div>

        <div class="grid grid-cols-1 md:grid-cols-3 gap-8 mt-12">
          <div
            class="text-center p-8 rounded-lg bg-gray-300 shadow-sm transition-shadow duration-300 ease-out hover:shadow-md"
          >
            <div
              class="w-16 h-16 bg-gray-700 rounded-full flex items-center justify-center mx-auto mb-6"
            >
              <svg
                class="w-8 h-8 text-white"
                fill="none"
                stroke="currentColor"
                viewBox="0 0 24 24"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M12 6.253v13m0-13C10.832 5.477 9.246 5 7.5 5S4.168 5.477 3 6.253v13C4.168 18.477 5.754 18 7.5 18s3.332.477 4.5 1.253m0-13C13.168 5.477 14.754 5 16.5 5c1.746 0 3.332.477 4.5 1.253v13C19.832 18.477 18.246 18 16.5 18c-1.746 0-3.332.477-4.5 1.253"
                ></path>
              </svg>
            </div>
            <h3 class="text-2xl font-semibold text-gray-800 mb-4">
              Bible Study
            </h3>
            <p class="text-gray-700">
              Deepen your understanding of God's Word through our engaging Bible
              study programs for all ages.
            </p>
          </div>

          <div
            class="text-center p-8 rounded-lg bg-gray-300 shadow-sm transition-shadow duration-300 ease-out hover:shadow-md"
          >
            <div
              class="w-16 h-16 bg-gray-700 rounded-full flex items-center justify-center mx-auto mb-6"
            >
              <svg
                class="w-8 h-8 text-white"
                fill="none"
                stroke="currentColor"
                viewBox="0 0 24 24"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z"
                ></path>
              </svg>
            </div>
            <h3 class="text-2xl font-semibold text-gray-800 mb-4">Community</h3>
            <p class="text-gray-700">
              Build meaningful relationships in our supportive community where
              faith and friendship come together.
            </p>
          </div>

          <div
            class="text-center p-8 rounded-lg bg-gray-300 shadow-sm transition-shadow duration-300 ease-out hover:shadow-md"
          >
            <div
              class="w-16 h-16 bg-gray-700 rounded-full flex items-center justify-center mx-auto mb-6"
            >
              <svg
                class="w-8 h-8 text-white"
                fill="none"
                stroke="currentColor"
                viewBox="0 0 24 24"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M4.318 6.318a4.5 4.5 0 000 6.364L12 20.364l7.682-7.682a4.5 4.5 0 00-6.364-6.364L12 7.636l-1.318-1.318a4.5 4.5 0 00-6.364 0z"
                ></path>
              </svg>
            </div>
            <h3 class="text-2xl font-semibold text-gray-800 mb-4">Worship</h3>
            <p class="text-gray-700">
              Experience inspiring worship services that uplift your spirit and
              draw you closer to God.
            </p>
          </div>
        </div>
      </div>
    </section>

    <!-- Service Times Section -->
    <section
      class="py-20 bg-gray-200"
      style="
        background: linear-gradient(rgba(0, 0, 0, 0.5), rgba(0, 0, 0, 0.5)),
          url('/images/worship.jpg') no-repeat center center;
        background-size: cover;
      "
    >
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="text-center max-w-4xl mx-auto">
          <h2 class="text-4xl md:text-5xl font-bold text-white mb-6">
            Join Us for <span class="text-blue-300">Worship</span>
          </h2>
          <p class="text-xl text-gray-200 mb-8">
            We welcome you to join us for our weekly services.<br />
            Come as you are and experience the love of Christ in our community.
          </p>

          <div class="grid grid-cols-1 md:grid-cols-3 gap-6 max-w-4xl mx-auto">
            <div
              v-for="service in serviceTimes"
              :key="service.day"
              class="text-center p-6 bg-white rounded-xl shadow-sm transition-all duration-300 ease-out hover:shadow-lg hover:scale-105"
            >
              <div
                class="w-16 h-16 rounded-full flex items-center justify-center mx-auto mb-4"
                :class="
                  service.day === 'Sunday'
                    ? 'bg-blue-600'
                    : service.day === 'Wednesday'
                    ? 'bg-green-600'
                    : 'bg-purple-600'
                "
              >
                <svg
                  class="w-8 h-8 text-white"
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
              </div>
              <h4 class="text-xl font-semibold text-gray-900 mb-2">
                {{ service.day }} {{ service.service }}
              </h4>
              <p
                class="text-2xl font-bold mb-2"
                :class="
                  service.day === 'Sunday'
                    ? 'text-blue-600'
                    : service.day === 'Wednesday'
                    ? 'text-green-600'
                    : 'text-purple-600'
                "
              >
                {{ service.time }}
              </p>
              <p class="text-gray-600 text-sm">{{ service.service }}</p>
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- Events Section -->
    <section class="py-20 bg-gray-100">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="text-center mb-16">
          <h2 class="text-4xl md:text-5xl font-bold text-gray-900 mb-6">
            Upcoming <span class="text-blue-600">Events</span>
          </h2>
          <p class="text-xl text-gray-600 max-w-3xl mx-auto">
            Join us for special events, celebrations, and community gatherings
            throughout the year.
          </p>
        </div>

        <div
          v-if="events.length > 0"
          class="grid grid-cols-1 md:grid-cols-3 gap-8"
        >
          <div
            v-for="(event, index) in events.slice(0, 3)"
            :key="event.id"
            class="bg-white rounded-xl shadow-lg overflow-hidden"
          >
            <div
              class="h-48 bg-cover bg-center hover:scale-105 transition-transform duration-500"
              :style="`background-image: url('/images/${
                index === 0
                  ? 'christmas_eve.jpg'
                  : index === 1
                  ? 'potluck.jpg'
                  : 'bible_study.jpg'
              }')`"
            ></div>
            <div class="p-6">
              <h3 class="text-xl font-bold text-gray-900 mb-3">
                {{ event.title }}
              </h3>
              <p class="text-gray-600 mb-4">{{ event.description }}</p>
              <button
                @click="$emit('navigate', 'events')"
                class="text-blue-600 font-semibold hover:text-blue-700 transition duration-300"
              >
                Learn More →
              </button>
            </div>
          </div>
        </div>

        <div class="text-center mt-12">
          <button
            @click="$emit('navigate', 'events')"
            class="inline-block bg-blue-600 text-white px-8 py-4 rounded-lg font-semibold hover:bg-blue-700 transition duration-300"
          >
            View All Events
          </button>
        </div>
      </div>
    </section>

    <!-- Call to Action -->
    <section class="py-20 bg-gray-100">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 text-center">
        <h2 class="text-4xl md:text-5xl font-bold text-gray-900 mb-6">
          Ready to Join Our <span class="text-blue-600">Community?</span>
        </h2>
        <p class="text-xl text-gray-600 mb-8 max-w-3xl mx-auto">
          We'd love to welcome you to our church family. Come experience the
          love of Christ in our community.
        </p>
        <div class="flex flex-col sm:flex-row gap-4 justify-center">
          <button
            @click="$emit('navigate', 'contact')"
            class="bg-blue-600 text-white px-8 py-4 rounded-lg font-semibold hover:bg-blue-700 hover:shadow-lg transition-all duration-300"
          >
            Contact Us
          </button>
          <button
            @click="$emit('navigate', 'about')"
            class="border-2 border-blue-600 text-blue-600 px-8 py-4 rounded-lg font-semibold hover:bg-blue-600 hover:text-white hover:shadow-lg transition-all duration-300"
          >
            Learn More
          </button>
        </div>
      </div>
    </section>
  </div>
</template>

<script>
export default {
  name: 'HomePage',
  props: {
    events: Array,
    churchSettings: Object,
    serviceTimes: Array,
  },
};
</script>
