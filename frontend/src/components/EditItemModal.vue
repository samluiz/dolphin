<script setup lang="ts">
import { Ref, ref, toRefs } from 'vue';
import Dialog from './ui/Dialog.vue';
import Input from './ui/Input.vue';
import { Item } from '@/shared/types';

const props = defineProps({
    isOpen: Boolean,
    selectedTabId: Number,
    current: Object
})

const { isOpen, selectedTabId, current } = toRefs(props)

const emit = defineEmits(["editItem", "on-cancel"]);

const formData: Ref<Item> = ref(
    current?.value as Item,
)

function onSubmit(): void {
    formData.value.amount = Number(formData.value.amount);
    emit("editItem", formData.value);
}

</script>

<template>
    <Dialog :isOpen="isOpen">
        <form class="flex flex-col gap-4" @submit.prevent="onSubmit">
            <h1 class="text-2xl">Edit {{ selectedTabId === 1 ? 'earning' : 'expense' }}</h1>
            <Input :required="true" :name="'description'" :title="'Description'" :type="'text'" :model-value="formData.description" @update:model-value="newValue => formData.description = newValue" />
            <Input :required="true" :name="'amount'" :title="'Amount'" :type="'number'" :model-value="formData.amount" @update:model-value="newValue => formData.amount = newValue" />
            <Input :required="true" v-if="selectedTabId === 2" :name="'category'" :title="'Category'" :type="'text'" :model-value="formData.category" @update:model-value="newValue => formData.category = newValue" />
            <div class="flex gap-4">
                <button @click="$emit('on-cancel')" type="button" class="bg-gray-500 text-white px-4 py-2 rounded-lg">Cancel</button>
                <button type="submit" class="bg-blue-500 text-white px-4 py-2 rounded-lg">Edit</button>
            </div>
        </form>
    </Dialog>
</template>