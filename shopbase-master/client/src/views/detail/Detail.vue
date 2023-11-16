<template>
  <div v-if="product" class="cart">
    <div class="main-cart">
      <div class="path">
        <router-link to="/">Trang chủ</router-link>
        <p>/</p>
        <span>{{ product.name }}</span>
      </div>
      <div class="add-to-cart">
        <div class="img-sp">
          <img :src="sp1" />
        </div>
        <div class="detail-product">
          <h2>{{ product.name }}</h2>
          <div class="price">
            <span>{{ product.price }} 000 đ</span>
            <p>Ở đâu rẻ hơn, Happy Tea đến dẹp tiệm</p>
          </div>
          <span>Chưa có đánh giá</span>
          <div class="brand">
            <table>
              <tr>
                <td>Thương hiệu</td>
                <td>Sharetea</td>
              </tr>
              <tr>
                <td>Popping</td>
                <td>Cà phê đen</td>
              </tr>
              <tr>
                <td>Vận chuyển</td>
                <td>Có</td>
              </tr>
              <tr>
                <td>Đã bán</td>
                <td>113</td>
              </tr>
            </table>
          </div>
          <div class="to-add">
            <i
              class="fa fa-caret-left"
              :disabled="quantity <= 1"
              @click="sub_quantity"
            ></i>
            <span>{{ quantity }}</span>
            <i class="fa fa-caret-right" @click="add_quantity"></i>
          </div>

          <button @click="create_order">Thêm vào giỏ hàng</button>
        </div>
      </div>
      <div class="related">
        <h4>Sản phẩm liên quan</h4>
        <div class="related-sp">
          <img :src="related" />
          <p>Không có sản phẩm liên quan</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios'
import { mapActions } from 'vuex'
import Cookies from 'universal-cookie'

import sp1 from './../../assets/product/sp1.jpg'
import related from './../../assets/related/empty.svg'

// <span>{{ product.id }}</span> id sp
// <span>{{ product.desc }}</span> chi tiet sp

export default {
  name: 'Detail',
  setup() {
    const cookies = new Cookies()
    return {
      cookies,
    }
  },
  data() {
    return {
      product: null,
      quantity: 1,
      sp1,
      related,
    }
  },
  mounted() {
    this.$nextTick(async function() {
      const id = this.$route.params.id
      const { data } = await axios.get(
        `http://localhost:3000/api/products/${id}`,
      )
      this.product = data
    })
  },
  methods: {
    ...mapActions(['addCart']),
    sub_quantity: function() {
      if (this.quantity > 1) {
        this.quantity = this.quantity - 1
      }
    },
    add_quantity: function() {
      this.quantity = this.quantity + 1
    },
    create_order: async function() {
      let productId = this.$route.params.id
      let payload = { product_id: productId, amount: this.quantity }
      this.addCart(payload)
    },
  },
}
</script>
<style lang="scss" scoped>
@import './detail.scss';
</style>
