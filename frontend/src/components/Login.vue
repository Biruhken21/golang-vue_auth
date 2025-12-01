<template>
  <div class="min-h-screen flex items-center justify-center bg-gradient-to-br from-gray-900 via-gray-800 to-gray-900 text-white py-16 px-4">
    <div class="w-full max-w-md bg-gray-800/60 backdrop-blur rounded-2xl shadow-2xl border border-white/10 p-8">
      <h2 class="text-3xl font-extrabold text-center tracking-tight">Sign in to your account</h2>

      <form class="mt-10 space-y-6" @submit.prevent="login">
        <div class="space-y-4">
          <div>
            <label for="email-address" class="sr-only">Email address</label>
            <input
              id="email-address"
              name="email"
              type="email"
              autocomplete="email"
              required
              v-model="email"
              class="w-full px-4 py-3 rounded-lg bg-white/5 border border-white/10 placeholder-gray-400 text-white focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-transparent"
              placeholder="Email address"
            >
          </div>
          <div>
            <label for="password" class="sr-only">Password</label>
            <input
              id="password"
              name="password"
              type="password"
              autocomplete="current-password"
              required
              v-model="password"
              class="w-full px-4 py-3 rounded-lg bg-white/5 border border-white/10 placeholder-gray-400 text-white focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-transparent"
              placeholder="Password"
            >
          </div>
        </div>

        <button
          type="submit"
          class="w-full inline-flex justify-center items-center px-6 py-3 rounded-lg font-semibold bg-indigo-600 text-white shadow-md hover:bg-indigo-700 hover:shadow-lg transition-all duration-200 focus:outline-none focus:ring-4 focus:ring-indigo-500/40"
        >
          Sign in
        </button>
      </form>

      <p class="mt-6 text-center text-sm text-gray-300">
        Don't have an account?
        <router-link to="/register" class="text-indigo-400 hover:text-indigo-300 underline">Register</router-link>
      </p>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'

const email = ref('')
const password = ref('')
const router = useRouter()

const GQL = async (query, variables = {}) => {
  const token = localStorage.getItem('token')
  const headers = { 'Content-Type': 'application/json' }
  if (token) headers['Authorization'] = `Bearer ${token}`

  const res = await fetch('/query', {
    method: 'POST',
    headers,
    body: JSON.stringify({ query, variables }),
  })
  return res.json()
}

const login = async () => {
  const query = `
    mutation Login($input: LoginInput!) {
      login(input: $input) {
        token
        user { id email username createdAt updatedAt }
      }
    }
  `
  const variables = { input: { email: email.value, password: password.value } }

  try {
    const resp = await GQL(query, variables)
    if (resp.errors && resp.errors.length) {
      alert(resp.errors[0].message)
      return
    }
    const data = resp.data.login
    if (!data || !data.token) {
      alert('Login failed')
      return
    }
    localStorage.setItem('token', data.token)
    localStorage.setItem('user', JSON.stringify(data.user))
    router.push('/')
  } catch (err) {
    console.error(err)
    alert('Unexpected error')
  }
}
</script>
