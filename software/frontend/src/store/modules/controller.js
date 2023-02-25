import {defineStore} from 'pinia'
import {computed, ref, watch} from 'vue'
import {useWsStore} from "./ws";

export const useControllerStore = defineStore('controller', () => {

    const id = ref("")
    const railwayModules = ref({})
    const locomotives = ref({})
    const started = ref(false)
    const reboot = ref(false)
    const poweroff = ref(false)
    const connected = ref(false)
    const isDisconnected = computed(() => {
        return !connected.value
    })
    const getEnabledRailwayModules = computed(() => {
        let enabledRailwayModules = {}
        Object.keys(railwayModules.value).forEach(key => {
            if (railwayModules.value[key] && railwayModules.value[key].enabled) {
                getEnabledRailwayModules[key] = railwayModules.value[key]
            }
        })
        return enabledRailwayModules
    })
    const getEnabledLocomotives = computed(() => {
        let enabledLocomotives = {}
        Object.keys(locomotives.value).forEach(key => {
            if (locomotives.value[key] && locomotives.value[key].enabled) {
                enabledLocomotives[key] = locomotives.value[key]
            }
        })
        return enabledLocomotives
    })

    function getState() {
        return {
            'id': id.value,
            'locomotives': locomotives.value,
            'railwayModules': railwayModules.value,
            'started': started.value,
            'reboot': reboot.value,
            'poweroff': poweroff.value,
        }
    }

    function newMessage(message) {
        locomotives.value = message.locomotives
        started.value = message.started
        railwayModules.value = message.railwayModules
        id.value = message.id
    }

    function sendDataToServer() {
        started.value = true;
        const wsStore = useWsStore()
        wsStore.sendMessage(getState())
    }

    function stopAll() {
        Object.keys(locomotives.value).forEach(key => {
            locomotives.value[key].speed = 0;
        })
        started.value = false;
        const wsStore = useWsStore()
        wsStore.sendMessage(getState())
    }

    function setReboot() {
        reboot.value = true;
        const wsStore = useWsStore()
        wsStore.sendMessage(getState())
    }

    function setPowerOff() {
        poweroff.value = true;
        const wsStore = useWsStore()
        wsStore.sendMessage(getState())
    }

    function setConnected() {
        connected.value = true;
    }

    function setDisconnected() {
        connected.value = false;
    }

    return {
        railwayModules,
        isDisconnected,
        locomotives,
        started,
        reboot,
        poweroff,
        connected,
        getEnabledRailwayModules,
        getEnabledLocomotives,
        newMessage,
        stopAll,
        setReboot,
        setPowerOff,
        setConnected,
        setDisconnected,
        sendDataToServer
    }
})