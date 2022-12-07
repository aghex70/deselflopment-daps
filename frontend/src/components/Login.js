import React, {useState} from 'react'
import {Button, Container, FloatingLabel, Form} from "react-bootstrap";
import UserService from "../services/user";
import {hashPassword, skipLogin} from "../utils/helpers";


const Login = () => {
  skipLogin();
  document.title = 'deselflopment - daps'
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [login, setLogin] = useState(false);

  const handleSubmit = (e) => {
    e.preventDefault();
    const hashedPassword = hashPassword(password);
    UserService.login(email, hashedPassword).then(
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
        error = new Error("Login failed!");
      }
    )
  }
    if (login) {
      window.location.href = "/categories";
    }

  return (
    // Create a container with a class name "contenedor" and background color red
    <Container
        style={{
          display: 'flex',
          justifyContent:'center',
          alignItems:'center',
          height: '50vh',
    }}>
      <Form  onSubmit={(e)=>handleSubmit(e)}>
        <h1 style={{ margin: '0px 0px 32px' }} className="text-center">Login</h1>
        <FloatingLabel
          controlId="floatingEmail"
          label="Email address"
          value={email}
          onChange={(e) => setEmail(e.target.value)}
        >
          <Form.Control type="email" placeholder="Email" />
        </FloatingLabel>

        <FloatingLabel
          controlId="floatingPassword"
          label="Password"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
        >
          <Form.Control type="password" placeholder="Password" />
        </FloatingLabel>

        <Button
          variant="success"
          type="submit"
        >
          Login
        </Button>
      </Form>
    </Container>
  )
}

export default Login;
