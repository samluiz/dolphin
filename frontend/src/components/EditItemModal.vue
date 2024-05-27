<script setup lang="ts">
import { PropType, Ref, onMounted, ref, toRefs, watch } from "vue";
import Dialog from "./ui/Dialog.vue";
import Input from "./ui/Input.vue";
import Select from "./ui/Select.vue";
import { FindAll as findAllCategories } from "../../wailsjs/go/category/service";
import { FindByID as findEarningById } from "../../wailsjs/go/earning/service";
import { FindByID as findExpenseById } from "../../wailsjs/go/expense/service";
import { types } from "wailsjs/go/models";
import ConfirmButton from "./ui/ConfirmButton.vue";
import CancelButton from "./ui/CancelButton.vue";
import { Tab } from "@/shared/types";

const props = defineProps({
  isOpen: Boolean,
  selectedTab: String as PropType<Tab>,
  selectedRowId: Number,
});

const { isOpen, selectedTab, selectedRowId } = toRefs(props);

const emit = defineEmits(["editItem", "on-cancel"]);

const currentRow = ref<types.EarningToUpdate | types.ExpenseToUpdate>();
const categories = ref<types.Category[]>([]);

onMounted(() => {
  fetchData(selectedTab?.value!, selectedRowId?.value!);
});

const fetchData = async (tab: Tab, rowId: number) => {
  await fetchCategories();
  if (tab === Tab.EARNING) {
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
        {{ $t("edit") }}
        {{ selectedTab === Tab.EARNING ? $t("earning") : $t("expense") }}
      </h1>
      <Input
        :required="true"
        :name="'description'"
        :title="$t('description')"
        :type="'text'"
        :model-value="formData.description"
        @update:model-value="
          (newValue: string) => (formData.description = newValue)
        "
      />
      <Input
        :required="true"
        :name="'amount'"
        :title="$t('amount')"
        :type="'number'"
        :model-value="formData.amount"
        @update:model-value="(newValue: number) => (formData.amount = newValue)"
      />
      <Select
        :options="categories"
        :required="true"
        v-if="selectedTab === Tab.EXPENSE"
        :name="'category'"
        :title="$t('category')"
        :model-value="formData && (formData as types.ExpenseUpdate).category_id"
        @update:model-value="
          (newValue: number) =>
            ((formData as types.ExpenseUpdate).category_id = newValue)
        "
      />
      <div class="flex gap-4">
        <CancelButton @click="emit('on-cancel')">{{
          $t("cancel")
        }}</CancelButton>
        <ConfirmButton>{{ $t("save") }}</ConfirmButton>
      </div>
    </form>
  </Dialog>
</template>
