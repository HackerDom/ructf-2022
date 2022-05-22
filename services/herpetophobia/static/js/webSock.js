let sock = new WebSocket(`ws://${window.location.host}/play`);
sock.onopen = function (event) {
    let id = document.getElementById("game_id").innerText;
    sock.send(JSON.stringify({"id": id}));
    let msg = JSON.parse(event.data);
    if (msg.gameMap) {
        makeMap(msg.gameMap);
        makeInfo(msg);
    }
};

sock.onmessage = function (event) {
    let msg = JSON.parse(event.data);
    if (msg.gameMap) {
        makeMap(msg.gameMap);
        makeInfo(msg);
    } else {
        makeRes(msg);
    }
};

sock.onerror = function (event) {
    alert(event.data);
};
document.addEventListener('keydown', function (event) {
    let _direction = '';
    if (event.code === 'KeyD') {
        _direction = 'd';
    } else if (event.code === 'KeyS') {
        _direction = 's';
    } else if (event.code === 'KeyA') {
        _direction = 'a';
    } else if (event.code === 'KeyW') {
        _direction = 'w';
    }
    if (_direction) {
        sock.send(JSON.stringify({direction: _direction, closeGame: false, newGame: true}));
    }
});

function makeMap(map) {
    for (let i = 0; i < 16; i++) {
        for (let j = 0; j < 16; j++) {
            if (map[i][j] === '#') {
                document.getElementById(`${i}I${j}`).style.backgroundColor = 'green'
            } else if (map[i][j] === '@') {
                document.getElementById(`${i}I${j}`).style.backgroundColor = 'red';
            } else if (map[i][j] === '*') {
                document.getElementById(`${i}I${j}`).style.backgroundColor = 'blue';
            } else {
                document.getElementById(`${i}I${j}`).style.backgroundColor = 'white'
            }
        }
    }
}

function makeRes(res) {
    document.getElementById('info').innerText = `result: ${res.gameResult}\ncounter: ${res.counter}`;
}

function makeInfo(msg) {
    document.getElementById('info').innerText = `step: ${msg.step}\ncounter: ${msg.counter}`;
}