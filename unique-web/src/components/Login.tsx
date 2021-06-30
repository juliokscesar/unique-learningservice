import React from "react";
import { Link } from "react-router-dom";
import { UserForm } from "./UserForm";

export const Login = () => {
    return (
        <div className="loginPage">
          <Link to="/register">Register</Link>

          <UserForm title="Login" />
        </div>
    )
}
