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

  async getGameState(gameId) {
    return BaseService.get(`/api/games/${gameId}`);
  },

  async getPlayerNameFromToken() {
    let token = localStorage.getItem('token')
    if (token) {
      return BaseService.get(`/api/players/token/${token}`);
    }else {
      return ""
    }
  },

  async drawCard(gameId) {
    return BaseService.post(`/api/games/${gameId}/draw`);
  },
  
  async playCard(gameId, cardValue, cardColor) {
    return BaseService.post(`/api/games/${gameId}/play`, {value: cardValue, color: cardColor});
  },
  
  async startGame(gameId) {
    return BaseService.post(`/api/games/${gameId}/start`);
  },

  async gotoHelp(tag) {
    return BaseService.post(`/help${tag}`)
  },
  
  async sendMessage(gameId, playerId, message) {
    return BaseService.post(`/api/chat/${gameId}/add`, { player: playerId, message: message});
  },

  async callUno(gameId, calledOnPlayerId) {
    console.log(`gameId`, gameId);
    console.log(`calledOnPlayerId`, calledOnPlayerId);
    return BaseService.post(`/api/games/${gameId}/call`, calledOnPlayerId);
  },

}
