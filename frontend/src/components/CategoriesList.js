import React, {useEffect, useState} from 'react';
import {
  faPencil, faPlus,
  faShareNodes,
  faTrash
} from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import CategoryService from "../services/category";
import {Button, ButtonGroup, Container, FloatingLabel, Form, Modal, ModalBody} from "react-bootstrap";
import {useNavigate} from "react-router-dom";
import './CategoriesList.css';
import BootstrapTable from "react-bootstrap-table-next";
import DapsHeader from "./Header";

const CategoriesList = () => {
  const [categories, setCategories] = useState([]);
  const [showModal, setShowModal] = useState(false);
  const [showUnshareModal, setUnshareShowModal] = useState(false);
  const [shareId, setShareId] = useState("");
  const [unshareId, setUnshareId] = useState("");
  const [shareEmail, setShareEmail] = useState("");
  const [categorySpan, setCategorySpan] = useState({
    textAlign: "center",
    display: "none",
  });
  const navigate = useNavigate();

  const userId = parseInt(localStorage.getItem("user_id"))
  // Color code the todo based on its priority
  const rowTextColor = (cell, row, rowIndex) => {
    const colors = ["red", "grey", "blue", "green", "orange"];
    return <div
      style={{color : colors[rowIndex % 5]}}
      onClick={() => navigateToCategory(row.id, row.name)}
    >
      {row.name}
    </div>;
  }
  // Color code the todo based on its priority
  const rowTasksFormatter = (cell, row) => {
    return <p
      style={{fontWeight: "bold" , color: "red", cursor: "pointer", margin: "0"}}
      title="Highest priority tasks"
      onClick={() => navigateToCategory(row.id, row.name)}>
      {row.highest_priority_tasks}
      <span title="Total number of tasks" style={{fontWeight: "bold", color: "black"}}>/{row.tasks}</span>
    </p>;
  }

  const columns = [
    {
      dataField: 'tasks',
      text: 'Tasks',
      style:{'width' : '15%', cursor: "pointer", verticalAlign: "middle", justifyContent: "center"},
      formatter: rowTasksFormatter,
    },
    {
      dataField: 'name',
      text: 'Category',
      style:{'width' : '55%', cursor: "pointer", verticalAlign: "middle"},
      formatter: rowTextColor,
    },
    {
      dataField: 'link',
      text: 'Actions',
      style:{'width' : '30%', verticalAlign: "middle"},
      formatter: actionsFormatter,
      headerAlign: 'center',
    }];

  if (!localStorage.getItem("access_token")) {
    window.location.href = "/login";
  }

  const navigateToCategory = (categoryId, categoryName) => {
    navigate("/todos", {state: {categoryId: categoryId, categoryName: categoryName}});
  }

  const navigateToCreateCategory = () => {
    navigate("/create-category");
  }

  const getCategory = (id) => {
    navigate("/category/" + id);
  }

  const deleteCategory = (id) => {
    CategoryService.deleteCategory(id).then(
      (response) => {
        console.log(response);
        if (response.status === 204) {
          window.location.reload();
        }
      }
    ).catch(
      (error) => {
        console.log(error);
        error = new Error("Deletion failed!");
      })
  }

  const toggleModal = () => {
    setShowModal(!showModal);
  }

  const toggleUnshareModal = () => {
    setUnshareShowModal(!showUnshareModal);
  }

  const shareCategory = (id) => {
    console.log("shareId: " + id);
    setShareId(id);
    setShowModal(true);
  }

  const unshareCategory = (id) => {
    console.log("shareId: " + id);
    setUnshareId(id);
    setUnshareShowModal(true);
  }

  const confirmUnshareCategory = () => {
    CategoryService.unshareCategory(unshareId).then(
      (response) => {
        console.log(response);
        if (response.status === 200) {
          window.location.reload();
        }
      }
    ).catch(
      (error) => {
        console.log(error);
        error = new Error("Unsharing failed!");
      })
  }

  const confirmShareCategory = () => {
    CategoryService.shareCategory(shareId, shareEmail).then(
      (response) => {
        console.log(response);
        if (response.status === 200) {
          setShowModal(false);
        }
      }
    ).catch(
      (error) => {
        console.log(error);
        console.log("AGPAGPAPGAPGPG");
        error = new Error("Sharing failed!");
        setShowModal(false);
      })
  }

  useEffect(() => {
    console.log("hola")
    if (!categories || categories.length === 0) {
      CategoryService.getCategories().then(
        (response) => {
          console.log(response);
          if (response.status === 200 && response.data) {
            console.log("response.data -----------> " + response.data);
            setCategories(response.data);
            setCategorySpan({
              textAlign: "center",
              display: "block",
            }
            );
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

  function isOwner(rowOwnerId) {
    return rowOwnerId === userId;
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
          {isOwner(row.owner_id)? (
            <Button style={{width: "15%", margin: "auto", display: "block", padding: "0", textAlign: "center"}}
                    title="Share"
                    variant="outline-success" onClick={() => shareCategory(row.id)}>
              <FontAwesomeIcon icon={faShareNodes} />
            </Button>
          ) : (
            <Button style={{width: "15%", margin: "auto", display: "block", padding: "0", textAlign: "center"}}
                    title="Unsubscribe"
                    variant="outline-dark" onClick={() => unshareCategory(row.id)}>
              <FontAwesomeIcon style={{rotate: "180deg"}} icon={faShareNodes}/>
            </Button>
          )}

          <Button style={{width: "15%", margin: "auto", padding: "0", textAlign: "center"}}
                  title="Edit"
                  variant="outline-primary" onClick={() => getCategory(row.id, row.name)}>
            <FontAwesomeIcon icon={faPencil} />
          </Button>
          <Button style={{width: "15%", margin: "auto", display: "block", padding: "0", textAlign: "center"}}
                  title="Delete"
                  variant="outline-danger" onClick={() => deleteCategory(row.id)}>
            <FontAwesomeIcon icon={faTrash} />
          </Button>
        </ButtonGroup>
      </div>
    );
  }

  function indication() {
    return <span className="createIcon" onClick={() => navigateToCreateCategory()}>
      <FontAwesomeIcon className="createIcon" icon={faPlus} />Create a new Category</span>
  }

  console.log("CATEGORIES: " + categories, typeof categories);
  return (
    <Container>
      <DapsHeader />
      <h1 className="text-center">Categories</h1>
      <span style={categorySpan} className="createIcon" onClick={() => navigateToCreateCategory()}>
      <FontAwesomeIcon className="createIcon" icon={faPlus} />Create a new Category</span>
      <BootstrapTable
        keyField='id'
        data={ categories }
        columns={ columns }
        noDataIndication={ indication }
        trStyle={rowTextColor}
        hover={true}
      />
      <Modal className='successModal text-center' show={showModal} open={showModal} centered={true} size='lg'>
        <ModalBody>
          <div>
            <div className='container my-4'>
            </div>
          </div>
          <Form  onSubmit={(e) => confirmShareCategory(e)}>
            <Form.Group controlId="formCategoryName">
              <h3 style={{margin: "32px"}}>Share category</h3>
              <FloatingLabel
                controlId="floatingEmail"
                label="Email"
                value={shareEmail}
                onChange={(e) => setShareEmail(e.target.value)}
                placeholder="Email"
              >
                <Form.Control type="email" placeholder="Email" />
              </FloatingLabel>
            </Form.Group>
          </Form>
          <ButtonGroup style={{width: "80%"}}>
          <Button
            variant="success"
            type="submit"
            onClick={(e) => confirmShareCategory(e)}
            style={{margin: "auto", display: "block", padding: "0", textAlign: "center"}}
          >Share</Button>
          <Button
            variant="danger"
            onClick={(e) => toggleModal(e)}
            style={{margin: "auto", display: "block", padding: "0", textAlign: "center"}}
          >Cancel</Button>
          </ButtonGroup>
        </ModalBody>
      </Modal>

      <Modal className='unshareModal text-center' show={showUnshareModal} open={showUnshareModal} centered={true} size='lg'>
        <ModalBody>
          <div>
            Are you sure you want to unsubscribe from this category? This operation cannot be undone.
            <div className='container my-4'>
            </div>
          </div>
          <ButtonGroup style={{width: "80%"}}>
            <Button
              variant="success"
              type="submit"
              onClick={(e) => confirmUnshareCategory(e)}
              style={{margin: "auto", display: "block", padding: "0", textAlign: "center"}}
            >Unsubscribe</Button>
            <Button
              variant="danger"
              onClick={(e) => toggleUnshareModal(e)}
              style={{margin: "auto", display: "block", padding: "0", textAlign: "center"}}
            >Cancel</Button>
          </ButtonGroup>
        </ModalBody>
      </Modal>
    </Container>
  );
};

export default CategoriesList;
