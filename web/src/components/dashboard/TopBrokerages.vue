<template>
  <v-card class="h-100">
    <v-card-title class="pa-4">
      <span class="text-lg font-weight-bold">Top Brokerages</span>
    </v-card-title>
    <v-card-text>
      <v-list class="pa-0">
        <v-list-item v-for="(brokerage, index) in data.slice(0, 10)" :key="brokerage.name" class="px-0 py-3">
          <template #prepend>
            <v-avatar size="32" color="primary">
              <span class="text-sm font-weight-medium">{{ index + 1 }}</span>
            </v-avatar>
          </template>

          <v-list-item-title class="font-weight-medium text-truncate">
            {{ brokerage.name }}
          </v-list-item-title>

          <template #append>
            <v-chip size="small" color="primary" variant="outlined" class="font-mono ml-2">
              {{ brokerage.stockCount }} stocks
            </v-chip>
          </template>
        </v-list-item>
      </v-list>
    </v-card-text>
  </v-card>
</template>

<script setup lang="ts">
interface BrokerageData {
  name: string
  stockCount: number
}

interface TopBrokeragesProps {
  data: BrokerageData[]
}

defineProps<TopBrokeragesProps>()
</script>

<style scoped>
.font-mono {
  font-family: 'IBM Plex Mono', monospace;
}
</style>
