import React from "react";
import {
  Button,
  ButtonGroup,
  Container,
  Modal,
  ModalBody,
} from "react-bootstrap";
import { LoginButtonText, ThankYouHeaderText } from "../utils/texts";
import { goToLogin } from "../utils/helpers";

const Logout = () => {
  localStorage.removeItem("access_token");
  localStorage.removeItem("user_id");

  const login = (e) => {
    e.preventDefault();
    goToLogin();
  };

  return (
    <Container>
      <Modal
        className="unshareModal text-center"
        show={true}
        centered={true}
        size="lg"
      >
        <ModalBody>
          <h3 style={{ margin: "32px" }}>{ThankYouHeaderText}</h3>
          <ButtonGroup style={{ width: "80%" }}>
            <Button
              variant="success"
              onClick={(e) => login(e)}
              style={{
                margin: "auto",
                display: "block",
                padding: "0",
                textAlign: "center",
              }}
            >
              {LoginButtonText}
            </Button>
          </ButtonGroup>
        </ModalBody>
      </Modal>
    </Container>
  );
};

export default Logout;
