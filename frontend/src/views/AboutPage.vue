<template>
  <div>
    <!-- Hero Section -->
    <div
      class="relative h-96 md:h-[500px] lg:h-[600px] flex items-center justify-center pt-24 lg:pt-32 bg-[position:center_50%] md:bg-[position:center_40%] lg:bg-[position:center_30%] lg:bg-[length:120%_auto]"
      style="
        background: linear-gradient(rgba(0, 0, 0, 0.5), rgba(0, 0, 0, 0.5)),
          url('/images/community.jpg') no-repeat center 30%;
        background-size: cover;
      "
    >
      <div
        class="relative max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 text-center flex flex-col items-center justify-center h-full"
      >
        <h1
          class="text-3xl md:text-5xl font-extrabold text-white tracking-tight mb-4"
        >
          About Us
        </h1>
        <p class="text-lg md:text-xl text-gray-200/95 max-w-3xl mx-auto">
          Learn more about our church, our mission, and our community.
        </p>
      </div>
    </div>

    <!-- Church Info Section -->
    <section class="pt-16 lg:pt-16 pb-12 bg-white">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="max-w-4xl mx-auto">
          <h2
            class="text-2xl md:text-3xl font-bold text-gray-900 mb-2 text-center"
          >
            Our <span class="text-blue-600">Mission</span>
          </h2>
          <div class="h-1 w-16 bg-blue-600 rounded mb-6 mx-auto"></div>
          <p
            class="text-base md:text-lg text-gray-700 leading-8 mb-4 text-center"
          >
            At {{ churchSettings.church_name }}, our mission is to glorify God
            by making disciples of Jesus Christ through worship, fellowship, and
            service to our community.
          </p>
          <p
            class="text-base md:text-lg text-gray-700 leading-8 mb-4 text-center"
          >
            We believe in creating a welcoming environment where everyone can
            grow in their faith, build meaningful relationships, and discover
            their God-given purpose.
          </p>
          <p class="text-base md:text-lg text-gray-700 leading-8 text-center">
            Our church is a place where you can come as you are, experience
            God's love, and be part of a caring community that supports you on
            your spiritual journey.
          </p>
        </div>
      </div>
    </section>

    <!-- Our Values Section -->
    <section class="py-12 bg-gradient-to-b from-white to-gray-50">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="text-center">
          <h2 class="text-2xl md:text-3xl font-bold text-gray-900 mb-2">
            Our <span class="text-blue-600">Values</span>
          </h2>
          <div class="h-1 w-16 bg-blue-600 rounded mb-6 mx-auto"></div>
          <p class="text-gray-600 max-w-3xl mx-auto">
            Gentle, welcoming, and Christ-centered—these shape everything we do.
          </p>
        </div>
        <div class="grid md:grid-cols-3 gap-4 md:gap-6 p-3 md:p-4">
          <div
            class="rounded-2xl border border-gray-100 bg-white/70 backdrop-blur-sm shadow-sm p-2"
          >
            <h3 class="text-xl font-semibold text-gray-900 mb-1">Worship</h3>
            <p class="text-gray-600 text-base leading-snug">
              We delight in God's presence with reverence and joy.
            </p>
          </div>
          <div
            class="rounded-2xl border border-gray-100 bg-white/70 backdrop-blur-sm shadow-sm p-2"
          >
            <h3 class="text-xl font-semibold text-gray-900 mb-1">Community</h3>
            <p class="text-gray-600 text-base leading-snug">
              We care for one another with grace, humility, and love.
            </p>
          </div>
          <div
            class="rounded-2xl border border-gray-100 bg-white/70 backdrop-blur-sm shadow-sm p-2"
          >
            <h3 class="text-xl font-semibold text-gray-900 mb-1">Service</h3>
            <p class="text-gray-600 text-base leading-snug">
              We serve our city with compassion and generous hearts.
            </p>
          </div>
        </div>
      </div>
    </section>

    <!-- What We Believe Section -->
    <section class="py-12 bg-gray-50">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="text-center mb-2">
          <h2 class="text-xl md:text-2xl font-bold text-gray-900 mb-2">
            What We <span class="text-blue-600">Believe</span>
          </h2>
          <div class="h-1 w-16 bg-blue-600 rounded mb-4 mx-auto"></div>
          <p class="text-base text-gray-600 max-w-3xl mx-auto">
            Our faith is built on the foundational truths of Christianity.
          </p>
        </div>

        <div class="grid md:grid-cols-2 lg:grid-cols-3 gap-4 md:gap-6 p-4">
          <div class="bg-white p-2 rounded-lg shadow-sm">
            <h3 class="text-xl font-bold text-gray-900 mb-3">The Bible</h3>
            <p class="text-gray-600">
              We believe the Bible is the inspired Word of God and our ultimate
              authority for faith and life.
            </p>
          </div>

          <div class="bg-white p-2 rounded-lg shadow-sm">
            <h3 class="text-xl font-bold text-gray-900 mb-3">God</h3>
            <p class="text-gray-600">
              We believe in one God, eternally existing in three persons:
              Father, Son, and Holy Spirit.
            </p>
          </div>

          <div class="bg-white p-2 rounded-lg shadow-sm">
            <h3 class="text-xl font-bold text-gray-900 mb-3">Jesus Christ</h3>
            <p class="text-gray-600">
              We believe Jesus Christ is the Son of God, who died for our sins
              and rose again for our salvation.
            </p>
          </div>

          <div class="bg-white p-2 rounded-lg shadow-sm">
            <h3 class="text-xl font-bold text-gray-900 mb-3">Salvation</h3>
            <p class="text-gray-600">
              We believe salvation is a gift of God's grace, received through
              faith in Jesus Christ alone.
            </p>
          </div>

          <div class="bg-white p-2 rounded-lg shadow-sm">
            <h3 class="text-xl font-bold text-gray-900 mb-3">The Church</h3>
            <p class="text-gray-600">
              We believe the Church is the body of Christ, called to worship,
              grow, and serve together.
            </p>
          </div>

          <div class="bg-white p-2 rounded-lg shadow-sm">
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
    <section class="py-12 bg-white">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 text-center">
        <h2 class="text-xl md:text-2xl font-bold text-gray-900 mb-2">
          Visit Us <span class="text-blue-600">This Sunday</span>
        </h2>
        <div class="h-1 w-16 bg-blue-600 rounded mb-6 mx-auto"></div>
        <p class="text-base text-gray-600 mb-8 max-w-3xl mx-auto">
          We'd love to meet you! Join us for worship and experience our
          welcoming community.
        </p>
        <div class="flex flex-col sm:flex-row gap-4 justify-center">
          <button
            @click="$router.push('/contact')"
            class="border border-blue-600 text-blue-600 px-6 py-3 rounded-lg text-base font-semibold hover:bg-blue-600 hover:text-white transition duration-300 cursor-pointer"
          >
            Get Directions
          </button>
          <button
            @click="$router.push('/our-ministries')"
            class="border border-blue-600 text-blue-600 px-6 py-3 rounded-lg text-base font-semibold hover:bg-blue-600 hover:text-white transition duration-300 cursor-pointer"
          >
            Our Ministries
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
