<template>
  <div>
    You lost! But do you know who didn't lose?
    <p>{{winner}} </p>
    <button v-on:click="counter += 1">Add 1</button>
    <div id = "buttonContainer">
    <div id="playAgain">
      <button id = "buttons" v-on:click="home"> Play Again?</button>
    </div>
  </div>
  </div> 
</template>

<script>
export default {
  name: "Loser",
  props: ["winner"],
  methods: {
    redirectCall() {
      this.$router.push({ name: 'Login'});
    }
 }
}
</script>

<style scoped>
#buttonContainer{
    columns: 2;
    width: 650px;
    margin: 0 auto;
  }
</style>
