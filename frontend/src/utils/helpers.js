import bcrypt from 'bcryptjs';

const hashPassword = (string) => {
    return bcrypt.hashSync(string, 10);
}

const checkAccess = () => {
  const token = getUserToken();
  if (!token) {
    goToLogin();
  }
}

const skipLogin = () => {
  if (getUserToken()) {
    goToCategories();
  }
}
const getUserToken = () => {
  return localStorage.getItem("access_token");
}

const clearLocalStorage = (excludedKeys= []) => {
    let baseExcludedKeys = ["access_token", "language", "auto-suggest"];
    let untouchedKeys = baseExcludedKeys.concat(excludedKeys);
    for (let key in localStorage) {
        if (!untouchedKeys.includes(key)) {
            localStorage.removeItem(key);
        }
    }
}

const setLanguage = (language) => {
  localStorage.setItem("language", language);
}

const setAutoSuggest = (autosuggest) => {
    localStorage.setItem("auto-suggest", autosuggest);
}

const goToCreateCategory = () => {
    window.location.href = "/create-category";
}

const goToCategories = () => {
    clearLocalStorage([]);
    window.location.href = "/categories";
}

const goToReportABug = () => {
    clearLocalStorage([]);
    window.location.href = "/report-bug";
}

const goToRecurringTodos = () => {
    clearLocalStorage([]);
    window.location.href = "/recurring-todos";
}

const goToSuggestedTodos = () => {
    clearLocalStorage([]);
    window.location.href = "/suggested-todos";
}

const goToCompletedTodos = () => {
    clearLocalStorage([]);
    window.location.href = "/completed-todos";
}

const goToProfile = () => {
    clearLocalStorage([]);
    window.location.href = "/profile";
}

const goToProvisionDemoUser = () => {
    clearLocalStorage([]);
    window.location.href = "/provision";
}

const goToListOfUsers = () => {
    clearLocalStorage([]);
    window.location.href = "/users";
}

const goToImportTodos = () => {
    clearLocalStorage([]);
    window.location.href = "/import";
}

const goToLogout = () => {
    clearLocalStorage([]);
    window.location.href = "/logout";
}

const goToLogin = () => {
    clearLocalStorage([]);
    window.location.href = "/login";
}

const goToRegister = () => {
    clearLocalStorage([]);
    window.location.href = "/register";
}

const sortArrayByField = (array, field, ascending) => {
    return array.sort((a, b) => {
        if (a[field] > b[field]) {
          return ascending ? 1 : -1;
        }
        if (a[field] < b[field]) {
          return ascending ? -1 : 1;
        }
        return 0;
    });


}

export default checkAccess;

export {
    hashPassword,
    setAutoSuggest,
    setLanguage,
    skipLogin,
    goToCreateCategory,
    goToCategories,
    goToReportABug,
    goToRecurringTodos,
    goToSuggestedTodos,
    goToCompletedTodos,
    goToProfile,
    goToProvisionDemoUser,
    goToListOfUsers,
    goToImportTodos,
    goToLogout,
    goToLogin,
    goToRegister,
    sortArrayByField,
    clearLocalStorage,
};