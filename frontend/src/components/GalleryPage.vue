<template>
  <div>
    <!-- Hero Section -->
    <div class="py-20 bg-gray-50">
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
        class="absolute left-4 text-white text-4xl hover:text-gray-300"
      >
        ‹
      </button>
      <div class="max-w-4xl max-h-full" @click.stop>
        <img
          :src="galleryImages[currentImageIndex].src"
          :alt="galleryImages[currentImageIndex].alt"
          class="max-w-full max-h-[80vh] object-contain"
        />
        <p class="text-white text-center mt-4">
          {{ galleryImages[currentImageIndex].alt }}
        </p>
      </div>
      <button
        @click.stop="nextImage"
        class="absolute right-4 text-white text-4xl hover:text-gray-300"
      >
        ›
      </button>
      <button
        @click="closeLightbox"
        class="absolute top-4 right-4 text-white text-4xl hover:text-gray-300"
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
          src: '/images/christmas_eve.jpg',
          alt: 'Christmas Eve Service',
        },
        {
          src: '/images/church_building_1.jpg',
          alt: 'Church Building',
        },
        {
          src: '/images/potluck.jpg',
          alt: 'Community Potluck',
        },
        {
          src: '/images/bible_study.jpg',
          alt: 'Bible Study',
        },
        {
          src: '/images/worship.jpg',
          alt: 'Worship Service',
        },
        {
          src: '/images/children-ministry.jpg',
          alt: 'Children Ministry',
        },
        {
          src: '/images/youth-ministry.jpg',
          alt: 'Youth Ministry',
        },
        {
          src: '/images/adult-bible-study.jpg',
          alt: 'Adult Bible Study',
        },
        {
          src: '/images/worship-ministry.jpg',
          alt: 'Worship Ministry',
        },
        {
          src: '/images/outreach-ministry.jpg',
          alt: 'Outreach Ministry',
        },
        {
          src: '/images/prayer-ministry.jpg',
          alt: 'Prayer Ministry',
        },
        {
          src: '/images/Annual Church Picnic.jpg',
          alt: 'Annual Church Picnic',
        },
        {
          src: '/images/Christmas Celebration.jpg',
          alt: 'Christmas Celebration',
        },
        {
          src: '/images/Easter Sunday Service.jpg',
          alt: 'Easter Sunday Service',
        },
        {
          src: '/images/New Years Prayer Service.jpg',
          alt: 'New Years Prayer Service',
        },
        {
          src: '/images/give.jpg',
          alt: 'Giving',
        },
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
