<template>
  <component :is="link ? RouterLink : 'div'" :to="link" class="kpi-card-wrapper">
    <v-card class="h-100 transition-all hover:shadow-md" elevation="4">
      <v-card-title class="d-flex justify-space-between align-center pa-4 pb-2">
        <span class="text-sm font-medium">{{ title }}</span>
        <v-icon :icon="icon" size="16" color="primary" />
      </v-card-title>
      <v-card-text class="pa-4 pt-0">
        <div class="text-2xl font-weight-bold">
          {{ value }}
        </div>
      </v-card-text>
    </v-card>
  </component>
</template>

<script setup lang="ts">
import { RouterLink } from 'vue-router'

interface KpiCardProps {
  title: string
  value: string | number
  icon: string
  link?: string
}

defineProps<KpiCardProps>()
</script>

<style scoped>
.kpi-card-wrapper {
  display: block;
  height: 100%;
  text-decoration: none;
}
</style>
