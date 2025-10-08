<template>
  <div>
    <!-- Hero Section -->
    <div
      class="relative h-64 md:h-80 lg:h-96 flex items-center justify-center pt-16"
      style="
        background: linear-gradient(rgba(0, 0, 0, 0.5), rgba(0, 0, 0, 0.5)),
          url('/images/worship.jpg') no-repeat center center;
        background-size: cover;
      "
    >
      <div class="relative max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 text-center">
        <h1 class="text-4xl md:text-6xl font-bold text-white mb-6">
          Our Services & Ministries
        </h1>
        <p class="text-xl md:text-2xl text-gray-200 max-w-3xl mx-auto">
          Join us for worship, fellowship, and spiritual growth. Find your place
          in our community.
        </p>
      </div>
    </div>

    <!-- Service Times Section -->
    <section class="pt-24 lg:pt-28 pb-20 bg-gray-50">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="text-center mb-16">
          <h2 class="text-3xl md:text-4xl font-bold text-gray-900 mb-6">
            Service Times
          </h2>
          <p class="text-xl text-gray-600 max-w-3xl mx-auto">
            Join us for worship and fellowship. All are welcome!
          </p>
        </div>

        <div class="grid md:grid-cols-2 lg:grid-cols-3 gap-8">
          <!-- Sunday Morning Service -->
          <div class="bg-white p-8 rounded-lg shadow-lg">
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
                  d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"
                ></path>
              </svg>
            </div>
            <h3 class="text-2xl font-bold text-gray-900 mb-4 text-center">
              Sunday Morning
            </h3>
            <div class="text-center space-y-3">
              <p class="text-3xl font-bold text-gray-700">10:00 AM</p>
              <p class="text-gray-600">Main Worship Service</p>
              <p class="text-gray-600">Main Sanctuary</p>
              <p class="text-sm text-gray-500">
                Traditional & Contemporary Worship
              </p>
            </div>
          </div>

          <!-- Sunday School -->
          <div class="bg-white p-8 rounded-lg shadow-lg">
            <div
              class="w-16 h-16 bg-green-600 rounded-full flex items-center justify-center mx-auto mb-6"
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
            <h3 class="text-2xl font-bold text-gray-900 mb-4 text-center">
              Sunday School
            </h3>
            <div class="text-center space-y-3">
              <p class="text-3xl font-bold text-green-600">9:00 AM</p>
              <p class="text-gray-600">Bible Study Classes</p>
              <p class="text-gray-600">Education Building</p>
              <p class="text-sm text-gray-500">All Ages Welcome</p>
            </div>
          </div>

          <!-- Wednesday Prayer -->
          <div class="bg-white p-8 rounded-lg shadow-lg">
            <div
              class="w-16 h-16 bg-purple-600 rounded-full flex items-center justify-center mx-auto mb-6"
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
                  d="M8 12h.01M12 12h.01M16 12h.01M21 12c0 4.418-4.03 8-9 8a9.863 9.863 0 01-4.255-.949L3 20l1.395-3.72C3.512 15.042 3 13.574 3 12c0-4.418 4.03-8 9-8s9 3.582 9 8z"
                ></path>
              </svg>
            </div>
            <h3 class="text-2xl font-bold text-gray-900 mb-4 text-center">
              Wednesday Prayer
            </h3>
            <div class="text-center space-y-3">
              <p class="text-3xl font-bold text-purple-600">7:00 PM</p>
              <p class="text-gray-600">Prayer Meeting</p>
              <p class="text-gray-600">Prayer Room</p>
              <p class="text-sm text-gray-500">Community Prayer Time</p>
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- Ministries Section -->
    <section class="py-20 bg-white">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="text-center mb-16">
          <h2 class="text-3xl md:text-4xl font-bold text-gray-900 mb-6">
            Our Ministries
          </h2>
          <p class="text-xl text-gray-600 max-w-3xl mx-auto">
            Discover opportunities to grow, serve, and connect in our church
            community.
          </p>
        </div>

        <div class="grid md:grid-cols-2 lg:grid-cols-3 gap-8">
          <!-- Children's Ministry -->
          <div
            class="bg-white rounded-lg shadow-lg overflow-hidden hover:shadow-xl hover:scale-105 transition-all duration-300 ease-out"
          >
            <div
              class="h-48 bg-cover bg-center"
              style="
                background-image: url('/images/children-ministry.jpg');
                background-position: center 30%;
              "
            ></div>
            <div class="p-6">
              <h3 class="text-2xl font-bold text-gray-900 mb-3">
                Children's Ministry
              </h3>
              <p class="text-gray-600 mb-4">
                Nurturing the faith of our youngest members through creative
                Bible lessons, music, crafts, and fun activities.
              </p>
              <ul class="text-sm text-gray-600 space-y-1">
                <li>• Sunday School (Ages 3-12)</li>
                <li>• Vacation Bible School</li>
                <li>• Children's Choir</li>
                <li>• Special Events & Activities</li>
              </ul>
            </div>
          </div>

          <!-- Youth Ministry -->
          <div
            class="bg-white rounded-lg shadow-lg overflow-hidden hover:shadow-xl hover:scale-105 transition-all duration-300 ease-out"
          >
            <div
              class="h-48 bg-cover bg-center"
              style="
                background-image: url('/images/youth-ministry.jpg');
                background-position: center 25%;
              "
            ></div>
            <div class="p-6">
              <h3 class="text-2xl font-bold text-gray-900 mb-3">
                Youth Ministry
              </h3>
              <p class="text-gray-600 mb-4">
                Creating engaging programs for teens and young adults to grow in
                faith and build meaningful relationships.
              </p>
              <ul class="text-sm text-gray-600 space-y-1">
                <li>• Friday Night Youth Group</li>
                <li>• Bible Study & Discussion</li>
                <li>• Mission Trips & Retreats</li>
                <li>• Community Service Projects</li>
              </ul>
            </div>
          </div>

          <!-- Adult Bible Study -->
          <div
            class="bg-white rounded-lg shadow-lg overflow-hidden hover:shadow-xl hover:scale-105 transition-all duration-300 ease-out"
          >
            <div
              class="h-48 bg-cover bg-center"
              style="
                background-image: url('/images/adult-bible-study.jpg');
                background-position: center 30%;
              "
            ></div>
            <div class="p-6">
              <h3 class="text-2xl font-bold text-gray-900 mb-3">
                Adult Bible Study
              </h3>
              <p class="text-gray-600 mb-4">
                Deepen your understanding of God's Word through interactive
                Bible study and meaningful discussions.
              </p>
              <ul class="text-sm text-gray-600 space-y-1">
                <li>• Sunday School Classes</li>
                <li>• Wednesday Bible Study</li>
                <li>• Small Group Discussions</li>
                <li>• Book Studies & Seminars</li>
              </ul>
            </div>
          </div>

          <!-- Worship Ministry -->
          <div
            class="bg-white rounded-lg shadow-lg overflow-hidden hover:shadow-xl hover:scale-105 transition-all duration-300 ease-out"
          >
            <div
              class="h-48 bg-cover bg-center"
              style="
                background-image: url('/images/worship-ministry.jpg');
                background-position: center 30%;
              "
            ></div>
            <div class="p-6">
              <h3 class="text-2xl font-bold text-gray-900 mb-3">
                Worship Ministry
              </h3>
              <p class="text-gray-600 mb-4">
                Leading our congregation in worship through music, creating an
                atmosphere of praise and adoration.
              </p>
              <ul class="text-sm text-gray-600 space-y-1">
                <li>• Adult Choir</li>
                <li>• Praise Team</li>
                <li>• Instrumentalists</li>
                <li>• Sound & Technical Team</li>
              </ul>
            </div>
          </div>

          <!-- Outreach Ministry -->
          <div
            class="bg-white rounded-lg shadow-lg overflow-hidden hover:shadow-xl hover:scale-105 transition-all duration-300 ease-out"
          >
            <div
              class="h-48 bg-cover bg-center"
              style="
                background-image: url('/images/outreach-ministry.jpg');
                background-position: center 30%;
              "
            ></div>
            <div class="p-6">
              <h3 class="text-2xl font-bold text-gray-900 mb-3">
                Outreach Ministry
              </h3>
              <p class="text-gray-600 mb-4">
                Serving our community and sharing God's love through various
                outreach programs and mission projects.
              </p>
              <ul class="text-sm text-gray-600 space-y-1">
                <li>• Food Bank & Pantry</li>
                <li>• Homeless Ministry</li>
                <li>• Mission Trips</li>
                <li>• Community Events</li>
              </ul>
            </div>
          </div>

          <!-- Prayer Ministry -->
          <div
            class="bg-white rounded-lg shadow-lg overflow-hidden hover:shadow-xl hover:scale-105 transition-all duration-300 ease-out"
          >
            <div
              class="h-48 bg-cover bg-center"
              style="
                background-image: url('/images/prayer-ministry.jpg');
                background-position: center 30%;
              "
            ></div>
            <div class="p-6">
              <h3 class="text-2xl font-bold text-gray-900 mb-3">
                Prayer Ministry
              </h3>
              <p class="text-gray-600 mb-4">
                Supporting our church and community through prayer,
                intercession, and spiritual encouragement.
              </p>
              <ul class="text-sm text-gray-600 space-y-1">
                <li>• Wednesday Prayer Meeting</li>
                <li>• Prayer Chain</li>
                <li>• Prayer Requests</li>
                <li>• Prayer Counseling</li>
              </ul>
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- Call to Action -->
    <section class="py-20 bg-gray-100">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 text-center">
        <h2 class="text-3xl md:text-4xl font-bold text-gray-900 mb-6">
          Join Us This <span class="text-blue-600">Sunday</span>
        </h2>
        <p class="text-xl text-gray-600 mb-8 max-w-3xl mx-auto">
          We'd love to welcome you to our church family. Come experience the
          love of Christ in our community.
        </p>
        <div class="flex flex-col sm:flex-row gap-4 justify-center">
          <button
            @click="$emit('navigate', 'contact')"
            class="bg-blue-600 text-white px-6 py-3 rounded-lg text-base font-semibold hover:bg-blue-700 transition duration-300 transform hover:scale-105 cursor-pointer"
          >
            Get Directions
          </button>
          <button
            @click="$emit('navigate', 'about')"
            class="border-2 border-blue-600 text-blue-600 px-6 py-3 rounded-lg text-base font-semibold hover:bg-blue-600 hover:text-white transition duration-300 transform hover:scale-105 cursor-pointer"
          >
            Learn More About Us
          </button>
        </div>
      </div>
    </section>
  </div>
</template>

<script>
export default {
  name: 'ServicesPage',
};
</script>
