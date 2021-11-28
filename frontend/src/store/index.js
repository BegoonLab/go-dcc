import Vue from 'vue'
import Vuex from 'vuex'
import controller from './modules/controller'

Vue.use(Vuex)

const debug = process.env.NODE_ENV !== 'production'

const store = new Vuex.Store({
  modules: {
    controller,
  },
  strict: debug,
  plugins: debug ? [] : []
})

let ws = {};

const wsHelper = {
  connectToWebsocket() {
    ws = new WebSocket(`ws://${location.hostname}:3000/ws`);
    ws.addEventListener('open', (event) => { this.onWebsocketOpen(event) });
    ws.addEventListener('message', (event) => { this.handleNewMessage(event) });
  },
  onWebsocketOpen() {
    console.log("connected to WS!");
  },
  handleNewMessage(event) {
    let data = event.data;
    data = data.split(/\r?\n/);

    for (let i = 0; i < data.length; i++) {
      let msg = JSON.parse(data[i]);
      store.dispatch('controller/newMessage', msg);
    }
  },
  sendMessage(state) {
    ws.send(JSON.stringify(state));
  }
}

wsHelper.connectToWebsocket();

store.subscribe((mutation, state) => {
  if (mutation.type === 'controller/setLocomotive' || mutation.type === 'controller/stopAll') {
    wsHelper.sendMessage(state.controller)
  }
})

export default store