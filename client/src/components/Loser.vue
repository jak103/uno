<template>
  <div>
    You lost! But do you know who didn't lose?
    <p>{{winner}} </p>
  </div>
</template>

<script>
export default {
  name: "Loser",
  props: ["winner"]
}
</script>

<style scoped>

</style>
