import React, { useState } from "react";
import {
  Button,
  ButtonGroup,
  Container,
  Form,
  FloatingLabel,
} from "react-bootstrap";
import TodoService from "../services/todo";
import DapsHeader from "./Header";
import checkAccess, { checkValidToken, goToCategories } from "../utils/helpers";
import {
  CancelButtonText,
  DescriptionLabelText,
  ReportABugHeaderText,
  ReportButtonText,
} from "../utils/texts";

const ReportBug = () => {
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
    };

    TodoService.createTodo(data)
      .then((response) => {
        goToCategories();
      })
      .catch((error) => {
        checkValidToken(error);
      });
  };

  return (
    <Container>
      <DapsHeader />
      <h1 style={{ margin: "0px 0px 32px" }} className="text-center">
        {ReportABugHeaderText}
      </h1>
      <Form onSubmit={handleSubmit}>
        <FloatingLabel
          controlId="floatingDescription"
          label={DescriptionLabelText}
        >
          <Form.Control
            as="textarea"
            placeholder="Description"
            style={{ height: "300px", margin: "0px 0px 32px" }}
            type="description"
            value={todoDescription}
            onChange={(e) => setTodoDescription(e.target.value)}
          />
        </FloatingLabel>

        <ButtonGroup
          style={{ width: "100%", paddingLeft: "10%", paddingRight: "10%" }}
        >
          <Button
            variant="danger"
            onClick={() => (window.location.href = "/categories")}
            style={{
              margin: "auto",
              display: "block",
              padding: "0",
              textAlign: "center",
            }}
          >
            {CancelButtonText}
          </Button>
          <Button
            variant="success"
            type="submit"
            style={{
              margin: "auto",
              display: "block",
              padding: "0",
              textAlign: "center",
            }}
          >
            {ReportButtonText}
          </Button>
        </ButtonGroup>
      </Form>
    </Container>
  );
};

export default ReportBug;
