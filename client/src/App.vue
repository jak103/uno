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
    //
  }),
  methods: {
    //
  },
};
</script>
