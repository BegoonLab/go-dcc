// initial state
const state = () => ({
  locomotives: [],
  started: false
})

// getters
const getters = {
  locomotives: (state, getters, rootState) => {
    return state.locomotives
  },

  cartTotalPrice: (state, getters) => {
    // return getters.cartProducts.reduce((total, product) => {
    //   return total + product.price * product.quantity
    // }, 0)
  }
}

// actions
const actions = {
  newMessage ({ commit, state }, message) {
    commit('setData', {data: message})
  },

  setLocomotiveState ({ commit, state }, {name, value, where}) {
    commit('setLocomotive', {name: name, value: value, where: where})
    // TODO: Send message here
  },

  addProductToCart ({ state, commit }, product) {
//     commit('setCheckoutStatus', null)
//     if (product.inventory > 0) {
//       const cartItem = state.items.find(item => item.id === product.id)
//       if (!cartItem) {
//         commit('pushProductToCart', { id: product.id })
//       } else {
//         commit('incrementItemQuantity', cartItem)
//       }
//       // remove 1 item from stock
//       commit('products/decrementProductInventory', { id: product.id }, { root: true })
//     }
  }
}

// mutations
const mutations = {
//   pushProductToCart (state, { id }) {
//     state.items.push({
//       id,
//       quantity: 1
//     })
//   },

//   incrementItemQuantity (state, { id }) {
//     const cartItem = state.items.find(item => item.id === id)
//     cartItem.quantity++
//   },

  setData (state, { data }) {
    state.locomotives = data.locomotives
    state.started = data.started
  },

  setLocomotive(state, { name, value, where }) {
    state.locomotives[name][where] = value
  }

//   setCheckoutStatus (state, status) {
//     state.checkoutStatus = status
//   }
}

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations
}