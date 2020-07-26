<template>
  <div>
    <v-container>
      <v-row :class="'mb-9'">
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
            <v-card :class="'ma-3 pa-6'" outlined tile> 
                <!-- <v-div class="hintbtn"> -->
              <v-btn @click.native="hint">Hint</v-btn>
                <!-- </v-div>     -->
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
          <!-- Help menu button -->

          <v-col cols= "6">
            <v-row :class="'mb-6'">
              <v-card :class="'ma-3 pa-6'" outlined tile> 
                <v-div class="dropdown">
                  <v-btn @click.native="helpMenu" class="helpDropBtn">Need Help?</v-btn>
                  <v-div class="dropdown_content">
                    <router-link to="help#section-one" @click.native="help('#section-one')">Rules</router-link>
                    <router-link to="help#section-two" @click.native="help('#section-two')">Tutorials</router-link>
                    <router-link to="help#section-three" @click.native="help('#section-three')">Card Abilities</router-link>
                    <!-- <a href="help" v-scroll-to="'#section-one'">Rules</a>
                    <a href="help" v-scroll-to="'#section-two'">Tutorials</a>
                    <a href="help" v-scroll-to="'#section-three'">Card Abilities</a> -->
                  </v-div>
                </v-div>    
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
            <!-- Click to play a card from your hand or -->
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
import unoService from "../services/unoService";
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
      unoService.update(this.game_id, this.username).then(res => {
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
    hint(){
      var color = this.current_card[0].color
      var number = this.current_card[0].number
      alert("Play a card with the number " + number + " or a card that is the color " + color + ".")
    },
    helpMenu(){
      window.onclick = function(event) {
        if (!event.target.matches('.helpDropBtn')) {
          var dropdowns = document.getElementsByClassName("dropdown_content");
          var i;
          for (i = 0; i < dropdowns.length; i++) {
            var openDropdown = dropdowns[i];
            if (openDropdown.classList.contains('show')) {
              openDropdown.classList.remove('show');
            }
          }
        }
      }
    },

    startGame() {
        unoService.startGame(this.game_id, this.username)
        .then(() => {
          this.updateData();
        });
    },
    playCard(card) {
        unoService.playCard(this.game_id, this.username, card.number, card.color)
        .then(() => {
          this.updateData();
        });
    },
    drawCard() {
      unoService.drawCard(this.game_id, this.username)
        .then(this.updateData());
    }
  },
  created() {
    setInterval(() => {
      this.updateData();
    }, 2000);
  }
};
</script>


<style scoped>
@import url(https://fonts.googleapis.com/css?family=Source+Sans+Pro:900);
  
  .hintbtn{
    background-color: white;
  }
  .helpDropBtn {
    background-color: #4CAF50;
    color: white;
    padding: 16px;
    font-size: 16px;
    border: none;
  }

/* The container <div> - needed to position the dropdown content */
  .helpDropBtn {
    position: relative;
    display: inline-block;
  }

/* Dropdown Content (Hidden by Default) */
  .dropdown_content {
    display: none;
    position: absolute;
    background-color:white;
    min-width: 160px;
    box-shadow: 0px 8px 16px 0px rgba(0,0,0,0.2);
    z-index: 1;
  }

  /* Links inside the dropdown */
  .hintbtn a, .dropdown_content a {
    color: black;
    padding: 6px 8px;
  }

  .dropdown_content a {
    color: black;
    padding: 12px 16px;
    text-decoration: none;
    display: block;
  }

  /* Change color of dropdown links on hover */

  .dropdown_content a:hover, .hintbtn a:hover {background-color: #ddd;}

  /* Show the dropdown menu on hover */
  .dropdown:hover .dropdown_content, .hintbtn a:hover {display: block;}

  /* Change the background color of the dropdown button when the dropdown content is shown */
  .dropdown:hover .helpDropBtn, .hintbtn a:hover  {background-color: #3e8e41;}
</style>