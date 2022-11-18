import axios from "axios";
// import requestHeaders from "../common/Headers";

const CATEGORY_URL = "http://localhost:11001/category";
const CATEGORIES_URL = "http://localhost:11001/summary";

const options = {
  headers: {
    'Accept': 'application/json',
    'Content-Type': 'application/json',
    'Origin': 'http://localhost:3000',
    'Authorization': 'Bearer ' + localStorage.getItem("access_token")
  }
}

const payload = {
  ...options,
  data: {},
  params: {},
}

const createCategory = (payload) => {
  return axios.post(CATEGORY_URL, payload, options);
}

const getCategory = (id) => {
  return axios.get(`${CATEGORY_URL}/${id}`, payload);
}

const getCategories = () => {
  return axios.get(CATEGORIES_URL, options);
}

const deleteCategory = (id) => {
  return axios.delete(`${CATEGORY_URL}/${id}`, options);
}

const updateCategory = (id, payload) => {
  return axios.put(`${CATEGORY_URL}/${id}`, payload, options);
}

const shareCategory = (id, email) => {
  return axios.put(`${CATEGORY_URL}/${id}`, {
    shared: true,
    email: email,
  }, options);
}

const unshareCategory = (id) => {
  return axios.put(`${CATEGORY_URL}/${id}`, {
    shared: false,
  }, options);
}

const CategoryService = {
  createCategory,
  getCategory,
  getCategories,
  deleteCategory,
  updateCategory,
  shareCategory,
  unshareCategory,
}

export default CategoryService;