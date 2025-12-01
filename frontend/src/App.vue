<script setup>
// No script needed for basic App.vue with router-view
</script>

<template>
  <nav class="bg-gray-800 text-white px-4 py-3 shadow-md">
    <div class="max-w-7xl mx-auto grid grid-cols-3 items-center">
      <div class="flex items-center gap-2 justify-start">
        <img src="/vite.svg" alt="Logo" class="h-8 w-8" />
        <span class="text-lg font-semibold">Awesome App</span>
      </div>
      <div class="flex items-center justify-center">
        <router-link to="/" class="px-4 py-2 rounded-md text-gray-200 hover:bg-gray-700 hover:text-white transition-colors duration-200">Home</router-link>
      </div>
      <div class="flex items-center gap-4 justify-end">
        <router-link to="/login" class="px-4 py-2 rounded-md text-gray-300 hover:bg-gray-700 hover:text-white transition-colors duration-200">Login</router-link>
        <router-link to="/register" class="px-4 py-2 rounded-md text-gray-300 hover:bg-gray-700 hover:text-white transition-colors duration-200">Register</router-link>
      </div>
    </div>
  </nav>
  <router-view />
</template>
