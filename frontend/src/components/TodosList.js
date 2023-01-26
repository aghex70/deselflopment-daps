import React, {useEffect, useState} from 'react';
import {
    faPencil,
    faPlay,
    faTrash,
    faCheck,
    faPlus,
    faArrowDown19,
    faArrowDown91,
} from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import TodoService from "../services/todo";
import {Button, ButtonGroup, Container, Modal, ModalBody} from "react-bootstrap";
import {useLocation, useNavigate} from "react-router-dom";
import './TodosList.css';
import BootstrapTable from 'react-bootstrap-table-next';
import DapsHeader from "./Header";
import checkAccess, {
    clearLocalStorage,
    sortTodosByField,
} from "../utils/helpers";
import {
    CancelButtonText,
    CompleteIconText,
    CreateIconText,
    DeleteButtonText,
    DeleteIconText,
    DeletingTodoText,
    EditIconText,
    HeaderActionsText,
    HeaderNameText,
    SortByDateButtonText,
    SortByPriorityButtonText,
    StartIconText
} from "../utils/texts";


const TodosList = () => {
  checkAccess();
  const [todos, setTodos] = useState([]);
  const [ascendingPriority, setAscendingPriority] = useState(false);
  const [ascendingPriorityIcon, setAscendingPriorityIcon] = useState(faArrowDown91);
  const [ascendingDate, setAscendingDate] = useState(true);
  const [ascendingDateIcon, setAscendingDateIcon] = useState(faArrowDown91);
  const [showDeleteTodoModal, setShowDeleteTodoModal] = useState(false);
  const [deleteId, setDeleteId] = useState("");
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

    const toggleConfirmDeleteTodoModal = (id) => {
        setDeleteId(id);
        setShowDeleteTodoModal(!showDeleteTodoModal);
    }


  const navigateToTodo = (id, categoryId, categoryName, action) => {
    clearLocalStorage([]);
    navigate("/todo/" + id, {state: {categoryId: location.state.categoryId, categoryName: location.state.categoryName, action: action}});
  }

  const deleteTodo = (id) => {
    TodoService.deleteTodo(id, categoryId).then(
      (response) => {
        if (response.status === 204) {
            clearLocalStorage([]);
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
            clearLocalStorage([]);
            window.location.reload();
        }
      }
    ).catch(
      (error) => {
      })
  }

  const startTodo = (id) => {
    TodoService.startTodo(id, categoryId).then(
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

  const createTodo = () => {
        clearLocalStorage([]);
        navigate("/create-todo", {state: {categoryId: location.state.categoryId, categoryName: location.state.categoryName}});
  }

  const sortByDate = () => {
      if (ascendingDate === true) {
          sortTodosByField("creation_date", true, setTodos, null);
          setAscendingDate(false)
          setAscendingDateIcon(faArrowDown91);
      } else {
          sortTodosByField("creation_date", false, setTodos, null);
          setAscendingDate(true)
          setAscendingDateIcon(faArrowDown19);
      }
  }

    const sortByPriority = () => {
        if (ascendingPriority === true) {
            sortTodosByField("priority", true, setTodos, null);
            setAscendingPriority(false);
            setAscendingPriorityIcon(faArrowDown91);
        } else {
            sortTodosByField("priority", false, setTodos, null);
            setAscendingPriority(true);
            setAscendingPriorityIcon(faArrowDown19);
        }
    }

  useEffect(() => {
    let todos = JSON.parse(localStorage.getItem("todos"));
    if (!todos) {
      TodoService.getTodos(categoryId).then(
        (response) => {
          if (response.status === 200 && response.data) {
            localStorage.setItem("todos", JSON.stringify(response.data));
            sortTodosByField("name", true, setTodos, setTodoSpan);
          }
        }
      ).catch(
        (error) => {
        })
    }
    else {
        sortTodosByField("name", true, setTodos, setTodoSpan);
    }
  },[categoryId]);

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
          <Button style={{width: "15%", margin: "auto", display: "block", padding: "0", textAlign: "center", }}
                  title={StartIconText} variant="outline-warning" onClick={() => startTodo(row.id)}>
            <FontAwesomeIcon icon={faPlay} />
          </Button>
        ) : (
          <Button style={{width: "15%", margin: "auto", display: "block", padding: "0", textAlign: "center"}}
                  title={CompleteIconText} variant="outline-success" onClick={() => completeTodo(row.id)}>
            <FontAwesomeIcon icon={faCheck} />
          </Button>
        )
        }

        <Button style={{width: "15%", margin: "auto", padding: "0", textAlign: "center"}}
                title={EditIconText} variant="outline-primary" onClick={() => navigateToTodo(row.id, categoryId, 0,"edit")}>
          <FontAwesomeIcon icon={faPencil} />
        </Button>

        <Button style={{width: "15%", margin: "auto", display: "block", padding: "0", textAlign: "center"}}
                title={DeleteIconText} variant="outline-danger" onClick={() => toggleConfirmDeleteTodoModal(row.id)}>
          <FontAwesomeIcon icon={faTrash} />
        </Button>
      </ButtonGroup>
      </div>
    );
  }

  function indication() {
    return <span className="createIcon" onClick={() => createTodo()}>
        <FontAwesomeIcon className="createIcon" icon={faPlus} />{CreateIconText}</span>
  }

      return (
        <Container>
          <DapsHeader />
          <h1 className="text-center">{location.state.categoryName}</h1>
          <span style={todoSpan} className="createIcon" onClick={() => createTodo()}>
          <FontAwesomeIcon className="createIcon" icon={faPlus} />{CreateIconText}</span>
            <ButtonGroup style={{width: "100%", marginTop: "10px"}}>
            <Button style={{width: "50%", margin: "auto", display: "block", padding: "0", textAlign: "center"}}
                    title={SortByPriorityButtonText}
                    onClick={() => sortByPriority()}
                    variant="primary">{SortByPriorityButtonText}
                <FontAwesomeIcon style={{marginLeft: "5px"}} icon={ascendingPriorityIcon} />
            </Button>
            <Button style={{width: "50%", margin: "auto", display: "block", padding: "0", textAlign: "center"}}
                    title={SortByDateButtonText}
                    onClick={() => sortByDate()}
                    variant="secondary">{SortByDateButtonText}
                <FontAwesomeIcon style={{marginLeft: "5px"}} icon={ascendingDateIcon} />
            </Button>
            </ButtonGroup>
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

        <Modal className='unshareModal text-center' show={showDeleteTodoModal} open={showDeleteTodoModal} centered={true} size='lg'>
            <ModalBody>
                <h4 style={{margin: "32px"}}>{DeletingTodoText}</h4>
                <ButtonGroup style={{width: "80%"}}>
                    <Button
                        variant="danger"
                        onClick={(e) => toggleConfirmDeleteTodoModal(e)}
                        style={{margin: "auto", display: "block", padding: "0", textAlign: "center"}}
                    >{CancelButtonText}</Button>
                    <Button
                        variant="success"
                        type="submit"
                        onClick={() => deleteTodo(deleteId)}
                        style={{margin: "auto", display: "block", padding: "0", textAlign: "center"}}
                    >{DeleteButtonText}</Button>
                </ButtonGroup>
            </ModalBody>
        </Modal>
        </Container>
        );
};

export default TodosList;
