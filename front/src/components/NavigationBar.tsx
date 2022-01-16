import React, {useContext} from 'react';
import {Nav} from "react-bootstrap";
import AuthContext from "./authContext";

function NavigationBar() {
  const auth = useContext(AuthContext);

  return (
    <Nav defaultActiveKey="/login" className="flex-column">
      {auth.user && <>
        <Nav.Link href="/">Your Profile</Nav.Link>
        <Nav.Link href="/people">People</Nav.Link>
        <Nav.Link onClick={auth.logout}>Logout</Nav.Link>
      </>}
      {!auth.user && <>
        <Nav.Link href="/login">Login</Nav.Link>
        <Nav.Link href="/register">Register</Nav.Link>
      </>}
    </Nav>
  );
}

export default NavigationBar;
