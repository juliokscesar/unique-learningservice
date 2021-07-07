import React from "react";
import { Redirect, RouteComponentProps } from "react-router-dom";
import { NavBar } from "../components/NavBar";
import { getUserFromPublicId } from "../apiCommunication";
import { setTitle } from "../utils";

type ProfileState = {
  isLoaded: boolean;
  error: string | null;
  userInfo:
    | {
        name: string;
      }
    | undefined;
};

interface ProfileInfo {
  uid: string;
}

interface ProfileProps extends RouteComponentProps<ProfileInfo> {}

export class Profile extends React.Component<ProfileProps, ProfileState> {
  state: ProfileState = {
    isLoaded: false,
    error: null,
    userInfo: undefined,
  };

  componentDidMount() {
    setTitle("Profile");

    if (this.props.match.params.uid !== undefined) {
      getUserFromPublicId(this.props.match.params.uid).then(
        (result) => {
          if (result["error"] !== undefined) {
            this.setState({
              isLoaded: true,
              error: result["error"],
            });
          } else {
            this.setState({
              isLoaded: true,
              userInfo: {
                name: result["name"],
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
  }

  render() {
    const { isLoaded, error, userInfo } = this.state;

    if (error) {
      if (error.includes("invalid Object ID")) {
        return <Redirect to="/error?err=Invalid+Profile+ID." />;
      } else {
        return <Redirect to={"/error?err=" + error} />;
      }
    } else if (!isLoaded) {
      return (
        <div>
          <h1>Loading the informations...</h1>
        </div>
      );
    } else if (userInfo !== undefined) {
      return (
        <div className="profilePage">
          <NavBar />

          <h1>{userInfo.name}'s Profile</h1>
        </div>
      );
    } else {
      return <Redirect to="/login" />;
    }
  }
}
