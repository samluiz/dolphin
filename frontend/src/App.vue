<script lang="ts" setup>
import { computed, ref } from 'vue';


interface Tab {
  id: number
  name: string
}

const tabs: Tab[] = [
  { id: 1, name: 'Earnings' },
  { id: 2, name: 'Expenses' },
]

const earnings = [
  { id: 1, description: 'Earnings 1', amount: 100 },
  { id: 2, description: 'Earnings 2', amount: 200 },
  { id: 3, description: 'Earnings 3', amount: 300 },
]

const expenses = [
  { id: 1, description: 'Expenses 1', category: 'Fun', amount: 100 },
  { id: 2, description: 'Expenses 2', category: 'Fun', amount: 200 },
  { id: 3, description: 'Expenses 3', category: 'Fun', amount: 300 },
]

const filteredData = computed((): any => {
  if (selectedTabId.value === 1) {
    return earnings;
  } else {
    return expenses;
  }
});

const selectedTabId = ref(1);
</script>

<template>
  <div class="h-full w-full grid place-items-center p-8 gap-4">
    <div class="w-full grid place-items-start">
      <h1 class="text-white text-5xl">Dashboard</h1>
    </div>
    <div class="flex mb-4 gap-2">
      <button v-for="tab in tabs" :key="tab.id"
        @click="selectedTabId = tab.id"
        :class="{ 'bg-gray-200 text-gray-600 font-bold': selectedTabId === tab.id }"
      >
        {{ tab.name }}
      </button>
    </div>

    <table class="min-w-full leading-normal">
      <thead>
        <tr class="text-center">
          <th class="px-5 py-3 border-b-2 
          border-gray-200 bg-gray-100 
          text-xs font-semibold 
          text-gray-700 uppercase 
          tracking-wider">Description</th>
          <th class="px-5 py-3 border-b-2 
          border-gray-200 bg-gray-100 
          text-xs font-semibold 
          text-gray-700 uppercase 
          tracking-wider">Amount</th>
          <th class="px-5 py-3 border-b-2 
          border-gray-200 bg-gray-100 
          text-xs font-semibold 
          text-gray-700 uppercase 
          tracking-wider" v-if="selectedTabId === 2">Category</th>
          <th class="px-5 py-3 border-b-2 
          border-gray-200 bg-gray-100 
          text-xs font-semibold 
          text-gray-700 uppercase 
          tracking-wider">Amount</th>
        </tr>
      </thead>
      <tbody>
        <tr class="text-center" v-for="item in filteredData" :key="item.id">
          <td class="px-5 py-5 border-b border-gray-200 bg-transparent text-sm">{{ item.description }}</td>
          <td class="px-5 py-5 border-b border-gray-200 bg-transparent text-sm">{{ item.amount }}</td>
          <td class="px-5 py-5 border-b border-gray-200 bg-transparent text-sm" v-if="selectedTabId === 2">{{ item.category }}</td>
          <td class="px-5 py-5 border-b border-gray-200 bg-transparent text-sm">{{ item.amount }}</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

