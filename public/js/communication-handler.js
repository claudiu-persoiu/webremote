var communicationHandler = function (url) {
    var socket = new WebSocket(url), socketOpened = false;

    socket.onopen = function (e) {
        socketOpened = true;
    };

    socket.onmessage = function (event) {
        console.log(event.data);
    };

    socket.onclose = function () {
        socketOpened = false;
    };

    return {
        send: function (message) {
            if (socketOpened) {
                return socket.send(JSON.stringify(message));
            }
            return false;
        },
        isConnected: function () {
            return socketOpened;
        }
    };
};
