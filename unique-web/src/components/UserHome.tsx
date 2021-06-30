import React from "react";
import { useCookies } from "react-cookie";
import { Redirect } from "react-router-dom";

export const UserHome = () => {
  const [cookies] = useCookies(["user_id"])

  if (cookies["user_id"] !== undefined) {
    return (
      <div>
        <p>Hello, guest!</p>
      </div>
    )
  } else {
      return (
        <Redirect to="/login"/>
      )
  }
}
