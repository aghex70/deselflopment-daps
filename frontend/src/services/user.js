import axios from "axios";

const REGISTER_URL = "http://localhost:11001/register";
const LOGIN_URL = "http://localhost:11001/login";
const REFRESH_TOKEN_URL = "http://localhost:11001/refresh-token";
const RECOVER_PASSWORD_URL = "http://localhost:11001/recover-password";

const options = {
  headers: {
    'Accept': 'application/json',
    'Content-Type': 'application/json',
    'Origin': 'http://localhost:3000',
    'Authorization': 'Bearer ' + localStorage.getItem("access_token")
  }
}

const register = async (name, email, password, repeat_password) => {
  return await axios.post(REGISTER_URL, {
    name,
    email,
    password,
    repeat_password,
  });
}

const login = async (email, password) => {
  return await axios
    .post(LOGIN_URL, {
      email,
      password,
    });
}

const refreshToken = async () => {
  return await axios.post(REFRESH_TOKEN_URL, {
  }, {options})
  .then((response) => {
    return response.data
  });
}

const recoverPassword = async (email) => {
  return await axios.post(RECOVER_PASSWORD_URL, {
    email,
  });
}

const getCurrentUser = () => {
  return localStorage.getItem("access_token");
}

const logout = () => {
  localStorage.removeItem("access_token");
}

const UserService = {
  register,
  login,
  refreshToken,
  recoverPassword,
  getCurrentUser,
  logout,

}

export default UserService;