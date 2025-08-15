<template>
  <div class="app">
    <Sidebar />
    
    <main class="main-content">
      <RequestForm @request-sent="handleRequestSent" />
      <ResponsePanel :response="currentResponse" />
    </main>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import Sidebar from '../components/Sidebar.vue'
import RequestForm from '../components/RequestForm.vue'
import ResponsePanel from '../components/ResponsePanel.vue'

const currentResponse = ref(null)

const handleRequestSent = (response) => {
  currentResponse.value = response
}
</script>

<style scoped>
.main-content {
  flex: 1;
  display: flex;
  flex-direction: column;
  overflow: hidden;
}
</style>