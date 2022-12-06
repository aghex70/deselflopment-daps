import React, { useState } from "react";
import { Button, ButtonGroup, Container, Form, FloatingLabel } from "react-bootstrap";
import TodoService from "../services/todo";
import DapsHeader from "./Header";
import checkAccess from "../utils/helpers";

const ReportBug = () => {
    checkAccess();
    const [todoDescription, setTodoDescription] = useState("");

    const handleSubmit = (e) => {
        e.preventDefault();

        const data = {
            name: new Date().getTime().toString(),
            description: todoDescription,
            priority: 5,
            recurring: false,
            category_id: 1,
        };

        TodoService.createTodo(data)
            .then((response) => {
                window.location.href = "/categories";
            })
            .catch((error) => {
                error = new Error("Update todo failed!");
            });
    };

    return (
        <Container>
            <DapsHeader />
            <h1 style={{ margin: "0px 0px 32px" }} className="text-center">
                Report a bug
            </h1>
            <Form onSubmit={handleSubmit}>
                <FloatingLabel controlId="floatingDescription" label="Description">
                    <Form.Control
                        as="textarea"
                        placeholder="Description"
                        style={{ height: "300px", margin: "0px 0px 32px" }}
                        type="description"
                        value={todoDescription}
                        onChange={(e) => setTodoDescription(e.target.value)}
                    />
                </FloatingLabel>

                <ButtonGroup
                    style={{ width: "100%", paddingLeft: "10%", paddingRight: "10%" }}
                >
                    <Button
                        variant="success"
                        type="submit"
                        style={{
                            margin: "auto",
                            display: "block",
                            padding: "0",
                            textAlign: "center",
                        }}
                    >
                        Report
                    </Button>
                    <Button
                        variant="danger"
                        onClick={() => (window.location.href = "/categories")}
                        style={{
                            margin: "auto",
                            display: "block",
                            padding: "0",
                            textAlign: "center",
                        }}
                    >
                        Cancel
                    </Button>
                </ButtonGroup>
            </Form>
        </Container>
    );
};

export default ReportBug;
