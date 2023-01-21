import React, {useState} from 'react'
import {Button, ButtonGroup, Container, FloatingLabel, Form, Modal, ModalBody} from "react-bootstrap";
import UserService from "../services/user";
import {
  CancelButtonText,
  EmailAddressLabelText,
  EnterEmailText,
  ForgotPasswordHeaderText,
  PasswordLinkResetText,
  RegisterButtonText,
  ResetPasswordButtonText,
  UserNotFoundText,
} from "../utils/texts";


const ResetLink = () => {
  localStorage.removeItem("access_token");
  document.title = 'deselflopment - daps'
  const [email, setEmail] = useState("");
  const [showModalEmailNotFilled, setShowModalEmailNotFilled] = useState(false);
  const [showModalPasswordReset, setShowModalPasswordReset] = useState(false);
  const [showModalUserDoesNotExist, setShowModalUserDoesNotExist] = useState(false);

  const toggleModalEmailNotFilled = () => {
    setShowModalEmailNotFilled(!showModalEmailNotFilled);
  }

  const toggleModalPasswordReset = () => {
    setShowModalPasswordReset(!showModalPasswordReset);
  }

  const toggleModalUserDoesNotExist = () => {
    setShowModalUserDoesNotExist(!showModalUserDoesNotExist);
  }

  const handleSubmit = (e) => {
    e.preventDefault();

    if (email.length < 1) {
      setShowModalEmailNotFilled(true);
      return;
    }

    UserService.createResetLink(email).then(
        (response) => {
          setShowModalPasswordReset(true);
        }
    ).catch(
        (error) => {
          if (error.response.data.message === "record not found") {
            setShowModalUserDoesNotExist(true);
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
      <Form onSubmit={(e) => handleSubmit(e)}>
        <h1 style={{ margin: '0px 0px 32px' }} className="text-center">{ForgotPasswordHeaderText}</h1>
        <FloatingLabel
          controlId="floatingEmail"
          label={EmailAddressLabelText}
          value={email}
          onChange={(e) => setEmail(e.target.value)}
        >
          <Form.Control type="email" placeholder="Email" />
        </FloatingLabel>

        <Button
          variant="success"
          type="submit"
        >
          {ResetPasswordButtonText}
        </Button>
      </Form>

      <Modal className='successModal text-center' show={showModalEmailNotFilled} open={showModalEmailNotFilled} centered={true} size='lg'>
        <ModalBody>
          <h4 style={{margin: "32px"}}>{EnterEmailText}</h4>
          <ButtonGroup style={{width: "40%"}}>
            <Button
                variant="danger"
                onClick={() => toggleModalEmailNotFilled()}
                style={{margin: "auto", display: "block", padding: "0", textAlign: "center"}}
            >{CancelButtonText}</Button>
          </ButtonGroup>
        </ModalBody>
      </Modal>

      <Modal className='successModal text-center' show={showModalPasswordReset} open={showModalPasswordReset} centered={true} size='lg'>
        <ModalBody>
          <h4 style={{margin: "32px"}}>{PasswordLinkResetText}</h4>
          <ButtonGroup style={{width: "40%"}}>
            <Button
                variant="success"
                onClick={() => toggleModalPasswordReset()}
                style={{margin: "auto", display: "block", padding: "0", textAlign: "center"}}
            >{CancelButtonText}</Button>
          </ButtonGroup>
        </ModalBody>
      </Modal>

      <Modal className='successModal text-center' show={showModalUserDoesNotExist} open={showModalUserDoesNotExist} centered={true} size='lg'>
        <ModalBody>
          <h4 style={{margin: "32px"}}>{UserNotFoundText}</h4>
          <ButtonGroup style={{width: "100%", paddingLeft: "10%", paddingRight: "10%"}}>
            <Button
                variant="danger"
                onClick={() => toggleModalUserDoesNotExist()}
                style={{margin: "auto", display: "block", padding: "0", textAlign: "center"}}
            >{CancelButtonText}</Button>
            <Button
                variant="success"
                type="submit"
                onClick={(e) => window.location.href = "/register"}
                style={{margin: "auto", display: "block", padding: "0", textAlign: "center"}}
            >{RegisterButtonText}</Button>
          </ButtonGroup>
        </ModalBody>
      </Modal>

    </Container>
  )
}


export default ResetLink;
