import React from "react";
import { cookies } from "..";
import { changeUserPassword } from "../apiCommunication";
import { NavBar } from "../components/NavBar";
import { BasicForm } from "../components/BasicForm";

type ChangePassState = {
  uid: string;
  oldPass: string;
  newPass: string;
};

export class ChangePasswordForm extends React.Component<{}, ChangePassState> {
  state: ChangePassState = {
    uid: "",
    oldPass: "",
    newPass: "",
  };

  changeOldPass = (pass: string) => {
    this.setState({
      oldPass: pass,
    });
  };

  changeNewPass = (pass: string) => {
    this.setState({
      newPass: pass,
    });
  };

  changePassSubmit = async () => {
    return await changeUserPassword(this.state.uid, this.state);
  };

  componentDidMount() {
    this.setState({
      uid: cookies.get("luid"),
    });
  }

  render() {
    const inputs = [
      {
        name: "oldPass",
        type: "password",
        required: true,
        minLength: 8,
        onChangeFn: (e: React.ChangeEvent<HTMLInputElement>) =>
          this.changeOldPass(e.target.value),
      },
      {
        name: "newPass",
        type: "password",
        required: true,
        minLength: 8,
        onChangeFn: (e: React.ChangeEvent<HTMLInputElement>) =>
          this.changeNewPass(e.target.value),
      },
    ];

    return (
      <div className="changePassForm">
        <NavBar />

        <BasicForm
          title="Change Password"
          inputs={inputs}
          submitDisabled={false}
          submitFn={async () => {
            const result = await this.changePassSubmit();

            if (result["error"] !== undefined) {
              document.location.href = "/error?err=" + result["error"];
            } else {
              document.location.href = "/";
            }
          }}
        />
      </div>
    );
  }
}
