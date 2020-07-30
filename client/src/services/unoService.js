import BaseService from "./baseService";

export default {
    async login(userName) {
        return BaseService.post(`/api/login/${userName}`);
    },

    async newGame() {
        return BaseService.get(`/api/newgame/`);
    },

    getGameState(gameId) {
        return BaseService.get(`/api/game/${gameId}`);
    },

    startGame() {
        return BaseService.post(`/api/startgame`);
    },

    playCard(cardNumber, cardColor) {
        return BaseService.post(`/api/play/${cardNumber}/${cardColor}`);
    },

    drawCard() {
        return BaseService.post(`/api/draw`);
    }
}