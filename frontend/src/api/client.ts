import axios from "axios";

const API_URL = "http://localhost:8000";

const apiClient = axios.create({
  baseURL: API_URL,
  headers: {
    "Content-Type": "application/json",
  },
  withCredentials: true,
});

// レスポンスインターセプタ（デバッグログ記録）
apiClient.interceptors.response.use(
  (response) => {
    console.log(
      `APIClient: [${
        response.status
      }] ${response.config.method?.toUpperCase()} ${response.config.url}`,
      response.data
    );
    return response;
  },
  (error) => {
    console.error(
      `APIClient: Error [${
        error.response?.status || "NETWORK"
      }] ${error.config?.method?.toUpperCase()} ${error.config?.url}`,
      error.response?.data || error.message
    );
    return Promise.reject(error);
  }
);

// リクエストインターセプター（認証関連のログ出力）
apiClient.interceptors.request.use(
  (config) => {
    // トークン認証は使用しない
    console.log(
      `APIClient: Request to ${config.method?.toUpperCase()} ${config.url}`
    );
    return config;
  },
  (error) => {
    console.error("APIClient: Request error:", error);
    return Promise.reject(error);
  }
);

// ユーザー関連API
export const userApi = {
  // ユーザー登録
  register: (data: { name: string; email: string; password: string }) =>
    apiClient.post("/users", data),

  // ログイン
  login: (data: { email: string; password: string }) =>
    apiClient.post("/login", data),

  // ユーザー情報取得
  getUser: (userId: number) => apiClient.get(`/users/${userId}`),

  // ユーザー情報更新
  updateUser: (userId: number, data: { name?: string; email?: string }) =>
    apiClient.put(`/users/${userId}`, data),

  // ユーザー削除
  deleteUser: (userId: number) => apiClient.delete(`/users/${userId}`),
};

// Todo関連API
export const todoApi = {
  // Todo一覧取得
  listTodos: (userId?: number) => {
    // OpenAPI仕様に従い、user_idが必要
    const currentUser = JSON.parse(localStorage.getItem("user") || "{}");
    const userIdToUse = userId || currentUser.id;
    if (!userIdToUse) {
      console.error("APIClient: User ID not available for listTodos");
      return Promise.reject(new Error("User ID is required"));
    }
    return apiClient.get(`/users/${userIdToUse}/todos`);
  },

  // Todo詳細取得
  getTodo: (todoId: number) => {
    const currentUser = JSON.parse(localStorage.getItem("user") || "{}");
    if (!currentUser.id) {
      console.error("APIClient: User ID not available for getTodo");
      return Promise.reject(new Error("User ID is required"));
    }
    return apiClient.get(`/users/${currentUser.id}/todos/${todoId}`);
  },

  // Todo作成
  createTodo: (data: { title: string; description?: string }) => {
    const currentUser = JSON.parse(localStorage.getItem("user") || "{}");
    if (!currentUser.id) {
      console.error("APIClient: User ID not available for createTodo");
      return Promise.reject(new Error("User ID is required"));
    }
    return apiClient.post(`/users/${currentUser.id}/todos`, data);
  },

  // Todo更新
  updateTodo: (
    todoId: number,
    data: { title?: string; description?: string; completed?: boolean }
  ) => {
    const currentUser = JSON.parse(localStorage.getItem("user") || "{}");
    if (!currentUser.id) {
      console.error("APIClient: User ID not available for updateTodo");
      return Promise.reject(new Error("User ID is required"));
    }
    return apiClient.put(`/users/${currentUser.id}/todos/${todoId}`, data);
  },

  // Todo削除
  deleteTodo: (todoId: number) => {
    const currentUser = JSON.parse(localStorage.getItem("user") || "{}");
    if (!currentUser.id) {
      console.error("APIClient: User ID not available for deleteTodo");
      return Promise.reject(new Error("User ID is required"));
    }
    return apiClient.delete(`/users/${currentUser.id}/todos/${todoId}`);
  },
};

// メモ関連API
export const memoApi = {
  // メモ一覧取得
  getMemos: (todoId: number) => {
    const currentUser = JSON.parse(localStorage.getItem("user") || "{}");
    if (!currentUser.id) {
      console.error("APIClient: User ID not available for getMemos");
      return Promise.reject(new Error("User ID is required"));
    }
    return apiClient.get(`/users/${currentUser.id}/todos/${todoId}/memos`);
  },

  // メモ詳細取得
  getMemo: (todoId: number, memoId: number) => {
    const currentUser = JSON.parse(localStorage.getItem("user") || "{}");
    if (!currentUser.id) {
      console.error("APIClient: User ID not available for getMemo");
      return Promise.reject(new Error("User ID is required"));
    }
    return apiClient.get(
      `/users/${currentUser.id}/todos/${todoId}/memos/${memoId}`
    );
  },

  // メモ作成
  createMemo: (todoId: number, data: { content: string }) => {
    const currentUser = JSON.parse(localStorage.getItem("user") || "{}");
    if (!currentUser.id) {
      console.error("APIClient: User ID not available for createMemo");
      return Promise.reject(new Error("User ID is required"));
    }
    return apiClient.post(
      `/users/${currentUser.id}/todos/${todoId}/memos`,
      data
    );
  },

  // メモ更新
  updateMemo: (todoId: number, memoId: number, data: { content: string }) => {
    const currentUser = JSON.parse(localStorage.getItem("user") || "{}");
    if (!currentUser.id) {
      console.error("APIClient: User ID not available for updateMemo");
      return Promise.reject(new Error("User ID is required"));
    }
    return apiClient.put(
      `/users/${currentUser.id}/todos/${todoId}/memos/${memoId}`,
      data
    );
  },

  // メモ削除
  deleteMemo: (todoId: number, memoId: number) => {
    const currentUser = JSON.parse(localStorage.getItem("user") || "{}");
    if (!currentUser.id) {
      console.error("APIClient: User ID not available for deleteMemo");
      return Promise.reject(new Error("User ID is required"));
    }
    return apiClient.delete(
      `/users/${currentUser.id}/todos/${todoId}/memos/${memoId}`
    );
  },
};

export default apiClient;
