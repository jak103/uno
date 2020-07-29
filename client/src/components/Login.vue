<template>
  <v-app id="inspire">
    <v-content>
      <v-container class="fill-height" fluid>
        <v-row align="center" justify="center">
          <v-col cols="12" sm="8" md="4">
            <v-card class="elevation-12">
              <v-toolbar color="primary" dark flat>
                <v-toolbar-title>Uno Login</v-toolbar-title>
                <v-spacer></v-spacer>
                <v-tooltip bottom></v-tooltip>
              </v-toolbar>
              <v-card-text>
                <v-form>
                  <v-text-field test-id="login-user-name" label="USERNAME" type="text" v-model="user_name"></v-text-field>
                  <v-btn test-id="login-join-lobby" @click.native="login" color="primary">Join Lobby</v-btn>
                  <v-card test-id="login-status" color="secondary" v-if="status != ''">{{ status }}</v-card>
                </v-form>
              </v-card-text>
            </v-card>
          </v-col>
        </v-row>
      </v-container>
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
      // Perform client-side validation
      if (this.user_name == "") {
        this.status = "Please enter a username. It will be displayed to other players";
        return;
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