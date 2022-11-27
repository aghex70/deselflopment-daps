import React from 'react';
import {
  faClockRotateLeft,
  faChartSimple,
  faHome,
  faPowerOff,
  faCog, faCheck,
} from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import {Button, ButtonGroup, Container} from "react-bootstrap";
import checkAccess from "../utils/helpers";

const DapsHeader = () => {
  document.title = 'deselflopment - daps'
  checkAccess();

  const navigateCategories = () => {
    window.location.href = "/categories";
  }

  const navigateCompletedTodos = () => {
    window.location.href = "/completed-todos";
  }

  const navigateRecurringTodos = () => {
    window.location.href = "/recurring-todos";
  }

  const logout = () => {
    window.location.href = "/logout";
  }

      return (
        <Container>
          <ButtonGroup style={{width: "45%", marginTop: "15px", marginBottom: "15px"}}>
            <Button style={{width: "15%", margin: "auto", padding: "0", textAlign: "center"}}
                    variant="primary"
                    onClick={() => navigateCategories()}
                    title="Categories"
            >
              <FontAwesomeIcon icon={faHome} />
            </Button>
            <Button style={{width: "15%", margin: "auto", padding: "0", textAlign: "center"}}
                    variant="success"
                    onClick={() => navigateCompletedTodos()}
                    title="Completed Todos"
            >
              <FontAwesomeIcon icon={faCheck} />
            </Button>
            <Button style={{width: "15%", margin: "auto", display: "block", padding: "0", textAlign: "center"}}
                    variant="secondary"
                    onClick={() => navigateRecurringTodos()}
                    title="Recurring Todos"
            >

              <FontAwesomeIcon icon={faClockRotateLeft} />
            </Button>
            <Button disabled={true} style={{width: "15%", margin: "auto", display: "block", padding: "0", textAlign: "center"}}
                    variant="outline-warning" title="Statistics (coming soon)"
            >

              <FontAwesomeIcon icon={faChartSimple} />
            </Button>
            <Button disabled={true} style={{width: "15%", margin: "auto", display: "block", padding: "0", textAlign: "center"}}
                    variant="outline-info" title="Configuration (coming soon)"
            >

              <FontAwesomeIcon icon={faCog} />
            </Button>
            <Button style={{width: "15%", margin: "auto", display: "block", padding: "0", textAlign: "center"}}
                    variant="danger" title="Logout"
                    onClick={() => logout()}
            >

              <FontAwesomeIcon style={{color: "white"}} icon={faPowerOff} />
            </Button>
          </ButtonGroup>
        </Container>
      );
    };

export default DapsHeader;
