import React, {useState} from 'react'
import {Button, Container, FloatingLabel, Form} from "react-bootstrap";
import UserService from "../services/user";

export default function Register() {
  const [name, setName] = useState("");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [repeatPassword, setRepeatPassword] = useState("");
  const [register, setRegister] = useState(false);

  const styles = {
    display: 'flex',
    justifyContent:'center',
    alignItems:'center',
    height: '70vh'
  }

  const handleSubmit = (e) => {
    e.preventDefault();

    UserService.register(name, email, password, repeatPassword).then(
      (response) => {
        if (response.status === 201) {
          setRegister(true);
          window.location.href = "/login";
        } else {
          setRegister(false);
        }
      }
    ).catch(
      (error) => {
        console.log(error);
        window.location.reload();
      }
    )
  }


  return (
    <Container style={styles}>
      <Form  onSubmit={(e)=>handleSubmit(e)}>
        <h1 style={{ margin: '0px 0px 32px' }} className="text-center">Register</h1>
        <FloatingLabel
          controlId="floatingName"
          label="Name"
          value={name}
          onChange={(e) => setName(e.target.value)}
        >
          <Form.Control type="name" placeholder="Name" />
        </FloatingLabel>

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

        <FloatingLabel
          controlId="floatingRepeatPassword"
          label="RepeatPassword"
          value={repeatPassword}
          onChange={(e) => setRepeatPassword(e.target.value)}
        >
          <Form.Control type="password" placeholder="RepeatPassword" />
        </FloatingLabel>


        <Button
          variant="success"
          type="submit"
          onClick={(e) => handleSubmit(e)}
        >
          Register
        </Button>

      </Form>
    </Container>
  )
}
