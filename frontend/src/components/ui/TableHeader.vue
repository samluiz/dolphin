<script setup lang="ts"></script>

<template>
  <th
    class="px-4 py-3
           bg-gradient-to-br from-primary-50 to-blue-50 dark:from-surface-dark dark:to-bg-dark
           border-b-2 border-primary-200 dark:border-border-dark
           text-xs font-bold text-primary-700 dark:text-primary-300
           uppercase tracking-wider
           whitespace-nowrap
           text-left
           first:rounded-tl-lg last:rounded-tr-lg
           sticky top-0 z-10
           backdrop-blur-sm"
  >
    <div class="flex items-center gap-2">
      <slot></slot>
    </div>
  </th>
</template>
