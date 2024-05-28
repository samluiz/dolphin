<script setup lang="ts">
import { Ref, onMounted, ref, toRefs } from "vue";
import Input from "./ui/Input.vue";
import Checkbox from "./ui/Checkbox.vue";
import ConfirmButton from "./ui/ConfirmButton.vue";
import CancelButton from "./ui/CancelButton.vue";
import DangerButton from "./ui/DangerButton.vue";
import Dialog from "./ui/Dialog.vue";
import { useProfileStore } from "@/stores/ProfileStore";
import { types } from "../../wailsjs/go/models";
import DestructiveActionModal from "./DestructiveActionModal.vue";

const profileStore = useProfileStore();

const emit = defineEmits([
  "on-cancel",
  "on-profile-update",
  "on-profile-delete",
]);

const props = defineProps({
  isOpen: Boolean,
  profile: Object as () => types.Profile,
});

const { isOpen, profile } = toRefs(props);

const isDeleteProfileModalOpen = ref(false);
const isOnlyProfile = ref(false);

const formData: Ref<types.ProfileInput> = ref({
  description: profile?.value?.description!,
  is_default: profile?.value?.is_default!,
});

function onSubmit(): void {
  console.log(formData?.value);

  profileStore
    .update(profile?.value?.id!, {
      description: formData?.value?.description!,
      is_default: !isOnlyProfile.value ? formData?.value?.is_default! : true,
    })
    .then(() => {
      emit("on-profile-update", profile?.value!);
      emit("on-cancel");
    });
}

function deleteProfile() {
  profileStore.remove(profile?.value?.id!).then(() => {
    emit("on-cancel");
  });
}

const openDeleteProfileModal = () => {
  isDeleteProfileModalOpen.value = true;
};

const closeDeleteProfileModal = () => {
  isDeleteProfileModalOpen.value = false;
};

onMounted(() => {
  if (profileStore.isOnlyProfile()) {
    console.log("isOnlyProfile", profileStore.isOnlyProfile());
    isOnlyProfile.value = true;
  }
});
</script>

<template>
  <Dialog :isOpen="isOpen">
    <form class="flex flex-col gap-4" @submit.prevent="onSubmit">
      <h1 class="text-2xl text-black dark:text-light-text">
        {{ $t("edit_profile") }}
      </h1>
      <Input
        :required="true"
        :name="'description'"
        :title="$t('description')"
        :type="'text'"
        :model-value="formData.description"
        @update:model-value="(newValue) => (formData.description = newValue)"
      />
      <Checkbox
        :disabled="profile?.is_default! && !isOnlyProfile"
        :required="false"
        :name="'default'"
        :title="$t('main_profile')"
        :type="'checkbox'"
        :checked="formData.is_default"
        :model-value="formData.is_default"
        @change="
          (e) => {
            console.log(e.target.checked);
            formData.is_default = e.target.checked;
          }
        "
      />
      <div class="flex gap-4">
        <ConfirmButton>{{ $t("save") }}</ConfirmButton>
        <CancelButton @click="$emit('on-cancel')">{{
          $t("cancel")
        }}</CancelButton>
        <DangerButton
          v-if="!isOnlyProfile && !profile?.is_default"
          @click="openDeleteProfileModal"
          >{{ $t("delete") }}</DangerButton
        >
      </div>
    </form>
  </Dialog>

  <DestructiveActionModal
    v-if="isDeleteProfileModalOpen"
    :is-open="isDeleteProfileModalOpen"
    @on-confirm="deleteProfile"
    @on-cancel="closeDeleteProfileModal"
    :title="$t('delete') + ' ' + $t('profile')"
  />
</template>
