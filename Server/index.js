
window.onload = function () {
    if (window["WebSocket"]) {
        const wsAddr = "ws://" + document.location.host + "/ws";
        console.log("opening " + wsAddr);
        const conn = new WebSocket(wsAddr);
        conn.onclose = function (evt) {
            console.log(evt.data);
            console.log("Connection closed.");
        };
        conn.onmessage = function (evt) {
            if(evt.data != "tick"){
                console.log(evt.data);
                var audio = new Audio('dmrdrn2.mp3');
                audio.play();
            }
        };
    } else {
        console.log("Your browser does not support WebSockets.");
    }
};