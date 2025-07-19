<template>
  <div class="pagination-container">
    <div class="pagination-info">
      <span class="text-body-2 text-medium-emphasis">
        Showing {{ startItem }}-{{ endItem }} of {{ total }} results
      </span>
    </div>

    <div class="pagination-controls">
      <v-btn
        variant="outlined"
        size="small"
        :disabled="currentPage === 1 || loading"
        @click="goToPage(1)"
        class="mr-2"
      >
        <v-icon size="16">fas fa-angle-double-left</v-icon>
      </v-btn>

      <v-btn
        variant="outlined"
        size="small"
        :disabled="currentPage === 1 || loading"
        @click="goToPage(currentPage - 1)"
        class="mr-2"
      >
        <v-icon size="16">fas fa-angle-left</v-icon>
        <span class="d-none d-sm-inline ml-1">Previous</span>
      </v-btn>

      <div class="pagination-pages d-none d-md-flex">
        <v-btn
          v-for="page in visiblePages"
          :key="page"
          :variant="page === currentPage ? 'flat' : 'text'"
          :color="page === currentPage ? 'primary' : 'default'"
          size="small"
          :disabled="loading"
          @click="goToPage(page)"
          class="mx-1"
        >
          {{ page }}
        </v-btn>
      </div>

      <div class="d-flex d-md-none align-center mx-3">
        <span class="text-body-2">Page {{ currentPage }} of {{ totalPages }}</span>
      </div>

      <v-btn
        variant="outlined"
        size="small"
        :disabled="currentPage === totalPages || loading"
        @click="goToPage(currentPage + 1)"
        class="ml-2"
      >
        <span class="d-none d-sm-inline mr-1">Next</span>
        <v-icon size="16">fas fa-angle-right</v-icon>
      </v-btn>

      <v-btn
        variant="outlined"
        size="small"
        :disabled="currentPage === totalPages || loading"
        @click="goToPage(totalPages)"
        class="ml-2"
      >
        <v-icon size="16">fas fa-angle-double-right</v-icon>
      </v-btn>
    </div>

    <div class="pagination-size d-none d-lg-flex">
      <v-select
        v-model="localPageSize"
        :items="pageSizeOptions"
        :disabled="loading"
        variant="outlined"
        density="compact"
        hide-details
        style="min-width: 100px;"
        @update:model-value="handlePageSizeChange"
      >
        <template #selection="{ item }">
          {{ item.title }}
        </template>
      </v-select>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, ref, watch } from 'vue'

interface PaginationProps {
  currentPage: number
  total: number
  pageSize?: number
  loading?: boolean
  showPageSize?: boolean
  maxVisiblePages?: number
  pageSizes?: number[]
}

const props = withDefaults(defineProps<PaginationProps>(), {
  pageSize: 20,
  loading: false,
  showPageSize: true,
  maxVisiblePages: 5,
  pageSizes: () => [10, 20, 50, 100]
})

const emit = defineEmits<{
  'page-change': [page: number]
  'page-size-change': [size: number]
}>()

const localPageSize = ref(props.pageSize)

const totalPages = computed(() => Math.ceil(props.total / props.pageSize))
const startItem = computed(() => (props.currentPage - 1) * props.pageSize + 1)
const endItem = computed(() => Math.min(props.currentPage * props.pageSize, props.total))

const pageSizeOptions = computed(() =>
  props.pageSizes.map(size => ({
    title: `${size} per page`,
    value: size
  }))
)

const visiblePages = computed(() => {
  const pages: number[] = []
  const half = Math.floor(props.maxVisiblePages / 2)
  let start = Math.max(1, props.currentPage - half)
  const end = Math.min(totalPages.value, start + props.maxVisiblePages - 1)

  if (end - start < props.maxVisiblePages - 1) {
    start = Math.max(1, end - props.maxVisiblePages + 1)
  }

  for (let i = start; i <= end; i++) {
    pages.push(i)
  }

  return pages
})

function goToPage(page: number) {
  if (page >= 1 && page <= totalPages.value && page !== props.currentPage) {
    emit('page-change', page)
  }
}

function handlePageSizeChange(newSize: number) {
  emit('page-size-change', newSize)
}

watch(() => props.pageSize, (newSize) => {
  localPageSize.value = newSize
})
</script>

<style scoped>
.pagination-container {
  display: flex;
  justify-content: space-between;
  align-items: center;
  gap: 1rem;
  padding: 1rem 0;
  font-family: 'Inter', system-ui, sans-serif;
  flex-wrap: wrap;
}

.pagination-controls {
  display: flex;
  align-items: center;
}

.pagination-pages {
  margin: 0 1rem;
}

.pagination-info {
  min-width: 150px;
}

.pagination-size {
  min-width: 120px;
}

@media (max-width: 768px) {
  .pagination-container {
    flex-direction: column;
    gap: 0.5rem;
  }

  .pagination-info {
    order: 2;
    text-align: center;
  }

  .pagination-controls {
    order: 1;
  }
}
</style>
