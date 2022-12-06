import React from 'react';
import {
  faClockRotateLeft,
  faChartSimple,
  faHome,
  faPowerOff,
  faCog,
  faCheck,
    faEnvelope,
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

  const navigateReportBug = () => {
    window.location.href = "/report-bug";
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
          <ButtonGroup style={{width: "100%", marginTop: "15px", marginBottom: "15px"}}>
            <Button style={{height: "50px", width: "100%", margin: "auto", padding: "0", textAlign: "center"}}
                    variant="primary"
                    onClick={() => navigateCategories()}
                    title="Categories"
            >
              <FontAwesomeIcon style={{height: "50%"}} icon={faHome} />
            </Button>
            <Button style={{height: "50px", width: "100%", margin: "auto", padding: "0", textAlign: "center"}}
                    variant="success"
                    onClick={() => navigateCompletedTodos()}
                    title="Completed Todos"
            >
              <FontAwesomeIcon style={{height: "50%"}} icon={faCheck} />
            </Button>
            <Button style={{height: "50px", width: "100%", margin: "auto", padding: "0", textAlign: "center"}}
                    variant="secondary"
                    onClick={() => navigateRecurringTodos()}
                    title="Recurring Todos"
            >

              <FontAwesomeIcon style={{height: "50%"}} icon={faClockRotateLeft} />
            </Button>
            <Button disabled={true} style={{height: "50px", width: "100%", margin: "auto", padding: "0", textAlign: "center"}}
                    variant="outline-warning" title="Statistics (coming soon)"
            >

              <FontAwesomeIcon icon={faChartSimple} />
            </Button>
            <Button style={{height: "50px", width: "100%", margin: "auto", padding: "0", textAlign: "center"}}
                    variant="info" title="Report a bug"
                    onClick={() => navigateReportBug()}
            >

              <FontAwesomeIcon style={{height: "50%", color: "white"}} icon={faEnvelope} />
            </Button>
            <Button disabled={true} style={{height: "50px", width: "100%", margin: "auto", display: "block", padding: "0", textAlign: "center"}}
                    variant="outline-dark" title="Configuration (coming soon)"
            >

              <FontAwesomeIcon icon={faCog} />

            </Button>
            <Button style={{height: "50px", width: "100%", margin: "auto", padding: "0", textAlign: "center"}}
                    variant="danger" title="Logout"
                    onClick={() => logout()}
            >

              <FontAwesomeIcon style={{height: "50%", color: "white"}} icon={faPowerOff} />
            </Button>
          </ButtonGroup>
        </Container>
      );
    };

export default DapsHeader;
