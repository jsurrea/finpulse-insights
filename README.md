# FinPulse Insights

FinPulse Insights is a full-stack stock-analytics platform that aggregates brokerage recommendations, live market data and AI-driven commentary into a single Vue 3 + Vuetify SPA backed by a lightweight Go API.

---

## ✨ Core Features

• **Dashboard KPIs & Charts** – instant view of total tickers tracked, buy/sell ratios, confidence averages, distribution pie and top-brokerage ranking.  
• **Stocks Module** – paginated, filterable list of 2.8 K tickers plus detail view with latest action and rating trail.  
• **Recommendations Module** – multi-field filters (ticker, brokerage, date range, action); confidence badges and deep-links.  
• **AI Financial Analyst** – dropdown to pick any stock; hits an n8n-cloud webhook that pulls fresh NewsAPI headlines and uses OpenAI GPT-4o-mini to return summary, BUY/SELL/HOLD, confidence, key factors and risk. Last 10 analyses persist in localStorage.  
• **System Health** – API / DB status cards with auto-refresh.

---

## 🛠️ Tech Stack

| Layer         | Tech                                              |
| ------------- | ------------------------------------------------- |
| Front-end     | Vue 3 · TypeScript · Vuetify 3 · Pinia · Chart.js |
| Back-end      | Go 1.24 (Gin) · GORM · CockroachDB                |
| External Data | Brokerage dataset (2.8 K rows) · NewsAPI          |
| AI Flows      | n8n Cloud · Google Gemini 2.5 Pro                 |

---

## 📂 Monorepo Layout

```
admin/ → data-loader & enrichment helpers
api/ → Go backend (handlers, models, utils)
infra/ → Terraform scripts for GCP deployment
web/ → Vue SPA (components, pages, stores, utils)
```

Each directory contains its own quick-start commands; check the nested READMEs for details.
