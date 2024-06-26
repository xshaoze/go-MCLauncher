import {createApp} from 'vue'
import App from './App.vue'

import Idux from "./idux";

const app = createApp(App)
app.mount('#app')
app.use(Idux)
