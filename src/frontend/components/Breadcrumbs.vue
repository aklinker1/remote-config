<template>
  <div class="flex space-x-2 items-center">
    <template v-for="(p, i) in path" :key="p">
      <h4 v-if="i > 0" class="path divider">▸</h4>
      <h4 v-if="i === path.length - 1" class="path active">{{ p }}</h4>
      <router-link v-else :to="`/${path.slice(0, i).join('/')}`"
        ><h4 class="path hover:underline">{{ p }}</h4></router-link
      >
    </template>
  </div>
</template>

<script lang="ts" setup>
import { useRoute } from "vue-router";

const route = useRoute();
const pathStr = ref(route.path);
const path = ref(getPath());
watch(
  () => route.path,
  () => {
    const newPath = route.path;
    pathStr.value = newPath;
    path.value = getPath();
  }
);

function getPath(): string[] {
  const items = route.path
    .replace(/^\//, "Remote Config/")
    .replace(/\/$/g, "")
    .split("/")
    .map((part) => decodeURIComponent(part));
  return items;
}
</script>

<style>
.path {
  @apply font-semibold text-2xl dark:text-blue-300;
}

.path.active {
  @apply dark:text-white text-opacity-90;
}
.path.divider {
  @apply text-lg;
}
</style>
