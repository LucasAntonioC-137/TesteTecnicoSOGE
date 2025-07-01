<script setup>
import { RouterLink, RouterView } from 'vue-router'
import Welcome from './components/Welcome.vue'
</script>

<template>
  <header>
    <img alt="logo" class="logo" src="@/assets/logo.png" width="280" height="125" />
    <div class="welcomemsg">
      <Welcome msg="Bem vindo(a)!" />
    </div>
  </header>

  <RouterView />
</template>

<style scoped>
header {
  line-height: 1.5;
  max-height: 200vh;
}

.logo {
  display: block;
  margin: 0 auto 2rem;
}

@media (min-width: 1024px) {
  header {
    display: flex;
    place-items: center;
    padding-right: calc(var(--section-gap) / 2);
  }

  .logo {
    margin: 0 2rem 0 0;
  }
}
</style>
