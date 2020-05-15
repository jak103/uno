<template>
  <div class="about">
    <h1>This is an about page</h1>
    <h2>Current Game id: {{ game_id }}</h2>
    <div v-if="valid">
      Game is valid!
    </div>
    <div v-else>
      Game is invalid
    </div>
  </div>
</template>

<script>
// This will eventually be the game page
  export default {
    props: {
      game_id: {
        type: String,
        required: true
      },
      valid: {
        type: Boolean,
        required: true
      }
    },
  };
</script>