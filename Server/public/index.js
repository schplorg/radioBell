var audios = [];
var bell = new Audio("doorbell.mp3");
var vid = document.createElement("video");
var but = document.createElement("button");
window.onload = function () {
    document.body.setAttribute("style","background: black; margin: 0");

    vid.src = "animated-logo.mp4";
    vid.setAttribute("style","position: absolute;top: 50%;left: 50%;transform: translate(-50%, -50%);");
    document.body.append(vid);


    let t = document.createElement("p");
    t.innerText = "▶️";
    t.setAttribute("style","position: absolute;top: 50%;left: 50%;transform: translate(-50%, -50%); color: white; font-size: 50vw; margin: 0;");
    but.append(t);

    but.setAttribute("style","position: absolute; width: 100%; height:100%; margin: 0; background: none; border: 0; padding: 0;");
    document.body.append(but);
    initAudio();
    initWebSocket();
    initVideo();
};
function initVideo(){
    document.addEventListener('click', function () {
        vid.play();
    }, false);
    vid.addEventListener('play', function () {
        // When the audio is ready to play, immediately pause.
        vid.pause();
        vid.removeEventListener('play', arguments.callee, false);
    }, false);
}
var presses = 0;
var last = -1;
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
                    while(ran == last){
                        ran = Math.floor(Math.random() * audios.length);
                    }
                    last = ran;
                    let audio = audios[ran];
                    audio.currentTime = 0
                    audio.play();
                    console.log("play " + audio.src);
                    presses = Math.min(presses+1,15);
                }else{
                    bell.currentTime = 0
                    bell.play();
                    console.log("play bell");
                    vid.currentTime = 0;
                    vid.play();
                    presses = Math.min(presses+5,15);
                }
            }else{
                presses = Math.max(presses-1,0);
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
        bell.pause();
        console.log("paused bell");
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
        "waow.mp3",      
    ];
    for(let s of aPaths){
        let audio = new Audio(s);
        audios.push(audio);
        audio.addEventListener('play', function () {
            audio.pause();
            console.log("paused " +s);
            audio.removeEventListener('play', arguments.callee, false);
        }, false);
    }
    document.addEventListener('click', function () {
        but.style["display"] = "none";
        document.removeEventListener('click', arguments.callee, false);
        bell.play();
        for(let a of audios){
            a.play();
        }
    }, false);
}