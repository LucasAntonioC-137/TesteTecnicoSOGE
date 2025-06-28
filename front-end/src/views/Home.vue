<template>
  <div class="container">
    <h1 class="title">Fibralink Sugestões</h1>

    <form @submit.prevent="enviarSugestao" class="formulario">
      <input v-model="nome" type="text" placeholder="Nome do colaborador" />
      <input v-model="setor" type="text" placeholder="Setor" />
      <textarea v-model="descricao" placeholder="Descrição da sugestão"></textarea>
      <button type="submit">Enviar Sugestão</button>
    </form>
  </div>
</template>

<script>
export default {
  name: 'Home',
  data() {
    return {
      nome: '',
      setor: '',
      descricao: ''
    }
  },
  methods: {
    enviarSugestao() {
      console.log('Sugestão enviada:', {
        nome: this.nome,
        setor: this.setor,
        descricao: this.descricao
      });
      // Aqui você pode integrar com sua API futuramente
    }
  }
}
</script>

<style scoped>
.container {
  background-color: white;
  min-height: 100vh;
  padding: 40px;
  display: flex;
  flex-direction: column;
  align-items: center;
}

.title {
  color: red;
  font-size: 32px;
  margin-bottom: 30px;
}

.formulario {
  width: 100%;
  max-width: 500px;
  display: flex;
  flex-direction: column;
  gap: 15px;
}

input,
textarea {
  padding: 12px;
  font-size: 16px;
  border: 2px solid blue;
  border-radius: 8px;
  outline: none;
}

textarea {
  resize: vertical;
  min-height: 100px;
}

button {
  background-color: blue;
  color: white;
  padding: 12px;
  border: none;
  border-radius: 8px;
  font-weight: bold;
  cursor: pointer;
  transition: 0.3s;
}

button:hover {
  background-color: darkblue;
}
</style>
