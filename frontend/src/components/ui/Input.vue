<script setup lang="ts">

defineProps({
    name: String,
    title: String,
    type: String,
})

</script>

<template>
    <div class="flex flex-col justify-start items-center gap-2">
        <label class="w-full text-sm" for="{{ name }}">{{ title }}</label>
        <input class="w-full p-1 rounded-sm shadow-lg bg-slate-600" type="{{ type }}" id="{{ name }}" />
    </div>
</template>