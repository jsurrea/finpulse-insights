<template>
  <v-card class="analysis-form">
    <v-card-title class="pb-2">
      <span class="text-lg font-weight-bold">Analyze Stock</span>
    </v-card-title>
    <v-card-subtitle class="pb-4">
      Enter a stock ticker and company name to begin.
    </v-card-subtitle>

    <v-card-text>
      <v-form @submit.prevent="onSubmit">
        <div class="form-fields">
          <v-text-field
            v-model="ticker"
            label="Ticker Symbol"
            placeholder="e.g., AAPL, GOOGL"
            variant="outlined"
            density="comfortable"
            :error-messages="errors.ticker"
            :disabled="loading"
            @blur="validateField('ticker')"
            class="mb-4"
          />

          <v-text-field
            v-model="company"
            label="Company Name"
            placeholder="e.g., Apple Inc."
            variant="outlined"
            density="comfortable"
            :error-messages="errors.company"
            :disabled="loading"
            @blur="validateField('company')"
            class="mb-6"
          />

          <v-btn
            type="submit"
            color="primary"
            block
            size="large"
            :loading="loading"
            :disabled="!isFormValid || loading"
          >
            <template #prepend>
              <v-icon>fas fa-sparkles</v-icon>
            </template>
            {{ loading ? 'Analyzing...' : 'Generate Analysis' }}
          </v-btn>
        </div>
      </v-form>
    </v-card-text>
  </v-card>
</template>

<script setup lang="ts">
import { ref, computed, watch, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { useForm } from 'vee-validate'
import { toTypedSchema } from '@vee-validate/yup'
import { useAIAnalyst } from '@/stores/aiAnalyst'
import { aiAnalysisSchema, type AIAnalysisFormData } from '@/utils/validation'

const emit = defineEmits<{
  'analysis-submitted': [data: AIAnalysisFormData]
}>()

const route = useRoute()
const store = useAIAnalyst()

const { handleSubmit, errors, validateField, setFieldValue, meta } = useForm<AIAnalysisFormData>({
  validationSchema: toTypedSchema(aiAnalysisSchema),
  initialValues: {
    ticker: (route.query.ticker as string) || '',
    company: (route.query.company as string) || ''
  }
})

const ticker = ref((route.query.ticker as string) || '')
const company = ref((route.query.company as string) || '')

const loading = computed(() => store.loading)
const isFormValid = computed(() => meta.value.valid)

// Sync with form validation
watch(ticker, (newValue) => {
  setFieldValue('ticker', newValue.toUpperCase())
})

watch(company, (newValue) => {
  setFieldValue('company', newValue)
})

const onSubmit = handleSubmit(async (values) => {
  emit('analysis-submitted', values)
})

onMounted(() => {
  // Set initial values from query params
  if (route.query.ticker) {
    ticker.value = route.query.ticker as string
  }
  if (route.query.company) {
    company.value = route.query.company as string
  }
})
</script>

<style scoped>
.analysis-form {
  font-family: 'Inter', system-ui, sans-serif;
}

.form-fields {
  display: flex;
  flex-direction: column;
}
</style>
