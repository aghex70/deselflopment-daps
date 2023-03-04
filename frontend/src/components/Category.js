import React, {useEffect, useState} from 'react';
import {Button, ButtonGroup, Container, FloatingLabel, Form} from "react-bootstrap";
import {useLocation, useParams} from 'react-router-dom'
import CategoryService from "../services/category";
import DapsHeader from "./Header";
import checkAccess, {goToCategories} from "../utils/helpers";
import {
  CancelButtonText,
  DescriptionLabelText,
  EditButtonText,
  EditCategoryHeaderText,
  NameLabelText, NoText, NotifiableLabelText,
  ViewCategoryHeaderText, YesText
} from "../utils/texts";
import toBoolean from "validator/es/lib/toBoolean";

const Category = () => {
  checkAccess();
  const [categoryName, setCategoryName] = useState("");
  const [categoryDescription, setCategoryDescription] = useState("");
  const [categoryNotifiable, setCategoryNotifiable] = useState("false");
  const { id } = useParams();
  const location = useLocation();
  const enableEdit = location.state.action === "edit";

  const handleSubmit = (e) => {
    e.preventDefault();

    const data = {
      name: categoryName,
      description: categoryDescription,
      notifiable: typeof(categoryNotifiable) == "boolean" ? categoryNotifiable : toBoolean(categoryNotifiable),
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
      }
    )
  }

  useEffect(() => {
    CategoryService.getCategory(id).then(
      (response) => {
        if (response.status === 200) {
          setCategoryName(response.data.name);
          setCategoryDescription(response.data.description);
          setCategoryNotifiable(response.data.notifiable);
        }
      }
    ).catch(
      (error) => {
      }
    )
  }
  ,[id]);

  return (
    <Container>
      <DapsHeader />
      <h1 className="text-center">{enableEdit ? EditCategoryHeaderText : ViewCategoryHeaderText}</h1>
      <Form onSubmit={(e) => handleSubmit(e)}>
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

        <FloatingLabel controlId="floatingNotifiable" label={NotifiableLabelText}>
          <Form.Select
              name="notifiable"
              value={categoryNotifiable}
              onChange={(e) => setCategoryNotifiable(e.target.value)}
              style={{ margin: '0px 0px 32px' }}
              disabled={!enableEdit}>
            <option value="false">{NoText}</option>
            <option value="true">{YesText}</option>
          </Form.Select>
        </FloatingLabel>

        {enableEdit ?
          (
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
          >{EditButtonText}</Button>
        </ButtonGroup>
          ) : (
            <Button
              variant="danger"
              onClick={() => goToCategories()}
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


