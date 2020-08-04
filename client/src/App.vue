<template>
  <v-app>
    <v-content>
      <!-- Imports the login. See index.js -->
      <v-app-bar color="#00263A">
          <v-btn text large to='/'>
            <div class="headline">USU DevOps :: UNO</div>
          </v-btn>

          <v-btn text :to="redirect()" >
            <div>LastGame</div>
          </v-btn>

          <v-btn text :to="'/help'" >
            <div>Help</div>
          </v-btn>
      </v-app-bar>
      
      <router-view @sendGameID="getGameID"/>
    
    </v-content>
  </v-app>
</template>

<script>
export default {
  name: "App",
 
  components: {},
  data: () => ({
    gameID: null,
  }),
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
