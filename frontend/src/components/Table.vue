<script setup lang="ts">
import { Ref, computed, ref, watch } from 'vue';
import EditIcon from './ui/icons/EditIcon.vue';
import DeleteIcon from './ui/icons/DeleteIcon.vue';
import IconButton from './ui/IconButton.vue';
import AddIcon from './ui/icons/AddIcon.vue';
import { Expense, Earning, Item, ExpenseInput, EarningInput } from '@/shared/types'
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
  selectedRowId.value = id;
  isDeleteModalOpen.value = true;
}

const closeDeleteModal = () => {
  selectedRowId.value = 0;
  isDeleteModalOpen.value = false;
}

const deleteItem = () => {
  if (selectedTabId.value === 1) {
    earnings.value.splice(earnings.value.findIndex((earning: Earning) => earning.id === selectedRowId.value), 1);
  } else {
    expenses.value.splice(expenses.value.findIndex((expense: Expense) => expense.id === selectedRowId.value), 1);
  }
  isDeleteModalOpen.value = false;
  trigger.value += 1;
}

const openEditModal = (id: number) => {
  selectedRowId.value = id;
  isEditModalOpen.value = true;
}

const closeEditModal = () => {
  selectedRowId.value = 0;
  isEditModalOpen.value = false;
}

const editItem = (formData: Item) => {
  if (selectedTabId.value === 1) {
    editEarning(formData);
  } else {
    editExpense(formData as ExpenseInput);
  }
  selectedRowId.value = 0;
  isEditModalOpen.value = false;
}

function editEarning(formData: EarningInput) {
  earnings.value.forEach((earning: Earning) => {
    if (earning.id === selectedRowId.value) {
      earning.description = formData.description;
      earning.amount = formData.amount;
    }
  });
}

function editExpense(formData: ExpenseInput) {
  expenses.value.forEach((expense: Expense) => {
    if (expense.id === selectedRowId.value) {
      expense.description = formData.description;
      expense.amount = formData.amount;
      expense.category = formData.category;
    }
  });
}

const openAddModal = () => {
  isAddModalOpen.value = true;
}

const closeAddModal = () => {
  isAddModalOpen.value = false;
}

const addItem = (formData: Item) => {
  if (selectedTabId.value === 1) {
    addEarning(formData);
  } else {
    addExpense(formData);
  }
  isAddModalOpen.value = false;
  trigger.value += 1;
}

function addEarning(formData: Item) {
  console.log(formData)
  earnings.value.push({ id: earnings.value.length + 1, ...formData });
}

function addExpense(formData: Item) {
  expenses.value.push({ ...formData, id: expenses.value.length + 1, category: formData.category as string});
}

interface Tab {
  id: number
  name: string
}

const tabs: Tab[] = [
  { id: 1, name: 'Earnings' },
  { id: 2, name: 'Expenses' },
]

const earnings: Ref<Earning[]> = ref([])

const expenses: Ref<Expense[]> = ref([])

const tableData = computed((): Earning[] | Expense[] => {
  return selectedTabId.value === 1 ? earnings.value : expenses.value;
});

const earningsTotal = computed(() => {
  let total = 0;
  earnings.value.forEach((earning: Earning) => {
    total += earning.amount;
  });
  return total;
});

const expensesTotal = computed(() => {
  let total = 0;
  expenses.value.forEach((expense: Expense) => {
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
    <EditItemModal v-if="isEditModalOpen && selectedTabId" :current="tableData.find((item: Earning | Expense) => item.id === selectedRowId)" @editItem="editItem" :selected-tab-id="selectedTabId" :is-open="isEditModalOpen" @on-cancel="closeEditModal" />
    <AddItemModal v-if="isAddModalOpen && selectedTabId" :selected-tab-id="selectedTabId" :is-open="isAddModalOpen" @addItem="addItem" @on-cancel="closeAddModal" />
</template>