import React from "react";
import { BrowserRouter, Switch, Route, Redirect } from "react-router-dom";
import { Toast } from "./components/Toast";
import { Error } from "./components/Error";
import { UserHome } from "./components/UserHome";
import { Login } from "./components/Login";
import { Register } from "./components/Register";

class AppRouter extends React.Component {
  render() {
    return (
      <div className="App">
        <BrowserRouter>
          <Switch>
            <Route exact path="/">
              <UserHome />
            </Route>

            <Route path="/error" component={Error} />

            <Route path="/login">
              <Login />
            </Route>

            <Route path="/register">
              <Register />
            </Route>

            <Route path="/secret">
              <Toast
                message="Eureka! You have just discovered my secret!"
                author="Julio"
              />
            </Route>

            <Redirect from="*" to="/error?err=Not+Found" />
          </Switch>
        </BrowserRouter>
      </div>
    );
  }
}

export default AppRouter;
