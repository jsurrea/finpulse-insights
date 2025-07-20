<template>
  <section class="health-page">
    <div class="mb-8">
      <div class="d-flex justify-space-between align-center mb-2">
        <div>
          <h1 class="text-h4 font-weight-bold mb-1">System Health</h1>
          <p class="text-body-2 text-medium-emphasis">Real-time status of the FinPulse API.</p>
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
          :status="store.healthStatus.status"
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

    <!-- System Overview Card -->
    <v-card v-if="store.healthStatus" class="mt-6">
      <v-card-title class="pa-4">
        <span class="text-lg font-weight-bold">System Overview</span>
      </v-card-title>
      <v-card-text class="pa-4">
        <v-row>
          <v-col cols="12" sm="6" md="3">
            <div class="text-center">
              <v-icon :color="store.isHealthy ? 'success' : 'error'" size="32" class="mb-2">
                {{ store.isHealthy ? 'fas fa-heart' : 'fas fa-heart-broken' }}
              </v-icon>
              <div class="text-body-2 font-weight-medium">
                {{ store.isHealthy ? 'Healthy' : 'Issues Detected' }}
              </div>
            </div>
          </v-col>
          <v-col cols="12" sm="6" md="3">
            <div class="text-center">
              <v-icon
                :color="store.isDatabaseConnected ? 'primary' : 'error'"
                size="32"
                class="mb-2"
              >
                fas fa-database
              </v-icon>
              <div class="text-body-2 font-weight-medium">
                Database {{ store.isDatabaseConnected ? 'Online' : 'Offline' }}
              </div>
            </div>
          </v-col>
          <v-col cols="12" sm="6" md="3">
            <div class="text-center">
              <v-icon color="info" size="32" class="mb-2"> fas fa-clock </v-icon>
              <div class="text-body-2 font-weight-medium">
                {{ formatUptimeDetailed(store.healthStatus.uptime) }}
              </div>
            </div>
          </v-col>
          <v-col cols="12" sm="6" md="3">
            <div class="text-center">
              <v-icon color="success" size="32" class="mb-2"> fas fa-shield-check </v-icon>
              <div class="text-body-2 font-weight-medium">Monitoring Active</div>
            </div>
          </v-col>
        </v-row>
      </v-card-text>
    </v-card>

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
import { computed } from 'vue'
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

function formatUptime(seconds: number): string {
  const now = Date.now()
  const startTime = now - seconds * 1000
  return formatDistanceToNowStrict(startTime, { addSuffix: true })
}

function formatUptimeDetailed(seconds: number): string {
  const hours = Math.floor(seconds / 3600)
  const days = Math.floor(hours / 24)

  if (days > 0) {
    return `${days}d ${hours % 24}h uptime`
  } else if (hours > 0) {
    return `${hours}h ${Math.floor((seconds % 3600) / 60)}m uptime`
  } else {
    return `${Math.floor(seconds / 60)}m uptime`
  }
}

function formatLastUpdated(date: Date): string {
  return format(date, 'MMM d, yyyy HH:mm:ss')
}
</script>

<style scoped>
.health-page {
  max-width: 1200px;
  margin: 0 auto;
  font-family: 'Inter', system-ui, sans-serif;
}
</style>
