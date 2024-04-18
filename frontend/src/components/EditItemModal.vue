<script setup lang="ts">
import { ref, toRefs } from 'vue';
import Dialog from './ui/Dialog.vue';
import Input from './ui/Input.vue';

const props = defineProps({
    isOpen: Boolean,
    selectedTabId: Number,
})

const { isOpen, selectedTabId } = toRefs(props)

console.log(selectedTabId?.value)

const emit = defineEmits(["updateCurrentItem", "on-cancel"]);

const description = ref('');
const amount = ref(0);
const category = ref('');

const updateValue = (value: any, field: string) => {
    switch (field) {
        case 'description':
            description.value = value;
            break;
        case 'amount':
            amount.value = value;
            break;
        case 'category':
            category.value = value;
            break;
        default:
            break;
    }
};

function onSubmit(): void {
    console.log(description.value, amount.value, category.value)
  emit("updateCurrentItem", description.value, amount.value, category.value);
}

</script>

<template>
    <div v-if="isOpen" class="fixed z-50">
        <Dialog>
            <div class="flex flex-col gap-4">
                <h1 class="text-2xl">Edit earning</h1>
                <Input @input="updateValue($event.target.value, 'description')" name="description" title="Description" type="text" v-model="description" />
                <Input @input="updateValue($event.target.value, 'amount')" name="amount" title="Amount" type="number" v-model="amount" />
                <Input @input="updateValue($event.target.value, 'category')" v-if="selectedTabId === 2" name="category" title="Category" type="text" v-model="category" />
                <div class="flex gap-4">
                    <button @click="$emit('on-cancel')" class="bg-gray-500 text-white px-4 py-2 rounded-lg">Cancel</button>
                    <button @click="onSubmit" class="bg-blue-500 text-white px-4 py-2 rounded-lg">Edit</button>
                </div>
            </div>
        </Dialog>
    </div>
</template>