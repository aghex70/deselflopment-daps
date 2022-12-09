import React, {useState} from 'react'
import {Button, ButtonGroup, Container, FloatingLabel, Form, Modal, ModalBody} from "react-bootstrap";
import UserService from "../services/user";
import {skipLogin} from "../utils/helpers";
import {
  CancelButtonText,
  EmailAddressLabelText, EnterEmailText, IncorrectPasswordText, InvalidEmailText,
  LoginButtonText, LoginHeaderText,
  PasswordLabelText, PasswordNotLongEnoughText,
  RegisterButtonText, ReturnButtonText,
  UserNotFoundText
} from "../utils/texts";


const Login = () => {
  skipLogin();
  document.title = 'deselflopment - daps'
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [showModalUserDoesNotExist, setShowModalUserDoesNotExist] = useState(false);
  const [showModalPasswordNotLongEnough, setShowModalPasswordNotLongEnough] = useState(false);
  const [showModalEmailNotFilled, setShowModalEmailNotFilled] = useState(false);
  const [showModalIncorrectPassword, setShowModalIncorrectPassword] = useState(false);
  const [showModalInvalidEmail, setShowModalInvalidEmail] = useState(false);

  const toggleModalUserDoesNotExist = () => {
    setShowModalUserDoesNotExist(!showModalUserDoesNotExist);
  }

  const toggleModalPasswordNotLongEnough = () => {
    setShowModalPasswordNotLongEnough(!showModalPasswordNotLongEnough);
  }

  const toggleModalEmailNotFilled = () => {
    setShowModalEmailNotFilled(!showModalEmailNotFilled);
  }

  const toggleModalIncorrectPassword = () => {
    setShowModalIncorrectPassword(!showModalIncorrectPassword);
  }

  const toggleModalInvalidEmail = () => {
    setShowModalInvalidEmail(!showModalInvalidEmail);
  }

  const handleSubmit = (e) => {
    e.preventDefault();
    if (password.length < 13) {
      setShowModalPasswordNotLongEnough(true);
      return;
    }

    if (email.length < 1) {
        setShowModalEmailNotFilled(true);
        return;
    }

    UserService.login(email, password).then(
      (response) => {
        if (response.status === 200) {
          localStorage.setItem("access_token", response.data.access_token);
          localStorage.setItem("user_id", response.data.user_id);
          localStorage.setItem("language", "es");
          window.location.href = "/categories";
        }
      }
    ).catch(
      (error) => {
        if (error.response.data.message === "record not found") {
          setShowModalUserDoesNotExist(true);
        } else if (error.response.data.message === "invalid credentials") {
          setShowModalIncorrectPassword(true);
        } else if (error.response.data.message.includes("Field validation for 'Email' failed on the 'email' tag")) {
          setShowModalInvalidEmail(true);
        }
      }
    )
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
      <Form onSubmit={(e)=>handleSubmit(e)}>
        <h1 style={{ margin: '0px 0px 32px' }} className="text-center">{LoginHeaderText}</h1>
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

        <Button
          variant="success"
          type="submit"
        >
          {LoginButtonText}
        </Button>
      </Form>
      <Modal className='successModal text-center' show={showModalUserDoesNotExist} open={showModalUserDoesNotExist} centered={true} size='lg'>
      <ModalBody>
        <h4 style={{margin: "32px"}}>{UserNotFoundText}</h4>
        <ButtonGroup style={{width: "100%", paddingLeft: "10%", paddingRight: "10%"}}>
          <Button
              variant="success"
              type="submit"
              onClick={(e) => window.location.href = "/register"}
              style={{margin: "auto", display: "block", padding: "0", textAlign: "center"}}
          >{RegisterButtonText}</Button>
          <Button
              variant="danger"
              onClick={(e) => toggleModalUserDoesNotExist(e)}
              style={{margin: "auto", display: "block", padding: "0", textAlign: "center"}}
          >{CancelButtonText}</Button>
        </ButtonGroup>
      </ModalBody>
    </Modal>

      <Modal className='successModal text-center' show={showModalEmailNotFilled} open={showModalEmailNotFilled} centered={true} size='lg'>
        <ModalBody>
          <h4 style={{margin: "32px"}}>{EnterEmailText}</h4>
          <ButtonGroup style={{width: "40%"}}>
            <Button
                variant="danger"
                onClick={(e) => toggleModalEmailNotFilled(e)}
                style={{margin: "auto", display: "block", padding: "0", textAlign: "center"}}
            >{ReturnButtonText}</Button>
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
            >{ReturnButtonText}</Button>
          </ButtonGroup>
        </ModalBody>
      </Modal>

      <Modal className='successModal text-center' show={showModalIncorrectPassword} open={showModalIncorrectPassword} centered={true} size='lg'>
        <ModalBody>
          <h4 style={{margin: "32px"}}>{IncorrectPasswordText}</h4>
          <ButtonGroup style={{width: "40%"}}>
            <Button
                variant="danger"
                onClick={(e) => toggleModalIncorrectPassword(e)}
                style={{margin: "auto", display: "block", padding: "0", textAlign: "center"}}
            >{ReturnButtonText}</Button>
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
            >{ReturnButtonText}</Button>
          </ButtonGroup>
        </ModalBody>
      </Modal>
    </Container>
  )
}

export default Login;
