import Vue from 'vue';
import App from './components/App.vue';
import Train from './components/Train.vue';
import vuetify from './plugins/vuetify'
import Vuex from 'vuex'
import store from './store'

Vue.use(Vuex)
Vue.component('train', Train)

new Vue({
  vuetify,
  store,
  el: '#app',
  render: h => h(App),
});