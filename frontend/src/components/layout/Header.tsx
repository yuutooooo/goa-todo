import React from "react";
import { AppBar, Toolbar, Typography, Button, Box } from "@mui/material";
import { useNavigate } from "react-router-dom";
import { useAuth } from "../../contexts/AuthContext";

const Header: React.FC = () => {
  const { user, logout } = useAuth();
  const navigate = useNavigate();

  const handleLogout = () => {
    logout();
    navigate("/login");
  };

  return (
    <AppBar position="static">
      <Toolbar>
        <Typography
          variant="h6"
          component="div"
          sx={{ flexGrow: 1, cursor: "pointer" }}
          onClick={() => navigate("/")}
        >
          Todo管理アプリ
        </Typography>

        {user ? (
          <Box>
            <Typography variant="body1" component="span" sx={{ mr: 2 }}>
              {user.name}さん
            </Typography>
            <Button color="inherit" onClick={handleLogout}>
              ログアウト
            </Button>
          </Box>
        ) : (
          <Box>
            <Button
              color="inherit"
              onClick={() => navigate("/login")}
              sx={{ mr: 1 }}
            >
              ログイン
            </Button>
            <Button color="inherit" onClick={() => navigate("/register")}>
              登録
            </Button>
          </Box>
        )}
      </Toolbar>
    </AppBar>
  );
};

export default Header;
