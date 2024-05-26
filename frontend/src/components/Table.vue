<script setup lang="ts">
import { computed, onMounted, ref, watch } from "vue";
import EditIcon from "./ui/icons/EditIcon.vue";
import DeleteIcon from "./ui/icons/DeleteIcon.vue";
import IconButton from "./ui/IconButton.vue";
import AddIcon from "./ui/icons/AddIcon.vue";
import { maskCurrency, maskDate } from "@/shared/utils";
import SortArrowIcon from "./ui/icons/SortArrowIcon.vue";
import PaginationTab from "./ui/PaginationTab.vue";
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
import { Tab } from "@/shared/types";

const profileStore = useProfileStore();

const { profile } = storeToRefs(profileStore);

const earningsPagination = ref<types.PaginationOutput>({
  page: 1,
  size: 5,
  total_pages: 1,
  total_items: 1,
  next_page: 1,
  prev_page: 1,
  order_by: "",
  sort_by: "",
});

const expensesPagination = ref<types.PaginationOutput>({
  page: 1,
  size: 5,
  total_pages: 1,
  total_items: 1,
  next_page: 1,
  prev_page: 1,
  order_by: "",
  sort_by: "",
});

const getSelectedTabFromLocalStorage = () => {
  const tab = localStorage.getItem("selectedTab");
  if (tab) {
    selectedTab.value = tab as Tab;
  }
};

const setSelectedTabOnLocalStorage = (tab: Tab) => {
  localStorage.setItem("selectedTab", tab);
};

const isDeleteTableDataModalOpen = ref(false);
const isUpdateProfileModalOpen = ref(false);
const isEditModalOpen = ref(false);
const isAddModalOpen = ref(false);
const selectedRowId = ref(0);
const selectedTab = ref();

const openDeleteTableDataModal = (id: number) => {
  selectedRowId.value = id;
  isDeleteTableDataModalOpen.value = true;
};

const closeDeleteTableDataModal = () => {
  selectedRowId.value = 0;
  isDeleteTableDataModalOpen.value = false;
};

const deleteTableData = async () => {
  if (selectedTab.value === Tab.EARNING) {
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
  if (selectedTab.value === Tab.EARNING) {
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
  if (profile?.value === null) {
    console.error("Profile is null");
    return;
  }

  if (selectedTab.value === Tab.EARNING) {
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

const tabs: Tab[] = [Tab.EARNING, Tab.EXPENSE];

const earnings = ref<types.EarningOutput[]>([]);
const expenses = ref<types.ExpenseOutput[]>([]);
const isLoading = ref(false);

const fetchEarnings = async () => {
  try {
    const e = await findAllEarnings(
      profile?.value?.id || 1,
      earningsPagination.value,
    );
    earnings.value = e.data;
    earningsPagination.value = e.pagination;
  } catch (error) {
    console.error(error);
  }
};

const fetchExpenses = async () => {
  try {
    const e = await findAllExpenses(
      profile?.value?.id || 1,
      expensesPagination.value,
    );
    expenses.value = e.data;
    expensesPagination.value = e.pagination;
  } catch (error) {
    console.error(error);
  }
};

const changePage = (page: number) => {
  if (selectedTab.value === Tab.EARNING) {
    earningsPagination.value.page = page;
  } else {
    expensesPagination.value.page = page;
  }
  fetchData();
};

const fetchData = async () => {
  isLoading.value = true;
  await fetchEarnings();
  await fetchExpenses();
  console.log(earningsPagination.value);
  console.log(expensesPagination.value);
  isLoading.value = false;
};

onMounted(() => {
  getSelectedTabFromLocalStorage();

  watch(selectedTab, (newVal) => {
    setSelectedTabOnLocalStorage(newVal);
  });

  fetchData();
});

const handleChangeTab = (tab: Tab) => {
  selectedTab.value = tab;
  fetchData();
};

const tableData = computed(
  (): types.EarningOutput[] | types.ExpenseOutput[] => {
    return selectedTab.value === Tab.EARNING ? earnings.value : expenses.value;
  },
);

const sortTableData = (orderBy: string) => {
  if (selectedTab.value === Tab.EARNING) {
    earningsPagination.value.order_by = orderBy;
    earningsPagination.value.sort_by =
      earningsPagination.value.sort_by === "ASC" ? "DESC" : "ASC";
  } else {
    expensesPagination.value.order_by = orderBy;
    expensesPagination.value.sort_by =
      expensesPagination.value.sort_by === "ASC" ? "DESC" : "ASC";
  }
  fetchData();
};

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
      :key="tab"
      @click="handleChangeTab(tab)"
      class="py-1.5 px-2 duration-200 text-black dark:text-white rounded-sm"
      :class="{
        'bg-primary dark:bg-secondary text-black dark:text-black':
          selectedTab === tab,
        'hover:bg-dark dark:hover:bg-light hover:bg-opacity-30':
          selectedTab !== tab,
      }"
    >
      {{ tab }}
    </button>
    <div class="grid place-items-center grid-flow-col absolute right-10">
      <ProfileSelector
        :profile="profile!"
        @on-profile-create="fetchData()"
        @on-profile-change="fetchData()"
      />
      <button
        @click="openUpdateProfileModal"
        class="py-1.5 px-2 duration-100 grid w-24 place-items-center grid-flow-col rounded-sm text-black dark:text-white hover:bg-dark dark:hover:bg-light hover:bg-opacity-30 dark:hover:bg-opacity-30"
      >
        Edit profile
      </button>
    </div>
  </div>
  <div v-if="tableData && tableData.length > 0">
    <table class="min-w-full rounded-lg leading-normal">
      <thead>
        <tr class="text-center">
          <TableHeader
            >Description
            <SortArrowIcon
              @click="sortTableData('description')"
              :class="{
                'rotate-180':
                  (earningsPagination.order_by.toLowerCase() ===
                    'description' &&
                    earningsPagination.sort_by.toLowerCase() === 'asc') ||
                  (expensesPagination.order_by.toLowerCase() ===
                    'description' &&
                    expensesPagination.sort_by.toLowerCase() === 'asc'),
              }"
          /></TableHeader>
          <TableHeader
            >Amount
            <SortArrowIcon
              @click="sortTableData('amount')"
              :class="{
                'rotate-180':
                  (earningsPagination.order_by.toLowerCase() === 'amount' &&
                    earningsPagination.sort_by.toLowerCase() === 'asc') ||
                  (expensesPagination.order_by.toLowerCase() === 'amount' &&
                    expensesPagination.sort_by.toLowerCase() === 'asc'),
              }"
          /></TableHeader>
          <TableHeader v-if="selectedTab === Tab.EXPENSE"
            >Category
            <SortArrowIcon
              @click="sortTableData('category')"
              :class="{
                'rotate-180':
                  expensesPagination.order_by.toLowerCase() === 'category' &&
                  expensesPagination.sort_by.toLowerCase() === 'asc',
              }"
          /></TableHeader>
          <TableHeader
            >Created
            <SortArrowIcon
              @click="sortTableData('created_at')"
              :class="{
                'rotate-180':
                  (earningsPagination.order_by.toLowerCase() === 'created_at' &&
                    earningsPagination.sort_by.toLowerCase() === 'asc') ||
                  (expensesPagination.order_by.toLowerCase() === 'created_at' &&
                    expensesPagination.sort_by.toLowerCase() === 'asc'),
              }"
          /></TableHeader>
          <TableHeader
            >Updated
            <SortArrowIcon
              @click="sortTableData('updated_at')"
              :class="{
                'rotate-180':
                  (earningsPagination.order_by.toLowerCase() === 'updated_at' &&
                    earningsPagination.sort_by.toLowerCase() === 'asc') ||
                  (expensesPagination.order_by.toLowerCase() === 'updated_at' &&
                    expensesPagination.sort_by.toLowerCase() === 'asc'),
              }"
          /></TableHeader>
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
          <TableData v-if="selectedTab === Tab.EXPENSE">{{
            (item as types.ExpenseOutput).category
          }}</TableData>
          <TableData>{{ maskDate(item.created_at) }}</TableData>
          <TableData>{{ maskDate(item.updated_at) }}</TableData>
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
  </div>
  <div v-else class="text-center text-black dark:text-white py-4">
    No data available yet
  </div>

  <div class="w-full mt-4 flex items-center justify-center">
    <div class="flex-1">
      <IconButton class="text-black dark:text-white" @click="openAddModal">
        <AddIcon />
      </IconButton>
    </div>
    <div class="flex-1 self-center">
      <PaginationTab
        :pagination="
          selectedTab === Tab.EARNING ? earningsPagination : expensesPagination
        "
        @on-page-change="changePage"
      />
    </div>
    <div class="flex-1 grid grid-flow-col gap-2 text-sm min-w-fit">
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
    v-if="isEditModalOpen && selectedTab"
    :selected-row-id="selectedRowId"
    @editItem="editItem"
    :selected-tab="selectedTab"
    :is-open="isEditModalOpen"
    @on-cancel="closeEditModal"
  />
  <AddItemModal
    v-if="isAddModalOpen && selectedTab"
    :selected-tab="selectedTab"
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
