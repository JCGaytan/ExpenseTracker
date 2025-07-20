<template>
  <div class="w-full max-w-3xl mx-auto bg-white rounded shadow p-6 mt-6">
    <div class="flex justify-between items-center mb-4">
      <h2 class="text-xl font-bold">Categories</h2>
      <button @click="showModal = true" class="bg-blue-600 text-white px-4 py-2 rounded hover:bg-blue-700">Add</button>
    </div>
    <div class="overflow-x-auto">
      <table class="min-w-full text-sm">
        <thead>
          <tr class="bg-gray-100">
            <th class="px-4 py-2">Name</th>
            <th class="px-4 py-2">Actions</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="cat in categories" :key="cat.id" class="border-b">
            <td class="px-4 py-2">{{ cat.name }}</td>
            <td class="px-4 py-2 flex gap-2">
              <button @click="edit(cat)" class="bg-blue-500 text-white px-3 py-1 rounded hover:bg-blue-600">Edit</button>
              <button @click="remove(cat.id)" class="bg-red-500 text-white px-3 py-1 rounded hover:bg-red-600">Delete</button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
    <!-- Modal -->
    <div v-if="showModal" class="fixed inset-0 flex items-center justify-center bg-black bg-opacity-40 z-50">
      <div class="bg-white rounded shadow p-6 w-full max-w-md">
        <h3 class="text-lg font-bold mb-4">{{ editId ? 'Edit' : 'Add' }} Category</h3>
        <form @submit.prevent="submit">
          <input v-model="form.name" type="text" placeholder="Name" class="w-full mb-2 px-3 py-2 border rounded" required />
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
        <p>Are you sure you want to delete this category?</p>
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

const categories = ref([])
const showModal = ref(false)
const editId = ref(null)
const form = ref({ name: '' })
const showDeleteModal = ref(false)
let deleteId = null

function getToken() {
  return localStorage.getItem('token')
}

async function fetchCategories() {
  const res = await fetch('http://localhost:8080/categories', {
    headers: { Authorization: 'Bearer ' + getToken() }
  })
  categories.value = await res.json()
}

function edit(cat) {
  editId.value = cat.id
  form.value = { ...cat }
  showModal.value = true
}

function close() {
  showModal.value = false
  editId.value = null
  form.value = { name: '' }
}

async function submit() {
  const method = editId.value ? 'PUT' : 'POST'
  const url = editId.value ? `http://localhost:8080/categories/${editId.value}` : 'http://localhost:8080/categories'
  await fetch(url, {
    method,
    headers: {
      'Content-Type': 'application/json',
      Authorization: 'Bearer ' + getToken()
    },
    body: JSON.stringify(form.value)
  })
  await fetchCategories()
  close()
}

function remove(id) {
  deleteId = id
  showDeleteModal.value = true
}

async function confirmDelete() {
  if (deleteId) {
    await fetch(`http://localhost:8080/categories/${deleteId}`, {
      method: 'DELETE',
      headers: {
        Authorization: 'Bearer ' + getToken()
      }
    })
    categories.value = categories.value.filter(cat => cat.id !== deleteId)
    deleteId = null
    showDeleteModal.value = false
  }
}

onMounted(() => {
  fetchCategories()
})
</script>
