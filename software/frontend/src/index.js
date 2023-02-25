import {createApp} from 'vue'
import {createPinia} from 'pinia'
import App from './components/App.vue'
import Train from './components/Train.vue'
import RailwayModule from './components/RailwayModule.vue'

// Vuetify
import 'vuetify/styles'
import './styles/styles.sass'
import {createVuetify} from 'vuetify'
import * as components from 'vuetify/components'
import * as directives from 'vuetify/directives'

const vuetify = createVuetify({
    components,
    directives,
})

// Pinia
const pinia = createPinia()

// Create Vue App
const app = createApp(App)

app.use(vuetify)
app.use(pinia)

app.component('train', Train)
app.component('railway-module', RailwayModule)

app.mount('#app')