var socket = new WebSocket("ws://localhost:8080/ws");

let connect = (callback: { (msg: string): void }) => {
    console.log("Attempting Connection...");

    socket.onopen = () => {
        console.log("Successfully Connected");
    };

    socket.onmessage = event => {
        console.log(event);
        callback(event.data)
    };

    socket.onclose = event => {
        console.log("Socket Closed Connection: ", event);
    };

    socket.onerror = error => {
        console.log("Socket Error: ", error);
    };
};

let sendMsg = (msg: string) => {
    console.log("sending msg: ", msg);
    socket.send(msg);
};

export { connect, sendMsg };