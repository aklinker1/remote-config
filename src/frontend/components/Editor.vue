<template>
  <div class="space-y-4 shadow-lg p-6 rounded bg-gray-900 relative">
    <div class="flex space-x-2">
      <button class="primary" @click="save()">Save</button>
      <button class="secondary" @click="discard()">Discard</button>
      <div class="flex-1" />
      <button class="primary" @click="format()">Format</button>
    </div>
    <prism-editor
      class="h-96 font-mono dark:text-white -mx-3"
      v-model="code"
      :highlight="highlighter"
      line-numbers
      :insert-spaces="true"
      :readonly="loading"
    />
    <p v-if="error" class="text-red-300">{{ error }}</p>

    <div
      class="
        absolute
        inset-0
        bg-gray-900 bg-opacity-90
        flex
        items-center
        justify-center
        transition-opacity
        duration-200
      "
      :class="{
        'opacity-100 pointer-events-auto cursor-auto': loading,
        'opacity-0 pointer-events-none cursor-not-allowed': !loading,
      }"
    >
      <p class="text-white">{{ loadingText }}</p>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { PrismEditor } from "vue-prism-editor";
import Prism from "prismjs";
import "vue-prism-editor/dist/prismeditor.min.css"; // Editor styles
import "prism-themes/themes/prism-atom-dark.css"; // syntax highlighting styles

const props = defineProps<{
  data: object;
  loading: boolean;
  loadingText: string;
}>();

const emits = defineEmits<{
  (event: "save", config: object): void;
}>();

const error = ref<string>();
function stringify(obj: object, fallback = ""): string {
  try {
    error.value = undefined;
    return JSON.stringify(obj, null, 2);
  } catch (err) {
    if (err instanceof Error) {
      error.value = err.message;
    } else {
      error.value = "Unknown error: " + JSON.stringify(err);
    }
    return fallback;
  }
}
function parse(str: string): object {
  try {
    error.value = undefined;
    return JSON.parse(str);
  } catch (err) {
    if (err instanceof Error) {
      error.value = err.message;
    } else {
      error.value = "Unknown error: " + JSON.stringify(err);
    }
    throw err;
  }
}

const code = ref(stringify(props.data));

const highlighter = (code: string) => {
  return Prism.highlight(code, Prism.languages.javascript, "javascript");
};

function save() {
  emits("save", parse(code.value));
}

function discard() {
  code.value = stringify(props.data);
}

function format() {
  code.value = stringify(parse(code.value), code.value);
}
</script>

<style>
button {
  @apply rounded px-3 py-1.5;
}
button.primary {
  @apply bg-blue-300 rounded;
}
button.secondary {
  @apply text-white text-opacity-70 hover:bg-gray-700 ring-1 ring-gray-700 transition-all hover:text-opacity-100;
}
</style>
