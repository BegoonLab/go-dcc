import Vue from 'vue';
import App from './components/App.vue';
import Train from './components/Train.vue';
import vuetify from './plugins/vuetify'

Vue.component('train', Train)

new Vue({
  vuetify,
  el: '#app',
  render: h => h(App),
});