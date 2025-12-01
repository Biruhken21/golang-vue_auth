# Frontend Guide (Vue 3 + Vite + Tailwind)

## 1) Scaffold the app
```bash
npm create vite@latest my-app -- --template vue
cd my-app
npm install
```

## 2) Add Tailwind CSS
```bash
npm install -D tailwindcss postcss autoprefixer
npx tailwindcss init -p
```

## 3) Configure Tailwind
Edit `tailwind.config.cjs`:
```js
module.exports = {
  content: ['./index.html','./src/**/*.{vue,js,ts,jsx,tsx}'],
  theme: { extend: {} },
  plugins: [],
}
```

## 4) Add Tailwind directives
Create or edit `src/index.css`:
```css
@tailwind base;
@tailwind components;
@tailwind utilities;
```

## 5) Wire styles in main
Edit `src/main.js`:
```js
import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import './index.css'

createApp(App).use(router).mount('#app')
```

## 6) Install and configure the router
```bash
npm i vue-router@4
```
Create `src/router/index.js`:
```js
import { createRouter, createWebHistory } from 'vue-router'
import Home from '../components/Home.vue'
import Login from '../components/Login.vue'
import Registration from '../components/Registration.vue'

export default createRouter({
  history: createWebHistory(),
  routes: [
    { path: '/', component: Home },
    { path: '/login', component: Login },
    { path: '/register', component: Registration },
  ],
})
```

## 7) Create core files
`src/App.vue`:
```vue
<template>
  <nav class="bg-gray-800 text-white px-4 py-3 shadow-md">
    <div class="max-w-7xl mx-auto grid grid-cols-3 items-center">
      <div class="flex items-center gap-2">
        <img src="/vite.svg" alt="Logo" class="h-8 w-8" />
        <span class="text-lg font-semibold">Awesome App</span>
      </div>
      <div class="flex justify-center">
        <router-link to="/" class="px-4 py-2 rounded-md hover:bg-gray-700">Home</router-link>
      </div>
      <div class="flex justify-end gap-3">
        <router-link to="/login" class="px-4 py-2 rounded-md hover:bg-gray-700">Login</router-link>
        <router-link to="/register" class="px-4 py-2 rounded-md hover:bg-gray-700">Register</router-link>
      </div>
    </div>
  </nav>
  <router-view />
</template>
```

`src/components/Home.vue`:
```vue
<template>
  <section class="min-h-[70vh] flex items-center justify-center bg-gradient-to-br from-gray-900 to-gray-800 text-white px-4 py-16">
    <div class="max-w-4xl w-full text-center">
      <h1 class="text-4xl sm:text-5xl font-extrabold mb-4">Welcome to Our Awesome App!</h1>
      <p class="text-lg sm:text-xl text-gray-300 mb-8">Discover features and streamline your daily tasks.</p>
      <div class="flex flex-col sm:flex-row items-center justify-center gap-4">
        <router-link to="/register" class="px-6 py-3 rounded-lg bg-indigo-600 hover:bg-indigo-700 text-white font-semibold">Get Started</router-link>
        <router-link to="/login" class="px-6 py-3 rounded-lg border border-indigo-400 text-indigo-300 hover:bg-indigo-600 hover:text-white">I have an account</router-link>
      </div>
    </div>
  </section>
</template>
```

`src/components/Login.vue`:
```vue
<template>
  <div class="min-h-[70vh] flex items-center justify-center bg-gradient-to-br from-gray-900 to-gray-800 px-4 py-16">
    <form class="w-full max-w-md bg-gray-800/60 p-8 rounded-2xl shadow-2xl space-y-4">
      <h2 class="text-white text-2xl font-bold text-center">Sign in</h2>
      <input class="w-full px-4 py-3 rounded-lg bg-white/5 border border-white/10 text-white" type="email" placeholder="Email" />
      <input class="w-full px-4 py-3 rounded-lg bg-white/5 border border-white/10 text-white" type="password" placeholder="Password" />
      <button class="w-full px-6 py-3 rounded-lg bg-indigo-600 hover:bg-indigo-700 text-white font-semibold">Login</button>
    </form>
  </div>
</template>
```

`src/components/Registration.vue`:
```vue
<template>
  <div class="min-h-[70vh] flex items-center justify-center bg-gradient-to-br from-gray-900 to-gray-800 px-4 py-16">
    <form class="w-full max-w-md bg-gray-800/60 p-8 rounded-2xl shadow-2xl space-y-4">
      <h2 class="text-white text-2xl font-bold text-center">Register</h2>
      <input class="w-full px-4 py-3 rounded-lg bg-white/5 border border-white/10 text-white" type="email" placeholder="Email" />
      <input class="w-full px-4 py-3 rounded-lg bg-white/5 border border-white/10 text-white" type="text" placeholder="Username" />
      <input class="w-full px-4 py-3 rounded-lg bg-white/5 border border-white/10 text-white" type="password" placeholder="Password" />
      <input class="w-full px-4 py-3 rounded-lg bg-white/5 border border-white/10 text-white" type="password" placeholder="Confirm Password" />
      <button class="w-full px-6 py-3 rounded-lg bg-indigo-600 hover:bg-indigo-700 text-white font-semibold">Create account</button>
    </form>
  </div>
</template>
```

## 8) Run the app
```bash
npm run dev
```

## 9) Build & preview
```bash
npm run build
npm run preview
```

## Notes
- Use Tailwind classes directly in templates.
- Add new pages by creating a `.vue` file under `src/components/` (or `src/pages/`) and registering a route in `src/router/index.js`.
- Replace `/vite.svg` with your own logo in `public/` when ready.
