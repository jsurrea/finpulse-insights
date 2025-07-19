<template>
  <div class="confidence-gauge">
    <div class="d-flex justify-space-between align-center mb-2">
      <span class="text-body-2 font-weight-medium">Confidence Level</span>
      <span class="text-body-2 font-weight-bold font-mono">{{ displayValue }}%</span>
    </div>
    <v-progress-linear
      :model-value="percentage"
      :color="progressColor"
      height="8"
      rounded
      bg-opacity="0.2"
    />
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'

interface ConfidenceGaugeProps {
  confidence: number // Expected to be between 0 and 1
}

const props = defineProps<ConfidenceGaugeProps>()

const percentage = computed(() => props.confidence * 100)
const displayValue = computed(() => Math.round(percentage.value))

const progressColor = computed(() => {
  if (percentage.value >= 80) return 'success'
  if (percentage.value >= 60) return 'primary'
  if (percentage.value >= 40) return 'warning'
  return 'error'
})
</script>

<style scoped>
.confidence-gauge {
  width: 100%;
}
.font-mono {
  font-family: 'IBM Plex Mono', monospace;
}
</style>
