<script setup lang="ts">
import { toRefs } from 'vue';
import Dialog from './ui/Dialog.vue';

const props = defineProps({
    isOpen: Boolean
})

const { isOpen } = toRefs(props)
</script>

<template>
    <div v-if="isOpen" class="fixed z-50">
        <Dialog>
            <div class="flex flex-col gap-4">
                <h1 class="text-lg">Are you sure you want to delete this item? This action cannot be undone.</h1>
                <div class="flex gap-4">
                    <button @click="$emit('on-cancel')" class="bg-gray-500 text-white px-4 py-2 rounded-lg">Cancel</button>
                    <button @click="$emit('on-confirm')" class="bg-red-500 text-white px-4 py-2 rounded-lg">Delete</button>
                </div>
            </div>
        </Dialog>
    </div>
</template>