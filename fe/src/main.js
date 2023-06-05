import { createApp } from 'vue'
import App from './App.vue'
import router from './router'

// Vuetify
import 'vuetify/styles'
import { createVuetify } from 'vuetify'
import * as components from 'vuetify/components'
import * as directives from 'vuetify/directives'

const vuetify = createVuetify({
  components,
  directives,
})

import eventBus from './eventbus'

import VueSimpleWebsocket from "vue-simple-websocket";


let app = createApp(App)
app.use(VueSimpleWebsocket, "ws://localhost:3000/ws",{
  reconnectEnabled: true,
  reconnectInterval: 1000,
})
app.use(router)
app.use(vuetify)
app.mount('#app')
