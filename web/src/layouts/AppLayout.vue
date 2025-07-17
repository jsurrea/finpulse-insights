<template>
  <AppHeader />
  <AppSidebar />
  <v-main>
    <router-view />
  </v-main>
  <!-- Toaster, Dialogs, etc. -->
</template>

<script setup lang="ts">
import AppHeader from '@/components/layout/AppHeader.vue'
import AppSidebar from '@/components/layout/AppSidebar.vue'
</script>
