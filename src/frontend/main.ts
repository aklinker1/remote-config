import 'vue-global-api';
import './tailwind.css';
import './styles.css';
import { createApp } from 'vue';
import { router } from './router';
import App from './App.vue';

createApp(App).use(router).mount('#app');
