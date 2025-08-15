<template>
  <div class="response-panel">
    <div class="response-header">
      <h3>Response</h3>
      <div class="response-status" v-if="response">
        Status: {{ response.status_code }}
      </div>
    </div>
    
    <div class="response-content">
      <div v-if="!response" class="empty-state">
        Send a request to see the response
      </div>
      
      <div v-else class="response-data">
        <div class="response-meta">
          <div><strong>Name:</strong> {{ response.name }}</div>
          <div><strong>Endpoint:</strong> {{ response.api_endpoint }}</div>
          <div><strong>Created:</strong> {{ formatDate(response.created_at) }}</div>
        </div>
        <pre class="response-json">{{ formattedResponse }}</pre>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, defineProps } from 'vue'

const props = defineProps({
  response: {
    type: Object,
    default: null
  }
})

const formattedResponse = computed(() => {
  if (!props.response || !props.response.response) return ''
  try {
    return JSON.stringify(props.response.response, null, 2)
  } catch {
    return props.response.response.toString()
  }
})

const formatDate = (dateString) => {
  if (!dateString) return ''
  return new Date(dateString).toLocaleString()
}
</script>

<style scoped src="./ResponsePanel.css"></style>