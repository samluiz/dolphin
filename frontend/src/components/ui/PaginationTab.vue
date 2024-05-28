<script setup lang="ts">
import { computed, toRefs } from "vue";
import { types } from "wailsjs/go/models";
import IconButton from "./IconButton.vue";

const props = defineProps({
  pagination: Object as () => types.PaginationOutput,
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
  <div class="flex justify-center gap-2 text-sm">
    <button
      @click="handlePageChange(pagination!.page - 1)"
      class="px-2 text-black dark:text-light-text duration-200"
      :disabled="pagination && pagination.prev_page === 0"
      :class="{
        'opacity-30': pagination && pagination.prev_page === 0,
      }"
    >
      &lt;
    </button>
    <button
      v-for="page in visiblePages"
      :key="page"
      @click="handlePageChange(page)"
      class="px-2 text-black dark:text-light-text rounded-lg duration-200"
      :class="{
        'opacity-30 pointer-events-none':
          pagination && pagination.page === page,
      }"
    >
      {{ page }}
    </button>
    <button
      @click="handlePageChange(pagination!.page + 1)"
      class="px-2 text-black dark:text-light-text rounded-lg duration-200"
      :disabled="pagination && pagination.next_page === 0"
      :class="{
        'opacity-30': pagination && pagination.next_page === 0,
      }"
    >
      &gt;
    </button>
  </div>
</template>
