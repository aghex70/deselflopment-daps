import React, { useEffect, useState } from "react";
import {
  Button,
  ButtonGroup,
  Container,
  FloatingLabel,
  Form,
} from "react-bootstrap";
import { useParams } from "react-router-dom";
import DapsHeader from "./Header";
import checkAccess, {
  goToCategories,
  setAutoSuggest,
  setAutoRemind,
  setLanguage,
  checkValidToken,
} from "../utils/helpers";
import {
  AutoSuggestLabelText,
  CancelButtonText,
  EditButtonText,
  EmailAddressLabelText,
  EnglishLanguageText,
  LanguageLabelText,
  NoText,
  ProfileHeaderText,
  SpanishLanguageText,
  YesText,
  AutoRemindLabelText,
} from "../utils/texts";
import UserService from "../services/user";
import toBoolean from "validator/es/lib/toBoolean";

const Profile = () => {
  checkAccess();
  const [userEmail, setUserEmail] = useState("");
  const [profileLanguage, setProfileLanguage] = useState("en");
  const [profileAutoSuggest, setProfileAutoSuggest] = useState("false");
  const [profileAutoRemind, setProfileAutoRemind] = useState("false");
  const { id } = useParams();

  const handleSubmit = (e) => {
    e.preventDefault();

    const data = {
      language: profileLanguage,
      auto_suggest:
        typeof profileAutoSuggest == "boolean"
          ? profileAutoSuggest
          : toBoolean(profileAutoSuggest),
      // auto_remind:
      //   typeof profileAutoRemind == "boolean"
      //     ? profileAutoRemind
      //     : toBoolean(profileAutoRemind),
    };

    UserService.editProfile(data)
      .then((response) => {
        if (response.status === 200) {
          setLanguage(data.language);
          setAutoSuggest(data.auto_suggest);
          // setAutoRemind(data.auto_remind);
          goToCategories();
        }
      })
      .catch((error) => {
        if (error.response.data.message === "no changes were made") {
          goToCategories();
        }
        console.log("error", error); // checkValidToken(error);
      });
  };

  useEffect(() => {
    UserService.getProfile()
      .then((response) => {
        if (response.status === 200) {
          setProfileLanguage(response.data.language);
          setProfileAutoSuggest(response.data.auto_suggest);
          // setProfileAutoRemind(response.data.auto_remind);
          setUserEmail(response.data.email);
        }
      })
      .catch((error) => {
        console.log("error", error); // checkValidToken(error);
      });
  }, []);

  return (
    <Container>
      <DapsHeader />
      <h1 className="text-center">{ProfileHeaderText}</h1>
      <Form onSubmit={(e) => handleSubmit(e)}>
        <FloatingLabel
          controlId="floatingEmail"
          label={EmailAddressLabelText}
          value={userEmail}
        >
          <Form.Control
            type="email"
            placeholder="Email"
            value={userEmail}
            disabled={true}
          />
        </FloatingLabel>

        <FloatingLabel controlId="floatingLanguage" label={LanguageLabelText}>
          <Form.Select
            name="language"
            value={profileLanguage}
            onChange={(e) => setProfileLanguage(e.target.value)}
            style={{ margin: "0px 0px 32px" }}
          >
            <option style={{ color: "red" }} value="es">
              {SpanishLanguageText}
            </option>
            <option style={{ color: "blue" }} value="en">
              {EnglishLanguageText}
            </option>
          </Form.Select>
        </FloatingLabel>

        <FloatingLabel
          controlId="floatingAutoSuggest"
          label={AutoSuggestLabelText}
        >
          <Form.Select
            name="auto-suggest"
            value={profileAutoSuggest}
            onChange={(e) => setProfileAutoSuggest(e.target.value)}
            style={{ margin: "0px 0px 32px" }}
          >
            <option value="false">{NoText}</option>
            <option value="true">{YesText}</option>
          </Form.Select>
        </FloatingLabel>

        {/*<FloatingLabel*/}
        {/*  controlId="floatingAutoRemind"*/}
        {/*  label={AutoRemindLabelText}*/}
        {/*>*/}
        {/*  <Form.Select*/}
        {/*    name="auto-remind"*/}
        {/*    value={profileAutoRemind}*/}
        {/*    onChange={(e) => setProfileAutoRemind(e.target.value)}*/}
        {/*    style={{ margin: "0px 0px 32px" }}*/}
        {/*  >*/}
        {/*    <option value="false">{NoText}</option>*/}
        {/*    <option value="true">{YesText}</option>*/}
        {/*  </Form.Select>*/}
        {/*</FloatingLabel>*/}

        <ButtonGroup
          style={{ width: "100%", paddingLeft: "10%", paddingRight: "10%" }}
        >
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
            {CancelButtonText}
          </Button>
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
            {EditButtonText}
          </Button>
        </ButtonGroup>
      </Form>
    </Container>
  );
};
export default Profile;
