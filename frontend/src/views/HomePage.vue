<template>
  <div>
    <!-- Full Screen Hero Section -->
    <div
      class="relative min-h-screen flex items-center justify-center"
      style="
        background: linear-gradient(rgba(0, 0, 0, 0.5), rgba(0, 0, 0, 0.5)),
          url('/images/church_building_1.jpg') no-repeat center center;
        background-size: cover;
      "
    >
      <!-- Hero Content -->
      <div class="relative z-10 text-center text-white px-4 sm:px-6 lg:px-8">
        <h1
          class="text-2xl md:text-4xl font-bold mb-4 leading-tight text-gray-200"
          style="text-shadow: 0 2px 4px rgba(0, 0, 0, 0.5)"
        >
          Welcome to
        </h1>
        <h1
          class="text-5xl md:text-7xl xl:text-8xl font-extrabold mb-6 leading-tight text-white"
          style="text-shadow: 0 2px 8px rgba(0, 0, 0, 0.6)"
        >
          {{ churchSettings.church_name || 'Muang Thai Korat Church' }}
        </h1>
        <p
          class="text-xl md:text-2xl text-white mb-3 max-w-3xl mx-auto leading-relaxed"
        >
          Join us in worship, fellowship, and growing together in faith.
        </p>
        <p
          class="text-xl md:text-2xl text-white mb-12 max-w-3xl mx-auto leading-relaxed"
        >
          Experience the love of Christ in our community.
        </p>

        <!-- Call to Action Links -->
        <div
          class="flex flex-col sm:flex-row gap-4 sm:gap-8 justify-center items-center"
        >
          <a
            @click="$router.push('/contact')"
            class="group text-white text-base md:text-lg cursor-pointer flex items-center gap-2 hover:text-blue-400 transition-all duration-300"
          >
            <span>Plan Your Visit</span>
            <svg
              class="w-4 h-4 transform group-hover:translate-x-1 transition-transform duration-300"
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M17 8l4 4m0 0l-4 4m4-4H3"
              />
            </svg>
          </a>
          <a
            @click="$router.push('/about')"
            class="group text-white text-base md:text-lg cursor-pointer flex items-center gap-2 hover:text-blue-400 transition-all duration-300"
          >
            <span>Learn More</span>
            <svg
              class="w-4 h-4 transform group-hover:translate-x-1 transition-transform duration-300"
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M17 8l4 4m0 0l-4 4m4-4H3"
              />
            </svg>
          </a>
        </div>
      </div>
    </div>

    <!-- Welcome Section -->
    <section class="py-20 relative bg-white">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="text-center">
          <h2
            class="text-3xl md:text-4xl lg:text-5xl font-bold text-gray-900 mb-6"
          >
            Welcome to Our <span class="text-blue-600">Faith Community</span>
          </h2>
          <p
            class="text-lg md:text-xl text-gray-700 max-w-3xl mx-auto leading-relaxed mb-16"
          >
            Join us in worship, fellowship, and growing together in faith.
            <br />
            Experience the love of Christ in our welcoming community where
            everyone belongs.
          </p>
        </div>

        <div
          ref="welcomeCards"
          class="grid grid-cols-1 md:grid-cols-3 gap-8 mt-12"
        >
          <div
            :class="['welcome-card', { visible: welcomeCardsVisible }]"
            class="relative text-center p-8 rounded-lg shadow-lg transition-shadow duration-300 ease-out hover:shadow-xl overflow-hidden bg-cover bg-center min-h-[300px] flex flex-col justify-end"
            style="
              background-image: linear-gradient(
                  rgba(0, 0, 0, 0.3),
                  rgba(0, 0, 0, 0.6)
                ),
                url('/images/bible_study.jpg');
            "
          >
            <h3 class="text-2xl font-semibold text-white mb-4 drop-shadow-lg">
              Bible Study
            </h3>
            <p class="text-white drop-shadow-md">
              Deepen your understanding of God's Word through our engaging Bible
              study programs for all ages.
            </p>
          </div>

          <div
            :class="['welcome-card', { visible: welcomeCardsVisible }]"
            class="relative text-center p-8 rounded-lg shadow-lg transition-shadow duration-300 ease-out hover:shadow-xl overflow-hidden bg-cover bg-center min-h-[300px] flex flex-col justify-end"
            style="
              background-image: linear-gradient(
                  rgba(0, 0, 0, 0.3),
                  rgba(0, 0, 0, 0.6)
                ),
                url('/images/community.jpg');
            "
          >
            <h3 class="text-2xl font-semibold text-white mb-4 drop-shadow-lg">
              Community
            </h3>
            <p class="text-white drop-shadow-md">
              Build meaningful relationships in our supportive community where
              faith and friendship come together.
            </p>
          </div>

          <div
            :class="['welcome-card', { visible: welcomeCardsVisible }]"
            class="relative text-center p-8 rounded-lg shadow-lg transition-shadow duration-300 ease-out hover:shadow-xl overflow-hidden bg-cover bg-center min-h-[300px] flex flex-col justify-end"
            style="
              background-image: linear-gradient(
                  rgba(0, 0, 0, 0.3),
                  rgba(0, 0, 0, 0.6)
                ),
                url('/images/worship.jpg');
            "
          >
            <h3 class="text-2xl font-semibold text-white mb-4 drop-shadow-lg">
              Worship
            </h3>
            <p class="text-white drop-shadow-md">
              Experience inspiring worship services that uplift your spirit and
              draw you closer to God.
            </p>
          </div>
        </div>
      </div>
    </section>

    <!-- Service Times Section -->
    <section
      class="py-20 bg-gray-200"
      style="
        background: linear-gradient(rgba(0, 0, 0, 0.5), rgba(0, 0, 0, 0.5)),
          url('/images/worship.jpg') no-repeat center center;
        background-size: cover;
      "
    >
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="text-center max-w-4xl mx-auto">
          <h2
            class="text-3xl md:text-4xl lg:text-5xl font-bold text-white mb-6"
          >
            Join Us for <span class="text-blue-300">Worship</span>
          </h2>
          <p class="text-xl text-gray-200 mb-8">
            We welcome you to join us for our weekly services.<br />
            Come as you are and experience the love of Christ in our community.
          </p>

          <div class="grid grid-cols-1 md:grid-cols-3 gap-6 max-w-4xl mx-auto">
            <div
              v-for="service in serviceTimes"
              :key="service.day"
              class="text-center p-6 bg-white/20 backdrop-blur-sm rounded-xl shadow-sm transition-shadow duration-200 ease-out hover:shadow-md border border-white/30"
            >
              <h4 class="text-xl font-semibold text-white mb-2 drop-shadow-md">
                {{ service.day }} {{ service.service }}
              </h4>
              <p
                class="text-2xl font-bold mb-2 drop-shadow-md"
                :class="
                  service.day === 'Sunday'
                    ? 'text-blue-300'
                    : service.day === 'Wednesday'
                    ? 'text-green-300'
                    : 'text-purple-300'
                "
              >
                {{ service.time }}
              </p>
              <p class="text-gray-100 text-sm drop-shadow">
                {{ service.service }}
              </p>
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- Events Section -->
    <section class="py-20 bg-white">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="text-center mb-16">
          <h2
            class="text-3xl md:text-4xl lg:text-5xl font-bold text-gray-900 mb-6"
          >
            Our Church <span class="text-blue-600">Events</span>
          </h2>
          <p class="text-xl text-gray-600 max-w-3xl mx-auto">
            Join us for special events, celebrations, and community gatherings
            throughout the year.
          </p>
        </div>

        <div
          v-if="events.length > 0"
          class="grid grid-cols-1 md:grid-cols-3 gap-8"
        >
          <div
            v-for="(event, index) in events.slice(0, 3)"
            :key="event.id"
            class="bg-white rounded-xl shadow-lg overflow-hidden"
          >
            <div
              class="h-48 bg-cover bg-center hover:scale-105 transition-transform duration-500"
              :style="`background-image: url('/images/${
                index === 0
                  ? 'christmas_eve.jpg'
                  : index === 1
                  ? 'bible_study.jpg'
                  : 'community.jpg'
              }')`"
            ></div>
            <div class="p-6">
              <h3 class="text-xl font-bold text-gray-900 mb-3">
                {{ event.title }}
              </h3>
              <p class="text-gray-600 mb-4">{{ event.description }}</p>
              <button
                @click="$router.push('/events')"
                class="text-blue-600 font-semibold hover:text-blue-700 transition duration-300"
              >
                Learn More →
              </button>
            </div>
          </div>
        </div>

        <div class="text-center mt-12">
          <button
            @click="$router.push('/events')"
            class="inline-block bg-blue-600 text-white px-6 py-3 rounded-lg text-base font-semibold hover:bg-white hover:text-blue-600 border border-blue-600 hover:border-blue-600 transition duration-300 cursor-pointer"
          >
            View All Events
          </button>
        </div>
      </div>
    </section>

    <!-- Call to Action -->
    <section class="py-20 bg-blue-50">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 text-center">
        <h2
          class="text-3xl md:text-4xl lg:text-5xl font-bold text-gray-900 mb-6"
        >
          Ready to Join Our <span class="text-blue-600">Community?</span>
        </h2>
        <p class="text-xl text-gray-600 mb-8 max-w-3xl mx-auto">
          We'd love to welcome you to our church family. Come experience the
          love of Christ in our community.
        </p>
        <div class="flex flex-col sm:flex-row gap-4 justify-center">
          <button
            @click="$router.push('/contact')"
            class="border border-blue-600 text-blue-600 px-6 py-3 rounded-lg text-base font-semibold hover:bg-blue-600 hover:text-white hover:shadow-lg transition-all duration-300 cursor-pointer"
          >
            Contact Us
          </button>
          <button
            @click="$router.push('/about')"
            class="border border-blue-600 text-blue-600 px-6 py-3 rounded-lg text-base font-semibold hover:bg-blue-600 hover:text-white hover:shadow-lg transition-all duration-300 cursor-pointer"
          >
            Learn More
          </button>
        </div>
      </div>
    </section>

    <!-- Recent Photos Carousel Section -->
    <section class="py-20 bg-white">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="text-center mb-12">
          <h2
            class="text-3xl md:text-4xl lg:text-5xl font-bold text-gray-900 mb-4"
          >
            Recent <span class="text-blue-600">Photos</span>
          </h2>
          <p class="text-xl text-gray-600">Moments from our church community</p>
        </div>

        <!-- Photo Grid Gallery -->
        <div class="grid grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-4">
          <div
            v-for="(photo, index) in recentPhotos.slice(0, 8)"
            :key="index"
            class="relative aspect-square overflow-hidden rounded-lg shadow-lg hover:shadow-2xl transition-all duration-300 group cursor-pointer"
            @click="$router.push('/gallery')"
          >
            <img
              :src="photo.src"
              :alt="photo.alt"
              class="w-full h-full object-cover group-hover:scale-110 transition-transform duration-500"
            />
            <!-- Hover Overlay -->
            <div
              class="absolute inset-0 bg-black/0 group-hover:bg-black/20 transition-all duration-300"
            ></div>
          </div>
        </div>

        <!-- View All Button -->
        <div class="text-center mt-12">
          <button
            @click="$router.push('/gallery')"
            class="text-blue-600 cursor-pointer underline hover:text-blue-800 transition-colors duration-300"
          >
            View All Photos
          </button>
        </div>
      </div>
    </section>
  </div>
</template>

<script>
export default {
  name: 'HomePage',
  props: {
    events: Array,
    churchSettings: Object,
    serviceTimes: Array,
  },
  data() {
    return {
      welcomeCardsVisible: false,
      currentPhotoIndex: 0,
      autoPlayInterval: null,
      recentPhotos: [
        {
          src: '/images/gallery/482960224_1003870451838173_3546241293535029221_n.jpg',
          alt: 'Church Photo',
          title: 'Church Community',
        },
        {
          src: '/images/gallery/482978676_1003870675171484_304448253131289558_n.jpg',
          alt: 'Church Photo',
          title: 'Fellowship Time',
        },
        {
          src: '/images/gallery/482979160_1003870368504848_9150588619849555659_n.jpg',
          alt: 'Church Photo',
          title: 'Worship Service',
        },
        {
          src: '/images/gallery/482985442_1003870645171487_6192367423611323304_n.jpg',
          alt: 'Church Photo',
          title: 'Church Gathering',
        },
        {
          src: '/images/gallery/482987292_1003870758504809_8454530624892264720_n.jpg',
          alt: 'Church Photo',
          title: 'Sunday Service',
        },
        {
          src: '/images/gallery/483559632_1003870378504847_5029709997006364920_n.jpg',
          alt: 'Church Photo',
          title: 'Community Event',
        },
        {
          src: '/images/gallery/484205993_1003870648504820_3802464983297733437_n.jpg',
          alt: 'Church Photo',
          title: 'Church Activities',
        },
        {
          src: '/images/gallery/484370333_1003870545171497_5659281909436006415_n.jpg',
          alt: 'Church Photo',
          title: 'Church Life',
        },
        {
          src: '/images/gallery/484458621_1003870455171506_3784833642305727976_n.jpg',
          alt: 'Church Photo',
          title: 'Ministry Event',
        },
        {
          src: '/images/gallery/484631233_1006243248267560_212969535223826768_n.jpg',
          alt: 'Church Photo',
          title: 'Church Family',
        },
        {
          src: '/images/gallery/488014982_1123881696435587_1041300931432242562_n.jpg',
          alt: 'Church Photo',
          title: 'Sunday Worship',
        },
        {
          src: '/images/gallery/488508070_1123881866435570_8999656856006383753_n.jpg',
          alt: 'Church Photo',
          title: 'Community Fellowship',
        },
        {
          src: '/images/gallery/488646670_1123881543102269_793173319572942970_n.jpg',
          alt: 'Church Photo',
          title: 'Church Event',
        },
        {
          src: '/images/gallery/488687040_1123881833102240_6675477191327402878_n.jpg',
          alt: 'Church Photo',
          title: 'Worship Time',
        },
        {
          src: '/images/gallery/488764889_1123881816435575_152737993779818058_n.jpg',
          alt: 'Church Photo',
          title: 'Church Gathering',
        },
        {
          src: '/images/gallery/488836307_1123881576435599_3105616996385475809_n.jpg',
          alt: 'Church Photo',
          title: 'Ministry Time',
        },
        {
          src: '/images/gallery/488927948_1123881699768920_1656124403816842703_n.jpg',
          alt: 'Church Photo',
          title: 'Church Service',
        },
        {
          src: '/images/gallery/489034415_24031724733097560_5843822842879051548_n.jpg',
          alt: 'Church Photo',
          title: 'Church Activities',
        },
        {
          src: '/images/gallery/489043282_1123881903102233_1917235248441307686_n.jpg',
          alt: 'Church Photo',
          title: 'Fellowship Event',
        },
        {
          src: '/images/gallery/489103997_1123881899768900_7650300014201284869_n.jpg',
          alt: 'Church Photo',
          title: 'Worship Service',
        },
        {
          src: '/images/gallery/489547625_1123881746435582_2783083233404649943_n.jpg',
          alt: 'Church Photo',
          title: 'Church Community',
        },
        {
          src: '/images/gallery/489618858_24031723906430976_7118238866147766557_n.jpg',
          alt: 'Church Photo',
          title: 'Church Life',
        },
        {
          src: '/images/gallery/489695632_1123881676435589_9063802118461382255_n.jpg',
          alt: 'Church Photo',
          title: 'Ministry Event',
        },
        {
          src: '/images/gallery/489948083_24031722219764478_2524803143642510094_n.jpg',
          alt: 'Church Photo',
          title: 'Church Gathering',
        },
        {
          src: '/images/gallery/495387377_24269504432652921_3921599915490744244_n.jpg',
          alt: 'Church Photo',
          title: 'Youth Ministry',
        },
        {
          src: '/images/gallery/495557497_24304689689134395_6364684146169796891_n.jpg',
          alt: 'Church Photo',
          title: 'Church Service',
        },
        {
          src: '/images/gallery/495609837_24304689475801083_8071265327311301657_n.jpg',
          alt: 'Church Photo',
          title: 'Fellowship Time',
        },
        {
          src: '/images/gallery/495795721_24304689529134411_91612552111306801_n.jpg',
          alt: 'Church Photo',
          title: 'Worship Service',
        },
        {
          src: '/images/gallery/495904414_24304689329134431_7486531290480889195_n.jpg',
          alt: 'Church Photo',
          title: 'Church Event',
        },
        {
          src: '/images/gallery/495995265_24304689342467763_450040876158480599_n.jpg',
          alt: 'Church Photo',
          title: 'Community Time',
        },
        {
          src: '/images/gallery/496024485_24304689532467744_6307468164738969804_n.jpg',
          alt: 'Church Photo',
          title: 'Church Activities',
        },
        {
          src: '/images/gallery/496062442_24304689699134394_467819688883526969_n.jpg',
          alt: 'Church Photo',
          title: 'Ministry Event',
        },
        {
          src: '/images/gallery/496100932_24269504355986262_3407648028291644391_n.jpg',
          alt: 'Church Photo',
          title: 'Church Gathering',
        },
        {
          src: '/images/gallery/496521576_24269504349319596_1874423298565187237_n.jpg',
          alt: 'Church Photo',
          title: 'Worship Time',
        },
        {
          src: '/images/gallery/496922520_24304689315801099_7928334991253994282_n.jpg',
          alt: 'Church Photo',
          title: 'Church Life',
        },
        {
          src: '/images/gallery/496925927_24304689662467731_6049132363213750978_n.jpg',
          alt: 'Church Photo',
          title: 'Church Family',
        },
        {
          src: '/images/gallery/church.jpg',
          alt: 'Church Photo',
          title: 'Our Church',
        },
      ],
    };
  },
  mounted() {
    this.setupIntersectionObserver();
    this.startAutoPlay();
  },
  beforeUnmount() {
    this.stopAutoPlay();
  },
  methods: {
    setupIntersectionObserver() {
      const observer = new IntersectionObserver(
        (entries) => {
          entries.forEach((entry) => {
            if (entry.isIntersecting) {
              this.welcomeCardsVisible = true;
              observer.unobserve(entry.target);
            }
          });
        },
        { threshold: 0.1 }
      );

      const welcomeSection = this.$refs.welcomeCards;
      if (welcomeSection) {
        observer.observe(welcomeSection);
      }
    },
    nextPhoto() {
      this.currentPhotoIndex =
        (this.currentPhotoIndex + 1) % this.recentPhotos.length;
      this.resetAutoPlay();
    },
    previousPhoto() {
      this.currentPhotoIndex =
        this.currentPhotoIndex === 0
          ? this.recentPhotos.length - 1
          : this.currentPhotoIndex - 1;
      this.resetAutoPlay();
    },
    startAutoPlay() {
      this.autoPlayInterval = setInterval(() => {
        this.nextPhoto();
      }, 5000); // Change photo every 5 seconds
    },
    stopAutoPlay() {
      if (this.autoPlayInterval) {
        clearInterval(this.autoPlayInterval);
      }
    },
    resetAutoPlay() {
      this.stopAutoPlay();
      this.startAutoPlay();
    },
  },
};
</script>

<style scoped>
.welcome-card {
  opacity: 0;
  transform: translateY(30px);
  transition: all 0.6s ease-out;
}

.welcome-card.visible {
  opacity: 1;
  transform: translateY(0);
}

.welcome-card:nth-child(1) {
  transition-delay: 0.1s;
}

.welcome-card:nth-child(2) {
  transition-delay: 0.3s;
}

.welcome-card:nth-child(3) {
  transition-delay: 0.5s;
}

/* Hover effects */
.welcome-card:hover {
  transform: translateY(-4px) scale(1.01);
  transition: all 0.3s ease;
}

.welcome-card::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(255, 255, 255, 0);
  transition: background 0.3s ease;
  pointer-events: none;
}

.welcome-card:hover::before {
  background: rgba(255, 255, 255, 0.05);
}
</style>
