import BaseService from "./baseService";

export default {
    async notifyAll(gameID, notification) {
        return BaseService.post(`/api/snack/${gameID}/notify`, {notification: notification});
    }
}