import React, { useState, useEffect } from "react";
import { useParams, useNavigate } from "react-router-dom";
import {
  Container,
  Typography,
  List,
  ListItem,
  ListItemText,
  IconButton,
  Button,
  Box,
  Divider,
  CircularProgress,
  Alert,
  Paper,
  Breadcrumbs,
  Link,
} from "@mui/material";
import {
  Delete as DeleteIcon,
  Edit as EditIcon,
  Add as AddIcon,
  ArrowBack as ArrowBackIcon,
} from "@mui/icons-material";
import { memoApi, todoApi } from "../../api/client";
import { Memo, Todo } from "../../types";
import { useAuth } from "../../contexts/AuthContext";

const MemoList: React.FC = () => {
  const { todoId } = useParams<{ todoId: string }>();
  const todoIdNum = parseInt(todoId || "0");
  const { user } = useAuth();
  const userId = user?.id || 0;

  const [memos, setMemos] = useState<Memo[]>([]);
  const [todo, setTodo] = useState<Todo | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const navigate = useNavigate();

  useEffect(() => {
    const fetchData = async () => {
      if (!user) {
        navigate("/login");
        return;
      }

      try {
        setLoading(true);

        // TODOの情報を取得
        const todoResponse = await todoApi.getTodo(todoIdNum);
        setTodo(todoResponse.data);

        // メモリストの取得
        const memosResponse = await memoApi.getMemos(todoIdNum);
        setMemos(memosResponse.data.items || []);

        setError(null);
      } catch (err) {
        setError("データの取得に失敗しました");
        console.error("Error fetching data:", err);
      } finally {
        setLoading(false);
      }
    };

    fetchData();
  }, [todoIdNum, userId, user, navigate]);

  const handleDelete = async (memoId: number) => {
    try {
      await memoApi.deleteMemo(todoIdNum, memoId);
      setMemos(memos.filter((memo) => memo.id !== memoId));
    } catch (err) {
      setError("メモの削除に失敗しました");
      console.error("Error deleting memo:", err);
    }
  };

  const handleEdit = (memoId: number) => {
    navigate(`/todos/${todoIdNum}/memos/${memoId}/edit`);
  };

  return (
    <Container maxWidth="md">
      <Box sx={{ my: 4 }}>
        <Breadcrumbs aria-label="breadcrumb" sx={{ mb: 2 }}>
          <Link
            underline="hover"
            color="inherit"
            onClick={() => navigate("/todos")}
            sx={{ cursor: "pointer" }}
          >
            TODOリスト
          </Link>
          <Typography color="text.primary">{todo?.title || "メモ"}</Typography>
        </Breadcrumbs>

        <Box
          sx={{
            display: "flex",
            justifyContent: "space-between",
            alignItems: "center",
            mb: 2,
          }}
        >
          <Box sx={{ display: "flex", alignItems: "center" }}>
            <IconButton onClick={() => navigate("/todos")} sx={{ mr: 1 }}>
              <ArrowBackIcon />
            </IconButton>
            <Typography variant="h4" component="h1">
              メモリスト
            </Typography>
          </Box>
          <Button
            variant="contained"
            color="primary"
            startIcon={<AddIcon />}
            onClick={() => navigate(`/todos/${todoIdNum}/memos/new`)}
          >
            新規メモ作成
          </Button>
        </Box>

        {todo && (
          <Paper sx={{ p: 2, mb: 3 }}>
            <Typography variant="h6">{todo.title}</Typography>
            {todo.description && (
              <Typography variant="body2" color="text.secondary">
                {todo.description}
              </Typography>
            )}
          </Paper>
        )}

        {error && (
          <Alert severity="error" sx={{ mb: 2 }}>
            {error}
          </Alert>
        )}

        {loading ? (
          <Box sx={{ display: "flex", justifyContent: "center", p: 3 }}>
            <CircularProgress />
          </Box>
        ) : memos.length === 0 ? (
          <Typography variant="body1" sx={{ textAlign: "center", p: 3 }}>
            メモがありません。新しいメモを作成してください。
          </Typography>
        ) : (
          <List>
            {memos.map((memo) => (
              <React.Fragment key={memo.id}>
                <ListItem>
                  <ListItemText
                    primary={memo.content}
                    secondary={new Date(memo.created_at).toLocaleString()}
                  />
                  <Box>
                    <IconButton
                      edge="end"
                      aria-label="edit"
                      onClick={() => handleEdit(memo.id)}
                      sx={{ mr: 1 }}
                    >
                      <EditIcon />
                    </IconButton>
                    <IconButton
                      edge="end"
                      aria-label="delete"
                      onClick={() => handleDelete(memo.id)}
                    >
                      <DeleteIcon />
                    </IconButton>
                  </Box>
                </ListItem>
                <Divider />
              </React.Fragment>
            ))}
          </List>
        )}
      </Box>
    </Container>
  );
};

export default MemoList;
