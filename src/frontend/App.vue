<template>
  <router-view />
</template>

<script lang="ts" setup>
import { provideAuthToken } from "./state/auth-token";

provideAuthToken();
</script>
