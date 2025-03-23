import React, { useState, useEffect } from 'react';
import { useNavigate, useParams } from 'react-router-dom';
import { 
  Container, Typography, TextField, Button, Box, 
  CircularProgress, Alert 
} from '@mui/material';
import { todoApi } from '../../api/client';
import { TodoCreatePayload, TodoUpdatePayload } from '../../types';

const TodoForm: React.FC = () => {
  const { id } = useParams<{ id: string }>();
  const isEditMode = !!id;
  const todoId = id ? parseInt(id) : 0;
  
  const [title, setTitle] = useState('');
  const [description, setDescription] = useState('');
  const [loading, setLoading] = useState(isEditMode);
  const [submitting, setSubmitting] = useState(false);
  const [error, setError] = useState<string | null>(null);
  
  const navigate = useNavigate();

  useEffect(() => {
    const fetchTodo = async () => {
      if (!isEditMode) return;
      
      try {
        setLoading(true);
        const response = await todoApi.getTodo(todoId);
        const todo = response.data;
        setTitle(todo.title);
        setDescription(todo.description || '');
      } catch (err) {
        setError('TODOの取得に失敗しました');
        console.error('Error fetching todo:', err);
      } finally {
        setLoading(false);
      }
    };

    fetchTodo();
  }, [isEditMode, todoId]);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    
    if (!title.trim()) {
      setError('タイトルを入力してください');
      return;
    }

    try {
      setSubmitting(true);
      
      if (isEditMode) {
        const updatePayload: TodoUpdatePayload = {
          title,
          description: description || undefined
        };
        await todoApi.updateTodo(todoId, updatePayload);
      } else {
        const createPayload: TodoCreatePayload = {
          title,
          description: description || undefined
        };
        await todoApi.createTodo(createPayload);
      }
      
      navigate('/todos');
    } catch (err) {
      setError(`TODOの${isEditMode ? '更新' : '作成'}に失敗しました`);
      console.error(`Error ${isEditMode ? 'updating' : 'creating'} todo:`, err);
    } finally {
      setSubmitting(false);
    }
  };

  if (loading) {
    return (
      <Container maxWidth="sm">
        <Box sx={{ display: 'flex', justifyContent: 'center', p: 5 }}>
          <CircularProgress />
        </Box>
      </Container>
    );
  }

  return (
    <Container maxWidth="sm">
      <Box sx={{ my: 4 }}>
        <Typography variant="h4" component="h1" gutterBottom>
          {isEditMode ? 'TODOを編集' : '新しいTODOを作成'}
        </Typography>
        
        {error && <Alert severity="error" sx={{ mb: 2 }}>{error}</Alert>}
        
        <Box component="form" onSubmit={handleSubmit} sx={{ mt: 2 }}>
          <TextField
            fullWidth
            label="タイトル"
            value={title}
            onChange={(e) => setTitle(e.target.value)}
            margin="normal"
            required
          />
          <TextField
            fullWidth
            label="説明"
            value={description}
            onChange={(e) => setDescription(e.target.value)}
            margin="normal"
            multiline
            rows={4}
          />
          <Box sx={{ mt: 3, display: 'flex', justifyContent: 'space-between' }}>
            <Button 
              variant="outlined" 
              onClick={() => navigate('/todos')}
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
              {submitting ? <CircularProgress size={24} /> : (isEditMode ? '更新' : '作成')}
            </Button>
          </Box>
        </Box>
      </Box>
    </Container>
  );
};

export default TodoForm; 