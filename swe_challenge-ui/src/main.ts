import { createApp } from 'vue'
import App from './App.vue'

// Pinia
import { createPinia } from 'pinia'
const pinia = createPinia()

// Importa los estilos de Tailwind (asegúrate de que tu tailwind.config.js está bien)
import './assets/main.css'

// PrimeVue y estilos
import PrimeVue from 'primevue/config';
import 'primevue/resources/themes/lara-light-blue/theme.css'; // Tema moderno
import 'primevue/resources/primevue.min.css'; 
import 'primeicons/primeicons.css';
import 'primeflex/primeflex.css';

const app = createApp(App)
app.use(pinia)
app.use(PrimeVue);
app.mount('#app')
