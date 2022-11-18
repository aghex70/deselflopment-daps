import {Container, Col, Row, Button, ModalBody, ButtonGroup, Modal} from "react-bootstrap";
import "./App.css";
import React from "react";

export default function App() {
  // Check if access_token is in local storage. If it is, navigate to the categories page.
  // if (localStorage.getItem("access_token")) {
  //   window.location.href = "/categories";
  // }
  // setData(localStorage.getItem("access_token"));

  const login = (e) => {
    e.preventDefault();
    window.location.href = "/login";
    // };
  }
  const register = (e) => {
    e.preventDefault();
    window.location.href = "/register";
    // };
  }

  return (
    <Container>
      <Modal className='unshareModal text-center' show={true}
             centered={true} size='lg'>
        <ModalBody>
          <div>
            Welcome to the DAPS application :]
            <div className='container my-4'>
            </div>
          </div>
          <ButtonGroup style={{width: "80%"}}>
            <Button
              variant="success"
              type="submit"
              onClick={(e) => register(e)}
              style={{margin: "auto", display: "block", padding: "0", textAlign: "center"}}
            >Register</Button>
            <Button
              variant="primary"
              onClick={(e) => login(e)}
              style={{margin: "auto", display: "block", padding: "0", textAlign: "center"}}
            >Login</Button>
          </ButtonGroup>
        </ModalBody>
      </Modal>
      {/*<Button*/}
      {/*  variant="primary"*/}
      {/*  type="submit"*/}
      {/*  onClick={(e) => login(e)}*/}
      {/*>*/}
      {/*  Login*/}
      {/*</Button>*/}
      {/*<Button*/}
      {/*  variant="primary"*/}
      {/*  type="submit"*/}
      {/*  onClick={(e) => register(e)}*/}
      {/*>*/}
      {/*  Register*/}
      {/*</Button>*/}
    </Container>
  );
}
