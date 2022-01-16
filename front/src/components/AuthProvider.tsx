import React, {ReactNode, useState} from "react";
import authService from "../services/auth";
import AuthContext from "./authContext";

function AuthProvider({ children }: { children: ReactNode }) {
  let [user, setUser] = useState<string | undefined>(authService.getUser());

  let login = (token: string, expireAt: Date, callback: VoidFunction) => {
    authService.login(token, expireAt);
    setUser(token);
    callback();
  };

  let logout = () => {
    authService.logout();
    setUser(undefined);
  };

  let value = { user, login, logout };

  return <AuthContext.Provider value={value}>{children}</AuthContext.Provider>;
}

export default AuthProvider;
