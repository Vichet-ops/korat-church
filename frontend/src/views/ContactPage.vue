<template>
  <div>
    <!-- Hero Section -->
    <div
      class="relative h-[28rem] md:h-[500px] lg:h-[600px] flex items-center justify-center pt-24 lg:pt-32 bg-[position:center_50%] md:bg-[position:center_40%] lg:bg-[position:center_45%]"
      style="
        background: linear-gradient(rgba(0, 0, 0, 0.5), rgba(0, 0, 0, 0.5)),
          url('/images/contact-us.png') no-repeat center center;
        background-size: cover;
      "
    >
      <div
        class="relative max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 text-center flex flex-col items-center justify-center h-full"
      >
        <h1
          class="text-5xl md:text-7xl xl:text-8xl font-extrabold text-white tracking-tight mb-6"
        >
          Contact Us
        </h1>
        <p class="text-xl md:text-2xl text-white max-w-3xl mx-auto">
          We'd love to hear from you. Get in touch with us today!
        </p>
      </div>
    </div>

    <!-- Contact Section -->
    <section class="py-20 bg-white">
      <div class="max-w-4xl mx-auto px-4 sm:px-6 lg:px-8">
        <!-- Contact Form -->
        <div class="bg-gray-50 p-8 rounded-lg">
          <h3 class="text-2xl md:text-3xl font-bold text-gray-900 mb-6">
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
                class="w-full px-4 py-2 border border-gray-300 rounded-lg outline-none"
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
                class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:outline-none"
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
                class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:outline-none"
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
                class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:outline-none"
                placeholder="Your message"
              ></textarea>
            </div>

            <button
              type="submit"
              class="w-full border border-blue-600 text-blue-600 py-3 px-6 rounded-lg font-semibold hover:bg-blue-600 hover:text-white outline-none transition duration-300 cursor-pointer"
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
    </section>

    <!-- Map Section -->
    <section class="py-20 bg-blue-50">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="text-center mb-12">
          <h2
            class="text-3xl md:text-4xl lg:text-5xl font-bold text-gray-900 mb-6"
          >
            Visit <span class="text-blue-600">Us</span>
          </h2>
          <p class="text-xl text-gray-600 max-w-3xl mx-auto">
            We're located in the heart of Nakhon Ratchasima and would love to
            see you!
          </p>
        </div>
        <div class="rounded-lg overflow-hidden shadow">
          <iframe
            :src="embedMapSrc"
            width="100%"
            height="400"
            style="border: 0"
            allowfullscreen=""
            loading="lazy"
            referrerpolicy="no-referrer-when-downgrade"
          ></iframe>
        </div>
        <div class="text-center mt-4">
          <a
            :href="streetViewLink"
            target="_blank"
            rel="noopener"
            class="text-blue-600 hover:text-blue-700 underline-offset-2 hover:underline"
          >
            Open Street View
          </a>
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
      try {
        // Get API URL from environment or use default
        const apiUrl = import.meta.env.VITE_API_URL || 'https://vichetkeo.com';

        // Send message to backend
        const response = await fetch(`${apiUrl}/api/contact/send`, {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify(this.contactForm),
        });

        const data = await response.json();

        if (!response.ok) {
          throw new Error(data.error || 'Failed to send message');
        }

        // Show success message
        this.messageSent = true;

        // Reset form after delay
        setTimeout(() => {
          this.messageSent = false;
          this.contactForm = { name: '', email: '', subject: '', message: '' };
        }, 5000);
      } catch (error) {
        alert(
          'Sorry, there was an error sending your message. Please try again or contact us directly.'
        );
      }
    },
  },
  computed: {
    googleMapsLink() {
      // Use the precise Street View URL provided by the user
      return 'https://www.google.com/maps/place/Muangthai+Korat+Church/@14.9592904,102.0389158,3a,75y,275.19h,90t/data=!3m7!1e1!3m5!1s423Hso2quX9FJ4xNDGLhDw!2e0!6shttps:%2F%2Fstreetviewpixels-pa.googleapis.com%2Fv1%2Fthumbnail%3Fcb_client%3Dmaps_sv.tactile%26w%3D900%26h%3D600%26pitch%3D0%26panoid%3D423Hso2quX9FJ4xNDGLhDw%26yaw%3D275.19135!7i16384!8i8192!4m14!1m7!3m6!1s0x311eb31037797a4f:0x6e1c1237a91516b9!2sMuangthai+Korat+Church!8m2!3d14.9593033!4d102.0387701!16s%2Fg%2F1hm4whygs!3m5!1s0x311eb31037797a4f:0x6e1c1237a91516b9!8m2!3d14.9593033!4d102.0387701!16s%2Fg%2F1hm4whygs?authuser=0&entry=ttu&g_ep=EgoyMDI1MTAwMS4wIKXMDSoASAFQAw%3D%3D';
    },
    embedMapSrc() {
      // Convert the Street View place URL to an embeddable maps URL
      // This uses a generic embed that opens the given place; for precise embeds, a Maps Embed API key can be used.
      return 'https://www.google.com/maps?q=Muangthai+Korat+Church&output=embed';
    },
    streetViewLink() {
      return this.googleMapsLink;
    },
  },
};
</script>
