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
                  <v-btn color="primary">Create new game</v-btn>
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
import axios from 'axios'
export default {
  name: "LoginPage",
  data: () => {
    return {
      game_id: null,
      valid_game: false,
      to: {}
    };
  },
  methods: {
    async login () {
      let res = await axios.post('/login/' + this.game_id)
      console.log(res.data)
      if (res.data.valid) {
        this.to = {name: 'About', params: {'game_id': this.game_id, 'valid': res.data.valid}};
      }
      // figure the else out
    },
  }
};
</script>