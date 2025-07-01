"use client";

import type React from "react";

import { useEffect, useState } from "react";
import { Button } from "@/components/ui/button";
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Alert, AlertDescription } from "@/components/ui/alert";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { Link, useNavigate } from "react-router-dom";
import { apiService } from "@/services/api";
import type { CreateUserRequest } from "@/types/api";

export default function RegisterPage() {
  const [formData, setFormData] = useState({
    name: "",
    email: "",
    password: "",
    confirmPassword: "",
    userType: "" as 'participant' | 'organizer' | "",
  });
  const [error, setError] = useState("");
  const [isLoading, setIsLoading] = useState(false);
  const navigate = useNavigate();

  const handleInputChange = (field: string, value: string) => {
    setFormData((prev) => ({ ...prev, [field]: value }));
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setIsLoading(true);
    setError("");

    // Validações
    if (
      !formData.name ||
      !formData.email ||
      !formData.password ||
      !formData.userType
    ) {
      setError("Por favor, preencha todos os campos");
      setIsLoading(false);
      return;
    }

    if (formData.password !== formData.confirmPassword) {
      setError("As senhas não coincidem");
      setIsLoading(false);
      return;
    }

    if (formData.password.length < 6) {
      setError("A senha deve ter pelo menos 6 caracteres");
      setIsLoading(false);
      return;
    }

    try {
      // Preparar dados para API
      const userData: CreateUserRequest = {
        name: formData.name,
        email: formData.email,
        password: formData.password,
        userType: formData.userType as 'participant' | 'organizer',
      };

      // Criar usuário via API
      await apiService.createUser(userData);
      console.log('User created successfully');
      
      // Após criar o usuário, fazer login automaticamente para obter o token e dados completos
      console.log('Attempting automatic login...');
      const loginResponse = await apiService.login({
        email: formData.email,
        password: formData.password,
      });
      console.log('Login successful, token received:', loginResponse.token ? 'Yes' : 'No');
      
      // Salvar o token primeiro
      localStorage.setItem("authToken", loginResponse.token);
      console.log('Token saved to localStorage');
      
      // Buscar dados completos do usuário
      try {
        console.log('Fetching user data...');
        const userDataResponse = await apiService.getCurrentUser();
        console.log('User data received:', userDataResponse);
        localStorage.setItem("userId", userDataResponse.id);
        localStorage.setItem("userName", userDataResponse.name);
        localStorage.setItem("userEmail", userDataResponse.email);
        localStorage.setItem("userType", userDataResponse.userType);
        
        // Teste adicional: verificar se o token realmente funciona
        console.log('Testing token with events endpoint...');
        try {
          await apiService.getEventsByUser();
          console.log('Token test successful - user can access protected endpoints');
        } catch (testError) {
          console.warn('Token test failed:', testError);
        }
      } catch (userError) {
        console.warn('Não foi possível buscar dados do usuário, usando dados do formulário:', userError);
        
        // Se falhar aqui, pode ser que o token não esteja funcionando
        if (userError instanceof Error && userError.message.includes("401")) {
          throw new Error("Falha na autenticação após criar conta. Tente fazer login manualmente.");
        }
        
        // Fallback para dados do formulário
        localStorage.setItem("userEmail", formData.email);
        localStorage.setItem("userName", formData.name);
        localStorage.setItem("userType", formData.userType);
        localStorage.setItem("userId", "temp-id");
      }
      
      localStorage.setItem("isLoggedIn", "true");
      navigate("/dashboard");
    } catch (error) {
      console.error("Registration error:", error);
      
      // Mensagens de erro mais específicas
      let errorMessage = "Erro ao criar conta. Tente novamente.";
      
      if (error instanceof Error) {
        if (error.message.includes("already exists") || error.message.includes("duplicate")) {
          errorMessage = "Este e-mail já está cadastrado. Use outro e-mail ou faça login.";
        } else if (error.message.includes("invalid email")) {
          errorMessage = "E-mail inválido. Verifique o formato do e-mail.";
        } else if (error.message.includes("password")) {
          errorMessage = "Erro na senha. Verifique se atende aos requisitos.";
        } else {
          errorMessage = error.message;
        }
      }
      
      setError(errorMessage);
    } finally {
      setIsLoading(false);
    }
  };

  // Testar conectividade quando o componente carregar
  useEffect(() => {
    const testConnection = async () => {
      const isConnected = await apiService.testConnection();
      if (!isConnected) {
        setError("Não foi possível conectar ao servidor. Verifique se o backend está rodando.");
      }
    };
    testConnection();
  }, []);

  return (
    <div className="min-h-screen bg-gray-50 flex items-center justify-center py-12 px-4 sm:px-6 lg:px-8">
      <Card className="w-full max-w-md">
        <CardHeader className="text-center">
          <CardTitle className="text-2xl font-bold">Criar Conta</CardTitle>
          <CardDescription>
            Cadastre-se para começar a participar de eventos
          </CardDescription>
        </CardHeader>
        <form onSubmit={handleSubmit}>
          <CardContent className="space-y-4">
            {error && (
              <Alert variant="destructive">
                <AlertDescription>{error}</AlertDescription>
              </Alert>
            )}
            <div className="space-y-2">
              <Label htmlFor="name">Nome completo</Label>
              <Input
                id="name"
                type="text"
                placeholder="Seu nome completo"
                value={formData.name}
                onChange={(e) => handleInputChange("name", e.target.value)}
                required
              />
            </div>
            <div className="space-y-2">
              <Label htmlFor="email">E-mail</Label>
              <Input
                id="email"
                type="email"
                placeholder="seu@email.com"
                value={formData.email}
                onChange={(e) => handleInputChange("email", e.target.value)}
                required
              />
            </div>
            <div className="space-y-2">
              <Label htmlFor="userType">Tipo de usuário</Label>
              <Select
                value={formData.userType}
                onValueChange={(value) => handleInputChange("userType", value)}
              >
                <SelectTrigger>
                  <SelectValue placeholder="Selecione o tipo" />
                </SelectTrigger>
                <SelectContent>
                  <SelectItem value="participant">Participante</SelectItem>
                  <SelectItem value="organizer">Organizador</SelectItem>
                </SelectContent>
              </Select>
            </div>
            <div className="space-y-2">
              <Label htmlFor="password">Senha</Label>
              <Input
                id="password"
                type="password"
                placeholder="Mínimo 6 caracteres"
                value={formData.password}
                onChange={(e) => handleInputChange("password", e.target.value)}
                required
              />
            </div>
            <div className="space-y-2">
              <Label htmlFor="confirmPassword">Confirmar senha</Label>
              <Input
                id="confirmPassword"
                type="password"
                placeholder="Confirme sua senha"
                value={formData.confirmPassword}
                onChange={(e) =>
                  handleInputChange("confirmPassword", e.target.value)
                }
                required
              />
            </div>
          </CardContent>
          <CardFooter className="flex flex-col space-y-4">
            <Button type="submit" className="w-full" disabled={isLoading}>
              {isLoading ? "Criando conta..." : "Criar conta"}
            </Button>
            <p className="text-sm text-center text-gray-600">
              Já tem uma conta?{" "}
              <Link to="/login" className="text-blue-600 hover:underline">
                Faça login
              </Link>
            </p>
          </CardFooter>
        </form>
      </Card>
    </div>
  );
}
