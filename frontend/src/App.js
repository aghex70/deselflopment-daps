import {
  Container,
  Button,
  ModalBody,
  ButtonGroup,
  Modal,
} from "react-bootstrap";
import "./App.css";
import React from "react";
import { goToLogin, goToRegister, skipLogin } from "./utils/helpers";
import {
  LoginButtonText,
  RegisterButtonText,
  WelcomeToDapsText,
} from "./utils/texts";

export default function App() {
  skipLogin();
  return (
    <Container>
      <Modal
        className="unshareModal text-center"
        show={true}
        centered={true}
        size="lg"
      >
        <ModalBody>
          <h3 style={{ margin: "32px" }}>{WelcomeToDapsText}</h3>
          <ButtonGroup style={{ width: "80%" }}>
            <Button
              variant="success"
              onClick={() => goToLogin()}
              style={{
                margin: "auto",
                display: "block",
                padding: "0",
                textAlign: "center",
              }}
            >
              {LoginButtonText}
            </Button>
            <Button
              variant="primary"
              type="submit"
              onClick={() => goToRegister()}
              style={{
                margin: "auto",
                display: "block",
                padding: "0",
                textAlign: "center",
              }}
            >
              {RegisterButtonText}
            </Button>
          </ButtonGroup>
        </ModalBody>
      </Modal>
    </Container>
  );
}
