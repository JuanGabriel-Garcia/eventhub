import type { 
  CreateUserRequest, 
  CreateUserResponse, 
  CreateEventRequest, 
  CreateEventResponse,
  EventWithAttendeesResponse,
  LoginRequest,
  LoginResponse
} from '@/types/api';

const API_BASE_URL = '/api';

// Verificar se o backend está rodando em outra porta
// const API_BASE_URL = 'http://localhost:3001'; // Se estiver em outra porta

class ApiService {
  private async request<T>(
    endpoint: string, 
    options: RequestInit = {}
  ): Promise<T> {
    const url = `${API_BASE_URL}${endpoint}`;
    
    const config: RequestInit = {
      headers: {
        'Content-Type': 'application/json',
        ...options.headers,
      },
      ...options,
    };

    // Adicionar token de autenticação se existir
    const token = localStorage.getItem('authToken');
    console.log('Token from localStorage:', token ? `${token.substring(0, 20)}...` : 'null');
    
    if (token) {
      config.headers = {
        ...config.headers,
        'Authorization': `Bearer ${token}`,
      };
    } else {
      console.warn('No auth token found in localStorage');
    }

    try {
      console.log('Making API request to:', url);
      console.log('Request config:', config);
      
      const response = await fetch(url, config);
      
      if (!response.ok) {
        let errorData = {};
        const contentType = response.headers.get('content-type');
        
        // Só tentar fazer parse do JSON se houver content-type application/json
        if (contentType && contentType.includes('application/json')) {
          try {
            errorData = await response.json();
          } catch (e) {
            console.warn('Failed to parse error response as JSON');
          }
        }
        
        console.error('API Error Response:', errorData);
        const errorMessage = (errorData as any)?.message || `HTTP error! status: ${response.status} - ${response.statusText}`;
        throw new Error(errorMessage);
      }
      
      // Verificar se a resposta tem conteúdo antes de fazer parse
      const contentLength = response.headers.get('content-length');
      if (contentLength === '0' || response.status === 204) {
        return {} as T; // Retornar objeto vazio para respostas sem conteúdo
      }
      
      return await response.json();
    } catch (error) {
      console.error('API request failed:', error);
      console.error('Error type:', typeof error);
      console.error('Error message:', error instanceof Error ? error.message : 'Unknown error');
      
      // Verificar se é um erro de conexão
      if (error instanceof TypeError && error.message.includes('fetch')) {
        throw new Error('Não foi possível conectar ao servidor. Possível problema de CORS ou servidor indisponível.');
      }
      
      throw error;
    }
  }

  private async requestWithoutAuth<T>(
    endpoint: string, 
    options: RequestInit = {}
  ): Promise<T> {
    const url = `${API_BASE_URL}${endpoint}`;
    
    const config: RequestInit = {
      headers: {
        'Content-Type': 'application/json',
        ...options.headers,
      },
      ...options,
    };

    // NÃO adicionar token de autenticação

    try {
      console.log('Making API request (without auth) to:', url);
      console.log('Request config:', config);
      
      const response = await fetch(url, config);
      
      if (!response.ok) {
        let errorData = {};
        const contentType = response.headers.get('content-type');
        
        // Só tentar fazer parse do JSON se houver content-type application/json
        if (contentType && contentType.includes('application/json')) {
          try {
            errorData = await response.json();
          } catch (e) {
            console.warn('Failed to parse error response as JSON');
          }
        }
        
        console.error('API Error Response:', errorData);
        const errorMessage = (errorData as any)?.message || `HTTP error! status: ${response.status} - ${response.statusText}`;
        throw new Error(errorMessage);
      }
      
      // Verificar se a resposta tem conteúdo antes de fazer parse
      const contentLength = response.headers.get('content-length');
      if (contentLength === '0' || response.status === 204) {
        return {} as T; // Retornar objeto vazio para respostas sem conteúdo
      }
      
      return await response.json();
    } catch (error) {
      console.error('API request failed:', error);
      console.error('Error type:', typeof error);
      console.error('Error message:', error instanceof Error ? error.message : 'Unknown error');
      
      // Verificar se é um erro de conexão
      if (error instanceof TypeError && error.message.includes('fetch')) {
        throw new Error('Não foi possível conectar ao servidor. Possível problema de CORS ou servidor indisponível.');
      }
      
      throw error;
    }
  }

  // Autenticação
  async login(data: LoginRequest): Promise<LoginResponse> {
    return this.request<LoginResponse>('/auth/login', {
      method: 'POST',
      body: JSON.stringify(data),
    });
  }

  // Usuários
  async createUser(data: CreateUserRequest): Promise<CreateUserResponse> {
    // Criar usuário sem token de autenticação
    return this.requestWithoutAuth<CreateUserResponse>('/users/', {
      method: 'POST',
      body: JSON.stringify(data),
    });
  }

  async getCurrentUser(): Promise<CreateUserResponse> {
    return this.request<CreateUserResponse>('/users/me');
  }

  // Eventos
  async createEvent(data: CreateEventRequest): Promise<CreateEventResponse> {
    return this.request<CreateEventResponse>('/events/', {
      method: 'POST',
      body: JSON.stringify(data),
    });
  }

  async getEvents(): Promise<CreateEventResponse[]> {
    return this.request<CreateEventResponse[]>('/events/');
  }

  async getEventsByOrganizer(): Promise<CreateEventResponse[]> {
    return this.request<CreateEventResponse[]>('/events/organizer');
  }

  async getEventsByUser(): Promise<CreateEventResponse[]> {
    return this.request<CreateEventResponse[]>('/events/registered');
  }

  async getEventById(id: string): Promise<EventWithAttendeesResponse> {
    return this.request<EventWithAttendeesResponse>(`/events/${id}`);
  }

  async getUserById(id: string): Promise<CreateUserResponse> {
    return this.request<CreateUserResponse>(`/users/${id}`);
  }

  async registerToEvent(eventId: string) {
    return this.request(`/events/${eventId}/register`, {
      method: 'POST',
    });
  }

  async cancelRegistration(eventId: string) {
    return this.request(`/events/${eventId}/register`, {
      method: 'DELETE',
    });
  }

  // Função para testar conectividade
  async testConnection(): Promise<boolean> {
    try {
      // Usar um endpoint público que existe no backend
      const response = await fetch(`${API_BASE_URL}/events/`, {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
        },
      });
      // Se retornar 200 ou 401 (token not provided), significa que o servidor está respondendo
      return response.status === 200 || response.status === 401;
    } catch (error) {
      console.error('Connection test failed:', error);
      return false;
    }
  }
}

export const apiService = new ApiService();
