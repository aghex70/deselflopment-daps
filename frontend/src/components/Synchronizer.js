import { useEffect } from 'react';

const QueryParamsToLocalStorage = () => {
  useEffect(() => {
    // Get query params from URL
    const queryParams = new URLSearchParams(window.location.search);

    // Store every query param in localStorage
    queryParams.forEach((value, key) => {
      localStorage.setItem(key, value);
    });

    window.location.href = "/categories";
  }, []);

  return null; // This component doesn't render anything
};

export default QueryParamsToLocalStorage;
