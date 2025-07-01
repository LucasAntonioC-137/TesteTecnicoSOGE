<script setup>
defineProps({
  msg: {
    type: String,
    required: true,
  },
})
</script>

<template>
  <div class="welcome">
    <h1 class="blue">{{ msg }}</h1>
    <h3>
      Você está no sistema interno de sugestão de melhorias da Fibralink
    </h3>
  </div>
</template>

<style scoped>

</style>
