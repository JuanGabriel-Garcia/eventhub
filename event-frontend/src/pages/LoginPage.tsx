"use client";

import type React from "react";

import { useState } from "react";
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
import { Link, useNavigate } from "react-router-dom";
import { Eye, EyeOff } from "lucide-react";
import { apiService } from "@/services/api";
import type { LoginRequest } from "@/types/api";
import { validateLoginForm } from "@/utils/validation";

export default function LoginPage() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState("");
  const [isLoading, setIsLoading] = useState(false);
  const [showPassword, setShowPassword] = useState(false);
  const navigate = useNavigate();

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setIsLoading(true);
    setError("");

    if (!email || !password) {
      setError("Por favor, preencha todos os campos");
      setIsLoading(false);
      return;
    }

    // Usar validação centralizada
    const validationError = validateLoginForm(email, password);
    if (validationError) {
      setError(validationError);
      setIsLoading(false);
      return;
    }

    // Teste de conectividade e login
    try {
      const loginData: LoginRequest = { email, password };
      
      const response = await apiService.login(loginData);
      
      // O backend agora retorna { token: "..." }
      if (!response || !response.token) {
        throw new Error('Falha na autenticação');
      }
      
      // Verificar se o token não está vazio
      if (response.token.trim() === '') {
        throw new Error('Token inválido recebido');
      }
      
      // Salvar o token primeiro
      localStorage.setItem("authToken", response.token);
      
      // Tentar buscar dados completos do usuário
      try {
        const userData = await apiService.getCurrentUser();
        localStorage.setItem("userId", userData.id);
        localStorage.setItem("userName", userData.name);
        localStorage.setItem("userEmail", userData.email);
        localStorage.setItem("userType", userData.userType);
      } catch (userError) {
        // Fallback para dados temporários
        localStorage.setItem("userEmail", email);
        localStorage.setItem("userName", email.split("@")[0]);
        localStorage.setItem("userType", "participant");
        localStorage.setItem("userId", "temp-id");
      }
      
      localStorage.setItem("isLoggedIn", "true");
      navigate("/dashboard");
    } catch (error) {
      // Limpar qualquer dado de autenticação em caso de erro
      apiService.clearAuthData();
      
      // Melhor tratamento de erros
      let errorMessage = "Erro ao fazer login. Tente novamente.";
      
      if (error instanceof Error) {
        if (error.message.includes('fetch')) {
          errorMessage = "Não foi possível conectar ao servidor.";
        } else if (error.message.includes('401') || error.message.includes('Unauthorized')) {
          errorMessage = "Email ou senha incorretos.";
        } else if (error.message.includes('400') || error.message.includes('Bad Request')) {
          errorMessage = "Dados de login inválidos.";
        } else if (error.message.includes('500')) {
          errorMessage = "Erro interno do servidor. Tente novamente.";
        }
      }
      
      setError(errorMessage);
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <div className="min-h-screen bg-gray-50 flex items-center justify-center py-12 px-4 sm:px-6 lg:px-8">
      <Card className="w-full max-w-md">
        <CardHeader className="text-center">
          <CardTitle className="text-2xl font-bold">Entrar</CardTitle>
          <CardDescription>
            Acesse sua conta para gerenciar seus eventos
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
              <Label htmlFor="email">E-mail</Label>
              <Input
                id="email"
                type="email"
                placeholder="seu@email.com"
                value={email}
                onChange={(e) => setEmail(e.target.value)}
                required
              />
            </div>
            <div className="space-y-2">
              <Label htmlFor="password">Senha</Label>
              <div className="relative">
                <Input
                  id="password"
                  type={showPassword ? "text" : "password"}
                  placeholder="Sua senha"
                  value={password}
                  onChange={(e) => setPassword(e.target.value)}
                  required
                />
                <button
                  type="button"
                  className="absolute inset-y-0 right-0 pr-3 flex items-center"
                  onClick={() => setShowPassword(!showPassword)}
                >
                  {showPassword ? (
                    <EyeOff className="h-4 w-4 text-gray-400 hover:text-gray-600" />
                  ) : (
                    <Eye className="h-4 w-4 text-gray-400 hover:text-gray-600" />
                  )}
                </button>
              </div>
            </div>
          </CardContent>
          <CardFooter className="flex flex-col space-y-4 pt-6">
            <Button type="submit" className="w-full" disabled={isLoading}>
              {isLoading ? "Entrando..." : "Entrar"}
            </Button>
            <p className="text-sm text-center text-gray-600">
              Não tem uma conta?{" "}
              <Link to="/register" className="text-blue-600 hover:underline">
                Cadastre-se
              </Link>
            </p>
          </CardFooter>
        </form>
      </Card>
    </div>
  );
}
