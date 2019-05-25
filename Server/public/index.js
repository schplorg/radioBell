var audios = [];
var bell = new Audio("doorbell.mp3");
window.onload = function () {
    document.body.setAttribute("style","background: black; margin: 0");
    initAudio();
    initWebSocket();
    initVideo();
};
function initVideo(){
    let vid = document.createElement("video");
    document.body.append(vid);
    vid.src = "animated-logo.mp4";
    vid.loop = true;
    vid.setAttribute("style","position: absolute;top: 50%;left: 50%;transform: translate(-50%, -50%);");
    document.addEventListener('click', function () {
        vid.play();
    }, false);
}
var presses = 0;
function initWebSocket(){
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
                if(presses > 10){
                    let ran = Math.floor(Math.random() * audios.length);
                    audios[ran].currentTime = 0
                    audios[ran].play();
                }else{
                    bell.currentTime = 0
                    bell.play();
                    presses += 3;
                }
            }else{
                presses--;
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
}
function initAudio(audio) {
    bell.addEventListener('play', function () {
        // When the audio is ready to play, immediately pause.
        bell.pause();
        bell.removeEventListener('play', arguments.callee, false);
    }, false);
    let aPaths = [
        "airhorn.mp3",
        "bark.mp3",
        "bark2.mp3",
        "bark3.mp3",
        "cat.mp3",
        "dmdrn1.mp3",
        "dmrdrn2.mp3",
        "duck.mp3",
        "duck2.mp3",
        "honk.mp3",
        "jeopardy.mp3",
        "waow.mp3",      
    ];
    for(let s of aPaths){
        let audio = new Audio(s);
        audios.push(audio);
        audio.addEventListener('play', function () {
            // When the audio is ready to play, immediately pause.
            audio.pause();
            audio.removeEventListener('play', arguments.callee, false);
        }, false);
    }
    document.addEventListener('click', function () {
        // Start playing audio when the user clicks anywhere on the page,
        // to force Mobile Safari to load the audio.
        document.removeEventListener('click', arguments.callee, false);
        bell.play();
        for(let a in audio){
            a.play();
        }
    }, false);
}