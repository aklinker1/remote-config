<template>
  <h1 class="font-semibold text-2xl dark:text-white text-opacity-90">Login</h1>
  <form @submit.prevent="login()" class="flex space-x-4">
    <input
      class="flex-1"
      v-model="authToken"
      placeholder="Auth token"
      type="password"
      autocomplete="current-password"
    />
    <input type="submit" value="Login" class="primary" />
  </form>
</template>

<script lang="ts" setup>
import { useRoute, useRouter } from 'vue-router';
import { getAuthToken, setAuthToken } from '../state/auth-token';

const authToken = ref(getAuthToken());

const router = useRouter();
const route = useRoute();

function login() {
  setAuthToken(authToken.value);
  router.push((route.params.redirect as string | undefined) || '/');
}
</script>
