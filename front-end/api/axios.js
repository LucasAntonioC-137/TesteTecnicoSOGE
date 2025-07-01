import axios from 'axios';

const api = axios.create({
  baseURL: 'http://localhost:5000', // Nome do serviço no docker + porta interna
});

export default api;

