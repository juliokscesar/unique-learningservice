import React from "react";
import Cookies from "universal-cookie/es6";
import ReactDOM from "react-dom";
import AppRouter from "./AppRouter";

export const cookies = new Cookies();

ReactDOM.render(
  <AppRouter />,
  document.getElementById("root")
);
