<script>
    import List from "./List.svelte";
    import Data from "./Data.svelte";
    import Upload from "./Upload.svelte";
</script>

<div class="grid">
    <List />
    <Data />
    <Upload />
</div>

<style>
    .grid {
        width: 100%;
        height: 100%;
        display: grid;
        grid-template-columns: minmax(300px, min-content) auto;
        grid-template-rows: auto min-content;
    }
</style>
