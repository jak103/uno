import BaseService from "./baseService";

export default {
    async login(gameId, userName) {
        return BaseService.post(`/login/${gameId}/${userName}`);
    },

    async newGame() {
        return BaseService.get("/newgame");
    },

    update(gameId, userName) {
        return BaseService.get(`/update/${gameId}/${userName}`);
    },

    getHandCount(gameId) {
        return BaseService.get(`/cardcount/${gameId}`);
    },

    startGame(gameId, userName) {
        return BaseService.post(`/startgame/${gameId}/${userName}`);
    },

    playCard(gameId, userName, cardNumber, cardColor) {
        return BaseService.post(`/play/${gameId}/${userName}/${cardNumber}/${cardColor}`);
    },

    drawCard(gameId, userName) {
        return BaseService.post(`/draw/${gameId}/${userName}`);
    }
}