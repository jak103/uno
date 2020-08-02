<template>
<div id="chatcontainer" class='container'>    
    <v-card id='chatcard' outlined tile>
        <div class="col">
            <div v-for="message in messages" 
             :key="message.id"
             :class="[message.name === gameState.current_player.name ? 'from-me' : 'from-them', 'message']">
                <div class="message-author"><small>{{ message.name }}</small></div>
                <div :id="'message-'+message.id" :class="[message.name !== gameState.current_player.name ? messageColors[message.id] : '', 'message-content']">{{ message.message }}</div>
            </div>
        </div>
    </v-card>
    <v-card id='message-box' outlined tile>
        <div class="container">
            <div class="row">
                <div class="col-9">
                    <v-text-field
                        label="Message"
                        outlined
                        clearable
                        v-model="newMessage"
                    > </v-text-field>
                </div>
                <div class="col-3">
                    <v-card-actions>
                        <v-btn color="blue darken-1" text @click="sendMessage">Send</v-btn>
                    </v-card-actions>
                </div>
            </div>
        </div>
    </v-card>
</div>
</template>

<script> 
// import unoService from '../services/unoService';
// import localStorage from '../util/localStorage';

export default {
    name: "Chat",
    components: {

    },
    props: ['gameState'],
    data() {
        return {
            newMessage: '',
            messages: [],
            messageColors: ['primary', 'success', 'warning', 'indigo', 'purple', 'pink', 'red', 'orange', 'yellow', 'green', 'teal', 'cyan',
                            'primary', 'success', 'warning', 'indigo', 'purple', 'pink', 'red', 'orange', 'yellow', 'green', 'teal', 'cyan'],
        }
    },
    methods: {
        scroll() {
            var div = document.getElementById('chatcard');
            div.scrollTop = div.scrollHeight - div.clientHeight;
        },
        sendMessage() {
            // unoService.sendMessage(this.$route.params.id, this.gameState.current_player);
            console.log('Added a new message: ', { name: this.gameState.current_player.name, message: this.newMessage, });
            this.messages.push({ name: this.gameState.current_player.name, message: this.newMessage, });
            this.newMessage = '';
            this.scroll()
        }
    },
    async created() {
        // let res = await unoService.getChat(this.$route.params.id);
        // this.messages = res.data

        // Ideas [This would need to update every interval in Game.vue]
        // this.messages = this.gameState.messages

        // Fake Data
        this.messages = [
            { id: 0, name: 'Matthew',  message: 'What\'s up guys!', },
            { id: 1, name: 'Patrick',  message: 'Nothing much how are you?', },
            { id: 2, name: 'Benjamin', message: 'Hey guys I\'m good!', },
            { id: 3, name: 'Andrew',   message: 'I\'m watching Doctor Who', },
            { id: 4, name: 'Andrew',   message: 'I\'m going to win this game', },
            { id: 5, name: 'Benjamin', message: 'Doubt full', },
            { id: 6, name: 'fjkjlk', message: 'Haha', },
            { id: 7, name: 'fjgjlk', message: 'Haha', },
            { id: 8, name: 'fjghkk', message: 'Haha', },
            { id: 9, name: 'fjghlk', message: 'Haha', },
            { id: 10, name: 'fjkjlk', message: 'Haha', },
            { id: 11, name: 'fjghlk', message: 'Haha', },
            { id: 12, name: 'fhkjl', message: 'Haha', },
            { id: 13, name: 'fhkjl', message: 'Haha', },
            { id: 14, name: 'fjgjlk', message: 'Haha', },
        ]
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
    /* width: 100%; */
    /* object-fit: cover; */
    /* display: flex; */
    /* max-width: 100%; */
}
.message {
    width: 100%;
    min-height: 80px;
}


@use postcss-nested;


.message-content {
  max-width: 255px;
  word-wrap: break-word;
  margin-bottom: 12px;
  line-height: 24px;
  position:relative;
  padding:10px 20px;
  border-radius:25px;
}

.message-content::before, :after {
    content:"";
    position:absolute;
    bottom:-2px;
    height:20px;
}



.from-me > .message-author {
    text-align: right;
}

.from-me > .message-content {
	color:white; 
	background:#0B93F6;
    float: right;
    /* width:260px; */

}

.from-me::before {
    right:-7px;
    /* border-right:20px solid #0B93F6; */
    border-bottom-left-radius: 16px 14px;
    transform:translate(0, -2px);
}

.from-me::after {
    right:-56px;
    width:26px;
    border-bottom-left-radius: 10px;
    transform:translate(-30px, -2px);
}


.from-them > .message-content{
	background:#E5E5EA;
	color:black;
}

.from-them::before {
    left:-7px;
    /* border-left:20px solid #E5E5EA; */
    border-bottom-right-radius: 16px 14px;
    transform:translate(0, -2px);
}

.from-them::after {
    left:4px;
    width:26px;
    border-bottom-right-radius: 10px;
    transform:translate(-30px, -2px);
}

</style>