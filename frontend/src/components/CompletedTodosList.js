import React, {useEffect, useState} from 'react';
import TodoService from "../services/todo";
import {Button, ButtonGroup, Container} from "react-bootstrap";
import './TodosList.css';
import BootstrapTable from 'react-bootstrap-table-next';
import DapsHeader from "./Header";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {faTrashRestore, faTrash} from "@fortawesome/free-solid-svg-icons";
import {useNavigate} from "react-router-dom";
import checkAccess from "../utils/helpers";

const CompletedTodosList = () => {
  checkAccess();
  const [todos, setTodos] = useState([]);
  const navigate = useNavigate();

  // Color code the todo based on its priority
  const rowTextColor = (cell, row) => {
    const colors = ["red", "grey", "blue", "green", "orange"];
    return <div
      style={{color : colors[row.priority % 5], textDecoration: "line-through", cursor: "pointer"}}
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
    },
    {
      dataField: 'link',
      text: 'Actions',
      style:{width: '10%', verticalAlign: "middle"},
      formatter: actionsFormatter,
      headerAlign: 'center',
    }];

  const navigateToTodo = (id, categoryId, categoryName, action) => {
    navigate("/todo/" + id, {state: {categoryId: categoryId, action: action}});
  }


  const deleteTodo = (id, categoryId) => {
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

  const activateTodo = (id, categoryId) => {
    TodoService.activateTodo(id, categoryId).then(
      (response) => {
        if (response.status === 200) {
          window.location.reload();
        }
      }
    ).catch(
      (error) => {
        error = new Error("Activation failed!");
      })
  }

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
          <Button style={{width: "15%", margin: "auto", padding: "0", textAlign: "center"}}
                  title="Reactivate"
                  variant="outline-success"
                  onClick={() => activateTodo(row.id, row.category_id)}
          >
            <FontAwesomeIcon icon={faTrashRestore} />
          </Button>
          <Button style={{width: "15%", margin: "auto", display: "block", padding: "0", textAlign: "center"}}
                  title="Delete"
                  variant="outline-danger"
                  onClick={() => deleteTodo(row.id, row.category_id)}
          >
            <FontAwesomeIcon icon={faTrash} />
          </Button>
        </ButtonGroup>
      </div>
    );
  }

  if (!localStorage.getItem("access_token")) {
    window.location.href = "/login";
  }

  useEffect(() => {
    if (!todos || todos.length === 0) {
      TodoService.getCompletedTodos().then(
        (response) => {
          if (response.status === 200 && response.data) {
            setTodos(response.data);
          }
        }
      ).catch(
        (error) => {
          error = new Error("Login failed!");
        })
    }
  },[todos]);

  function indication() {
    return "You better complete some Todos first!!!";
  }

  return (
    <Container>
      <DapsHeader />
      <h1 className="text-center">Completed Todos</h1>
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

export default CompletedTodosList;
