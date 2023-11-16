import axios from 'axios'
import { reactive } from 'vue'

const state = reactive({
  cart: [],
  quantity: 0,
})
const getters = {
  getCart: state => state,
}

const actions = {
  testthui() {
    console.log('hello world')
  },
  // fetch gio hang
  async fetchCartById({ commit }) {
    const response = await axios.get(`http://localhost:3000/api/cartitems`, {
      withCredentials: true,
    })
    commit('FETCH_CART_BY_ID', response.data)
  },
  // them vao gio hang
  async addCart({ commit }, payload) {
    console.log(payload)
    const response = await axios.post(
      `http://localhost:3000/api/cartitems`,
      {
        product_id: payload.product_id,
        quantity: payload.amount,
      },
      { withCredentials: true },
    )
    if (typeof response === 'string') {
      console.log(response)
    }
    commit('ADD_CART', payload)
  },
  // xoa bot san pham
  async subCart({ commit }, payload) {
    if (payload.amount === 1) {
      const response = await axios.delete(
        `http://localhost:3000/api/cartitems/${payload.cart_id}`,
        { withCredentials: true },
      )
      if (typeof response === 'string') {
        console.log(response)
      }
      commit('REMOVE_CART', payload)
      return
    }
    const response = await axios.put(
      `http://localhost:3000/api/cartitems/${payload.cart_id}`,
      {
        quantity: payload.amount - 1,
      },
      { withCredentials: true },
    )
    if (typeof response === 'string') {
      console.log(response)
    }
    commit('SUB_CART', payload.cart_id)
  },
  // xoa san pham khoi gio hang
  async removeCart({ commit }, payload) {
    const response = await axios.delete(
      `http://localhost:3000/api/cartitems/${payload.cart_id}`,
      { withCredentials: true },
    )
    if (typeof response === 'string') {
      console.log(response)
    }
    commit('REMOVE_CART', payload)
  },
}

const mutations = {
  FETCH_CART_BY_ID: (state, res) => {
    let msg = Object.freeze(res) // ngăn ko cho nó thành proxy object
    let total = 0
    if (typeof msg === 'object') {
      Array.prototype.forEach.call(msg, val => {
        total += val.quantity
      })
      state.cart = msg
    }
    state.quantity = total
  },
  ADD_CART: (state, payload) => {
    Array.prototype.forEach.call(state.cart, item => {
      if (item.product_id === payload.product_id) {
        item.quantity = item.quantity + payload.amount
      }
    })
    state.quantity += payload.amount
  },
  SUB_CART: (state, cart_id) => {
    Array.prototype.forEach.call(state.cart, item => {
      if (item.id === cart_id) {
        item.quantity -= 1
      }
    })
    state.quantity -= 1
  },
  REMOVE_CART: (state, payload) => {
    console.log(payload)
    let updateCart = state.cart.filter(item => item.id !== payload.cart_id)
    state.cart = updateCart
    state.quantity = state.quantity - payload.amount
  },
}

export default {
  state,
  getters,
  actions,
  mutations,
}
