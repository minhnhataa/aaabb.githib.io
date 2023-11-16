import { createStore } from 'vuex'
import cart from './modules/cart'
import order from './modules/order'
import history from './modules/history'

export default createStore({
  state: {},
  mutations: {},
  actions: {},
  modules: {
    cart,
    order,
    history,
  },
})
