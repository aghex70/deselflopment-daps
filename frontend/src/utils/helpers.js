const checkAccess = () => {
  const token = localStorage.getItem("access_token");
  if (!token) {
    window.location.href = "/login";
  }
}

export default checkAccess;