import 'vuetify/styles'
import { createVuetify } from 'vuetify'
import { aliases, mdi } from 'vuetify/iconsets/mdi'

const myTheme = {
  dark: false,
  colors: {
    background: '#F8F9FC',
    surface: '#FFFFFF',
    primary: '#05478A',
    secondary: '#048CFC',
    success: '#2E865F',
    warning: '#FF6B35',
    neutral: '#6B7280',
    error: '#B00020',
  }
}

export default createVuetify({
  theme: {
    defaultTheme: 'myTheme',
    themes: {
      myTheme,
    },
  },
  icons: {
    defaultSet: 'mdi',
    aliases,
    sets: { mdi }
  },
  defaults: {
    global: {
      style: {
        fontFamily: 'Inter, system-ui, sans-serif',
      },
    },
  }
})
