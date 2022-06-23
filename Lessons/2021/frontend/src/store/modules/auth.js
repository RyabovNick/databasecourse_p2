// import axios from 'axios'
import jwtDecode from 'jwt-decode'

// initial state
const state = () => ({
  user: {},
  isAuth: false,
})

// getters
const getters = {}

// actions
const actions = {
  setAuth({ commit }, token) {
    this._vm.$axios.defaults.headers.common['Authorization'] = token
    const user = jwtDecode(token)
    commit('setUser', user)
  },

  async signIn({ commit }, credentials) {
    const resp = await this._vm.$axios.post('/sign_in', credentials)
    const token = resp.data.token
    // добавлю к axios header по умолчанию
    // чтобы все запросы к бэку отправлялись с токеном
    this._vm.$axios.defaults.headers.common['Authorization'] = `Bearer ${token}`
    localStorage.setItem('token', `Bearer ${token}`)

    const user = jwtDecode(token)
    commit('setUser', user)
  },

  logout({ commit }) {
    delete this._vm.$axios.defaults.headers.common['Authorization']
    localStorage.removeItem('token')
    commit('deleteUser')
  },
}

// mutations
const mutations = {
  setUser(state, user) {
    state.user = user
    state.isAuth = true
  },

  deleteUser(state) {
    state.user = {}
    state.isAuth = false
  },
}

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations,
}
