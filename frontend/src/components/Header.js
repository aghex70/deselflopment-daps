import {faClockRotateLeft, faBookOpen, faChartSimple, faThumbsUp, faHome} from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import {Button, ButtonGroup, Container} from "react-bootstrap";
// import './TodosList.css';
// import BootstrapTable from 'react-bootstrap-table-next';
// import createTodo from "./CreateTodo";

const DapsHeader = () => {

  const navigateCategories = () => {
    window.location.href = "/categories";
  }

  const navigateCompletedTodos = () => {
    window.location.href = "/completed-todos";
  }

  const navigateRecurringTodos = () => {
    window.location.href = "/recurring-todos";
  }

      return (
        <Container>
          <ButtonGroup style={{width: "35%", marginTop: "15px", marginBottom: "15px"}}>
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
              <FontAwesomeIcon icon={faThumbsUp} />
            </Button>
            <Button style={{width: "15%", margin: "auto", display: "block", padding: "0", textAlign: "center"}}
                    variant="secondary"
                    onClick={() => navigateRecurringTodos()}
                    title="Recurring Todos"
            >

              <FontAwesomeIcon icon={faClockRotateLeft} />
            </Button>
            <Button style={{width: "15%", margin: "auto", display: "block", padding: "0", textAlign: "center"}}
                    variant="warning" title="Statistics (coming soon)"
            >

              <FontAwesomeIcon style={{color: "white"}} icon={faChartSimple} />
            </Button>
          </ButtonGroup>
        </Container>
      );
    };

export default DapsHeader;
