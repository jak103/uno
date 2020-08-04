<template>
  <div >
    You lost! But do you know who didn't lose? asdfasdf
    <p>1 {{gameState.GameOver}} 
      asdf
      2 {{curPlayer}}
    </p>
    <!-- <button v-on:click="counter += 1">Add 1</button> -->
    <div id = "buttonContainer">
    <div id="playAgain">
      <button id = "buttons" v-on:click="home"> Play Again?</button>
    </div>
  </div>
  </div> 
</template>

<script>
export default {
  name: "Results",
  // props: ["winner", "curPlayer"],
  methods: {
    redirectCall() {
      this.$router.push({ name: 'Lobby'});
    }
 }
}
</script>

<style scoped>
  #displayMessage{
    text-align: center;
    font-size: 300%;
    padding-top: 300px;
    padding-right: 30px;
    padding-bottom: 50px;
    padding-left: 80px;
  }
  #buttonContainer{
    columns: 2;
    width: 650px;
    margin: 0 auto;
  }
  #playAgain{
    padding-left: 150px;
    font-size: 150%;
  }
  #buttons{
    border-style: solid;
    border-color: white;
    padding-top: 10px;
    padding-right: 10px;
    padding-bottom: 10px;
    padding-left: 10px;
  }
  #goHome{
    padding-left: 50px;
    font-size: 150%;
  }
</style>
