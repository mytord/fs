import React, {useContext} from 'react';
import {useState} from "react";
import {publicApi} from "../services/api";
import AuthContext from "./authContext";
import {useNavigate} from "react-router-dom";
import {Alert, Button, Form} from "react-bootstrap";
import axios, {AxiosError} from "axios";
import {ErrorResponse} from "typescript-axios";

function Login() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [errors, setErrors] = useState<string[]>([]);
  const auth = useContext(AuthContext);
  const navigate = useNavigate();

  const handleLogin = async (e: any) => {
    e.preventDefault();

    try {
      const response = await publicApi.login({email, password});

      console.log(response);

      const token = response.headers["x-set-token"];
      const expireAt = new Date(response.headers["x-token-expires"]);

      if (token && expireAt) {
        auth.login(token, expireAt, () => {
          navigate("/");
        });
      }
    } catch (e) {
      if (axios.isAxiosError(e)) {
        const errorResponse = (e as AxiosError<ErrorResponse, ErrorResponse>).response!;
        if (errorResponse.status === 400) {
          setErrors(errorResponse.data.errors!.map(err => err.message!));
        }

        if (errorResponse.status === 401) {
          setErrors(["invalid credentials"])
        }
      }
    }
  };

  return (
    <Form onSubmit={handleLogin}>
      <h2>Login</h2>
      <Alert show={errors.length>0} variant="danger">
        {errors.map(error => error)}
      </Alert>
      <Form.Group className="mb-3" controlId="formGroupEmail">
        <Form.Control type="email" placeholder="Email" onChange={(e) => setEmail(e.target.value)} />
      </Form.Group>
      <Form.Group className="mb-3" controlId="formGroupPassword">
        <Form.Control type="password" placeholder="Password" onChange={(e) => setPassword(e.target.value)} />
      </Form.Group>
      <Button variant="primary" type="submit">Log in</Button>
    </Form>
  );
}

export default Login;
