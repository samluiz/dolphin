<template>
    <button @click="$emit('click')" type="button" class="duration-200 p-2 hover:bg-gray-500 [clip-path:circle(50%)]">
        <slot></slot>
    </button>
</template>