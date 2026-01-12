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
import BalanceTab from "./BalanceTab.vue";
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
import { pagination, types } from "../../wailsjs/go/models";
import ProfileSelector from "./ProfileSelector.vue";
import { useProfileStore } from "@/stores/ProfileStore";
import { storeToRefs } from "pinia";
import EditProfileModal from "./EditProfileModal.vue";
import { Tab } from "@/shared/types";
import { useI18n } from "vue-i18n";

const { t, d, n } = useI18n();

const profileStore = useProfileStore();

const { profile } = storeToRefs(profileStore);

const itemsPerPage = parseInt(localStorage.getItem("itemsPerPage") || "5");

console.log(itemsPerPage);

const earningsPagination = ref<pagination.PaginationOutput>({
  page: 1,
  size: itemsPerPage,
  total_pages: 1,
  total_items: 1,
  next_page: 1,
  prev_page: 1,
  order_by: "",
  sort_by: "",
});

const expensesPagination = ref<pagination.PaginationOutput>({
  page: 1,
  size: itemsPerPage,
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
  } else {
    selectedTab.value = Tab.EARNING;
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

const handleSizeChange = (n: number) => {
  if (selectedTab.value === Tab.EARNING) {
    earningsPagination.value.size = n;
  } else {
    expensesPagination.value.size = n;
  }
  fetchData();
};
</script>

<template>
  <div class="flex justify-between w-full items-center my-6 gap-2">
    <div class="w-full flex justify-start">
      <ProfileSelector
        :profile="profile!"
        @on-profile-create="fetchData()"
        @on-profile-change="fetchData()"
      />
      <button
        @click="openUpdateProfileModal"
        class="py-1.5 px-2 duration-100 grid w-24 place-items-center grid-flow-col rounded-sm text-black dark:text-light-text hover:bg-dark dark:hover:bg-light hover:bg-opacity-30 dark:hover:bg-opacity-30"
      >
        {{ t("edit_profile") }}
      </button>
    </div>
    <div class="w-full flex justify-center gap-2">
      <button
        v-for="tab in tabs"
        :key="tab"
        @click="handleChangeTab(tab)"
        class="relative px-4 py-2 rounded-lg font-medium text-sm transition-all duration-200 ease-out transform hover:scale-105"
        :class="{
          'bg-gradient-ocean text-white shadow-glow-blue':
            selectedTab && selectedTab === tab,
          'text-text-primary-light dark:text-text-primary-dark hover:bg-primary-50 dark:hover:bg-surface-dark':
            selectedTab && selectedTab !== tab,
        }"
      >
        {{ tab === Tab.EARNING ? t("earning", 2) : t("expense", 2) }}
      </button>
    </div>
    <div
      class="w-full flex justify-end gap-2 text-sm text-black dark:text-light-text"
    ></div>
  </div>

  <div
    v-if="tableData && tableData.length > 0"
    class="overflow-x-auto rounded-xl border border-border-light dark:border-border-dark shadow-lg-blue"
  >
    <table class="w-full bg-white dark:bg-surface-dark">
      <thead>
        <tr class="text-center">
          <TableHeader>
            <button
              @click="sortTableData('description')"
              class="flex items-center gap-1 cursor-pointer hover:text-primary-500 transition-colors"
            >
              <span>{{ t("description") }}</span>
              <span class="text-xs">
                {{
                  earningsPagination.order_by.toLowerCase() === "description" ||
                  expensesPagination.order_by.toLowerCase() === "description"
                    ? earningsPagination.sort_by.toLowerCase() === "asc" ||
                      expensesPagination.sort_by.toLowerCase() === "asc"
                      ? "↑"
                      : "↓"
                    : "↕"
                }}
              </span>
            </button>
          </TableHeader>
          <TableHeader>
            <button
              @click="sortTableData('amount')"
              class="flex items-center gap-1 cursor-pointer hover:text-primary-500 transition-colors"
            >
              <span>{{ t("amount") }}</span>
              <span class="text-xs">
                {{
                  earningsPagination.order_by.toLowerCase() === "amount" ||
                  expensesPagination.order_by.toLowerCase() === "amount"
                    ? earningsPagination.sort_by.toLowerCase() === "asc" ||
                      expensesPagination.sort_by.toLowerCase() === "asc"
                      ? "↑"
                      : "↓"
                    : "↕"
                }}
              </span>
            </button>
          </TableHeader>
          <TableHeader>
            <button
              @click="sortTableData('created_at')"
              class="flex items-center gap-1 cursor-pointer hover:text-primary-500 transition-colors"
            >
              <span>{{ t("created") }}</span>
              <span class="text-xs">
                {{
                  earningsPagination.order_by.toLowerCase() === "created_at" ||
                  expensesPagination.order_by.toLowerCase() === "created_at"
                    ? earningsPagination.sort_by.toLowerCase() === "asc" ||
                      expensesPagination.sort_by.toLowerCase() === "asc"
                      ? "↑"
                      : "↓"
                    : "↕"
                }}
              </span>
            </button>
          </TableHeader>
          <TableHeader>
            <button
              @click="sortTableData('updated_at')"
              class="flex items-center gap-1 cursor-pointer hover:text-primary-500 transition-colors"
            >
              <span>{{ t("updated") }}</span>
              <span class="text-xs">
                {{
                  earningsPagination.order_by.toLowerCase() === "updated_at" ||
                  expensesPagination.order_by.toLowerCase() === "updated_at"
                    ? earningsPagination.sort_by.toLowerCase() === "asc" ||
                      expensesPagination.sort_by.toLowerCase() === "asc"
                      ? "↑"
                      : "↓"
                    : "↕"
                }}
              </span>
            </button>
          </TableHeader>
          <TableHeader v-if="selectedTab === Tab.EXPENSE">
            <button
              @click="sortTableData('category')"
              class="flex items-center gap-1 cursor-pointer hover:text-primary-500 transition-colors"
            >
              <span>{{ t("category") }}</span>
              <span class="text-xs">
                {{
                  expensesPagination.order_by.toLowerCase() === "category"
                    ? expensesPagination.sort_by.toLowerCase() === "asc"
                      ? "↑"
                      : "↓"
                    : "↕"
                }}
              </span>
            </button>
          </TableHeader>
          <TableHeader v-else></TableHeader>
          <TableHeader></TableHeader>
        </tr>
      </thead>
      <tbody>
        <tr
          class="text-center transition-all duration-200 ease-out hover:bg-primary-50 dark:hover:bg-surface-dark/50 odd:bg-white even:bg-gray-50/50 dark:odd:bg-surface-dark dark:even:bg-bg-dark/50 group"
          v-for="item in tableData"
          :key="item.id"
        >
          <TableData>{{ item.description }}</TableData>
          <TableData>
            <span class="font-semibold text-primary-600 dark:text-primary-400">
              {{ n(item.amount, "currency") }}
            </span>
          </TableData>
          <TableData>{{ d(new Date(item.created_at), "short") }}</TableData>
          <TableData>{{ d(new Date(item.updated_at), "short") }}</TableData>
          <TableData v-if="selectedTab === Tab.EXPENSE">
            <span
              class="px-2 py-1 rounded-full text-xs font-medium bg-primary-100 dark:bg-primary-900/30 text-primary-700 dark:text-primary-300"
            >
              {{ t((item as types.ExpenseOutput).category.toLowerCase()) }}
            </span>
          </TableData>
          <TableData v-else></TableData>
          <TableData>
            <div
              class="flex flex-row justify-end gap-1 opacity-0 group-hover:opacity-100 transition-opacity duration-200"
            >
              <IconButton @click="openEditModal(item.id)">
                <EditIcon />
              </IconButton>
              <IconButton @click="openDeleteTableDataModal(item.id)">
                <DeleteIcon />
              </IconButton>
            </div>
          </TableData>
        </tr>
      </tbody>
    </table>
  </div>
  <div v-else class="text-center text-black dark:text-light-text py-4">
    {{ t("no_data") }}
  </div>

  <div class="w-full mt-4 flex justify-between items-center">
    <div class="w-1/3 flex justify-start items-center">
      <button
        class="py-1.5 px-2 duration-100 grid gap-2 place-items-center grid-flow-col rounded-sm text-black dark:text-light-text hover:bg-dark dark:hover:bg-light hover:bg-opacity-30 dark:hover:bg-opacity-30"
        @click="openAddModal"
      >
        <AddIcon />
        <span class="text-black dark:text-light-text">{{
          selectedTab === Tab.EARNING ? t("new_earning") : t("new_expense")
        }}</span>
      </button>
    </div>
    <div class="w-1/3 flex justify-center">
      <PaginationTab
        :pagination="
          selectedTab === Tab.EARNING ? earningsPagination : expensesPagination
        "
        @on-page-change="changePage"
      />
    </div>
    <div class="w-1/3 flex justify-end">
      <BalanceTab
        :earnings_subtotal="
          earnings && earnings.length > 0 ? earnings[0].sub_total : 0
        "
        :expenses_subtotal="
          expenses && expenses.length > 0 ? expenses[0].sub_total : 0
        "
      />
    </div>
  </div>

  <DestructiveActionModal
    v-if="isDeleteTableDataModalOpen"
    :is-open="isDeleteTableDataModalOpen"
    @on-confirm="deleteTableData"
    @on-cancel="closeDeleteTableDataModal"
    :title="
      selectedTab === Tab.EARNING
        ? t('delete') + ' ' + t('earning')
        : t('delete') + ' ' + t('expense')
    "
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
