<template>
  <div>
    <v-container>
      <v-row :class="mb-6">
        <v-col>
          <!-- Game stats -->
          <v-row>
            <v-card class="ma-3 pa-6" outlined tile>
              Current Game id: {{ game_id }}
              <span v-if="!valid">Invalid Game Id</span>
            </v-card>
            <v-card class="ma-3 pa-6" outlined tile>
              <v-btn v-if="valid" @click.native="startGame">Start Game</v-btn>
              <v-btn v-else to="/">Create a new game</v-btn>
            </v-card>
          </v-row>
          <!-- Game Players -->
          <v-row caption="Current Players">
            <v-card
              v-for="player in players"
              :key="player"
              color="current_player == player ? '#1F7087' : ''"
              class="ma-3 pa-6"
              outlined
              tile
            >
              {{ player }}
              <span v-if="current_player == player">
                <br />On clock
              </span>
            </v-card>
          </v-row>
          <!-- Current Card and actions -->
          <v-row v-if="current_player != ''">
            <v-card class="ma-3 pa-6" outlined tile>
              Current Card
              <br>
              <Card
                v-for="card in current_card"
                :number="card.number"
                :key="card.color"
                :color="card.color"
              />
            </v-card>
            <v-card class="ma-3 pa-6" outlined tile>
              <v-btn>Take from pile</v-btn>
              <br />
              <br />or
              <br />
              <br />
              <v-btn>Draw from deck</v-btn>
            </v-card>
          </v-row>
        </v-col>
        <v-col class="mb-6" v-if="current_card != ''">
          <v-card>Click to play a card from your hand</v-card>
          <br />
          <Card v-for="(card, i) in cards" :key="i" :number="card.number" :color="card.color" @click.native="playCard"></Card>
        </v-col>
      </v-row>
    </v-container>
  </div>
</template>

<script>
import axios from "axios";
import Card from "../components/Card";
export default {
  props: {
    game_id: {
      required: true
    },
    valid: {
      type: Boolean,
      required: true
    }
  },
  data() {
    return {
      cards: [],
      current_player: "",
      players: [],
      current_card: []
    };
  },
  components: {
    Card
  },
  methods: {
    async startGame() {
      let res = await axios.post("/startgame/" + this.game_id);
      if (res.data.valid) {
        this.valid = res.data.valid;
        this.cards = res.data.payload.cards;
        this.current_player = res.data.payload.current_player;
        this.players = res.data.payload.players;
        this.current_card = res.data.payload.current_card;
      }
    },
    playCard() {
      console.log("Playing card!")
    }
  }
};
</script>