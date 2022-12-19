import React, {useEffect, useState} from 'react';
import {Button, ButtonGroup, Container, FloatingLabel, Form} from "react-bootstrap";
import DapsHeader from "./Header";
import checkAccess from "../utils/helpers";
import {
    CancelButtonText,
    DeleteButtonText,
    EmailAddressLabelText,
    NameLabelText,
    RegistrationDateText,
    ViewUserHeaderText,
} from "../utils/texts";
import UserService from "../services/user";
import {useParams} from "react-router-dom";

const User = () => {
        checkAccess();
        const { id } = useParams();
        const [userName, setUserName] = useState("");
        const [userEmail, setUserEmail] = useState("");
        const [userRegistrationDate, setUserRegistrationDate] = useState("");

        const navigateUsers = () => {
            window.location.href = "/users";
        }

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

            if (!userEmail) {
                UserService.getUser(id).then(
                    (response) => {
                        if (response.status === 200) {
                            setUserName(response.data.name);
                            setUserEmail(response.data.email);
                            setUserRegistrationDate(response.data.registration_date);
                        }
                    }
                ).catch(
                    (error) => {
                    })
            }
        }, [userName]);

        return (
            <Container>
                <DapsHeader/>
                <h1 className="text-center">{ViewUserHeaderText}</h1>
                <Form>
                    <FloatingLabel
                        controlId="floatingName"
                        label={NameLabelText}
                        value={userName}
                    >
                        <Form.Control type="name" placeholder="Name" value={userName} disabled={true}/>
                    </FloatingLabel>

                    <FloatingLabel
                        controlId="floatingEmail"
                        label={EmailAddressLabelText}
                        value={userEmail}
                    >
                        <Form.Control type="description" placeholder="Description" value={userEmail} disabled={true}/>
                    </FloatingLabel>

                    <FloatingLabel
                        controlId="floatingRegistrationDate"
                        label={RegistrationDateText}
                        value={userRegistrationDate}
                    >
                        <Form.Control type="registration_date" placeholder="Registration date"
                                      value={userRegistrationDate} disabled={true}/>
                    </FloatingLabel>

                    <ButtonGroup style={{width: "100%", paddingLeft: "10%", paddingRight: "10%"}}>
                        <Button
                            variant="success"
                            onClick={() => navigateUsers()}
                            style={{margin: "auto", display: "block", padding: "0", textAlign: "center"}}
                        >{CancelButtonText}</Button>
                        <Button
                            variant="danger"
                            onClick={() => navigateUsers()}
                            style={{margin: "auto", display: "block", padding: "0", textAlign: "center"}}
                        >{DeleteButtonText}</Button>
                    </ButtonGroup>
                </Form>
            </Container>
        )
    }
;


export default User;