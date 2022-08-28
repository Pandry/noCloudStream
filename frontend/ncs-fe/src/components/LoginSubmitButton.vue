<template>
    <button type="submit">
        {{value}}
    </button>
</template>

<script setup>
const props = defineProps(["value"])
</script>

<style scoped>
    button {
        padding: 10px 30px;
        background-color: hsla(160, 100%, 37%, 1);
        color: white;
        border-radius: 10px 30px;
        margin: auto;
        border-color: transparent;
        left: 70%;
        font-size: 14px;
        font-weight: bold;
    } 
</style>