import React, {useState} from 'react';
import {
  faClockRotateLeft,
  faChartSimple,
  faHome,
  faPowerOff,
  faCheck,
  faEnvelope,
  faUser, faUserPlus,
} from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import {Button, ButtonGroup, Container} from "react-bootstrap";
import checkAccess from "../utils/helpers";
import {
  CategoriesIconText,
  CompletedTodosIconText,
  LogoutIconText,
  ProfileIconText, ProvisionDemoUserIconText,
  RecurringTodosIconText,
  ReportABugIconText,
  StatisticsIconText
} from "../utils/texts";
import UserService from "../services/user";

const DapsHeader = () => {
  document.title = 'deselflopment - daps'
  checkAccess();

  const [isHoverProfile, setIsHoverProfile] = useState(false);

  const handleMouseEnter = () => {
    setIsHoverProfile(true);
  };
  const handleMouseLeave = () => {
    setIsHoverProfile(false);
  };

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

  const navigateProfile = () => {
    window.location.href = "/profile";
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
                    title={CategoriesIconText}
            >
              <FontAwesomeIcon style={{height: "50%"}} icon={faHome} />
            </Button>
            <Button style={{height: "50px", width: "100%", margin: "auto", padding: "0", textAlign: "center"}}
                    variant="success"
                    onClick={() => navigateCompletedTodos()}
                    title={CompletedTodosIconText}
            >
              <FontAwesomeIcon style={{height: "50%"}} icon={faCheck} />
            </Button>
            <Button style={{height: "50px", width: "100%", margin: "auto", padding: "0", textAlign: "center"}}
                    variant="secondary"
                    onClick={() => navigateRecurringTodos()}
                    title={RecurringTodosIconText}
            >
              <FontAwesomeIcon style={{height: "50%"}} icon={faClockRotateLeft} />
            </Button>
            <Button disabled={true} style={{height: "50px", width: "100%", margin: "auto", padding: "0", textAlign: "center"}}
                    variant="outline-warning" title={StatisticsIconText}
            >

              <FontAwesomeIcon icon={faChartSimple} />
            </Button>
            <Button style={{height: "50px", width: "100%", margin: "auto", padding: "0", textAlign: "center"}}
                    variant="info" title={ReportABugIconText}
                    onClick={() => navigateReportBug()}
            >

              <FontAwesomeIcon style={{height: "50%", color: "white"}} icon={faEnvelope} />
            </Button>
            <Button style={{
              height: "50px",
              width: "100%",
              margin: "auto",
              display: "block",
              padding: "0",
              textAlign: "center",
              backgroundColor: isHoverProfile ? "#eab676": "orange",
              borderColor: isHoverProfile ? "#eab676": "orange",
            }}
                    title={ProfileIconText}
                    onClick={() => navigateProfile()}
                    onMouseEnter={handleMouseEnter}
                    onMouseLeave={handleMouseLeave}
            >

              <FontAwesomeIcon style={{height: "50%", color: "white"}} icon={faUser} />

            </Button>
            <Button style={{height: "50px", width: "100%", margin: "auto", padding: "0", textAlign: "center"}}
                    variant="danger" title={LogoutIconText}
                    onClick={() => logout()}
            >

              <FontAwesomeIcon style={{height: "50%", color: "white"}} icon={faPowerOff} />
            </Button>

            <Button style={{height: "50px", width: "100%", margin: "auto", padding: "0", textAlign: "center"}}
                    variant="danger" title={ProvisionDemoUserIconText}
                    onClick={() => UserService.provisionDemoUser("demo@demo.com", "RU")}
            >

              <FontAwesomeIcon style={{height: "50%", color: "white"}} icon={faUserPlus} />
            </Button>
          </ButtonGroup>
        </Container>
      );
    };

export default DapsHeader;
