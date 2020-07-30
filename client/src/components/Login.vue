<template>
  <v-app id="inspire">
    <v-content>
      
    </v-content>
  </v-app>
</template>

<script>
import unoService from "../services/unoService";
export default {
  name: "LoginPage",
  data: () => {
    return {
      user_name: "",
      status: "",
    };
  },
  methods: {
    async login() {
<<<<<<< HEAD
      if (this.user_name != "") {
        let res = await unoService.login(this.game_id, this.user_name);
        if (res.data.valid) {
          this.$router.push({name: "About", params:{ game_id: this.game_id, valid: res.data.valid, username: this.user_name}});
        }
      } else {
          // TODO: Move this to a snack
        alert("Please enter a username. This will be displayed to other players")
=======
      // Perform client-side validation
      if (this.user_name == "") {
        this.status = "Please enter a username. It will be displayed to other players";
        return;
>>>>>>> origin/jwt
      }

      // Attempt to login
      let res = await unoService.login(this.user_name);
      if (!res.data.valid) {
        return;
      }

      // Set JWT token
      let token = res.data.payload.JWT;
      window.localStorage.setItem("token", token);
      unoService.setToken(token);

      // Send client to the "Lobby" object
      this.$router.push({
        name: "Lobby",
        params: {username: this.user_name}
      });
    },
    created() {
      // Check again if the user is logged in.
      // Other check is on the app start
      let token = window.localStorage.getItem("token");
      if (token != null) {
        unoService.setToken(token);
      }

      // Send client to the "Lobby" object
      this.$router.push({
        name: "Lobby",
        params: {username: this.user_name}
      });
    }
  }
};
</script>

<style scoped>
.v-btn {
  margin: 5px 10px 20px 10px;
}
</style>