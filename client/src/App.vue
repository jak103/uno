<template>
  <v-app>
    <v-content>
      <!-- Imports the login. See index.js -->
      <v-app-bar color="#00263A">
          <v-btn text large to='/'>
            <div class="headline">USU DevOps :: UNO</div>
          </v-btn>
          <v-spacer></v-spacer>
          <v-btn text :to="redirect()" >
            <div>Game</div>
          </v-btn>
      </v-app-bar>
      
      <router-view @sendGameID="getGameID"/>
      <v-snackbar
        v-model="snackbar"
        :color="color"
        :timeout='0'>
        {{snackbarText}}

        <v-btn text @click="snackbar=false">
          Close
        </v-btn>
      </v-snackbar>
    </v-content>
  </v-app>
</template>

<script>
import bus from "./helpers/bus"
export default {
  name: "App",
 
  components: {},
  data: () => ({
    snackbar: false,
    snackbarText: "",
    color: "info",
    gameID: null,
  }),
  mounted() {
    bus.$on('updateSnack', (message, color) => {
      color = color || "info";
      this.color = color;
      this.snackbarText=message;
      this.snackbar=true;
    }); 
  },
  methods: {
    getGameID(value) {
      this.gameID = value;
    },
    redirect() {
      if (this.gameID != null){
        return '/game/' + this.gameID;
      }else{
        return '/'
      }
    }
  }
};
</script>
