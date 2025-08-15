<template>
  <Dashboard />
</template>

<script setup>
import Dashboard from './views/Dashboard.vue'
</script>

<style>
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

body {
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
  background-color: #f5f5f5;
}

.app {
  display: flex;
  height: 100vh;
  overflow: hidden;
}
</style>