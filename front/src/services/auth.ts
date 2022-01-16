const login = (token: string, expireAt: Date): void => {
  localStorage.setItem("currentUser", JSON.stringify({token, expireAt}));
}

const logout = () => {
  localStorage.removeItem("currentUser");
}

const getUser = (): string | undefined => {
  const item = localStorage.getItem("currentUser");

  if (!item) {
    return undefined;
  }

  const currentUser = JSON.parse(item);

  if (new Date(currentUser.expireAt) <= new Date()) {
    logout();
    return undefined;
  }

  return currentUser.token;
}

const authService = {
  login,
  logout,
  getUser,
}

export default authService;
