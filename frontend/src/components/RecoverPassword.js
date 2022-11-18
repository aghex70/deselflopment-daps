import React, {useState} from 'react'
import {Button, Container, Form} from "react-bootstrap";
import UserService from "../services/user";

const Register = () =>{
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [login, setLogin] = useState(false);

  const handleSubmit = (e) => {
    e.preventDefault();

    UserService.login(email, password).then(
      (response) => {
        if (response.status === 200) {
          localStorage.setItem("access_token", response.data.access_token);
          localStorage.setItem("user_id", response.data.user_id);
          setLogin(true);
        } else {
          setLogin(false);
        }
      }
    ).catch(
      (error) => {
        console.log(error);
        error = new Error("Login failed!");
      }
    )
  }
    if (login) {
      window.location.href = "/categories";
    }

  return (
    <Container>
      <h2>Login</h2>
      <Form  onSubmit={(e)=>handleSubmit(e)}>
        {/* email */}
        <Form.Group controlId="formBasicEmail">
          <Form.Label>Email address</Form.Label>
          <Form.Control
            type="email"
            name="email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            placeholder="Enter email" />
        </Form.Group>

        {/* password */}
        <Form.Group controlId="formBasicPassword">
          <Form.Label>Password</Form.Label>
          <Form.Control
            type="password"
            name="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            placeholder="Password" />
        </Form.Group>

        {/* submit button */}
        <Button
          variant="primary"
          type="submit"
          onClick={(e) => handleSubmit(e)}
        >
          Submit
        </Button>

        {/* display success message */}
          <p className="text-danger">Please login</p>
      </Form>
    </Container>
  )
}

export default Register;
