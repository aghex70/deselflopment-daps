import React, {useEffect, useState} from 'react';
import {Button, ButtonGroup, Container} from "react-bootstrap";
import DapsHeader from "./Header";
import checkAccess from "../utils/helpers";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import {
    DeleteIconText,
    HeaderActionsText,
    HeaderUserText,
    UsersHeaderText,
    ViewIconText,
} from "../utils/texts";
import UserService from "../services/user";
import BootstrapTable from "react-bootstrap-table-next";
import {faEye, faTrash} from "@fortawesome/free-solid-svg-icons";
import {useNavigate} from "react-router-dom";

const UsersList = () => {
    checkAccess();
    const navigate = useNavigate();
    const [users, setUsers] = useState([]);

    const navigateToUser = (id) => {
        navigate("/user/" + id);
    }

    const deleteUser = (id) => {
        UserService.deleteUser(id).then(
            (response) => {
                if (response.status === 204) {
                    window.location.reload();
                }
            }
        ).catch(
            (error) => {
                error = new Error("Deletion failed!");
            })
    }

    const columns = [
        {
            dataField: 'email',
            text: HeaderUserText,
            style:{'width' : '70%', cursor: "pointer", verticalAlign: "middle"},
        },
        {
            dataField: 'link',
            text: HeaderActionsText,
            style:{'width' : '30%', verticalAlign: "middle"},
            formatter: actionsFormatter,
            headerAlign: 'center',
        }];

    useEffect(() => {
        UserService.checkAdminAccess().then(
            (response) => {
                if (response.status !== 200) {
                    window.location.href = "/categories";
                }
            }
        ).catch(
            (error) => {
                window.location.href = "/categories";

            }
        )

        if (!users || users.length === 0) {
            UserService.getUsers().then(
                (response) => {
                    if (response.status === 200) {
                        setUsers(response.data.users);
                    }
                }
            ).catch(
                (error) => {
                })
        }
    }, [users]);

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
                    <Button style={{width: "15%", margin: "auto", padding: "0", textAlign: "center"}}
                            title={ViewIconText}
                            variant="outline-primary"
                            onClick={() => navigateToUser(row.id)}
                    >
                        <FontAwesomeIcon icon={faEye} />
                    </Button>

                    <Button style={{width: "15%", margin: "auto", display: "block", padding: "0", textAlign: "center"}}
                            title={DeleteIconText}
                            variant="outline-danger"
                           onClick={() => deleteUser(row.id)}
                    >
                        <FontAwesomeIcon icon={faTrash} />
                    </Button>
                </ButtonGroup>
            </div>
        );
    }

    return (
        <Container>
        <DapsHeader />
        <h1 className="text-center">{UsersHeaderText}</h1>
        <BootstrapTable
            keyField='id'
            data={ users }
            columns={ columns }
            hover={true}
        />
        </Container>
    )
}
;

export default UsersList;
