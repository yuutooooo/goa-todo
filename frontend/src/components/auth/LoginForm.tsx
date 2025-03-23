import React, { useState, useEffect } from "react";
import { useNavigate, useLocation } from "react-router-dom";
import {
  Button,
  TextField,
  Container,
  Typography,
  Box,
  Alert,
} from "@mui/material";
import { useAuth } from "../../contexts/AuthContext";

const LoginForm: React.FC = () => {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [localError, setLocalError] = useState<string | null>(null);
  const { login, error: authError, user } = useAuth();
  const navigate = useNavigate();
  const location = useLocation();

  // ユーザーがログインしている場合、リダイレクト
  useEffect(() => {
    if (user) {
      console.log("LoginForm: User already logged in, redirecting to /todos");
      navigate("/todos", { replace: true });
    }
  }, [user, navigate]);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setLocalError(null);

    // 入力バリデーション
    if (!email) {
      setLocalError("メールアドレスを入力してください");
      return;
    }
    if (!password) {
      setLocalError("パスワードを入力してください");
      return;
    }

    console.log("LoginForm: Attempting login...");
    try {
      await login({ email, password });
      console.log("LoginForm: Login successful, redirecting to /todos");

      // ログイン後のリダイレクト先を決定
      // stateから遷移元URLを取得するか、デフォルトで/todosに遷移
      const fromPath = (location.state as any)?.from?.pathname || "/todos";
      console.log(`LoginForm: Redirecting to ${fromPath}`);
      navigate(fromPath, { replace: true });
    } catch (error) {
      console.error("LoginForm: Login failed", error);
      // エラーはAuthContextで処理済み
    }
  };

  return (
    <Container maxWidth="xs">
      <Box
        sx={{
          marginTop: 8,
          display: "flex",
          flexDirection: "column",
          alignItems: "center",
        }}
      >
        <Typography component="h1" variant="h5">
          ログイン
        </Typography>
        {(localError || authError) && (
          <Alert severity="error" sx={{ width: "100%", mt: 2 }}>
            {localError || authError}
          </Alert>
        )}
        <Box component="form" onSubmit={handleSubmit} sx={{ mt: 1 }}>
          <TextField
            margin="normal"
            required
            fullWidth
            id="email"
            label="メールアドレス"
            name="email"
            autoComplete="email"
            autoFocus
            value={email}
            onChange={(e) => setEmail(e.target.value)}
          />
          <TextField
            margin="normal"
            required
            fullWidth
            name="password"
            label="パスワード"
            type="password"
            id="password"
            autoComplete="current-password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
          />
          <Button
            type="submit"
            fullWidth
            variant="contained"
            sx={{ mt: 3, mb: 2 }}
          >
            ログイン
          </Button>
          <Button
            fullWidth
            variant="text"
            onClick={() => navigate("/register")}
          >
            アカウント登録はこちら
          </Button>
        </Box>
      </Box>
    </Container>
  );
};

export default LoginForm;
