import PrimeVue from "primevue/config";
import Ripple from 'primevue/ripple';
import Tooltip from 'primevue/tooltip';
import FocusTrap from 'primevue/focustrap';

import "primevue/resources/themes/aura-light-indigo/theme.css";
import "primeflex/primeflex.css";
import "primeicons/primeicons.css";
import "@/assets/css/app.css";

import { createApp } from "vue";
import { createPinia } from "pinia";
import piniaPluginPersistedstate from 'pinia-plugin-persistedstate';

import App from "./App.vue";

const app = createApp(App);
app.use(createPinia().use(piniaPluginPersistedstate));
app.use(PrimeVue, { ripple: true });
app.use(PrimeVue);

app.directive('focustrap', FocusTrap);
app.directive('ripple', Ripple);
app.directive('tooltip', Tooltip);

app.mount("#app");
