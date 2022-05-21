const blockList = document.getElementById('block-list');
const pageSizeBlock = document.getElementById('page_size')
const pageNumBlock = document.getElementById('page_num')

async function apiList(pageSize, pageNum) {
    return await fetch(`/api/demo/list?page_size=${pageSize}&page_num=${pageNum}`, {
        method: 'GET'
    }).then(r => r.json());
}

async function createLink(id) {
    const link = document.createElement('div');

    link.className = 'link';
    link.innerText = id;

    return link;
}

async function initList() {
    const demos = await apiList(parseInt());

    console.log(demos);

    for (const demo of demos) {
        console.log(demo);
        // const link = await createLink(demo);

        // blockList.appendChild(link);
    }

    blockList.hidden = false;
}

async function init() {
    await initList();
}

init()
