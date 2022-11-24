import axios from "axios";

const TODO_URL = "http://api/api/todo";
const TODOS_URL = "http://api/api/todos";
const RECURRING_TODOS_URL = "http://api/api/recurring-todos";
const COMPLETED_TODOS_URL = "http://api/api/completed-todos";

const options = {
  headers: {
    'Accept': 'application/json',
    'Content-Type': 'application/json',
    'Origin': 'http://deselflopment.com',
    'Authorization': 'Bearer ' + localStorage.getItem("access_token")
  }
}

const createTodo = (payload) => {
  return axios.post(TODO_URL, payload, options);
}

const getTodo = (id, categoryId) => {
  return axios.get(`${TODO_URL}/${id}?category_id=${categoryId}`, options);
}

const getTodos = (id) => {
  return axios.get(TODOS_URL, {
    ...options,
    params: {
      category_id: id,
    }
  });
}

const getRecurringTodos = () => {
  return axios.get(RECURRING_TODOS_URL, options);
}

const getCompletedTodos = () => {
  return axios.get(COMPLETED_TODOS_URL, options);
}

const deleteTodo = (id, categoryId) => {
  return axios.delete(`${TODO_URL}/${id}?category_id=${categoryId}`, options);
}

const updateTodo = (id, payload) => {
  return axios.put(`${TODO_URL}/${id}`,
    payload
  , options);
}

const completeTodo = (id, categoryId) => {
  return axios.put(`${TODO_URL}/${id}/complete`, {
    category_id: categoryId,
  }, options);
}

const activateTodo = (id, categoryId) => {
  return axios.put(`${TODO_URL}/${id}/activate`, {
    category_id: categoryId,
  }, options);
}

const startTodo = (id, categoryId, payload) => {
  return axios.put(`${TODO_URL}/${id}/start`, {
    ...payload,
    category_id: categoryId,
  }, options);
}

const TodoService = {
  createTodo,
  getTodo,
  getTodos,
  getRecurringTodos,
  getCompletedTodos,
  deleteTodo,
  updateTodo,
  completeTodo,
  activateTodo,
  startTodo,
}

export default TodoService;