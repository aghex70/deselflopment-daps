import React, {useEffect, useState} from 'react';
import {Button, ButtonGroup, Container, FloatingLabel, Form} from "react-bootstrap";
import {useLocation, useNavigate, useParams} from 'react-router-dom'
import TodoService from "../services/todo";
import DapsHeader from "./Header";
import checkAccess from "../utils/helpers";
import toBoolean from "validator/es/lib/toBoolean";

const Todo = () => {
    checkAccess();
    const [todoName, setTodoName] = useState("");
    const [todoDescription, setTodoDescription] = useState("");
    const [todoLink, setTodoLink] = useState("");
    const [todoPriority, setTodoPriority] = useState("");
    const [todoRecurring, setTodoRecurring] = useState("");
    const [todoCategoryId, setTodoCategoryId] = useState();
    const [todoCategoryName, setTodoCategoryName] = useState();
    const location = useLocation();
    const categoryId = location.state.categoryId;
    const categoryName = location.state.categoryName;
    const enableEdit = location.state.action === "edit";
    const { id } = useParams();
    const navigate = useNavigate();

    const navigateTodos = () => {
      navigate("/todos", {state: {categoryId: location.state.categoryId, categoryName: location.state.categoryName}});
    }

    const handleSubmit = (e) => {
      e.preventDefault();

      const data = {
        name: todoName,
        description: todoDescription,
        link: todoLink,
        priority: parseInt(todoPriority),
        recurring: toBoolean(todoRecurring),
        category_id: todoCategoryId,
      }

      TodoService.updateTodo(id, data).then(
        (response) => {
          if (response.status === 200) {
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

    useEffect(() => {
        TodoService.getTodo(id, categoryId).then(
          (response) => {
            if (response.status === 200) {
              setTodoName(response.data.name);
              setTodoDescription(response.data.description);
              setTodoLink(response.data.link);
              setTodoPriority(response.data.priority);
              setTodoRecurring(response.data.recurring);
              setTodoCategoryId(response.data.category_id);
              setTodoCategoryName(response.data.category_name);
            }
          }
        ).catch(
          (error) => {
            // window.location.href = "/categories";
          }
        )
      }
      ,[id, categoryId]);

    return (
      <Container>
        <DapsHeader />
        <h1 className="text-center">{enableEdit ? "Edit todo" : "View todo"}</h1>
        <Form onSubmit={(e) => handleSubmit(e)}>
          <FloatingLabel
            controlId="floatingName"
            label="Name"
            value={todoName}
            onChange={(e) => setTodoName(e.target.value)}
          >
            <Form.Control type="name" placeholder="Name" value={todoName} disabled={!enableEdit}/>
          </FloatingLabel>

          <FloatingLabel
            controlId="floatingDescription"
            label="Description"
            value={todoDescription}
            onChange={(e) => setTodoDescription(e.target.value)}
          >
            <Form.Control type="description" placeholder="Description" value={todoDescription} disabled={!enableEdit}/>
          </FloatingLabel>

          <FloatingLabel
            controlId="floatingLink"
            label="Link"
            value={todoLink}
            onChange={(e) => setTodoLink(e.target.value)}
          >
            <Form.Control type="link" placeholder="Link" value={todoLink} disabled={!enableEdit}/>
          </FloatingLabel>

          <FloatingLabel controlId="floatingPriority" label="Priority">
            <Form.Select
              name="priority"
              value={todoPriority}
              onChange={(e) => setTodoPriority(e.target.value)}
              style={{ margin: '0px 0px 32px' }}
              disabled={!enableEdit}
            >
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
              style={{ margin: '0px 0px 32px' }}
              disabled={!enableEdit}
            >
              <option value="false">No</option>
              <option value="true">Yes</option>
            </Form.Select>
          </FloatingLabel>

          <FloatingLabel
            controlId="floatingCategory"
            label="Category"
            value={todoCategoryName}
            onChange={(e) => setTodoCategoryName(e.target.value)}
            placeholder={todoCategoryName}
          >
            <Form.Control disabled={true} type="name" placeholder="Name" value={todoCategoryName}/>
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
                onClick={() => navigateTodos()}
                style={{margin: "auto", display: "block", padding: "0", textAlign: "center"}}
              >Cancel</Button>
            </ButtonGroup>
            ) : (
              <ButtonGroup style={{width: "20%", marginLeft: "40%"}}>
              <Button
                variant="success"
                onClick={() => navigateTodos()}
                style={{margin: "auto", display: "block", padding: "0", textAlign: "center"}}
              >Return</Button>
            </ButtonGroup>
            )
          }
        </Form>
      </Container>
    )
  }
;


export default Todo;


