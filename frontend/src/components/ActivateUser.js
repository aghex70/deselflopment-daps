import React, {useState} from 'react'
import {Button, ButtonGroup, Container, Modal, ModalBody} from "react-bootstrap";
import {
    ActivationCodeRefreshedText,
    InvalidActivationLinkText, LoginButtonText,
    RefreshCodeButtonText,
    RegisterButtonText,
} from "../utils/texts";
import {goToLogin, goToRegister} from "../utils/helpers";
import UserService from "../services/user";


const ActivateUser = () => {
  const [showModalUserNotFound, setShowModalUserNotFound] = useState(false);
  const [showModalRefreshedActivationCode, setShowModalRefreshedActivationCode] = useState(false);
  const uuid = window.location.pathname.split("activate/")[1];

    if (!showModalUserNotFound) {
      UserService.activateUser(uuid).then(
          (response) => {
              goToLogin();
        }
      ).catch(
          (error) => {
              console.log("error" + error);
              console.log("error.response: " + error.response);
              console.log("error.response.data: " + error.response.data);
              console.log("error.response.data.message: " + error.response.data.message);
              setShowModalUserNotFound(true);
          }
      )
    }

    const refreshActivationCode = () => {
        UserService.refreshActivationCode(uuid).then(
            (response) => {
                setShowModalRefreshedActivationCode(true);
            }
        ).catch(
            (error) => {
                console.log("error" + error);
                console.log("error.response: " + error.response);
                // console.log("error.response.data: " + error.response.data);
                // console.log("error.response.data.message: " + error.response.data.message);
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

      <Modal className='activateUser text-center' show={true}
             centered={true} size='lg'>
          <ModalBody>
              <h3 style={{margin: "32px"}}>{InvalidActivationLinkText}</h3>
              <ButtonGroup style={{width: "80%"}}>
                  <Button
                      variant="warning"
                      type="submit"
                      onClick={() => goToRegister()}
                      style={{margin: "auto", display: "block", padding: "0", textAlign: "center"}}
                  >{RegisterButtonText}</Button>
                  <Button
                      variant="success"
                      onClick={() => refreshActivationCode()}
                      style={{margin: "auto", display: "block", padding: "0", textAlign: "center"}}
                  >{RefreshCodeButtonText}</Button>
              </ButtonGroup>
          </ModalBody>
      </Modal>

      <Modal className='successModal text-center' show={showModalRefreshedActivationCode} open={showModalRefreshedActivationCode} centered={true} size='lg'>
          <ModalBody>
              <h4 style={{margin: "32px"}}>{ActivationCodeRefreshedText}</h4>
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

export default ActivateUser;
