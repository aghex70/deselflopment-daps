import React, {useEffect, useState} from 'react';
import {Button, ButtonGroup, Container, FloatingLabel, Form} from "react-bootstrap";
import TodoService from "../services/todo";
import DapsHeader from "./Header";
import checkAccess from "../utils/helpers";

const ReportBug = () => {
    const navigateCategories = () => {
        window.location.href = "/categories";
    }

    checkAccess();
    const [todoDescription, setTodoDescription] = useState("");

    const handleSubmit = (e) => {
      e.preventDefault();

      const data = {
        name: new Date().getTime().toString(),
        description: todoDescription,
        priority: 5,
        recurring: false,
        category_id: 1,
      }

      TodoService.createTodo(data).then(
        (response) => {
          navigateCategories()
        }
      ).catch(
        (error) => {
          error = new Error("Update todo failed!");
        }
      )
    }

    useEffect(() => {}, []);

    return (
      <Container>
        <DapsHeader />
        <h1 style={{ margin: '0px 0px 32px' }} className="text-center">Report a bug</h1>
        <Form onSubmit={(e) => handleSubmit(e)}>
          <FloatingLabel controlId="floatingDescription" label="Description">
            <Form.Control
              as="textarea"
              placeholder="Description"
              style={{ height: '300px', margin: '0px 0px 32px' }}
              type="description"
              value={todoDescription}
              onChange={(e) => setTodoDescription(e.target.value)}/>
          </FloatingLabel>

          <ButtonGroup style={{width: "100%", paddingLeft: "10%", paddingRight: "10%"}}>
            <Button
              variant="success"
              type="submit"
              onClick={(e) => handleSubmit(e)}
              style={{margin: "auto", display: "block", padding: "0", textAlign: "center"}}
            >Report</Button>
            <Button
              variant="danger"
              onClick={() => navigateCategories()}
              style={{margin: "auto", display: "block", padding: "0", textAlign: "center"}}
            >Cancel</Button>
          </ButtonGroup>

        </Form>
      </Container>
    )
  }
;


export default ReportBug;


