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

const setAutoSuggest = (autosuggest) => {
    localStorage.setItem("auto-suggest", autosuggest);
}

const sortArray = (array, key, ascending) => {
    return array.sort((a, b) => {
        if (a[key] > b[key]) {
          return ascending ? 1 : -1;
        }
        if (a[key] < b[key]) {
          return ascending ? -1 : 1;
        }
        return 0;
    });
}

export default checkAccess;
export { hashPassword, setAutoSuggest, setLanguage, skipLogin, sortArray };