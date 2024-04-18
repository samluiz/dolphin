<script setup lang="ts">
import { computed, ref, watch } from 'vue';
import EditIcon from './ui/icons/EditIcon.vue';
import DeleteIcon from './ui/icons/DeleteIcon.vue';
import IconButton from './ui/IconButton.vue';
import AddIcon from './ui/icons/AddIcon.vue';
import { Expense, Earning } from '@/shared/types'
import { maskCurrency } from '@/shared/utils'
import DestructiveActionModal from './DestructiveActionModal.vue';
import EditItemModal from './EditItemModal.vue';
import AddItemModal from './AddItemModal.vue';
import TableHeader from './ui/TableHeader.vue';
import TableData from './ui/TableData.vue';

const isDeleteModalOpen = ref(false)
const isEditModalOpen = ref(false)
const isAddModalOpen = ref(false)
const selectedRowId = ref(0);
const selectedTabId = ref(1);
const trigger = ref(0);

const openDeleteModal = (id: number) => {
  console.log('open modal')
  selectedRowId.value = id;
  isDeleteModalOpen.value = true;
}

const closeDeleteModal = () => {
  console.log('close modal')
  selectedRowId.value = 0;
  isDeleteModalOpen.value = false;
}

const deleteItem = () => {
  console.log('delete item')
  if (selectedTabId.value === 1) {
    earnings.splice(earnings.findIndex((earning: Earning) => earning.id === selectedRowId.value), 1);
  } else {
    expenses.splice(expenses.findIndex((expense: Expense) => expense.id === selectedRowId.value), 1);
  }
  isDeleteModalOpen.value = false;
  // recomputing the total values
  trigger.value += 1;
}

const openEditModal = (id: number) => {
  console.log('open edit modal')
  selectedRowId.value = id;
  isEditModalOpen.value = true;
}

const closeEditModal = () => {
  console.log('close edit modal')
  selectedRowId.value = 0;
  isEditModalOpen.value = false;
}

const editItem = (description: string, amount: number, category: string) => {
  console.log('edit item')
  console.log(description, amount, category)
  if (selectedTabId.value === 1) {
    editEarning(description, amount);
  } else {
    editExpense(description, amount, category);
  }
  selectedRowId.value = 0;
  isEditModalOpen.value = false;
}

function editEarning(description: string, amount: number) {
  console.log('edit earning')
  earnings.forEach((earning: Earning) => {
    if (earning.id === selectedRowId.value) {
      earning.description = description;
      earning.amount = amount;
    }
  });

  console.log(earnings)
}

function editExpense(description: string, amount: number, category: string) {
  console.log('edit expense')
  expenses.forEach((expense: Expense) => {
    if (expense.id === selectedRowId.value) {
      expense.description = description;
      expense.amount = amount;
      expense.category = category;
    }
  });
}

const openAddModal = () => {
  console.log('open add modal')
  isAddModalOpen.value = true;
}

const closeAddModal = () => {
  console.log('close add modal')
  isAddModalOpen.value = false;
}

const addItem = (description: string, amount: number, category: string) => {
  console.log('add item')
  console.log(description, amount, category)
  if (selectedTabId.value === 1) {
    addEarning(description, amount);
  } else {
    addExpense(description, amount, category);
  }
  isAddModalOpen.value = false;
  // recomputing the total values
  trigger.value += 1;
}

function addEarning(description: string, amount: number) {
  console.log('add earning')
  earnings.push({ id: earnings.length + 1, description, amount });
}

function addExpense(description: string, amount: number, category: string) {
  console.log('add expense')
  expenses.push({ id: expenses.length + 1, description, category, amount });
}

interface Tab {
  id: number
  name: string
}

const tabs: Tab[] = [
  { id: 1, name: 'Earnings' },
  { id: 2, name: 'Expenses' },
]

const earnings: Earning[] = [
  { id: 1, description: 'Earnings 1', amount: 100 },
  { id: 2, description: 'Earnings 2', amount: 200 },
  { id: 3, description: 'Earnings 3', amount: 300 },
]

const expenses: Expense[] = [
  { id: 1, description: 'Expenses 1', category: 'Fun', amount: 100 },
  { id: 2, description: 'Expenses 2', category: 'Fun', amount: 200 },
  { id: 3, description: 'Expenses 3', category: 'Fun', amount: 300 },
]

const tableData = computed((): Earning[] | Expense[] => {
  return selectedTabId.value === 1 ? earnings : expenses;
});

const earningsTotal = computed(() => {
  let total = 0;
  earnings.forEach((earning: Earning) => {
    total += earning.amount;
  });
  return total;
});

const expensesTotal = computed(() => {
  let total = 0;
  expenses.forEach((expense: Expense) => {
    total += expense.amount;
  });
  return total;
});

const balance = computed(() => {
  return earningsTotal.value - expensesTotal.value;
});
</script>

<template>
 <div class="flex mb-4 gap-2">
      <button v-for="tab in tabs" :key="tab.id"
        @click="selectedTabId = tab.id"
        class="p-2 font-semibold rounded-md"
        :class="{ 'duration-200 bg-gray-200 text-gray-600 font-bold': selectedTabId === tab.id }"
      >
        {{ tab.name }}
      </button>
    </div>

    <table :key="trigger" class="min-w-full leading-normal">
      <thead>
        <tr class="text-center">
          <TableHeader>Description</TableHeader>
          <TableHeader>Amount</TableHeader>
          <TableHeader v-if="selectedTabId === 2">Category</TableHeader>
          <TableHeader></TableHeader>
        </tr>
      </thead>
      <tbody>
        <tr class="text-center" v-for="item in tableData" :key="item.id">
          <TableData>{{ item.description }}</TableData>
          <TableData>{{ maskCurrency(item.amount)}}</TableData>
          <TableData v-if="selectedTabId === 2">{{ (item as Expense).category }}</TableData>
          <TableData class="flex flex-row justify-end">
              <div class="flex flex-row justify-evenly w-20">
                <IconButton @click="openEditModal(item.id)">
                    <EditIcon />
                </IconButton>
                <IconButton @click="openDeleteModal(item.id)">
                    <DeleteIcon />
                </IconButton>
              </div>
          </TableData>
        </tr>
      </tbody>
    </table>

    <div class="w-full flex items-start justify-center">
      <div class="flex-1">
        <IconButton @click="openAddModal">
          <AddIcon />
        </IconButton>
      </div>
      <div class="flex-2 grid grid-flow-col gap-2">
        <span class="p-2 rounded-md bg-green-600 bg-opacity-30">Total Earnings: {{ maskCurrency(earningsTotal) }}</span>

        <span class="p-2 rounded-md bg-red-600 bg-opacity-30">Total Expenses: {{ maskCurrency(expensesTotal) }}</span>

        <span class="p-2 rounded-md" 
        :class="{ 'bg-red-600': balance < 0, 'bg-green-600': balance > 0, 'bg-gray-600': balance == 0 }">Balance: {{ maskCurrency(balance) }}</span>
      </div>
    </div>

    <DestructiveActionModal v-if="isDeleteModalOpen" :is-open="isDeleteModalOpen" @on-confirm="deleteItem" @on-cancel="closeDeleteModal" />
    <EditItemModal v-if="isEditModalOpen && selectedTabId" @updateCurrentItem="editItem" :selected-tab-id="selectedTabId" :is-open="isEditModalOpen" @on-cancel="closeEditModal" />
    <AddItemModal v-if="isAddModalOpen && selectedTabId" :selected-tab-id="selectedTabId" :is-open="isAddModalOpen" @addItem="addItem" @on-cancel="closeAddModal" />
</template>