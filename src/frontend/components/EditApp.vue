<template>
  <p v-if="error" class="text-red-300">{{ error }}</p>
  <editor
    v-if="data"
    :data="data"
    :loading="isSaving"
    loading-text="Saving..."
    @save="saveNewConfig"
    @discard="$router.back()"
  />
</template>

<script lang="ts" setup>
import { useRoute } from 'vue-router';
import { api } from '../api';
import Editor from './Editor.vue';

const data = ref<object>();
const error = ref<string>();

const route = useRoute();
const appName = route.params.app as string;

onMounted(loadConfig);

async function loadConfig() {
  try {
    error.value = undefined;
    data.value = await api.getConfig(appName);
  } catch (err) {
    error.value = err instanceof Error ? err.message : 'Unknown error: ' + JSON.stringify(err);
  }
}

const isSaving = ref(false);
async function saveNewConfig(config: object) {
  const animation = new Promise(res => setTimeout(res, 400));
  try {
    isSaving.value = true;
    error.value = undefined;
    data.value = await api.updateConfig(appName, config);
  } catch (err) {
    if (err instanceof Error) {
      error.value = err.message;
    } else {
      error.value = 'Unknown error: ' + JSON.stringify(err);
    }
  } finally {
    isSaving.value = false;
  }
  await animation;
}
</script>
