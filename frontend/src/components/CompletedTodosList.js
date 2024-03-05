import React, { useEffect, useState } from "react";
import TodoService from "../services/todo";
import { Button, ButtonGroup, Container } from "react-bootstrap";
import "./TodosList.css";
import BootstrapTable from "react-bootstrap-table-next";
import DapsHeader from "./Header";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faTrashRestore, faTrash } from "@fortawesome/free-solid-svg-icons";
import { useNavigate } from "react-router-dom";
import checkAccess, {
  checkValidToken,
  clearLocalStorage,
  sortTodosByField,
} from "../utils/helpers";
import {
  CompletedTodosHeaderText,
  CompletedTodosIndicationText,
  DeleteIconText,
  HeaderActionsText,
  NameLabelText,
  ReactivateIconText,
} from "../utils/texts";

const CompletedTodosList = () => {
  checkAccess();
  const [todos, setTodos] = useState([]);
  const navigate = useNavigate();

  // Color code the todo based on its priority
  const rowTextColor = (cell, row) => {
    const colors = ["red", "grey", "blue", "green", "orange"];
    return (
      <div
        style={{
          color: colors[row.priority % 5],
          textDecoration: "line-through",
          cursor: "pointer",
        }}
        onClick={() =>
          navigateToTodo(row.id, row.category_id, row.category_id, "view")
        }
      >
        {row.name}
      </div>
    );
  };

  const columns = [
    {
      dataField: "name",
      text: NameLabelText,
      style: { width: "70%" },
      formatter: rowTextColor,
    },
    {
      dataField: "actions",
      text: HeaderActionsText,
      style: { width: "10%", verticalAlign: "middle" },
      formatter: actionsFormatter,
      headerAlign: "center",
    },
  ];

  const navigateToTodo = (id, categoryId, categoryName, action) => {
    clearLocalStorage([]);
    navigate("/todo/" + id, {
      state: { categoryId: categoryId, action: action },
    });
  };

  const deleteTodo = (id) => {
    TodoService.deleteTodo(id)
      .then((response) => {
        if (response.status === 204) {
          clearLocalStorage([]);
          window.location.reload();
        }
      })
      .catch((error) => {
        console.log("error", error); // checkValidToken(error);
      });
  };

  const activateTodo = (id) => {
    TodoService.activateTodo(id)
      .then((response) => {
        if (response.status === 200) {
          clearLocalStorage([]);
          window.location.reload();
        }
      })
      .catch((error) => {
        console.log("error", error); // checkValidToken(error);
      });
  };

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
        <ButtonGroup style={{ width: "100%" }}>
          <Button
            style={{
              width: "15%",
              margin: "auto",
              padding: "0",
              textAlign: "center",
            }}
            title={ReactivateIconText}
            variant="outline-success"
            onClick={() => activateTodo(row.id, row.category_id)}
          >
            <FontAwesomeIcon icon={faTrashRestore} />
          </Button>
          <Button
            style={{
              width: "15%",
              margin: "auto",
              display: "block",
              padding: "0",
              textAlign: "center",
            }}
            title={DeleteIconText}
            variant="outline-danger"
            onClick={() => deleteTodo(row.id, row.category_id)}
          >
            <FontAwesomeIcon icon={faTrash} />
          </Button>
        </ButtonGroup>
      </div>
    );
  }

  useEffect(() => {
    let todos = JSON.parse(localStorage.getItem("todos"));
    if (!todos) {
      let fields = "completed=1"
      TodoService.getTodos(fields)
        .then((response) => {
          if (response.status === 200 && response.data) {
            localStorage.setItem("todos", JSON.stringify(response.data));
            sortTodosByField("end_date", false, setTodos, null);
          }
        })
        .catch((error) => {
          console.log("error", error); // checkValidToken(error);
        });
    } else {
      sortTodosByField("end_date", false, setTodos, null);
    }
  }, []);

  function indication() {
    return CompletedTodosIndicationText;
  }

  return (
    <Container>
      <DapsHeader />
      <h1 className="text-center">{CompletedTodosHeaderText}</h1>
      <BootstrapTable
        keyField="id"
        data={todos}
        columns={columns}
        noDataIndication={indication}
        trStyle={rowTextColor}
        hover={true}
        striped={true}
        style={{
          display: "block",
          minHeight: "80%",
          width: "10%",
          overflow: "auto",
        }}
      />
    </Container>
  );
};

export default CompletedTodosList;
