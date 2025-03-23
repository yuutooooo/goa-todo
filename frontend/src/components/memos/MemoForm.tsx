import React, { useState, useEffect } from "react";
import { useNavigate, useParams } from "react-router-dom";
import {
  Container,
  Typography,
  TextField,
  Button,
  Box,
  CircularProgress,
  Alert,
  Breadcrumbs,
  Link,
  IconButton,
} from "@mui/material";
import { ArrowBack as ArrowBackIcon } from "@mui/icons-material";
import { memoApi, todoApi } from "../../api/client";
import { useAuth } from "../../contexts/AuthContext";

const MemoForm: React.FC = () => {
  const { todoId, memoId } = useParams<{ todoId: string; memoId: string }>();
  const todoIdNum = parseInt(todoId || "0");
  const memoIdNum = memoId ? parseInt(memoId) : 0;
  const isEditMode = !!memoId;

  const { user } = useAuth();
  const userId = user?.id || 0;

  const [content, setContent] = useState("");
  const [todoTitle, setTodoTitle] = useState("");
  const [loading, setLoading] = useState(isEditMode);
  const [submitting, setSubmitting] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const navigate = useNavigate();

  useEffect(() => {
    if (!user) {
      navigate("/login");
      return;
    }

    const fetchData = async () => {
      try {
        setLoading(true);

        // TODOのタイトルを取得
        const todoResponse = await todoApi.getTodo(todoIdNum);
        setTodoTitle(todoResponse.data.title);

        // メモ編集時の場合、現在のメモを取得
        if (isEditMode) {
          const memoResponse = await memoApi.getMemo(todoIdNum, memoIdNum);
          setContent(memoResponse.data.content);
        }
      } catch (err) {
        setError("データの取得に失敗しました");
        console.error("Error fetching data:", err);
      } finally {
        setLoading(false);
      }
    };

    fetchData();
  }, [isEditMode, todoIdNum, memoIdNum, user, navigate]);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    if (!content.trim()) {
      setError("内容を入力してください");
      return;
    }

    try {
      setSubmitting(true);

      if (isEditMode) {
        await memoApi.updateMemo(todoIdNum, memoIdNum, { content });
      } else {
        await memoApi.createMemo(todoIdNum, { content });
      }

      navigate(`/todos/${todoIdNum}/memos`);
    } catch (err) {
      setError(`メモの${isEditMode ? "更新" : "作成"}に失敗しました`);
      console.error(`Error ${isEditMode ? "updating" : "creating"} memo:`, err);
    } finally {
      setSubmitting(false);
    }
  };

  if (loading) {
    return (
      <Container maxWidth="sm">
        <Box sx={{ display: "flex", justifyContent: "center", p: 5 }}>
          <CircularProgress />
        </Box>
      </Container>
    );
  }

  return (
    <Container maxWidth="sm">
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
          <Link
            underline="hover"
            color="inherit"
            onClick={() => navigate(`/todos/${todoIdNum}/memos`)}
            sx={{ cursor: "pointer" }}
          >
            {todoTitle} のメモ
          </Link>
          <Typography color="text.primary">
            {isEditMode ? "メモを編集" : "新規メモ"}
          </Typography>
        </Breadcrumbs>

        <Box sx={{ display: "flex", alignItems: "center", mb: 3 }}>
          <IconButton
            onClick={() => navigate(`/todos/${todoIdNum}/memos`)}
            sx={{ mr: 1 }}
          >
            <ArrowBackIcon />
          </IconButton>
          <Typography variant="h4" component="h1">
            {isEditMode ? "メモを編集" : "新規メモ作成"}
          </Typography>
        </Box>

        {error && (
          <Alert severity="error" sx={{ mb: 2 }}>
            {error}
          </Alert>
        )}

        <Box component="form" onSubmit={handleSubmit} sx={{ mt: 2 }}>
          <TextField
            fullWidth
            label="メモ内容"
            value={content}
            onChange={(e) => setContent(e.target.value)}
            margin="normal"
            required
            multiline
            rows={6}
          />
          <Box sx={{ mt: 3, display: "flex", justifyContent: "space-between" }}>
            <Button
              variant="outlined"
              onClick={() => navigate(`/todos/${todoIdNum}/memos`)}
              disabled={submitting}
            >
              キャンセル
            </Button>
            <Button
              type="submit"
              variant="contained"
              color="primary"
              disabled={submitting}
            >
              {submitting ? (
                <CircularProgress size={24} />
              ) : isEditMode ? (
                "更新"
              ) : (
                "作成"
              )}
            </Button>
          </Box>
        </Box>
      </Box>
    </Container>
  );
};

export default MemoForm;
