import axios from "axios";

const REGISTER_URL = "http://3.75.160.227:11001/api/register";
const LOGIN_URL = "http://3.75.160.227:11001/api/login";
const REFRESH_TOKEN_URL = "http://3.75.160.227:11001/api/refresh-token";
const RECOVER_PASSWORD_URL = "http://3.75.160.227:11001/api/recover-password";

const options = {
  headers: {
    'Accept': 'application/json',
    'Content-Type': 'application/json',
    'Origin': 'http://3.75.160.227',
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