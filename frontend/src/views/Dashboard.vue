<template>
  <div class="app">
    <Sidebar />
    
    <main class="main-content">
      <RequestForm />
      <ResponsePanel />
    </main>
  </div>
</template>

<script setup>
import Sidebar from '../components/Sidebar.vue'
import RequestForm from '../components/RequestForm.vue'
import ResponsePanel from '../components/ResponsePanel.vue'
</script>

<style scoped>
.main-content {
  flex: 1;
  display: flex;
  flex-direction: column;
  overflow: hidden;
}
</style>