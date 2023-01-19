import bcrypt from 'bcryptjs';

// Function that ciphers a string using bcryptjs and store it in variable agp
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

const clearLocalStorage = (keys) => {
    for (let key in localStorage) {
        if (!keys.includes(key)) {
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

const goToCreateCategory = () => {
    window.location.href = "/create-category";
}

const goToCategories = () => {
    window.location.href = "/categories";
}

const goToReportABug = () => {
    window.location.href = "/report-bug";
}

const goToRecurringTodos = () => {
    window.location.href = "/recurring-todos";
}

const goToSuggestedTodos = () => {
    window.location.href = "/suggested-todos";
}

const goToCompletedTodos = () => {
    window.location.href = "/completed-todos";
}

const goToProfile = () => {
    window.location.href = "/profile";
}

const goToProvisionDemoUser = () => {
    window.location.href = "/provision";
}

const goToListOfUsers = () => {
    window.location.href = "/users";
}

const goToImportTodos = () => {
    window.location.href = "/import";
}

const goToLogout = () => {
    window.location.href = "/logout";
}

const goToLogin = () => {
    window.location.href = "/login";
}

const goToRegister = () => {
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

export {
    hashPassword,
    setAutoSuggest,
    setLanguage,
    skipLogin,
    sortArray,
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
    checkAccess,
    clearLocalStorage,
};