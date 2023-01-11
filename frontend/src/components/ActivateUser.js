import React, {useState} from 'react'
import {Button, ButtonGroup, Container, Modal, ModalBody} from "react-bootstrap";
import {
    InvalidActivationLinkText,
    RegisterButtonText,
} from "../utils/texts";
import {goToLogin, goToRegister} from "../utils/helpers";
import UserService from "../services/user";


const ActivateUser = () => {
  document.title = 'deselflopment - daps'
  const [showModalUserNotFound, setShowModalUserNotFound] = useState(false);
  const uuid = window.location.pathname.split("activate/")[1];

    if (!showModalUserNotFound) {
      UserService.activateUser(uuid).then(
          () => {
              goToLogin();
        }
      ).catch(
          (error) => {
              if (error.response.data.message === "record not found") {
                  setShowModalUserNotFound(true);
              } else if (error.response.data.message === "user is alredy active") {
                  goToLogin();
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

      <Modal className='activateUser text-center' show={showModalUserNotFound}
             centered={true} size='lg'>
          <ModalBody>
              <h3 style={{margin: "32px"}}>{InvalidActivationLinkText}</h3>
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

export default ActivateUser;
