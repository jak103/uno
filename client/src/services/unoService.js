import BaseService from "./baseService";

export default {
  async getAllGames() {
    return BaseService.get(`/games`);
  },

  async newGame(gameName, creatorName) {
      return BaseService.post(`/games`, {name: gameName, creator: creatorName});
  },

  async login(userName) {
      return BaseService.post(`/api/login/${userName}`);
  },

  getGameState(gameId) {
      return BaseService.get(`/api/games/${gameId}`);
  },

  startGame(gameId) {
      return BaseService.post(`/api/games/${gameId}/start`);
  },

  playCard(cardNumber, cardColor) {
      return BaseService.post(`/api/play/${cardNumber}/${cardColor}`);
  },

  drawCard() {
      return BaseService.post(`/api/draw`);
  },

}