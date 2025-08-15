<template>
  <aside class="sidebar">
    <div class="sidebar-header">
      <h1>Collections & History</h1>
    </div>
    
    <nav class="sidebar-nav">
      <ul>
        <li>
          <button 
            class="nav-item"
            :class="{ active: activeTab === 'history' }"
            @click="activeTab = 'history'"
          >
            <span class="nav-icon">⏱️</span>
            History
          </button>
        </li>
        <li>
          <button 
            class="nav-item"
            :class="{ active: activeTab === 'collections' }"
            @click="activeTab = 'collections'"
          >
            <span class="nav-icon">📂</span>
            Collections
          </button>
        </li>
      </ul>
    </nav>
  </aside>
</template>

<script setup>
import { ref } from 'vue'

const activeTab = ref('history')
</script>

<style scoped src="./Sidebar.css"></style>