<template>
  <div class="request-form">
    <div class="request-header">
      <!-- Поле для имени теста -->
      <div class="test-name">
        <input 
          v-model="testName"
          type="text" 
          placeholder="Test name" 
          class="test-name-input"
        />
      </div>
      
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
        <button @click="sendRequest" class="send-btn" :disabled="loading">
          <span v-if="loading">⏱️</span>
          <span v-else>📤</span>
          {{ loading ? 'Sending...' : 'Send' }}
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
      <button 
        class="tab-btn"
        :class="{ active: activeTab === 'params' }"
        @click="activeTab = 'params'"
      >
        Parameters
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

      <!-- Parameters Section -->
      <div v-show="activeTab === 'params'" class="params-section">
        <div class="section-header">
          <span class="section-title">Query Parameters</span>
          <button @click="addParam" class="add-btn">+ Add Parameter</button>
        </div>
        
        <div class="params-list">
          <div 
            v-for="(param, index) in params" 
            :key="index"
            class="param-row"
          >
            <input 
              v-model="param.key"
              type="text" 
              placeholder="Key"
              class="param-input"
            />
            <input 
              v-model="param.value"
              type="text" 
              placeholder="Value"
              class="param-input"
            />
            <button @click="removeParam(index)" class="remove-btn">×</button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, defineEmits } from 'vue'

const emit = defineEmits(['requestSent'])

const testName = ref('')
const method = ref('GET')
const url = ref('')
const activeTab = ref('headers')
const headers = ref([{ key: '', value: '' }])
const body = ref('')
const params = ref([{ key: '', value: '' }])
const loading = ref(false)

const addHeader = () => {
  headers.value.push({ key: '', value: '' })
}

const removeHeader = (index) => {
  headers.value.splice(index, 1)
}

const addParam = () => {
  params.value.push({ key: '', value: '' })
}

const removeParam = (index) => {
  params.value.splice(index, 1)
}

const sendRequest = async () => {
  if (!testName.value.trim()) {
    alert('Please enter a test name')
    return
  }

  if (!url.value.trim()) {
    alert('Please enter a URL')
    return
  }

  loading.value = true

  try {
    // Формируем данные для отправки
    const requestData = {
      name: testName.value,
      api_endpoint: url.value,
      method: method.value,
      parameters: {},
      headers: {},
      body: body.value ? JSON.parse(body.value) : null
    }

    // Добавляем параметры
    params.value.forEach(param => {
      if (param.key && param.value) {
        requestData.parameters[param.key] = param.value
      }
    })

    // Добавляем заголовки
    headers.value.forEach(header => {
      if (header.key && header.value) {
        requestData.headers[header.key] = header.value
      }
    })

    const response = await fetch('http://localhost:8080/api/tests', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(requestData)
    })

    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`)
    }

    const result = await response.json()
    console.log('Request sent successfully:', result)
    
    // Отправляем результат в родительский компонент
    emit('requestSent', result)
    
    // Сбрасываем форму (опционально)
    // resetForm()
    
  } catch (error) {
    console.error('Error sending request:', error)
    alert('Error sending request: ' + error.message)
  } finally {
    loading.value = false
  }
}

const saveRequest = () => {
  console.log('Сохранение запроса:', {
    testName: testName.value,
    method: method.value,
    url: url.value,
    headers: headers.value.filter(h => h.key && h.value),
    body: body.value,
    params: params.value.filter(p => p.key && p.value)
  })
}

const resetForm = () => {
  testName.value = ''
  method.value = 'GET'
  url.value = ''
  headers.value = [{ key: '', value: '' }]
  body.value = ''
  params.value = [{ key: '', value: '' }]
}
</script>

<style scoped src="./RequestForm.css"></style>