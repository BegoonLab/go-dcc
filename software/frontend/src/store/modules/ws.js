import {defineStore} from "pinia";
import {useControllerStore} from "./controller";

export const useWsStore = defineStore('ws', () => {
    let ws = {};
    let lastMsgTime = Date.now();

    function connectToWebsocket() {
        ws = new WebSocket(`ws://${location.hostname}:3000/ws`);
        ws.addEventListener('open', (event) => { onWebsocketOpen(event) });
        ws.addEventListener('message', (event) => { handleNewMessage(event) });
        ws.addEventListener('error', (event) => { handleError(event) });

        setInterval(() => {
            if (Date.now() - lastMsgTime > 5000) {
                const controllerStore = useControllerStore()
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
        let data = event.data;
        data = data.split(/\r?\n/);
        for (let i = 0; i < data.length; i++) {
            let msg = JSON.parse(data[i]);
            controllerStore.newMessage(msg)
        }
        lastMsgTime = Date.now();
        controllerStore.setConnected()
    }
    function sendMessage(msg) {
        const controllerStore = useControllerStore()
        if (ws.readyState !== WebSocket.OPEN) {
            console.log("Connection Error!");
            controllerStore.setDisconnected()
        }
        ws.send(JSON.stringify(msg));
    }

    return {
        connectToWebsocket,
        sendMessage
    }
})