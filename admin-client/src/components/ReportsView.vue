<template>
  <div class="w-full max-w-3xl mx-auto bg-white rounded shadow p-6 mt-6">
    <h2 class="text-xl font-bold mb-4">Monthly Summary</h2>
    <div v-if="summary" class="grid grid-cols-1 md:grid-cols-2 gap-4">
      <div class="bg-blue-50 p-4 rounded">
        <div class="text-gray-600">Total Income</div>
        <div class="text-2xl font-bold text-green-600">{{ summary.total_income }}</div>
      </div>
      <div class="bg-red-50 p-4 rounded">
        <div class="text-gray-600">Total Expense</div>
        <div class="text-2xl font-bold text-red-600">{{ summary.total_expense }}</div>
      </div>
      <div class="bg-gray-50 p-4 rounded col-span-2">
        <div class="text-gray-600">Balance</div>
        <div :class="{'text-green-600': balance >= 0, 'text-red-600': balance < 0}" class="text-2xl font-bold">{{ balance }}</div>
      </div>
      <div class="col-span-2">
        <h3 class="font-semibold mt-4 mb-2">By Category</h3>
        <table class="min-w-full text-sm">
          <thead>
            <tr class="bg-gray-100">
              <th class="px-4 py-2">Category</th>
              <th class="px-4 py-2">Amount</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="(amount, cat) in summary.by_category" :key="cat" :class="{'bg-blue-50': amount > 0, 'bg-red-50': amount < 0}" class="border-b">
              <td class="px-4 py-2">{{ cat }}</td>
              <td class="px-4 py-2 flex justify-between items-center">
                <span>{{ amount }}</span>
                <span :class="{'text-green-600': amount > 0, 'text-red-600': amount < 0}" class="text-xl font-bold">
                  {{ amount > 0 ? '+' : '-' }}
                </span>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
    <div v-else class="text-gray-500">Loading...</div>
  </div>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue'

const summary = ref(null)
const balance = computed(() => summary.value ? summary.value.total_income + summary.value.total_expense : 0) // Correct calculation

function getToken() {
  return localStorage.getItem('token')
}

async function fetchSummary() {
  const res = await fetch('http://localhost:8080/reports/summary', {
    headers: { Authorization: 'Bearer ' + getToken() }
  })
  summary.value = await res.json()
}

onMounted(() => {
  fetchSummary()
})
</script>
