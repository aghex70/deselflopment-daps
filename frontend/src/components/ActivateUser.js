import React from 'react'
import {Button, ButtonGroup, Container, Modal, ModalBody} from "react-bootstrap";
import {
    RegisterButtonText,
    UserNotFoundText,
} from "../utils/texts";
import {goToLogin, goToRegister} from "../utils/helpers";
import {useParams} from "react-router-dom";
import UserService from "../services/user";


const ActivateUser = () => {
  const { id } = useParams();
  const [showModalUserNotFound, setShowModalUserNotFound] = React.useState(false);

  UserService.activateUser(id).then(
      (response) => {
      if (response.status === 200) {
          goToLogin();
      }
    }
  ).catch(
      (error) => {
          console.log("response: " + error.response);
          console.log("response.data: " + error.response.data);
          console.log("response.data.message: " + error.response.data.message);
          setShowModalUserNotFound(true);
      }
  )

  return (
      <Container
          style={{
            display: 'flex',
            justifyContent:'center',
            alignItems:'center',
            height: '50vh',
          }}>

      <Modal className='successModal text-center' show={showModalUserNotFound} open={showModalUserNotFound} centered={true} size='lg'>
          <ModalBody>
              <h4 style={{margin: "32px"}}>{UserNotFoundText}</h4>
              <ButtonGroup style={{width: "40%"}}>
                  <Button
                      variant="danger"
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
