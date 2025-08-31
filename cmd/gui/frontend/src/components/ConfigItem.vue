<script setup>
import { ref, computed } from "vue";

const props = defineProps(["label", "modelValue", "help", "show"]);
const emit = defineEmits(["update:modelValue"]);
const model = ref(props.modelValue);

const emitValue = () => emit("update:modelValue", model.value);

const formatedLabel = computed(
  () => props.label?.replace(/([A-Z])/g, " $1").trim() ?? props.label
);
</script>
<template>
  <li class="flex align-items-center py-3 px-2 border-top-1 surface-border">
    <div class="flex flex-column align-items-flexstart flex-wrap w-4">
      <div class="text-700 w-12 md:w-4 font-medium font-bold">
        {{ formatedLabel }}
      </div>
      <div v-if="help" class="text-400 text-sm w-10 mt-2">{{ help }}</div>
    </div>

    <div class="flex text-900 w-8 md:w-8">
      <InputNumber
        v-if="modelValue.constructor.name == 'Number'"
        mode="decimal"
        showButtons
        :min="0"
        :max="99"
        v-model="model"
        @update:modelValue="emitValue"
      />
      <InputSwitch
        v-else-if="modelValue.constructor.name == 'Boolean'"
        v-model="model"
        @update:modelValue="emitValue"
      />
      <Textarea
        v-else-if="modelValue.constructor.name == 'Array'"
        v-model="model"
        @update:modelValue="emitValue"
        rows="5"
        cols="30"
      />
      <InputText
        v-else="modelValue.constructor.name == 'String'"
        v-model="model"
        @update:modelValue="emitValue"
      />
    </div>
  </li>
</template>
<style scoped>
input,
textarea {
  width: 100%;
}
</style>
