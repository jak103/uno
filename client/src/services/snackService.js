import BaseService from "./baseService";

export default {
    async notifyAll(gameID, notification) {
        console.log('the notification: ', notification);
        return BaseService.post(`/api/snack/${gameID}/notify`, {notification: notification});
    }
}