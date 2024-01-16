import axios from "axios";
import {userAccessToken} from "../utils/helpers";

const DAPS_BASE_URL = process.env.REACT_APP_API_URL;

const TODOS_URL = `${DAPS_BASE_URL}api/todos`;
const RECURRING_TODOS_URL = `${DAPS_BASE_URL}api/recurring-todos`;
const COMPLETED_TODOS_URL = `${DAPS_BASE_URL}api/completed-todos`;
const SUGGESTED_TODOS_URL = `${DAPS_BASE_URL}api/suggested-todos`;
const SUGGEST_TODOS_URL = `${DAPS_BASE_URL}api/suggest`;

const options = {
  headers: {
    "Content-Type": "application/json",
    Authorization: "Bearer " + userAccessToken,
  },
};

const createTodo = (payload) => {
  return axios.post(TODOS_URL, payload, options);
};

const getTodo = (id) => {
  return axios.get(`${TODOS_URL}/${id}`, options);
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

const deleteTodo = (id) => {
  return axios.delete(`${TODOS_URL}/${id}`, options);
};

const updateTodo = (id, payload) => {
  return axios.put(`${TODOS_URL}/${id}`, payload, options);
};

const completeTodo = (id) => {
  return axios.post(
    `${TODOS_URL}/${id}/complete`,
    {},
    options
  );
};

const activateTodo = (id) => {
  return axios.post(
    `${TODOS_URL}/${id}/activate`,
    {},
    options
  );
};

const startTodo = (id) => {
  return axios.post(
    `${TODOS_URL}/${id}/start`,
    {},
    options
  );
};

const restartTodo = (id) => {
  return axios.post(
    `${TODOS_URL}/${id}/restart`,
    {},
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
