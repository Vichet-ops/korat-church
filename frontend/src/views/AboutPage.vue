<template>
  <div>
    <!-- Hero Section -->
    <div
      class="relative h-64 md:h-80 lg:h-96 flex items-center justify-center"
      style="
        background: linear-gradient(
            rgba(11, 31, 58, 0.45),
            rgba(27, 47, 88, 0.55)
          ),
          url('/images/church_building_1.jpg') no-repeat center center;
        background-size: cover;
      "
    >
      <div class="relative max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 text-center">
        <h1
          class="text-4xl md:text-5xl font-extrabold text-white tracking-tight mb-4"
        >
          About Us
        </h1>
        <p class="text-lg md:text-xl text-gray-200/95 max-w-3xl mx-auto">
          Learn more about our church, our mission, and our community.
        </p>
      </div>
    </div>

    <!-- Church Info Section -->
    <section class="pt-24 lg:pt-28 pb-20 bg-white">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="grid grid-cols-1 md:grid-cols-2 gap-12">
          <div>
            <h2 class="text-3xl md:text-4xl font-bold text-gray-900 mb-4">
              Our <span class="text-blue-600">Mission</span>
            </h2>
            <div class="h-1 w-16 bg-blue-600 rounded mb-6"></div>
            <p class="text-base md:text-lg text-gray-700 leading-8 mb-4">
              At {{ churchSettings.church_name }}, our mission is to glorify God
              by making disciples of Jesus Christ through worship, fellowship,
              and service to our community.
            </p>
            <p class="text-base md:text-lg text-gray-700 leading-8 mb-4">
              We believe in creating a welcoming environment where everyone can
              grow in their faith, build meaningful relationships, and discover
              their God-given purpose.
            </p>
            <p class="text-base md:text-lg text-gray-700 leading-8">
              Our church is a place where you can come as you are, experience
              God's love, and be part of a caring community that supports you on
              your spiritual journey.
            </p>
          </div>

          <div>
            <h3 class="text-2xl font-bold text-gray-900 mt-0 mb-4">
              Service Times
            </h3>
            <div class="space-y-3">
              <div
                v-for="service in serviceTimes"
                :key="service.day"
                class="bg-white p-4 rounded-lg border border-gray-100 shadow-sm"
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

    <!-- Our Values Section -->
    <section class="py-20 bg-gradient-to-b from-white to-gray-50">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="text-center mb-12">
          <h2 class="text-3xl md:text-4xl font-bold text-gray-900 mb-4">
            Our <span class="text-blue-600">Values</span>
          </h2>
          <p class="text-gray-600 max-w-3xl mx-auto">
            Gentle, welcoming, and Christ-centered—these shape everything we do.
          </p>
        </div>
        <div class="grid md:grid-cols-3 gap-6">
          <div
            class="rounded-2xl border border-gray-100 bg-white/70 backdrop-blur-sm p-6 shadow-sm hover:shadow-md transition"
          >
            <h3 class="text-xl font-semibold text-gray-900 mb-2">Worship</h3>
            <p class="text-gray-600">
              We delight in God’s presence with reverence and joy.
            </p>
          </div>
          <div
            class="rounded-2xl border border-gray-100 bg-white/70 backdrop-blur-sm p-6 shadow-sm hover:shadow-md transition"
          >
            <h3 class="text-xl font-semibold text-gray-900 mb-2">Community</h3>
            <p class="text-gray-600">
              We care for one another with grace, humility, and love.
            </p>
          </div>
          <div
            class="rounded-2xl border border-gray-100 bg-white/70 backdrop-blur-sm p-6 shadow-sm hover:shadow-md transition"
          >
            <h3 class="text-xl font-semibold text-gray-900 mb-2">Service</h3>
            <p class="text-gray-600">
              We serve our city with compassion and generous hearts.
            </p>
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
            @click="$router.push('/contact')"
            class="bg-blue-600 text-white px-6 py-3 rounded-lg text-base font-semibold hover:bg-blue-700 transition duration-300 cursor-pointer"
          >
            Get Directions
          </button>
          <button
            @click="$router.push('/services')"
            class="border-2 border-blue-600 text-blue-600 px-6 py-3 rounded-lg text-base font-semibold hover:bg-blue-600 hover:text-white transition duration-300 cursor-pointer"
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
