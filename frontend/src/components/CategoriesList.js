import React, { useEffect, useState } from "react";
import {
  faPencil,
  faPlus,
  faShareNodes,
  faTrash,
  faEye,
} from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import CategoryService from "../services/category";
import {
  Button,
  ButtonGroup,
  Container,
  FloatingLabel,
  Form,
  Modal,
  ModalBody,
} from "react-bootstrap";
import { useNavigate } from "react-router-dom";
import "./CategoriesList.css";
import BootstrapTable from "react-bootstrap-table-next";
import DapsHeader from "./Header";
import checkAccess, {
  checkValidToken,
  clearLocalStorage,
  goToCreateCategory,
  sortCategoriesByField,
  getUserId,
} from "../utils/helpers";
import {
  CancelButtonText,
  CategoriesHeaderText,
  ConfirmUnshareCategoryText,
  CreateCategoryIconText,
  DeleteButtonText,
  DeleteIconText,
  DeletingCategoryText,
  EditIconText,
  HeaderActionsText,
  HeaderNameText,
  HeaderTasksText,
  HighPriorityTasksText,
  language,
  OnlyOwnerCanEditCategoryText,
  OnlyOwnersCanDeleteCategoryText,
  ShareButtonText,
  ShareCategoryHeaderText,
  ShareIconText,
  TotalNumberOfTasksText,
  UnSubscribeIconText,
  UnsuscribeButtonText,
  UserAlreadySubscribedText,
  ViewIconText,
} from "../utils/texts";

const CategoriesList = () => {
  checkAccess();

  const [categories, setCategories] = useState([]);
  const [showModal, setShowModal] = useState(false);
  const [showModalUserAlreadySubscribed, setShowModalUserAlreadySubscribed] =
    useState(false);
  const [showModalCannotDeleteCategory, setShowModalCannotDeleteCategory] =
    useState(false);
  const [showModalCannotEditCategory, setShowModalCannotEditCategory] =
    useState(false);
  const [showUnshareModal, setUnshareShowModal] = useState(false);
  const [showDeleteCategoryModal, setShowDeleteCategoryModal] = useState(false);
  const [deleteId, setDeleteId] = useState("");
  const [shareId, setShareId] = useState("");
  const [unshareId, setUnshareId] = useState("");
  const [shareEmail, setShareEmail] = useState("");
  const [categorySpan, setCategorySpan] = useState({
    textAlign: "center",
    display: "none",
  });
  const navigate = useNavigate();

  const sharedTag = (shared) => {
    if (!shared) {
      return "";
    }
    return language === "en" ? "(s) " : "(c) ";
  };

  const userId = parseInt(getUserId());
  // Color code the todo based on its priority
  const rowTextColor = (cell, row) => {
    return (
      <div onClick={() => navigateToCategory(row.id, row.name)}>
        {sharedTag(row.shared)}
        {row.name}
      </div>
    );
  };
  // Color code the todo based on its priority
  const rowTasksFormatter = (cell, row) => {
    return (
      <p
        style={{
          fontWeight: "bold",
          color: "red",
          cursor: "pointer",
          margin: "0",
        }}
        title={HighPriorityTasksText}
        onClick={() => navigateToCategory(row.id, row.name)}
      >
        {row.highest_priority_tasks}
        <span
          title={TotalNumberOfTasksText}
          style={{ fontWeight: "bold", color: "black" }}
        >
          /{row.tasks}
        </span>
      </p>
    );
  };

  const columns = [
    {
      dataField: "tasks",
      text: HeaderTasksText,
      style: {
        width: "15%",
        cursor: "pointer",
        verticalAlign: "middle",
        justifyContent: "center",
      },
      formatter: rowTasksFormatter,
    },
    {
      dataField: "name",
      text: HeaderNameText,
      style: { width: "55%", cursor: "pointer", verticalAlign: "middle" },
      formatter: rowTextColor,
    },
    {
      dataField: "link",
      text: HeaderActionsText,
      style: { width: "30%", verticalAlign: "middle" },
      formatter: actionsFormatter,
      headerAlign: "center",
    },
  ];

  const navigateToCategory = (categoryId, categoryName) => {
    clearLocalStorage([]);
    navigate("/todos", {
      state: { categoryId: categoryId, categoryName: categoryName },
    });
  };

  const getCategory = (id, action, categoryName) => {
    navigate("/category/" + id, {
      state: { action: action, categoryName: categoryName },
    });
  };

  const confirmDeleteCategory = () => {
    CategoryService.deleteCategory(deleteId)
      .then((response) => {
        if (response.status === 204) {
          window.location.reload();
        }
      })
      .catch((error) => {
        checkValidToken(error);
        if (error.response.data.message === "cannot remove category") {
          setShowModalCannotDeleteCategory(true);
        }
      });
  };

  const toggleConfirmDeleteCategoryModal = (id) => {
    setDeleteId(id);
    setShowDeleteCategoryModal(!showDeleteCategoryModal);
  };

  const toggleModal = () => {
    setShowModal(!showModal);
  };

  const toggleUserAlreadySubscribedModal = () => {
    setShowModalUserAlreadySubscribed(!showModalUserAlreadySubscribed);
  };

  const toggleCannotDeleteCategoryModal = () => {
    setShowModalCannotDeleteCategory(!setShowModalCannotDeleteCategory);
  };

  const toggleCannotEditCategoryModal = () => {
    setShowModalCannotEditCategory(!setShowModalCannotEditCategory);
  };

  const toggleUnshareModal = () => {
    setUnshareShowModal(!showUnshareModal);
  };

  const shareCategory = (id) => {
    setShareId(id);
    setShowModal(true);
  };

  const unshareCategory = (id) => {
    setUnshareId(id);
    setUnshareShowModal(true);
  };

  const confirmUnshareCategory = () => {
    CategoryService.unshareCategory(unshareId)
      .then((response) => {
        if (response.status === 200) {
          window.location.reload();
        }
      })
      .catch((error) => {
        checkValidToken(error);
      });
  };

  const confirmShareCategory = () => {
    CategoryService.shareCategory(shareId, shareEmail)
      .then((response) => {
        if (response.status === 200) {
          setShowModal(false);
        }
      })
      .catch((error) => {
        checkValidToken(error);
        if (
          error.response.data.message ===
          "user already subscribed to that category"
        ) {
          setShowModal(false);
          setShowModalUserAlreadySubscribed(true);
        }
        setShowModal(false);
      });
  };

  useEffect(() => {
    if (!categories || categories.length === 0) {
      CategoryService.getCategories()
        .then((response) => {
          if (response.status === 200 && response.data) {
            sortCategoriesByField(
              response.data,
              "name",
              true,
              setCategories,
              setCategorySpan
            );
          }
        })
        .catch((error) => {
          checkValidToken(error);
        });
    }
  }, [categories]);

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
        <ButtonGroup style={{ width: "100%" }}>
          {isOwner(row.owner_id) ? (
            <Button
              style={{
                width: "15%",
                margin: "auto",
                display: "block",
                padding: "0",
                textAlign: "center",
              }}
              title={ShareIconText}
              variant="outline-success"
              onClick={() => shareCategory(row.id)}
            >
              <FontAwesomeIcon icon={faShareNodes} />
            </Button>
          ) : (
            <Button
              style={{
                width: "15%",
                margin: "auto",
                display: "block",
                padding: "0",
                textAlign: "center",
              }}
              title={UnSubscribeIconText}
              variant="outline-dark"
              onClick={() => unshareCategory(row.id)}
            >
              <FontAwesomeIcon
                style={{ rotate: "180deg" }}
                icon={faShareNodes}
              />
            </Button>
          )}

          {isOwner(row.owner_id) ? (
            <Button
              style={{
                width: "15%",
                margin: "auto",
                padding: "0",
                textAlign: "center",
              }}
              title={EditIconText}
              variant="outline-primary"
              onClick={() => getCategory(row.id, "edit", row.name)}
            >
              <FontAwesomeIcon icon={faPencil} />
            </Button>
          ) : (
            <Button
              style={{
                width: "15%",
                margin: "auto",
                padding: "0",
                textAlign: "center",
              }}
              title={ViewIconText}
              variant="outline-primary"
              onClick={() => getCategory(row.id, "view", row.name)}
            >
              <FontAwesomeIcon icon={faEye} />
            </Button>
          )}

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
            onClick={() => toggleConfirmDeleteCategoryModal(row.id)}
          >
            <FontAwesomeIcon icon={faTrash} />
          </Button>
        </ButtonGroup>
      </div>
    );
  }

  function indication() {
    return (
      <span className="createIcon" onClick={() => goToCreateCategory()}>
        <FontAwesomeIcon className="createIcon" icon={faPlus} />
        {CreateCategoryIconText}
      </span>
    );
  }

  return (
    <Container>
      <DapsHeader />
      <h1 className="text-center">{CategoriesHeaderText}</h1>
      <span
        style={categorySpan}
        className="createIcon"
        onClick={() => goToCreateCategory()}
      >
        <FontAwesomeIcon className="createIcon" icon={faPlus} />
        {CreateCategoryIconText}
      </span>
      <BootstrapTable
        keyField="id"
        data={categories}
        columns={columns}
        noDataIndication={indication}
        trStyle={rowTextColor}
        hover={true}
        striped={true}
      />
      <Modal
        className="successModal text-center"
        show={showModal}
        open={showModal}
        centered={true}
        size="lg"
      >
        <ModalBody>
          <h4 style={{ margin: "32px" }}>{ShareCategoryHeaderText}</h4>
          <Form onSubmit={(e) => confirmShareCategory(e)}>
            <Form.Group controlId="formCategoryName">
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
          <ButtonGroup style={{ width: "80%" }}>
            <Button
              variant="danger"
              onClick={(e) => toggleModal(e)}
              style={{
                margin: "auto",
                display: "block",
                padding: "0",
                textAlign: "center",
              }}
            >
              {CancelButtonText}
            </Button>
            <Button
              variant="success"
              type="submit"
              onClick={(e) => confirmShareCategory(e)}
              style={{
                margin: "auto",
                display: "block",
                padding: "0",
                textAlign: "center",
              }}
            >
              {ShareButtonText}
            </Button>
          </ButtonGroup>
        </ModalBody>
      </Modal>

      <Modal
        className="successModal text-center"
        show={showModalUserAlreadySubscribed}
        open={showModalUserAlreadySubscribed}
        centered={true}
        size="lg"
      >
        <ModalBody>
          <h4 style={{ margin: "32px" }}>{UserAlreadySubscribedText}</h4>
          <ButtonGroup style={{ width: "40%" }}>
            <Button
              variant="danger"
              onClick={(e) => toggleUserAlreadySubscribedModal(e)}
              style={{
                margin: "auto",
                display: "block",
                padding: "0",
                textAlign: "center",
              }}
            >
              {CancelButtonText}
            </Button>
          </ButtonGroup>
        </ModalBody>
      </Modal>

      <Modal
        className="successModal text-center"
        show={showModalCannotDeleteCategory}
        open={showModalCannotDeleteCategory}
        centered={true}
        size="lg"
      >
        <ModalBody>
          <h4 style={{ margin: "32px" }}>{OnlyOwnersCanDeleteCategoryText}</h4>
          <ButtonGroup style={{ width: "40%" }}>
            <Button
              variant="danger"
              onClick={(e) => toggleCannotDeleteCategoryModal(e)}
              style={{
                margin: "auto",
                display: "block",
                padding: "0",
                textAlign: "center",
              }}
            >
              {CancelButtonText}
            </Button>
          </ButtonGroup>
        </ModalBody>
      </Modal>

      <Modal
        className="successModal text-center"
        show={showModalCannotEditCategory}
        open={showModalCannotEditCategory}
        centered={true}
        size="lg"
      >
        <ModalBody>
          <h4 style={{ margin: "32px" }}>{OnlyOwnerCanEditCategoryText}</h4>
          <ButtonGroup style={{ width: "40%" }}>
            <Button
              variant="danger"
              onClick={(e) => toggleCannotEditCategoryModal(e)}
              style={{
                margin: "auto",
                display: "block",
                padding: "0",
                textAlign: "center",
              }}
            >
              {CancelButtonText}
            </Button>
          </ButtonGroup>
        </ModalBody>
      </Modal>

      <Modal
        className="unshareModal text-center"
        show={showUnshareModal}
        open={showUnshareModal}
        centered={true}
        size="lg"
      >
        <ModalBody>
          <h4 style={{ margin: "32px" }}>{ConfirmUnshareCategoryText}</h4>
          <ButtonGroup style={{ width: "80%" }}>
            <Button
              variant="danger"
              onClick={(e) => toggleUnshareModal(e)}
              style={{
                margin: "auto",
                display: "block",
                padding: "0",
                textAlign: "center",
              }}
            >
              {CancelButtonText}
            </Button>
            <Button
              variant="success"
              type="submit"
              onClick={(e) => confirmUnshareCategory(e)}
              style={{
                margin: "auto",
                display: "block",
                padding: "0",
                textAlign: "center",
              }}
            >
              {UnsuscribeButtonText}
            </Button>
          </ButtonGroup>
        </ModalBody>
      </Modal>

      <Modal
        className="unshareModal text-center"
        show={showDeleteCategoryModal}
        open={showDeleteCategoryModal}
        centered={true}
        size="lg"
      >
        <ModalBody>
          <h4 style={{ margin: "32px" }}>{DeletingCategoryText}</h4>
          <ButtonGroup style={{ width: "80%" }}>
            <Button
              variant="danger"
              onClick={(e) => toggleConfirmDeleteCategoryModal(e)}
              style={{
                margin: "auto",
                display: "block",
                padding: "0",
                textAlign: "center",
              }}
            >
              {CancelButtonText}
            </Button>
            <Button
              variant="success"
              type="submit"
              onClick={(e) => confirmDeleteCategory(e)}
              style={{
                margin: "auto",
                display: "block",
                padding: "0",
                textAlign: "center",
              }}
            >
              {DeleteButtonText}
            </Button>
          </ButtonGroup>
        </ModalBody>
      </Modal>
    </Container>
  );
};

export default CategoriesList;
