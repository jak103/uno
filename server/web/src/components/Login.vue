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
      to: {},
      status: ""
    };
  },
  methods: {
    async login() {
      let res = await axios.post("/login/" + this.game_id);
      if (res.data.valid) {
        this.to = {
          name: "About",
          params: { game_id: this.game_id, valid: res.data.valid }
        };
      }
    },
    async newGame() {
      console.log("New game!");
      let res = await axios.post("/newgame");
      this.game_id = res.data.payload.game_id;
      this.status = "New game id is: " + this.game_id;
    }
  }
};
</script>