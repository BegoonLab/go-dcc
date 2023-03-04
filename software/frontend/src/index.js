import {createApp} from 'vue'
import {createPinia} from 'pinia'
import App from './components/App.vue'
import Train from './components/Train.vue'
import RailwayModule from './components/RailwayModule.vue'

// Vuetify
import '@mdi/font/css/materialdesignicons.css'
import 'material-design-icons-iconfont/dist/material-design-icons.css'
import '@fortawesome/fontawesome-free/css/all.css'
import 'vuetify/styles'
import './styles/styles.sass'
import {createVuetify} from 'vuetify'
import * as components from 'vuetify/components'
import * as directives from 'vuetify/directives'
import {aliases, mdi} from 'vuetify/iconsets/mdi'
import {fa} from 'vuetify/iconsets/fa'
import {md} from 'vuetify/iconsets/md'
import {mdiAccount} from '@mdi/js'

const vuetify = createVuetify({
    components,
    directives,
    icons: {
        defaultSet: 'mdi',
        aliases: {
            ...aliases,
            account: mdiAccount,
        },
        sets: {
            md,
            fa,
            mdi,
        },
    },
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