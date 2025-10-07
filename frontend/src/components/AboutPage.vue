<template>
  <div>
    <!-- Hero Section -->
    <div
      class="relative h-64 md:h-80 lg:h-96 flex items-center justify-center"
      style="
        background: linear-gradient(rgba(0, 0, 0, 0.5), rgba(0, 0, 0, 0.5)),
          url('/images/church_building_1.jpg') no-repeat center center;
        background-size: cover;
      "
    >
      <div class="relative max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 text-center">
        <h1 class="text-4xl md:text-6xl font-bold text-white mb-6">About Us</h1>
        <p class="text-xl md:text-2xl text-gray-200 max-w-3xl mx-auto">
          Learn more about our church, our mission, and our community.
        </p>
      </div>
    </div>

    <!-- Church Info Section -->
    <section class="py-20 bg-white">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="grid grid-cols-1 md:grid-cols-2 gap-12">
          <div>
            <h2 class="text-3xl md:text-4xl font-bold text-gray-900 mb-6">
              Our <span class="text-blue-600">Mission</span>
            </h2>
            <p class="text-lg text-gray-600 mb-4">
              At {{ churchSettings.church_name }}, our mission is to glorify God
              by making disciples of Jesus Christ through worship, fellowship,
              and service to our community.
            </p>
            <p class="text-lg text-gray-600 mb-4">
              We believe in creating a welcoming environment where everyone can
              grow in their faith, build meaningful relationships, and discover
              their God-given purpose.
            </p>
            <p class="text-lg text-gray-600">
              Our church is a place where you can come as you are, experience
              God's love, and be part of a caring community that supports you on
              your spiritual journey.
            </p>
          </div>

          <div>
            <h3 class="text-2xl font-bold text-gray-900 mb-6">
              Contact Information
            </h3>
            <div class="space-y-4">
              <div class="flex items-start">
                <svg
                  class="w-6 h-6 text-blue-600 mr-3 mt-1"
                  fill="currentColor"
                  viewBox="0 0 20 20"
                >
                  <path
                    fill-rule="evenodd"
                    d="M5.05 4.05a7 7 0 119.9 9.9L10 18.9l-4.95-4.95a7 7 0 010-9.9zM10 11a2 2 0 100-4 2 2 0 000 4z"
                    clip-rule="evenodd"
                  />
                </svg>
                <div>
                  <strong class="text-gray-900">Address:</strong>
                  <p class="text-gray-600">
                    {{ churchSettings.church_address }}
                  </p>
                </div>
              </div>
              <div class="flex items-start">
                <svg
                  class="w-6 h-6 text-blue-600 mr-3 mt-1"
                  fill="currentColor"
                  viewBox="0 0 20 20"
                >
                  <path
                    d="M2 3a1 1 0 011-1h2.153a1 1 0 01.986.836l.74 4.435a1 1 0 01-.54 1.06l-1.548.773a11.037 11.037 0 006.105 6.105l.774-1.548a1 1 0 011.059-.54l4.435.74a1 1 0 01.836.986V17a1 1 0 01-1 1h-2C7.82 18 2 12.18 2 5V3z"
                  />
                </svg>
                <div>
                  <strong class="text-gray-900">Phone:</strong>
                  <p class="text-gray-600">{{ churchSettings.church_phone }}</p>
                </div>
              </div>
              <div class="flex items-start">
                <svg
                  class="w-6 h-6 text-blue-600 mr-3 mt-1"
                  fill="currentColor"
                  viewBox="0 0 20 20"
                >
                  <path
                    d="M2.003 5.884L10 9.882l7.997-3.998A2 2 0 0016 4H4a2 2 0 00-1.997 1.884z"
                  />
                  <path
                    d="M18 8.118l-8 4-8-4V14a2 2 0 002 2h12a2 2 0 002-2V8.118z"
                  />
                </svg>
                <div>
                  <strong class="text-gray-900">Email:</strong>
                  <p class="text-gray-600">{{ churchSettings.church_email }}</p>
                </div>
              </div>
            </div>

            <h3 class="text-2xl font-bold text-gray-900 mt-8 mb-6">
              Service Times
            </h3>
            <div class="space-y-3">
              <div
                v-for="service in serviceTimes"
                :key="service.day"
                class="bg-gray-50 p-4 rounded-lg"
              >
                <div class="flex justify-between items-center">
                  <div>
                    <strong class="text-gray-900">{{ service.day }}</strong>
                    <p class="text-gray-600 text-sm">{{ service.service }}</p>
                  </div>
                  <span class="text-blue-600 font-bold">{{
                    service.time
                  }}</span>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- What We Believe Section -->
    <section class="py-20 bg-gray-50">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="text-center mb-16">
          <h2 class="text-3xl md:text-4xl font-bold text-gray-900 mb-6">
            What We <span class="text-blue-600">Believe</span>
          </h2>
          <p class="text-xl text-gray-600 max-w-3xl mx-auto">
            Our faith is built on the foundational truths of Christianity.
          </p>
        </div>

        <div class="grid md:grid-cols-2 lg:grid-cols-3 gap-8">
          <div class="bg-white p-6 rounded-lg shadow-sm">
            <h3 class="text-xl font-bold text-gray-900 mb-3">The Bible</h3>
            <p class="text-gray-600">
              We believe the Bible is the inspired Word of God and our ultimate
              authority for faith and life.
            </p>
          </div>

          <div class="bg-white p-6 rounded-lg shadow-sm">
            <h3 class="text-xl font-bold text-gray-900 mb-3">God</h3>
            <p class="text-gray-600">
              We believe in one God, eternally existing in three persons:
              Father, Son, and Holy Spirit.
            </p>
          </div>

          <div class="bg-white p-6 rounded-lg shadow-sm">
            <h3 class="text-xl font-bold text-gray-900 mb-3">Jesus Christ</h3>
            <p class="text-gray-600">
              We believe Jesus Christ is the Son of God, who died for our sins
              and rose again for our salvation.
            </p>
          </div>

          <div class="bg-white p-6 rounded-lg shadow-sm">
            <h3 class="text-xl font-bold text-gray-900 mb-3">Salvation</h3>
            <p class="text-gray-600">
              We believe salvation is a gift of God's grace, received through
              faith in Jesus Christ alone.
            </p>
          </div>

          <div class="bg-white p-6 rounded-lg shadow-sm">
            <h3 class="text-xl font-bold text-gray-900 mb-3">The Church</h3>
            <p class="text-gray-600">
              We believe the Church is the body of Christ, called to worship,
              grow, and serve together.
            </p>
          </div>

          <div class="bg-white p-6 rounded-lg shadow-sm">
            <h3 class="text-xl font-bold text-gray-900 mb-3">Eternity</h3>
            <p class="text-gray-600">
              We believe in the resurrection of both the saved and the lost, and
              in eternal life with God.
            </p>
          </div>
        </div>
      </div>
    </section>

    <!-- Call to Action -->
    <section class="py-20 bg-white">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 text-center">
        <h2 class="text-3xl md:text-4xl font-bold text-gray-900 mb-6">
          Visit Us <span class="text-blue-600">This Sunday</span>
        </h2>
        <p class="text-xl text-gray-600 mb-8 max-w-3xl mx-auto">
          We'd love to meet you! Join us for worship and experience our
          welcoming community.
        </p>
        <div class="flex flex-col sm:flex-row gap-4 justify-center">
          <button
            @click="$emit('navigate', 'contact')"
            class="bg-blue-600 text-white px-8 py-4 rounded-lg text-lg font-semibold hover:bg-blue-700 transition duration-300"
          >
            Get Directions
          </button>
          <button
            @click="$emit('navigate', 'services')"
            class="border-2 border-blue-600 text-blue-600 px-8 py-4 rounded-lg text-lg font-semibold hover:bg-blue-600 hover:text-white transition duration-300"
          >
            View Services
          </button>
        </div>
      </div>
    </section>
  </div>
</template>

<script>
export default {
  name: 'AboutPage',
  props: {
    churchSettings: Object,
    serviceTimes: Array,
  },
};
</script>
