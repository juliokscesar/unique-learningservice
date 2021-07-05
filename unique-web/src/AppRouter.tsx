import React from "react";
import { BrowserRouter, Switch, Route, Redirect } from "react-router-dom";
import { Toast } from "./components/Toast";
import { Error } from "./components/Error";
import { UserHome } from "./pages/UserHome";
import { Login } from "./pages/Login";
import { Register } from "./pages/Register";
import { cookies } from "./index";
import { Profile } from "./pages/Profile";
import { Logout } from "./pages/Logout";

class AppRouter extends React.Component {
  render() {
    const luid = cookies.get("luid");
    const puid = cookies.get("puid");
    return (
      <div className="App">
        <BrowserRouter>
          <Switch>
            <Route exact path="/">
              <UserHome />
            </Route>

            <Route path="/error" component={Error} />

            <Route path="/login">
              {luid === undefined ? <Login /> : <Redirect to="/" />}
            </Route>

            <Route path="/register">
              {luid === undefined ? <Register /> : <Redirect to="/" />}
            </Route>

            <Route path="/logout">
              <Logout />
            </Route>

            <Route path="/user/profile/:uid" component={Profile} />
            <Route path="/user/profile">
              <Redirect
                to={luid === undefined ? "/login" : "/user/profile/" + puid}
              />
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
