<template>
<div id="chatcontainer" class='container'>  

    <!-- Scrollable messaging Card   -->
    <v-card id='chatcard' outlined tile>
        <div class="col">
            <div v-for="(message) in messages" 
                :key="message.id"
                :class="[message.player.id === gameState.player_id ? 'from-me' : 'from-them', 'message']">
                <div class="message-author"><small>{{ message.player.name }}</small><small v-show="message.player.color === undefined"> ... ( watching )</small></div>
                <div :class="[message.player.color != undefined ? message.player.color : '', 'message-content']">{{ message.message }}</div>
            </div>
        </div>
    </v-card>

    <!-- Type and Send Card -->
    <v-card id='message-box' outlined tile>
        <div class="container">
            <div v-if="!displayInfo" class="row">
                <div class="col-9" @keypress.enter="sendMessage">
                    <!-- <v-on  -->
                    <v-text-field
                        @keydown.enter="sendMessage"
                        autofocus
                        label="Message"
                        outlined
                        clearable
                        v-model="newMessage"
                    > </v-text-field>
                    <!-- </v-on> -->
                </div>
                <div class="col-3">
                    <v-card-actions>
                        <v-btn id="send-btn" color="blue darken-1" text @click.native="sendMessage">Send</v-btn>
                    </v-card-actions>
                </div>
            </div>
            <div v-else id="message-info" class="text-center">
                <p> {{ info }} </p>
            </div>
        </div>
    </v-card>
</div>
</template>

<script> 
import unoService from '../services/unoService';

export default {
    name: "Chat",
    components: {

    },
    props: ['gameState'],
    data() {
        return {
            newMessage: '',
            messages: [],

            info: null,
            displayInfo: false,
            
            players: [],
            // It would be best to know a max number of people able to play
            // TODO: get rid of the blue option similar to the `from-me`
            messageColors: ['success', 'warning', 'indigo', 'purple', 'pink', 'red', 'orange', 'yellow', 'green', 'teal', 'cyan'],

            loop_scroll: true,
            skipfirstGameState: false,
        }
    },
    watch: {
        gameState(newVal, oldVal) {
            if (this.skipfirstGameState) {
                this.messages = newVal.messages;
                this.players = newVal.all_players;

                // There are .length errors if the messages is null
                if (this.messages != null) {
                    // Assign Message Colors to the players
                    for (var i = 0; i < this.players.length; i++) {
                        for (var j = 0; j < this.messages.length; j++) {
                            if (this.messages[j].player.id === this.players[i].id) {
                                this.messages[j].player.color = this.messageColors[i]
                            }
                        }
                    }

                    let len__new = this.messages.length - 1;
                    var len__old;
                    try {
                        len__old = oldVal.messages.length - 1;
                    }catch(err){
                        len__old = len__new
                    }

                    // If we have a new message scroll down and tell the game
                    if ( len__new !== len__old ) {
                        this.loop_scroll = true;
                        this.$emit('snackbarText', this.messages[len__new].player.name, this.messages[len__new].message)
                    }
                }
                this.info = null;
                this.displayInfo = false;
            }
            this.skipfirstGameState = true;
        }
    },
    methods: {
        scroll() {
            var div = document.getElementById('chatcard');
            div.scrollTop = div.scrollHeight - div.clientHeight;
        },
        async sendMessage() {
            if (this.newMessage != "") {
                this.info = 'Sending to Server';
                this.displayInfo = true;
                let res = await unoService.sendMessage(this.$route.params.id, this.gameState.current_player.id, this.newMessage);
                if (res.data) {
                    this.newMessage = '';
                }
            }else{
                this.info = 'Message is Empty, No message will be sent';
                this.displayInfo = true;
                this.loop_scroll = true;
            }
        }
    },
    created() {
        this.scrollInterval = setInterval(() => {
            if (this.loop_scroll) {
                this.scroll();
                this.loop_scroll = false
            }
        }, 100);
    },
    beforeDestroy (){
        if(this.scrollInterval){
            clearInterval(this.scrollInterval);
        }
    }
}
</script>


<style scoped>

#chatcontainer {
    height: 100%;
    width: 100%;
    
    font-family: "Helvetica Neue", Helvetica, sans-serif;
	font-size: 15px;
	font-weight: normal;
    max-width: 450px;
	margin: auto;
    display: flex;
    flex-direction: column;
}

#chatcard {
    height: 100%;
    width: 100%;
    object-fit: cover;
    display: flex;
    overflow: auto;
    max-width: 100%;
    overflow-x: hidden;
    overflow-anchor: auto;
}

#chatcard::-webkit-scrollbar {
    width: 0px;
    background: transparent;
} 

#message-box {
    height: 100px;
}

#message-info {
    height: 100px;
}

.message {
    width: 100%;
    min-height: 80px;
}

.message-content {
  max-width: 255px;
  word-wrap: break-word;
  margin-bottom: 12px;
  line-height: 24px;
  position:relative;
  padding:10px 20px;
  border-radius:25px;
}

.from-me > .message-author {
    text-align: right;
}

.from-me > .message-content {
	color:white; 
	background:#0B93F6;
    float: right;
}

.from-them > .message-content{
	background:#E5E5EA;
	color:black;
    float: left;
}
</style>