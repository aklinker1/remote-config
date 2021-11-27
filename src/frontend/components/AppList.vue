<template>
  <create-app />
  <p v-if="error" class="text-red-300">{{ error }}</p>
  <ul class="space-y-4">
    <li v-for="app in apps" :key="app" class="list-item">
      <router-link :to="`/${app}`"
        ><p class="font-medium text-lg dark:text-white">
          {{ app }}
        </p></router-link
      >
      <div class="flex-1" />
      <img
        v-if="canDelete"
        class="
          w-10
          h-10
          p-2
          -m-2
          opacity-50
          hover:opacity-90
          dark:bg-white dark:bg-opacity-0 dark:hover:bg-opacity-20
          rounded-full
          text-white text-opacity-70 text-xs
          cursor-pointer
          transition-all
        "
        :src="TrashIcon"
        :alt="`Delete ${app}`"
        :title="`Delete ${app}`"
        @click="deleteApp(app)"
      />
    </li>
  </ul>
  <p v-if="isLoading" class="dark:text-white text-opacity-50">Loading...</p>
</template>

<script lang="ts" setup>
import { api } from "../api";
import createApp from "./CreateApp.vue";
import TrashIcon from "../assets/ic-trash.svg";
import { getAuthToken } from "../state/auth-token";

const isLoading = ref(false);
const apps = ref<string[]>([]);
const error = ref<string | undefined>();

const canDelete = !!getAuthToken();

onMounted(loadApps);

async function loadApps() {
  try {
    isLoading.value = true;
    apps.value = await api.getApps();
  } catch (err) {
    if (err instanceof Error) {
      error.value = err.message;
    } else {
      error.value = "Unknown error: " + JSON.stringify(err);
    }
  } finally {
    isLoading.value = false;
  }
}
async function deleteApp(app: string) {
  try {
    isLoading.value = true;
    await api.deleteConfig(app);
    await loadApps();
  } catch (err) {
    if (err instanceof Error) {
      error.value = err.message;
    } else {
      error.value = "Unknown error: " + JSON.stringify(err);
    }
  } finally {
    isLoading.value = false;
  }
}
</script>

<style>
.list-item {
  @apply bg-gray-700 shadow-md rounded p-4 flex items-center;
}
</style>
