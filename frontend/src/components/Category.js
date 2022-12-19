import React, {useEffect, useState} from 'react';
import {Button, ButtonGroup, Container, FloatingLabel, Form} from "react-bootstrap";
import {useLocation, useParams} from 'react-router-dom'
import CategoryService from "../services/category";
import DapsHeader from "./Header";
import checkAccess from "../utils/helpers";
import {
  CancelButtonText,
  DescriptionLabelText,
  EditButtonText,
  EditCategoryHeaderText,
  NameLabelText,
  ViewCategoryHeaderText
} from "../utils/texts";

const Category = () => {
  checkAccess();
  const [categoryName, setCategoryName] = useState("");
  const [categoryDescription, setCategoryDescription] = useState("");
  const { id } = useParams();
  const location = useLocation();
  const enableEdit = location.state.action === "edit";

  const navigateCategories = () => {
    window.location.href = "/categories";
  }

  const handleSubmit = (e) => {
    e.preventDefault();

    const data = {
      name: categoryName,
      description: categoryDescription,
    }

    CategoryService.updateCategory(id, data).then(
      (response) => {
        if (response.status === 200) {
          window.location.href = "/categories";
        } else {
          window.location.reload()
        }
      }
    ).catch(
      (error) => {
        error = new Error("Update category failed!");
      }
    )
  }

  useEffect(() => {
    CategoryService.getCategory(id).then(
      (response) => {
        if (response.status === 200) {
          setCategoryName(response.data.name);
          setCategoryDescription(response.data.description);
        }
      }
    ).catch(
      (error) => {
        // window.location.href = "/categories";
      }
    )
  }
  ,[id]);

  return (
    <Container>
      <DapsHeader />
      <h1 className="text-center">{enableEdit ? EditCategoryHeaderText : ViewCategoryHeaderText}</h1>
      <Form  onSubmit={(e) => handleSubmit(e)}>
        <FloatingLabel
          controlId="floatingName"
          label={NameLabelText}
        >
          <Form.Control
            type="name"
            value={categoryName}
            onChange={(e) => setCategoryName(e.target.value)}
            disabled={!enableEdit}
          />
        </FloatingLabel>

        <FloatingLabel controlId="floatingDescription" label={DescriptionLabelText}>
          <Form.Control
            as="textarea"
            placeholder="Description"
            style={{ height: '100px', margin: '0px 0px 32px' }}
            type="description"
            value={categoryDescription}
            onChange={(e) => setCategoryDescription(e.target.value)}
            disabled={!enableEdit}/>
        </FloatingLabel>

        {enableEdit ?
          (
        <ButtonGroup style={{width: "100%", paddingLeft: "10%", paddingRight: "10%"}}>
          <Button
            variant="success"
            type="submit"
            onClick={(e) => handleSubmit(e)}
            style={{margin: "auto", display: "block", padding: "0", textAlign: "center"}}
          >{EditButtonText}</Button>
          <Button
            variant="danger"
            onClick={() => navigateCategories()}
            style={{margin: "auto", display: "block", padding: "0", textAlign: "center"}}
          >{CancelButtonText}</Button>
        </ButtonGroup>
          ) : (
            <Button
              variant="danger"
              onClick={() => navigateCategories()}
              style={{margin: "auto", display: "block", padding: "0", textAlign: "center"}}
            >{CancelButtonText}</Button>
          )
        }
      </Form>
  </Container>
  )
}
;


export default Category;


