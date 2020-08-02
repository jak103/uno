<template>
  <div>
    <v-container>
      <v-row class="mb-6">
      <!-- <v-row> -->

        <v-col :cols="6">
          <!-- Game stats -->
          <v-row>
            <v-card class="ma-3 pa-6" outlined tile>
              Current Game id: {{ gameState.game_id }}
              <br />
              Status: {{gameState.status}}
            </v-card>
          </v-row>

          <!-- Game Players -->
          <h3>Players</h3>
          <v-row v-if="gameState.all_players !== undefined">
            <v-card
              v-for="player in gameState.all_players"
              :key="player.name"
              :color=" player.id == gameState.current_player.id ? '#1F7087' : ''"
              class="ma-3 pa-6"
              outlined
              tile
            >{{ player.name }}</v-card>
          </v-row>

          <!-- Current Card and actions -->
          <v-col cols="12" v-if="gameState.status === 'Playing'">
            <v-row v-if="gameState.current_card != undefined">
              <v-card class="center-text ma-3 pa-6" outlined tile>
              <center>Current Card</center>
                <Card
                  :number="gameState.current_card.value"
                  :key="gameState.current_card.color"
                  :color="gameState.current_card.color"
                />
              </v-card>
            </v-row>
          </v-col>
        </v-col>

        <!-- Current cards in the deck -->
        <v-col class="mb-6" v-if="gameState.status === 'Playing'">
          <v-card
            v-if="gameState.status !== 'Finished'"
            class="ma-3 pa-6"
            outlined
            tile
          >
            Click to play a card from your hand or
            <v-btn v-if="gameState.current_player != undefined &&  gameState.player_id === gameState.current_player.id" @click.native="drawCard">Draw from deck</v-btn>
          </v-card>
          <v-card v-else-if="gameState.status === 'Finished'">The game is finished!</v-card>

          <v-card v-else class="ma-3 pa-6" outlined tile>Waiting for {{ gameState.current_player.name }}</v-card>
          <Card
            v-for="(card, i) in gameState.player_cards"
            :key="i"
            :number="card.value"
            :color="card.color"
            @click.native="playCard(card)"
          ></Card>
        </v-col>
        <v-col v-else>
          <v-card
            class="ma-3 pa-6"
            outlined
            tile
          >
            <v-row v-if="gameState.creator != undefined && gameState.creator.id == gameState.player_id">
                You are the creator of the game. When you are ready: <v-btn @click.native="startGame">Start Game</v-btn>
            </v-row>
            <v-row v-else>
                Please wait for the creator to start the game.
            </v-row>
          </v-card>
        </v-col>
      </v-row>
    </v-container>
  </div>
</template>

<script>
import unoService from "../services/unoService";
import Card from "../components/Card";
export default {
  name: "Game",
  components: {
    Card
  },
  data() {
    return {
      gameState: {},
      cards: []
    };
  },
  methods: {
    async updateData() {      
      let res = await unoService.getGameState(this.$route.params.id);
      
      if (res.data != null) {
        this.gameState = res.data;
      }
    },

    async startGame() {
      await unoService.startGame(this.$route.params.id);
      // TODO make sure startGame endpoint returns the game state and then remove this call to updateData()
      this.updateData(); 
    },

    async playCard(card) { 
      console.log("Playing card", card);     
      let res = await unoService.playCard(this.$route.params.id, card.value, card.color);
      
      if (res.data) {
        this.gameState = res.data;
      }
    },

    async drawCard() {
      let res = await unoService.drawCard(this.$route.params.id);
      
      if (res.data) {
        this.gameState = res.data;
      }
    }
  },
  created() {
    this.updateData();
    this.updateInterval = setInterval(() => {
      this.updateData();
    }, 2000);
  },
  beforeDestroy (){
    if(this.updateInterval){
      clearInterval(this.updateInterval);
    }
  }
};
</script>