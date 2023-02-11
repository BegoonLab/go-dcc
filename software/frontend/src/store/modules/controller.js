// initial state
const state = () => ({
    locomotives: {},
    started: false,
    reboot: false,
    poweroff: false,
    connected: false,
})

// getters
const getters = {
    getLocomotives: (state, getters, rootState) => {
        return state.locomotives ? state.locomotives : {}
    },
    getEnabledLocomotives: (state, getters, rootState) => {
        let enabledLocomotives = {}
        Object.keys(state.locomotives).forEach(key => {
            if (state.locomotives[key].enabled) {
                enabledLocomotives[key] = state.locomotives[key]
            }
        })
        return enabledLocomotives
    },
}

// actions
const actions = {
    newMessage({ commit, state }, message) {
        commit('setData', { data: message })
    },

    setLocomotiveState({ commit, state }, { name, value, where }) {
        commit('setLocomotive', { name: name, value: value, where: where })
    },

    stopAll({ state, commit }) {
        commit('stopAll')
    },

    reboot({ state, commit }) {
        commit('reboot')
    },

    poweroff({ state, commit }) {
        commit('poweroff')
    },

    connected({ state, commit }) {
        commit('connected')
    },

    disconnected({ state, commit }) {
        commit('disconnected')
    }
}

// mutations
const mutations = {
    setData(state, { data }) {
        state.locomotives = data.locomotives
        state.started = data.started
    },

    setLocomotive(state, { name, value, where }) {
        state.locomotives[name][where] = value
        state.started = true;
    },

    stopAll(state) {
        Object.keys(state.locomotives).forEach(key => {
            state.locomotives[key].speed = 0;
        })
        state.started = false;
    },

    reboot(state) {
        state.reboot = true;
    },

    poweroff(state) {
        state.poweroff = true;
    },

    connected(state) {
        state.connected = true;
    },

    disconnected(state) {
        state.connected = false;
    },

}

export default {
    namespaced: true,
    state,
    getters,
    actions,
    mutations
}