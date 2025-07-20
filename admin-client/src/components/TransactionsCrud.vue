<template>
  <div class="w-full max-w-3xl mx-auto bg-white rounded shadow p-6 mt-6">
    <div class="flex justify-between items-center mb-4">
      <h2 class="text-xl font-bold">Transactions</h2>
      <button @click="showModal = true" class="bg-blue-600 text-white px-4 py-2 rounded hover:bg-blue-700">Add</button>
    </div>
    <div class="overflow-x-auto">
      <table class="min-w-full text-sm">
        <thead>
          <tr class="bg-gray-100">
            <th class="px-4 py-2">Amount</th>
            <th class="px-4 py-2">Date</th>
            <th class="px-4 py-2">Category</th>
            <th class="px-4 py-2">Description</th>
            <th class="px-4 py-2">Actions</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="tx in transactions" :key="tx.id" class="border-b">
            <td class="px-4 py-2">{{ tx.amount }}</td>
            <td class="px-4 py-2">{{ tx.date.split('T')[0] }}</td>
            <td class="px-4 py-2">{{ getCategoryName(tx.category_id) }}</td>
            <td class="px-4 py-2">{{ tx.description }}</td>
            <td class="px-4 py-2 flex gap-2">
              <button @click="edit(tx)" class="bg-blue-500 text-white px-3 py-1 rounded hover:bg-blue-600">Edit</button>
              <button @click="remove(tx.id)" class="bg-red-500 text-white px-3 py-1 rounded hover:bg-red-600">Delete</button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
    <!-- Modal -->
    <div v-if="showModal" class="fixed inset-0 flex items-center justify-center bg-black bg-opacity-40 z-50">
      <div class="bg-white rounded shadow p-6 w-full max-w-md">
        <h3 class="text-lg font-bold mb-4">{{ editId ? 'Edit' : 'Add' }} Transaction</h3>
        <form @submit.prevent="submit">
          <input v-model.number="form.amount" type="number" step="0.01" placeholder="Amount" class="w-full mb-2 px-3 py-2 border rounded" required />
          <input v-model="form.date" type="date" placeholder="Date" class="w-full mb-2 px-3 py-2 border rounded" required />
          <select v-model.number="form.category_id" class="w-full mb-2 px-3 py-2 border rounded" required>
            <option value="" disabled>Select Category</option>
            <option v-for="cat in categories" :key="cat.id" :value="cat.id">{{ cat.name }}</option>
          </select>
          <input v-model="form.description" type="text" placeholder="Description" class="w-full mb-2 px-3 py-2 border rounded" />
          <div class="flex justify-end gap-2 mt-4">
            <button type="button" @click="close" class="px-4 py-2 rounded bg-gray-200">Cancel</button>
            <button type="submit" class="px-4 py-2 rounded bg-blue-600 text-white">Save</button>
          </div>
        </form>
      </div>
    </div>
    <!-- Delete Confirmation Modal -->
    <div v-if="showDeleteModal" class="fixed inset-0 flex items-center justify-center bg-black bg-opacity-40 z-50">
      <div class="bg-white rounded shadow p-6 w-full max-w-md">
        <h3 class="text-lg font-bold mb-4">Confirm Delete</h3>
        <p>Are you sure you want to delete this transaction?</p>
        <div class="flex justify-end gap-2 mt-4">
          <button type="button" @click="showDeleteModal = false" class="px-4 py-2 rounded bg-gray-200">Cancel</button>
          <button type="button" @click="confirmDelete" class="px-4 py-2 rounded bg-red-600 text-white">Delete</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'

const transactions = ref([])
const categories = ref([])
const showModal = ref(false)
const editId = ref(null)
const form = ref({ amount: '', date: '', category_id: '', description: '' })
const showDeleteModal = ref(false)
let deleteId = null

function getToken() {
  return localStorage.getItem('token')
}

function getCategoryName(id) {
  const cat = categories.value.find(c => c.id === id)
  return cat ? cat.name : ''
}

async function fetchTransactions() {
  const res = await fetch('http://localhost:8080/transactions', {
    headers: { Authorization: 'Bearer ' + getToken() }
  })
  transactions.value = await res.json()
}

async function fetchCategories() {
  const res = await fetch('http://localhost:8080/categories', {
    headers: { Authorization: 'Bearer ' + getToken() }
  })
  categories.value = await res.json()
}

function edit(tx) {
  editId.value = tx.id
  form.value = { ...tx, date: tx.date.split('T')[0] } // Ensure date is formatted correctly
  showModal.value = true
}

function close() {
  showModal.value = false
  editId.value = null
  form.value = { amount: '', date: '', category_id: '', description: '' }
}

async function submit() {
  const method = editId.value ? 'PUT' : 'POST'
  const url = editId.value ? `http://localhost:8080/transactions/${editId.value}` : 'http://localhost:8080/transactions'
  await fetch(url, {
    method,
    headers: {
      'Content-Type': 'application/json',
      Authorization: 'Bearer ' + getToken()
    },
    body: JSON.stringify(form.value)
  })
  await fetchTransactions()
  close()
}

function remove(id) {
  deleteId = id
  showDeleteModal.value = true
}

async function confirmDelete() {
  if (deleteId) {
    await fetch(`http://localhost:8080/transactions/${deleteId}`, {
      method: 'DELETE',
      headers: {
        Authorization: 'Bearer ' + getToken()
      }
    })
    transactions.value = transactions.value.filter(transaction => transaction.id !== deleteId)
    deleteId = null
    showDeleteModal.value = false
  }
}

onMounted(() => {
  fetchCategories()
  fetchTransactions()
})
</script>
