<template>
  <div class="min-h-screen flex flex-col bg-background text-foreground">
    <AppHeader />
    <div class="flex flex-1">
      <AppSidebar />
      <main class="flex-1 p-4">
        <router-view />
      </main>
    </div>
    <!-- Toaster, Dialogs, etc. -->
  </div>
</template>

<script setup lang="ts">
import AppHeader from '@/components/layout/AppHeader.vue'
import AppSidebar from '@/components/layout/AppSidebar.vue'
</script>
