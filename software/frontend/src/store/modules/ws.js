import {defineStore} from "pinia";
import {useControllerStore} from "./controller";

export const useWsStore = defineStore('ws', () => {
    let ws = {};
    let lastMsgTime = Date.now();
    let ControllerType = {};

    function connectToWebsocket() {
        const controllerStore = useControllerStore()
        const pbRoot = require('./../../pb/controller.proto');
        ControllerType = pbRoot.lookupType('controller.Controller')
        ws = new WebSocket(`ws://${location.hostname}:3000/ws`);
        ws.binaryType = 'arraybuffer' // Important!
        ws.addEventListener('open', (event) => { onWebsocketOpen(event) });
        ws.addEventListener('message', (event) => { handleNewMessage(event) });
        ws.addEventListener('error', (event) => { handleError(event) });

        setInterval(() => {
            if (Date.now() - lastMsgTime > 5000) {
                console.log("Connection Error!");
                controllerStore.setDisconnected();
            }
        }, 1000)
    }
    function handleError(event) {
        const controllerStore = useControllerStore()
        console.log("Connection Error!");
        console.log(event)
        controllerStore.setDisconnected()
    }
    function onWebsocketOpen(store) {
        const controllerStore = useControllerStore()
        console.log("connected to WS!");
        controllerStore.setConnected()
    }
    function handleNewMessage(event) {
        const controllerStore = useControllerStore()
        let decoded = ControllerType.decode(new Uint8Array(event.data));
        let data = ControllerType.toObject(decoded, {
            arrays: true,
            defaults: true,
        });
        controllerStore.newMessage(data)
        lastMsgTime = Date.now();
        controllerStore.setConnected()
    }
    function sendMessage(msg) {
        const controllerStore = useControllerStore()
        if (ws.readyState !== WebSocket.OPEN) {
            console.log("Connection Error!");
            controllerStore.setDisconnected()
        }

        const buffer = ControllerType.encode(msg).finish();
        ws.send(buffer);
    }

    return {
        connectToWebsocket,
        sendMessage
    }
})