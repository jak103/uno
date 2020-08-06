
<template>
  <div class="mb-0 game-wrapper">
    <v-card 
      class="overflow-down"
    >
      <!-- Game Players Drawer -->
      <v-navigation-drawer>
        <v-list
          nav
          dense
        >
          <v-list-item two-line >
            <v-list-item-icon>
              <v-icon class="pt-3">
              mdi-account-group
              </v-icon>
            </v-list-item-icon>

            <v-list-item-content>
              <v-list-item-title>Players</v-list-item-title>
            </v-list-item-content>

            <v-list-item-icon>
              <v-icon class="pt-3" v-if="gameState.direction === true">
                mdi-arrow-down-bold
              </v-icon>
              <v-icon class="pt-3" v-else>
                mdi-arrow-up-bold
              </v-icon>
            </v-list-item-icon>
          </v-list-item>

          <v-divider
            class="mb-2"
          ></v-divider>
            
          <div
            v-if="gameState.all_players !== undefined"            
          >  
            <v-card
              v-for="player in gameState.all_players"
              :key="player.name"
              :color="getTileBackgroundColor(player)"
              class="pa-1"
            >
              <v-card-title
                class="pa-0 pl-1 drawer-card-title"
              >
                {{ player.name }}
                <v-btn
                  v-if="player.cards !== undefined && player.cards !== null"
                  :class="player.protection ? 'protected_call_button' : 'unprotected_call_button'" 
                  @click.native="callUno(player)"
                  :disabled="player.cards.length > 1"
                  class="pa-0"
                >
                  Uno!
                </v-btn>   
              </v-card-title>
              <v-card-text class="pa-0 pl-1">
                <span>
                  <span v-for="(card, index) of player.cards" :key="index">🃏</span>
                </span>
              </v-card-text>
            </v-card>
          </div>
        </v-list>
      </v-navigation-drawer>
    </v-card>
    <v-container>
      <v-row >
      <!-- <v-row> -->

        <v-col>
          <!-- Game stats -->
          <v-row>
            <v-card class="ma-3 pa-6" outlined tile>
              <p>
                Current Game id: {{ gameState.game_id }}
              </p>
              <p>
                Status: {{ gameState.status }}
              </p>
              <p>
                Your Name: {{ playerName }}
              </p>
              <p v-if="gameState.draw_pile != undefined">
                Cards Remaining in Draw Pile: {{ gameState.draw_pile.length }}
              </p>              

              <!-- Need Help? button -->
              <v-btn @click.native="helpMenu = !helpMenu" >Need Help?</v-btn>
              <v-card v-show="helpMenu" class="mt-5 pa-2" outlined tile >                  
                <router-link to="/help#rules">Rules</router-link> |
                <router-link to="/help#tutorials">Tutorials</router-link> |
                <router-link to="/help#cardAbilities">Card Abilities</router-link> |

                <!-- Hint Button -->
                <v-btn @click.native="hint">Hint</v-btn>
              </v-card>

            </v-card>
          </v-row>

          <!-- Current Card and actions -->
          <v-col cols="12" v-if="gameState.status === 'Playing'">
            <v-row v-if="gameState.current_card != undefined">
              <v-card class="center-text ma-3 pa-6" outlined tile>
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
        <v-col class="mb-6">
          <v-card
            class="ma-3 pa-6"
            outlined
            tile
          >
            <v-card-text v-if="gameState.status === 'Waiting For Players'">
              <v-row v-if="gameState.creator != undefined && gameState.creator.id == gameState.player_id">
                You are the creator of the game. When you are ready: <v-btn @click.native="startGame">Start Game</v-btn>
              </v-row>
              <v-row v-else>
                Please wait for the creator to start the game.
              </v-row>
            </v-card-text>

            <v-card-text v-if="gameState.status === 'Playing' && gameState.player_id === gameState.current_player.id">
              Click to play a card from your hand or <v-btn @click.native="drawCard">Draw from deck</v-btn>
            </v-card-text>

            <v-card-text v-else-if="gameState.status === 'Playing'">
              Waiting for {{ gameState.current_player.name }}
            </v-card-text>
          </v-card>

          <div v-if="gameState.status === 'Playing'" >

            <!-- Organize Cards -->
            <v-card :class="'ma-3 pl-6 pa-4'" outlined tile>
              <v-row v-if="loadingHand">
                Loading Original Hand Layout
              </v-row>

              <v-row v-else class="pl-3">
                Organize Cards
                <v-btn class="org-btn" @click.native="orgByColor">by Color</v-btn>
                <v-btn class="org-btn" @click.native="orgByNum">by Number</v-btn>
                <v-btn class="org-btn" @click.native="orgOff">Off</v-btn>
              </v-row>
            </v-card>

            <v-container
              class="card-container"
            >
              <Card
                v-for="(card, i) in gameState.player_cards"
                :key="i"
                :number="card.value"
                :color="card.color"
                :showColorDialog="card.showColorDialog"
                :ref="'player_cards'"
                @click.native="(card.value == 'W' || card.value == 'W4') ? selectWildColor(i) : playCard(card)"
                v-on:playWild="(color)=>playWildCard(color, i)"
              ></Card>
            </v-container>
          </div>
          <div v-if="gameState.status === 'Finished'">
            <Results v-bind:players="{
                winner: gameState.gameOver,
                curPlayer: playerName 
                }"/>
          </div>
        </v-col>

        <!-- Chat -->
        <v-col v-show="chatOpen" class="float-chat">
          <Chat @snackbarText="runsnackbar" :gameState="gameState"/>
        </v-col>
      </v-row>

    </v-container>

    <div 
      v-if="gameState.status === 'Playing' && gameState.current_player != undefined" 
      @click="chatOpen = !chatOpen"
      class="float-button">
        Chat
    </div>

      <v-snackbar
        v-model="snackbar"
        color="info"
        :timeout='4000'
        v-show="gameState.status === 'Playing' && playerName !== newMessageName"
        >{{snackbarText}}
        <v-btn text @click="snackbar=false">
          Close
        </v-btn>
      </v-snackbar>

  </div>
</template>

<script>
import unoService from "../services/unoService";
import Card from "../components/Card";
import Chat from "../components/Chat";
import Results from "../components/Results";

export default {
  
  name: "Game",
  title:"Greatest Uno",
  components: {
    Card,
    Chat,
    Results,
  },
  data() {
    
    return {
      gameState: {},
      cards: [],

      chatOpen: false,
      
      playerName: "",

      sortByNum: false,
      sortByColor: false,
      loadingHand: false,
      colors: { 'red': 0, 'blue': 1 , 'green': 2, 'yellow': 3, 'wild': 4},
      values: { '1' : 0, '2' : 1, '3' : 2, '4' : 3, '5' : 4, '6' : 5, '7' : 6, '8' : 7, '9' : 8, 'S' : 9, 'R' : 10, 'W' : 11, 'D2' : 12, 'W4' : 13},
    
      helpMenu: false,
      
      snackbar: false,
      snackbarText: "",
      newMessageName: "",
    };
  },

  methods: {
    async updateData() {
      let res = await unoService.getGameState(this.$route.params.id);

      if (res.data != null) {
        this.gameState = res.data;
      }
      this.decideSort()
    },  
    async startGame() {
      await unoService.startGame(this.$route.params.id);
      // TODO make sure startGame endpoint returns the game state and then remove this call to updateData()
      this.updateData(); 
    },

    // Methods for organizing the Cards, added by Andrew McMullin for the organize-cards issue
    decideSort() {
      if (this.sortByColor) {
        this.orgByColor()
      }else if (this.sortByNum) {
        this.orgByNum()
      }else{
        this.loadingHand = false
      }
    },
    orgOff() {
      if (this.sortByColor == true || this.sortByNum == true) {
        this.loadingHand = true;
      }
      this.sortByNum = false;
      this.sortByColor = false;
    },
    orgByNum() {
      if (this.gameState.player_cards != undefined) {
        this.gameState.player_cards.sort((a, b) => (this.colors[a.color] > this.colors[b.color]) ? 1 : -1 );
        this.gameState.player_cards.sort((a, b) => (this.values[a.value] > this.values[b.value]) ? 1 : -1 );
      }

      this.sortByNum = true;
      this.sortByColor = false;
    },
    orgByColor() {
      if (this.gameState.player_cards != undefined) {
        this.gameState.player_cards.sort((a, b) => (this.values[a.value] < this.values[b.value]) ? 1 : -1 );
        this.gameState.player_cards.sort((a, b) => (this.colors[a.color] < this.colors[b.color]) ? 1 : -1 );
      }

      this.sortByNum = false;
      this.sortByColor = true;
    },

    runsnackbar(name, message) {
      this.newMessageName = name;
      this.snackbarText = name + " says: " + message;
      this.snackbar = true;
    },

    selectWildColor(index)
    {
      this.$refs.player_cards[index].showColorDialog = true;
    },

    async playWildCard(color, i) {
      
      this.$refs.player_cards[i].showColorDialog = false;
      this.$refs.player_cards[i].color = color;

      let res = await unoService.playCard(
        this.$route.params.id, 
        this.$refs.player_cards[i].number, 
        this.$refs.player_cards[i].color
      );
     
      if (res.data) {
        this.gameState = res.data;
      }
    },

    async playCard(card) { 
      let res = await unoService.playCard(this.$route.params.id, card.value, card.color);
     
      if (res.data) {
        this.gameState = res.data;
      }
    },

    sendMessage() {
      this.snackbarText = this.username + " says: " + this.newMessage;
      this.snackbar = true;
    },

    async drawCard() {
      let res = await unoService.drawCard(this.$route.params.id);
      
      if (res.data) {
        this.gameState = res.data;
      }
    },

    async callUno(calledOnPlayer) {      
      let res = await unoService.callUno(this.gameState.game_id, calledOnPlayer)
      
      if (res.data) {
        this.gameState = res.data;
      }
    },

    // Getting a hint, added by the creator of the Help Button
    hint(){
      var color = this.gameState.current_card.color
      var number = this.gameState.current_card.value
      this.snackbarText = "Play a card with the number " + number + " or a card that is the color " + color + ".";
      this.snackbar = true;
    }, 

    getTileBackgroundColor(player) {
      if (player.isActive !== undefined && player.isActive !== null && player.isActive === false) {
        return "error";
      } else if (this.gameState.current_player !== undefined && this.gameState.current_player !== null && player.id === this.gameState.current_player.id) {
        return "info"
      } else {
        return ""
      }
    },
  }, 

  created() {
    this.updateData();
    this.updateInterval = setInterval(() => {
      this.updateData();
    }, 2000);
  },
  mounted() {
    this.$emit('sendGameID', this.$route.params.id)

    unoService.getPlayerNameFromToken()
    .then( resp => {
        this.playerName = resp?.data?.name
    })
    .catch(err => {
      console.err("Could not get player name from assigned token\n", err)
    })
  },
  beforeDestroy (){
    if(this.updateInterval){
      clearInterval(this.updateInterval);
    }
  },
};
</script>

<style scoped>

.float-chat {
  position:fixed;
	width:450px;
	height:550px;
  bottom:125px;
  right:50px;
	background-color:#00263A;
	color:#FFF;
	border-radius:10px;
  /* border-width: 5px;
  border-color: #FFF;
  outline: #FFF; */
  padding: 5px 5px 5px 5px;
}

.float-button{
	position:fixed;
	width:60px;
	height:60px;
	bottom:50px;
	right:50px;
	background-color:#00263A;
	color:#FFF;
	border-radius:50px;
  text-align:center;
  padding: 20px 0px 0px 0px;
	/* box-shadow: 2px 2px 3px #999; */
}

.game-wrapper {
  display: flex; 
  min-height: 100vh;
  height: 100vh;
}

.overflow-down
{
  overflow: auto;
  max-height: 100vh;
}

.v-btn {
  margin: 0px 10px 0px 10px;
}

.drawer-card-title {
  display: flex; 
  justify-content: space-between
}

.card-container {
  overflow-y: auto; 
  max-height: 60vh;
}

/* CSS for the Help buttons */
@import url(https://fonts.googleapis.com/css?family=Source+Sans+Pro:900);
  
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
    min-width: 160px;
    box-shadow: 0px 8px 16px 0px rgba(0,0,0,0.2);
    z-index: 1;
  }

  .unprotected_call_button {
    color: red;
    align-content: center;
    margin: 10px;
  }

  .protected_call_button {
    color: green;
    align-content: center;
    margin: 10px;
  }

  .org-btn {
    margin: 10px;
  }
  /* Show the dropdown menu on hover */
  .dropdown:hover .dropdown_content, .hintbtn a:hover {display: block;}
</style>
