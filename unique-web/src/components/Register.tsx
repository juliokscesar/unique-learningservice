import React from "react";
import { Link } from "react-router-dom";
import { UserForm } from "./UserForm";

export const Register = () => {
    return (
        <div className="registerPage">
          <Link to="/login">Login</Link>

          <UserForm title="Register" />
        </div>
    )
}
