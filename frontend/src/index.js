import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import 'bootstrap/dist/css/bootstrap.min.css';
import App from './App';
import {BrowserRouter, Route, Routes} from "react-router-dom";
import Login from './components/Login';
import Register from './components/Register';
import CategoriesList from './components/CategoriesList';
import TodosList from './components/TodosList';
import RecurringTodosList from './components/RecurringTodosList';
import Category from "./components/Category";
import CreateCategory from "./components/CreateCategory";
import CreateTodo from "./components/CreateTodo";
import Todo from "./components/Todo";
import CompletedTodosList from "./components/CompletedTodosList";
import Logout from "./components/Logout";
import ReportBug from "./components/ReportBug";
import Profile from "./components/UserConfiguration";


const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
  <React.StrictMode>
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<App />} />
        <Route path="login" element={<Login />} />
        <Route path="logout" element={<Logout />} />
        <Route path="register" element={<Register />} />
        <Route path="recover-password" element={<Register />} />
        <Route path="report-bug" element={<ReportBug />} />
        <Route path="categories" element={<CategoriesList />} />
        <Route path="todos" element={<TodosList />} />
        <Route path="category/:id" element={<Category />} />
        <Route path="todo/:id" element={<Todo />} />
        <Route path="create-category" element={<CreateCategory />} />
        <Route path="create-todo" element={<CreateTodo />} />
        <Route path="recurring-todos" element={<RecurringTodosList />} />
        <Route path="suggested-todos" element={<RecurringTodosList />} />
        <Route path="completed-todos" element={<CompletedTodosList />} />
        <Route path="profile" element={<Profile />} />
        <Route path="statistics" element={<CategoriesList />} />
      </Routes>
    </BrowserRouter>
  </React.StrictMode>
);

