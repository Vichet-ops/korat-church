<template>
  <div>
    <!-- Hero Section -->
    <div
      class="relative h-64 md:h-80 lg:h-96 flex items-center justify-center"
      style="
        background: linear-gradient(rgba(0, 0, 0, 0.5), rgba(0, 0, 0, 0.5)),
          url('/images/give.jpg') no-repeat center center;
        background-size: cover;
      "
    >
      <div class="relative max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 text-center">
        <h1 class="text-4xl md:text-6xl font-bold text-white mb-6">
          Support Our Ministry
        </h1>
        <p class="text-xl md:text-2xl text-gray-200 max-w-3xl mx-auto">
          Your generous giving enables us to serve our community and spread
          God's love through various ministries and outreach programs.
        </p>
      </div>
    </div>

    <!-- Why Give Section -->
    <section class="py-20 bg-white">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="text-center mb-16">
          <h2 class="text-3xl md:text-4xl font-bold text-gray-900 mb-6">
            Why <span class="text-blue-600">Give?</span>
          </h2>
          <p class="text-xl text-gray-600 max-w-3xl mx-auto">
            Your donations make a real difference in our community and help us
            fulfill our mission to serve God and others.
          </p>
        </div>

        <div class="grid md:grid-cols-2 lg:grid-cols-4 gap-8">
          <!-- Support Our Mission -->
          <div
            class="bg-gradient-to-br from-red-50 to-red-100 rounded-xl shadow-lg p-8 text-center hover:shadow-2xl hover:scale-105 transition-all duration-300 border-2 border-red-200"
          >
            <div
              class="w-16 h-16 bg-gradient-to-br from-red-500 to-red-600 rounded-full flex items-center justify-center mx-auto mb-6 shadow-lg"
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
            <h3 class="text-xl font-bold text-red-800 mb-4">
              Support Our Mission
            </h3>
            <p class="text-gray-600">
              Help us spread the love of Christ and serve our community with
              compassion and faith.
            </p>
          </div>

          <!-- Community Outreach -->
          <div
            class="bg-gradient-to-br from-green-50 to-green-100 rounded-xl shadow-lg p-8 text-center hover:shadow-2xl hover:scale-105 transition-all duration-300 border-2 border-green-200"
          >
            <div
              class="w-16 h-16 bg-gradient-to-br from-green-500 to-green-600 rounded-full flex items-center justify-center mx-auto mb-6 shadow-lg"
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
            <h3 class="text-xl font-bold text-green-800 mb-4">
              Community Outreach
            </h3>
            <p class="text-gray-600">
              Fund our programs that help those in need, including food banks,
              counseling, and support services.
            </p>
          </div>

          <!-- Church Operations -->
          <div
            class="bg-gradient-to-br from-blue-50 to-blue-100 rounded-xl shadow-lg p-8 text-center hover:shadow-2xl hover:scale-105 transition-all duration-300 border-2 border-blue-200"
          >
            <div
              class="w-16 h-16 bg-gradient-to-br from-blue-500 to-blue-600 rounded-full flex items-center justify-center mx-auto mb-6 shadow-lg"
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
                  d="M19 21V5a2 2 0 00-2-2H7a2 2 0 00-2 2v16m14 0h2m-2 0h-5m-9 0H3m2 0h5M9 7h1m-1 4h1m4-4h1m-1 4h1m-5 10v-5a1 1 0 011-1h2a1 1 0 011 1v5m-4 0h4"
                ></path>
              </svg>
            </div>
            <h3 class="text-xl font-bold text-blue-800 mb-4">
              Church Operations
            </h3>
            <p class="text-gray-600">
              Support our daily operations, facility maintenance, and
              administrative needs.
            </p>
          </div>

          <!-- Youth & Education -->
          <div
            class="bg-gradient-to-br from-purple-50 to-purple-100 rounded-xl shadow-lg p-8 text-center hover:shadow-2xl hover:scale-105 transition-all duration-300 border-2 border-purple-200"
          >
            <div
              class="w-16 h-16 bg-gradient-to-br from-purple-500 to-purple-600 rounded-full flex items-center justify-center mx-auto mb-6 shadow-lg"
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
            <h3 class="text-xl font-bold text-purple-800 mb-4">
              Youth & Education
            </h3>
            <p class="text-gray-600">
              Help fund our youth programs, Bible studies, and educational
              initiatives for all ages.
            </p>
          </div>
        </div>
      </div>
    </section>

    <!-- Ways to Give Section -->
    <section class="py-20 bg-gray-50">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="text-center mb-16">
          <h2 class="text-3xl md:text-4xl font-bold text-gray-900 mb-6">
            Ways to <span class="text-blue-600">Give</span>
          </h2>
          <p class="text-xl text-gray-600 max-w-3xl mx-auto">
            Choose the method that works best for you. All donations are
            tax-deductible and greatly appreciated.
          </p>
        </div>

        <div class="grid lg:grid-cols-2 gap-12">
          <!-- Online Giving -->
          <div
            class="bg-white rounded-2xl shadow-xl p-8 border border-gray-100"
          >
            <div class="text-center mb-8">
              <div
                class="w-16 h-16 bg-gray-100 rounded-full flex items-center justify-center mx-auto mb-4"
              >
                <svg
                  class="w-8 h-8 text-gray-600"
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
                    class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                    placeholder="$0.00"
                  />
                </div>
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2"
                    >Category</label
                  >
                  <select
                    v-model="donationForm.category"
                    class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                  >
                    <option value="general">General Fund</option>
                    <option value="missions">Missions</option>
                    <option value="youth">Youth Ministry</option>
                    <option value="building">Building Fund</option>
                  </select>
                </div>
                <button
                  type="submit"
                  class="w-full bg-blue-600 text-white py-3 px-6 rounded-lg font-semibold hover:bg-blue-700 transition duration-300"
                >
                  Proceed to Payment
                </button>
              </form>
            </div>
          </div>

          <!-- Other Giving Methods -->
          <div class="space-y-6">
            <!-- Mail Check -->
            <div
              class="bg-white rounded-xl shadow-lg p-6 border-l-4 border-purple-600 hover:shadow-xl transition-shadow duration-300"
            >
              <div class="flex items-start">
                <div
                  class="flex-shrink-0 w-12 h-12 bg-purple-100 rounded-lg flex items-center justify-center"
                >
                  <svg
                    class="w-6 h-6 text-purple-600"
                    fill="none"
                    stroke="currentColor"
                    viewBox="0 0 24 24"
                  >
                    <path
                      stroke-linecap="round"
                      stroke-linejoin="round"
                      stroke-width="2"
                      d="M3 8l7.89 4.26a2 2 0 002.22 0L21 8M5 19h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z"
                    ></path>
                  </svg>
                </div>
                <div class="ml-4">
                  <h3 class="text-lg font-semibold text-gray-900 mb-2">
                    Mail a Check
                  </h3>
                  <p class="text-gray-600 mb-2">Send your donation to:</p>
                  <p class="text-gray-600 font-medium">
                    Muang Thai Korat Church<br />
                    123 Main St<br />
                    Nakhon Ratchasima, Thailand 30000
                  </p>
                </div>
              </div>
            </div>

            <!-- Bank Transfer -->
            <div
              class="bg-white rounded-xl shadow-lg p-6 border-l-4 border-green-600 hover:shadow-xl transition-shadow duration-300"
            >
              <div class="flex items-start">
                <div
                  class="flex-shrink-0 w-12 h-12 bg-green-100 rounded-lg flex items-center justify-center"
                >
                  <svg
                    class="w-6 h-6 text-green-600"
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
                <div class="ml-4">
                  <h3 class="text-lg font-semibold text-gray-900 mb-2">
                    Bank Transfer
                  </h3>
                  <p class="text-gray-600">
                    Contact the church office for bank account details for
                    direct transfers.
                  </p>
                </div>
              </div>
            </div>

            <!-- Sunday Offering -->
            <div
              class="bg-white rounded-xl shadow-lg p-6 border-l-4 border-orange-600 hover:shadow-xl transition-shadow duration-300"
            >
              <div class="flex items-start">
                <div
                  class="flex-shrink-0 w-12 h-12 bg-orange-100 rounded-lg flex items-center justify-center"
                >
                  <svg
                    class="w-6 h-6 text-orange-600"
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
                <div class="ml-4">
                  <h3 class="text-lg font-semibold text-gray-900 mb-2">
                    Sunday Offering
                  </h3>
                  <p class="text-gray-600">
                    Place your donation in the offering plate during our Sunday
                    service.
                  </p>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- Call to Action Section -->
    <section class="py-20 bg-white">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 text-center">
        <h2 class="text-3xl md:text-4xl font-bold text-gray-900 mb-6">
          Make a Difference <span class="text-blue-600">Today</span>
        </h2>
        <p class="text-xl text-gray-600 mb-8 max-w-3xl mx-auto">
          Your generous giving enables us to serve our community, support
          families, and spread God's love. Every gift, no matter the size, makes
          a real impact.
        </p>
        <button
          @click="$emit('navigate', 'contact')"
          class="border-2 border-blue-600 text-blue-600 px-8 py-4 rounded-lg text-lg font-semibold hover:bg-blue-600 hover:text-white transition duration-300 transform hover:scale-105"
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
      // TODO: Implement actual payment processing (e.g., PayPal, Stripe)
      alert(
        `Processing donation of $${this.donationForm.amount} to ${this.donationForm.category}`
      );
      console.log('Donation details:', this.donationForm);
    },
  },
};
</script>
