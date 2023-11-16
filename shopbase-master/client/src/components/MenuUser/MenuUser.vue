<template>
  <div class="user-menu">
    <div class="user-name">
      <img src="" alt="" />
      <p>Tên tài khoản</p>
    </div>
    <ul>
      <li id="TTTK">Thông tin tài khoản</li>
      <li class="LSMH">Lịch sử mua hàng</li>
      <li id="SDC">Sổ địa chỉ của bạn</li>
    </ul>
  </div>
</template>

<script>
export default {
  name: 'MenuUser',
  setup() {},
}
</script>

<style lang="scss" scoped>
.user-menu {
  width: 30%;
  height: 100%;
  ul {
    list-style-type: none;
  }
  .user-name {
    width: 100%;
    img {
      width: 5rem;
      height: 5rem;
      border-radius: 5rem;
      border: 1px solid red;
    }
    text-align: center;
  }
  li {
    height: 4rem;
    line-height: 5rem;
    text-align: center;
  }
}
</style>
