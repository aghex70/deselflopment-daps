import React, {useState} from 'react';
import {Button, ButtonGroup, Container, Form, Modal, ModalBody} from "react-bootstrap";
import DapsHeader from "./Header";
import checkAccess, {goToImportTodos} from "../utils/helpers";
import {
    CancelButtonText,
    CheckFileFormatText,
    ImportButtonText,
    ImportTodosHeaderText,
} from "../utils/texts";
import UserService from "../services/user";

const ImportCSV = () => {
    checkAccess();
    const [file, setFile] = useState();
    const [showIncorrectFileFormat, setShowIncorrectFileFormat] = useState(false);

    const toggleIncorrectFileFormat = () => {
        setShowIncorrectFileFormat(!showIncorrectFileFormat);
    }

    const handleSubmit = (event) => {
        event.preventDefault();
        UserService.importCSV(file)
            .then(response => {
                console.log(response.data);
            })
            .catch(error => {
                setShowIncorrectFileFormat(true);
            });
    }

    const handleFileChange = (event) => {
        setFile(event.target.files[0]);
    }

    return (
        <Container>
            <DapsHeader/>
            <h1 className="text-center">{ImportTodosHeaderText}</h1>
            <Form onSubmit={handleSubmit}>
                <Form.Group controlId="formFile">
                    <Form.Control
                        type="file"
                        onChange={handleFileChange}
                        style={{width: "80%", marginLeft: "10%", marginRight: "10%"}}
                    />
                </Form.Group>

            <ButtonGroup style={{width: "100%", paddingLeft: "10%", paddingRight: "10%"}}>
                <Button
                    variant="danger"
                    onClick={() => goToImportTodos()}
                    style={{margin: "auto", display: "block", padding: "0", textAlign: "center"}}
                >{CancelButtonText}</Button>
                <Button
                    variant="success"
                    type="submit"
                    style={{margin: "auto", display: "block", padding: "0", textAlign: "center"}}
                >{ImportButtonText}</Button>
            </ButtonGroup>
            </Form>

        <Modal className='successModal text-center' show={showIncorrectFileFormat} open={showIncorrectFileFormat} centered={true} size='lg'>
            <ModalBody>
                <h4 style={{margin: "32px"}}>{CheckFileFormatText}</h4>
                <ButtonGroup style={{width: "40%"}}>
                    <Button
                        variant="danger"
                        onClick={(e) => toggleIncorrectFileFormat(e)}
                        style={{margin: "auto", display: "block", padding: "0", textAlign: "center"}}
                    >{CancelButtonText}</Button>
                </ButtonGroup>
            </ModalBody>
        </Modal>
        </Container>
    )
}
;


export default ImportCSV;