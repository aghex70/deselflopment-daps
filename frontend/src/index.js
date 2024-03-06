import React from "react";
import ReactDOM from "react-dom/client";
import "./index.css";
import "bootstrap/dist/css/bootstrap.min.css";
import App from "./App";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import CategoriesList from "./components/CategoriesList";
import TodosList from "./components/TodosList";
import RecurringTodosList from "./components/RecurringTodosList";
import Category from "./components/Category";
import CreateCategory from "./components/CreateCategory";
import CreateTodo from "./components/CreateTodo";
import Todo from "./components/Todo";
import CompletedTodosList from "./components/CompletedTodosList";
import ReportBug from "./components/ReportBug";
import Profile from "./components/UserProfile";
import ProvisionDemoUser from "./components/ProvisionDemoUser";
import UsersList from "./components/UsersList";
import User from "./components/User";
import ImportCSV from "./components/ImportCSV";
import SuggestedTodosList from "./components/SuggestedTodosList";
import QueryParamsToLocalStorage from "./components/Synchronizer";
import Logout from "./components/Logout";
import Login from "./components/Login";
import Register from "./components/Register";

const production =
  process.env.REACT_APP_API_URL === "https://daps.deselflopment.com/";
const root = ReactDOM.createRoot(document.getElementById("root"));

root.render(
  production ? (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<App />} />
        <Route path="login" element={<Login/>}/>
        <Route path="logout" element={<Logout/>}/>
        <Route path="register" element={<Register/>}/>
        <Route path="sync" element={<QueryParamsToLocalStorage />} />
        <Route path="report-bug" element={<ReportBug />} />
        <Route path="categories" element={<CategoriesList />} />
        <Route path="todos" element={<TodosList />} />
        <Route path="category/:id" element={<Category />} />
        <Route path="todo/:id" element={<Todo />} />
        <Route path="create-category" element={<CreateCategory />} />
        <Route path="create-todo" element={<CreateTodo />} />
        <Route path="recurring-todos" element={<RecurringTodosList />} />
        <Route path="suggested-todos" element={<SuggestedTodosList />} />
        <Route path="completed-todos" element={<CompletedTodosList />} />
        <Route path="profile" element={<Profile />} />
        <Route path="provision" element={<ProvisionDemoUser />} />
        <Route path="users" element={<UsersList />} />
        <Route path="import" element={<ImportCSV />} />
        <Route path="user/:id" element={<User />} />
      </Routes>
    </BrowserRouter>
  ) : (
    <React.StrictMode>
      <BrowserRouter>
        <Routes>
          <Route path="/" element={<App />} />
          <Route path="login" element={<Login/>}/>
          <Route path="logout" element={<Logout/>}/>
          <Route path="register" element={<Register/>}/>
          <Route path="sync" element={<QueryParamsToLocalStorage />} />
          <Route path="report-bug" element={<ReportBug />} />
          <Route path="categories" element={<CategoriesList />} />
          <Route path="todos" element={<TodosList />} />
          <Route path="category/:id" element={<Category />} />
          <Route path="todo/:id" element={<Todo />} />
          <Route path="create-category" element={<CreateCategory />} />
          <Route path="create-todo" element={<CreateTodo />} />
          <Route path="recurring-todos" element={<RecurringTodosList />} />
          <Route path="suggested-todos" element={<SuggestedTodosList />} />
          <Route path="completed-todos" element={<CompletedTodosList />} />
          <Route path="profile" element={<Profile />} />
          <Route path="provision" element={<ProvisionDemoUser />} />
          <Route path="users" element={<UsersList />} />
          <Route path="import" element={<ImportCSV />} />
          <Route path="user/:id" element={<User />} />
        </Routes>
      </BrowserRouter>
    </React.StrictMode>
  )
);
