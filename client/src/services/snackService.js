import BaseService from "./baseService";

export default {

    //example of how to use 
    //snackService.notifyAll(this.$route.params.id, "Wild Played!!");
    //TODO: make work when called multiple times in a row with same message.
    //currently requires different messages in order to show the snack
    async notifyAll(gameID, notification) {
        console.log('the notification: ', notification);
        return BaseService.post(`/api/snack/${gameID}/notify`, {notification: notification});
    }
}