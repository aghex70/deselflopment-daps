import React, {useState} from 'react';
import CategoryService from "../services/category";
import {Button, ButtonGroup, Container, FloatingLabel, Form, Modal, ModalBody} from "react-bootstrap";
import DapsHeader from "./Header";
import checkAccess, {checkValidToken, goToCategories} from "../utils/helpers";
import {
  CancelButtonText,
  CategoryAlreadyExistsText,
  CreateButtonText,
  CreateCategoryHeaderText,
  DescriptionLabelText,
  NameLabelText,
  NoText,
  NotifiableLabelText,
  PleaseEnterCategoryNameText,
  YesText,
} from "../utils/texts";
import toBoolean from "validator/es/lib/toBoolean";

const CreateCategory = () => {
  checkAccess();
  const [categoryName, setCategoryName] = useState("");
  const [categoryDescription, setCategoryDescription] = useState("");
  const [categoryNotifiable, setCategoryNotifiable] = useState("false");
  const [disableNotifiable, setDisableNotifiable] = useState(false);
  const [showModalCategoryAlreadyExists, setShowModalCategoryAlreadyExists] = useState(false);
  const [showEnterCategoryModal, setShowEnterCategoryModal] = useState(false);

  const disableNotifiableSelect = () => {
    if (!disableNotifiable) {
      if (categoryNotifiable === "true" || categoryNotifiable === "false") {
        setDisableNotifiable(true);
      }
    }
  }
  disableNotifiableSelect();

  const toggleModalCategoryAlreadyExists = () => {
    setShowModalCategoryAlreadyExists(!showModalCategoryAlreadyExists);
  }

  const toggleEnterCategoryModal = () => {
    setShowEnterCategoryModal(!showEnterCategoryModal);
  }

  const handleSubmit = (e) => {
    e.preventDefault();

    if (categoryName === "") {
      toggleEnterCategoryModal();
      return;
    }

    const data = {
      name: categoryName,
      description: categoryDescription,
      notifiable: typeof(categoryNotifiable) == "boolean" ? categoryNotifiable : toBoolean(categoryNotifiable),
    }

    CategoryService.createCategory(data).then(
      (response) => {
        if (response.status === 201) {
          window.location.href = "/categories";
        }
      }
    ).catch(
      (error) => {
        checkValidToken(error)
        if (error.response.data.message === "already existent category with that user and name") {
          setShowModalCategoryAlreadyExists(true);
        }
      }
    )
  }

  return (
    <Container>
      <DapsHeader />
      <h1 style={{ margin: '0px 0px 32px' }} className="text-center">{CreateCategoryHeaderText}</h1>
        <Form onSubmit={(e) => handleSubmit(e)}>
          <FloatingLabel
            controlId="floatingName"
            label={NameLabelText}
            value={categoryName}
            onChange={(e) => setCategoryName(e.target.value)}
          >
            <Form.Control
              type="name"
              placeholder="Name"
              value={categoryName}
              onChange={(e) => setCategoryName(e.target.value)} />
          </FloatingLabel>

          <FloatingLabel controlId="floatingDescription" label={DescriptionLabelText}>
            <Form.Control
              as="textarea"
              placeholder={DescriptionLabelText}
              style={{ height: '100px', margin: '0px 0px 32px' }}
              type="description"
              value={categoryDescription}
              onChange={(e) => setCategoryDescription(e.target.value)}/>
          </FloatingLabel>

          <FloatingLabel controlId="floatingNotifiable" label={NotifiableLabelText}>
            <Form.Select
                name="notifiable"
                value={categoryNotifiable}
                onChange={(e) => setCategoryNotifiable(e.target.value)}
                style={{ margin: '0px 0px 32px' }}>>
              <option value="false">{NoText}</option>
              <option value="true">{YesText}</option>
            </Form.Select>
          </FloatingLabel>

        <ButtonGroup style={{width: "100%", paddingLeft: "10%", paddingRight: "10%"}}>
          <Button
              variant="danger"
              onClick={() => goToCategories()}
              style={{margin: "auto", display: "block", padding: "0", textAlign: "center"}}
          >{CancelButtonText}</Button>
          <Button
            variant="success"
            type="submit"
            style={{margin: "auto", display: "block", padding: "0", textAlign: "center"}}
          >{CreateButtonText}</Button>
        </ButtonGroup>

      </Form>
      <Modal className='successModal text-center' show={showModalCategoryAlreadyExists} open={showModalCategoryAlreadyExists} centered={true} size='lg'>
        <ModalBody>
          <h4 style={{margin: "32px"}}>{CategoryAlreadyExistsText}</h4>
          <ButtonGroup style={{width: "40%"}}>
            <Button
                variant="danger"
                onClick={(e) => toggleModalCategoryAlreadyExists(e)}
                style={{margin: "auto", display: "block", padding: "0", textAlign: "center"}}
            >{CancelButtonText}</Button>
          </ButtonGroup>
        </ModalBody>
      </Modal>

      <Modal className='successModal text-center' show={showEnterCategoryModal} open={showEnterCategoryModal} centered={true} size='lg'>
        <ModalBody>
          <h4 style={{margin: "32px"}}>{PleaseEnterCategoryNameText}</h4>
          <ButtonGroup style={{width: "40%"}}>
            <Button
                variant="danger"
                onClick={(e) => toggleEnterCategoryModal(e)}
                style={{margin: "auto", display: "block", padding: "0", textAlign: "center"}}
            >{CancelButtonText}</Button>
          </ButtonGroup>
        </ModalBody>
      </Modal>

    </Container>
  )
}

export default CreateCategory;
