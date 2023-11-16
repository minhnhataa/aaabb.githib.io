<template>
  <div class="page-user">
    <router-link to="">Home</router-link>
    >
    <router-link to="#">User</router-link>
    <div class="user-main">
      <MenuUser />
      <div class="user-content">
        <h2>Thông tin tài khoản</h2>
        <form action="">
          <ul>
            <li>
              <p>Họ và tên</p>
              <input type="text" value="Nguyen" />
            </li>
            <li>
              <p>Số điện thoại</p>
              <input type="text" value="01254" />
            </li>
            <li>
              <p>Email</p>
              <input type="text" value="a@gmai.com" />
            </li>
            <li>
              <p>Giới tính</p>
              <input type="radio" id="Nam" value="true" />
              <label for="Nam">Nam</label>
              <input type="radio" id="Nu" />
              <label for="Nu">Nu</label>
            </li>
            <li>
              <p>Ngày sinh</p>
              <input type="date" value="1/1/2021" />
            </li>
          </ul>
          <button type="submit">Cập nhật</button>
        </form>
      </div>
    </div>
  </div>
</template>

<script>
import Cookies from 'universal-cookie'
import MenuUser from '../../components/MenuUser/MenuUser.vue'
export default {
  name: 'User',
  setup() {
    const cookies = new Cookies()
    return {
      cookies,
    }
  },
  components: {
    MenuUser,
  },
}
</script>

<style lang="scss" scoped>
@import './user.scss';
</style>
