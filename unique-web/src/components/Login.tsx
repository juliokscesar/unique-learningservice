import React from "react";
import { Link } from "react-router-dom";
import { UserForm } from "./UserForm";

const testSubmit = (email: string, password: string) => {
  alert("Email: " + email + " Password: " + password);
}

export const Login = () => {
  return (
      <div className="loginPage">
        <Link to="/register">Register</Link>

        <UserForm title="Login" submitFn={ testSubmit } />
      </div>
  )
}
