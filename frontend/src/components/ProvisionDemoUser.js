import React, {useEffect, useState} from 'react';
import {Button, ButtonGroup, Container, FloatingLabel, Form, Modal, ModalBody} from "react-bootstrap";
import DapsHeader from "./Header";
import checkAccess, {goToCategories, goToListOfUsers, hashPassword} from "../utils/helpers";
import {
    CancelButtonText,
    CreateButtonText,
    DemoUserAlreadyProvisionedText,
    EmailAddressLabelText,
    EnglishLanguageText,
    LanguageLabelText,
    PleaseEnterAnEmailText,
    ProvisionDemoUserIconText,
    SpanishLanguageText,
} from "../utils/texts";
import UserService from "../services/user";

const ProvisionDemoUser = () => {
    checkAccess();
    const [userEmail, setUserEmail] = useState("");
    const [userLanguage, setUserLanguage] = useState("en");
    const [showModalDemoUserAlreadyCreated, setShowModalDemoUserAlreadyCreated] = useState(false);
    const [showModalEmptyEmail, setShowModalEmptyEmail] = useState(false);

    const toggleModalDemoUserAlreadyCreated = () => {
        setShowModalDemoUserAlreadyCreated(!showModalDemoUserAlreadyCreated);
    }

    const toggleModalEmptyEmail = () => {
        setShowModalEmptyEmail(!showModalEmptyEmail);
    }

    useEffect(() => {
        UserService.checkAdminAccess().then(
            (response) => {
                if (response.status !== 200) {
                    goToCategories();
                }
            }
        ).catch(
            (error) => {
                goToCategories();

            }
        )
    }, []);

    const handleSubmit = (e) => {
        e.preventDefault();

        if (!userEmail) {
            toggleModalEmptyEmail();
            return;
        }

        const hashedPassword = hashPassword(process.env.REACT_APP_DEMO_USER_PASSWORD);
        UserService.provisionDemoUser(userEmail, hashedPassword, userLanguage).then(
            (response) => {
                if (response.status === 201) {
                    goToListOfUsers();
                }
            }
        ).catch(
            (error) => {
                if (error.response.data.message === "unauthorized") {
                    goToCategories();
                }
                setShowModalDemoUserAlreadyCreated(true);
            }
        )
    }

    return (
        <Container>
            <DapsHeader />
            <h1 className="text-center">{ProvisionDemoUserIconText}</h1>
            <Form onSubmit={(e) => handleSubmit(e)}>
                <FloatingLabel
                    controlId="floatingEmail"
                    label={EmailAddressLabelText}
                    value={userEmail}
                    onChange={(e) => setUserEmail(e.target.value)}
                >
                    <Form.Control type="email" placeholder="Email" value={userEmail} onChange={(e) => setUserEmail(e.target.value)}/>
                </FloatingLabel>

                <FloatingLabel controlId="floatingLanguage" label={LanguageLabelText}>
                    <Form.Select
                        name="language"
                        value={userLanguage}
                        onChange={(e) => setUserLanguage(e.target.value)}
                        style={{ margin: '0px 0px 32px' }}
                    >
                        <option style={{color: "red"}} value="es">{SpanishLanguageText}</option>
                        <option style={{color: "blue"}} value="en">{EnglishLanguageText}</option>
                    </Form.Select>
                </FloatingLabel>

                <ButtonGroup style={{width: "100%", paddingLeft: "10%", paddingRight: "10%"}}>
                    <Button
                        variant="danger"
                        onClick={() => window.location.href = "/categories"}
                        style={{margin: "auto", display: "block", padding: "0", textAlign: "center"}}
                    >{CancelButtonText}</Button>
                    <Button
                        variant="success"
                        type="submit"
                        style={{margin: "auto", display: "block", padding: "0", textAlign: "center"}}
                    >{CreateButtonText}</Button>
                </ButtonGroup>
            </Form>

            <Modal className='successModal text-center' show={showModalDemoUserAlreadyCreated} open={showModalDemoUserAlreadyCreated} centered={true} size='lg'>
                <ModalBody>
                    <h4 style={{margin: "32px"}}>{DemoUserAlreadyProvisionedText}</h4>
                    <ButtonGroup style={{width: "40%"}}>
                        <Button
                            variant="danger"
                            onClick={(e) => toggleModalDemoUserAlreadyCreated(e)}
                            style={{margin: "auto", display: "block", padding: "0", textAlign: "center"}}
                        >{CancelButtonText}</Button>
                    </ButtonGroup>
                </ModalBody>
            </Modal>

            <Modal className='successModal text-center' show={showModalEmptyEmail} open={showModalEmptyEmail} centered={true} size='lg'>
                <ModalBody>
                    <h4 style={{margin: "32px"}}>{PleaseEnterAnEmailText}</h4>
                    <ButtonGroup style={{width: "40%"}}>
                        <Button
                            variant="danger"
                            onClick={(e) => toggleModalEmptyEmail(e)}
                            style={{margin: "auto", display: "block", padding: "0", textAlign: "center"}}
                        >{CancelButtonText}</Button>
                    </ButtonGroup>
                </ModalBody>
            </Modal>
        </Container>
    )
}
;


export default ProvisionDemoUser;


