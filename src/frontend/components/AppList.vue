<template>
  <p v-if="error">{{ error }}</p>
  <ul class="space-y-4">
    <li v-for="app in apps" :key="app" class="list-item">
      <router-link :to="`/${app}`"
        ><p class="font-medium text-lg dark:text-white">
          {{ app }}
        </p></router-link
      >
      <div class="flex-1" />
      <img
        class="
          w-8
          h-8
          p-1
          dark:bg-white dark:bg-opacity-0
          hover:bg-opacity-20
          transition-opacity
          rounded-full
          text-white text-opacity-70 text-xs
        "
        alt="Delete App"
      />
    </li>
  </ul>
  <p v-if="isLoading" class="dark:text-white text-opacity-50">Loading...</p>
</template>

<script lang="ts" setup>
import axios from "axios";

const isLoading = ref(false);
const apps = ref<string[]>([]);
const error = ref<string | undefined>();

onMounted(loadApps);

async function loadApps() {
  try {
    isLoading.value = true;
    const res = await axios.get<string[]>("/api/apps");
    apps.value = res.data;
    isLoading.value = false;
  } catch (err) {
    if (err instanceof Error) {
      error.value = err.message;
    } else {
      error.value = "Unknown error: " + JSON.stringify(err);
    }
  }
}
</script>

<style>
.list-item {
  @apply bg-gray-700 shadow-md rounded p-4 flex;
}
</style>
