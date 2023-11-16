<template>
  <div style="width: 500px; margin:200px auto; padding:auto;">
    <div id="id01">
      <h1
        style="width:100%; text-align: center; background:#214559; color: white;"
      >
        Log in
      </h1>
      <form @submit.prevent="submit">
        <input
          v-model="email"
          type="text"
          placeholder="example@gmail.com"
          required
        />
        <input
          v-model="password"
          type="password"
          placeholder="password"
          required
        />
        <span>{{ msg }}</span>
        <ul>
          <li><button type="reset">Reset</button></li>
          <li><button type="submit">Login</button></li>
        </ul>
        <a
          href="#"
          onclick="{document.getElementById('id01').style.display='none'
        document.getElementById('id02').style.display='block'}"
          style="width:auto;"
        >
          <p>Sign in</p>
        </a>
      </form>
    </div>

    <div id="id02" style="display:none">
      <h1
        style="width:100%; text-align: center; background:#214559; color: white;"
      >
        Sign in
      </h1>
      <form @submit.prevent="submit">
        <input
          v-model="email"
          type="text"
          placeholder="example@gmail.com"
          required
        />
        <input
          v-model="password"
          type="password"
          placeholder="password"
          required
        />
        <span>{{ msg }}</span>
        <ul>
          <li><button type="reset">Reset</button></li>
          <li><button type="submit">Sign in</button></li>
        </ul>
        <a
          href="#"
          onclick="{document.getElementById('id01').style.display='block'
        document.getElementById('id02').style.display='none'}"
          style="width:auto;"
        >
          <p>Log in</p>
        </a>
      </form>
    </div>
  </div>
</template>
<script>
import axios from 'axios'
import Cookies from 'universal-cookie'
import { mapActions } from 'vuex'

export default {
  name: 'Auth',
  setup() {
    const cookies = new Cookies()
    return {
      cookies,
    }
  },
  data() {
    return {
      email: '',
      password: '',
      msg: '',
    }
  },
  methods: {
    ...mapActions(['fetchCartById']),
    submit: async function() {
      const d = JSON.stringify({
        email: this.email,
        password: this.password,
      })
      const { data } = await axios.post('http://localhost:3000/api/login', d)
      this.msg = data.msg
      this.cookies.set(
        'user_id',
        data.user_id,
        { path: '/' },
        { expires: '1h' },
      )
      if (data.msg === 'success') await this.$router.push('/')
      this.fetchCartById()
    },
  },
}
</script>

<style lang="scss" scoped>
* {
  box-sizing: border-box;
  margin: 0;
  padding: 0;
  font-family: 'Courier New', Courier, monospace;
}
a {
  text-decoration: none;
}
li {
  display: block;
  float: left;
  width: 50%;
}
.center {
  display: flex;
  justify-content: center;
  align-items: center;
}
form {
  border: 3px solid #f1f1f1;
}
input[type='text'],
input[type='password'] {
  width: 100%;
  padding: 12px 10px;
  margin: 8px 0;
  display: inline-block;
  box-sizing: border-box;
}
button {
  padding: 0;
  margin: 0;
  border: 1px solid #ccc;
  box-sizing: border-box;
  background-color: #214559;
  color: white;
  font-size: 14px;
  width: 100%;
  height: 40px;
}
button:hover {
  background-color: #6d91a5;
}
</style>
