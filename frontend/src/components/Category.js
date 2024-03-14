import React, { useEffect, useState } from "react";
import {
  Button,
  ButtonGroup,
  Container,
  FloatingLabel,
  Form,
} from "react-bootstrap";
import BootstrapTable from "react-bootstrap-table-next";
import { useLocation, useParams } from "react-router-dom";
import CategoryService from "../services/category";
import DapsHeader from "./Header";
import checkAccess, { checkValidToken, goToCategories } from "../utils/helpers";
import {
  CancelButtonText, DeleteIconText,
  DescriptionLabelText,
  EditButtonText,
  EditCategoryHeaderText, HeaderActionsText, HeaderUserText,
  NameLabelText,
  NoText,
  NotifiableLabelText,
  ViewCategoryHeaderText, ViewIconText,
  YesText,
} from "../utils/texts";
import toBoolean from "validator/es/lib/toBoolean";
import { faEye, faXmark } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { useNavigate } from "react-router-dom";


const Category = () => {
  checkAccess();
  const [categoryName, setCategoryName] = useState("");
  const [categoryDescription, setCategoryDescription] = useState("");
  const [categoryNotifiable, setCategoryNotifiable] = useState("false");
  const { id } = useParams();
  const location = useLocation();
  const enableEdit = location.state.action === "edit";

  const navigate = useNavigate();
  const [users, setUsers] = useState([]);

  const navigateToUser = (id) => {
    navigate("/user/" + id);
  };

  const columns = [
    {
      dataField: "email",
      text: HeaderUserText + "s",
      style: { width: "70%", cursor: "pointer", verticalAlign: "middle" },
    },
    {
      dataField: "link",
      text: HeaderActionsText,
      style: { width: "30%", verticalAlign: "middle" },
      formatter: actionsFormatter,
      headerAlign: "center",
    },
  ];

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
          {enableEdit && !row.is_owner &&
              <ButtonGroup style={{ width: "100%" }}>
            <Button
                style={{
                  width: "15%",
                  margin: "auto",
                  padding: "0",
                  textAlign: "center",
                }}
                title={ViewIconText}
                variant="outline-primary"
                onClick={() => navigateToUser(row.user_id)}
            >
                  <FontAwesomeIcon icon={faEye}/>
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
                onClick={() => unshareCategory(id, row.email)}
            >
              <FontAwesomeIcon icon={faXmark} />
            </Button>
          </ButtonGroup>
            }
        </div>
    );
  }

  const unshareCategory = (id, email) => {
    CategoryService.unshareCategory(id, email)
        .then((response) => {
            window.location.reload();
        })
        .catch((error) => {
          checkValidToken(error);
        });
  };

  const handleSubmit = (e) => {
    e.preventDefault();

    const data = {
      name: categoryName,
      description: categoryDescription,
      notifiable:
        typeof categoryNotifiable == "boolean"
          ? categoryNotifiable
          : toBoolean(categoryNotifiable),
    };

    CategoryService.updateCategory(id, data)
      .then((response) => {
        if (response.status === 200) {
          window.location.href = "/categories";
        } else {
          window.location.reload();
        }
      })
      .catch((error) => {
        checkValidToken(error);
      });
  };

  useEffect(() => {
    CategoryService.getCategory(id)
      .then((response) => {
        if (response.status === 200) {
          setCategoryName(response.data.name);
          setCategoryDescription(response.data.description);
          setCategoryNotifiable(response.data.notifiable);
        }
      })
      .catch((error) => {
        checkValidToken(error);
      });

    if (enableEdit && (!users || users.length === 0)) {
      CategoryService.getCategoryUsers(id)
          .then((response) => {
            if (response.status === 200) {
              setUsers(response.data.users);
            }
          })
          .catch((error) => {
            checkValidToken(error);
          });
    }

  }, [users, id]);

  return (
    <Container>
      <DapsHeader />
      <h1 className="text-center">
        {enableEdit ? EditCategoryHeaderText : ViewCategoryHeaderText}
      </h1>
      <Form onSubmit={(e) => handleSubmit(e)}>
        <FloatingLabel controlId="floatingName" label={NameLabelText}>
          <Form.Control
            type="name"
            value={categoryName}
            onChange={(e) => setCategoryName(e.target.value)}
            disabled={!enableEdit}
          />
        </FloatingLabel>

        <FloatingLabel
          controlId="floatingDescription"
          label={DescriptionLabelText}
        >
          <Form.Control
            as="textarea"
            placeholder="Description"
            style={{ height: "100px", margin: "0px 0px 32px" }}
            type="description"
            value={categoryDescription}
            onChange={(e) => setCategoryDescription(e.target.value)}
            disabled={!enableEdit}
          />
        </FloatingLabel>

        <FloatingLabel
          controlId="floatingNotifiable"
          label={NotifiableLabelText}
        >
          <Form.Select
            name="notifiable"
            value={categoryNotifiable}
            onChange={(e) => setCategoryNotifiable(e.target.value)}
            style={{ margin: "0px 0px 32px" }}
            disabled={!enableEdit}
          >
            <option value="false">{NoText}</option>
            <option value="true">{YesText}</option>
          </Form.Select>
        </FloatingLabel>
        {enableEdit && (
          <BootstrapTable
              keyField="user_id"
              data={users}
              columns={columns}
              hover={true}
              striped={true}
          />
        )}

        {enableEdit ? (
          <ButtonGroup
            style={{ width: "100%", paddingLeft: "10%", paddingRight: "10%" }}
          >
            <Button
              variant="danger"
              onClick={() => goToCategories()}
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
              style={{
                margin: "auto",
                display: "block",
                padding: "0",
                textAlign: "center",
              }}
            >
              {EditButtonText}
            </Button>
          </ButtonGroup>
        ) : (
          <Button
            variant="danger"
            onClick={() => goToCategories()}
            style={{
              margin: "auto",
              display: "block",
              padding: "0",
              textAlign: "center",
            }}
          >
            {CancelButtonText}
          </Button>
        )}
      </Form>
    </Container>
  );
};
export default Category;
