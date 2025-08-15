<template>
  <div class="response-panel">
    <div class="response-header">
      <h3>Response</h3>
      <div class="response-status" v-if="response.status">
        Status: {{ response.status }} {{ response.statusText }}
      </div>
    </div>
    
    <div class="response-content">
      <div v-if="!response.data" class="empty-state">
        Send a request to see the response
      </div>
      
      <div v-else class="response-data">
        <pre class="response-json">{{ formattedResponse }}</pre>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'

const response = ref({
  status: null,
  statusText: '',
  data: null
})

const formattedResponse = computed(() => {
  if (!response.value.data) return ''
  try {
    return JSON.stringify(response.value.data, null, 2)
  } catch {
    return response.value.data.toString()
  }
})
</script>

<style scoped>
.response-panel {
  flex: 1;
  display: flex;
  flex-direction: column;
  background: white;
  overflow: hidden;
}

.response-header {
  padding: 16px 20px;
  border-bottom: 1px solid #e1e5e9;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.response-header h3 {
  font-size: 14px;
  font-weight: 600;
  color: #24292e;
}

.response-status {
  font-size: 13px;
  color: #586069;
  background: #f6f8fa;
  padding: 4px 8px;
  border-radius: 4px;
}

.response-content {
  flex: 1;
  overflow: auto;
  padding: 20px;
}

.empty-state {
  display: flex;
  align-items: center;
  justify-content: center;
  height: 100%;
  color: #6a737d;
  font-size: 14px;
}

.response-data {
  height: 100%;
}

.response-json {
  background: #f6f8fa;
  padding: 16px;
  border-radius: 6px;
  font-family: 'Courier New', monospace;
  line-height: 1.4;
  overflow: auto;
  max-height: 100%;
  margin: 0;
}
</style>