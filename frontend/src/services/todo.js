import axios from "axios";

const DAPS_BASE_URL = process.env.REACT_APP_API_URL;

const TODO_URL = `${DAPS_BASE_URL}api/todo`;
const TODOS_URL = `${DAPS_BASE_URL}api/todos`;
const RECURRING_TODOS_URL = `${DAPS_BASE_URL}api/recurring-todos`;
const COMPLETED_TODOS_URL = `${DAPS_BASE_URL}api/completed-todos`;
const SUGGESTED_TODOS_URL = `${DAPS_BASE_URL}api/suggested-todos`;
const SUGGEST_TODOS_URL = `${DAPS_BASE_URL}api/suggest`;

const options = {
  headers: {
    "Content-Type": "application/json",
    Authorization: "Bearer " + localStorage.getItem("access_token"),
  },
};

const createTodo = (payload) => {
  return axios.post(TODO_URL, payload, options);
};

const getTodo = (id, categoryId) => {
  return axios.get(`${TODO_URL}/${id}?category_id=${categoryId}`, options);
};

const suggestTodos = () => {
  return axios.post(SUGGEST_TODOS_URL, {}, options);
};

const getTodos = (id) => {
  return axios.get(TODOS_URL, {
    ...options,
    params: {
      category_id: id,
    },
  });
};

const getRecurringTodos = () => {
  return axios.get(RECURRING_TODOS_URL, options);
};

const getCompletedTodos = () => {
  return axios.get(COMPLETED_TODOS_URL, options);
};

const getSuggestedTodos = () => {
  return axios.get(SUGGESTED_TODOS_URL, options);
};

const deleteTodo = (id, categoryId) => {
  return axios.delete(`${TODO_URL}/${id}?category_id=${categoryId}`, options);
};

const updateTodo = (id, payload) => {
  return axios.put(`${TODO_URL}/${id}`, payload, options);
};

const completeTodo = (id, categoryId) => {
  return axios.put(
    `${TODO_URL}/${id}/complete`,
    {
      category_id: categoryId,
    },
    options
  );
};

const activateTodo = (id, categoryId) => {
  return axios.put(
    `${TODO_URL}/${id}/activate`,
    {
      category_id: categoryId,
    },
    options
  );
};

const startTodo = (id, categoryId, payload) => {
  return axios.put(
    `${TODO_URL}/${id}/start`,
    {
      ...payload,
      category_id: categoryId,
    },
    options
  );
};

const restartTodo = (id, categoryId, payload) => {
  return axios.put(
    `${TODO_URL}/${id}/restart`,
    {
      ...payload,
      category_id: categoryId,
    },
    options
  );
};

const TodoService = {
  createTodo,
  getTodo,
  getTodos,
  getRecurringTodos,
  getCompletedTodos,
  getSuggestedTodos,
  deleteTodo,
  updateTodo,
  completeTodo,
  activateTodo,
  startTodo,
  restartTodo,
  suggestTodos,
};

export default TodoService;
