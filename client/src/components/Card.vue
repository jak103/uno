<template>
  <div v-if="!isNaN(number)" :class="card_classes">
    <span class="inner">
      <span class="mark">{{ number }}</span>
    </span>
  </div>
  <div v-else-if="number === 'S'" :class="card_classes">
    <span class="inner">
      <span class="mark">
        <img v-if="color === 'green'" class="s-inner" src="@/assets/card_symbols/s-green.svg">
        <img v-else-if="color === 'yellow'" class="s-inner" src="@/assets/card_symbols/s-yellow.svg">
        <img v-else-if="color === 'red'" class="s-inner" src="@/assets/card_symbols/s-red.svg">
        <img v-else-if="color === 'blue'" class="s-inner" src="@/assets/card_symbols/s-blue.svg">
      </span>
    </span>
  </div>
  <div v-else-if="number === 'D2'" :class="card_classes">
    <span class="inner">
      <span class="mark">
        <img v-if="color === 'green'" class="d2-inner" src="@/assets/card_symbols/d2-green.svg">
        <img v-else-if="color === 'yellow'" class="d2-inner" src="@/assets/card_symbols/d2-yellow.svg">
        <img v-else-if="color === 'red'" class="d2-inner" src="@/assets/card_symbols/d2-red.svg">
        <img v-else-if="color === 'blue'" class="d2-inner" src="@/assets/card_symbols/d2-blue.svg">
      </span>
    </span>
  </div>
  <div v-else-if="number === 'W4'" :class="card_classes">
    <span class="inner">
      <span class="mark">
        <img class="w4-inner" src="@/assets/card_symbols/d4.svg">
      </span>
    </span>
  </div>
  <div v-else-if="number === 'W'" :class="card_classes">
    <span class="inner">
      <span class="mark">
        <img class="w-inner" src="@/assets/card_symbols/w.svg">
      </span>
    </span>
  </div>
  <div v-else-if="number === 'R'" :class="card_classes">
    <span class="inner">
      <span class="mark">
        <img v-if="color === 'green'" class="r-inner" src="@/assets/card_symbols/r-green.svg">
        <img v-else-if="color === 'yellow'" class="r-inner" src="@/assets/card_symbols/r-yellow.svg">
        <img v-else-if="color === 'red'" class="r-inner" src="@/assets/card_symbols/r-red.svg">
        <img v-else-if="color === 'blue'" class="r-inner" src="@/assets/card_symbols/r-blue.svg">
      </span>
    </span>
  </div>
  <div v-else :class="card_classes">
    <span class="inner">
      <span class="mark">E{{number}}</span>
    </span>
  </div>
</template>

<script>
export default {
  name: "Card",
  props: {
    number: {
      required: true
    },
    color: {
      type: String,
      required: true
    },
  },
  computed: {
    card_classes: function() {
      if(!isNaN(this.number))
        return `card num-${this.number} ${this.color.toLowerCase()}`;
      else
        return `card ${this.number.toLowerCase()} ${this.color.toLowerCase()}`;
    },
  }
};
</script>

<style scoped>
@import url(https://fonts.googleapis.com/css?family=Source+Sans+Pro:900);

.card {
  cursor: pointer;
  width: 116px;
  height: 178px;
  float: left;
  background: #fff;
  border-radius: 5px;
  display: table;
  box-sizing: border-box;
  padding: 5px;
  margin: 10px;
  font-family: "Source Sans Pro", sans-serif;
  font-size: 100px;
  text-shadow: 1px 1px 0 #000000, -1px -1px 0 #000000, -1px 1px 0 #000000,
    1px -1px 0 #000000, 1px 0 0 #000000, -1px 0 0 #000000, 0 -1px 0 #000000,
    0 1px 0 #000000, 4px 4px 0 #000000;
  box-shadow: 0 0 10px #aaaaaa;
  text-align: center;
  position: relative;
}
.card .inner {
  display: table-cell;
  vertical-align: middle;
  border-radius: 5px;
  overflow: hidden;
}

.card .mark {
  display: inline-block;
  vertical-align: middle;
  background: #ffffff;
  margin: auto;
  padding: 0 26px;
  border-radius: 100px 60px / 120px 60px;
  line-height: 1.4;
}

.card:before,
.card:after {
  display: inline-block;
  position: absolute;
  line-height: 0;
  font-size: 20px;
  color: #ffffff;
  text-shadow: 1px 1px 0 #000000, -1px -1px 0 #000000, -1px 1px 0 #000000,
    1px -1px 0 #000000, 1px 0 0 #000000, -1px 0 0 #000000, 0 -1px 0 #000000,
    0 1px 0 #000000, 2px 2px 0 #000000;
}

.card:before {
  top: 15px;
  left: 10px;
}
.card:after {
  bottom: 15px;
  right: 10px;
  -webkit-transform: rotate(180deg);
}

.card.blue {
  color: #0063b3;
}

.card.blue .inner {
  background: #0063b3;
}

.card.green {
  color: #18a849;
}

.card.green .inner {
  background: #18a849;
}

.card.red {
  color: #c72a18;
}

.card.red .inner {
  background: #c72a18;
}

.card.yellow {
  color: #e6ca1e;
}

.card.yellow .inner {
  background: #e6ca1e;
}

.card.num-0:before,
.card.num-0:after {
  content: "0";
}

.card.num-1:before,
.card.num-1:after {
  content: "1";
}

.card.num-2:before,
.card.num-2:after {
  content: "2";
}

.card.num-3:before,
.card.num-3:after {
  content: "3";
}

.card.num-4:before,
.card.num-4:after {
  content: "4";
}

.card.num-5:before,
.card.num-5:after {
  content: "5";
}

.card.num-6:before,
.card.num-6:after {
  content: "6";
}

.card.num-7:before,
.card.num-7:after {
  content: "7";
}

.card.num-8:before,
.card.num-8:after {
  content: "8";
}

.card.num-9:before,
.card.num-9:after {
  content: "9";
}

.card.num-6 .mark:after,
.card.num-9 .mark:after {
  display: block;
  content: "";
  position: relative;
  top: -25px;
  left: 5px;
  width: 80%;
  border: 1px solid #000000;
  height: 4px;
  box-shadow: 1px 1px 0 #000000;
}
.card.num-6.blue .mark:after,
.card.num-9.blue .mark:after {
  background: #0063b3;
}
.card.num-6.green .mark:after,
.card.num-9.green .mark:after {
  background: #18a849;
}
.card.num-6.red .mark:after,
.card.num-9.red .mark:after {
  background: #c72a18;
}
.card.num-6.yellow .mark:after,
.card.num-9.yellow .mark:after {
  background: #e6ca1e;
}
/* Start Reverse Arrow CSS */
.r:before,
.r:after {
  content: "reverse";
}
.r-inner{
  height: 100%;
  width: 100%;
  margin-bottom: -10%;
}
/* End Reverse Arrow CSS */
/* Start Skip CSS */
.s:before,
.s:after {
  content: "skip";
}

.s-inner{
  height: 150%;
  width: 150%;
  margin-left: -25%;
  margin-bottom: -10%;
}
/* End Skip CSS */

/* Start Draw 2 CSS */
.d2:before,
.d2:after {
  content: "draw 2";
}

.d2-inner {
  height: 100%;
  width: 100%;
}
/* End Draw 2 CSS */

/* Start Wild CSS */
.w:before,
.w:after {
  content: "wild";
}
.w-inner {
  height: 125%;
  width: 125%;
  margin-left: -63%;
  margin-top: -104%;
  position: absolute;
}
/* End Wild CSS */

/* Start Draw 4 CSS */
.w4:before,
.w4:after {
  content: "draw 4";
}
.mark.w4-mark {
  background-image: linear-gradient(to left, red, yellow, green, blue) !important;
}
.w4-inner {
  height: 120%;
  width: 120%;
  margin-left: -10%;
  margin-bottom: -15%;
}
/* End Draw 4 CSS */

.card.black .inner {
  background: black;
}
</style>