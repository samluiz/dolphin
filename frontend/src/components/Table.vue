<script setup lang="ts">
import { computed, onMounted, ref} from "vue";
import EditIcon from "./ui/icons/EditIcon.vue";
import DeleteIcon from "./ui/icons/DeleteIcon.vue";
import IconButton from "./ui/IconButton.vue";
import AddIcon from "./ui/icons/AddIcon.vue";
import { maskCurrency } from "@/shared/utils";
import DestructiveActionModal from "./DestructiveActionModal.vue";
import EditItemModal from "./EditItemModal.vue";
import AddItemModal from "./AddItemModal.vue";
import TableHeader from "./ui/TableHeader.vue";
import TableData from "./ui/TableData.vue";
import {
  Create as createEarning,
  Update as updateEarning,
  FindAllByProfileID as findAllEarnings,
  Delete as deleteEarning,
} from "../../wailsjs/go/earning/service";
import {
  Create as createExpense,
  Update as updateExpense,
  FindAllByProfileID as findAllExpenses,
  Delete as deleteExpense,
} from "../../wailsjs/go/expense/service";
import { types } from "../../wailsjs/go/models";
import ProfileSelector from "./ProfileSelector.vue";
import { useProfileStore } from "@/stores/ProfileStore";
import { storeToRefs } from "pinia";
import EditProfileModal from "./EditProfileModal.vue";

const profileStore = useProfileStore();

const { profile } = storeToRefs(profileStore);

const isDeleteTableDataModalOpen = ref(false);
const isUpdateProfileModalOpen = ref(false);
const isEditModalOpen = ref(false);
const isAddModalOpen = ref(false);
const selectedRowId = ref(0);
const selectedTabId = ref(1);

const openDeleteTableDataModal = (id: number) => {
  selectedRowId.value = id;
  isDeleteTableDataModalOpen.value = true;
};

const closeDeleteTableDataModal = () => {
  selectedRowId.value = 0;
  isDeleteTableDataModalOpen.value = false;
};

const deleteTableData = async () => {
  if (selectedTabId.value === 1) {
    await deleteEarning(selectedRowId.value);
  } else {
    await deleteExpense(selectedRowId.value);
  }
  isDeleteTableDataModalOpen.value = false;
  fetchData();
};

const openEditModal = (id: number) => {
  selectedRowId.value = id;
  isEditModalOpen.value = true;
};

const closeEditModal = () => {
  selectedRowId.value = 0;
  isEditModalOpen.value = false;
};

const openUpdateProfileModal = () => {
  isUpdateProfileModalOpen.value = true;
};

const closeUpdateProfileModal = () => {
  isUpdateProfileModalOpen.value = false;
};

const editItem = (formData: types.EarningUpdate | types.ExpenseUpdate) => {
  if (selectedTabId.value === 1) {
    const updatedEarning: types.EarningUpdate = {
      description: formData.description,
      amount: formData.amount,
    };
    updateEarning(selectedRowId.value, updatedEarning)
      .then(() => {
        fetchData();
      })
      .catch((e: any) => {
        console.error(e);
      });
  } else {
    const updatedExpense: types.ExpenseUpdate = {
      description: formData.description,
      amount: formData.amount,
      category_id: (formData as types.ExpenseUpdate).category_id,
    };
    updateExpense(selectedRowId.value, updatedExpense)
      .then(() => {
        fetchData();
      })
      .catch((e: any) => {
        console.error(e);
      });
  }
  selectedRowId.value = 0;
  isEditModalOpen.value = false;
};

const openAddModal = () => {
  isAddModalOpen.value = true;
};

const closeAddModal = () => {
  isAddModalOpen.value = false;
};

const addItem = (formData: types.EarningUpdate | types.ExpenseUpdate) => {
  if (selectedTabId.value === 1) {
    addEarning(formData);
  } else {
    addExpense(formData);
  }
  isAddModalOpen.value = false;
};

function addEarning(formData: types.EarningUpdate | types.ExpenseUpdate) {
  const newEarning: types.EarningInput = {
    description: formData.description,
    amount: formData.amount,
    profile_id: profile?.value?.id || 1,
  };
  createEarning(newEarning)
    .then(() => {
      fetchData();
    })
    .catch((e: any) => {
      console.error(e);
    });
}

function addExpense(formData: types.EarningUpdate | types.ExpenseUpdate) {
  const newExpense: types.ExpenseInput = {
    description: formData.description,
    amount: formData.amount,
    category_id: (formData as types.ExpenseUpdate).category_id,
    profile_id: profile?.value?.id || 1,
  };

  createExpense(newExpense)
    .then(() => {
      fetchData();
    })
    .catch((e: any) => {
      console.error(e);
    });
}

interface Tab {
  id: number;
  name: string;
}

const tabs: Tab[] = [
  { id: 1, name: "Earnings" },
  { id: 2, name: "Expenses" },
];

const earnings = ref<types.EarningOutput[]>([]);
const expenses = ref<types.ExpenseOutput[]>([]);
const isLoading = ref(false);

const fetchEarnings = async () => {
  try {
    const e = await findAllEarnings(profile?.value?.id || 1);
    earnings.value = e;
  } catch (error) {
    console.error(error);
  }
};

const fetchExpenses = async () => {
  try {
    const e = await findAllExpenses(profile?.value?.id || 1);
    expenses.value = e;
  } catch (error) {
    console.error(error);
  }
};

const fetchData = async () => {
  isLoading.value = true;
  await fetchEarnings();
  await fetchExpenses();
  isLoading.value = false;
};

onMounted(() => {
  fetchData();
});

const tableData = computed(
  (): types.EarningOutput[] | types.ExpenseOutput[] => {
    return selectedTabId.value === 1 ? earnings.value : expenses.value;
  },
);

const totalEarnings = computed(() => {
  return earnings.value && earnings.value.length > 0
    ? earnings.value[0].sub_total
    : 0;
});

const totalExpenses = computed(() => {
  return expenses.value && expenses.value.length > 0
    ? expenses.value[0].sub_total
    : 0;
});

const balance = computed(() => {
  return totalEarnings.value - totalExpenses.value;
});
</script>

<template>
  <div class="flex justify-center items-center mb-4 gap-2">
    <button
      v-for="tab in tabs"
      :key="tab.id"
      @click="selectedTabId = tab.id"
      class="py-1.5 px-2 text-black dark:text-white rounded-sm"
      :class="{
        'duration-200 bg-primary dark:bg-secondary text-black dark:text-black':
          selectedTabId === tab.id,
      }"
    >
      {{ tab.name }}
    </button>
    <div class="grid place-items-center grid-flow-col absolute right-10">
      <ProfileSelector
        @on-select="profileStore.setActiveProfile"
        @on-profile-create="fetchData()"
        @on-cancel="fetchData()"
      />
      <button
        @click="openUpdateProfileModal"
        class="py-1.5 px-2 duration-100 grid w-24 place-items-center grid-flow-col rounded-sm text-black dark:text-white hover:bg-dark dark:hover:bg-light hover:bg-opacity-30 dark:hover:bg-opacity-30"
      >
        Edit profile
      </button>
    </div>
  </div>

  <table
    v-if="tableData && tableData.length > 0"
    class="min-w-full rounded-lg leading-normal"
  >
    <thead>
      <tr class="text-center">
        <TableHeader>Description</TableHeader>
        <TableHeader>Amount</TableHeader>
        <TableHeader v-if="selectedTabId === 2">Category</TableHeader>
        <TableHeader></TableHeader>
      </tr>
    </thead>
    <tbody>
      <tr
        class="text-center odd:bg-primary even:bg-secondary text-black"
        v-for="item in tableData"
        :key="item.id"
      >
        <TableData>{{ item.description }}</TableData>
        <TableData>{{ maskCurrency(item.amount) }}</TableData>
        <TableData v-if="selectedTabId === 2">{{
          (item as types.ExpenseOutput).category
        }}</TableData>
        <TableData class="flex flex-row justify-end">
          <div class="flex flex-row justify-evenly w-20">
            <IconButton class="text-black" @click="openEditModal(item.id)">
              <EditIcon />
            </IconButton>
            <IconButton
              class="text-black"
              @click="openDeleteTableDataModal(item.id)"
            >
              <DeleteIcon />
            </IconButton>
          </div>
        </TableData>
      </tr>
    </tbody>
  </table>
  <div v-else class="text-center text-black dark:text-white py-4">
    No data available yet
  </div>

  <div class="w-full mt-4 flex items-start justify-center">
    <div class="flex-1">
      <IconButton class="text-black dark:text-white" @click="openAddModal">
        <AddIcon />
      </IconButton>
    </div>
    <div class="flex-2 grid grid-flow-col gap-2">
      <span
        class="p-2 rounded-md bg-green-600 bg-opacity-50 hover:bg-opacity-100 duration-200"
        >Total Earnings: {{ maskCurrency(totalEarnings) }}</span
      >

      <span class="p-2 rounded-md bg-red-600 bg-opacity-50 hover:bg-opacity-100"
        >Total Expenses: {{ maskCurrency(totalExpenses) }}</span
      >

      <span
        class="p-2 rounded-md"
        :class="{
          'bg-red-600': balance < 0,
          'bg-green-600': balance > 0,
          'bg-gray-600': balance == 0,
        }"
        >Balance: {{ maskCurrency(balance) }}</span
      >
    </div>
  </div>

  <DestructiveActionModal
    v-if="isDeleteTableDataModalOpen"
    :is-open="isDeleteTableDataModalOpen"
    @on-confirm="deleteTableData"
    @on-cancel="closeDeleteTableDataModal"
  />
  <EditItemModal
    v-if="isEditModalOpen && selectedTabId"
    :selected-row-id="selectedRowId"
    @editItem="editItem"
    :selected-tab-id="selectedTabId"
    :is-open="isEditModalOpen"
    @on-cancel="closeEditModal"
  />
  <AddItemModal
    v-if="isAddModalOpen && selectedTabId"
    :selected-tab-id="selectedTabId"
    :is-open="isAddModalOpen"
    @addItem="addItem"
    @on-cancel="closeAddModal"
  />
  <EditProfileModal
    v-if="isUpdateProfileModalOpen"
    :profile="profile!"
    :is-open="isUpdateProfileModalOpen"
    @on-profile-update="fetchData()"
    @on-profile-delete="fetchData()"
    @on-cancel="closeUpdateProfileModal"
  />
</template>
