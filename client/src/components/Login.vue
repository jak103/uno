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
                  <v-btn test-id="login-join-game" @click="login" color="primary">Join Game</v-btn>
                  <v-btn test-id="login-new-game" color="primary" @click.native="newGame">Create new game</v-btn>
                  <v-card test-id="login-status" v-if="status != ''">{{ status }}</v-card>
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
      game_id: null,
      user_name: "",
      to: {},
      status: "",
      sim: true // Only true while debugging
    };
  },
  methods: {
    async login() {
        console.log("Running login");
      if (this.user_name != "") {
        console.log("Hitting API");
        let res = await unoService.login(this.game_id, this.user_name);
        if (res.data.valid) {
          console.log("Reponse was valid");
          this.$router.push({name: "About", params:{ game_id: this.game_id, valid: res.data.valid, username: this.user_name}});
          /*
          this.to = {
            name: "About",
            params: { game_id: this.game_id, valid: res.data.valid, username: this.user_name}
          };
          */
        }
      } else {
          // TODO: Move this to a snack
        alert("Please enter a username. This will be displayed to other players")
      }
    },
    async newGame() {
      let res = await unoService.newGame();
      this.game_id = res.data.payload.game_id;
      this.status = "New game id is: " + this.game_id;
    }
  }
};
</script>