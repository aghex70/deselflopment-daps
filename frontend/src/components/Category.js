import React, {useEffect, useState} from 'react';
import {Button, ButtonGroup, Container, FloatingLabel, Form} from "react-bootstrap";
import {useLocation, useParams} from 'react-router-dom'
import CategoryService from "../services/category";
import DapsHeader from "./Header";
import checkAccess from "../utils/helpers";

const Category = () => {
  checkAccess();
  const [categoryName, setCategoryName] = useState("");
  const [categoryDescription, setCategoryDescription] = useState("");
  const { id } = useParams();
  const location = useLocation();
  const enableEdit = location.state.action === "edit";
  console.log("EDIT: " + enableEdit);

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
          console.log("success!!!")
          window.location.href = "/categories";
        } else {
          console.log("NO success!!!")
          window.location.reload()
        }
      }
    ).catch(
      (error) => {
        console.log(error);
        error = new Error("Update category failed!");
      }
    )
  }

  useEffect(() => {
    CategoryService.getCategory(id).then(
      (response) => {
        console.log(response);
        if (response.status === 200) {
          setCategoryName(response.data.name);
          console.log(response.data.name);
          setCategoryDescription(response.data.description);
          console.log(response.data);
        }
      }
    ).catch(
      (error) => {
        // window.location.href = "/categories";
        console.log(error)
      }
    )
  }
  ,[]);

  return (
    <Container>
      <DapsHeader />
      <h1 className="text-center">{enableEdit ? "Edit category" : "View category"}</h1>
      <Form  onSubmit={(e) => handleSubmit(e)}>
        <FloatingLabel
          controlId="floatingName"
          label="Name"
        >
          <Form.Control
            type="name"
            value={categoryName}
            onChange={(e) => setCategoryName(e.target.value)}
            disabled={!enableEdit}
          />
        </FloatingLabel>

        <FloatingLabel controlId="floatingDescription" label="Description">
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
            onClick={() => handleSubmit()}
            style={{margin: "auto", display: "block", padding: "0", textAlign: "center"}}
          >Edit</Button>
          <Button
            variant="danger"
            onClick={() => navigateCategories()}
            style={{margin: "auto", display: "block", padding: "0", textAlign: "center"}}
          >Cancel</Button>
        </ButtonGroup>
          ) : (
            <Button
              variant="success"
              onClick={() => navigateCategories()}
              style={{margin: "auto", display: "block", padding: "0", textAlign: "center"}}
            >Return</Button>
          )
        }
      </Form>
  </Container>
  )
}
;


export default Category;


