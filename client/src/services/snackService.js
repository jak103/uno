import BaseService from "./baseService";

export default {
    async sendMessage(gameID, message, showSnack) {
        //console.log('snack status:', showSnack);
        return BaseService.post(`/api/games/${gameID}/sendMessage`, {value: message, showSnack});
    }
}