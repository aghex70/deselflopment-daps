import axios from "axios";

const DAPS_BASE_URL = process.env.REACT_APP_API_URL

const USER_CONFIGURATION_URL = `${DAPS_BASE_URL}api/user-configuration`;

const options = {
  headers: {
    'Accept': 'application/json',
    'Content-Type': 'application/json',
    'Authorization': 'Bearer ' + localStorage.getItem("access_token")
  }
}

const getUserConfiguration = () => {
  return axios.get(`${USER_CONFIGURATION_URL}/`, options);
}

const updateUserConfiguration = (payload) => {
  return axios.put(`${USER_CONFIGURATION_URL}/`,
      payload
      , options);
}

const UserConfigurationService = {
  getUserConfiguration,
  updateUserConfiguration,
}

export default UserConfigurationService;