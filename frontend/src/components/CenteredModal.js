import React, {useEffect, useState} from 'react';
import {Button, Container, Form, Modal} from "react-bootstrap";
import { useParams } from 'react-router-dom'
import CategoryService from "../services/category";

const MyVerticallyCenteredModal = () =>{
  return (
    <Modal
      {...props}
      size="lg"
      aria-labelledby="contained-modal-title-vcenter"
      centered
    >
      <Modal.Header closeButton>
        <Modal.Title id="contained-modal-title-vcenter">
          Modal heading
        </Modal.Title>
      </Modal.Header>
      <Modal.Body>
        <h4>Centered Modal</h4>
        <p>
          Cras mattis consectetur purus sit amet fermentum. Cras justo odio,
          dapibus ac facilisis in, egestas eget quam. Morbi leo risus, porta ac
          consectetur ac, vestibulum at eros.
        </p>
      </Modal.Body>
      <Modal.Footer>
        {/*Hide element when clicking button*/}
        <Button onClick={props.onHide}>Close</Button>
        <Button onClick={(e) => e}>Close</Button>
      </Modal.Footer>
    </Modal>
  );
}

export default MyVerticallyCenteredModal;