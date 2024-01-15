import axios from "axios";
import {userAccessToken} from "../utils/helpers";

const DAPS_BASE_URL = process.env.REACT_APP_API_URL;

const CATEGORIES_URL = `${DAPS_BASE_URL}api/summary`;

const options = {
  headers: {
    "Content-Type": "application/json",
    Authorization: "Bearer " + userAccessToken,
  },
};

const payload = {
  ...options,
  data: {},
  params: {},
};

const createCategory = (payload) => {
  return axios.post(CATEGORIES_URL, payload, options);
};

const getCategory = (id) => {
  return axios.get(`${CATEGORIES_URL}/${id}`, payload);
};

const getCategories = () => {
  return axios.get(CATEGORIES_URL, options);
};

const deleteCategory = (id) => {
  return axios.delete(`${CATEGORIES_URL}/${id}`, options);
};

const updateCategory = (id, payload) => {
  return axios.put(`${CATEGORIES_URL}/${id}`, payload, options);
};

const shareCategory = (id, email) => {
  return axios.put(
    `${CATEGORIES_URL}/${id}`,
    {
      shared: true,
      email: email,
    },
    options
  );
};

const unshareCategory = (id) => {
  return axios.put(
    `${CATEGORIES_URL}/${id}`,
    {
      shared: false,
    },
    options
  );
};

const CategoryService = {
  createCategory,
  getCategory,
  getCategories,
  deleteCategory,
  updateCategory,
  shareCategory,
  unshareCategory,
};

export default CategoryService;
