import {createContext} from "react";

interface AuthContextType {
  user: any;
  login: (token: string, expireAt: Date, callback: VoidFunction) => void;
  logout: () => void;
}

const AuthContext = createContext<AuthContextType>(null!);

export default AuthContext;
