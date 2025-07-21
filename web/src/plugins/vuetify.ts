import 'vuetify/styles'
import { createVuetify } from 'vuetify'
import { aliases, fa } from 'vuetify/iconsets/fa'

// Paleta personalizada LIGHT
const myLight = {
  dark: false,
  colors: {
    background: '#F8F9FC',
    surface: '#FFFFFF',
    primary: '#05478A',
    secondary: '#048CFC',
    success: '#2E865F',
    warning: '#FF6B35',
    error: '#B00020',
    onPrimary: '#FFFFFF',
    neutral: '#6B7280',
    text: '#333333',
  },
}

// Paleta personalizada DARK
const myDark = {
  dark: true,
  colors: {
    background: '#1E2228',
    surface: '#23272f',
    primary: '#048CFC',
    secondary: '#05478A',
    success: '#4BB1A5',
    warning: '#FFA366',
    error: '#fff1f1',
    onPrimary: '#23272f',
    neutral: '#AAB0BB',
    text: '#F8F9FC',
  },
}

export default createVuetify({
  theme: {
    defaultTheme: 'myLight',
    themes: { myLight, myDark },
  },
  icons: {
    defaultSet: 'fa',
    aliases,
    sets: { fa },
  },
  defaults: {
    global: {
      style: {
        fontFamily: 'Inter, system-ui, sans-serif',
      },
    },
  },
})
