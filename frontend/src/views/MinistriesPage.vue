<template>
  <div>
    <!-- Hero Section -->
    <div
      class="relative h-[28rem] md:h-[500px] lg:h-[600px] flex items-center justify-center pt-24 lg:pt-32"
      style="
        background: linear-gradient(rgba(0, 0, 0, 0.5), rgba(0, 0, 0, 0.5)),
          url('/images/our-ministries.jpg') no-repeat center center;
        background-size: cover;
      "
    >
      <div
        class="relative max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 text-center flex flex-col items-center justify-center h-full"
      >
        <h1
          class="text-5xl md:text-7xl xl:text-8xl font-extrabold text-white tracking-tight mb-6"
        >
          Our Ministries
        </h1>
        <p class="text-xl md:text-2xl text-white max-w-3xl mx-auto">
          Join us for worship, fellowship, and spiritual growth. Find your place
          in our community where faith comes alive through meaningful
          connections, inspiring worship, and opportunities to serve others in
          love.
        </p>
      </div>
    </div>

    <!-- Ministries Section -->
    <section class="py-20 bg-white">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="text-center mb-16">
          <h2
            class="text-3xl md:text-4xl lg:text-5xl font-bold text-gray-900 max-w-3xl mx-auto"
          >
            Discover opportunities to grow, serve, and connect in our church
            community.
          </h2>
        </div>

        <div class="grid md:grid-cols-2 lg:grid-cols-3 gap-8">
          <!-- Children's Ministry -->
          <div
            class="bg-white rounded-lg shadow-lg overflow-hidden hover:shadow-xl transition-shadow duration-300 ease-out"
          >
            <div
              class="h-48 bg-cover bg-center"
              style="
                background-image: url('/images/children-ministry.jpg');
                background-position: center 30%;
              "
            ></div>
            <div class="p-6">
              <h3 class="text-xl font-bold text-gray-900 mb-3">
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
            class="bg-white rounded-lg shadow-lg overflow-hidden hover:shadow-xl transition-shadow duration-300 ease-out"
          >
            <div
              class="h-48 bg-cover bg-center"
              style="
                background-image: url('/images/youth-ministry.jpg');
                background-position: center 25%;
              "
            ></div>
            <div class="p-6">
              <h3 class="text-xl font-bold text-gray-900 mb-3">
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
            class="bg-white rounded-lg shadow-lg overflow-hidden hover:shadow-xl transition-shadow duration-300 ease-out"
          >
            <div
              class="h-48 bg-cover bg-center"
              style="
                background-image: url('/images/adult-bible-study.jpg');
                background-position: center 30%;
              "
            ></div>
            <div class="p-6">
              <h3 class="text-xl font-bold text-gray-900 mb-3">
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
            class="bg-white rounded-lg shadow-lg overflow-hidden hover:shadow-xl transition-shadow duration-300 ease-out"
          >
            <div
              class="h-48 bg-cover bg-center"
              style="
                background-image: url('/images/worship-ministry.jpg');
                background-position: center 30%;
              "
            ></div>
            <div class="p-6">
              <h3 class="text-xl font-bold text-gray-900 mb-3">
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
            class="bg-white rounded-lg shadow-lg overflow-hidden hover:shadow-xl transition-shadow duration-300 ease-out"
          >
            <div
              class="h-48 bg-cover bg-center"
              style="
                background-image: url('/images/outreach-ministry.jpg');
                background-position: center 30%;
              "
            ></div>
            <div class="p-6">
              <h3 class="text-xl font-bold text-gray-900 mb-3">
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
            class="bg-white rounded-lg shadow-lg overflow-hidden hover:shadow-xl transition-shadow duration-300 ease-out"
          >
            <div
              class="h-48 bg-cover bg-center"
              style="
                background-image: url('/images/prayer-ministry.jpg');
                background-position: center 30%;
              "
            ></div>
            <div class="p-6">
              <h3 class="text-xl font-bold text-gray-900 mb-3">
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
    <section class="py-20 bg-blue-50">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 text-center">
        <h2
          class="text-3xl md:text-4xl lg:text-5xl font-bold text-gray-900 mb-6"
        >
          Join Us This <span class="text-blue-600">Sunday</span>
        </h2>
        <p class="text-base text-gray-600 mb-8 max-w-3xl mx-auto">
          We'd love to welcome you to our church family. Come experience the
          love of Christ in our community.
        </p>
        <div class="flex flex-col sm:flex-row gap-4 justify-center">
          <button
            @click="$router.push('/contact')"
            class="border border-blue-600 text-blue-600 px-6 py-3 rounded-lg text-base font-semibold hover:bg-blue-600 hover:text-white hover:shadow-lg transition-all duration-300 cursor-pointer"
          >
            Get Directions
          </button>
        </div>
      </div>
    </section>
  </div>
</template>

<script>
export default {
  name: 'MinistriesPage',
};
</script>
