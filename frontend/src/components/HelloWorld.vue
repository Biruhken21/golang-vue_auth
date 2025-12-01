<script setup>
import { ref } from 'vue'

defineProps({
  msg: String,
})

const count = ref(0)
</script>

<template>
  <div class="container mx-auto p-4 text-center">
    <h1 class="text-5xl font-extrabold mb-6 text-transparent bg-clip-text bg-gradient-to-r from-indigo-500 to-purple-500">{{ msg }}</h1>

    <div class="card bg-vt-c-bg-soft p-8 rounded-2xl shadow-xl hover:shadow-2xl transition-all duration-300 transform hover:-translate-y-1 mb-10 border border-vt-c-divider-dark-1 dark:border-vt-c-divider-light-1">
      <button type="button" @click="count++" class="px-8 py-4 bg-indigo-600 text-white font-bold rounded-full shadow-lg hover:bg-indigo-700 focus:outline-none focus:ring-4 focus:ring-indigo-500 focus:ring-opacity-50 transition-all duration-300 transform hover:scale-105">
        count is {{ count }}
      </button>
      <p class="mt-6 text-lg text-vt-c-text-2">
        Edit
        <code class="bg-vt-c-bg-alt px-3 py-1 rounded-md text-vt-c-text-1 font-mono">components/HelloWorld.vue</code> to test HMR
      </p>
    </div>

    <p class="mb-4 text-xl text-vt-c-text-1">
      Check out
      <a href="https://vuejs.org/guide/quick-start.html#local" target="_blank"
        class="text-indigo-400 hover:text-indigo-300 transition-colors duration-200 underline"
        >create-vue</a
      >, the official Vue + Vite starter
    </p>
    <p class="mb-4 text-xl text-vt-c-text-1">
      Learn more about IDE Support for Vue in the
      <a
        href="https://vuejs.org/guide/scaling-up/tooling.html#ide-support"
        target="_blank"
        class="text-indigo-400 hover:text-indigo-300 transition-colors duration-200 underline"
        >Vue Docs Scaling up Guide</a
      >.
    </p>
    <p class="read-the-docs text-vt-c-text-3 text-base mt-8">
      Click on the Vite and Vue logos to learn more
    </p>
  </div>
</template>

<style scoped>
/* Removed custom styles, relying on Tailwind CSS */
</style>
