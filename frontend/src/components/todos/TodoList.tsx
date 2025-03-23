import React, { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import {
  Container,
  Typography,
  Button,
  Box,
  List,
  ListItem,
  ListItemText,
  IconButton,
  Checkbox,
  Divider,
  CircularProgress,
  Alert,
} from "@mui/material";
import {
  Add as AddIcon,
  Edit as EditIcon,
  Delete as DeleteIcon,
  Note as NoteIcon,
} from "@mui/icons-material";
import { todoApi } from "../../api/client";
import { Todo } from "../../types";
import { useAuth } from "../../contexts/AuthContext";

const TodoList: React.FC = () => {
  const [todos, setTodos] = useState<Todo[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const navigate = useNavigate();
  const { user } = useAuth();

  useEffect(() => {
    const fetchTodos = async () => {
      if (!user) {
        console.log("TodoList: No user found, redirecting to login");
        navigate("/login");
        return;
      }

      try {
        console.log(`TodoList: Fetching todos for user ID ${user.id}...`);
        setLoading(true);
        const response = await todoApi.listTodos(user.id);
        console.log("TodoList: Todos fetched successfully:", response.data);
        setTodos(response.data.items || []);
        setError(null);
      } catch (err) {
        console.error("TodoList: Error fetching todos:", err);
        setError("TODOの取得に失敗しました");
      } finally {
        setLoading(false);
      }
    };

    fetchTodos();
  }, [user, navigate]);

  const handleDelete = async (todoId: number) => {
    try {
      console.log(`TodoList: Deleting todo with ID ${todoId}`);
      await todoApi.deleteTodo(todoId);
      console.log("TodoList: Todo deleted successfully");
      setTodos(todos.filter((todo) => todo.id !== todoId));
    } catch (err) {
      console.error("TodoList: Error deleting todo:", err);
      setError("TODOの削除に失敗しました");
    }
  };

  const handleToggleComplete = async (todo: Todo) => {
    try {
      console.log(`TodoList: Toggling complete status for todo ID ${todo.id}`);
      const updatedTodo = { ...todo, completed: !todo.completed };
      await todoApi.updateTodo(todo.id, { completed: !todo.completed });
      console.log("TodoList: Todo updated successfully");
      setTodos(
        todos.map((t) =>
          t.id === todo.id ? { ...t, completed: !t.completed } : t,
        ),
      );
    } catch (err) {
      console.error("TodoList: Error updating todo:", err);
      setError("TODOの更新に失敗しました");
    }
  };

  return (
    <Container maxWidth="md">
      <Box sx={{ my: 4 }}>
        <Box
          sx={{
            display: "flex",
            justifyContent: "space-between",
            alignItems: "center",
            mb: 3,
          }}
        >
          <Typography variant="h4" component="h1">
            TODOリスト
          </Typography>
          <Button
            variant="contained"
            color="primary"
            startIcon={<AddIcon />}
            onClick={() => navigate("/todos/new")}
          >
            新規TODO
          </Button>
        </Box>

        {error && (
          <Alert severity="error" sx={{ mb: 2 }}>
            {error}
          </Alert>
        )}

        {loading ? (
          <Box sx={{ display: "flex", justifyContent: "center", p: 3 }}>
            <CircularProgress />
          </Box>
        ) : todos.length === 0 ? (
          <Typography variant="body1" sx={{ textAlign: "center", p: 3 }}>
            TODOがありません。新しいTODOを作成してください。
          </Typography>
        ) : (
          <List>
            {todos.map((todo) => (
              <React.Fragment key={todo.id}>
                <ListItem>
                  <Checkbox
                    edge="start"
                    checked={todo.completed}
                    onChange={() => handleToggleComplete(todo)}
                    sx={{ mr: 1 }}
                  />
                  <ListItemText
                    primary={todo.title}
                    secondary={todo.description}
                    sx={{
                      textDecoration: todo.completed ? "line-through" : "none",
                    }}
                  />
                  <Box>
                    <IconButton
                      edge="end"
                      aria-label="notes"
                      onClick={() => navigate(`/todos/${todo.id}/memos`)}
                      sx={{ mr: 1 }}
                    >
                      <NoteIcon />
                    </IconButton>
                    <IconButton
                      edge="end"
                      aria-label="edit"
                      onClick={() => navigate(`/todos/${todo.id}/edit`)}
                      sx={{ mr: 1 }}
                    >
                      <EditIcon />
                    </IconButton>
                    <IconButton
                      edge="end"
                      aria-label="delete"
                      onClick={() => handleDelete(todo.id)}
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

export default TodoList;
