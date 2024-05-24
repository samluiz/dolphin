<script setup lang="ts">
import { ref, toRefs } from "vue";
import Input from "./ui/Input.vue";
import Checkbox from "./ui/Checkbox.vue";
import ConfirmButton from "./ui/ConfirmButton.vue";
import CancelButton from "./ui/CancelButton.vue";
import Dialog from "./ui/Dialog.vue";
import { useProfileStore } from "@/stores/ProfileStore";

const profileStore = useProfileStore();

const emit = defineEmits(["on-cancel", "on-profile-create"]);

const props = defineProps({
  isOpen: Boolean,
});

const { isOpen } = toRefs(props);

const description = ref("");
const isDefault = ref(false);

function onSubmit(): void {
  profileStore
    .create({
      description: description.value,
      is_default: !profileStore.hasProfileCreated() ? true : isDefault.value,
    })
    .then((data) => {
      emit("on-profile-create", data);
      emit("on-cancel");
    });
}
</script>

<template>
  <Dialog :isOpen="isOpen">
    <form class="flex flex-col gap-4" @submit.prevent="onSubmit">
      <h1 class="text-2xl text-black dark:text-white">Create profile</h1>
      <Input
        :required="true"
        :name="'description'"
        :title="'Description'"
        :type="'text'"
        :model-value="description"
        @update:model-value="(newValue) => (description = newValue)"
      />
      <Checkbox
        v-if="profileStore.hasProfileCreated()"
        :required="false"
        :name="'default'"
        :title="'Main profile'"
        :type="'checkbox'"
        :checked="false"
        :model-value="isDefault"
        @change="
          (e) => {
            console.log(e.target.checked);
            isDefault = e.target.checked;
          }
        "
      />
      <div class="flex gap-4">
        <CancelButton @click="$emit('on-cancel')">Cancel</CancelButton>
        <ConfirmButton>Add</ConfirmButton>
      </div>
    </form>
  </Dialog>
</template>
