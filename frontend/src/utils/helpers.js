const checkAccess = () => {
  const token = getUserToken();
  if (!token) {
    window.location.href = "/login";
  }
}

const skipLogin = () => {
  if (getUserToken()) {
    window.location.href = "/categories";
  }
}
const getUserToken = () => {
  return localStorage.getItem("access_token");
}

export default checkAccess;
export { skipLogin };