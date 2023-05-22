import bcrypt from "bcryptjs";

const hashPassword = (string) => {
  return bcrypt.hashSync(string, 10);
};

const checkAccess = () => {
  if (!getUserToken() || !getUserId()) {
    goToLogin();
  }
};

const skipLogin = () => {
  if (getUserToken() && getUserId()) {
    goToCategories();
  }
};

const getUserToken = () => {
  return localStorage.getItem("access_token");
};

const checkValidToken = (error) => {
  if (error.response.data.message === "signature is invalid") {
    localStorage.removeItem("access_token");
    localStorage.removeItem("user_id");
    goToLogin();
  }
};

const getUserId = () => {
  return localStorage.getItem("user_id");
};

const clearLocalStorage = (excludedKeys = []) => {
  let baseExcludedKeys = [
    "access_token",
    "language",
    "auto-suggest",
    "user_id",
  ];
  let untouchedKeys = baseExcludedKeys.concat(excludedKeys);
  for (let key in localStorage) {
    if (!untouchedKeys.includes(key)) {
      localStorage.removeItem(key);
    }
  }
};

const setLanguage = (language) => {
  localStorage.setItem("language", language);
};

const setAutoSuggest = (autoSuggest) => {
  localStorage.setItem("auto-suggest", autoSuggest);
};

const setAutoRemind = (autoRemind) => {
  localStorage.setItem("auto-remind", autoRemind);
};

const goToCreateCategory = () => {
  window.location.href = "/create-category";
};

const goToCategories = () => {
  clearLocalStorage([]);
  window.location.href = "/categories";
};

const goToReportABug = () => {
  clearLocalStorage([]);
  window.location.href = "/report-bug";
};

const goToRecurringTodos = () => {
  clearLocalStorage([]);
  window.location.href = "/recurring-todos";
};

const goToSuggestedTodos = () => {
  clearLocalStorage([]);
  window.location.href = "/suggested-todos";
};

const goToCompletedTodos = () => {
  clearLocalStorage([]);
  window.location.href = "/completed-todos";
};

const goToProfile = () => {
  clearLocalStorage([]);
  window.location.href = "/profile";
};

const goToProvisionDemoUser = () => {
  clearLocalStorage([]);
  window.location.href = "/provision";
};

const goToListOfUsers = () => {
  clearLocalStorage([]);
  window.location.href = "/users";
};

const goToImportTodos = () => {
  clearLocalStorage([]);
  window.location.href = "/import";
};

const goToLogin = () => {
  clearLocalStorage([]);
  // Rather than redirecting to the login page, we redirect to the desync page in order to clear localStorage on domain change.
  window.location.href =
    getHost() === "daps.localhost"
      ? "http://localhost/desync"
      : "https://deselflopment.com/desync";
};

const goToRegister = () => {
  clearLocalStorage([]);
  window.location.href =
    getHost() === "daps.localhost"
      ? "http://localhost/register"
      : "https://deselflopment.com/register";
};

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
};

const sortTodosByField = (field, ascending, f, setSpan) => {
  let todos = JSON.parse(localStorage.getItem("todos"));
  if (!todos) {
    return;
  }
  todos = sortArrayByField(todos, field, ascending);
  localStorage.setItem("todos", JSON.stringify(todos));
  f(todos);
  if (setSpan) {
    setSpan({
      textAlign: "center",
      display: "block",
    });
  }
};

const sortCategoriesByField = (data, field, ascending, f, setSpan) => {
  let categories = sortArrayByField(data, field, ascending);
  f(categories);
  if (setSpan) {
    setSpan({
      textAlign: "center",
      display: "block",
    });
  }
};

function getHost() {
  return window.location.host;
}

export default checkAccess;

export {
  hashPassword,
  setAutoSuggest,
  setAutoRemind,
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
  goToLogin,
  goToRegister,
  sortArrayByField,
  clearLocalStorage,
  getUserId,
  sortTodosByField,
  checkValidToken,
  sortCategoriesByField,
};
