import React, {useState} from 'react'
import {Button, ButtonGroup, Container, FloatingLabel, Form, Modal, ModalBody} from "react-bootstrap";
import UserService from "../services/user";
import {goToLogin, hashPassword, skipLogin} from "../utils/helpers";
import {
  ActivateUserText,
  CancelButtonText,
  EmailAddressLabelText,
  InvalidEmailText,
  LoginButtonText,
  NameLabelText,
  PasswordLabelText,
  PasswordNotLongEnoughText,
  PasswordsDoNotMatchText,
  RegisterButtonText,
  RegisterHeaderText,
  RepeatPasswordLabelText,
  UserAlreadyRegisteredText
} from "../utils/texts";

const Register = ()  =>{
  skipLogin();
  document.title = 'deselflopment - daps'
  const [name, setName] = useState("");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [repeatPassword, setRepeatPassword] = useState("");
  const [showModalPasswordsDoNotMatch, setShowModalPasswordsDoNotMatch] = useState(false);
  const [showModalUserAlreadyExists, setShowModalUserAlreadyExists] = useState(false);
  const [showModalPasswordNotLongEnough, setShowModalPasswordNotLongEnough] = useState(false);
  const [showModalInvalidEmail, setShowModalInvalidEmail] = useState(false);
  const [showModalActivateUser, setShowModalActivateUser] = useState(false);


  const toggleModalPasswordsDoNotMatch = () => {
    setShowModalPasswordsDoNotMatch(!showModalPasswordsDoNotMatch);
  }

  const toggleModalUserAlreadyExists = () => {
    setShowModalUserAlreadyExists(!showModalUserAlreadyExists);
  }

  const toggleModalPasswordNotLongEnough = () => {
    setShowModalPasswordNotLongEnough(!showModalPasswordNotLongEnough);
  }

  const toggleModalInvalidEmail = () => {
    setShowModalInvalidEmail(!showModalInvalidEmail);
  }

  const styles = {
    display: 'flex',
    justifyContent:'center',
    alignItems:'center',
    height: '70vh'
  }

  const handleSubmit = (e) => {
    e.preventDefault();
    if (password !== repeatPassword) {
      setShowModalPasswordsDoNotMatch(true);
      return;
    }

    if (password.length < 13) {
      setShowModalPasswordNotLongEnough(true);
      return;
    }

    const hashedPassword = hashPassword(password);
    UserService.register(name, email, hashedPassword).then(
      (response) => {
        if (response.status === 201) {
          localStorage.setItem("language", "en");
          setShowModalActivateUser(true);
        }
      }
    ).catch(
      (error) => {
        if (error.response.data.message === "user already registered") {
          setShowModalUserAlreadyExists(true);
        }
      }
    )
  }

  return (
    <Container style={styles}>
      <Form  onSubmit={(e) => handleSubmit(e)}>
        <h1 style={{ margin: '0px 0px 32px' }} className="text-center">{RegisterHeaderText}</h1>
        <FloatingLabel
          controlId="floatingName"
          label={NameLabelText}
          value={name}
          onChange={(e) => setName(e.target.value)}
        >
          <Form.Control type="name" placeholder="Name" />
        </FloatingLabel>

        <FloatingLabel
          controlId="floatingEmail"
          label={EmailAddressLabelText}
          value={email}
          onChange={(e) => setEmail(e.target.value)}
        >
          <Form.Control type="email" placeholder="Email" />
        </FloatingLabel>

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
          {RegisterButtonText}
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

      <Modal className='successModal text-center' show={showModalUserAlreadyExists} open={showModalUserAlreadyExists} centered={true} size='lg'>
        <ModalBody>
          <h4 style={{margin: "32px"}}>{UserAlreadyRegisteredText}</h4>
          <ButtonGroup style={{width: "100%", paddingLeft: "10%", paddingRight: "10%"}}>
            <Button
                variant="danger"
                onClick={(e) => toggleModalUserAlreadyExists(e)}
                style={{margin: "auto", display: "block", padding: "0", textAlign: "center"}}
            >{CancelButtonText}</Button>
            <Button
                variant="success"
                type="submit"
                onClick={(e) => window.location.href = "/login"}
                style={{margin: "auto", display: "block", padding: "0", textAlign: "center"}}
            >{LoginButtonText}</Button>
          </ButtonGroup>
        </ModalBody>
      </Modal>

      <Modal className='successModal text-center' show={showModalPasswordNotLongEnough} open={showModalPasswordNotLongEnough} centered={true} size='lg'>
        <ModalBody>
          <h4 style={{margin: "32px"}}>{PasswordNotLongEnoughText}</h4>
          <ButtonGroup style={{width: "40%"}}>
            <Button
                variant="danger"
                onClick={(e) => toggleModalPasswordNotLongEnough(e)}
                style={{margin: "auto", display: "block", padding: "0", textAlign: "center"}}
            >{CancelButtonText}</Button>
          </ButtonGroup>
        </ModalBody>
      </Modal>

      <Modal className='successModal text-center' show={showModalInvalidEmail} open={showModalInvalidEmail} centered={true} size='lg'>
        <ModalBody>
          <h4 style={{margin: "32px"}}>{InvalidEmailText}</h4>
          <ButtonGroup style={{width: "40%"}}>
            <Button
                variant="danger"
                onClick={(e) => toggleModalInvalidEmail(e)}
                style={{margin: "auto", display: "block", padding: "0", textAlign: "center"}}
            >{CancelButtonText}</Button>
          </ButtonGroup>
        </ModalBody>
      </Modal>

      <Modal className='activateUser text-center' show={showModalActivateUser}
             centered={true} size='lg'>
        <ModalBody>
          <h3 style={{margin: "32px"}}>{ActivateUserText}</h3>
          <ButtonGroup style={{width: "40%"}}>
            <Button
                variant="success"
                onClick={() => goToLogin()}
                style={{margin: "auto", display: "block", padding: "0", textAlign: "center"}}
            >{LoginButtonText}</Button>
          </ButtonGroup>
        </ModalBody>
      </Modal>
    </Container>
  )
}

export default Register;
