import React, {useContext} from 'react';
import './App.css';
import {Routes, Route} from "react-router-dom";
import Login from "./components/Login";
import Profile from "./components/Profile";
import AuthProvider from "./components/AuthProvider";
import RequireAuth from "./components/RequireAuth";
import Register from "./components/Register";
import {Col, Container, Nav, Row} from "react-bootstrap";
import NavigationBar from "./components/NavigationBar";
import People from "./components/People";

function App() {
  return (
    <AuthProvider>
      <Container>
        <Row>
          <Col lg={2} className="pt-3">
            <NavigationBar/>
          </Col>
          <Col className="pt-3">
            <Routes>
              <Route path="/login" element={<Login/>}/>
              <Route path="/register" element={<Register/>}/>
              <Route path="/" element={
                <RequireAuth>
                  <Profile/>
                </RequireAuth>
              }/>
              <Route path="/profile/:id" element={
                <RequireAuth>
                  <Profile/>
                </RequireAuth>
              }/>
              <Route path="/people" element={
                <RequireAuth>
                  <People/>
                </RequireAuth>
              }/>
            </Routes>
          </Col>
        </Row>
      </Container>
    </AuthProvider>
  );
}

export default App;
