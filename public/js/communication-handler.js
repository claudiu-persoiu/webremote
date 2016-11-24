var communicationHandler = function (url, errorBlock) {
    var socket, socketOpened = false;

    var connect = function () {
        socket = new WebSocket(url);

        socket.onopen = function () {
            socketOpened = true;
            errorBlock.style.display = 'none';
        };

        socket.onmessage = function (event) {
            console.log(event.data);
        };

        socket.onclose = function () {
            socketOpened = false;
            errorBlock.style.display = 'block';
        };
    };

    connect();

    return {
        send: function (message) {
            if (socketOpened) {
                return socket.send(JSON.stringify(message));
            }
            return false;
        },
        isConnected: function () {
            return socketOpened;
        },
        reconnect: function () {
            connect();
        }
    };
};
