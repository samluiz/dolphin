<script setup lang="ts">
import { ref, defineEmits, computed } from "vue";
import { useProfileStore } from "@/stores/ProfileStore";
import { types } from "../../wailsjs/go/models";
import CreateProfileModal from "./CreateProfileModal.vue";
import ArrowIcon from "./ui/icons/ArrowIcon.vue";
import { storeToRefs } from "pinia";
import AddIcon from "./ui/icons/AddIcon.vue";

const emit = defineEmits(["on-select", "on-profile-create", "on-cancel"]);

const profileStore = useProfileStore();

const { profile, profiles } = storeToRefs(profileStore);

const selectedProfile = ref<types.Profile | null>(null);
const isCreateProfileModalOpen = ref(false);
const isDropdownOpen = ref(false);

const selectProfile = (newProfile: types.Profile) => {
  profileStore.setActiveProfile(newProfile);
  selectedProfile.value = newProfile;
  isDropdownOpen.value = false;
  emit("on-select", newProfile);
};

const openCreateProfileModal = () => {
  isCreateProfileModalOpen.value = true;
};

const closeCreateProfileModal = () => {
  isCreateProfileModalOpen.value = false;
  emit("on-cancel");
};

const onProfileCreate = () => {
  emit("on-profile-create", profile.value);
  emit("on-cancel");
};

const toggleDropdown = () => {
  isDropdownOpen.value = !isDropdownOpen.value;
};

const otherProfiles = computed(() => {
  return (
    profiles.value &&
    profiles.value.filter((profile) => profile.id !== profileStore.profile?.id)
  );
});
</script>

<template>
  <div @click="toggleDropdown">
    <button
      class="py-1.5 px-2 duration-100 grid min-w-24 place-items-center grid-flow-col rounded-sm text-black dark:text-white hover:bg-dark dark:hover:bg-light hover:bg-opacity-30 dark:hover:bg-opacity-30"
    >
      {{ profile ? profile.description : "Select profile" }} <ArrowIcon />
    </button>
    <ul
      v-if="isDropdownOpen"
      class="shadow-lg flex flex-col rounded-sm ring-1 ring-black ring-opacity-5 focus:outline-none absolute min-w-24 bg-light dark:bg-dark text-black dark:text-white z-50"
    >
      <li
        v-for="profile in otherProfiles"
        :key="profile.id"
        @click="selectProfile(profile)"
        class="cursor-pointer p-2 text-black min-w-24 dark:text-white hover:bg-dark dark:hover:bg-light hover:bg-opacity-30 dark:hover:bg-opacity-30"
      >
        {{ profile.description }}
      </li>
      <li
        @click="openCreateProfileModal"
        class="cursor-pointer flex flex-col min-w-24 p-2 text-black dark:text-white hover:bg-dark dark:hover:bg-light hover:bg-opacity-30 dark:hover:bg-opacity-30"
      >
        <AddIcon
          class="text-black dark:text-white justify-self-center self-center"
          @click="openCreateProfileModal"
        />
      </li>
    </ul>
  </div>

  <CreateProfileModal
    v-if="isCreateProfileModalOpen"
    :has-profile-created="true"
    :is-open="isCreateProfileModalOpen"
    @on-profile-create="onProfileCreate"
    @on-cancel="closeCreateProfileModal"
  />
</template>
