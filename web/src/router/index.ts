import { createRouter, createWebHistory, type RouteRecordRaw } from 'vue-router'

import AppLayout from '@/layouts/AppLayout.vue'

import Dashboard from '@/pages/dashboard/DashboardPage.vue'
import Health from '@/pages/health/HealthPage.vue'

import StockList from '@/pages/stocks/StockList.vue'
import StockDetail from '@/pages/stocks/StockDetail.vue'

import RecommendationsIndex from '@/pages/recommendations/RecommendationList.vue'
import RecommendationDetail from '@/pages/recommendations/RecommendationDetail.vue'
import AiAnalyst from '@/pages/AiAnalyst.vue'

const routes: RouteRecordRaw[] = [
  {
    path: '/',
    component: AppLayout,
    children: [
      { path: '', redirect: '/dashboard' },
      { path: 'dashboard', name: 'dashboard', component: Dashboard },
      { path: 'ai-analyst', name: 'ai-analyst', component: AiAnalyst },
      { path: 'health', name: 'health', component: Health },
      {
        path: 'stocks',
        children: [
          { path: '', name: 'stocks', component: StockList },
          { path: ':ticker', name: 'stock-detail', component: StockDetail, props: true },
        ],
      },
      {
        path: 'recommendations',
        children: [
          { path: '', name: 'recommendations', component: RecommendationsIndex },
          {
            path: ':id',
            name: 'recommendation-detail',
            component: RecommendationDetail,
            props: true,
          },
        ],
      },
    ],
  },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router
