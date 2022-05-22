async function start() {
    let path = location.hash.replace('#', '')

    let data = await fetch(path, {
        method: 'GET'
    });

    if (data.status !== 200) {
        alert('invalid rom path!');

        return;
    }

    let rom = (await data.body.getReader().read()).value;
    let run = Module.cwrap("run_chip8_b", null, ["number", "number"]);
    let converted_value = new Uint8Array(rom);
    let value_ptr = Module._malloc(rom.length);

    console.log(run);

    Module.HEAPU8.set(converted_value, value_ptr);

    run(value_ptr, rom.length);
}

function toUTF8Array(str) {
    let utf8 = [];
    for (let i = 0; i < str.length; i++) {
        let charcode = str.charCodeAt(i);
        if (charcode < 0x80) utf8.push(charcode);
        else if (charcode < 0x800) {
            utf8.push(0xc0 | (charcode >> 6), 0x80 | (charcode & 0x3f));
        } else if (charcode < 0xd800 || charcode >= 0xe000) {
            utf8.push(0xe0 | (charcode >> 12), 0x80 | ((charcode >> 6) & 0x3f), 0x80 | (charcode & 0x3f));
        } else {
            i++;
            charcode = 0x10000 + (((charcode & 0x3ff) << 10) | (str.charCodeAt(i) & 0x3ff));
            utf8.push(0xf0 | (charcode >> 18), 0x80 | ((charcode >> 12) & 0x3f), 0x80 | ((charcode >> 6) & 0x3f), 0x80 | (charcode & 0x3f));
        }
    }
    return utf8;
}
