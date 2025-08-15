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

<style scoped>
.sidebar {
  width: 280px;
  background-color: #f8f9fa;
  border-right: 1px solid #e1e5e9;
  display: flex;
  flex-direction: column;
}

.sidebar-header {
  padding: 16px;
  border-bottom: 1px solid #e1e5e9;
}

.sidebar-header h1 {
  font-size: 16px;
  font-weight: 600;
  color: #24292e;
}

.sidebar-nav ul {
  list-style: none;
  padding: 8px;
}

.nav-item {
  width: 100%;
  padding: 8px 12px;
  border: none;
  background: none;
  text-align: left;
  border-radius: 6px;
  cursor: pointer;
  font-size: 14px;
  color: #586069;
  display: flex;
  align-items: center;
  transition: background-color 0.2s;
}

.nav-item:hover {
  background-color: #e1e5e9;
}

.nav-item.active {
  background-color: #0366d6;
  color: white;
}

.nav-icon {
  margin-right: 8px;
  font-size: 16px;
}
</style>