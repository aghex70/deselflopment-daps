import React, {useState} from 'react';
import {Button, ButtonGroup, Container, FloatingLabel, Form, Modal, ModalBody} from "react-bootstrap";
import {useLocation, useNavigate} from 'react-router-dom'
import TodoService from "../services/todo";
import toBoolean from "validator/es/lib/toBoolean";
import DapsHeader from "./Header";
import checkAccess, {clearLocalStorage} from "../utils/helpers";
import {
    BiweeklyText,
    CancelButtonText,
    CreateButtonText,
    CreateTodoHeaderText,
    DailyText,
    DescriptionLabelText,
    HighestPriorityText,
    HighPriorityText,
    LinkLabelText,
    LowestPriorityText,
    LowPriorityText,
    MediumPriorityText,
    MonthlyText,
    NameLabelText,
    NoText,
    PleaseEnterTodoNameText,
    PriorityLabelText,
    RecurrencyLabelText,
    RecurringLabelText,
    SelectPriorityText,
    SelectRecurringText,
    WeekdaysText,
    WeekendsText,
    WeeklyText,
    YesText
} from "../utils/texts";

const CreateTodo = () => {
    checkAccess();
    const [todoName, setTodoName] = useState("");
    const [todoDescription, setTodoDescription] = useState("");
    const [todoLink, setTodoLink] = useState("");
    const [todoPriority, setTodoPriority] = useState("3");
    const [todoRecurring, setTodoRecurring] = useState("false");
    const [todoRecurrencyPeriod, setTodoRecurrencyPeriod] = useState("daily");
    const [disablePriority, setDisablePriority] = useState(false);
    const [disableRecurring, setDisableRecurring] = useState(false);
    const [showEnterTodoNameModal, setShowEnterTodoNameModal] = useState(false);
    const location = useLocation();
    const categoryId = location.state.categoryId;
    const categoryName = location.state.categoryName;
    const navigate = useNavigate();

    const disablePrioritySelect = () => {
      if (!disablePriority) {
        if (todoPriority in ["1", "2", "3", "4", "5"]) {
          setDisablePriority(true);
        }
      }
    }
    disablePrioritySelect();

    const disableRecurringSelect = () => {
      if (!disableRecurring) {
        if (todoRecurring === "true" || todoRecurring === "false") {
          setDisableRecurring(true);
        }
      }
    }
    disableRecurringSelect();

    const navigateTodos = () => {
        clearLocalStorage([]);
        navigate("/todos", {state: {categoryId: location.state.categoryId, categoryName: location.state.categoryName}});
    }

    const toggleEnterTodoNameModal = () => {
        setShowEnterTodoNameModal(!showEnterTodoNameModal);
    }

    const mapRecurrencyPeriod = () => {
        if (todoRecurring === "true" || todoRecurring === true) {
            if (todoRecurrencyPeriod.length === 0) {
                return "daily";
            }
            return todoRecurrencyPeriod;
        }
        return "";
    }

    const handleSubmit = (e) => {
      e.preventDefault();

      if (todoName === "") {
        toggleEnterTodoNameModal();
        return;
      }

      const data = {
        name: todoName,
        description: todoDescription,
        link: todoLink,
        // priority will be casted to int
        priority: typeof(todoPriority) === "number" ? todoPriority : parseInt(todoPriority),
        recurring: toBoolean(todoRecurring),
        category_id: categoryId,
        recurrency: mapRecurrencyPeriod(),
      }

      TodoService.createTodo(data).then(
        (response) => {
          if (response.status === 201) {
                clearLocalStorage([]);
                navigateTodos(categoryId, categoryName);
          } else {
            window.location.reload()
          }
        }
      ).catch(
        (error) => {
          error = new Error("Update todo failed!");
        }
      )
    }

    return (
      <Container>
        <DapsHeader />
        <h1 style={{ margin: '0px 0px 32px' }} className="text-center">{CreateTodoHeaderText}</h1>
        <Form onSubmit={(e) => handleSubmit(e)}>
          <FloatingLabel
            controlId="floatingName"
            label={NameLabelText}
            value={todoName}
            onChange={(e) => setTodoName(e.target.value)}
          >
            <Form.Control type="name" placeholder="Name" />
          </FloatingLabel>

          <FloatingLabel controlId="floatingDescription" label={DescriptionLabelText}>
            <Form.Control
              as="textarea"
              placeholder="Description"
              style={{ height: '100px', margin: '0px 0px 32px' }}
              type="description"
              value={todoDescription}
              onChange={(e) => setTodoDescription(e.target.value)}/>
          </FloatingLabel>


            <FloatingLabel controlId="floatingPriority" label={PriorityLabelText}>
                <Form.Select
              name="priority"
              value={todoPriority}
              onChange={(e) => setTodoPriority(e.target.value)}
              style={{ margin: '0px 0px 32px' }}
            >
                    <option disabled={disablePriority}>{SelectPriorityText}</option>
                    <option style={{color: "grey"}} value="1">{LowestPriorityText}</option>
                    <option style={{color: "blue"}} value="2">{LowPriorityText}</option>
                    <option style={{color: "green"}} value="3">{MediumPriorityText}</option>
                    <option style={{color: "orange"}} value="4">{HighPriorityText}</option>
                    <option style={{color: "red"}} value="5">{HighestPriorityText}</option>
                </Form.Select>
            </FloatingLabel>

            <FloatingLabel controlId="floatingRecurring" label={RecurringLabelText}>
                <Form.Select
              name="recurring"
              value={todoRecurring}
              onChange={(e) => setTodoRecurring(e.target.value)}
                    style={{ margin: '0px 0px 32px' }}>>
                    <option value="false">{NoText}</option>
                    <option value="true">{YesText}</option>
                </Form.Select>
            </FloatingLabel>
            <FloatingLabel
                controlId="floatingRecurringPeriod"
                style={{ display: todoRecurring === "false" || todoRecurring === false ? "none" : "block" }}
                label={RecurrencyLabelText}>
                <Form.Select
                    name="recurring"
                    value={todoRecurrencyPeriod}
                    onChange={(e) => setTodoRecurrencyPeriod(e.target.value)}
                    style={{ margin: '0px 0px 32px' }}>>
                    <option disabled={true}>{SelectRecurringText}</option>
                    <option value="daily">{DailyText}</option>
                    <option value="weekly">{WeeklyText}</option>
                    <option value="biweekly">{BiweeklyText}</option>
                    <option value="monthly">{MonthlyText}</option>
                    <option value="weekdays">{WeekdaysText}</option>
                    <option value="weekends">{WeekendsText}</option>
                </Form.Select>
            </FloatingLabel>
            <FloatingLabel
                controlId="floatingLink"
                label={LinkLabelText}
                value={todoLink}
                onChange={(e) => setTodoLink(e.target.value)}
            >
                <Form.Control type="link" placeholder="Link" />
            </FloatingLabel>

            <ButtonGroup style={{width: "100%", paddingLeft: "10%", paddingRight: "10%"}}>
            <Button
              variant="danger"
              onClick={() => navigateTodos()}
              style={{margin: "auto", display: "block", padding: "0", textAlign: "center"}}
            >{CancelButtonText}</Button>
            <Button
                variant="success"
                type="submit"
                style={{margin: "auto", display: "block", padding: "0", textAlign: "center"}}
            >{CreateButtonText}</Button>
          </ButtonGroup>

        </Form>

          <Modal className='successModal text-center' show={showEnterTodoNameModal} open={showEnterTodoNameModal} centered={true} size='lg'>
              <ModalBody>
                  <h4 style={{margin: "32px"}}>{PleaseEnterTodoNameText}</h4>
                  <ButtonGroup style={{width: "40%"}}>
                      <Button
                          variant="danger"
                          onClick={(e) => toggleEnterTodoNameModal(e)}
                          style={{margin: "auto", display: "block", padding: "0", textAlign: "center"}}
                      >{CancelButtonText}</Button>
                  </ButtonGroup>
              </ModalBody>
          </Modal>

      </Container>
    )
  }
;


export default CreateTodo;


