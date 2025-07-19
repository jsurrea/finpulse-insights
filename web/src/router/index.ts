import { createRouter, createWebHistory, type RouteRecordRaw } from 'vue-router';

import AppLayout from '@/layouts/AppLayout.vue';

import Dashboard from '@/pages/Dashboard.vue';
import Analytics from '@/pages/Analytics.vue';
import Health from '@/pages/Health.vue';

import StocksIndex from '@/pages/Stocks/Index.vue';
import StockDetail from '@/pages/Stocks/_ticker.vue';

import RecommendationsIndex from '@/pages/Recommendations/Index.vue';
import RecommendationDetail from '@/pages/Recommendations/_id.vue';
import AiAnalyst from '@/pages/AiAnalyst.vue';

const routes: RouteRecordRaw[] = [
  {
    path: '/',
    component: AppLayout,
    children: [
      { path: '', redirect: '/dashboard' },
      { path: 'dashboard', name: 'dashboard', component: Dashboard },
      { path: 'analytics', name: 'analytics', component: Analytics },
      { path: 'ai-analyst', name: 'ai-analyst', component: AiAnalyst },
      { path: 'health', name: 'health', component: Health },
      {
        path: 'stocks',
        children: [
          { path: '', name: 'stocks', component: StocksIndex },
          { path: ':ticker', name: 'stock-detail', component: StockDetail, props: true },
        ]
      },
      {
        path: 'recommendations',
        children: [
          { path: '', name: 'recommendations', component: RecommendationsIndex },
          { path: ':id', name: 'recommendation-detail', component: RecommendationDetail, props: true },
        ]
      }
    ]
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
