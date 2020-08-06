<template>
<div>
  <div v-if="players.winner === players.curPlayer">
    <p style="font-size: 36px; text-align:center">
    WINNER! Click buttons to celebrate.
    </p>
    <template>
      <main>
        <button @click="start">Click to celebrate!</button>
        <br>
        <button @click="stop">Click to stop celebrating.</button>
      </main>
    </template>
    <p style="font-size: 24px;text-align:center"> 
      {{players.winner}} 
    </p>
  </div>
  <div v-else style="font-size: 24px;text-align:center">
    You lost! But do you know who didn't lose? 
    <p>
      {{players.winner}} 
    </p>
    <template>
      <main>
        <button @click="love">Click to show some love to the winner.</button>
        <br>
        <button @click="stop">Click to stop.</button>
      </main>
    </template>     
  </div>
    <div id = "buttonContainer">
      <div style="text-align:center" id="playAgain">
        <button style="width:150px" id = "buttons" v-on:click="redirectCall()"> Play Again?</button>
      </div>
    </div>
</div>
</template>

<script>
import Vue from 'vue'
import VueConfetti from 'vue-confetti'

Vue.use(VueConfetti)

export default {
  name: "Results",
  props: {
    players: {
      winner: "winner", 
      curPlayer: "curPlayer"
      }
  },
  methods: {
    redirectCall() {
      this.$router.push({ name: 'Lobby'});
    },
    
    start() {
      this.$confetti.start();
    },
 
    stop() {
      this.$confetti.stop();
    },
 
    love() {
      this.$confetti.start();
      this.$confetti.update({
        particles: [
          {
            type: 'heart',
          },
        ],
        defaultColors: [
          'red',
          'pink',
          '#ba0000'
        ],
        defaultSize: 15,
      });
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
    margin-left:250px;
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
