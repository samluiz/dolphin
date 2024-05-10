<script setup lang="ts">
import { Ref, onMounted, ref, toRefs, watch } from "vue";
import Dialog from "./ui/Dialog.vue";
import Input from "./ui/Input.vue";
import Select from "./ui/Select.vue";
import {
  Create as createCategory,
  FindAll as findAllCategories,
} from "../../wailsjs/go/category/service";
import { FindByID as findEarningById } from "../../wailsjs/go/earning/service";
import { FindByID as findExpenseById } from "../../wailsjs/go/expense/service";
import { types } from "wailsjs/go/models";
import ConfirmButton from "./ui/ConfirmButton.vue";
import CancelButton from "./ui/CancelButton.vue";

const props = defineProps({
  isOpen: Boolean,
  selectedTabId: Number,
  selectedRowId: Number,
});

const { isOpen, selectedTabId, selectedRowId } = toRefs(props);

const emit = defineEmits(["editItem", "on-cancel"]);

const currentRow = ref<types.EarningToUpdate | types.ExpenseToUpdate>();
const categories = ref<types.Category[]>([]);

onMounted(() => {
  fetchData(selectedTabId?.value!, selectedRowId?.value!);
});

const fetchData = async (tabId: number, rowId: number) => {
  await fetchCategories();
  if (tabId === 1) {
    try {
      const earning = await findEarningById(rowId);
      currentRow.value = earning;
      formData.value = { ...currentRow.value };
    } catch (error) {
      console.error(error);
    }
  } else {
    try {
      const expense = await findExpenseById(rowId);
      currentRow.value = expense;
      formData.value = {
        ...currentRow.value,
        category_id: categories.value.find(
          (cat) => cat.description === expense.category,
        )?.id,
      };
    } catch (error) {
      console.error(error);
    }
  }
};

const fetchCategories = async () => {
  try {
    const c = await findAllCategories();
    categories.value = c;
  } catch (error) {
    console.error(error);
  }
};

const formData: Ref<types.EarningUpdate | types.ExpenseUpdate> = ref({
  description: "",
  amount: undefined as unknown as number,
  category_id: 0,
});

function getCategoryId(description: string): number {
  let categoryId = 0;
  const category = categories.value.find(
    (cat) => cat.description === description,
  );
  if (category) {
    categoryId = category.id;
  }
  return categoryId;
}

function onSubmit(): void {
  formData.value.amount = Number(formData.value.amount);
  (formData.value as types.ExpenseUpdate).category_id = Number(
    (formData.value as types.ExpenseUpdate).category_id,
  );
  emit("editItem", formData.value);
}
</script>

<template>
  <Dialog :isOpen="isOpen">
    <form class="flex flex-col gap-4" @submit.prevent="onSubmit">
      <h1 class="text-2xl text-black dark:text-white">
        Edit {{ selectedTabId === 1 ? "earning" : "expense" }}
      </h1>
      <Input
        :step="0.01"
        :required="true"
        :name="'description'"
        :title="'Description'"
        :type="'text'"
        :model-value="formData && formData.description"
        @update:model-value="(newValue) => (formData.description = newValue)"
      />
      <Input
        :required="true"
        :name="'amount'"
        :title="'Amount'"
        :type="'number'"
        :model-value="formData && formData.amount"
        @update:model-value="(newValue) => (formData.amount = newValue)"
      />
      <Select
        :options="categories"
        :required="true"
        v-if="selectedTabId === 2"
        :name="'category'"
        :title="'Category'"
        :model-value="formData && (formData as types.ExpenseUpdate).category_id"
        @update:model-value="
          (newValue) =>
            ((formData as types.ExpenseUpdate).category_id = newValue)
        "
      />
      <div class="flex gap-4">
        <CancelButton @click="emit('on-cancel')">Cancel</CancelButton>
        <ConfirmButton>Edit</ConfirmButton>
      </div>
    </form>
  </Dialog>
</template>
