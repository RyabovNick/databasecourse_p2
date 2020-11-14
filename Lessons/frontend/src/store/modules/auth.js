// initial state
const state = () => ({
  user: {},
  is_auth: false,
})

// getters
const getters = {}

// actions
const actions = {
  signIn({ commit }) {
    // запрос к /sign_in API
    // shop.getProducts(products => {
    //   commit('setProducts', products)
    // })
  },
}

// mutations
const mutations = {
  setUser(state, user) {
    state.user = user
    state.is_auth = true
  },
}

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations,
}
