import React, {useState} from 'react'
import {Button, ButtonGroup, Container, FloatingLabel, Form, Modal, ModalBody} from "react-bootstrap";
import UserService from "../services/user";
import {hashPassword, skipLogin} from "../utils/helpers";

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


  const toggleModalPasswordsDoNotMatch = () => {
    setShowModalPasswordsDoNotMatch(!showModalPasswordsDoNotMatch);
  }

  const toggleModalUserAlreadyExists = () => {
    setShowModalUserAlreadyExists(!showModalUserAlreadyExists);
  }

  const toggleModalPasswordNotLongEnough = () => {
    setShowModalPasswordNotLongEnough(!showModalPasswordNotLongEnough);
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
          window.location.href = "/login";
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

      <Modal className='successModal text-center' show={showModalPasswordsDoNotMatch} open={showModalPasswordsDoNotMatch} centered={true} size='lg'>
        <ModalBody>
          <h4 style={{margin: "32px"}}>Passwords do not match! Please try again</h4>
          <ButtonGroup style={{width: "40%"}}>
            <Button
                variant="danger"
                onClick={(e) => toggleModalPasswordsDoNotMatch(e)}
                style={{margin: "auto", display: "block", padding: "0", textAlign: "center"}}
            >Return</Button>
          </ButtonGroup>
        </ModalBody>
      </Modal>

      <Modal className='successModal text-center' show={showModalUserAlreadyExists} open={showModalUserAlreadyExists} centered={true} size='lg'>
        <ModalBody>
          <h4 style={{margin: "32px"}}>User already registered! Please try with a different email</h4>
          <ButtonGroup style={{width: "40%"}}>
            <Button
                variant="danger"
                onClick={(e) => toggleModalUserAlreadyExists(e)}
                style={{margin: "auto", display: "block", padding: "0", textAlign: "center"}}
            >Return</Button>
          </ButtonGroup>
        </ModalBody>
      </Modal>

      <Modal className='successModal text-center' show={showModalPasswordNotLongEnough} open={showModalPasswordNotLongEnough} centered={true} size='lg'>
        <ModalBody>
          <h4 style={{margin: "32px"}}>Password must have more than 12 characters!</h4>
          <ButtonGroup style={{width: "40%"}}>
            <Button
                variant="danger"
                onClick={(e) => toggleModalPasswordNotLongEnough(e)}
                style={{margin: "auto", display: "block", padding: "0", textAlign: "center"}}
            >Return</Button>
          </ButtonGroup>
        </ModalBody>
      </Modal>
    </Container>
  )
}

export default Register;
