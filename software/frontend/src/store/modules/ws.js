import {defineStore} from "pinia";
import {useControllerStore} from "./controller";

export const useWsStore = defineStore('ws', () => {
    let ws = {};
    let ControllerType = {};

    function connectToWebsocket() {
        const controllerStore = useControllerStore()
        const pbRoot = require('./../../pb/controller.proto');
        ControllerType = pbRoot.lookupType('controller.Controller')
        setupWs()

        setInterval(() => {
            if (controllerStore.isDisconnected === false && ws.readyState !== WebSocket.OPEN) {
                console.log("Connection Error!");
                controllerStore.setDisconnected();
            }
        }, 1000)
    }
    function setupWs() {
        ws = new WebSocket(`ws://${location.hostname}:3000/ws`);
        ws.binaryType = 'arraybuffer' // Important!
        ws.addEventListener('open', (event) => { onWebsocketOpen(event) });
        ws.addEventListener('message', (event) => { handleNewMessage(event) });
        ws.addEventListener('error', (event) => { handleError(event) });
    }
    function handleError(event) {
        const controllerStore = useControllerStore()
        console.log("Connection Error!");
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