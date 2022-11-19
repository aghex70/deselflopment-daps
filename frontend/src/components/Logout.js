import React from 'react'
import {Button, ButtonGroup, Container, Modal, ModalBody} from "react-bootstrap";

const Logout = () => {
  const clearLocalStorage = () => {
    localStorage.clear();
  }

  clearLocalStorage();

  const login = (e) => {
    e.preventDefault();
    window.location.href = "/login";
  }

  return (
    <Container>
      <Modal className='unshareModal text-center' show={true}
             centered={true} size='lg'>
        <ModalBody>
          <h3 style={{margin: "32px"}}>Thanks for using DAPS!</h3>
          <ButtonGroup style={{width: "80%"}}>
            <Button
              variant="success"
              onClick={(e) => login(e)}
              style={{margin: "auto", display: "block", padding: "0", textAlign: "center"}}
            >Login</Button>
          </ButtonGroup>
        </ModalBody>
      </Modal>
    </Container>
  )
}

export default Logout;
