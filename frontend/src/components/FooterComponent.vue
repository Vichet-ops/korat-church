<template>
  <footer class="bg-gray-900 text-white py-12">
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
      <div class="grid grid-cols-1 md:grid-cols-4 gap-8">
        <!-- Church Info -->
        <div class="col-span-1 md:col-span-2">
          <div class="flex items-center mb-4">
            <a
              href="/"
              @click.prevent="onLogoClick"
              class="flex items-center hover:opacity-80 transition duration-300"
            >
              <img
                src="/images/cross-logo.png"
                alt="Muang Thai Korat Church Logo"
                class="h-12 w-auto mr-3"
              />
              <h3 class="text-base lg:text-lg font-semibold text-white">
                {{ churchName }}
              </h3>
            </a>
          </div>
        </div>

        <!-- Quick Links -->
        <div>
          <h4 class="text-lg font-semibold text-white mb-4">Quick Links</h4>
          <ul class="space-y-2">
            <li>
              <router-link
                to="/about"
                @click="onNavClick('/about')"
                class="text-sm text-gray-300 hover:text-white transition duration-300"
                >About Us</router-link
              >
            </li>
            <li>
              <router-link
                to="/our-ministries"
                class="text-sm text-gray-300 hover:text-white transition duration-300"
                >Our Ministries</router-link
              >
            </li>
            <li>
              <a
                href="/events"
                class="text-sm text-gray-300 hover:text-white transition duration-300"
                >Events</a
              >
            </li>
            <li>
              <a
                href="/contact"
                class="text-sm text-gray-300 hover:text-white transition duration-300"
                >Contact</a
              >
            </li>
            <li>
              <a
                href="/give"
                class="text-sm text-gray-300 hover:text-white transition duration-300"
                >Give</a
              >
            </li>
          </ul>
        </div>

        <!-- Contact Info -->
        <div>
          <h4 class="text-lg font-semibold text-white mb-4">Contact Info</h4>
          <div class="space-y-2 text-sm text-gray-300">
            <div class="flex items-start">
              <svg
                class="w-5 h-5 mr-2 mt-1 flex-shrink-0"
                fill="currentColor"
                viewBox="0 0 20 20"
              >
                <path
                  fill-rule="evenodd"
                  d="M5.05 4.05a7 7 0 119.9 9.9L10 18.9l-4.95-4.95a7 7 0 010-9.9zM10 11a2 2 0 100-4 2 2 0 000 4z"
                  clip-rule="evenodd"
                />
              </svg>
              <a
                :href="googleMapsLink"
                target="_blank"
                rel="noopener"
                class="hover:text-white underline-offset-2 hover:underline"
              >
                {{ churchSettings.church_address }}
              </a>
            </div>
            <div class="flex items-center">
              <svg
                class="w-5 h-5 mr-2 flex-shrink-0"
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
              <span>{{ churchSettings.church_email }}</span>
            </div>
            <div class="flex items-center">
              <svg
                class="w-5 h-5 mr-2 flex-shrink-0"
                fill="currentColor"
                viewBox="0 0 20 20"
              >
                <path
                  d="M2 3a1 1 0 011-1h2.153a1 1 0 01.986.836l.74 4.435a1 1 0 01-.54 1.06l-1.548.773a11.037 11.037 0 006.105 6.105l.774-1.548a1 1 0 011.059-.54l4.435.74a1 1 0 01.836.986V17a1 1 0 01-1 1h-2C7.82 18 2 12.18 2 5V3z"
                />
              </svg>
              <span>{{ churchSettings.church_phone }}</span>
            </div>
          </div>
        </div>
      </div>

      <!-- Bottom Bar -->
      <div class="border-t border-gray-800 mt-8 pt-8 text-center">
        <p class="text-xs md:text-sm text-gray-400">
          © {{ currentYear }} {{ churchName }}. All rights reserved.
        </p>
      </div>
    </div>
  </footer>
</template>

<script>
export default {
  name: 'FooterComponent',
  props: {
    churchName: {
      type: String,
      default: 'Muang Thai Korat Church',
    },
    churchSettings: {
      type: Object,
      default: () => ({
        church_address: '',
        church_email: '',
        church_phone: '',
      }),
    },
  },
  computed: {
    currentYear() {
      return new Date().getFullYear();
    },
    googleMapsLink() {
      // Use the precise Street View URL provided by the user
      return 'https://www.google.com/maps/place/Muangthai+Korat+Church/@14.9592904,102.0389158,3a,75y,275.19h,90t/data=!3m7!1e1!3m5!1s423Hso2quX9FJ4xNDGLhDw!2e0!6shttps:%2F%2Fstreetviewpixels-pa.googleapis.com%2Fv1%2Fthumbnail%3Fcb_client%3Dmaps_sv.tactile%26w%3D900%26h%3D600%26pitch%3D0%26panoid%3D423Hso2quX9FJ4xNDGLhDw%26yaw%3D275.19135!7i16384!8i8192!4m17!1m10!3m9!1s0x311eb31037797a4f:0x6e1c1237a91516b9!2sMuangthai+Korat+Church!8m2!3d14.9593033!4d102.0387701!10e5!14m1!1BCgIgARICGAI!16s%2Fg%2F1hm4whygs!3m5!1s0x311eb31037797a4f:0x6e1c1237a91516b9!8m2!3d14.9593033!4d102.0387701!16s%2Fg%2F1hm4whygs?authuser=0&entry=ttu&g_ep=EgoyMDI1MTAwMS4wIKXMDSoASAFQAw%3D%3D';
    },
  },
  methods: {
    onLogoClick() {
      if (window.location.pathname !== '/') {
        window.location.href = '/';
        return;
      }
      window.scrollTo({ top: 0, behavior: 'smooth' });
    },
    onNavClick(path) {
      if (this.$route && this.$route.path === path) {
        window.scrollTo({ top: 0, behavior: 'smooth' });
      }
    },
  },
};
</script>
