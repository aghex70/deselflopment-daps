import axios from "axios";

const DAPS_BASE_URL = process.env.REACT_APP_API_URL;

const REGISTER_URL = `${DAPS_BASE_URL}api/register`;
const LOGIN_URL = `${DAPS_BASE_URL}api/login`;
const REFRESH_TOKEN_URL = `${DAPS_BASE_URL}api/refresh-token`;
const RECOVER_PASSWORD_URL = `${DAPS_BASE_URL}api/recover-password`;
const USER_URL = `${DAPS_BASE_URL}api/user`;
const USERS_URL = `${DAPS_BASE_URL}api/users`;
const ADMIN_URL = `${DAPS_BASE_URL}api/user/admin`;
const PROVISION_DEMO_USER_URL = `${DAPS_BASE_URL}api/user/provision`;
const IMPORT_CSV_URL = `${DAPS_BASE_URL}api/import`;
const ACTIVATE_USER_URL = `${DAPS_BASE_URL}api/user/activate`;
const REFRESH_ACTIVATION_CODE_URL = `${DAPS_BASE_URL}api/user/refresh-activation-code`;
const RESET_PASSWORD_URL = `${DAPS_BASE_URL}api/reset-password`;
const RESET_LINK_URL = `${DAPS_BASE_URL}api/reset-link`;

const options = {
  headers: {
    "Content-Type": "application/json",
    Authorization: "Bearer " + localStorage.getItem("access_token"),
  },
};

const fileOptions = {
  headers: {
    "Content-Type": "multipart/form-data",
    Authorization: "Bearer " + localStorage.getItem("access_token"),
  },
};

const register = async (name, email, password) => {
  return await axios.post(REGISTER_URL, {
    name,
    email,
    password,
  });
};

const login = async (email, password) => {
  return await axios.post(LOGIN_URL, {
    email,
    password,
  });
};

const deleteUser = async (id) => {
  return await axios.delete(`${USER_URL}/${id}`, options);
};

const getUser = async (id) => {
  return await axios.get(`${USER_URL}/${id}`, options);
};

const getUsers = async () => {
  return await axios.get(USERS_URL, options);
};

const checkAdminAccess = async () => {
  return axios.post(ADMIN_URL, {}, options);
};

const refreshToken = async () => {
  return await axios.post(REFRESH_TOKEN_URL, {}, options).then((response) => {
    return response.data;
  });
};

const recoverPassword = async (email) => {
  return await axios.post(RECOVER_PASSWORD_URL, {
    email,
  });
};

const getCurrentUser = () => {
  return localStorage.getItem("access_token");
};

const logout = () => {
  localStorage.removeItem("access_token");
};

const provisionDemoUser = async (email, password, language) => {
  return await axios.post(
    PROVISION_DEMO_USER_URL,
    {
      email,
      password,
      language,
    },
    options
  );
};

const importCSV = async (file) => {
  const formData = new FormData();
  formData.append("todos.csv", file);
  return await axios.post(IMPORT_CSV_URL, formData, fileOptions);
};

const activateUser = async (uuid) => {
  return await axios.post(`${ACTIVATE_USER_URL}`, {
    activation_code: uuid,
  });
};

const refreshActivationCode = async (uuid) => {
  return await axios.post(`${REFRESH_ACTIVATION_CODE_URL}`, {
    activation_code: uuid,
  });
};

const createResetLink = async (email) => {
  return await axios.post(`${RESET_LINK_URL}`, {
    email,
  });
};

const resetPassword = async (uuid, password) => {
  return await axios.post(`${RESET_PASSWORD_URL}`, {
    reset_password_code: uuid,
    password,
    repeat_password: password,
  });
};

const UserService = {
  register,
  login,
  refreshToken,
  recoverPassword,
  getCurrentUser,
  logout,
  provisionDemoUser,
  checkAdminAccess,
  getUsers,
  getUser,
  deleteUser,
  importCSV,
  activateUser,
  refreshActivationCode,
  createResetLink,
  resetPassword,
};

export default UserService;
