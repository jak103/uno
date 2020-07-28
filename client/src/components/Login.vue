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
                  <v-text-field test-id="login-game-id" label="GAME ID" type="text" v-model="game_id"></v-text-field>
                  <v-text-field test-id="login-user-name" label="USERNAME" type="text" v-model="user_name"></v-text-field>
                  <v-btn test-id="login-join-game" @click.native="login" color="primary">Join Game</v-btn>
                  <v-btn test-id="login-new-game" color="primary" @click.native="newGame">Create new game</v-btn>
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
      valid_game: false,
      game_id: "",
      user_name: "",
      status: "",
      sim: true // Only true while debugging
    };
  },
  methods: {
    async login() {
      if (this.user_name == "") {
        this.status = "Please enter a username. It will be displayed to other players";
        return;
      }
      if (this.game_id == "") {
        this.status = "Please enter a game id or create a new game.";
        return;
      }

      let res = await unoService.login(this.game_id, this.user_name);
      if (res.data.valid) {
        unoService.setToken(res.data.payload.JWT);
        this.$router.push({
          name: "About",
          params: { game_id: this.game_id, valid: res.data.valid, username: this.user_name}
        });
      }
    },
    async newGame() {
      if (this.user_name == "") {
        this.status = "Please enter a username. It will be displayed to other players";
        return;
      }

      let res = await unoService.newGame(this.user_name);
      if (res.data.valid) {
        unoService.setToken(res.data.payload.JWT);
        this.$router.push({
          name: "About",
          params: { game_id: res.data.payload.game_id, valid: res.data.valid, username: this.user_name}
        });
      }
    }
  }
};
</script>