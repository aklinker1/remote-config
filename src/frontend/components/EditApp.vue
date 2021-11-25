<template>
  <p v-if="error">{{ error }}</p>
  <editor
    v-if="data"
    :data="data"
    :loading="isSaving"
    loading-text="Saving..."
    @save="saveNewConfig"
  />
</template>

<script lang="ts" setup>
import axios from "axios";
import { useRoute } from "vue-router";
import Editor from "./Editor.vue";

const data = ref<object>();
const error = ref<string>();

const route = useRoute();
const appName = route.params.app;

onMounted(loadConfig);

async function loadConfig() {
  try {
    error.value = undefined;
    const res = await axios.get(`/api/config/${appName}`);
    data.value = res.data;
  } catch (err) {
    error.value =
      err instanceof Error
        ? err.message
        : "Unknown error: " + JSON.stringify(err);
  }
}

const isSaving = ref(false);
async function saveNewConfig(config: object) {
  const animation = new Promise((res) => setTimeout(res, 400));
  try {
    isSaving.value = true;
    const res = await axios.put(`/api/config/${appName}`, config);
    data.value = res.data;
  } finally {
    isSaving.value = false;
  }
  await animation;
}
</script>
