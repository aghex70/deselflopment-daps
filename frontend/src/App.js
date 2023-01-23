import {Container, Button, ModalBody, ButtonGroup, Modal} from "react-bootstrap";
import "./App.css";
import React, {useEffect, useState} from "react";
import {goToLogin, goToRegister} from "./utils/helpers";
import {LoginButtonText, RegisterButtonText, WelcomeToDapsText} from "./utils/texts";
import OneSignal from 'react-onesignal';

export default function App() {
  const [isOneSignalInitialized, setIsOneSignalInitialized] = useState(false);

  useEffect(() => {
    if (!isOneSignalInitialized) {
      OneSignal.init({
        appId: "d70c53bb-aad1-461e-96a6-b6cec7b9a0d4"
      });
      setIsOneSignalInitialized(true);
    }
  }, [isOneSignalInitialized]);

  return (
    <Container>
      <Modal className='unshareModal text-center' show={true}
             centered={true} size='lg'>
        <ModalBody>
          <h3 style={{margin: "32px"}}>{WelcomeToDapsText}</h3>
          <ButtonGroup style={{width: "80%"}}>
            <Button
                variant="success"
                onClick={() => goToLogin()}
                style={{margin: "auto", display: "block", padding: "0", textAlign: "center"}}
            >{LoginButtonText}</Button>
            <Button
              variant="warning"
              type="submit"
              onClick={() => goToRegister()}
              style={{margin: "auto", display: "block", padding: "0", textAlign: "center"}}
            >{RegisterButtonText}</Button>
          </ButtonGroup>
        </ModalBody>
      </Modal>
    </Container>
  );
}
