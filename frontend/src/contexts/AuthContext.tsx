import React, { createContext, useState, useContext, useEffect, useCallback } from 'react';
import { User, LoginPayload, RegisterPayload } from '../types';
import { userApi } from '../api/client';

interface AuthContextType {
  user: User | null;
  loading: boolean;
  error: string | null;
  login: (data: LoginPayload) => Promise<void>;
  register: (data: RegisterPayload) => Promise<void>;
  logout: () => void;
}

const AuthContext = createContext<AuthContextType | undefined>(undefined);

export const useAuth = () => {
  const context = useContext(AuthContext);
  if (context === undefined) {
    throw new Error('useAuth must be used within an AuthProvider');
  }
  return context;
};

export const AuthProvider: React.FC<{ children: React.ReactNode }> = ({ children }) => {
  const [user, setUser] = useState<User | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    // ローカルストレージからユーザー情報を取得
    const loadUser = () => {
      console.log('AuthContext: Loading user from localStorage');
      try {
        const storedUser = localStorage.getItem('user');
        if (storedUser) {
          const parsedUser = JSON.parse(storedUser);
          console.log('AuthContext: Found user in localStorage', parsedUser);
          setUser(parsedUser);
        } else {
          console.log('AuthContext: No user found in localStorage');
        }
      } catch (err) {
        console.error('AuthContext: Error loading user from localStorage:', err);
      } finally {
        setLoading(false);
      }
    };

    loadUser();
  }, []);

  const setUserState = useCallback((newUser: User | null) => {
    console.log('AuthContext: Setting user state:', newUser);
    setUser(newUser);
  }, []);

  const login = async (data: LoginPayload) => {
    console.log('AuthContext: Login attempt with:', data.email);
    try {
      setError(null);
      setLoading(true);
      
      console.log('AuthContext: Sending login request to API');
      const response = await userApi.login(data);
      console.log('AuthContext: Login response received:', response);
      
      // APIレスポンスからユーザー情報を取得
      // OpenAPI定義に基づいて、responseデータは直接UserResultの形式になっています
      const userData = response.data;
      console.log('AuthContext: Login successful, user data:', userData);
      
      if (!userData || !userData.id) {
        throw new Error('ユーザー情報が含まれていません');
      }
      
      // ローカルストレージに保存
      // ※トークンは扱わない
      localStorage.setItem('user', JSON.stringify(userData));
      console.log('AuthContext: User data saved to localStorage');
      
      // 状態を更新
      setUserState(userData);
      console.log('AuthContext: User state updated, current user:', userData);
    } catch (err: any) {
      console.error('AuthContext: Login error:', err);
      const errorMessage = err.response?.data?.message || 'ログインに失敗しました';
      setError(errorMessage);
      throw new Error(errorMessage);
    } finally {
      setLoading(false);
    }
  };

  const register = async (data: RegisterPayload) => {
    console.log('AuthContext: Register attempt with:', data.email);
    try {
      setError(null);
      setLoading(true);
      
      console.log('AuthContext: Sending register request to API');
      const response = await userApi.register(data);
      console.log('AuthContext: Register response received:', response);
      
      // APIレスポンスからユーザー情報を取得
      // OpenAPI定義に基づいて、responseデータは直接UserResultの形式になっています
      const userData = response.data;
      console.log('AuthContext: Registration successful, user data:', userData);
      
      if (!userData || !userData.id) {
        throw new Error('ユーザー情報が含まれていません');
      }
      
      // ローカルストレージに保存
      // ※トークンは扱わない
      localStorage.setItem('user', JSON.stringify(userData));
      console.log('AuthContext: User data saved to localStorage');
      
      // 状態を更新
      setUserState(userData);
      console.log('AuthContext: User state updated, current user:', userData);
    } catch (err: any) {
      console.error('AuthContext: Registration error:', err);
      const errorMessage = err.response?.data?.message || 'ユーザー登録に失敗しました';
      setError(errorMessage);
      throw new Error(errorMessage);
    } finally {
      setLoading(false);
    }
  };

  const logout = () => {
    console.log('AuthContext: Logging out user');
    localStorage.removeItem('token'); // 互換性のために残す
    localStorage.removeItem('user');
    setUserState(null);
    console.log('AuthContext: User logged out, localStorage cleared');
  };

  return (
    <AuthContext.Provider value={{ user, loading, error, login, register, logout }}>
      {children}
    </AuthContext.Provider>
  );
};

export default AuthContext; 