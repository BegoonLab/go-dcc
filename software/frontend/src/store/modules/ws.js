import {defineStore} from "pinia";
import {useControllerStore} from "./controller";
import {demoData} from "./demo";

export const useWsStore = defineStore('ws', () => {
    let ws = {};
    let ControllerType = {};
    let skipVibro = true;

    function connectToWebsocket() {
        const controllerStore = useControllerStore()
        const pbRoot = require('./../../pb/controller.proto');
        ControllerType = pbRoot.lookupType('controller.Controller')

        if (process.env.NODE_ENV !== 'demo') {
            setupWs()

            setInterval(() => {
                if (controllerStore.isDisconnected === false && ws.readyState !== WebSocket.OPEN) {
                    console.log("Connection Error!");
                    controllerStore.setDisconnected();
                }
            }, 1000)
        } else {
            // Add Demo data
            handleNewMessage({})
        }
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
        controllerStore.setDisconnected();
    }
    function onWebsocketOpen() {
        const controllerStore = useControllerStore()
        console.log("connected to WS!");
        controllerStore.setConnected()
    }
    function handleNewMessage(event) {
        const controllerStore = useControllerStore()
        let decoded;
        if (process.env.NODE_ENV !== 'demo') {
            decoded = ControllerType.decode(new Uint8Array(event.data));
        } else {
            console.log("Loading Demo data...")
            decoded = ControllerType.decode(demoData);
        }

        let data = ControllerType.toObject(decoded, {
            arrays: true,
            defaults: true,
        });
        controllerStore.newMessage(data)
        controllerStore.setConnected();
        if (!skipVibro) {
            window.navigator.vibrate(10);
        }
        skipVibro = false;
    }
    function sendMessage(msg) {
        const controllerStore = useControllerStore()
        const buffer = ControllerType.encode(msg).finish();
        if (process.env.NODE_ENV !== 'demo') {
            if (ws.readyState !== WebSocket.OPEN) {
                console.log("Connection Error!");
                controllerStore.setDisconnected()
            }
            ws.send(buffer);
        }
        window.navigator.vibrate(20);
        skipVibro = true;
    }

    return {
        connectToWebsocket,
        sendMessage
    }
})