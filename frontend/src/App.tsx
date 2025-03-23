import React from 'react';
import { BrowserRouter as Router, Route, Routes, Navigate, useLocation } from 'react-router-dom';
import { Container, CssBaseline, Box, CircularProgress } from '@mui/material';
import { AuthProvider, useAuth } from './contexts/AuthContext';

// コンポーネントのインポート
import LoginForm from './components/auth/LoginForm';
import RegisterForm from './components/auth/RegisterForm';
import TodoList from './components/todos/TodoList';
import TodoForm from './components/todos/TodoForm';
import MemoList from './components/memos/MemoList';
import MemoForm from './components/memos/MemoForm';
import Header from './components/layout/Header';

// PrivateRouteコンポーネント
const PrivateRoute: React.FC<{ children: React.ReactNode }> = ({ children }) => {
  const { user, loading } = useAuth();
  const location = useLocation();

  if (loading) {
    return (
      <Box sx={{ display: 'flex', justifyContent: 'center', alignItems: 'center', height: '100vh' }}>
        <CircularProgress />
      </Box>
    );
  }

  if (!user) {
    console.log('PrivateRoute: No user found, redirecting to login', location.pathname);
    return <Navigate to="/login" state={{ from: location }} replace />;
  }

  console.log('PrivateRoute: User authenticated, rendering protected content');
  return <>{children}</>;
};

// ログイン済みの場合にリダイレクトするルート
const PublicRoute: React.FC<{ children: React.ReactNode, redirectTo: string }> = ({ 
  children, 
  redirectTo 
}) => {
  const { user, loading } = useAuth();
  
  if (loading) {
    return (
      <Box sx={{ display: 'flex', justifyContent: 'center', alignItems: 'center', height: '100vh' }}>
        <CircularProgress />
      </Box>
    );
  }
  
  if (user) {
    console.log('PublicRoute: User already logged in, redirecting to', redirectTo);
    return <Navigate to={redirectTo} replace />;
  }
  
  return <>{children}</>;
};

const App: React.FC = () => {
  return (
    <AuthProvider>
      <Router>
        <CssBaseline />
        <Box sx={{ display: 'flex', flexDirection: 'column', minHeight: '100vh' }}>
          <Header />
          <Container component="main" sx={{ flexGrow: 1, py: 3 }}>
            <Routes>
              {/* 認証ルート */}
              <Route path="/login" element={
                <PublicRoute redirectTo="/todos">
                  <LoginForm />
                </PublicRoute>
              } />
              <Route path="/register" element={
                <PublicRoute redirectTo="/todos">
                  <RegisterForm />
                </PublicRoute>
              } />
              
              {/* 保護されたルート */}
              <Route path="/todos" element={<PrivateRoute><TodoList /></PrivateRoute>} />
              <Route path="/todos/new" element={<PrivateRoute><TodoForm /></PrivateRoute>} />
              <Route path="/todos/:id/edit" element={<PrivateRoute><TodoForm /></PrivateRoute>} />
              <Route path="/todos/:todoId/memos" element={<PrivateRoute><MemoList /></PrivateRoute>} />
              <Route path="/todos/:todoId/memos/new" element={<PrivateRoute><MemoForm /></PrivateRoute>} />
              <Route path="/todos/:todoId/memos/:memoId/edit" element={<PrivateRoute><MemoForm /></PrivateRoute>} />
              
              {/* リダイレクト */}
              <Route path="/" element={<Navigate to="/todos" replace />} />
              <Route path="*" element={<Navigate to="/todos" replace />} />
            </Routes>
          </Container>
        </Box>
      </Router>
    </AuthProvider>
  );
};

export default App;
