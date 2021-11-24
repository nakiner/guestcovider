export const UPDATE_USER = () => `/user`;
export const SEARCH_USERS = (surname) => `user/search?surname=${encodeURIComponent(surname)}`;
