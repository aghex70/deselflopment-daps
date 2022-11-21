import React, {useEffect, useState} from 'react';
import {faPencil, faPlay, faTrash, faCheck, faPlus} from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import TodoService from "../services/todo";
import {Button, ButtonGroup, Container} from "react-bootstrap";
import {useLocation, useNavigate} from "react-router-dom";
import './TodosList.css';
import BootstrapTable from 'react-bootstrap-table-next';
import DapsHeader from "./Header";
import checkAccess from "../utils/helpers";

const TodosList = () => {
  checkAccess();
  const [todos, setTodos] = useState([]);
  const location = useLocation();
  const navigate = useNavigate();
  const categoryId = location.state.categoryId;
  const [todoSpan, setTodoSpan] = useState({
    textAlign: "center",
    display: "none",
  });

  // Color code the todo based on its priority
  const rowTextColor = (cell, row) => {
    const colors = ["red", "grey", "blue", "green", "orange"];
    return <div
      style={{color : colors[row.priority % 5], cursor: "pointer"}}
      onClick={() => navigateToTodo(row.id, row.category_id, row.category_id, "view")}
    >
      {row.name}
    </div>;
  }

  const columns = [
    {
    dataField: 'name',
    text: 'Name',
    style:{'width' : '70%'},
    formatter: rowTextColor,
  }, {
    // dataField: 'link',
    text: 'Actions',
    style:{'width' : '30%'},
    formatter: actionsFormatter,
    headerAlign: 'center',
  }];

  if (!localStorage.getItem("access_token")) {
    window.location.href = "/login";
  }

  const navigateToTodo = (id, categoryId, categoryName, action) => {
    navigate("/todo/" + id, {state: {categoryId: location.state.categoryId, categoryName: location.state.categoryName, action: action}});
  }

  const deleteTodo = (id) => {
    TodoService.deleteTodo(id, categoryId).then(
      (response) => {
        if (response.status === 204) {
          window.location.reload();
        }
      }
    ).catch(
      (error) => {
        error = new Error("Deletion failed!");
      })
  }

  const completeTodo = (id) => {
    TodoService.completeTodo(id, categoryId).then(
      (response) => {
        if (response.status === 200) {
          window.location.reload();
        }
      }
    ).catch(
      (error) => {
        error = new Error("Completion failed!");
      })
  }

  const startTodo = (id) => {
    TodoService.startTodo(id, categoryId).then(
      (response) => {
        if (response.status === 200) {
          window.location.reload();
        }
      }
    ).catch(
      (error) => {
        error = new Error("Completion failed!");
      })
  }

  const createTodo = () => {
    navigate("/create-todo", {state: {categoryId: location.state.categoryId, categoryName: location.state.categoryName}});
  }

  useEffect(() => {
    if (!todos || todos.length === 0) {
      TodoService.getTodos(categoryId).then(
        (response) => {
          if (response.status === 200 && response.data) {
            setTodos(response.data);
            setTodoSpan({
              textAlign: "center",
              display: "block",
            })
          }
        }
      ).catch(
        (error) => {
          error = new Error("Login failed!");
        })
    }
  },[]);

  function actionsFormatter(cell, row) {
    return (
      <div
        style={{
          textAlign: "center",
          cursor: "pointer",
          lineHeight: "normal",
          width: "100%",
          flexDirection: "row",
        }}
      >

      <ButtonGroup style={{width: "100%"}}>
        {row.active === false ? (
          <Button style={{width: "15%", margin: "auto", display: "block", padding: "0", textAlign: "center", }}
                  title="Start" variant="outline-warning" onClick={() => startTodo(row.id)}>
            <FontAwesomeIcon icon={faPlay} />
          </Button>
        ) : (
          <Button style={{width: "15%", margin: "auto", display: "block", padding: "0", textAlign: "center"}}
                  title="Complete" variant="outline-success" onClick={() => completeTodo(row.id)}>
            <FontAwesomeIcon icon={faCheck} />
          </Button>
        )
        }

        <Button style={{width: "15%", margin: "auto", padding: "0", textAlign: "center"}}
                title="Edit" variant="outline-primary" onClick={() => navigateToTodo(row.id, categoryId, 0,"edit")}>
          <FontAwesomeIcon icon={faPencil} />
        </Button>

        <Button style={{width: "15%", margin: "auto", display: "block", padding: "0", textAlign: "center"}}
                title="Delete" variant="outline-danger" onClick={() => deleteTodo(row.id)}>
          <FontAwesomeIcon icon={faTrash} />
        </Button>
      </ButtonGroup>
      </div>
    );
  }

  function indication() {
    return <span className="createIcon" onClick={() => createTodo()}>
      <FontAwesomeIcon className="createIcon" icon={faPlus} />Create a new Todo</span>
  }

      return (
        <Container>
          <DapsHeader />
          <h1 className="text-center">{location.state.categoryName}</h1>
          <span style={todoSpan} className="createIcon" onClick={() => createTodo()}>
          <FontAwesomeIcon className="createIcon" icon={faPlus} />Create a new Todo</span>
          <BootstrapTable
            keyField='id'
            data={ todos }
            columns={ columns }
            noDataIndication={ indication }
            trStyle={rowTextColor}
            hover={true}
            style={{display: "block", minHeight: "80%", width: "10%", overflow: "auto"}}
          />
        </Container>
      );
    };

export default TodosList;
