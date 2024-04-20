<script setup lang="ts">

import { defineProps } from 'vue';

interface Props {
    name: string;
    title: string;
    type: string;
    required: boolean;
    modelValue: any;
}

defineProps<Props>();

</script>

<template>
    <div class="flex flex-col justify-start items-center gap-2">
        <label class="w-full text-sm" :for="name">{{ title }}</label>
        <input class="w-full p-1 rounded-sm shadow-lg bg-slate-600" :required="required" :type="type" :id="name" :value="modelValue" @input="$emit('update:modelValue', ($event.target as HTMLInputElement)?.value)" />
    </div>
</template>