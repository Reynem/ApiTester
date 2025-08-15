<template>
  <div class="request-form">
    <div class="request-header">
      <div class="method-url">
        <select v-model="method" class="method-select">
          <option>GET</option>
          <option>POST</option>
          <option>PUT</option>
          <option>DELETE</option>
          <option>PATCH</option>
        </select>
        
        <input 
          v-model="url"
          type="text" 
          placeholder="Enter URL" 
          class="url-input"
        />
      </div>
      
      <div class="request-actions">
        <button @click="sendRequest" class="send-btn">
          <span>📤</span> Send
        </button>
        <button @click="saveRequest" class="save-btn">
          <span>💾</span> Save
        </button>
      </div>
    </div>
    
    <div class="request-tabs">
      <button 
        class="tab-btn"
        :class="{ active: activeTab === 'headers' }"
        @click="activeTab = 'headers'"
      >
        Headers
      </button>
      <button 
        class="tab-btn"
        :class="{ active: activeTab === 'body' }"
        @click="activeTab = 'body'"
      >
        Body
      </button>
    </div>
    
    <div class="request-body">
      <!-- Headers Section -->
      <div v-show="activeTab === 'headers'" class="headers-section">
        <div class="section-header">
          <span class="section-title">Headers</span>
          <button @click="addHeader" class="add-btn">+ Add Header</button>
        </div>
        
        <div class="headers-list">
          <div 
            v-for="(header, index) in headers" 
            :key="index"
            class="header-row"
          >
            <input 
              v-model="header.key"
              type="text" 
              placeholder="Key"
              class="header-input"
            />
            <input 
              v-model="header.value"
              type="text" 
              placeholder="Value"
              class="header-input"
            />
            <button @click="removeHeader(index)" class="remove-btn">×</button>
          </div>
        </div>
      </div>
      
      <!-- Body Section -->
      <div v-show="activeTab === 'body'" class="body-section">
        <div class="section-header">
          <span class="section-title">Body</span>
        </div>
        <textarea 
          v-model="body"
          placeholder="Enter request body (JSON, XML, etc.)"
          class="body-textarea"
        ></textarea>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'

const method = ref('GET')
const url = ref('')
const activeTab = ref('headers')
const headers = ref([
  { key: '', value: '' }
])
const body = ref('')

const addHeader = () => {
  headers.value.push({ key: '', value: '' })
}

const removeHeader = (index) => {
  headers.value.splice(index, 1)
}

const sendRequest = () => {
  console.log('Отправка запроса:', {
    method: method.value,
    url: url.value,
    headers: headers.value.filter(h => h.key && h.value),
    body: body.value
  })
}

const saveRequest = () => {
  console.log('Сохранение запроса')
}
</script>

<style scoped>
.request-form {
  background: white;
  border-bottom: 1px solid #e1e5e9;
  padding: 20px;
}

.request-header {
  display: flex;
  gap: 12px;
  margin-bottom: 16px;
  flex-wrap: wrap;
}

.method-url {
  display: flex;
  flex: 1;
  gap: 8px;
  min-width: 300px;
}

.method-select {
  padding: 8px 12px;
  border: 1px solid #e1e5e9;
  border-radius: 4px;
  background: white;
  font-size: 14px;
  cursor: pointer;
}

.method-select:focus {
  outline: none;
  border-color: #0366d6;
  box-shadow: 0 0 0 2px rgba(3, 102, 214, 0.2);
}

.url-input {
  flex: 1;
  padding: 8px 12px;
  border: 1px solid #e1e5e9;
  border-radius: 4px;
  font-size: 14px;
}

.url-input:focus {
  outline: none;
  border-color: #0366d6;
  box-shadow: 0 0 0 2px rgba(3, 102, 214, 0.2);
}

.request-actions {
  display: flex;
  gap: 8px;
}

.send-btn, .save-btn {
  padding: 8px 16px;
  border: none;
  border-radius: 4px;
  font-size: 14px;
  cursor: pointer;
  display: flex;
  align-items: center;
  gap: 6px;
  transition: background-color 0.2s;
}

.send-btn {
  background-color: #28a745;
  color: white;
}

.send-btn:hover {
  background-color: #218838;
}

.save-btn {
  background-color: white;
  color: #24292e;
  border: 1px solid #e1e5e9;
}

.save-btn:hover {
  background-color: #f6f8fa;
}

.request-tabs {
  display: flex;
  gap: 2px;
  margin-bottom: 16px;
}

.tab-btn {
  padding: 6px 12px;
  border: none;
  background: #f6f8fa;
  border-radius: 4px;
  font-size: 13px;
  cursor: pointer;
  transition: background-color 0.2s;
}

.tab-btn.active {
  background: #0366d6;
  color: white;
}

.request-body {
  min-height: 200px;
}

.section-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 12px;
}

.section-title {
  font-weight: 600;
  font-size: 14px;
  color: #24292e;
}

.add-btn {
  padding: 4px 8px;
  border: none;
  background: none;
  color: #0366d6;
  font-size: 13px;
  cursor: pointer;
  border-radius: 4px;
}

.add-btn:hover {
  background-color: rgba(3, 102, 214, 0.1);
}

.headers-list {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.header-row {
  display: flex;
  gap: 8px;
  align-items: center;
}

.header-input {
  flex: 1;
  padding: 6px 10px;
  border: 1px solid #e1e5e9;
  border-radius: 4px;
  font-size: 13px;
}

.header-input:focus {
  outline: none;
  border-color: #0366d6;
  box-shadow: 0 0 0 2px rgba(3, 102, 214, 0.2);
}

.remove-btn {
  width: 24px;
  height: 24px;
  border: none;
  background: #dc3545;
  color: white;
  border-radius: 4px;
  cursor: pointer;
  font-size: 16px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.remove-btn:hover {
  background: #c82333;
}

.body-textarea {
  width: 100%;
  height: 150px;
  padding: 12px;
  border: 1px solid #e1e5e9;
  border-radius: 4px;
  font-family: 'Courier New', monospace;
  font-size: 13px;
  resize: vertical;
}

.body-textarea:focus {
  outline: none;
  border-color: #0366d6;
  box-shadow: 0 0 0 2px rgba(3, 102, 214, 0.2);
}
</style>