<template>
  <footer class="footer">
    <div class="footer-content">
      <div class="customer-store-pay-follow">
        <h6>Chăm sóc khách hàng</h6>
        <router-link to="#">Trung tâm trợ giúp</router-link>
        <router-link to="#">Happy Tea Blog</router-link>
        <router-link to="#">Hướng dẫn mua hàng</router-link>
        <router-link to="#">Thanh toán</router-link>
        <router-link to="#">Vận chuyển</router-link>
        <router-link to="#">Trung tâm trợ giúp</router-link>
      </div>
      <div class="customer-store-pay-follow">
        <h6>Về cửa hàng</h6>
        <router-link to="#">Giới thiệu về Happy Tea</router-link>
        <router-link to="#">Giấy phép kinh doanh</router-link>
        <router-link to="#">Nguồn góc sản phẩm</router-link>
        <router-link to="#">Flash Sales</router-link>
        <router-link to="#">Đối tác</router-link>
      </div>
      <div class="customer-store-pay-follow">
        <h6>Thanh toán</h6>
        <router-link to="#">Tại cửa hàng</router-link>
        <router-link to="#">Khi nhận hàng</router-link>
        <router-link to="#">Thanh toán trực tuyến</router-link>
        <div class="pay">
          <div class="pay-1">
            <img :src="momo" />
            <img :src="airpay" />
            <img :src="zalopay" />
          </div>
          <div class="pay-1">
            <img :src="mastercard" />
            <img :src="visa" />
            <img :src="cod" />
          </div>
        </div>
      </div>
      <div class="customer-store-pay-follow">
        <h6>Liên hệ</h6>
        <router-link to="#"><i class="fa fa-facebook-square"></i>Facebook</router-link>
        <router-link to="#"><i class="fa fa-instagram"></i>Instagram</router-link>
        <router-link to="#"><i class="fa fa-twitter"></i>Twitter</router-link>
        <router-link to="#"><i class="fa fa-phone-square"></i>01991000</router-link>
      </div>
    </div>
  </footer>
</template>

<script>
  import momo from './../../assets/pay/momo.svg'
  import cod from './../../assets/pay/cod.svg'
  import airpay from './../../assets/pay/airpay.svg'
  import mastercard from './../../assets/pay/mastercard.svg'
  import visa from './../../assets/pay/visa.svg'
  import zalopay from './../../assets/pay/zalopay.png'

  export default {
  data() {
    return {
      momo,
      cod,
      airpay,
      mastercard,
      visa,
      zalopay
    }
  },
}
</script>

<style lang="scss" scoped>
@import './footer.scss';
footer {
  clear: both;
}
</style>

