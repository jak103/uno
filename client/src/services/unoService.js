import BaseService from "./baseService";

export default {
  async getAllGames() {
    return BaseService.get(`/api/games`);
  },

  async newGame(gameName, creatorName) {
    return BaseService.post(`/api/games`, {name: gameName, creator: creatorName});
  },

  async joinGame(gameId, playerName) {
    return BaseService.post(`/api/games/${gameId}/join`, { playerName: playerName });
  },

  getGameState(gameId) {
    return BaseService.get(`/api/games/${gameId}`);
  },

  startGame() {
    return BaseService.post(`/api/startgame`);
  },

  playCard(cardNumber, cardColor) {
    return BaseService.post(`/api/play/${cardNumber}/${cardColor}`);
  },

  drawCard() {
    return BaseService.post(`/api/draw`);
  },

}