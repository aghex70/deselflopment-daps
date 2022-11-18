import React, {useEffect, useState} from 'react';
import TodoService from "../services/todo";
import {Button, ButtonGroup, Container} from "react-bootstrap";
import './TodosList.css';
import BootstrapTable from 'react-bootstrap-table-next';
import DapsHeader from "./Header";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {faPencil, faTrash} from "@fortawesome/free-solid-svg-icons";
import {useLocation, useNavigate} from "react-router-dom";

const RecurringTodosList = () => {
  const [todos, setTodos] = useState([]);
  const location = useLocation();
  const navigate = useNavigate();

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
          <Button style={{width: "20%", margin: "auto", padding: "0", textAlign: "center"}}
                  variant="outline-primary"
                  title="Edit"
                  onClick={() => navigateToTodo(row.id, row.category_id, 0, "edit")}
          >
            <FontAwesomeIcon icon={faPencil} />
          </Button>

          <Button style={{width: "20%", margin: "auto", display: "block", padding: "0", textAlign: "center"}}
                  variant="outline-danger"
                  title="Delete"
                  onClick={() => deleteTodo(row.id, row.category_id)}
          >
            <FontAwesomeIcon icon={faTrash} />
          </Button>
        </ButtonGroup>
      </div>
    );
  }

  const columns = [
    {
      dataField: 'name',
      text: 'Name',
      style:{'width' : '80%'},
      formatter: rowTextColor,
    }, {
      // dataField: 'link',
      text: 'Actions',
      style:{'width' : '20%'},
      formatter: actionsFormatter,
      headerAlign: 'center',
    }];

  if (!localStorage.getItem("access_token")) {
    window.location.href = "/login";
  }

  const navigateToTodo = (id, categoryId, categoryName, action) => {
    console.log("TODOLIST navigatin to TODO id: " + id);
    console.log("TODOLIST categoryId: " + categoryId);
    console.log("TODOLIST categoryName: " + categoryName);
    // console.log("TODOLIST location.state: " + Object.keys(location.state), Object.values(location.state));
    navigate("/todo/" + id, {state: {categoryId: categoryId, action: action}});
  }

  const deleteTodo = (id, categoryId) => {
    TodoService.deleteTodo(id, categoryId).then(
      (response) => {
        console.log(response);
        if (response.status === 204) {
          console.log("Deleted");
          window.location.reload();
        }
      }
    ).catch(
      (error) => {
        console.log(error);
        error = new Error("Deletion failed!");
      })
  }

  useEffect(() => {
    console.log("hola")
    if (!todos || todos.length === 0) {
      TodoService.getRecurringTodos().then(
        (response) => {
          console.log(response);
          if (response.status === 200 && response.data) {
            console.log("response.data " + response.data);
            setTodos(response.data);
          } else {
            console.log("NO DATA MY FRIEND");
          }

        }
      ).catch(
        (error) => {
          console.log(error);
          error = new Error("Login failed!");
        })
    }
  },[]);

  return (
    <Container>
      <DapsHeader />
      <h1 className="text-center">Recurring Todos</h1>
      <BootstrapTable
        keyField='id'
        data={ todos }
        columns={ columns }
        trStyle={rowTextColor}
        hover={true}
        style={{display: "block", minHeight: "80%", width: "10%", overflow: "auto"}}
      />
    </Container>
  );
};

export default RecurringTodosList;
