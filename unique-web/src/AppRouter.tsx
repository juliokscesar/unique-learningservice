import React from "react";
import { BrowserRouter, Switch, Route } from "react-router-dom"; 
import { Toast } from "./components/Toast";
import { UserHome } from "./components/UserHome";
import { Login } from "./components/Login";
import { Register } from "./components/Register";

const AppRouter = () => {
  return (
    <div className="App">
      <BrowserRouter>
        <Switch>

          <Route exact path="/">
            <UserHome />
          </Route>

          <Route path="/login">
            <Login />
          </Route>

          <Route path="/register">
            <Register />
          </Route>

          <Route path="/secret">
            <Toast message="Eureka! You have just discovered my secret!" author="Julio"/>
          </Route>

        </Switch>
      </BrowserRouter>
    </div>
  );
}

export default AppRouter;
