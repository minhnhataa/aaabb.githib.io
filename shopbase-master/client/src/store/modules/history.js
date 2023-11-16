import { reactive } from 'vue'
import axios from 'axios'

const state = reactive({
  history: [],
})

const getters = {
  getHistory: state => state.history,
}

const actions = {
  async fetchHistory({ commit }) {
    const response = await axios.get(`http://localhost:3000/api/orders`, {
      withCredentials: true,
    })
    console.log(response)
    commit('ADD_HISTORY', response.data)
  },
}

const mutations = {
  ADD_HISTORY: (state, payload) => {
    state.history = payload
  },
}

export default {
  state,
  getters,
  actions,
  mutations,
}
