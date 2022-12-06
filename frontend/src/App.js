import {Container, Button, ModalBody, ButtonGroup, Modal} from "react-bootstrap";
import "./App.css";
import React from "react";
import {skipLogin} from "./utils/helpers";

export default function App() {
  skipLogin();
  document.title = 'deselflopment - daps';
  // Check if access_token is in local storage. If it is, navigate to the categories page.
  // if (localStorage.getItem("access_token")) {
  //   window.location.href = "/categories";
  // }
  // setData(localStorage.getItem("access_token"));

  const login = (e) => {
    e.preventDefault();
    window.location.href = "/login";
  }
  const register = (e) => {
    e.preventDefault();
    window.location.href = "/register";
  }

  return (
    <Container>
      <Modal className='unshareModal text-center' show={true}
             centered={true} size='lg'>
        <ModalBody>
          <h3 style={{margin: "32px"}}>Welcome to DAPS =]</h3>
          <ButtonGroup style={{width: "80%"}}>
            <Button
              variant="warning"
              type="submit"
              onClick={(e) => register(e)}
              style={{margin: "auto", display: "block", padding: "0", textAlign: "center"}}
            >Register</Button>
            <Button
              variant="success"
              onClick={(e) => login(e)}
              style={{margin: "auto", display: "block", padding: "0", textAlign: "center"}}
            >Login</Button>
          </ButtonGroup>
        </ModalBody>
      </Modal>
    </Container>
  );
}
