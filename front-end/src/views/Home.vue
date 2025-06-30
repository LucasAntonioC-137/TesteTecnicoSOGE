<template>
  <div class="container">
    <h1 class="title">Fomulário de sugestão</h1>

    <form @submit.prevent="enviarSugestao" class="formulario">
      <input v-model="nome" type="text" placeholder="Nome do colaborador" />
      <input v-model="setor" type="text" placeholder="Setor" />
      <textarea v-model="descricao" placeholder="Descrição da sugestão"></textarea>
      <button type="submit">Enviar Sugestão</button>
    </form>

    <!-- Botão Gerenciar sugestões -->
    <button class="manage-button" @click="gerenciarSugestoes">
      Gerenciar sugestões
      <svg xmlns="http://www.w3.org/2000/svg" class="icon-person" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2" width="20" height="20" >
        <path stroke-linecap="round" stroke-linejoin="round" d="M5.121 17.804A7 7 0 0112 15a7 7 0 016.879 2.804M15 11a3 3 0 11-6 0 3 3 0 016 0z" />
      </svg>
    </button>
  </div>
</template>

<script>
import api from '../../api/axios';

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
    async enviarSugestao() {
      try {
        await api.post('/register', {
            collaborator_name: this.nome,
            sector: this.setor,
            description: this.descricao
        });

        alert('Sugestão enviada com sucesso!');
        this.nome = '';
        this.setor = '';
        this.descricao = '';
      } catch (error) {
        console.error('Erro ao enviar sugestão:', error);
        alert('Erro ao enviar sugestão. Verifique os dados ou tente novamente.');
      }
    },
    gerenciarSugestoes() {
      this.$router.push('/suggestions');
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
  border: 2px solid #00008B;
  border-radius: 8px;
  outline: none;
}

textarea {
  resize: vertical;
  min-height: 100px;
}

button {
  background-color: #00008B;
  color: white;
  padding: 12px;
  border: none;
  border-radius: 8px;
  font-weight: bold;
  cursor: pointer;
  transition: 0.3s;
}

button:hover {
  background-color: blue;
}

button:active {
  background-color: #000080; /* cor enquanto estiver clicado */
}

/* Estilo específico para o botão Gerenciar sugestões */
.manage-button {
  margin-top: 30px;
  display: flex;
  align-items: center;
  gap: 8px;
  background-color: #008000; /* verde */
  padding: 12px 20px;
  border-radius: 8px;
  font-weight: bold;
  font-size: 16px;
}

.manage-button:hover {
  background-color: #006400; /* verde escuro */
}

.manage-button:active {
  background-color: #004d00;
}

.icon-person {
  stroke: white;
}

</style>
