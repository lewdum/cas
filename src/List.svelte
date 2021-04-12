<script>
    import { knownFiles, fileContents } from "./stores";

    const endpoint = "http://localhost:8080";

    let query = "";

    async function show(key) {
        fileContents.clear();
        fileContents.add("Pesquisando...");

        let response;

        const path = `${endpoint}/${key}`;

        try {
            response = await fetch(path);
        } catch (err) {
            fileContents.add(`Erro de conexão: ${err}`);
            return;
        }

        if (response.status === 404) {
            fileContents.add("Arquivo não encontrado.");
            return;
        }

        if (response.status !== 200) {
            fileContents.add(`Erro: ${response.body.text}`);
            return;
        }

        fileContents.clear(await response.text());

        knownFiles.add(key);
    }

    async function search(event) {
        if (event.key !== "Enter") {
            return;
        }
        show(query);
    }

    async function updateIndexQuietly() {
        let response;

        const path = `${endpoint}/list`;

        try {
            response = await fetch(path);
        } catch (err) {
            return;
        }

        if (response.status !== 200) {
            return;
        }

        const text = await response.text();
        for (let file of text.split("\n")) {
            knownFiles.add(file);
        }
    }

    updateIndexQuietly();
</script>

<div>
    <input
        type="search"
        on:keyup={search}
        bind:value={query}
        placeholder="Digite uma chave aqui para pesquisar."
    />

    <ul>
        {#each $knownFiles as file}
            <li>
                <a href="#" on:click={() => show(file)}>
                    {file}
                </a>
            </li>
        {/each}
    </ul>
</div>

<style>
    div {
        grid-column: 1;
        grid-row: 1 / span 2;

        background: linear-gradient(
            to bottom,
            var(--light-blue),
            var(--light-purple)
        );
        border-right: 1px solid gray;
        font-size: 14px;
        overflow-x: auto;
    }

    input {
        margin: 10px;
        width: calc(100% - 20px);
    }

    ul {
        list-style: none;
        padding-left: 10px;
        padding-right: 10px;
        margin: 0;
        font-family: var(--font-monospace);
    }

    a {
        display: inline-block;
        overflow: hidden;
        text-overflow: ellipsis;
        width: 100%;
    }
</style>
