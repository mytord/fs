import React, {useContext} from 'react';
import {useState} from "react";
import {publicApi} from "../services/api";
import AuthContext from "./authContext";
import {useNavigate} from "react-router-dom";
import {Alert, Button, Form} from "react-bootstrap";
import axios, {AxiosError} from "axios";
import {ErrorResponse} from "typescript-axios";

function Register() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [firstName, setFirstName] = useState("");
  const [lastName, setLastName] = useState("");
  const [age, setAge] = useState("");
  const [city, setCity] = useState("");
  const [interests, setInterests] = useState("");
  const [errors, setErrors] = useState<string[]>([]);
  const auth = useContext(AuthContext);
  const navigate = useNavigate();

  const handleRegister = async (e: any) => {
    e.preventDefault();
    try {
      setErrors([]);

      const response = await publicApi.createProfile({
        email,
        password,
        firstName,
        lastName,
        city,
        interests,
        age: Number(age),
      });

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
      }
    }
  };

  return (
    <div>
      <Form onSubmit={handleRegister}>
        <h2>Register new profile</h2>
        <Alert show={errors.length>0} variant="danger">
          {errors.map(error => error)}
        </Alert>
        <Form.Group className="mb-3">
          <Form.Control type="email" placeholder="Email" onChange={(e) => setEmail(e.target.value)}/>
        </Form.Group>
        <Form.Group className="mb-3">
          <Form.Control type="password" placeholder="Password" onChange={(e) => setPassword(e.target.value)}/>
        </Form.Group>
        <Form.Group className="mb-3">
          <Form.Control type="text" placeholder="First Name" onChange={(e) => setFirstName(e.target.value)}/>
        </Form.Group>
        <Form.Group className="mb-3">
          <Form.Control type="text" placeholder="Last Name" onChange={(e) => setLastName(e.target.value)}/>
        </Form.Group>
        <Form.Group className="mb-3">
          <Form.Control type="text" placeholder="City" onChange={(e) => setCity(e.target.value)}/>
        </Form.Group>
        <Form.Group className="mb-3">
          <Form.Control type="number" placeholder="Age" onChange={(e) => setAge(e.target.value)}/>
        </Form.Group>
        <Form.Group className="mb-3">
          <Form.Control type="text" placeholder="Interests" onChange={(e) => setInterests(e.target.value)}/>
        </Form.Group>
        <Button variant="primary" type="submit">Register</Button>
      </Form>
    </div>
  );
}

export default Register;
