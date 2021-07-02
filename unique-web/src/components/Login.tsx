import React from "react";
import { Link } from "react-router-dom";
import { UserForm } from "./UserForm";

type LoginState = {
  email: string;
  password: string;
}

export class Login extends React.Component<{}, LoginState> {
  state: LoginState = {
    email: "",
    password: "",
  }

  changeEmail = (newEmail: string) => {
    this.setState({email: newEmail});
  }

  changePassword = (newPassword: string) => {
    this.setState({password: newPassword});
  }
}
