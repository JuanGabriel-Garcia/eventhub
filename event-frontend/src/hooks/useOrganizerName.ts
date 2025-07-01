import { useState, useEffect } from 'react';
import { apiService } from '@/services/api';

interface OrganizerCache {
  [key: string]: string;
}

// Cache para evitar múltiplas requisições para o mesmo organizador
const organizerCache: OrganizerCache = {};

export const useOrganizerName = (organizerId: string) => {
  const [organizerName, setOrganizerName] = useState<string>('Carregando...');
  
  useEffect(() => {
    if (!organizerId) {
      setOrganizerName('Organizador não informado');
      return;
    }

    // Verificar se já temos o nome no cache
    if (organizerCache[organizerId]) {
      setOrganizerName(organizerCache[organizerId]);
      return;
    }

    const fetchOrganizerName = async () => {
      try {
        const user = await apiService.getUserById(organizerId);
        const name = user.name || 'Nome não informado';
        organizerCache[organizerId] = name; // Adicionar ao cache
        setOrganizerName(name);
      } catch (error) {
        console.error('Error fetching organizer name:', error);
        setOrganizerName('Organizador não encontrado');
      }
    };

    fetchOrganizerName();
  }, [organizerId]);

  return organizerName;
};
