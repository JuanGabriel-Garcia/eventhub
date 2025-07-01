"use client";

import { useState, useEffect } from "react";
import { Button } from "@/components/ui/button";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";
import { Plus } from "lucide-react";
import { Link, useNavigate } from "react-router-dom";
import { apiService } from "@/services/api";
import type { CreateEventResponse } from "@/types/api";
import { EventCard } from "@/components/EventCard";
import { LogoutModal } from "@/components/LogoutModal";

export default function DashboardPage() {
  const [user, setUser] = useState({ name: "", email: "", type: "" });
  const [myEvents, setMyEvents] = useState<CreateEventResponse[]>([]);
  const [myRegistrations, setMyRegistrations] = useState<CreateEventResponse[]>([]);
  const navigate = useNavigate();

  useEffect(() => {
    checkLoginAndLoadData();

    // Listener para mudanças no localStorage
    const handleStorageChange = () => {
      checkLoginAndLoadData();
    };

    // Listener para quando a janela ganha foco
    const handleFocus = () => {
      checkLoginAndLoadData();
    };

    window.addEventListener('storage', handleStorageChange);
    window.addEventListener('focus', handleFocus);

    return () => {
      window.removeEventListener('storage', handleStorageChange);
      window.removeEventListener('focus', handleFocus);
    };
  }, [navigate]);

  const checkLoginAndLoadData = () => {
    const isLoggedIn = localStorage.getItem("isLoggedIn");
    if (!isLoggedIn) {
      navigate("/login");
      return;
    }

    loadUserData();
  };

  const loadUserData = () => {
    // Limpar estados anteriores
    setMyEvents([]);
    setMyRegistrations([]);

    const userName = localStorage.getItem("userName") || "";
    const userEmail = localStorage.getItem("userEmail") || "";
    let userType = localStorage.getItem("userType") || "";

    if (!userType) {
      userType = "organizer";
      localStorage.setItem("userType", "organizer");
    }

    setUser({ name: userName, email: userEmail, type: userType });

    console.log("User data:", { name: userName, email: userEmail, type: userType });
    console.log("Current userId from localStorage:", localStorage.getItem("userId"));

    // Buscar eventos do usuário através da API
    if (userType === "organizer") {
      console.log("User is organizer, loading events...");
      loadMyEvents();
    } else {
      console.log("User is not organizer, skipping event loading");
    }

    // Carregar inscrições do usuário
    loadMyRegistrations();
  };

  const loadMyEvents = async () => {
    try {
      console.log("Loading events for organizer...");
      console.log("Auth token exists:", !!localStorage.getItem("authToken"));
      console.log("User ID:", localStorage.getItem("userId"));
      
      const events = await apiService.getEventsByOrganizer();
      console.log("Events loaded from API:", events);
      setMyEvents(events);
    } catch (error) {
      console.error("Error loading events:", error);
      
      // Se o erro for 401, limpar dados e redirecionar para login
      if (error instanceof Error && error.message.includes("401")) {
        console.log("Token inválido, redirecionando para login...");
        localStorage.clear();
        navigate("/login");
        return;
      }
      
      setMyEvents([]);
    }
  };

  const loadMyRegistrations = async () => {
    try {
      console.log("Loading user registrations...");
      console.log("Auth token exists:", !!localStorage.getItem("authToken"));
      console.log("User ID:", localStorage.getItem("userId"));
      
      const events = await apiService.getEventsByUser();
      console.log("Registrations loaded from API:", events);
      setMyRegistrations(events);
    } catch (error) {
      console.error("Error loading registrations:", error);
      
      // Se o erro for 401, limpar dados e redirecionar para login
      if (error instanceof Error && error.message.includes("401")) {
        console.log("Token inválido, redirecionando para login...");
        localStorage.clear();
        navigate("/login");
        return;
      }
      
      setMyRegistrations([]);
    }
  };

  return (
    <div className="min-h-screen bg-gray-50">
      {/* Header */}
      <header className="bg-white shadow-sm border-b">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="flex justify-between items-center h-16">
            <Link to="/" className="text-2xl font-bold text-gray-900">
              EventHub
            </Link>
            <nav className="flex items-center space-x-4">
              <Link to="/">
                <Button variant="ghost">Explorar Eventos</Button>
              </Link>
              <Link to="/create-event">
                <Button variant="ghost">
                  <Plus className="w-4 h-4 mr-2" />
                  Criar Evento
                </Button>
              </Link>
              <LogoutModal
                onLogout={() => {
                  localStorage.clear();
                  navigate("/");
                }}
              >
                <Button variant="outline">
                  Sair
                </Button>
              </LogoutModal>
            </nav>
          </div>
        </div>
      </header>

      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
        <div className="mb-8">
          <h1 className="text-3xl font-bold text-gray-900 mb-2">
            Olá, {user.name}!
          </h1>
          <p className="text-gray-600">
            Gerencie seus eventos e inscrições em um só lugar.
          </p>
        </div>

        <Tabs defaultValue="registrations" className="space-y-6">
          <TabsList className="grid w-full grid-cols-2">
            <TabsTrigger value="registrations">Minhas Inscrições</TabsTrigger>
            <TabsTrigger value="events">Meus Eventos</TabsTrigger>
          </TabsList>

          <TabsContent value="registrations" className="space-y-6">
            <div className="flex justify-between items-center">
              <h2 className="text-2xl font-semibold">Eventos Inscritos</h2>
            </div>

            <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
              {myRegistrations.map((registration) => (
                <EventCard key={registration.id} event={registration} showRegistrationBadge={true} />
              ))}
            </div>

            {myRegistrations.length === 0 && (
              <div className="text-center py-12">
                <p className="text-gray-500 text-lg mb-4">
                  Você ainda não se inscreveu em nenhum evento.
                </p>
                <Link to="/">
                  <Button>Explorar Eventos</Button>
                </Link>
              </div>
            )}
          </TabsContent>

          <TabsContent value="events" className="space-y-6">
            <div className="flex justify-between items-center">
              <h2 className="text-2xl font-semibold">
                {user.type === "organizer" ? "Eventos Organizados" : "Eventos"}
              </h2>
              <div className="flex gap-2">
                {user.type === "organizer" && (
                  <Link to="/create-event">
                    <Button>
                      <Plus className="w-4 h-4 mr-2" />
                      Novo Evento
                    </Button>
                  </Link>
                )}
              </div>
            </div>

            <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
              {myEvents.map((event) => (
                <EventCard key={event.id} event={event} />
              ))}
            </div>

            {myEvents.length === 0 && (
              <div className="text-center py-12">
                <p className="text-gray-500 text-lg mb-4">
                  {user.type === "organizer"
                    ? "Você ainda não criou nenhum evento."
                    : "Não há eventos disponíveis."}
                </p>
                {user.type === "organizer" ? (
                  <Link to="/create-event">
                    <Button>
                      <Plus className="w-4 h-4 mr-2" />
                      Criar Primeiro Evento
                    </Button>
                  </Link>
                ) : (
                  <Link to="/">
                    <Button>Explorar Eventos</Button>
                  </Link>
                )}
              </div>
            )}
          </TabsContent>
        </Tabs>
      </div>
    </div>
  );
}
