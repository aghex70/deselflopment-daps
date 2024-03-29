import React, { useEffect, useState } from "react";
import {
  Button,
  ButtonGroup,
  Container,
  FloatingLabel,
  Form,
  Nav,
} from "react-bootstrap";
import { useLocation, useNavigate, useParams } from "react-router-dom";
import TodoService from "../services/todo";
import DapsHeader from "./Header";
import checkAccess, {
  checkValidToken,
  clearLocalStorage,
  goToCategories,
} from "../utils/helpers";
import toBoolean from "validator/es/lib/toBoolean";
import {
  BiweeklyText,
  CancelButtonText,
  DailyText,
  DescriptionLabelText,
  EditButtonText,
  EditTodoHeaderText,
  HighestPriorityText,
  HighPriorityText,
  LinkLabelText,
  LowestPriorityText,
  LowPriorityText,
  MediumPriorityText,
  MonthlyText,
  NameLabelText,
  NoText,
  OpenLinkText,
  PriorityLabelText,
  RecurrencyLabelText,
  RecurringLabelText,
  SuggestableLabelText,
  ViewTodoHeaderText,
  WeekdaysText,
  WeekendsText,
  WeeklyText,
  YesText,
} from "../utils/texts";

const Todo = () => {
  checkAccess();
  const [todoName, setTodoName] = useState("");
  const [todoDescription, setTodoDescription] = useState("");
  const [todoLink, setTodoLink] = useState("");
  const [todoPriority, setTodoPriority] = useState("");
  const [todoRecurring, setTodoRecurring] = useState("");
  const [todoSuggestable, setTodoSuggestable] = useState("");
  const [todoRecurrencyPeriod, setTodoRecurrencyPeriod] = useState("");
  const [todoCategoryId, setTodoCategoryId] = useState();
  const [, setTodoCategoryName] = useState();
  const location = useLocation();
  const categoryId = location.state.categoryId;
  const categoryName = location.state.categoryName;
  const enableEdit = location.state.action === "edit";
  const { id } = useParams();
  const navigate = useNavigate();

  const navigateTodos = () => {
    clearLocalStorage([]);
    navigate("/todos", {
      state: {
        categoryId: location.state.categoryId,
        categoryName: location.state.categoryName,
      },
    });
  };

  const mapRecurrencyPeriod = () => {
    if (todoRecurring === "true" || todoRecurring === true) {
      if (todoRecurrencyPeriod.length === 0) {
        return "daily";
      }
      return todoRecurrencyPeriod;
    }
    return "";
  };

  const handleSubmit = (e) => {
    e.preventDefault();

    const data = {
      name: todoName,
      description: todoDescription,
      link: todoLink,
      priority: parseInt(todoPriority),
      recurring:
        typeof todoRecurring == "boolean"
          ? todoRecurring
          : toBoolean(todoRecurring),
      recurrency: mapRecurrencyPeriod(),
      suggestable:
        typeof todoSuggestable == "boolean"
          ? todoSuggestable
          : toBoolean(todoSuggestable),
    };

    TodoService.updateTodo(id, data)
      .then((response) => {
        if (response.status === 200) {
          clearLocalStorage([]);
          categoryName === "" || categoryName === undefined
            ? goToCategories()
            : navigateTodos(categoryId, categoryName);
        } else {
          window.location.reload();
        }
      })
      .catch((error) => {
        clearLocalStorage([]);
        checkValidToken(error);
        categoryName === "" || categoryName === undefined
          ? goToCategories()
          : navigateTodos(categoryId, categoryName);
      });
  };

  useEffect(() => {
    TodoService.getTodo(id)
      .then((response) => {
        if (response.status === 200) {
          setTodoName(response.data.name);
          setTodoDescription(response.data.description);
          setTodoLink(response.data.link);
          setTodoPriority(response.data.priority);
          setTodoRecurring(response.data.recurring);
          setTodoCategoryId(response.data.category_id);
          setTodoCategoryName(response.data.category_name);
          setTodoRecurrencyPeriod(response.data.recurrency);
          setTodoSuggestable(response.data.suggestable);
        }
      })
      .catch((error) => {
        checkValidToken(error);
      });
  }, [id, categoryId]);

  return (
    <Container>
      <DapsHeader />
      <h1 className="text-center">
        {enableEdit ? EditTodoHeaderText : ViewTodoHeaderText}
      </h1>
      <Form onSubmit={(e) => handleSubmit(e)}>
        <FloatingLabel
          controlId="floatingName"
          label={NameLabelText}
          value={todoName}
          onChange={(e) => setTodoName(e.target.value)}
        >
          <Form.Control
            type="name"
            placeholder="Name"
            value={todoName}
            disabled={!enableEdit}
          />
        </FloatingLabel>

        {!enableEdit && todoDescription && (
          <FloatingLabel
            controlId="floatingDescription"
            label={DescriptionLabelText}
            value={todoDescription}
            onChange={(e) => setTodoDescription(e.target.value)}
          >
            <Form.Control
              type="description"
              placeholder="Description"
              value={todoDescription}
              disabled={!enableEdit}
            />
          </FloatingLabel>
        )}

        <FloatingLabel controlId="floatingPriority" label={PriorityLabelText}>
          <Form.Select
            name="priority"
            value={todoPriority}
            onChange={(e) => setTodoPriority(e.target.value)}
            style={{ margin: "0px 0px 32px" }}
            disabled={!enableEdit}
          >
            <option style={{ color: "grey" }} value="1">
              {LowestPriorityText}
            </option>
            <option style={{ color: "blue" }} value="2">
              {LowPriorityText}
            </option>
            <option style={{ color: "green" }} value="3">
              {MediumPriorityText}
            </option>
            <option style={{ color: "orange" }} value="4">
              {HighPriorityText}
            </option>
            <option style={{ color: "red" }} value="5">
              {HighestPriorityText}
            </option>
          </Form.Select>
        </FloatingLabel>

        <FloatingLabel controlId="floatingRecurring" label={RecurringLabelText}>
          <Form.Select
            name="recurring"
            value={todoRecurring}
            onChange={(e) => setTodoRecurring(e.target.value)}
            style={{ margin: "0px 0px 32px" }}
            disabled={!enableEdit}
          >
            <option value="false">{NoText}</option>
            <option value="true">{YesText}</option>
          </Form.Select>
        </FloatingLabel>

        <FloatingLabel
          controlId="floatingSuggestable"
          label={SuggestableLabelText}
        >
          <Form.Select
            name="suggestable"
            value={todoSuggestable}
            onChange={(e) => setTodoSuggestable(e.target.value)}
            style={{ margin: "0px 0px 32px" }}
            disabled={!enableEdit}
          >
            <option value="false">{NoText}</option>
            <option value="true">{YesText}</option>
          </Form.Select>
        </FloatingLabel>

        <FloatingLabel
          controlId="floatingRecurringPeriod"
          style={{
            display:
              todoRecurring === "false" || todoRecurring === false
                ? "none"
                : "block",
          }}
          label={RecurrencyLabelText}
        >
          <Form.Select
            name="recurring"
            value={todoRecurrencyPeriod}
            onChange={(e) => setTodoRecurrencyPeriod(e.target.value)}
            style={{ margin: "0px 0px 32px" }}
            disabled={!enableEdit}
          >
            <option value="daily">{DailyText}</option>
            <option value="weekly">{WeeklyText}</option>
            <option value="biweekly">{BiweeklyText}</option>
            <option value="monthly">{MonthlyText}</option>
            <option value="weekdays">{WeekdaysText}</option>
            <option value="weekends">{WeekendsText}</option>
          </Form.Select>
        </FloatingLabel>

        {enableEdit && (
          <FloatingLabel
            controlId="floatingLink"
            label={LinkLabelText}
            value={todoLink}
            onChange={(e) => {
              const link = e.target.value;
              if (link === "") {
                setTodoLink("");
              } else if (!/^https?:\/\//i.test(link)) {
                // checks if http:// or https:// is already present
                setTodoLink("https://" + link); // prepends https:// if it's not present
              } else {
                setTodoLink(link); // otherwise, set the link as is
              }
            }}
          >
            <Form.Control
              type="link"
              placeholder="Link"
              value={todoLink}
              disabled={!enableEdit}
            />
          </FloatingLabel>
        )}

        {todoLink && (
          <Nav
            className="justify-content-center"
            style={{ marginBottom: "15px" }}
            activeKey={todoLink}
          >
            <Nav.Item className="font-size-lg">
              <Nav.Link href={todoLink} target="_blank">
                {OpenLinkText}
              </Nav.Link>
            </Nav.Item>
          </Nav>
        )}

        {enableEdit ? (
          <ButtonGroup
            style={{ width: "100%", paddingLeft: "10%", paddingRight: "10%" }}
          >
            <Button
              variant="danger"
              onClick={() => navigateTodos()}
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
          <ButtonGroup style={{ width: "40%", marginLeft: "30%" }}>
            <Button
              variant="danger"
              onClick={() => navigateTodos()}
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
        )}
      </Form>
    </Container>
  );
};
export default Todo;
