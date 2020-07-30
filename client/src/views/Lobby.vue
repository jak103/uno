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
                style="background-color: #33B0FF"                                 
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
              </v-card-title>
              <v-data-table
                :headers="headers"
                :items="games"
                :search="search"
              >
                <template v-slot:item.action="{item}">
                  <v-btn
                    text
                    @click="handleActionClick(item.id)"
                    color="primary"
                  >
                    {{item.status == "Playing" ? "WATCH" : "JOIN"}}
                  </v-btn>
                </template>
              </v-data-table>
            </v-card>
        </v-row>
      </v-container>
  </div>
</template>

<script>
// @ is an alias to /src
export default {
  name: 'Lobby',
  components: {

  },

  data() {
    return {
      search: "",
      headers:[
        { text: "Name", value: "name" },
        { text: "Creator", value: "creator" },
        { text: "# of Players", value: "playerCount" },
        { text: "Status", value: "status" },
        { text: "Action", value: "action" },

      ],
      games: [
        { id: "45223235", name: "mygame", creator: "Jak", playerCount: "4", status: "Waiting for players" },
        { id: "83732332", name: "Best game ever", creator: "Logan", playerCount: "1", status: "Playing" },
        { id: "23423553", name: "Friends only", creator: "Ryan", playerCount: "16", status: "Playing" },
        { id: "35939332", name: "Game", creator: "Brianna", playerCount: "2", status: "Waiting for players" },
      ]
    }
  },

  methods: {
    async getAllGames() {
      // TODO: get all the games and 
    }
  },

  mounted() {
    this.getAllGames();
  }
}
</script>
