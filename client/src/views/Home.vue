<template>
  <div class="home">
    <!-- <HelloWorld msg="Welcome to Your Vue.js App"/> -->
    <Login/>
  </div>
</template>

<script>
// @ is an alias to /src
import Login from '@/components/Login'
export default {
  name: 'Home',
  components: {
    Login
  }
}
</script>
