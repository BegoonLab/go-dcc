// initial state
const state = () => ({
  locomotives: [],
  started: false
})

// getters
const getters = {
  getLocomotives: (state, getters, rootState) => {
    return state.locomotives
  },
}

// actions
const actions = {
  newMessage ({ commit, state }, message) {
    commit('setData', {data: message})
  },

  setLocomotiveState ({ commit, state }, {name, value, where}) {
    commit('setLocomotive', {name: name, value: value, where: where})
  },

  stopAll ({state, commit }) {
    commit('stopAll')
  }
}

// mutations
const mutations = {
  setData (state, { data }) {
    state.locomotives = data.locomotives
    state.started = data.started
  },

  setLocomotive(state, { name, value, where }) {
    state.locomotives[name][where] = value
  },

  stopAll(state) {
    Object.keys(state.locomotives).forEach(key => {
      state.locomotives[key].speed = 0;
    })
  }
}

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations
}