<template>
  <div>
    <!-- Hero Section -->
    <div
      class="relative h-[28rem] md:h-[500px] lg:h-[600px] flex items-center justify-center pt-24 lg:pt-32"
      style="
        background: linear-gradient(rgba(0, 0, 0, 0.5), rgba(0, 0, 0, 0.5)),
          url('/images/give.jpg') no-repeat center center;
        background-size: cover;
      "
    >
      <div
        class="relative max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 text-center flex flex-col items-center justify-center h-full"
      >
        <h1
          class="text-5xl md:text-7xl xl:text-8xl font-extrabold text-white tracking-tight mb-6"
        >
          Support Our Ministry
        </h1>
        <p class="text-xl md:text-2xl text-white max-w-3xl mx-auto">
          Your generous giving enables us to serve our community and spread
          God's love through various ministries and outreach programs.
        </p>
      </div>
    </div>

    <!-- Why Give Section -->
    <section class="py-20 bg-white">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="text-center mb-16">
          <h2
            class="text-3xl md:text-4xl lg:text-5xl font-bold text-gray-900 mb-6"
          >
            Why <span class="text-blue-600">Give?</span>
          </h2>
          <p class="text-xl text-gray-600 max-w-3xl mx-auto">
            Your donations make a real difference in our community and help us
            fulfill our mission to serve God and others.
          </p>
        </div>

        <div class="grid md:grid-cols-2 lg:grid-cols-4 gap-8">
          <!-- Support Our Mission -->
          <div class="bg-white rounded-xl shadow-lg p-8">
            <h3 class="text-xl font-bold text-gray-900 mb-4">
              Support Our Mission
            </h3>
            <p class="text-gray-600 text-left">
              Help us spread the love of Christ and serve our community with
              compassion and faith.
            </p>
          </div>

          <!-- Community Outreach -->
          <div class="bg-white rounded-xl shadow-lg p-8">
            <h3 class="text-xl font-bold text-gray-900 mb-4">
              Community Outreach
            </h3>
            <p class="text-gray-600 text-left">
              Fund our programs that help those in need, including food banks,
              counseling, and support services.
            </p>
          </div>

          <!-- Church Operations -->
          <div class="bg-white rounded-xl shadow-lg p-8">
            <h3 class="text-xl font-bold text-gray-900 mb-4">
              Church Operations
            </h3>
            <p class="text-gray-600 text-left">
              Support our daily operations, facility maintenance, and
              administrative needs.
            </p>
          </div>

          <!-- Youth & Education -->
          <div class="bg-white rounded-xl shadow-lg p-8">
            <h3 class="text-xl font-bold text-gray-900 mb-4">
              Youth & Education
            </h3>
            <p class="text-gray-600 text-left">
              Help fund our youth programs, Bible studies, and educational
              initiatives for all ages.
            </p>
          </div>
        </div>
      </div>
    </section>

    <!-- Make a Donation Section -->
    <section class="py-20 bg-gray-50">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="text-center mb-16">
          <h2
            class="text-3xl md:text-4xl lg:text-5xl font-bold text-gray-900 mb-6"
          >
            Make a <span class="text-blue-600">Donation</span>
          </h2>
          <p class="text-xl text-gray-600 max-w-3xl mx-auto">
            Your generous contribution is tax-deductible and greatly
            appreciated. Thank you for supporting our ministry.
          </p>
        </div>

        <div class="max-w-2xl mx-auto">
          <!-- Online Giving -->
          <div
            class="bg-white rounded-2xl shadow-xl p-8 border border-gray-100"
          >
            <div class="text-center mb-8">
              <div
                class="w-16 h-16 bg-blue-100 rounded-full flex items-center justify-center mx-auto mb-4"
              >
                <svg
                  class="w-8 h-8 text-blue-600"
                  fill="none"
                  stroke="currentColor"
                  viewBox="0 0 24 24"
                >
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M3 10h18M7 15h1m4 0h1m-7 4h12a3 3 0 003-3V8a3 3 0 00-3-3H6a3 3 0 00-3 3v8a3 3 0 003 3z"
                  ></path>
                </svg>
              </div>
              <h3 class="text-2xl font-bold text-gray-900 mb-4">
                Online Giving
              </h3>
              <p class="text-gray-600">
                Make a secure donation online. Works worldwide with all major
                credit cards and payment methods.
              </p>
            </div>

            <!-- Online Giving Form -->
            <div class="bg-gray-50 p-6 rounded-lg">
              <form @submit.prevent="processDonation" class="space-y-4">
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2"
                    >Amount</label
                  >
                  <input
                    v-model="donationForm.amount"
                    type="number"
                    min="1"
                    step="0.01"
                    required
                    class="w-full px-4 py-3 border border-gray-300 rounded-lg outline-none"
                    placeholder="$0.00"
                  />
                </div>
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2"
                    >Category</label
                  >
                  <select
                    v-model="donationForm.category"
                    class="w-full px-4 py-3 border border-gray-300 rounded-lg outline-none"
                  >
                    <option value="general">General Fund</option>
                    <option value="missions">Missions</option>
                    <option value="youth">Youth Ministry</option>
                    <option value="building">Building Fund</option>
                  </select>
                </div>
                <button
                  type="submit"
                  class="w-full bg-blue-600 text-white py-3 px-6 rounded-lg font-semibold hover:bg-blue-700 transition duration-300 cursor-pointer"
                >
                  Proceed to Payment
                </button>
              </form>
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- Call to Action Section -->
    <section class="py-20 bg-blue-50">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 text-center">
        <h2
          class="text-3xl md:text-4xl lg:text-5xl font-bold text-gray-900 mb-6"
        >
          Make a Difference <span class="text-blue-600">Today</span>
        </h2>
        <p class="text-xl text-gray-600 mb-8 max-w-3xl mx-auto">
          Your generous giving enables us to serve our community, support
          families, and spread God's love. Every gift, no matter the size, makes
          a real impact.
        </p>
        <button
          @click="$router.push('/contact')"
          class="border border-blue-600 text-blue-600 px-8 py-3 rounded-lg text-lg font-semibold hover:bg-blue-600 hover:text-white transition duration-300 cursor-pointer"
        >
          Contact Us
        </button>
      </div>
    </section>
  </div>
</template>

<script>
export default {
  name: 'GivePage',
  data() {
    return {
      donationForm: {
        amount: '',
        category: 'general',
      },
    };
  },
  methods: {
    processDonation() {
      // Payment processing will be implemented when needed
      alert(
        `Processing donation of $${this.donationForm.amount} to ${this.donationForm.category}`
      );
    },
  },
};
</script>
