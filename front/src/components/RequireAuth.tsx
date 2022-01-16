import React from "react";
import {useContext} from "react";
import AuthContext from "./authContext";
import {Navigate} from "react-router-dom";

function RequireAuth({ children }: { children: JSX.Element }) {
  let auth = useContext(AuthContext);

  if (!auth.user) {
    return <Navigate to="/login" />;
  }

  return children;
}

export default RequireAuth;
