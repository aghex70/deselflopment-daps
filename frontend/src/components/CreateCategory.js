import React, {useState} from 'react';
import CategoryService from "../services/category";
import {Button, ButtonGroup, Container, FloatingLabel, Form, Modal, ModalBody} from "react-bootstrap";
import DapsHeader from "./Header";
import checkAccess from "../utils/helpers";

const CreateCategory = () => {
  checkAccess();
  const [categoryName, setCategoryName] = useState("");
  const [categoryDescription, setCategoryDescription] = useState("");
  const [showModalCategoryAlreadyExists, setShowModalCategoryAlreadyExists] = useState(false);

  const navigateCategories = () => {
    window.location.href = "/categories";
  }

  const toggleModalCategoryAlreadyExists = () => {
    setShowModalCategoryAlreadyExists(!showModalCategoryAlreadyExists);
  }

  const handleSubmit = (e) => {
    e.preventDefault();
    const data = {
      name: categoryName,
      description: categoryDescription,
    }

    CategoryService.createCategory(data).then(
      (response) => {
        if (response.status === 201) {
          window.location.href = "/categories";
        }
      }
    ).catch(
      (error) => {
        if (error.response.data.message === "already existent category with that user and name") {
            setShowModalCategoryAlreadyExists(true);
        }
      }
    )
  }

  return (
    <Container>
      <DapsHeader />
      <h1 style={{ margin: '0px 0px 32px' }} className="text-center">Create category</h1>
        <Form onSubmit={(e) => handleSubmit(e)}>
          <FloatingLabel
            controlId="floatingName"
            label="Name"
            value={categoryName}
            onChange={(e) => setCategoryName(e.target.value)}
          >
            <Form.Control
              type="name"
              placeholder="Name"
              value={categoryName}
              onChange={(e) => setCategoryName(e.target.value)} />
          </FloatingLabel>


          <FloatingLabel controlId="floatingDescription" label="Description">
            <Form.Control
              as="textarea"
              placeholder="Description"
              style={{ height: '100px', margin: '0px 0px 32px' }}
              type="description"
              value={categoryDescription}
              onChange={(e) => setCategoryDescription(e.target.value)}/>
          </FloatingLabel>

        <ButtonGroup style={{width: "100%", paddingLeft: "10%", paddingRight: "10%"}}>
          <Button
            variant="success"
            type="submit"
            onClick={(e) => handleSubmit(e)}
            style={{margin: "auto", display: "block", padding: "0", textAlign: "center"}}
          >Create</Button>
          <Button
            variant="danger"
            onClick={() => navigateCategories()}
            style={{margin: "auto", display: "block", padding: "0", textAlign: "center"}}
          >Cancel</Button>
        </ButtonGroup>

      </Form>
      <Modal className='successModal text-center' show={showModalCategoryAlreadyExists} open={showModalCategoryAlreadyExists} centered={true} size='lg'>
        <ModalBody>
          <h4 style={{margin: "32px"}}>Category already exists! Please try with a different name</h4>
          <ButtonGroup style={{width: "40%"}}>
            <Button
                variant="danger"
                onClick={(e) => toggleModalCategoryAlreadyExists(e)}
                style={{margin: "auto", display: "block", padding: "0", textAlign: "center"}}
            >Return</Button>
          </ButtonGroup>
        </ModalBody>
      </Modal>
    </Container>
  )
}

export default CreateCategory;
