<template>
  <div>
    <v-container>
      <v-row :class="'mb-6'">
        <v-col :cols="6">
          <!-- Game stats -->
          <v-row>
            <v-card :class="'ma-3 pa-6'" outlined tile>
              Current Game id: {{ game_id }}
              <span v-if="!valid">Invalid Game Id</span>
            </v-card>
            <v-card :class="'ma-3 pa-6'" outlined tile>
              <v-btn v-if="valid" @click.native="startGame">Start Game</v-btn>
              <v-btn v-else to="/">Create a new game</v-btn>
            </v-card>
          </v-row>

          <!-- Game Players -->
          <v-row>
            <v-card
              v-for="player in players"
              :key="player"
              color="current_player == player ? '#1F7087' : ''"
              :class="'ma-3 pa-6'"
              outlined
              tile
            >{{ player }}</v-card>
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
          <v-card v-else-if="game_over != ''">{{game_over}} has won the game!</v-card>

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
import axios from "axios";
import Card from "../components/Card";
export default {
  props: {
    game_id: {
      required: false
    },
    valid: {
      type: Boolean,
      required: false
    },
    username: {
      type: String,
      required: false
    }
  },
  data() {
    return {
      cards: [],
      current_player: "",
      players: [],
      current_card: [],
      game_over: ""
    };
  },
  components: {
    Card
  },
  methods: {
    updateData() {
      axios.get("http://localhost:8080/update/" + this.game_id + "/" + this.username).then(res => {
        if (res.data.valid) {
          this.valid = res.data.valid;
          this.cards = res.data.payload.deck;
          this.current_player = res.data.payload.current_player;
          this.players = res.data.payload.all_players;
          this.current_card = res.data.payload.current_card;
          if (res.data.game_over != "") {
            this.game_over = res.data.game_over;
          }
        }
      });
    },
    startGame() {
      axios
        .post("http://localhost:8080/startgame/" + this.game_id + "/" + this.username)
        .then(() => {
          this.updateData();
        });
    },
    playCard(card) {
      axios
        .post(
          "http://localhost:8080/play/" +
            this.game_id +
            "/" +
            this.username +
            "/" +
            card.number +
            "/" +
            card.color
        )
        .then(() => {
          this.updateData();
        });
    },
    drawCard() {
      var f = "http://localhost:8080/draw/" + this.game_id + "/" + this.username;
      console.log(f);
      axios.post(f).then(res => {
        console.log(res.data);
        this.updateData();
      });
    }
  },
  created() {
    setInterval(() => {
      this.updateData();
    }, 2000);
  }
};
</script>