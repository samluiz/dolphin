<script setup lang="ts">
import { computed, toRefs } from "vue";
import { pagination as p } from "../../../wailsjs/go/models";

const props = defineProps({
  pagination: Object as () => p.PaginationOutput,
});

const { pagination } = toRefs(props);

const emit = defineEmits(["on-page-change"]);

const handlePageChange = (page: number) => {
  if (page < 1 || page > pagination?.value?.total_pages!) return;
  emit("on-page-change", page);
};

const visiblePages = computed(() => {
  const pages = [];
  const totalPages = pagination?.value?.total_pages || 0;
  const currentPage = pagination?.value?.page || 1;

  if (totalPages <= 5) {
    for (let i = 1; i <= totalPages; i++) {
      pages.push(i);
    }
  } else {
    if (currentPage <= 3) {
      pages.push(1, 2, 3, 4, totalPages);
    } else if (currentPage >= totalPages - 2) {
      pages.push(1, totalPages - 3, totalPages - 2, totalPages - 1, totalPages);
    } else {
      pages.push(1, currentPage - 1, currentPage, currentPage + 1, totalPages);
    }
  }

  return pages;
});
</script>

<template>
  <div class="flex items-center gap-1">
    <!-- Previous button -->
    <button
      @click="handlePageChange(pagination!.page - 1)"
      class="p-2 rounded-lg
             text-text-primary-light dark:text-text-primary-dark
             hover:bg-primary-50 dark:hover:bg-surface-dark
             disabled:opacity-30 disabled:cursor-not-allowed
             transition-all duration-200 ease-out
             transform hover:scale-110 active:scale-95"
      :disabled="pagination && pagination.prev_page === 0"
    >
      <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7"/>
      </svg>
    </button>

    <!-- Page numbers -->
    <button
      v-for="page in visiblePages"
      :key="page"
      @click="handlePageChange(page)"
      class="min-w-[2.5rem] px-3 py-2 rounded-lg font-medium text-sm
             transition-all duration-200 ease-out
             transform hover:scale-110 active:scale-95"
      :class="{
        'bg-gradient-ocean text-white shadow-md-blue': pagination && pagination.page === page,
        'text-text-primary-light dark:text-text-primary-dark hover:bg-primary-50 dark:hover:bg-surface-dark': pagination && pagination.page !== page,
      }"
    >
      {{ page }}
    </button>

    <!-- Next button -->
    <button
      @click="handlePageChange(pagination!.page + 1)"
      class="p-2 rounded-lg
             text-text-primary-light dark:text-text-primary-dark
             hover:bg-primary-50 dark:hover:bg-surface-dark
             disabled:opacity-30 disabled:cursor-not-allowed
             transition-all duration-200 ease-out
             transform hover:scale-110 active:scale-95"
      :disabled="pagination && pagination.next_page === 0"
    >
      <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"/>
      </svg>
    </button>
  </div>
</template>
