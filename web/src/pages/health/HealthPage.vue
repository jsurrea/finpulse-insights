<template>
  <section class="health-page">
    <div class="mb-8">
      <div class="d-flex justify-space-between align-center mb-2">
        <div>
          <h1 class="text-h4 font-weight-bold mb-1">System Health</h1>
          <p class="text-medium-emphasis">Real-time status of the FinPulse API.</p>
        </div>
        <div class="d-flex align-center ga-2">
          <v-chip
            :color="systemStatusColor"
            variant="tonal"
            size="small"
            class="font-weight-medium"
          >
            <v-icon start :icon="systemStatusIcon" size="14" />
            {{ systemStatusText }}
          </v-chip>
          <v-btn
            icon
            variant="text"
            color="primary"
            :loading="store.loading"
            @click="store.fetchHealth"
            :disabled="store.loading"
            aria-label="Refresh health status"
          >
            <v-icon>fas fa-sync-alt</v-icon>
          </v-btn>
        </div>
      </div>

      <div v-if="store.lastUpdated" class="text-caption text-medium-emphasis">
        Last updated: {{ formatLastUpdated(store.lastUpdated) }}
      </div>
    </div>

    <v-row v-if="store.healthStatus">
      <v-col cols="12" md="4">
        <StatusCard
          title="API Status"
          :value="store.healthStatus.status"
          description="Main API endpoint health"
          :status="mapStatus(store.healthStatus.status)"
          icon="fas fa-check-circle"
        />
      </v-col>

      <v-col cols="12" md="4">
        <StatusCard
          title="Database"
          :value="store.healthStatus.database"
          description="Database connection status"
          :status="store.healthStatus.database"
          icon="fas fa-database"
        />
      </v-col>

      <v-col cols="12" md="4">
        <StatusCard
          title="Uptime"
          :value="formattedUptime"
          description="Since last restart"
          status="info"
          icon="fas fa-power-off"
        />
      </v-col>
    </v-row>

    <!-- Loading State -->
    <div v-if="store.loading && !store.healthStatus" class="d-flex justify-center py-12">
      <div class="text-center">
        <v-progress-circular indeterminate color="primary" size="64" class="mb-4" />
        <div class="text-body-1">Checking system health...</div>
      </div>
    </div>

    <!-- Error State -->
    <v-alert
      v-if="store.error"
      type="error"
      class="mt-6"
      variant="tonal"
      closable
      @click:close="store.error = null"
    >
      <template #title>Health Check Failed</template>
      {{ store.error }}
    </v-alert>

    <!-- No Data State -->
    <v-card v-if="!store.healthStatus && !store.loading && !store.error" class="mt-6">
      <v-card-text class="text-center py-12">
        <v-icon size="64" color="grey" class="mb-4"> fas fa-question-circle </v-icon>
        <div class="text-h6 mb-2">No Health Data Available</div>
        <div class="text-body-2 text-medium-emphasis mb-4">
          Unable to retrieve system health information.
        </div>
        <v-btn color="primary" @click="store.fetchHealth"> Try Again </v-btn>
      </v-card-text>
    </v-card>
  </section>
</template>

<script setup lang="ts">
import { computed, onMounted } from 'vue'
import { useHealth } from '@/stores/health'
import { formatDistanceToNowStrict, format } from 'date-fns'
import StatusCard from '@/components/health/StatusCard.vue'

const store = useHealth()

const formattedUptime = computed(() => {
  if (!store.healthStatus) return 'N/A'
  return formatUptime(store.healthStatus.uptime)
})

const systemStatusColor = computed(() => {
  switch (store.systemStatus) {
    case 'healthy':
      return 'success'
    case 'critical':
      return 'error'
    case 'degraded':
      return 'warning'
    default:
      return 'grey'
  }
})

const systemStatusIcon = computed(() => {
  switch (store.systemStatus) {
    case 'healthy':
      return 'fas fa-check-circle'
    case 'critical':
      return 'fas fa-times-circle'
    case 'degraded':
      return 'fas fa-exclamation-triangle'
    default:
      return 'fas fa-question-circle'
  }
})

const systemStatusText = computed(() => {
  switch (store.systemStatus) {
    case 'healthy':
      return 'All Systems Operational'
    case 'critical':
      return 'Critical Issues'
    case 'degraded':
      return 'Degraded Performance'
    default:
      return 'Status Unknown'
  }
})

onMounted(() => {
  store.fetchHealth()
})

function formatUptime(seconds: number): string {
  const now = Date.now()
  const startTime = now - seconds * 1000
  return formatDistanceToNowStrict(startTime, { addSuffix: true })
}

function formatLastUpdated(date: Date): string {
  return format(date, 'MMM d, yyyy HH:mm:ss')
}

function mapStatus(status: 'error' | 'ok' | 'degraded'): 'error' | 'info' | 'ok' | 'connected' | 'disconnected' {
  if (status === 'degraded') return 'info'
  return status
}
</script>

<style scoped>
.health-page {
  max-width: 1200px;
  margin: 0 auto;
  font-family: 'Inter', system-ui, sans-serif;
}
</style>
