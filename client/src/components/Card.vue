<template>
  <div v-if="typeof number === 'number'" :class="card_classes">
    <span class="inner">
      <span class="mark">{{ number }}</span>
    </span>
  </div>
  <div v-else-if="number === 'skip'" :class="card_classes">
    <span class="inner">
      <span class="mark">S</span>
    </span>
  </div>
  <div v-else-if="number === 'draw_two'" :class="card_classes">
    <span class="inner">
      <span class="mark">D</span>
    </span>
  </div>
  <div v-else-if="number === 'wild_draw_four'" :class="card_classes">
    <span class="inner">
      <span class="mark">4</span>
    </span>
  </div>
  <div v-else-if="number === 'wild'" :class="card_classes">
    <span class="inner">
      <span class="mark">w</span>
    </span>
  </div>
  <div v-else-if="number === 'reverse'" :class="card_classes">
    <span class="inner">

      <div class="reverse-mini reverse-mini-top">
        <div class="reverse-symbol shadow-mini">
          <div class="shadow-container">
            <div class="arrow-right">
              <div class="arrow-right-tail shadow"></div>
              <div class="arrow-right-head shadow"></div>
            </div>
            <div class="arrow-left">
              <div class="arrow-left-head shadow"></div>
              <div class="arrow-left-tail shadow"></div>
            </div>
          </div>
        </div>
        <div class="reverse-symbol">
          <div class="arrow-right">
            <div class="arrow-right-tail white"></div>
            <div class="arrow-right-head white"></div>
          </div>
          <div class="arrow-left">
            <div class="arrow-left-head white"></div>
            <div class="arrow-left-tail white"></div>
          </div>
        </div>
      </div>

      <div class="reverse-mini reverse-mini-bottom">
        <div class="reverse-symbol shadow-mini">
          <div class="shadow-container">
            <div class="arrow-right">
              <div class="arrow-right-tail shadow"></div>
              <div class="arrow-right-head shadow"></div>
            </div>
            <div class="arrow-left">
              <div class="arrow-left-head shadow"></div>
              <div class="arrow-left-tail shadow"></div>
            </div>
          </div>
        </div>
        <div class="reverse-symbol">
          <div class="arrow-right">
            <div class="arrow-right-tail white"></div>
            <div class="arrow-right-head white"></div>
          </div>
          <div class="arrow-left">
            <div class="arrow-left-head white"></div>
            <div class="arrow-left-tail white"></div>
          </div>
        </div>
      </div>

      <span class="mark">
        <div class="reverse">
          <div class="reverse-symbol shadow">
            <div class="shadow-container">
              <div class="arrow-right">
                <div class="arrow-right-tail shadow"></div>
                <div class="arrow-right-head shadow"></div>
              </div>
              <div class="arrow-left">
                <div class="arrow-left-head shadow"></div>
                <div class="arrow-left-tail shadow"></div>
              </div>
            </div>
          </div>
          <div class="reverse-symbol">
            <div class="arrow-right">
              <div class="arrow-right-tail" :class="this.color.toLowerCase()"></div>
              <div class="arrow-right-head" :class="this.color.toLowerCase()"></div>
            </div>
            <div class="arrow-left">
              <div class="arrow-left-head" :class="this.color.toLowerCase()"></div>
              <div class="arrow-left-tail" :class="this.color.toLowerCase()"></div>
            </div>
          </div>
        </div>
        
      </span>
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
      if(typeof this.number === 'number')
        return `card num-${this.number} ${this.color.toLowerCase()}`;
      else
        return `card ${this.number.toLowerCase()} ${this.color.toLowerCase()}`;
    }
  }
};
</script>

<style scoped>
@import url(https://fonts.googleapis.com/css?family=Source+Sans+Pro:900);

.card {
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
.reverse{
  display: flex;
  align-items: center;
  justify-content: center;
  height: 140px;
  width: 60px;
}
.reverse-mini {
  position: absolute;
  transform: scale(0.2);
}
.reverse-mini-top{
  top: -10px;
  left: -33px;
}
.reverse-mini-bottom{
  bottom: -10px;
  right: -33px;
  /* transform: rotate(180deg); */
}
.reverse-symbol {
  width: 100px;
  height: 55px;
  display: flex;
  flex-direction: column;
  transform: rotate(-45deg);
}
.arrow-right {
  margin-left: 30px;
  display: flex;
  align-items: center;
  width: 70px;
  height: 30px;
}
.arrow-right-tail{
  width: 60px;
  height: 15px;
  background: #0063b3;
  border-radius: 15px 0 0 0
}
.arrow-right-head{
  width: 0;
  height: 0;
  border-top: 20px solid transparent;
  border-left-width: 25px;
  border-left-style: solid;
  border-bottom: 20px solid transparent;
}
.arrow-left {
  margin-top: -13px;
  display: flex;
  align-items: center;
  width: 70px;
  height: 30px;
}
.arrow-left-tail{
  width: 60px;
  height: 15px;
  border-radius: 0 0 15px 0
}
.arrow-left-head{
  width: 0;
  height: 0;
  border-top: 20px solid transparent;
  border-right-width: 25px;
  border-right-style: solid;
  border-bottom: 20px solid transparent;
}   

.reverse-symbol.shadow {
  margin-right: -100px;
}
.shadow-container {
  padding: 5px 0 5px 0;
}
.reverse-symbol.shadow-mini {
  margin-right: -95px;
  margin-bottom: -57px;
}

.arrow-right-tail.shadow{
  background: black;
}
.arrow-right-head.shadow{
  border-left-color: black;
}
.arrow-left-tail.shadow{
  background: black;
}
.arrow-left-head.shadow{
  border-right-color: black;
}   

.arrow-right-tail.white{
  background: white;
}
.arrow-right-head.white{
  border-left-color: white;
}
.arrow-left-tail.white{
  background: white;
}
.arrow-left-head.white{
  border-right-color: white;
}   

.arrow-right-tail.blue{
  background: #0063b3;
}
.arrow-right-head.blue{
  border-left-color: #0063b3;
}
.arrow-left-tail.blue{
  background: #0063b3;
}
.arrow-left-head.blue{
  border-right-color: #0063b3;
}   

.arrow-right-tail.green{
  background: #18a849;
}
.arrow-right-head.green{
  border-left-color: #18a849;
}
.arrow-left-tail.green{
  background: #18a849;
}
.arrow-left-head.green{
  border-right-color: #18a849;
}   

.arrow-right-tail.red{
  background: #c72a18;
}
.arrow-right-head.red{
  border-left-color: #c72a18;
}
.arrow-left-tail.red{
  background: #c72a18;
}
.arrow-left-head.red{
  border-right-color: #c72a18;
}   

.arrow-right-tail.yellow{
  background: #e6ca1e;
}
.arrow-right-head.yellow{
  border-left-color: #e6ca1e;
}
.arrow-left-tail.yellow{
  background: #e6ca1e;
}
.arrow-left-head.yellow{
  border-right-color: #e6ca1e;
}   
/* End Reverse Arrow CSS */

/* Start Skip CSS */
.skip:before,
.skip:after {
  content: "skip";
}
/* End Skip CSS */

/* Start Draw 2 CSS */
.draw_two:before,
.draw_two:after {
  content: "draw 2";
}
/* End Draw 2 CSS */

/* Start Wild CSS */
.wild:before,
.wild:after {
  content: "wild";
}
.mark.wild-mark {
  background-image: linear-gradient(to left, red, yellow, green, blue) !important;
}
/* End Wild CSS */

/* Start Draw 4 CSS */
.wild_draw_four:before,
.wild_draw_four:after {
  content: "draw 4";
}
.mark.wild_draw_four-mark {
  background-image: linear-gradient(to left, red, yellow, green, blue) !important;
}
/* End Draw 4 CSS */

.card.black .inner {
  background: black;
}
</style>