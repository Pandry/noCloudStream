<template>
    <div class=bg>
        <input
            @input="handleInput"
            :model="modelValue"
            :class="'input-text ' + additionalClasses"
            placeholder=" "
            :type="type"
            />
        <span class="floating-label">{{placeholder}}</span>
    </div>
</template>

<script setup>
    import {ref, reactive} from 'vue'
    //const state = reactive({modelValue: ""})

    const props = defineProps({
        modelValue: {
            type: String,
            default: ''
        },
        additionalClasses: {
            type: String,
            default: ""
        },
        placeholder: {
            type: String,
            default: ""
        },
        type: {
            type: String,
            default: "text"
        },
    })
    const emit = defineEmits(['update:modelValue'])

    //const val = ref(props.modelValue)

    const handleInput = (e) =>emit('update:modelValue', e.target.value)
</script>

<style scoped>
input {
    border-radius: 20px;
    border-color: transparent;
    color: #444;
    font-size: 15px;
    font-weight: bold;
    width: 100%;
}

input:focus {
    outline: none
}

input:focus ~ .floating-label,
input:not(:placeholder-shown):not(:focus) ~ .floating-label {
  top: -4px;
  left: 14px;
  opacity: 1;
  font-size: 14px;
  font-weight: normal;
}

.floating-label {
  position: absolute;
  pointer-events: none;
  left: 20px;
  top: 5px;
  transition: 0.2s ease all;
  color: #333;
  font-size: 18px;
  font-weight: bold;
}

.bg{
    background-color: white;
    border-radius: 25px;
    padding: 6px 4px;
    width:100%;
    margin: 20px;
}
</style>
