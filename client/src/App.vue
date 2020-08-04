<template>
  <v-app>
    <v-content>
      <!-- Imports the login. See index.js -->
      <v-app-bar
        color="#00263A"      
      >
          <v-btn
            text 
            large
            to='/'
          >
            <div class="headline">
                USU DevOps :: UNO
            </div>
          </v-btn>
      </v-app-bar>
      
      <router-view />
      <v-snackbar
        v-model="snackbar"
        color="info"
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
// import HelloWorld from './components/HelloWorld';
// import axios from 'axios';

export default {
  name: "App",

  components: {},
  data: () => ({
    snackbar: false,
    snackbarText: "",
    //
  }),
  mounted() {
    this.$root.$on('updateSnack', (message) => {
      this.snackbarText=message;
      this.snackbar=true;
      console.log("you called me!!!!!!!");
    }); 
  },
  methods: {
    //
  },
};
</script>
