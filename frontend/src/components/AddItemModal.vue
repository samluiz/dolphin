<script setup lang="ts">
import { PropType, Ref, onMounted, ref, toRefs } from "vue";
import Dialog from "./ui/Dialog.vue";
import Input from "./ui/Input.vue";
import { FindAll as findAllCategories } from "../../wailsjs/go/category/service";
import { types } from "wailsjs/go/models";
import Select from "./ui/Select.vue";
import ConfirmButton from "./ui/ConfirmButton.vue";
import CancelButton from "./ui/CancelButton.vue";
import { Tab } from "@/shared/types";

const props = defineProps({
  isOpen: Boolean,
  selectedTab: String as PropType<Tab>,
});

const { isOpen, selectedTab } = toRefs(props);

const emit = defineEmits(["addItem", "on-cancel"]);

const categories = ref<types.Category[]>([]);

onMounted(() => {
  findAllCategories()
    .then((c: types.Category[]) => {
      categories.value = c;
    })
    .catch((e: any) => {
      console.error(e);
    });
});

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
  emit("addItem", formData.value);
}
</script>

<template>
  <Dialog :isOpen="isOpen">
    <form class="flex flex-col gap-4" @submit.prevent="onSubmit">
      <h1 class="text-2xl text-black dark:text-light-text">
        {{ $t("create") }}
        {{ selectedTab === Tab.EARNING ? $t("earning") : $t("expense") }}
      </h1>
      <Input
        :required="true"
        :name="'description'"
        :title="$t('description')"
        :type="'text'"
        :model-value="formData.description"
        @update:model-value="(newValue) => (formData.description = newValue)"
      />
      <Input
        :required="true"
        :name="'amount'"
        :title="$t('amount')"
        :type="'number'"
        :model-value="formData.amount"
        @update:model-value="(newValue) => (formData.amount = newValue)"
      />
      <Select
        :options="categories"
        :required="true"
        v-if="selectedTab === Tab.EXPENSE"
        :name="'category'"
        :title="$t('category')"
        :model-value="(formData as types.ExpenseUpdate).category_id"
        @update:model-value="
          (newValue) =>
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
