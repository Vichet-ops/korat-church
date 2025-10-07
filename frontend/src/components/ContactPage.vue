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
        <h1 class="text-4xl md:text-6xl font-bold text-white mb-6">
          Contact Us
        </h1>
        <p class="text-xl md:text-2xl text-gray-200 max-w-3xl mx-auto">
          We'd love to hear from you. Get in touch with us today!
        </p>
      </div>
    </div>

    <!-- Contact Section -->
    <section class="py-20 bg-white">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="grid grid-cols-1 md:grid-cols-2 gap-12">
          <!-- Contact Information -->
          <div>
            <h2 class="text-3xl font-bold text-gray-900 mb-6">Get In Touch</h2>
            <p class="text-gray-600 mb-8">
              Whether you have questions, prayer requests, or just want to learn
              more about our church, we're here for you.
            </p>

            <div class="space-y-6">
              <div class="flex items-start">
                <div
                  class="flex-shrink-0 w-12 h-12 bg-blue-100 rounded-lg flex items-center justify-center"
                >
                  <svg
                    class="w-6 h-6 text-blue-600"
                    fill="currentColor"
                    viewBox="0 0 20 20"
                  >
                    <path
                      fill-rule="evenodd"
                      d="M5.05 4.05a7 7 0 119.9 9.9L10 18.9l-4.95-4.95a7 7 0 010-9.9zM10 11a2 2 0 100-4 2 2 0 000 4z"
                      clip-rule="evenodd"
                    />
                  </svg>
                </div>
                <div class="ml-4">
                  <h3 class="text-lg font-semibold text-gray-900 mb-2">
                    Address
                  </h3>
                  <p class="text-gray-600">
                    {{ churchSettings.church_address }}
                  </p>
                </div>
              </div>

              <div class="flex items-start">
                <div
                  class="flex-shrink-0 w-12 h-12 bg-green-100 rounded-lg flex items-center justify-center"
                >
                  <svg
                    class="w-6 h-6 text-green-600"
                    fill="currentColor"
                    viewBox="0 0 20 20"
                  >
                    <path
                      d="M2 3a1 1 0 011-1h2.153a1 1 0 01.986.836l.74 4.435a1 1 0 01-.54 1.06l-1.548.773a11.037 11.037 0 006.105 6.105l.774-1.548a1 1 0 011.059-.54l4.435.74a1 1 0 01.836.986V17a1 1 0 01-1 1h-2C7.82 18 2 12.18 2 5V3z"
                    />
                  </svg>
                </div>
                <div class="ml-4">
                  <h3 class="text-lg font-semibold text-gray-900 mb-2">
                    Phone
                  </h3>
                  <p class="text-gray-600">{{ churchSettings.church_phone }}</p>
                </div>
              </div>

              <div class="flex items-start">
                <div
                  class="flex-shrink-0 w-12 h-12 bg-purple-100 rounded-lg flex items-center justify-center"
                >
                  <svg
                    class="w-6 h-6 text-purple-600"
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
                </div>
                <div class="ml-4">
                  <h3 class="text-lg font-semibold text-gray-900 mb-2">
                    Email
                  </h3>
                  <p class="text-gray-600">{{ churchSettings.church_email }}</p>
                </div>
              </div>

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
                    Office Hours
                  </h3>
                  <p class="text-gray-600">
                    Monday - Friday: 9:00 AM - 5:00 PM<br />
                    Saturday - Sunday: By Appointment
                  </p>
                </div>
              </div>
            </div>
          </div>

          <!-- Contact Form -->
          <div class="bg-gray-50 p-8 rounded-lg">
            <h3 class="text-2xl font-bold text-gray-900 mb-6">
              Send us a Message
            </h3>
            <form @submit.prevent="sendMessage" class="space-y-6">
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2"
                  >Name</label
                >
                <input
                  v-model="contactForm.name"
                  type="text"
                  required
                  class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                  placeholder="Your name"
                />
              </div>

              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2"
                  >Email</label
                >
                <input
                  v-model="contactForm.email"
                  type="email"
                  required
                  class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                  placeholder="your@email.com"
                />
              </div>

              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2"
                  >Subject</label
                >
                <input
                  v-model="contactForm.subject"
                  type="text"
                  required
                  class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                  placeholder="Message subject"
                />
              </div>

              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2"
                  >Message</label
                >
                <textarea
                  v-model="contactForm.message"
                  rows="5"
                  required
                  class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                  placeholder="Your message"
                ></textarea>
              </div>

              <button
                type="submit"
                class="w-full bg-blue-600 text-white py-3 px-6 rounded-lg font-semibold hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 transition duration-300"
              >
                Send Message
              </button>
            </form>
            <div
              v-if="messageSent"
              class="mt-4 p-4 bg-green-100 text-green-700 rounded-lg"
            >
              Message sent successfully! We'll get back to you soon.
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- Map Section (Optional) -->
    <section class="py-20 bg-gray-100">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="text-center mb-12">
          <h2 class="text-3xl md:text-4xl font-bold text-gray-900 mb-6">
            Visit <span class="text-blue-600">Us</span>
          </h2>
          <p class="text-xl text-gray-600">
            We're located in the heart of Nakhon Ratchasima and would love to
            see you!
          </p>
        </div>
        <div
          class="bg-gray-300 h-96 rounded-lg flex items-center justify-center"
        >
          <p class="text-gray-600">
            Map placeholder - Integrate Google Maps here
          </p>
        </div>
      </div>
    </section>
  </div>
</template>

<script>
export default {
  name: 'ContactPage',
  props: {
    churchSettings: Object,
  },
  data() {
    return {
      contactForm: {
        name: '',
        email: '',
        subject: '',
        message: '',
      },
      messageSent: false,
    };
  },
  methods: {
    async sendMessage() {
      // TODO: Implement actual message sending logic
      console.log('Sending message:', this.contactForm);
      this.messageSent = true;
      setTimeout(() => {
        this.messageSent = false;
        this.contactForm = { name: '', email: '', subject: '', message: '' };
      }, 3000);
    },
  },
};
</script>
