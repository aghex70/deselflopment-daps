import React, {useEffect, useState} from 'react';
import {Button, ButtonGroup, Container, FloatingLabel, Form} from "react-bootstrap";
import {useLocation, useNavigate, useParams} from 'react-router-dom'
import TodoService from "../services/todo";
import DapsHeader from "./Header";
import checkAccess from "../utils/helpers";
import toBoolean from "validator/es/lib/toBoolean";
import {
    CancelButtonText,
    DescriptionLabelText,
    EditButtonText,
    EditTodoHeaderText,
    HighestPriorityText,
    HighPriorityText,
    LinkLabelText,
    LowestPriorityText,
    LowPriorityText,
    MediumPriorityText,
    NameLabelText,
    NoRecurringText,
    PriorityLabelText,
    RecurringLabelText,
    ReturnButtonText,
    ViewTodoHeaderText,
    YesRecurringText
} from "../utils/texts";

const Todo = () => {
    checkAccess();
    const [todoName, setTodoName] = useState("");
    const [todoDescription, setTodoDescription] = useState("");
    const [todoLink, setTodoLink] = useState("");
    const [todoPriority, setTodoPriority] = useState("");
    const [todoRecurring, setTodoRecurring] = useState("");
    const [todoCategoryId, setTodoCategoryId] = useState();
    const [, setTodoCategoryName] = useState();
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
        recurring: typeof(todoRecurring) == "boolean" ? todoRecurring : toBoolean(todoRecurring),
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
          }
        )
      }
      ,[id, categoryId]);

    return (
      <Container>
        <DapsHeader />
        <h1 className="text-center">{enableEdit ? EditTodoHeaderText : ViewTodoHeaderText}</h1>
        <Form onSubmit={(e) => handleSubmit(e)}>
          <FloatingLabel
            controlId="floatingName"
            label={NameLabelText}
            value={todoName}
            onChange={(e) => setTodoName(e.target.value)}
          >
            <Form.Control type="name" placeholder="Name" value={todoName} disabled={!enableEdit}/>
          </FloatingLabel>

        {!enableEdit && todoDescription &&
          <FloatingLabel
            controlId="floatingDescription"
            label={DescriptionLabelText}
            value={todoDescription}
            onChange={(e) => setTodoDescription(e.target.value)}
          >
            <Form.Control type="description" placeholder="Description" value={todoDescription} disabled={!enableEdit}/>
          </FloatingLabel>
        }


            <FloatingLabel controlId="floatingPriority" label={PriorityLabelText}>
                <Form.Select
              name="priority"
              value={todoPriority}
              onChange={(e) => setTodoPriority(e.target.value)}
              style={{ margin: '0px 0px 32px' }}
              disabled={!enableEdit}
            >
                    <option style={{color: "grey"}} value="1">{LowestPriorityText}</option>
                    <option style={{color: "blue"}} value="2">{LowPriorityText}</option>
                    <option style={{color: "green"}} value="3">{MediumPriorityText}</option>
                    <option style={{color: "orange"}} value="4">{HighPriorityText}</option>
                    <option style={{color: "red"}} value="5">{HighestPriorityText}</option>
                </Form.Select>
            </FloatingLabel>

            <FloatingLabel controlId="floatingRecurring" label={RecurringLabelText}>
                <Form.Select
              name="recurring"
              value={todoRecurring}
              onChange={(e) => setTodoRecurring(e.target.value)}
              style={{ margin: '0px 0px 32px' }}
              disabled={!enableEdit}
            >
                    <option value="false">{NoRecurringText}</option>
                    <option value="true">{YesRecurringText}</option>
                </Form.Select>
            </FloatingLabel>

            <FloatingLabel
                controlId="floatingLink"
                label={LinkLabelText}
                value={todoLink}
                onChange={(e) => setTodoLink(e.target.value)}
            >
                <Form.Control type="link" placeholder="Link" value={todoLink} disabled={!enableEdit}/>
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
                onClick={() => navigateTodos()}
                style={{margin: "auto", display: "block", padding: "0", textAlign: "center"}}
              >{CancelButtonText}</Button>
            </ButtonGroup>
            ) : (
              <ButtonGroup style={{width: "40%", marginLeft: "30%"}}>
              <Button
                variant="success"
                onClick={() => navigateTodos()}
                style={{margin: "auto", display: "block", padding: "0", textAlign: "center"}}
              >{ReturnButtonText}</Button>
            </ButtonGroup>
            )
          }
        </Form>
      </Container>
    )
  }
;


export default Todo;


