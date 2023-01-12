import React, {useState} from 'react'
import {Button, ButtonGroup, Container, FloatingLabel, Form, Modal, ModalBody} from "react-bootstrap";
import {
    CancelButtonText,
    InvalidResetPasswordCodeText,
    PasswordLabelText,
    PasswordsDoNotMatchText,
    RegisterButtonText,
    RepeatPasswordLabelText,
    ResetPasswordHeaderText,
} from "../utils/texts";
import {goToLogin, goToRegister, hashPassword} from "../utils/helpers";
import UserService from "../services/user";


const ResetPassword = () => {
  localStorage.removeItem("access_token");
  document.title = 'deselflopment - daps'
  const [password, setPassword] = useState("");
  const [repeatPassword, setRepeatPassword] = useState("");
  const [showModalPasswordsDoNotMatch, setShowModalPasswordsDoNotMatch] = useState(false);
  const [showModalUserNotFound, setShowModalUserNotFound] = useState(false);
  const uuid = window.location.pathname.split("reset-password/")[1];

  const toggleModalPasswordsDoNotMatch = () => {
    setShowModalPasswordsDoNotMatch(!showModalPasswordsDoNotMatch);
  }

  const handleSubmit = (e) => {
    e.preventDefault();
    if (password !== repeatPassword) {
      setShowModalPasswordsDoNotMatch(true);
      return;
    }

    const hashedPassword = hashPassword(password);
    UserService.resetPassword(uuid, hashedPassword).then(
        () => {
          goToLogin();
        }
    ).catch(
        (error) => {
            if (error.response.data.message === "record not found") {
                setShowModalUserNotFound(true);
            }
        }
    )
  }

  return (
      <Container
          style={{
            display: 'flex',
            justifyContent:'center',
            alignItems:'center',
            height: '50vh',
          }}>

        <Form  onSubmit={(e) => handleSubmit(e)}>
          <h1 style={{ margin: '0px 0px 32px' }} className="text-center">{ResetPasswordHeaderText}</h1>
          <FloatingLabel
              controlId="floatingPassword"
              label={PasswordLabelText}
              value={password}
              onChange={(e) => setPassword(e.target.value)}
          >
            <Form.Control type="password" placeholder="Password" />
          </FloatingLabel>

          <FloatingLabel
              controlId="floatingRepeatPassword"
              label={RepeatPasswordLabelText}
              value={repeatPassword}
              onChange={(e) => setRepeatPassword(e.target.value)}
          >
            <Form.Control type="password" placeholder="RepeatPassword" />
          </FloatingLabel>
          <Button
              variant="success"
              type="submit"
          >
            {ResetPasswordHeaderText}
          </Button>

        </Form>

        <Modal className='successModal text-center' show={showModalPasswordsDoNotMatch} open={showModalPasswordsDoNotMatch} centered={true} size='lg'>
          <ModalBody>
            <h4 style={{margin: "32px"}}>{PasswordsDoNotMatchText}</h4>
            <ButtonGroup style={{width: "40%"}}>
              <Button
                  variant="danger"
                  onClick={(e) => toggleModalPasswordsDoNotMatch(e)}
                  style={{margin: "auto", display: "block", padding: "0", textAlign: "center"}}
              >{CancelButtonText}</Button>
            </ButtonGroup>
          </ModalBody>
        </Modal>
        <Modal className='activateUser text-center' show={showModalUserNotFound}
               centered={true} size='lg'>
          <ModalBody>
            <h3 style={{margin: "32px"}}>{InvalidResetPasswordCodeText}</h3>
            <ButtonGroup style={{width: "40%"}}>
              <Button
                  variant="success"
                  onClick={() => goToRegister()}
                  style={{margin: "auto", display: "block", padding: "0", textAlign: "center"}}
              >{RegisterButtonText}</Button>
            </ButtonGroup>
          </ModalBody>
        </Modal>
      </Container>
  )
}

export default ResetPassword;
