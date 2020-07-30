import BaseService from "./baseService";

export default {
    setToken(token) {
        BaseService.defaults.headers.common['Authorization'] = `Bearer ${token}`;
    },

    async login(userName) {
        return BaseService.post(`/login/${userName}`);
    },

    async newGame() {
        return BaseService.get(`/newgame/`);
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