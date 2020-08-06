<template>
  <div class="home">
    <v-container 
      fluid
      pt-16
    >
      <v-row justify="center">
        <v-card 
          class="elevation-12" 
          outlined
          width="1000"              
        >
          <v-card-title 
            style="background-color: #007BC7"                                 
          >
            Games  
            <v-spacer></v-spacer>
            <v-text-field
              v-model="search"
              append-icon="mdi-magnify"
              label="Search"
              single-line
              hide-details
            ></v-text-field>  
            <v-spacer></v-spacer>
            <v-btn
              icon
              @click="getAllGames()"
            >
              <v-icon>mdi-refresh</v-icon>
            </v-btn>
            <v-btn
              icon
              @click="createDialog.visible = true"
            >
              <v-icon>mdi-plus</v-icon>
            </v-btn>
            <v-spacer></v-spacer>
            <v-btn
              @click="deleteItems"
            >
              <small>Delete Selected Games</small>
            </v-btn>
          </v-card-title>

          <!-- Data Table -->
          <v-data-table
            :headers="headers"
            :items="games"
            :search="search"
            dense
          >

            <!-- CheckBox -->
            <template v-slot:item.check="{item}">
                <v-checkbox
                  class="mt-n1"
                  v-model="item.selected"
                  primary
                  hide-details
                ></v-checkbox>
            </template>

            <!-- Watch or Join Button -->
            <template v-slot:item.action="{item}">
              <v-btn
                text
                @click="handleActionClick(item)"
                color="primary"
              >
                {{item.status == "Playing" ? "WATCH" : "JOIN"}}
              </v-btn>
            </template>
          </v-data-table>

        </v-card>
      </v-row>
    </v-container>
    <v-dialog 
      v-model="joinDialog.visible" 
      persistent
      max-width="500px"
    >
      <v-card >
        <v-card-title
          class="blue"
        >
          <strong>Game Name: {{joinDialog.game.name}}</strong>
        </v-card-title>
        <v-card-text>
          <v-card-subtitle>
            <h3>Current Players</h3>
          </v-card-subtitle>
          <v-list>
            <v-list-item
              v-for="(player, i) in joinDialog.game.players"
              :key="i"
            >
              <v-list-item-content
                class="pl-4"
              >
                -- {{player.name}}
              </v-list-item-content>
            </v-list-item>
          </v-list>
          <v-text-field
            label="Your name"
            outlined
            v-model="joinDialog.yourname"
          >
          </v-text-field>
        </v-card-text>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn color="blue darken-1" text @click="closeJoinDialog">Cancel</v-btn>
          <v-btn color="blue darken-1" text @click="joinGame">Join</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
    <v-dialog 
      v-model="createDialog.visible" 
      persistent
      max-width="500px"
    >
      <v-card >
        <v-card-title
          class="blue"
        >
          Create Game
        </v-card-title>
        <v-card-text>          
          <v-text-field
            @keydown.enter="createGame"
            autofocus
            label="Game name"
            outlined
            v-model="createDialog.name"
            class="pt-4"
          > </v-text-field>
          <v-text-field
            @keydown.enter="createGame"
            label="Creator name"
            outlined
            v-model="createDialog.creator"            
          >
          </v-text-field>
        </v-card-text>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn color="blue darken-1" text @click="closeCreateDialog">Cancel</v-btn>
          <v-btn color="blue darken-1" text @click="createGame">Create & Join</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </div>
</template>

<script>
import unoService from '../services/unoService';
import localStorage from '../util/localStorage';
import bus from '../helpers/bus';

export default {
  name: 'Lobby',
  components: {

  },
  data() {
    return {
      search: "",
      headers:[
        { text: "Check", value: "check" },
        { text: "Name", value: "name" },
        { text: "Creator", value: "creator" },
        { text: "# of Players", value: "players.length" },
        { text: "Status", value: "status" },
        { text: "Action", value: "action" },
      ],
      games: [],
      joinDialog: {
        visible: false,
        headers: [
          { text: "Player Name" }
        ],
        game: {},
        yourname: ""
      },
      createDialog: {
        visible: false,
        name: "",
        creator: ""
      }
    }
  },

  methods: {
    async getAllGames() {
      let res = await unoService.getAllGames();
      this.games = res.data;
    },
    
    async joinGameOnLoad(gameid) {
      // pop up join to game they were invited to, if they are responding to an invite
      
      try {
          let res = await unoService.getGameSummary(gameid);
          
          if (res.data != null && res.data.status == "Waiting For Players") {
            // let the user join this game!
            this.joinDialog.game = res.data;
            this.joinDialog.visible = true;
          } else {
            // tell the user they cannot join that game!
            // invalid game name -- TODO use a snack bar for this
            alert("This game is not joinable!");
          }
        } catch {
          // invalid game name -- TODO use a snack bar for this
          alert("game was not found!");
        }
      
    },
    
    handleActionClick(game) {
      if (game.status == "Playing") {
        this.$router.push({path: `/game/${game.id}`});
      } else {
        this.joinDialog.game = game;
        this.joinDialog.visible = true;
      }
    },

    clearJoinDialog() {
      this.joinDialog.game = {};
      this.joinDialog.yourname = "";
    },

    closeJoinDialog() {
      this.joinDialog.visible = false;
      this.clearJoinDialog();
    },

    closeCreateDialog() {
      this.createDialog.visible = false;
    },

    async joinGame() {      
      this.joinDialog.visible = false;
      let res = await unoService.joinGame(this.joinDialog.game.id, this.joinDialog.yourname);

      this.clearJoinDialog();
      if (res.data.token && res.data.game) {
        localStorage.set('token', res.data.token);
        this.$router.push({path: `/game/${res.data.game.game_id}`});
      } else {
        //show the snack with your error message (just local)
        bus.$emit('updateSnack', "Failed to create & join game");
      }
    },

    async createGame() { 
      if (!this.createDialog.name || this.createDialog.name == "") {
        //show the snack with your error message (just local)
        bus.$emit('updateSnack', "Undefined Game name");
        return;
      }
     
      if (!this.createDialog.creator || this.createDialog.creator == "") {
        //show the snack with your error message (just local)
        bus.$emit('updateSnack', "Undefined Creator Name");
        return;
      }

      let res = await unoService.newGame(this.createDialog.name, this.createDialog.creator);
      
      if (res.data.token && res.data.game) {
        localStorage.set('token', res.data.token);
        this.$router.push({path: `/game/${res.data.game.game_id}`});
      } else {
        //show the snack with your error message (just local)
        bus.$emit('updateSnack', "Failed to create & join game");
      }
    },
    deleteItems () {
      let NotDeleted = [];
      for ( var i = this.games.length - 1; i >= 0; i--) {
        if (this.games[i].selected) {
          if (this.games[i].status == "Finished") {
            unoService.deleteGame(this.games[i].id);
            this.games.splice(i, 1);
          }else{
            NotDeleted.push(this.games[i].name);
          }
        }
      }
      if (NotDeleted.length > 0) {
        let plural = (NotDeleted.length > 1) ? ['Games', 'Are', 'them'] : ['Game', 'Is', 'it'];
        var notification = `${plural[0]}: [ ` + NotDeleted.toString() + ` ]. ${plural[1]} not Finished yet, You are unable to Delete ${plural[2]}.`;
        bus.$emit('updateSnack', notification);
      }
    },
  },

  mounted() {
    this.getAllGames();
  },
  
  created (){
    // get the game id they wish to join, if it exists
    let gameid = window.location.hash.substr(1);
    if(gameid.length > 0){
      // if there is a gameid we want to join(from an invite), get the game state to make sure it is still waiting for players to join.
      this.joinGameOnLoad(gameid);
    }
  },
}
</script>
