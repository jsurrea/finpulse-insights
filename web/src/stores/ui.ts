import { defineStore } from 'pinia'
import vuetify from '@/plugins/vuetify'

export const useUI = defineStore('ui', {
  state: () => ({
    dark: localStorage.getItem('theme') === 'dark',
    sidebar: window.innerWidth >= 768,
  }),
  actions: {
    toggleDarkMode() {
      this.dark = !this.dark
      vuetify.theme.global.name.value = this.dark ? 'myDark' : 'myLight'
      localStorage.setItem('theme', this.dark ? 'dark' : 'light')
    },
    toggleSidebar() {
      this.sidebar = !this.sidebar
    },
    setSidebar(val: boolean) {
      this.sidebar = val
    },
  },
})
