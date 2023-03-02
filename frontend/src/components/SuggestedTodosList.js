import React, {useEffect, useState} from 'react';
import TodoService from "../services/todo";
import {Button, ButtonGroup, Container, Modal, ModalBody} from "react-bootstrap";
import './TodosList.css';
import BootstrapTable from 'react-bootstrap-table-next';
import DapsHeader from "./Header";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {faCheck, faPencil, faPlay, faTrash, faBackwardFast} from "@fortawesome/free-solid-svg-icons";
import {useNavigate} from "react-router-dom";
import checkAccess, {
    clearLocalStorage,
    goToSuggestedTodos, sortTodosByField
} from "../utils/helpers";
import {
    CancelButtonText,
    CannotDeleteActiveTodoText,
    CompleteIconText,
    DeleteIconText,
    EditIconText,
    HeaderActionsText,
    HeaderNameText, RestartIconText, StartIconText,
    SuggestedTodosHeaderText,
    SuggestedTodosIndicationText
} from "../utils/texts";

const SuggestedTodosList = () => {
  checkAccess();
  const [todos, setTodos] = useState([]);
  const [suggested, setSuggested] = useState(false);
  const navigate = useNavigate();
  const [showCannotDeleteActiveTodoModal, setShowCannotDeleteActiveTodoModal] = useState(false);

  const toggleCannotDeleteStartedTodoModal = () => {
    setShowCannotDeleteActiveTodoModal(!showCannotDeleteActiveTodoModal);
  }

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

    const completeTodo = (id, categoryId) => {
        TodoService.completeTodo(id, categoryId).then(
            (response) => {
                if (response.status === 200) {
                    goToSuggestedTodos();
                }
            }
        ).catch(
            (error) => {
            })
    }

    const startTodo = (id, categoryId) => {
        TodoService.startTodo(id, categoryId).then(
            (response) => {
                if (response.status === 200) {
                    goToSuggestedTodos();
                }
            }
        ).catch(
            (error) => {
            })
    }

    const restartTodo = (id, categoryId) => {
        TodoService.restartTodo(id, categoryId).then(
            (response) => {
                if (response.status === 200) {
                    clearLocalStorage([]);
                    window.location.reload();
                }
            }
        ).catch(
            (error) => {
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
            {row.active === false? (
                <Button style={{width: "20%", margin: "auto", display: "block", padding: "0", textAlign: "center", }}
                        title={StartIconText} variant="outline-success" onClick={() => startTodo(row.id, row.category_id)}>
                    <FontAwesomeIcon icon={faPlay} />
                </Button>
            ) : (
                <Button style={{width: "20%", margin: "auto", display: "block", padding: "0", textAlign: "center"}}
                        title={CompleteIconText} variant="outline-success" onClick={() => completeTodo(row.id, row.category_id)}>
                    <FontAwesomeIcon icon={faCheck} />
                </Button>
            )
            }
            {row.active === true ? (
                <Button style={{width: "15%", margin: "auto", padding: "0", textAlign: "center",
                }}
                        title={RestartIconText}
                        variant="outline-primary"
                        onClick={() => restartTodo(row.id, row.category_id)}>
                    <FontAwesomeIcon
                        icon={faBackwardFast} />
                </Button>
            ) : (
                <Button style={{width: "15%", margin: "auto", padding: "0", textAlign: "center"}}
                        title={EditIconText} variant="outline-primary" onClick={() => navigateToTodo(row.id, row.category_id, 0,"edit")}>
                    <FontAwesomeIcon icon={faPencil} />
                </Button>
            )
            }

          <Button style={{width: "20%", margin: "auto", display: "block", padding: "0", textAlign: "center"}}
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
      dataField: 'name',
      text: HeaderNameText,
      style:{'width' : '70%'},
      formatter: rowTextColor,
    }, {
      dataField: 'actions',
      text: HeaderActionsText,
      style:{'width' : '30%'},
      formatter: actionsFormatter,
      headerAlign: 'center',
    }];

  const navigateToTodo = (id, categoryId, categoryName, action) => {
      clearLocalStorage([]);
      navigate("/todo/" + id, {state: {categoryId: categoryId, action: action}});
  }

  const deleteTodo = (id, categoryId) => {
    TodoService.deleteTodo(id, categoryId).then(
      (response) => {
        if (response.status === 204) {
            clearLocalStorage([]);
            window.location.reload();
        }
      }
    ).catch(
      (error) => {
      })
  }

  useEffect(() => {
      if (!suggested && localStorage.getItem("auto-suggest") === "true") {
          TodoService.suggestTodos().then(
              (response) => {
                  setSuggested(true);
              }
          ).catch(
              (error) => {
              }
          )
      }

      let todos = JSON.parse(localStorage.getItem("todos"));
      if (!todos) {
          TodoService.getSuggestedTodos().then(
              (response) => {
                  if (response.status === 200 && response.data) {
                      localStorage.setItem("todos", JSON.stringify(response.data));
                      sortTodosByField("priority", true, setTodos, null);
                  }
              }
          ).catch(
              (error) => {
              })
      }
      else {
          sortTodosByField("priority", true, setTodos, null);
      }
  },[]);

    function indication() {
        return SuggestedTodosIndicationText;
    }

  return (
    <Container>
      <DapsHeader />
      <h1 className="text-center">{SuggestedTodosHeaderText}</h1>
      <BootstrapTable
        keyField='id'
        data={ todos }
        columns={ columns }
        noDataIndication={ indication }
        trStyle={rowTextColor}
        hover={true}
        striped={true}
        style={{display: "block", minHeight: "80%", width: "10%", overflow: "auto"}}
      />

    <Modal className='deleteActiveTodoModal text-center' show={showCannotDeleteActiveTodoModal} open={showCannotDeleteActiveTodoModal} centered={true} size='lg'>
        <ModalBody>
            <h4 style={{margin: "32px"}}>{CannotDeleteActiveTodoText}</h4>
            <ButtonGroup style={{width: "80%"}}>
                <Button
                    variant="danger"
                    onClick={() => toggleCannotDeleteStartedTodoModal()}
                    style={{margin: "auto", display: "block", padding: "0", textAlign: "center"}}
                >{CancelButtonText}</Button>
            </ButtonGroup>
        </ModalBody>
    </Modal>
    </Container>


  );
};

export default SuggestedTodosList;
