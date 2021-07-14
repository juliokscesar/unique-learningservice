import React from "react";
import { Link } from "react-router-dom";
import { BasicForm } from "../components/BasicForm";
import { loginUser } from "../apiCommunication";
import { cookies } from "../index";
import {
  capitalizeFirstLetter,
  toggleElementById,
  validateEmail,
} from "../utils";
import "../style/Login.scss";

type LoginState = {
  email: string;
  password: string;
};

export class Login extends React.Component<{}, LoginState> {
  state: LoginState = {
    email: "",
    password: "",
  };

  changeEmail = (newEmail: string) => {
    this.setState({ email: newEmail });
  };

  changePassword = (newPassword: string) => {
    this.setState({ password: newPassword });
  };

  loginSubmit = async () => {
    return await loginUser(this.state);
  };

  onSuccess = (userId: string, publicId: string) => {
    cookies.set("luid", userId, {
      path: "/",
      expires: new Date(Date.now() + 60 * 60 * 24 * 365 * 10),
      maxAge: 60 * 60 * 24 * 365 * 10,
      sameSite: "lax",
    });
    cookies.set("puid", publicId, {
      path: "/",
      expires: new Date(Date.now() + 60 * 60 * 24 * 365 * 10),
      maxAge: 60 * 60 * 24 * 365 * 10,
      sameSite: "lax",
    });
  };

  render() {
    const inputFields = [
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
      <div className="loginForm">
        <BasicForm
          title="Login"
          inputs={inputFields}
          submitFn={async () => {
            const data = await this.loginSubmit();

            if (data["error"] !== undefined) {
              document.location.href =
                "/error?err=" + capitalizeFirstLetter(String(data["error"]));
            } else {
              this.onSuccess(data["id"], data["public_id"]);
              document.location.href = "/";
            }
          }}
          submitDisabled
        />

        <br />

        <p>
          Not registered yet? Register <Link to="/register">here.</Link>
        </p>
      </div>
    );
  }
}
