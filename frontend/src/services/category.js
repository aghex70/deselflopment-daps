import axios from "axios";
import {userAccessToken} from "../utils/helpers";
import {CancelButtonText, ShareButtonText, ShareCategoryHeaderText} from "../utils/texts";

const DAPS_BASE_URL = process.env.REACT_APP_API_URL;

const SUMMARY_URL = `${DAPS_BASE_URL}api/summary`;
const CATEGORIES_URL = `${DAPS_BASE_URL}api/categories`;

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

const getCategoryUsers = (id) => {
  return axios.get(`${CATEGORIES_URL}/${id}/users/`, payload);
};

const getCategories = () => {
  return axios.get(SUMMARY_URL, options);
};

const deleteCategory = (id) => {
  return axios.delete(`${CATEGORIES_URL}/${id}`, options);
};

const updateCategory = (id, payload) => {
  return axios.put(`${CATEGORIES_URL}/${id}`, payload, options);
};

const shareCategory = (id, email) => {
  return axios.post(
    `${CATEGORIES_URL}/${id}/share`,
    {
      email: email,
    },
    options
  );
};

const unshareCategory = (id, email) => {
  return axios.post(
    `${CATEGORIES_URL}/${id}/unshare`,
    {
      email: email,
    },
    options
  );
};

const CategoryService = {
  createCategory,
  getCategory,
  getCategoryUsers,
  getCategories,
  deleteCategory,
  updateCategory,
  shareCategory,
  unshareCategory,
};

export default CategoryService;