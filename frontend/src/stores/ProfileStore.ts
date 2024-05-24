import { defineStore } from "pinia";
import { Ref, ref } from "vue";
import { types } from "../../wailsjs/go/models";
import {
  FindAll as findAllProfiles,
  Create as createProfile,
  Update as updateProfile,
  Delete as deleteProfile,
  FindByID as findProfileByID,
} from "../../wailsjs/go/profile/service";

export const useProfileStore = defineStore(
  "profiles",
  () => {
    const profile: Ref<types.Profile | null> = ref(null);
    const profiles: Ref<types.Profile[]> = ref([]);

    const fetchProfiles = async () => {
      console.log("Fetching profiles");
      const response = await findAllProfiles();
      profiles.value = response;

      if (profiles.value && profiles.value.length > 0) {
        if (profile.value) {
          console.log("Profile already created");
          return;
        }
        profiles.value.map((p) => {
          if (p.is_default) {
            setActiveProfile(p);
            profile.value = p;
          }
        });
      } else {
        setActiveProfile(null);
      }
    };

    const isOnlyProfile = () =>
      profiles && profiles.value && profiles.value.length === 1;
    const hasProfileCreated = () =>
      profiles && profiles.value && profiles.value.length > 0;

    const setActiveProfile = (newProfile: types.Profile | null) => {
      console.log("Active profile set to", newProfile);
      profile.value = newProfile;
    };

    const create = async (profile: types.ProfileInput) => {
      const response = await createProfile(profile)
        .then(async (data) => {
          setActiveProfile(data);
          await fetchProfiles();
        })
        .catch((e: any) => {
          console.error(e);
        });
      return response;
    };

    const update = async (id: number, profile: types.ProfileInput) => {
      const response = await updateProfile(id, profile)
        .then(async (data) => {
          setActiveProfile(data);
          await fetchProfiles();
        })
        .catch((e: any) => {
          console.error(e);
        });
      return response;
    };

    const remove = async (id: number) => {
      const response = await deleteProfile(id)
        .then(async () => {
          await fetchProfiles();
          setActiveProfile(null);
        })
        .catch((e: any) => {
          console.error(e);
        });
      return response;
    };

    const findProfile = async (id: number) => {
      const response = await findProfileByID(id);
      profile.value = response;
      return response;
    };

    return {
      profile,
      profiles,
      isOnlyProfile,
      hasProfileCreated,
      setActiveProfile,
      fetchProfiles,
      create,
      update,
      remove,
      findProfile,
    };
  },
  {
    persist: true,
  },
);
