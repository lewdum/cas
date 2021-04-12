<script>
    import { fileContents } from "./stores";
</script>

<div>
    <textarea readonly>{$fileContents}</textarea>
</div>

<style>
    div {
        grid-column: 2;
        grid-row: 1;
    }

    textarea {
        width: 100%;
        height: 100%;
        overflow: scroll;
        border: none;
        outline: none;

        font-family: var(--font-monospace);
        font-size: 14px;
        white-space: pre;
        resize: none;
    }
</style>
