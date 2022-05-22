const blockList = document.getElementById('block-list');
const sendRomButton = document.getElementById('send-rom-button');
const romName = document.getElementById('rom-name');
const romAuthor = document.getElementById('rom-author');
const romSecret = document.getElementById('rom-secret');
const romFile = document.getElementById('rom-file');

async function apiList() {
    return await fetch(`/api/demo/list?page_size=10000&page_num=0`, {
        method: 'GET'
    }).then(r => r.json());
}

async function apiCreate(name, author, secret, rom) {
    let formData = new FormData();
    formData.append("name", rom);

    return await fetch('/api/demo', {
        method: 'POST',
        body: formData,
        headers: new Headers({'X-Svm-Name': btoa(name), 'X-Svm-Author': btoa(author), 'X-Svm-Secret': btoa(secret)})
    });
}

async function createLink(demo) {
    const div = document.createElement('div')
    const link = document.createElement('a');

    link.innerText = atob(demo["name"]);
    link.href = `/play.html?#${demo["rom_path"]}`;
    link.hidden = false;

    div.appendChild(link);

    return div;
}

async function initList() {
    const demos = await apiList();

    for (const demo of demos["demos"]) {
        const link = await createLink(demo);

        blockList.appendChild(link);
    }

    blockList.hidden = false;
}

async function onSendClick() {
    let name = romName.value;
    let author = romAuthor.value;
    let secret = romSecret.value;

    if (name === '' || author === '' || secret === '' || romFile.files.length === 0) {
        alert('You must fill all fields!');

        return;
    }

    let rom = romFile.files[0];

    let response = await apiCreate(name, author, secret, rom);

    if (response.status !== 200) {
        alert(`Demo creation failed with status ${response.status}`);

        return;
    }

    prompt("Success! Dont forget your access token for this rom!", response.headers.get('X-Svm-Key'));
}

async function init() {
    sendRomButton.addEventListener('click', onSendClick);

    await initList();
}

init()
