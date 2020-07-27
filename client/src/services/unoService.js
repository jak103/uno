import BaseService from "./baseService";

export default {
    setToken(token) {
        BaseService.defaults.headers.common['Authorization'] = `Bearer ${token}`;
    },

    async login(gameId, userName) {
        return BaseService.post(`/login/${gameId}/${userName}`);
    },

    async newGame(userName) {
        return BaseService.get(`/newgame/${userName}`);
    },

    update() {
        return BaseService.get(`/update`);
    },

    startGame() {
        return BaseService.post(`/startgame`);
    },

    playCard(cardNumber, cardColor) {
        return BaseService.post(`/play/${cardNumber}/${cardColor}`);
    },

    drawCard() {
        return BaseService.post(`/draw`);
    }
}