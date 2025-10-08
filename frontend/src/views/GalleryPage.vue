<template>
  <div>
    <!-- Hero Section -->
    <div class="pt-24 lg:pt-28 pb-20 bg-gray-50">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="text-center mb-10">
          <h1 class="text-4xl md:text-5xl font-bold text-gray-900">
            Photo Gallery
          </h1>
          <p class="text-gray-600 mt-3">
            Moments from our church life and events
          </p>
        </div>

        <div class="grid grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-4">
          <div
            v-for="(image, index) in galleryImages"
            :key="index"
            class="group block overflow-hidden rounded-xl shadow-sm bg-white cursor-pointer"
            @click="openLightbox(index)"
          >
            <img
              :src="image.src"
              :alt="image.alt"
              class="w-full h-36 object-cover group-hover:scale-105 transition-transform duration-500"
            />
          </div>
        </div>
      </div>
    </div>

    <!-- Lightbox -->
    <div
      v-if="lightboxOpen"
      class="fixed inset-0 z-50 bg-black bg-opacity-90 flex items-center justify-center p-4"
      @click="closeLightbox"
    >
      <button
        @click.stop="prevImage"
        class="absolute left-4 text-white text-4xl hover:text-gray-300 cursor-pointer"
      >
        ‹
      </button>
      <div class="max-w-4xl max-h-full relative" @click.stop>
        <img
          :src="galleryImages[currentImageIndex].src"
          :alt="galleryImages[currentImageIndex].alt"
          class="max-w-full max-h-[80vh] object-contain"
        />
        <!-- Image Counter -->
        <div
          class="absolute bottom-2 right-2 bg-black/60 text-white text-sm px-2 py-1 rounded"
        >
          {{ currentImageIndex + 1 }} / {{ galleryImages.length }}
        </div>
      </div>
      <button
        @click.stop="nextImage"
        class="absolute right-4 text-white text-4xl hover:text-gray-300 cursor-pointer"
      >
        ›
      </button>
      <button
        @click="closeLightbox"
        class="absolute top-4 right-4 text-white text-4xl hover:text-gray-300 cursor-pointer"
      >
        ×
      </button>
    </div>
  </div>
</template>

<script>
export default {
  name: 'GalleryPage',
  data() {
    return {
      lightboxOpen: false,
      currentImageIndex: 0,
      galleryImages: [
        {
          src: '/images/gallery/482960224_1003870451838173_3546241293535029221_n.jpg',
          alt: 'Church Community',
        },
        {
          src: '/images/gallery/482978676_1003870675171484_304448253131289558_n.jpg',
          alt: 'Fellowship Time',
        },
        {
          src: '/images/gallery/482979160_1003870368504848_9150588619849555659_n.jpg',
          alt: 'Worship Service',
        },
        {
          src: '/images/gallery/482985442_1003870645171487_6192367423611323304_n.jpg',
          alt: 'Church Gathering',
        },
        {
          src: '/images/gallery/482987292_1003870758504809_8454530624892264720_n.jpg',
          alt: 'Sunday Service',
        },
        {
          src: '/images/gallery/483559632_1003870378504847_5029709997006364920_n.jpg',
          alt: 'Community Event',
        },
        {
          src: '/images/gallery/484205993_1003870648504820_3802464983297733437_n.jpg',
          alt: 'Church Activities',
        },
        {
          src: '/images/gallery/484370333_1003870545171497_5659281909436006415_n.jpg',
          alt: 'Church Life',
        },
        {
          src: '/images/gallery/484458621_1003870455171506_3784833642305727976_n.jpg',
          alt: 'Ministry Event',
        },
        {
          src: '/images/gallery/484631233_1006243248267560_212969535223826768_n.jpg',
          alt: 'Church Family',
        },
        {
          src: '/images/gallery/488014982_1123881696435587_1041300931432242562_n.jpg',
          alt: 'Sunday Worship',
        },
        {
          src: '/images/gallery/488508070_1123881866435570_8999656856006383753_n.jpg',
          alt: 'Community Fellowship',
        },
        {
          src: '/images/gallery/488646670_1123881543102269_793173319572942970_n.jpg',
          alt: 'Church Event',
        },
        {
          src: '/images/gallery/488687040_1123881833102240_6675477191327402878_n.jpg',
          alt: 'Worship Time',
        },
        {
          src: '/images/gallery/488764889_1123881816435575_152737993779818058_n.jpg',
          alt: 'Church Gathering',
        },
        {
          src: '/images/gallery/488836307_1123881576435599_3105616996385475809_n.jpg',
          alt: 'Ministry Time',
        },
        {
          src: '/images/gallery/488927948_1123881699768920_1656124403816842703_n.jpg',
          alt: 'Church Service',
        },
        {
          src: '/images/gallery/489034415_24031724733097560_5843822842879051548_n.jpg',
          alt: 'Church Activities',
        },
        {
          src: '/images/gallery/489043282_1123881903102233_1917235248441307686_n.jpg',
          alt: 'Fellowship Event',
        },
        {
          src: '/images/gallery/489103997_1123881899768900_7650300014201284869_n.jpg',
          alt: 'Worship Service',
        },
        {
          src: '/images/gallery/489547625_1123881746435582_2783083233404649943_n.jpg',
          alt: 'Church Community',
        },
        {
          src: '/images/gallery/489618858_24031723906430976_7118238866147766557_n.jpg',
          alt: 'Church Life',
        },
        {
          src: '/images/gallery/489695632_1123881676435589_9063802118461382255_n.jpg',
          alt: 'Ministry Event',
        },
        {
          src: '/images/gallery/489948083_24031722219764478_2524803143642510094_n.jpg',
          alt: 'Church Gathering',
        },
        {
          src: '/images/gallery/495387377_24269504432652921_3921599915490744244_n.jpg',
          alt: 'Youth Ministry',
        },
        {
          src: '/images/gallery/495557497_24304689689134395_6364684146169796891_n.jpg',
          alt: 'Church Service',
        },
        {
          src: '/images/gallery/495609837_24304689475801083_8071265327311301657_n.jpg',
          alt: 'Fellowship Time',
        },
        {
          src: '/images/gallery/495795721_24304689529134411_91612552111306801_n.jpg',
          alt: 'Worship Service',
        },
        {
          src: '/images/gallery/495904414_24304689329134431_7486531290480889195_n.jpg',
          alt: 'Church Event',
        },
        {
          src: '/images/gallery/495995265_24304689342467763_450040876158480599_n.jpg',
          alt: 'Community Time',
        },
        {
          src: '/images/gallery/496024485_24304689532467744_6307468164738969804_n.jpg',
          alt: 'Church Activities',
        },
        {
          src: '/images/gallery/496062442_24304689699134394_467819688883526969_n.jpg',
          alt: 'Ministry Event',
        },
        {
          src: '/images/gallery/496100932_24269504355986262_3407648028291644391_n.jpg',
          alt: 'Church Gathering',
        },
        {
          src: '/images/gallery/496521576_24269504349319596_1874423298565187237_n.jpg',
          alt: 'Worship Time',
        },
        {
          src: '/images/gallery/496922520_24304689315801099_7928334991253994282_n.jpg',
          alt: 'Church Life',
        },
        {
          src: '/images/gallery/496925927_24304689662467731_6049132363213750978_n.jpg',
          alt: 'Church Family',
        },
        { src: '/images/gallery/church.jpg', alt: 'Our Church' },
      ],
    };
  },
  methods: {
    openLightbox(index) {
      this.currentImageIndex = index;
      this.lightboxOpen = true;
      document.body.style.overflow = 'hidden';
    },
    closeLightbox() {
      this.lightboxOpen = false;
      document.body.style.overflow = '';
    },
    nextImage() {
      this.currentImageIndex =
        (this.currentImageIndex + 1) % this.galleryImages.length;
    },
    prevImage() {
      this.currentImageIndex =
        (this.currentImageIndex - 1 + this.galleryImages.length) %
        this.galleryImages.length;
    },
  },
  mounted() {
    document.addEventListener('keydown', (e) => {
      if (!this.lightboxOpen) return;
      if (e.key === 'Escape') this.closeLightbox();
      if (e.key === 'ArrowRight') this.nextImage();
      if (e.key === 'ArrowLeft') this.prevImage();
    });
  },
};
</script>
