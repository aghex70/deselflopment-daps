import React, { useEffect, useState } from "react";
import TodoService from "../services/todo";
import { Button, ButtonGroup, Container } from "react-bootstrap";
import "./TodosList.css";
import BootstrapTable from "react-bootstrap-table-next";
import DapsHeader from "./Header";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faPencil, faTrash } from "@fortawesome/free-solid-svg-icons";
import { useNavigate } from "react-router-dom";
import checkAccess, { checkValidToken } from "../utils/helpers";
import {
  DeleteIconText,
  EditIconText,
  HeaderActionsText,
  HeaderNameText,
  RecurringTodosHeaderText,
  RecurringTodosIndicationText,
} from "../utils/texts";

const RecurringTodosList = () => {
  checkAccess();
  const [todos, setTodos] = useState([]);
  const navigate = useNavigate();

  // Color code the todo based on its priority
  const rowTextColor = (cell, row) => {
    const colors = ["red", "grey", "blue", "green", "orange"];
    return (
      <div
        style={{ color: colors[row.priority % 5], cursor: "pointer" }}
        onClick={() =>
          navigateToTodo(row.id, row.category_id, row.category_id, "view")
        }
      >
        {row.name}
      </div>
    );
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
              width: "20%",
              margin: "auto",
              padding: "0",
              textAlign: "center",
            }}
            variant="outline-primary"
            title={EditIconText}
            onClick={() => navigateToTodo(row.id, row.category_id, 0, "edit")}
          >
            <FontAwesomeIcon icon={faPencil} />
          </Button>

          <Button
            style={{
              width: "20%",
              margin: "auto",
              display: "block",
              padding: "0",
              textAlign: "center",
            }}
            variant="outline-danger"
            title={DeleteIconText}
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
      dataField: "name",
      text: HeaderNameText,
      style: { width: "80%" },
      formatter: rowTextColor,
    },
    {
      dataField: "actions",
      text: HeaderActionsText,
      style: { width: "20%" },
      formatter: actionsFormatter,
      headerAlign: "center",
    },
  ];

  const navigateToTodo = (id, categoryId, categoryName, action) => {
    navigate("/todo/" + id, {
      state: { categoryId: categoryId, action: action },
    });
  };

  const deleteTodo = (id) => {
    TodoService.deleteTodo(id)
      .then((response) => {
        if (response.status === 204) {
          window.location.reload();
        }
      })
      .catch((error) => {
        console.log("error", error); // checkValidToken(error);
      });
  };

  useEffect(() => {
    if (!todos || todos.length === 0) {
        let fields = "recurring=1&completed=0"
        TodoService.getTodos(fields)
        .then((response) => {
          if (response.status === 200 && response.data) {
            setTodos(response.data);
          }
        })
        .catch((error) => {
          console.log("error", error); // checkValidToken(error);
        });
    }
  }, [todos]);

  function indication() {
    return RecurringTodosIndicationText;
  }

  return (
    <Container>
      <DapsHeader />
      <h1 className="text-center">{RecurringTodosHeaderText}</h1>
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

export default RecurringTodosList;
