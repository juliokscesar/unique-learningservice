import React from "react";
import { Redirect } from "react-router-dom";
import { cookies } from "../index";
import { setTitle } from "../utils";

export class Logout extends React.Component {
  render() {
    setTitle("Logout");

    return <Redirect to="/login" />;
  }

  componentWillUnmount() {
    if (cookies.get("luid") !== undefined) {
      cookies.remove("luid");
    }
    if (cookies.get("puid") !== undefined) {
      cookies.remove("puid");
    }
  }
}
