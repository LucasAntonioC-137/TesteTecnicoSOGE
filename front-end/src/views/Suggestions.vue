<template>
  <div class="container">
    <h1 class="title">Gerenciamento de sugestões</h1>
    
    <!-- Botões de filtro -->
    <div class="filter-options">
      <div class="filter-bar">
        <div class="dropdown">
          <button 
            class="filtro-button"
            :class="{ ativo: showStatusDropdown }"
            @click="toggleStatusDropdown"
          >
            Filtrar por status
          </button>
          <div v-if="showStatusDropdown" class="dropdown-content">
            <div class="dropdown-item" @click="filtrarPorStatus(null)">Todos</div>
            <div class="dropdown-item" @click="filtrarPorStatus('open')">Aberto</div>
            <div class="dropdown-item" @click="filtrarPorStatus('under review')">Em análise</div>
            <div class="dropdown-item" @click="filtrarPorStatus('implemented')">Implementado</div>
          </div>
        </div>

        <button 
          class="filtro-button"
          :class="{ ativo: showSetorSearch }"
          @click="toggleSetorSearch"
        >
          Filtrar por setor
        </button>
      </div>

      <div v-if="showSetorSearch" class="setor-search">
        <input type="text" v-model="setorBusca" placeholder="Digite o setor..." />
        <button @click="buscarPorSetor" class="search-button">
          <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <circle cx="11" cy="11" r="8"></circle>
            <line x1="21" y1="21" x2="16.65" y2="16.65"></line>
          </svg>
        </button>
      </div>
    </div>

    <!-- Caixa de listagem de sugestôes -->
    <div class="lista-sugestoes">
      <div 
        class="sugestao" 
        v-for="(s, index) in sugestoes" 
        :key="s.id_suggestion"
        @click="toggleDescricao(index)"
      >
        <div class="linha">
          <span>{{ s.collaborator_name }}</span>
          <span>{{ s.sector }}</span>
          <span 
            class="status"
            @click.stop="s.exibeDescricao && alterarStatus(index)"
          >
            <template v-if="s.exibeDescricao">
              <button class="status-button">
                {{ s.status }}
              </button>
            </template>
            <template v-else>
              {{ s.status }}
            </template>
          </span>
        </div>

        <div v-if="s.exibeDescricao" class="descricao">
          {{ s.description }}
        </div>
      </div>
    </div>

    <!-- Botão Voltar -->
    <button class="back-button" @click="voltarParaFormulario">
      Voltar para o formulário
      <svg xmlns="http://www.w3.org/2000/svg" class="icon-person" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2" width="20" height="20">
        <path stroke-linecap="round" stroke-linejoin="round" d="M15 19l-7-7 7-7" />
      </svg>
    </button>
  </div>
</template>

<script>
import api from '../../api/axios';

export default {
  name: 'Suggestions',
  data() {
    return {
      showStatusDropdown: false,
      showSetorSearch: false,
      setorBusca: '',
      sugestoes: [],
      sugestoesOriginais: [], // Lista fresca do backend
      filtroStatus: null
    }
  },
  async mounted() {
    await this.carregarSugestoes();
  },
  methods: {
    async carregarSugestoes() {
      try {
        const response = await api.get('/suggestions');
        this.sugestoesOriginais = response.data.map(s => ({ ...s, exibeDescricao: false }));
        this.aplicarFiltros();
      } catch (error) {
        console.error('Erro ao carregar sugestões:', error);
      }
    },

    aplicarFiltros() {
      let resultado = this.sugestoesOriginais;

      if (this.filtroStatus) {
        resultado = resultado.filter(s => s.status === this.filtroStatus);
      }

      if (this.setorBusca.trim()) {
        // Busca exata, comparando string sem case sensitive
        resultado = resultado.filter(s => s.sector.toLowerCase() === this.setorBusca.toLowerCase());
      }

      this.sugestoes = resultado;
    },

    async filtrarPorStatus(status) {
      this.filtroStatus = status;
      this.setorBusca = ''; // Limpa busca por setor

      try {
        if (!status) {
          // Quando status for null (ou "Todos"), busca lista agrupada da API
          const response = await api.get('/suggestions/grouped-by-status');

          const agrupadas = response.data;

          this.sugestoesOriginais = agrupadas.map(s => ({ ...s, exibeDescricao: false }));

        } else {
          // Caso contrário, busca a lista completa (não agrupada)
          await this.carregarSugestoes();
        }

        // Aplica os filtros para exibir na tela
        this.aplicarFiltros();
        this.showStatusDropdown = false;

      } catch (error) {
        console.error('Erro ao filtrar por status:', error);
      }
    },

    async buscarPorSetor() {
      this.filtroStatus = null; // Limpa o filtro de status

      try {
        if (!this.setorBusca.trim()) {
          const response = await api.get('/suggestions/grouped-by-sector');

          const agrupado = response.data;

          // Junta todos os arrays de setores em um único array
          const sugestoes = Object.values(agrupado).flat();

          this.sugestoesOriginais = sugestoes.map(s => ({ ...s, exibeDescricao: false }));
          this.sugestoes = [...this.sugestoesOriginais]; // Aplica diretamente, pois é sem filtro de nome

        } else {
          await this.carregarSugestoes();

          const termo = this.setorBusca.trim().toLowerCase();
          this.sugestoes = this.sugestoesOriginais.filter(
            s => s.sector.toLowerCase() === termo
          );
        }

      } catch (error) {
        console.error('Erro ao buscar por setor:', error);
      }
    },

    toggleStatusDropdown() {
      this.showStatusDropdown = !this.showStatusDropdown;
      if (this.showStatusDropdown) {
        this.showSetorSearch = false;
      }
    },

    toggleSetorSearch() {
      this.showSetorSearch = !this.showSetorSearch;
      if (this.showSetorSearch) {
        this.showStatusDropdown = false;
      }
    },

    toggleDescricao(index) {
      this.sugestoes[index].exibeDescricao = !this.sugestoes[index].exibeDescricao;
    },

    voltarParaFormulario() {
      this.$router.push('/');
    },

    async alterarStatus(index) {

      const s = this.sugestoes[index];
      const statusList = ['open', 'under review', 'implemented'];
      const nextStatus = statusList[(statusList.indexOf(s.status) + 1) % statusList.length];

      try {
        await api.put(`/suggestions/${s.id_suggestion}/status`, { status: nextStatus });
        this.sugestoes[index].status = nextStatus;

        // Atualiza também sugestoesOriginais para manter sincronizado
        const idxOriginal = this.sugestoesOriginais.findIndex(sug => sug.id_suggestion === s.id_suggestion);
        if (idxOriginal !== -1) {
          this.sugestoesOriginais[idxOriginal].status = nextStatus;
        }
      } catch (error) {
        console.error('Erro ao alterar status:', error);
      }
    }
  }
}
</script>

<style scoped>
.container {
  background-color: white;
  padding: 40px;
  display: flex;
  flex-direction: column;
  align-items: center;
  height: 100vh;
  box-sizing: border-box;
  .container {
  min-height: 100vh;
  box-shadow: 0 0 12px rgba(0, 0, 0, 0.05); /* novo */
  border-radius: 12px; /* opcional para suavizar cantos */
}

}

.title {
  font-size: 36px;
  font-weight: 700;
  color: red;
  position: relative;
  padding-bottom: 8px;
  margin-bottom: 30px;
  text-align: center;
}

.title::after {
  content: '';
  position: absolute;
  left: 50%;
  transform: translateX(-50%);
  bottom: 0;
  width: 60%;
  height: 3px;
  background-color: #00008B; /* Azul escuro */
  border-radius: 2px;
}


.filter-options {
  display: flex;
  flex-direction: column;
  align-items: center;
  margin-bottom: 30px;
  gap: 10px;
}

.filter-bar {
  display: flex;
  gap: 20px;
  margin-bottom: 20px;
  justify-content: center;
}

.filtro-button {
  background-color: #ffffff;
  color: #000000;
  border: 2px solid #00008B;
  padding: 10px 16px;
  border-radius: 8px;
  font-weight: bold;
  cursor: pointer;
  transition: 0.3s;
  width: 180px;
  text-align: center;
}

.filtro-button:hover,
.filtro-button.ativo {
  background-color: #00008B;
  color: white;
}

.setor-search {
  display: flex;
  align-items: center;
  gap: 8px;
  margin-bottom: center;
}

.setor-search input {
  padding: 10px;
  font-size: 16px;
  border: 2px solid #00008B;
  border-radius: 8px;
  outline: none;
  width: 250px;
}

.search-button {
  background-color: #ffffff;
  border: 2px solid #00008B;
  border-radius: 8px;
  padding: 8px 10px;
  cursor: pointer;
  transition: 0.3s;
  display: flex;
  align-items: center;
  justify-content: center;
}

.search-button svg {
  stroke: #00008B;
}


.search-button:hover {
  background-color: #00008B;
}

.search-button:hover svg {
  stroke: white;
}

.dropdown {
  position: relative;
}

.dropdown-content {
  position: absolute;
  top: 100%;
  left: 0;
  background-color: #ffffff;
  border: 2px solid #00008B;
  border-radius: 8px;
  padding: 8px;
  z-index: 100;
  display: flex;
  flex-direction: column;
  width: 180px;
}

.dropdown-item {
  padding: 8px 12px;
  font-weight: bold;
  color: #000000;
  cursor: pointer;
  border-radius: 6px;
}

.dropdown-item:hover {
  background-color: #00008B;
  color: white;
}

.sugestoes {
  width: 100%;
  max-width: 700px;
  display: flex;
  flex-direction: column;
  gap: 15px;
}

.lista-sugestoes {
  background-color: #f0f0f0;
  padding: 20px;
  border-radius: 10px;
  width: 100%;
  max-width: 900px;
  margin-bottom: 30px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.05);

  flex: 1;
  overflow-y: auto;
  max-height: 100vh;
}

.sugestao {
  background-color: white;
  border-radius: 8px;
  padding: 16px;
  margin-bottom: 5px; /* aumenta o espaço entre as sugestões */
  cursor: pointer;
  transition: background-color 0.2s ease;
  color: black; /* garante texto legível */
}

.sugestao:hover {
  background-color: #f9f9f9;
}

.linha {
  display: flex;
  justify-content: space-between;
  flex-wrap: wrap;
  gap: 10px;
}

.linha span {
  color: black; /* garante que os campos Nome, Setor e Status estejam legíveis */
  font-weight: 500;
}

.descricao {
  margin-top: 8px;
  color: #333;
}

.status {
  color: black; /* cor padrão */
  font-weight: bold;
}

.status.clicavel {
  color: #00008B;
  text-decoration: underline;
  cursor: pointer;
}

.back-button {
  margin-top: 40px;
  display: flex;
  align-items: center;
  gap: 8px;
  background-color: #008000;
  padding: 12px 20px;
  border-radius: 8px;
  font-weight: bold;
  font-size: 16px;
  color: white;
  border: none;
  cursor: pointer;
  transition: 0.3s;
}

.back-button:hover {
  background-color: #006400;
}

.back-button:active {
  background-color: #004d00;
}

.icon-person {
  stroke: white;
}

.status-button {
  background-color: #ffffff;
  border: 2px solid #00008B;
  color: #000000;
  font-weight: bold;
  border-radius: 8px;
  padding: 6px 12px;
  cursor: pointer;
  transition: 0.3s;
}

.status-button:hover {
  background-color: #00008B;
  color: white;
}

</style>
