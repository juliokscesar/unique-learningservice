import React from "react";
import ReactDOM from "react-dom";
import { CookiesProvider } from "react-cookie";
import AppRouter from "./AppRouter";

ReactDOM.render(
  <CookiesProvider>
    <AppRouter />
  </CookiesProvider>,
  document.getElementById("root")
);
