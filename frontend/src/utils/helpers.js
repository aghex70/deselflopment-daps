import bcrypt from 'bcryptjs';

// Create a function that ciphers a string using bcryptjs and store it in variable agp
const hashPassword = (string) => {
    return bcrypt.hashSync(string, 10);
}

const checkAccess = () => {
  const token = getUserToken();
  if (!token) {
    window.location.href = "/login";
  }
}

// const checkSuperUser = () => {
//   const token = getUserToken();
//   if (!token) {
//     window.location.href = "/login";
//   }
// }

const skipLogin = () => {
  if (getUserToken()) {
    window.location.href = "/categories";
  }
}
const getUserToken = () => {
  return localStorage.getItem("access_token");
}

const setLanguage = (language) => {
  localStorage.setItem("language", language);
}

export default checkAccess;
export { hashPassword, setLanguage, skipLogin };