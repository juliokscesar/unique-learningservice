import React from "react";
import { Link } from "react-router-dom";
import { UserForm } from "./UserForm";

const testSubmit = (email: string, password: string) => {
  console.log("Email: " + email + " Password: " + password);
}

export const Register = () => {
    return (
        <div className="registerPage">
          <Link to="/login">Login</Link>

          <UserForm title="Register" submitFn={ testSubmit } />
        </div>
    )
}
