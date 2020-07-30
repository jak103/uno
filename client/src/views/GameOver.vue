<template>
  <div v-if="players.cur===players.winner">
    Winner winner
    <Winner />
  </div>
  <div v-else>
    <Loser v-bind:winner="players.winner" />
  </div>
</template>

<script>

import Loser from "../components/Loser";
import Winner from "../components/Winner";

export default {
  name: 'GameOver',
  components: {
    Loser,
    Winner
  },
  data() {
    return {
      players: {
        cur : this.$route.params.cur,
        winner : this.$route.params.winner
      }
    }
  }
}
</script>

<style scoped>

</style>>
