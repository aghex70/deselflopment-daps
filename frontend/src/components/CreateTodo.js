import React, {useEffect, useState} from 'react';
import {Button, ButtonGroup, Container, FloatingLabel, Form} from "react-bootstrap";
import {useLocation, useNavigate} from 'react-router-dom'
import TodoService from "../services/todo";
import toBoolean from "validator/es/lib/toBoolean";
import DapsHeader from "./Header";
import checkAccess from "../utils/helpers";

const CreateTodo = () => {
    checkAccess();
    const [todoName, setTodoName] = useState("");
    const [todoDescription, setTodoDescription] = useState("");
    const [todoLink, setTodoLink] = useState("");
    const [todoPriority, setTodoPriority] = useState("");
    const [todoRecurring, setTodoRecurring] = useState("");
    const [disablePriority, setDisablePriority] = useState(false);
    const [disableRecurring, setDisableRecurring] = useState(false);
    const location = useLocation();
    const categoryId = location.state.categoryId;
    const categoryName = location.state.categoryName;
    const navigate = useNavigate();

    console.log("categoryName: " + categoryName);
    const disablePrioritySelect = () => {
      if (!disablePriority) {
        if (todoPriority in ["1", "2", "3", "4", "5"]) {
          setDisablePriority(true);
        }
      }
    }
    disablePrioritySelect();

    const disableRecurringSelect = () => {
      if (!disableRecurring) {
        if (todoRecurring === "true" || todoRecurring === "false") {
          setDisableRecurring(true);
        }
      }
    }
    disableRecurringSelect();



    const navigateTodos = () => {
      navigate("/todos", {state: {categoryId: location.state.categoryId, categoryName: location.state.categoryName}});
    }

    const handleSubmit = (e) => {
      e.preventDefault();

      const data = {
        name: todoName,
        description: todoDescription,
        link: todoLink,
        // priority will be casted to int
        priority: parseInt(todoPriority),
        // priority: todoPriority.,
        recurring: toBoolean(todoRecurring),
        category_id: categoryId,
      }

      TodoService.createTodo(data).then(
        (response) => {
          if (response.status === 201) {
            navigateTodos(categoryId, categoryName);
          } else {
            window.location.reload()
          }
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
        <h1 style={{ margin: '0px 0px 32px' }} className="text-center">Create Todo</h1>
        <Form onSubmit={(e) => handleSubmit(e)}>
          <FloatingLabel
            controlId="floatingName"
            label="Name"
            value={todoName}
            onChange={(e) => setTodoName(e.target.value)}
          >
            <Form.Control type="name" placeholder="Name" />
          </FloatingLabel>

          <FloatingLabel controlId="floatingDescription" label="Description">
            <Form.Control
              as="textarea"
              placeholder="Description"
              style={{ height: '100px', margin: '0px 0px 32px' }}
              type="description"
              value={todoDescription}
              onChange={(e) => setTodoDescription(e.target.value)}/>
          </FloatingLabel>


            <FloatingLabel controlId="floatingPriority" label="Priority">
                <Form.Select
              name="priority"
              value={todoPriority}
              onChange={(e) => setTodoPriority(e.target.value)}
              style={{ margin: '0px 0px 32px' }}
            >
                    <option disabled={disablePriority}>Select priority</option>
                    <option style={{color: "grey"}} value="1">Lowest</option>
                    <option style={{color: "blue"}} value="2">Low</option>
                    <option style={{color: "green"}} value="3">Medium</option>
                    <option style={{color: "orange"}} value="4">High</option>
                    <option style={{color: "red"}} value="5">Highest</option>
                </Form.Select>
            </FloatingLabel>

            <FloatingLabel controlId="floatingRecurring" label="Recurring">
                <Form.Select
              name="recurring"
              value={todoRecurring}
              onChange={(e) => setTodoRecurring(e.target.value)}
                    style={{ margin: '0px 0px 32px' }}>>
                    <option disabled={disableRecurring}>Select recurring</option>
                    <option value="false">No</option>
                    <option value="true">Yes</option>
                </Form.Select>
            </FloatingLabel>

            <FloatingLabel
                controlId="floatingLink"
                label="Link"
                value={todoLink}
                onChange={(e) => setTodoLink(e.target.value)}
            >
                <Form.Control type="link" placeholder="Link" />
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
              onClick={() => navigateTodos()}
              style={{margin: "auto", display: "block", padding: "0", textAlign: "center"}}
            >Cancel</Button>
          </ButtonGroup>

        </Form>
      </Container>
    )
  }
;


export default CreateTodo;


