window.onload = function () {
    var audio = new Audio('dmrdrn2.mp3');
    initAudio(audio);
    if (window["WebSocket"]) {
        const wsAddr = "ws://" + document.location.host + "/ws";
        console.log("opening " + wsAddr);
        const conn = new WebSocket(wsAddr);
        conn.onclose = function (evt) {
            console.log(evt.data);
            console.log("Connection closed.");
        };
        conn.onmessage = function (evt) {
            console.log(evt.data);
            if(evt.data != "tick"){
                audio.play();
            }
        };
        conn.onerror = function (evt) {
            console.log("error!");
            console.log(evt);
        }
        console.log(conn);
    } else {
        console.log("Your browser does not support WebSockets.");
    }
};
function initAudio(audio) {
    audio.addEventListener('play', function () {
        // When the audio is ready to play, immediately pause.
        audio.pause();
        audio.removeEventListener('play', arguments.callee, false);
    }, false);
    document.addEventListener('click', function () {
        // Start playing audio when the user clicks anywhere on the page,
        // to force Mobile Safari to load the audio.
        document.removeEventListener('click', arguments.callee, false);
        audio.play();
    }, false);
}