import React, {useEffect, useState} from 'react';
import {
  faClockRotateLeft,
  faChartSimple,
  faHome,
  faPowerOff,
  faCheck,
  faEnvelope,
  faUserPlus,
  faCog,
  faList, faFileImport, faLightbulb,
} from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import {Button, ButtonGroup, Container} from "react-bootstrap";
import checkAccess, {
  goToCategories,
  goToCompletedTodos,
  goToListOfUsers,
  goToLogout,
  goToProfile,
  goToProvisionDemoUser,
  goToRecurringTodos,
  goToReportABug,
  goToSuggestedTodos
} from "../utils/helpers";
import {
  CategoriesIconText,
  CompletedTodosIconText,
  ImportTodosHeaderText,
  ListOfUsersIconText,
  LogoutIconText,
  ProfileIconText,
  ProvisionDemoUserIconText,
  RecurringTodosIconText,
  ReportABugIconText,
  StatisticsIconText, SuggestedTodosIconText
} from "../utils/texts";
import UserService from "../services/user";

const DapsHeader = () => {
  document.title = 'deselflopment - daps'
  checkAccess();

  const [isHoverProfile, setIsHoverProfile] = useState(false);
  const [isAdmin, setIsAdmin] = useState(false);

  const handleMouseEnter = () => {
    setIsHoverProfile(true);
  };
  const handleMouseLeave = () => {
    setIsHoverProfile(false);
  };

  useEffect(() => {
    UserService.checkAdminAccess().then(
        (response) => {
          if (response.status === 200) {
            setIsAdmin(true);
          }
        }
    ).catch(
        (error) => {
          setIsAdmin(false);
        }
    )
  }, [isAdmin]);

      return (
        <Container>
          <ButtonGroup style={{width: "100%", marginTop: "15px", marginBottom: "15px"}}>
            <Button style={{height: "50px", width: "100%", margin: "auto", padding: "0", textAlign: "center"}}
                    variant="primary"
                    onClick={() => goToCategories()}
                    title={CategoriesIconText}
            >
              <FontAwesomeIcon style={{height: "50%"}} icon={faHome} />
            </Button>
            <Button style={{height: "50px", width: "100%", margin: "auto", padding: "0", textAlign: "center"}}
                    variant="success"
                    onClick={() => goToCompletedTodos()}
                    title={CompletedTodosIconText}
            >
              <FontAwesomeIcon style={{height: "50%"}} icon={faCheck} />
            </Button>
            <Button style={{height: "50px", width: "100%", margin: "auto", padding: "0", textAlign: "center"}}
                    variant="warning"
                    onClick={() => goToSuggestedTodos()}
                    title={SuggestedTodosIconText}
            >

              <FontAwesomeIcon style={{height: "50%", color: "white"}} icon={faLightbulb} />
            </Button>
            <Button style={{height: "50px", width: "100%", margin: "auto", padding: "0", textAlign: "center"}}
                    variant="secondary"
                    onClick={() => goToRecurringTodos()}
                    title={RecurringTodosIconText}
            >
              <FontAwesomeIcon style={{height: "50%"}} icon={faClockRotateLeft} />
            </Button>
            <Button style={{height: "50px", width: "100%", margin: "auto", padding: "0", textAlign: "center"}}
                    variant="info" title={ReportABugIconText}
                    onClick={() => goToReportABug()}
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
                    onClick={() => goToProfile()}
                    onMouseEnter={handleMouseEnter}
                    onMouseLeave={handleMouseLeave}
            >

              <FontAwesomeIcon style={{height: "50%", color: "white"}} icon={faCog} />

            </Button>
            <Button style={{height: "50px", width: "100%", margin: "auto", padding: "0", textAlign: "center"}}
                    variant="danger" title={LogoutIconText}
                    onClick={() => goToLogout()}
            >

              <FontAwesomeIcon style={{height: "50%", color: "white"}} icon={faPowerOff} />
            </Button>
          </ButtonGroup>

          {isAdmin && (
          <ButtonGroup style={{width: "100%", marginTop: "15px", marginBottom: "15px"}}>
            <Button disabled={true} style={{height: "50px", width: "100%", margin: "auto", padding: "0", textAlign: "center"}}
                    variant="outline-warning" title={StatisticsIconText}
            >

              <FontAwesomeIcon icon={faChartSimple} />
            </Button>
            <Button style={{height: "50px", width: "100%", margin: "auto", padding: "0", textAlign: "center"}}
                    variant="danger" title={ProvisionDemoUserIconText}
                    onClick={() => goToProvisionDemoUser()}
            >

              <FontAwesomeIcon style={{height: "50%", color: "white"}} icon={faUserPlus} />
            </Button>
            <Button style={{height: "50px", width: "100%", margin: "auto", padding: "0", textAlign: "center"}}
                    variant="danger" title={ListOfUsersIconText}
                    onClick={() => goToListOfUsers()}
            >

              <FontAwesomeIcon style={{height: "50%", color: "white"}} icon={faList} />
            </Button>
            <Button style={{height: "50px", width: "100%", margin: "auto", padding: "0", textAlign: "center"}}
                    variant="danger" title={ImportTodosHeaderText}
                    onClick={() => goToImportTodos()}
            >

              <FontAwesomeIcon style={{height: "50%", color: "white"}} icon={faFileImport} />
            </Button>
          </ButtonGroup>
          )}

        </Container>
      );
    };

export default DapsHeader;
