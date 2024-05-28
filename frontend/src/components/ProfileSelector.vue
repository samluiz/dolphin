<script setup lang="ts">
import { ref, defineEmits, computed, toRefs } from "vue";
import { useProfileStore } from "@/stores/ProfileStore";
import { types } from "../../wailsjs/go/models";
import CreateProfileModal from "./CreateProfileModal.vue";
import SelectArrowIcon from "./ui/icons/SelectArrowIcon.vue";
import { storeToRefs } from "pinia";
import AddIcon from "./ui/icons/AddIcon.vue";

const emit = defineEmits(["on-profile-create", "on-profile-change"]);

const profileStore = useProfileStore();

const { profiles } = storeToRefs(profileStore);

const props = defineProps({
  profile: Object as () => types.Profile,
});

const { profile } = toRefs(props);

const isCreateProfileModalOpen = ref(false);
const isDropdownOpen = ref(false);

const selectProfile = (newProfile: types.Profile) => {
  isDropdownOpen.value = false;
  profileStore.setActiveProfile(newProfile);
  profileStore.fetchProfiles();
  emit("on-profile-change");
};

const openCreateProfileModal = () => {
  isCreateProfileModalOpen.value = true;
};

const closeCreateProfileModal = () => {
  isCreateProfileModalOpen.value = false;
  emit("on-profile-change");
};

const onProfileCreate = () => {
  emit("on-profile-create", profile?.value);
  emit("on-profile-change");
};

const openDropdown = () => {
  isDropdownOpen.value = true;
};

const closeDropdown = () => {
  setTimeout(() => (isDropdownOpen.value = false), 1000);
};

const otherProfiles = computed(() => {
  return (
    profiles.value &&
    profiles.value.filter(
      (profile: types.Profile) => profile.id !== profileStore.profile?.id,
    )
  );
});
</script>

<template>
  <div @mouseover="openDropdown" @mouseleave="closeDropdown">
    <button
      class="py-1.5 px-2 duration-100 grid min-w-24 place-items-center grid-flow-col rounded-sm text-black dark:text-light-text hover:bg-dark dark:hover:bg-light hover:bg-opacity-30 dark:hover:bg-opacity-30"
    >
      {{ profile ? profile.description : $t("select_profile") }}
      <SelectArrowIcon />
    </button>
    <transition name="fade">
      <ul
        v-if="isDropdownOpen"
        class="shadow-lg flex flex-col rounded-sm ring-1 ring-black ring-opacity-5 focus:outline-none absolute min-w-24 bg-light dark:bg-dark text-black dark:text-light-text z-50"
      >
        <li
          v-for="profile in otherProfiles"
          :key="profile.id"
          @click="selectProfile(profile)"
          class="cursor-pointer p-2 text-black min-w-24 dark:text-light-text hover:bg-dark dark:hover:bg-light hover:bg-opacity-30 dark:hover:bg-opacity-30"
        >
          {{ profile.description }}
        </li>
        <li
          @click="openCreateProfileModal"
          class="cursor-pointer flex flex-col min-w-24 p-2 text-black dark:text-light-text hover:bg-dark dark:hover:bg-light hover:bg-opacity-30 dark:hover:bg-opacity-30"
        >
          <AddIcon
            class="text-black dark:text-light-text justify-self-center self-center"
            @click="openCreateProfileModal"
          />
        </li>
      </ul>
    </transition>
  </div>

  <CreateProfileModal
    v-if="isCreateProfileModalOpen"
    :has-profile-created="true"
    :is-open="isCreateProfileModalOpen"
    @on-profile-create="onProfileCreate"
    @on-cancel="closeCreateProfileModal"
  />
</template>

<style scoped>
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.5s;
}
.fade-enter,
.fade-leave-to {
  opacity: 0;
}
</style>
