<script>
    import { knownFiles, fileContents } from "./stores";

    const endpoint = "http://localhost:8080";

    let form;

    let files;
    $: ready = files && files.length > 0;
    $: selected = ready ? files[0] : null;

    async function upload(event) {
        event.preventDefault();

        fileContents.clear();
        fileContents.add("Enviando...");

        let response;

        try {
            response = await fetch(endpoint, {
                method: "POST",
                body: new FormData(form),
            });
        } catch (err) {
            fileContents.add(`Erro de conexão: ${err}`);
            return;
        }

        if (response.status === 409) {
            const key = response.headers.get("Key");

            fileContents.add("Esse arquivo já existe no servidor.");
            fileContents.add();
            fileContents.add("Chave de acesso:");
            fileContents.add(key);

            knownFiles.add(key);

            return;
        }

        if (response.status !== 200) {
            fileContents.add(`Erro: ${await response.text()}`);
            return;
        }

        const key = await response.text();

        fileContents.add(`Arquivo enviado com sucesso.`);
        fileContents.add();
        fileContents.add("Chave de acesso:");
        fileContents.add(key);

        knownFiles.add(key);

        // TODO: show file
    }
</script>

<div class="outer">
    <form on:submit={upload} bind:this={form}>
        <h2>Novo Arquivo</h2>

        <div class="inner">
            <input name="data" type="file" bind:files />
            <input type="submit" value="Enviar" disabled={!ready} />
        </div>
    </form>

    {#if ready}
        <table>
            <tr>
                <td><strong>Size:</strong></td>
                <td>{selected.size} bytes</td>
            </tr>
            <tr>
                <td><strong>Type:</strong></td>
                <td>{selected.type || "?"}</td>
            </tr>
        </table>
    {/if}
</div>

<style>
    .outer {
        grid-column: 2;
        grid-row: 2;

        background: linear-gradient(
            to right,
            var(--light-blue),
            var(--dark-blue)
        );
        border-top: 1px solid gray;
        padding: 20px;
    }

    .inner {
        width: 100%;
        display: grid;
        grid-template-columns: auto 80px;
        column-gap: 5px;
    }

    h2 {
        margin-top: 0;
    }

    input[type="file"] {
        grid-column: 1;

        line-height: 24px;
        border: 1px solid gray;
    }

    input[type="submit"] {
        grid-column: 2;

        line-height: 24px;
    }

    table tr:nth-child(1) {
        text-align: right;
    }
</style>
