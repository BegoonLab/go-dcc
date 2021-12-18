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
let lastMsgTime = Date.now();

const wsHelper = {
  connectToWebsocket() {
    ws = new WebSocket(`ws://${location.hostname}:3000/ws`);
    ws.addEventListener('open', (event) => { this.onWebsocketOpen(event) });
    ws.addEventListener('message', (event) => { this.handleNewMessage(event) });
    ws.addEventListener('error', (event) => { this.handleError(event) });
  },
  handleError(event) {
    console.log("Connection Error!");
    console.log(event)
    store.dispatch('controller/disconnected');
  },
  onWebsocketOpen() {
    console.log("connected to WS!");
    store.dispatch('controller/connected');
  },
  handleNewMessage(event) {
    let data = event.data;
    data = data.split(/\r?\n/);

    for (let i = 0; i < data.length; i++) {
      let msg = JSON.parse(data[i]);
      store.dispatch('controller/newMessage', msg);
    }

    lastMsgTime = Date.now();

    store.dispatch('controller/connected');
  },
  sendMessage(state) {
    if (ws.readyState !== WebSocket.OPEN) {
      console.log("Connection Error!");
      store.dispatch('controller/disconnected');
    }
    ws.send(JSON.stringify(state));
  }
}

wsHelper.connectToWebsocket();

setInterval(() => {
  if (Date.now() - lastMsgTime > 5000) {
    console.log("Connection Error!");
    store.dispatch('controller/disconnected');
  }
}, 1000)

store.subscribe((mutation, state) => {
  if (mutation.type === 'controller/setLocomotive' ||
    mutation.type === 'controller/stopAll' ||
    mutation.type === 'controller/reboot' ||
    mutation.type === 'controller/poweroff') {
    wsHelper.sendMessage(state.controller)
  }
})

export default store