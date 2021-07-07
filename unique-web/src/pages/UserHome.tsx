import React from "react";
import { cookies } from "../index";
import { Redirect } from "react-router-dom";
import { NavBar } from "../components/NavBar";
import { setTitle } from "../utils";
import { getUserFromId } from "../apiCommunication";

type UserHomeState = {
  isLoaded: boolean;
  error: string | null;
  loggedUser:
    | {
        name: string;
        id: string;
      }
    | undefined;
};

export class UserHome extends React.Component {
  state: UserHomeState = {
    isLoaded: false,
    error: null,
    loggedUser: undefined,
  };

  componentDidMount() {
    const userId = cookies.get("luid");
    if (userId === undefined) {
      document.location.href = "/login";
    } else {
      getUserFromId(userId).then((result) => {
        if (result["error"] !== undefined) {
          this.setState({
            isLoaded: true,
            error: result["error"],
          });
        } else {
          this.setState({
            isLoaded: true,
            loggedUser: {
              name: result["name"],
              id: userId,
            },
          });
        }
      });
    }

    setTitle("Home");
  }

  render() {
    const { isLoaded, error, loggedUser } = this.state;

    if (error) {
      return <Redirect to={"/error?err=" + error} />;
    } else if (!isLoaded) {
      return (
        <div>
          <h1>Loading your informations...</h1>
        </div>
      );
    } else if (loggedUser !== undefined) {
      return (
        <div className="userHome">
          <NavBar />

          <h1>Hello, {loggedUser.name}!</h1>
        </div>
      );
    } else {
      return <Redirect to="/login" />;
    }
  }
}
