import {Container, Button, ModalBody, ButtonGroup, Modal} from "react-bootstrap";
import "./App.css";
import React from "react";
import {skipLogin} from "./utils/helpers";
import {LoginButtonText, RegisterButtonText, WelcomeToDapsText} from "./utils/texts";

export default function App() {
  skipLogin();
  document.title = 'deselflopment - daps';

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
          <h3 style={{margin: "32px"}}>{WelcomeToDapsText}</h3>
          <ButtonGroup style={{width: "80%"}}>
            <Button
              variant="warning"
              type="submit"
              onClick={(e) => register(e)}
              style={{margin: "auto", display: "block", padding: "0", textAlign: "center"}}
            >{RegisterButtonText}</Button>
            <Button
              variant="success"
              onClick={(e) => login(e)}
              style={{margin: "auto", display: "block", padding: "0", textAlign: "center"}}
            >{LoginButtonText}</Button>
          </ButtonGroup>
        </ModalBody>
      </Modal>
    </Container>
  );
}
