<template>
  <div class="min-h-screen flex items-center justify-center bg-gray-100">
    <form
      class="bg-white p-8 rounded shadow-md w-full max-w-sm space-y-6"
      @submit.prevent="handleLogin"
    >
      <h2 class="text-2xl font-bold text-center text-gray-800">Sign In</h2>
      <input
        v-model="email"
        type="email"
        placeholder="Email"
        class="w-full px-4 py-2 border rounded focus:outline-none focus:ring-2 focus:ring-blue-400"
        required
        autocomplete="username"
      />
      <input
        v-model="password"
        type="password"
        placeholder="Password"
        class="w-full px-4 py-2 border rounded focus:outline-none focus:ring-2 focus:ring-blue-400"
        required
        autocomplete="current-password"
      />
      <button
        type="submit"
        class="w-full bg-blue-600 hover:bg-blue-700 text-white font-semibold py-2 rounded transition-colors"
        :disabled="loading"
      >
        <span v-if="loading">Signing in...</span>
        <span v-else>Login</span>
      </button>
      <p v-if="error" class="text-red-500 text-center">{{ error }}</p>
    </form>
  </div>
</template>

<script setup>
import { ref } from 'vue'

const email = ref('')
const password = ref('')
const loading = ref(false)
const error = ref('')

async function handleLogin() {
  error.value = ''
  loading.value = true
  try {
    const res = await fetch('http://localhost:8080/auth/login', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ email: email.value, password: password.value })
    })
    if (!res.ok) {
      const data = await res.json().catch(() => ({}))
      throw new Error(data.error || 'Invalid email or password')
    }
    const data = await res.json()
    // Store JWT token (localStorage or cookie)
    localStorage.setItem('token', data.token)
    // Redirect or emit event (customize as needed)
    window.location.reload()
  } catch (e) {
    error.value = e.message || 'Login failed'
  } finally {
    loading.value = false
  }
}
</script>
