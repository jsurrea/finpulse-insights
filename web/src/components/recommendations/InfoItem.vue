<template>
  <v-card variant="outlined" class="info-item h-100">
    <v-card-text class="d-flex align-start ga-3 pa-3">
      <v-icon :icon="icon" size="20" class="mt-1 text-medium-emphasis" />
      <div class="flex-grow-1">
        <div class="text-body-2 text-medium-emphasis mb-1">{{ label }}</div>
        <div class="text-body-1 font-weight-medium">{{ value }}</div>
      </div>
    </v-card-text>
  </v-card>
</template>

<script setup lang="ts">
interface InfoItemProps {
  icon: string
  label: string
  value: string
}

defineProps<InfoItemProps>()
</script>

<style scoped>
.info-item {
  font-family: 'Inter', system-ui, sans-serif;
}
</style>
