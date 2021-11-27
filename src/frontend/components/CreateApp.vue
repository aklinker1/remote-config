<template>
  <div class="space-y-2">
    <form class="flex space-x-4" @submit.prevent="createApp()">
      <input
        type="text"
        placeholder="Create new..."
        v-model="appName"
        class="flex-1"
      />
      <input
        type="submit"
        value="Create"
        class="primary"
        :disabled="!appName"
      />
    </form>
    <p v-if="error" class="text-red-300">{{ error }}</p>
  </div>
</template>

<script lang="ts" setup>
import { useRouter } from "vue-router";
import { api } from "../api";

const appName = ref("");
const error = ref<string | undefined>();
const isSaving = ref(false);
const router = useRouter();

async function createApp() {
  if (!appName) return;

  try {
    isSaving.value = true;
    await api.createNewConfig(appName.value);
    router.push(`/${appName.value}`);
  } catch (err) {
    if (err instanceof Error) {
      error.value = err.message;
    } else {
      error.value = "Unknown error: " + JSON.stringify(err);
    }
  } finally {
    isSaving.value = false;
  }
}
</script>
