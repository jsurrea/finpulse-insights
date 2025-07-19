<template>
  <v-chip
    :color="chipColor"
    :variant="chipVariant"
    class="font-weight-bold recommendation-badge"
    size="small"
  >
    {{ recommendation }}
  </v-chip>
</template>

<script setup lang="ts">
import { computed } from 'vue'

interface RecommendationBadgeProps {
  recommendation: 'BUY' | 'SELL' | 'HOLD'
}

const props = defineProps<RecommendationBadgeProps>()

const chipColor = computed(() => {
  switch (props.recommendation) {
    case 'BUY': return 'success'
    case 'SELL': return 'error'
    case 'HOLD': return 'warning'
    default: return 'primary'
  }
})

const chipVariant = computed(() => 'tonal')
</script>

<style scoped>
.recommendation-badge {
  font-family: 'Inter', system-ui, sans-serif;
  letter-spacing: 0.05em;
}
</style>
