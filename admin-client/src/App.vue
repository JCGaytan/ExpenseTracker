
<script setup>
import { computed } from 'vue'
import Login from './components/Login.vue'
import AdminPortal from './components/AdminPortal.vue'
const isLoggedIn = computed(() => !!localStorage.getItem('token'))
</script>

<template>
  <component :is="isLoggedIn ? AdminPortal : Login" />
</template>

