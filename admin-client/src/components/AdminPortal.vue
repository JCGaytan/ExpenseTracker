<template>
  <div class="flex min-h-screen bg-gray-100">
    <!-- Sidebar -->
    <aside :class="['transition-all duration-200 bg-white shadow h-screen', collapsed ? 'w-16' : 'w-56']">
      <div class="flex flex-col h-full">
        <button @click="collapsed = !collapsed" class="p-2 focus:outline-none">
          <svg v-if="collapsed" class="w-6 h-6 mx-auto" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" d="M4 6h16M4 12h16M4 18h16"/></svg>
          <svg v-else class="w-6 h-6 mx-auto" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" d="M6 18L18 6M6 6l12 12"/></svg>
        </button>
        <nav class="flex-1 mt-4">
          <ul class="space-y-2">
            <li v-for="item in navItems" :key="item.name">
              <button @click="current = item.name" :class="['flex items-center w-full px-4 py-2 rounded hover:bg-blue-100', current === item.name ? 'bg-blue-100 text-blue-700' : 'text-gray-700']">
                <span class="w-6 h-6 flex items-center justify-center">
                  <component :is="item.icon" />
                </span>
                <span v-if="!collapsed" class="ml-3">{{ item.label }}</span>
              </button>
            </li>
          </ul>
        </nav>
        <!-- Logoff Button at the bottom -->
        <div class="mt-auto mb-4 flex flex-col items-center">
          <button @click="handleLogoff" class="flex items-center w-full px-4 py-2 rounded text-red-600 hover:bg-red-100">
            <span class="w-6 h-6 flex items-center justify-center">
              <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-6 h-6">
                <path stroke-linecap="round" stroke-linejoin="round" d="M15.75 9V5.25A2.25 2.25 0 0013.5 3h-6A2.25 2.25 0 005.25 5.25v13.5A2.25 2.25 0 007.5 21h6a2.25 2.25 0 002.25-2.25V15" />
                <path stroke-linecap="round" stroke-linejoin="round" d="M18 12H9m0 0l3-3m-3 3l3 3" />
              </svg>
            </span>
            <span v-if="!collapsed" class="ml-3">Logoff</span>
          </button>
        </div>
      </div>
    </aside>
    <!-- Main Content -->
    <main class="flex-1 flex flex-col items-center justify-start p-6 max-w-5xl mx-auto w-full">
      <component :is="currentComponent" @refresh="refresh" />
    </main>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'
import { UserIcon, CreditCardIcon, TagIcon, ChartBarIcon } from '@heroicons/vue/24/outline'
import TransactionsCrud from './TransactionsCrud.vue'
import CategoriesCrud from './CategoriesCrud.vue'
import ReportsView from './ReportsView.vue'

const collapsed = ref(false)
const current = ref('transactions')

function handleLogoff() {
  localStorage.removeItem('token');
  window.location.reload();
}

const navItems = [
  { name: 'transactions', label: 'Transactions', icon: CreditCardIcon },
  { name: 'categories', label: 'Categories', icon: TagIcon },
  { name: 'reports', label: 'Reports', icon: ChartBarIcon },
]

const currentComponent = computed(() => {
  switch (current.value) {
    case 'transactions': return TransactionsCrud
    case 'categories': return CategoriesCrud
    case 'reports': return ReportsView
    default: return TransactionsCrud
  }
})

function refresh() {
  // Placeholder for refresh event from children
}
</script>

<style scoped>
/* Responsive sidebar */
@media (max-width: 768px) {
  aside {
    position: fixed;
    z-index: 40;
    height: 100vh;
    left: 0;
    top: 0;
  }
  main {
    margin-left: 0 !important;
  }
}

/* Adjust main content layout */
main {
  flex-grow: 1;
  display: flex;
  flex-direction: column;
  justify-content: flex-start;
  align-items: flex-start;
  padding: 1rem;
}

/* Reduce empty space */
.component-container {
  width: calc(100% - 4rem); /* Adjust width to reduce space */
  max-width: 1200px;
  margin: 0 auto;
}

/* Align content closer to sidebar */
.component-container {
  margin-left: 1rem;
}
</style>
