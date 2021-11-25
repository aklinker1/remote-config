<template>
  <div class="max-w-lg mx-auto space-y-4 pt-16">
    <breadcrumbs />
    <router-view />
  </div>
</template>

<script lang="ts" setup>
import Breadcrumbs from "./Breadcrumbs.vue";
</script>
