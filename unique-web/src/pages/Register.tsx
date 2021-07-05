import React from "react";
// import * as $ from "jquery";
import { Link } from "react-router-dom";
import { toggleElementById, validateEmail } from "../utils";
import { UserForm } from "../components/UserForm";
import { API_BASE_URI } from "../constants";
import QueryString from "qs";

type RegisterState = {
  name: string;
  email: string;
  password: string;
  resultURL: string;
};

export class Register extends React.Component<{}, RegisterState> {
  state: RegisterState = {
    name: "",
    email: "",
    password: "",
    resultURL: "/",
  };

  changeName = (newName: string) => {
    this.setState({ name: newName });
  };

  changeEmail = (newEmail: string) => {
    this.setState({ email: newEmail });
  };

  changePassword = (newPassword: string) => {
    this.setState({ password: newPassword });
  };

  changeResultURL = (newURL: string) => {
    this.setState({ resultURL: newURL });
  };

  registerSubmit = async () => {
    let dataResult: string = await fetch(API_BASE_URI + "user/register", {
      method: "POST",
      headers: {
        "Content-Type": "application/x-www-form-urlencoded",
      },
      body: QueryString.stringify(this.state),
    })
      .then((res) => res.json())
      .then((data) => JSON.stringify(data));

    return JSON.parse(dataResult);
  };

  render() {
    const inputFields = [
      {
        name: "name",
        type: "text",
        required: true,
        onChangeFn: (e: React.ChangeEvent<HTMLInputElement>) =>
          this.changeName(e.target.value),
      },
      {
        name: "email",
        type: "text",
        required: true,
        onChangeFn: (e: React.ChangeEvent<HTMLInputElement>) => {
          const pass = (document.querySelector("#password") as HTMLInputElement)
            .value;
          toggleElementById(
            "submit",
            validateEmail(e.target.value) && pass.length >= 8
          );
          this.changeEmail(e.target.value);
        },
      },
      {
        name: "password",
        type: "password",
        required: true,
        minLength: 8,
        onChangeFn: (e: React.ChangeEvent<HTMLInputElement>) => {
          const email = (document.querySelector("#email") as HTMLInputElement)
            .value;
          toggleElementById(
            "submit",
            validateEmail(email) && e.target.value.length >= 8
          );
          this.changePassword(e.target.value);
        },
      },
    ];

    return (
      <div className="registerForm">
        <UserForm
          title="Register"
          inputs={inputFields}
          submitFn={async () => {
            const data = await this.registerSubmit();

            if (data["error"] !== undefined) {
              this.changeResultURL("/error?err=Email+Registered");
            } else {
              this.changeResultURL("/login");
            }

            document.location.href = this.state.resultURL;
          }}
        />

        <br />

        <p>
          Already have an account? Login <Link to="/login">here.</Link>
        </p>
      </div>
    );
  }
}
