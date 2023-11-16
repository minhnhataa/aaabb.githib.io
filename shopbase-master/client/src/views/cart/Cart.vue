<template>
  <div class="main">
    <div class="container">
      <div class="product-content">
        <div class="cart">
          <h2>Giỏ hàng ({{ getCart.quantity }})</h2>
        </div>
        <div class="product">
          <table>
            <tr class="title">
              <td>Mua</td>
              <td class="img">Hình ảnh</td>
              <td>Tên SP</td>
              <td>Đơn giá</td>
              <td>Số lượng</td>
              <td>thành tiền</td>
              <td>Xóa</td>
            </tr>
            <tr v-for="cart in getCart.cart" :key="cart.id" class="product">
              <td>
                <input
                  type="checkbox"
                  name="choose"
                  v-model="ordered"
                  :value="{ cart }"
                  @change="checkboxHandler"
                />
              </td>
              <td class="img"><img :src="sp1" alt="Hinh anh" /></td>
              <td>{{ cart.name }}</td>
              <td>{{ cart.price }}</td>
              <td>
                <i class="fa fa-caret-left" @click="sub_quantity(cart.id, cart.quantity)"></i>
                <span>{{ cart.quantity }}</span>
                <i class="fa fa-caret-right" @click="add_quantity(cart.product_id)"></i>
              </td>
              <td>{{ cart.quantity * cart.price }}</td>
              <td><button @click="remove_quantity(cart.id, cart.quantity)">xóa</button></td>
            </tr>
          </table>
        </div>
      </div>
      <div class="price">
        <div class="pro">
          <div class="km">
            <span>Khuyến mãi</span>
            <router-link to="#">
            <i class="fa fa-pencil-square-o"></i>
            Nhập khuyến mãi
            </router-link>
          </div>

        </div>
        <div class="bill">
          <div class="bill-content">
            <table>
              <tr>
                <td>Tạm tính :</td>
                <td>{{ getOrder }}đ</td>
              </tr>
              <tr>
                <td>Giảm giá :</td>
                <td>0đ</td>
              </tr>
              <tr class="sum">
                <td>Tổng cộng :</td>
                <td>{{ getOrder }}đ</td>
              </tr>
            </table>
          </div>
        </div>
        <div class="order">
          <button @click="orderHandler">Đặt hàng</button>
        </div>

      </div>
    </div>
  </div>
</template>

<script>
import { mapGetters, mapActions } from 'vuex'
import sp1 from './../../assets/product/sp1.jpg'

export default {
  name: 'Cart',
  data() {
    return {
      ordered: [],
      sp1
    }
  },
  computed: {
    ...mapGetters(['getCart', 'getOrder']),
  },
  created() {
    this.fetchCartById()
  },
  methods: {
    ...mapActions([
      'fetchCartById',
      'addCart',
      'subCart',
      'removeCart',
      'updateOrder',
      'createOrder',
    ]),
    sub_quantity(cart_id, quantity) {
      let payload = { cart_id: cart_id, amount: quantity }
      this.subCart(payload)
    },
    add_quantity(product_id) {
      let payload = { product_id: product_id, amount: 1 }
      this.addCart(payload)
    },
    remove_quantity(cart_id, quantity) {
      let payload = { cart_id: cart_id, amount: quantity }
      this.removeCart(payload)
    },
    checkboxHandler() {
      this.updateOrder(this.ordered)
    },
    orderHandler() {
      this.createOrder()
    },
  },
}
</script>
<style lang="scss" scoped>
@import './cart.scss';
</style>
