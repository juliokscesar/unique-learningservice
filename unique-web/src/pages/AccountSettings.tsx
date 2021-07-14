import React from "react";
import { Redirect } from "react-router-dom";
import { getUserFromId } from "../apiCommunication";
import { NavBar } from "../components/NavBar";
import { cookies } from "../index";

type UserInfo = {
  id: string;
  name: string;
  email: string;
};

type AccountSettingsState = {
  isLoaded: boolean;
  error: string | null;
  user: UserInfo | undefined;
};

export class AccountSettings extends React.Component<{}, AccountSettingsState> {
  state: AccountSettingsState = {
    isLoaded: false,
    error: null,
    user: undefined,
  };

  componentDidMount() {
    const uid = cookies.get("luid");
    if (uid === undefined) {
      document.location.href = "/login";
    }

    getUserFromId(uid).then(
      (result) => {
        if (result["error"] !== undefined) {
          this.setState({
            isLoaded: true,
            error: result["error"],
          });
        } else {
          this.setState({
            isLoaded: true,
            user: {
              id: uid,
              name: result["name"],
              email: result["email"],
            },
          });
        }
      },
      (err) => {
        this.setState({
          isLoaded: true,
          error: err,
        });
      }
    );
  }

  render() {
    const { isLoaded, error, user } = this.state;

    if (!isLoaded) {
      return <h1>Loading your informations...</h1>;
    } else if (error !== null) {
      return <Redirect to={"/error?err=" + error} />;
    } else if (user !== undefined) {
      return (
        <div className="accountSettings">
          <NavBar />
          <h1>{user.name}'s Account Settings</h1>

          <p>
            Name: {user.name} <br />
            Email: {user.email} <br />
            Password: <a href="/accountSettings/changePass">Change Password</a>
            <br />
          </p>
        </div>
      );
    } else {
      return <Redirect to="/login" />;
    }
  }
}
