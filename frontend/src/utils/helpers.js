import bcrypt from "bcryptjs";

const hashPassword = (string) => {
  return bcrypt.hashSync(string, 10);
};

const getUserData = () => {
  return JSON.parse(localStorage.getItem("deselflopment"))
}


const checkAccess = () => {
  if (!getUserToken() || !getUserId()) {
    goToLogin();
    // console.log("No user token or user id found.")
  }
};

const skipLogin = () => {
  // if (getUserToken() && getUserId()) {
  //   goToCategories();
  // }
};

const getUserToken = () => {
  // return getUserData()?.accessToken;
  let userData = getUserData();
  if (userData) {
      return userData.accessToken;
  }
};

let userAccessToken = getUserToken();

const checkValidToken = (error) => {
  if (error.response.data.message === "signature is invalid") {
    localStorage.removeItem("deselflopment");
    goToLogin();
  }
};

const getUserId = () => {
  // return getUserData()?.userId;
  let userData = getUserData();
  if (userData) {
    return userData.userId;
  }
};

const getIsAdmin = () => {
  // return getUserData()?.admin;
  let userData = getUserData();
  if (userData) {
    return userData.admin;
  }
};

const clearLocalStorage = (excludedKeys = []) => {
  let baseExcludedKeys = [
    "deselflopment",
  ];
  let untouchedKeys = baseExcludedKeys.concat(excludedKeys);
  for (let key in localStorage) {
    if (!untouchedKeys.includes(key)) {
      localStorage.removeItem(key);
    }
  }
};

// const retrieveFromLocalStorage = (key) => {
//   let deselflopmentData = localStorage.getItem('deselflopment');
//   let deselflopmentObject = JSON.parse(deselflopmentData);
//   return deselflopmentObject[key];
// }


const setLanguage = (language) => {
  console.log("Setting language to: ", language);
  let deselflopmentData = localStorage.getItem('deselflopment');

  // Parse the JSON data into a JavaScript object
  let deselflopmentObject = JSON.parse(deselflopmentData);

  deselflopmentObject.language = language;
  // Convert the updated object back to a JSON string
  let updatedDeselflopmentData = JSON.stringify(deselflopmentObject);

  // Store the updated data back into localStorage
  localStorage.setItem('deselflopment', updatedDeselflopmentData);
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
  window.location.href = "http://localhost:13001/login";
    // getHost() === "daps.localhost"
    //   ? "http://localhost/desync"
    //   : "https://deselflopment.com/desync";

};

const goToRegister = () => {
  clearLocalStorage([]);
  if (window.location.host.includes("localhost")) {
    window.location.href = "http://localhost:13001/register";
  }
  // window.location.href =
  //   getHost() === "daps.localhost"
  //     ? "http://localhost/register"
  //     : "https://deselflopment.com/register";
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

const retrieveFromLocalStorage = (key) => {
  let deselflopmentData = localStorage.getItem('deselflopment');
  if (!deselflopmentData) {
    return null;
  }
  let deselflopmentObject = JSON.parse(deselflopmentData);
  return deselflopmentObject[key];
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
  getIsAdmin,
  getUserData,
  userAccessToken,
  retrieveFromLocalStorage,
};
