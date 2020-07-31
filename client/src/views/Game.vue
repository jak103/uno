<template>
  <div>
    <v-container>
      <v-row :class='mb-6'>
        <v-col :cols="6">
          <!-- Game stats -->
          <v-row>
            <v-card :class="'ma-3 pa-6'" outlined tile>
              Current Game id: {{ $route.params.id }}              
            </v-card>
          </v-row>

          <!-- Game Players -->
          <v-row>
            <v-card
              v-for="player in players"
              :key="player"
              :color="current_player == player ? '#1F7087' : ''"
              :class="'ma-3 pa-6'"
              outlined
              tile
            >{{ player.name }}</v-card>
          </v-row>

          <!-- Current Card and actions -->
          <v-col cols="12">
            <v-row v-if="current_player != ''">
              <v-card :class="'ma-3 pa-6'" outlined tile class="center-text">
                <Card
                  v-for="card in current_card"
                  :number="card.number"
                  :key="card.color"
                  :color="card.color"
                />
              </v-card>
            </v-row>
          </v-col>
        </v-col>

        <!-- Current cards in the deck -->
        <v-col :class="'mb-6'" v-if="current_card != ''">
          <v-card
            v-if="username == current_player"
            :class="'ma-3 pa-6'"
            outlined
            tile
          >
            Click to play a card from your hand or
            <v-btn v-if="username == current_player" @click.native="drawCard">Draw from deck</v-btn>
          </v-card>
          <v-card v-else-if="!!game_over">{{game_over}} has won the game!</v-card>

          <v-card v-else :class="'ma-3 pa-6'" outlined tile>Waiting for {{ current_player }}</v-card>
          <Card
            v-for="(card, i) in cards"
            :key="i"
            :number="card.number"
            :color="card.color"
            @click.native="playCard(card)"
          ></Card>
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
      state: {}
    };
  },
  methods: {
    async updateData() {      
      let gameState = await unoService.getGameState(this.$route.params.id);
      if (gameState != null) {
        this.state = gameState;
      }
    },

    async startGame() {
      await unoService.startGame(this.$route.params.id);
      this.updateData();
      
    },

    async playCard(card) {
      await unoService.playCard(this.$route.params.id, this.username, card.number, card.color);
      this.updateData();      
    },

    async drawCard() {
      await unoService.drawCard(this.$route.params.id);
      this.updateData();
    }
  },
  created() {
    setInterval(() => {
      this.updateData();
    }, 2000);
  }

  // Thomas O'Reilly, stop type on destroy
};
</script>