import { defineStore } from 'pinia'

export const useUI = defineStore('ui', {
  state: () => ({
    darkMode: false,
    sidebarOpen: true
  }),
  actions: {
    toggleDarkMode() {
      this.darkMode = !this.darkMode
      if (this.darkMode) {
        document.documentElement.classList.add('dark')
      } else {
        document.documentElement.classList.remove('dark')
      }
      localStorage.setItem('theme', this.darkMode ? 'dark' : 'light')
    },
    initTheme() {
      const stored = localStorage.getItem('theme')
      if (stored === 'dark') {
        this.darkMode = true
        document.documentElement.classList.add('dark')
      } else if (stored === 'light') {
        this.darkMode = false
        document.documentElement.classList.remove('dark')
      } else {
        const prefersDark = window.matchMedia('(prefers-color-scheme: dark)').matches
        this.darkMode = prefersDark
        if (prefersDark) {
          document.documentElement.classList.add('dark')
        } else {
          document.documentElement.classList.remove('dark')
        }
        localStorage.setItem('theme', this.darkMode ? 'dark' : 'light')
      }
    },
    toggleSidebar() {
      this.sidebarOpen = !this.sidebarOpen
    },
    setSidebarOpen(open: boolean) {
      this.sidebarOpen = open
    }
  }
})
