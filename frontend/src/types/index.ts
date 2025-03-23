// ユーザー関連の型
export interface User {
  id: number;
  name: string;
  email: string;
  created_at: string;
  updated_at: string;
}

export interface LoginPayload {
  email: string;
  password: string;
}

export interface RegisterPayload {
  name: string;
  email: string;
  password: string;
}

// Todo関連の型
export interface Todo {
  id: number;
  user_id: number;
  title: string;
  description: string;
  completed: boolean;
  created_at: string;
  updated_at: string;
}

export interface TodoCollection {
  items: Todo[];
}

export interface TodoCreatePayload {
  title: string;
  description?: string;
  completed?: boolean;
}

export interface TodoUpdatePayload {
  title?: string;
  description?: string;
  completed?: boolean;
}

// メモ関連の型
export interface Memo {
  id: number;
  todo_id: number;
  content: string;
  created_at: string;
  updated_at: string;
}

export interface MemoCollection {
  items: Memo[];
}

export interface MemoCreatePayload {
  content: string;
}

export interface MemoUpdatePayload {
  content: string;
}

// エラーレスポンスの型
export interface APIError {
  name: string;
  message: string;
}
