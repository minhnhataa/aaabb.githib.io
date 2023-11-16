import axios from 'axios'
import { reactive } from 'vue'

const state = reactive({
  ordered: [],
  total: 0,
  payment_id: 1,
})

const getters = {
  getOrder: state => state.total,
}

const actions = {
  updateOrder({ commit }, payload) {
    commit('UPDATE_ORDER', payload)
  },
  async createOrder({ dispatch }) {
    let products = []
    let cartId = []
    state.ordered.forEach(item => {
      let product = {
        product_id: item.cart.product_id,
        quantity: item.cart.quantity,
        price: parseInt(item.cart.price),
      }
      products.push(product)
      cartId.push(item.cart.id)
    })
    if (state.ordered.length === 0) {
      alert('bạn chưa chọn sản phẩm để mua')
    } else {
      const response = await axios.post(
        `http://localhost:3000/api/orders`,
        {
          total: state.total,
          payment_id: state.payment_id,
          products: products,
        },
        { withCredentials: true },
      )
      console.log(response)
      state.total = 0
      cartId.forEach(item => {
        let payload = { cart_id: item }
        dispatch('removeCart', payload)
      })
    }
  },
}

const mutations = {
  UPDATE_ORDER: (state, payload) => {
    state.ordered = payload
    let tmp = 0
    state.ordered.forEach(item => {
      tmp += item.cart.price * item.cart.quantity
    })
    state.total = tmp
  },
}

export default {
  state,
  getters,
  actions,
  mutations,
}
