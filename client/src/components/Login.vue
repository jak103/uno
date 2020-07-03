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
                  <v-text-field label="GAME ID" type="text" v-model="game_id"></v-text-field>
                  <v-text-field label="USERNAME" type="text" v-model="user_name"></v-text-field>
                  <v-btn @click.native="login" color="primary" :to="to">Join Game</v-btn>
                  <v-btn color="primary" @click.native="newGame">Create new game</v-btn>
                  <v-card v-if="status != ''">{{ status }}</v-card>
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
import axios from "axios";
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
      if (this.user_name != "") {
        let res = await axios.post("http://localhost:8080/login/" + this.game_id + "/" + this.user_name);
        if (res.data.valid) {
          this.to = {
            name: "About",
            params: { game_id: this.game_id, valid: res.data.valid, username: this.user_name}
          };
        }
      } else {
        alert("Please enter a username. This will be displayed to other players")
      }
    },
    async newGame() {
      console.log("New game!");
      let res = await axios.get("http://localhost:8080/newgame");
      this.game_id = res.data.payload.game_id;
      this.status = "New game id is: " + this.game_id;
    }
  }
};
</script>